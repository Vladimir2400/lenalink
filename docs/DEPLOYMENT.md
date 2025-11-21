# LenaLink Production Deployment Guide

Руководство по деплою LenaLink на production VPS с использованием Docker и cron.

## Архитектура

```
┌─────────────────────────────────────────────┐
│              VPS Server                     │
│                                             │
│  ┌────────────────┐    ┌──────────────────┐ │
│  │  Docker Image  │    │  Host Cron       │ │
│  │   lenalink     │    │  (0 */6 * * *)   │ │
│  │                │    └────────┬─────────┘ │
│  │  • server   ◄───────────────┐│           │
│  │  • seed     ◄───────────────┘│           │
│  └────────────────┘             │           │
│         │                       │           │
│         ▼                       ▼           │
│  ┌──────────────┐      ┌──────────────┐    │
│  │  Container   │      │  Container   │    │
│  │   server     │      │   seed       │    │
│  │  (always on) │      │  (on-demand) │    │
│  └──────┬───────┘      └──────┬───────┘    │
│         │                     │             │
│         └──────────┬──────────┘             │
│                    ▼                         │
│             ┌─────────────┐                 │
│             │  PostgreSQL │                 │
│             │  Container  │                 │
│             └─────────────┘                 │
└─────────────────────────────────────────────┘
```

## Особенности архитектуры

✅ **Один Docker образ** - `lenalink` содержит оба бинарника (`server` и `seed`)
✅ **Два режима запуска** - определяется через `command` в docker-compose
✅ **Cron на хосте** - запускает `docker compose run --rm seed` по расписанию
✅ **Отдельные логи** - каждый запуск seed пишет в `/logs/sync.log`
✅ **Ручной запуск** - можно запустить синхронизацию в любой момент

## Предварительные требования

- VPS с Ubuntu 20.04+ / Debian 11+
- Docker 20.10+
- Docker Compose v2+
- 2GB+ RAM
- 10GB+ disk space

## Установка

### 1. Клонирование репозитория

```bash
# На VPS сервере
cd /opt
git clone https://github.com/your-org/lenalink.git
cd lenalink
```

### 2. Настройка окружения

```bash
# Создать .env файл
cp .env.example .env
nano .env
```

Обязательные переменные:

```bash
# Database
DATABASE_HOST=postgres
DATABASE_PORT=5432
DATABASE_USER=lenalink
DATABASE_PASSWORD=CHANGE_ME_IN_PRODUCTION
DATABASE_NAME=lenalink_db

# GARS (АвиБус) API
GARS_BASE_URL=https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata
GARS_USERNAME=ХАКАТОН
GARS_PASSWORD=123456

# Aviasales (опционально)
AVIASALES_TOKEN=your_token_here
AVIASALES_MARKER=your_marker

# Server
SERVER_HOST=0.0.0.0
SERVER_PORT=8080
```

⚠️ **ВАЖНО**: Измените `DATABASE_PASSWORD` на случайный пароль в production!

### 3. Сборка и запуск

```bash
# Собрать образ
docker compose build

# Запустить PostgreSQL и сервер
docker compose up -d

# Проверить логи
docker compose logs -f server

# Проверить здоровье
curl http://localhost:8080/health
```

### 4. Применение миграций

Миграции применяются автоматически при запуске сервера (`cmd/server/main.go`).

Проверить статус:

```bash
docker compose exec postgres psql -U lenalink -d lenalink_db -c "\dt"
```

### 5. Настройка cron для синхронизации

```bash
# Автоматическая настройка (рекомендуется)
./scripts/setup-cron.sh

# Или ручная настройка
crontab -e
# Добавить строку:
0 */6 * * * cd /opt/lenalink && docker compose run --rm seed >> /opt/lenalink/logs/sync.log 2>&1
```

Расписание синхронизации:
- `0 */6 * * *` - каждые 6 часов (00:00, 06:00, 12:00, 18:00)
- `0 */4 * * *` - каждые 4 часа
- `0 2 * * *` - каждый день в 02:00

## Использование

### Управление сервером

```bash
# Запустить все сервисы
docker compose up -d

# Остановить все сервисы
docker compose down

# Перезапустить сервер
docker compose restart server

# Просмотр логов
docker compose logs -f server
docker compose logs -f postgres

# Обновление кода
git pull
docker compose build
docker compose up -d
```

### Управление синхронизацией

```bash
# Запустить синхронизацию вручную
docker compose run --rm seed

# Синхронизировать только один провайдер
docker compose run --rm -e SYNC_PROVIDER=gars seed
docker compose run --rm -e SYNC_PROVIDER=aviasales seed
docker compose run --rm -e SYNC_PROVIDER=rzd seed

# Просмотр логов синхронизации
tail -f logs/sync.log

# Просмотр последних 100 строк
tail -n 100 logs/sync.log
```

### Проверка данных

```bash
# Подключиться к БД
docker compose exec postgres psql -U lenalink -d lenalink_db

# Посмотреть статистику
SELECT stop_type, COUNT(*) as count FROM stops GROUP BY stop_type;
SELECT transport_type, COUNT(*) as count FROM segments GROUP BY transport_type;

# Найти маршруты
SELECT * FROM segments WHERE start_stop_id LIKE '%Москва%' LIMIT 10;
```

## Мониторинг

### Проверка здоровья сервера

```bash
# HTTP health check
curl http://localhost:8080/health

# Из docker-compose
docker compose ps
```

### Метрики синхронизации

```bash
# Время последней синхронизации
ls -lht logs/sync.log

# Количество записей
docker compose exec postgres psql -U lenalink -d lenalink_db <<EOF
SELECT
  'Stops' as entity, COUNT(*) as count FROM stops
UNION ALL
SELECT
  'Segments' as entity, COUNT(*) as count FROM segments;
EOF
```

### Логи cron

```bash
# Системные логи cron
sudo journalctl -u cron -f

# Проверить что cron работает
sudo systemctl status cron

# Список cron заданий
crontab -l
```

## Обслуживание

### Очистка старых данных

Старые сегменты (>7 дней) автоматически удаляются при каждой синхронизации.

Ручная очистка:

```bash
docker compose exec postgres psql -U lenalink -d lenalink_db <<EOF
DELETE FROM segments
WHERE route_id IS NULL
  AND departure_time < NOW() - INTERVAL '7 days';
EOF
```

### Резервное копирование

```bash
# Создать backup
docker compose exec postgres pg_dump -U lenalink lenalink_db | gzip > backup-$(date +%Y%m%d).sql.gz

# Восстановить backup
gunzip < backup-20250122.sql.gz | docker compose exec -T postgres psql -U lenalink -d lenalink_db
```

### Ротация логов

Создать `/etc/logrotate.d/lenalink-sync`:

```
/opt/lenalink/logs/sync.log {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
    create 0640 root root
}
```

## Troubleshooting

### Сервер не запускается

```bash
# Проверить логи
docker compose logs server

# Проверить переменные окружения
docker compose config

# Проверить порты
sudo netstat -tulpn | grep 8080
```

### Синхронизация не работает

```bash
# Проверить cron задачи
crontab -l

# Проверить логи cron
sudo journalctl -u cron -n 50

# Запустить вручную
docker compose run --rm seed

# Проверить подключение к GARS
curl -u "ХАКАТОН:123456" "https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata/Stops?\$top=1"
```

### База данных переполнена

```bash
# Проверить размер БД
docker compose exec postgres psql -U lenalink -d lenalink_db -c "
SELECT pg_size_pretty(pg_database_size('lenalink_db'));"

# Очистить старые сегменты
docker compose exec postgres psql -U lenalink -d lenalink_db -c "
DELETE FROM segments WHERE departure_time < NOW() - INTERVAL '30 days';"

# VACUUM для освобождения места
docker compose exec postgres psql -U lenalink -d lenalink_db -c "VACUUM FULL;"
```

### Контейнер seed завис

```bash
# Найти процесс
docker ps -a | grep seed

# Убить контейнер
docker kill lenalink_seed
docker rm lenalink_seed
```

## Безопасность

### Firewall

```bash
# Разрешить только необходимые порты
sudo ufw allow 22/tcp    # SSH
sudo ufw allow 8080/tcp  # HTTP API
sudo ufw enable
```

### SSL/TLS

Рекомендуется использовать Nginx как reverse proxy с Let's Encrypt:

```bash
# Установить Nginx
sudo apt install nginx certbot python3-certbot-nginx

# Настроить proxy
sudo nano /etc/nginx/sites-available/lenalink

# Получить сертификат
sudo certbot --nginx -d api.lenalink.ru
```

### Обновление зависимостей

```bash
# В директории проекта
go get -u ./...
go mod tidy

# Пересобрать образ
docker compose build --no-cache
docker compose up -d
```

## CI/CD (опционально)

GitHub Actions для автоматического деплоя:

```yaml
# .github/workflows/deploy.yml
name: Deploy to Production

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Deploy to VPS
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.VPS_HOST }}
          username: ${{ secrets.VPS_USER }}
          key: ${{ secrets.SSH_KEY }}
          script: |
            cd /opt/lenalink
            git pull
            docker compose build
            docker compose up -d
```

## Полезные команды

```bash
# Статус всех сервисов
docker compose ps

# Использование ресурсов
docker stats

# Проверка размера образов
docker images | grep lenalink

# Очистка неиспользуемых образов
docker image prune -a

# Экспорт переменных окружения
set -a && source .env && set +a

# Проверка конфигурации docker-compose
docker compose config

# Проверка синтаксиса Dockerfile
docker build --check .
```

## Контакты

Для вопросов и поддержки:
- GitHub Issues: https://github.com/your-org/lenalink/issues
- Email: support@lenalink.ru
