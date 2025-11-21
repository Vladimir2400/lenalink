# Docker Quick Start Guide

Быстрый старт с Docker для локальной разработки и production деплоя.

## Архитектура

Один Docker образ содержит два бинарника:
- **server** - HTTP API сервер (работает постоянно)
- **seed** - синхронизация данных (запускается по требованию)

## Локальная разработка

### 1. Запуск всех сервисов

```bash
# Сборка и запуск
docker compose up -d

# Просмотр логов
docker compose logs -f server

# Проверка здоровья
curl http://localhost:8080/health
```

### 2. Загрузка тестовых данных

```bash
# Способ 1: Через docker-compose
docker compose run --rm seed

# Способ 2: Через make (если Docker не используется)
make seed
```

### 3. Проверка данных

```bash
# Подключиться к PostgreSQL
docker compose exec postgres psql -U lenalink -d lenalink_db

# SQL запросы
SELECT COUNT(*), stop_type FROM stops GROUP BY stop_type;
SELECT COUNT(*), transport_type FROM segments GROUP BY transport_type;
```

## Production деплой

### Структура сервисов

```yaml
services:
  server:          # Основной HTTP сервер
    command: ["server"]
    restart: always
    ports: ["8080:8080"]

  seed:            # Синхронизация (не стартует автоматически)
    command: ["seed"]
    profiles: ["tools"]

  postgres:        # База данных
  pgadmin:         # Веб-интерфейс для БД (опционально)
  redis:           # Кэш (пока не используется)
```

### Команды для production

```bash
# Запустить сервер (без seed)
docker compose up -d server

# Ручной запуск синхронизации
docker compose run --rm seed

# Настроить автоматическую синхронизацию
./scripts/setup-cron.sh
```

### Cron расписание

После запуска `setup-cron.sh` создаётся задание:

```cron
# Каждые 6 часов: 00:00, 06:00, 12:00, 18:00
0 */6 * * * cd /path/to/lenalink && docker compose run --rm seed >> logs/sync.log 2>&1
```

## Полезные команды

### Управление контейнерами

```bash
# Список запущенных контейнеров
docker compose ps

# Остановить все
docker compose down

# Перезапустить сервер
docker compose restart server

# Просмотр логов
docker compose logs -f server
docker compose logs -f seed
```

### Отладка

```bash
# Запустить shell в контейнере
docker compose exec server sh

# Проверить переменные окружения
docker compose exec server env

# Проверить сетевое подключение
docker compose exec server wget -O- http://postgres:5432
```

### Сборка

```bash
# Пересобрать образ
docker compose build

# Пересобрать без кэша
docker compose build --no-cache

# Проверить размер образа
docker images | grep lenalink
```

### Очистка

```bash
# Удалить все контейнеры и volumes
docker compose down -v

# Очистить неиспользуемые образы
docker image prune -a

# Очистить всё (осторожно!)
docker system prune -a --volumes
```

## Переменные окружения

Основные переменные в `.env`:

```bash
# Database
DATABASE_HOST=postgres          # имя сервиса в docker-compose
DATABASE_PORT=5432              # внутренний порт (не 15432!)
DATABASE_USER=lenalink
DATABASE_PASSWORD=password
DATABASE_NAME=lenalink_db

# GARS API
GARS_BASE_URL=https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata
GARS_USERNAME=ХАКАТОН
GARS_PASSWORD=123456

# Aviasales (опционально)
AVIASALES_TOKEN=
AVIASALES_MARKER=

# Server
SERVER_HOST=0.0.0.0
SERVER_PORT=8080
LOG_LEVEL=INFO
```

## Troubleshooting

### Ошибка "port is already allocated"

```bash
# Найти процесс на порту 8080
sudo lsof -i :8080

# Изменить порт в docker-compose.yml
ports:
  - "8081:8080"  # внешний:внутренний
```

### Ошибка "database connection refused"

```bash
# Проверить что PostgreSQL запущен
docker compose ps postgres

# Проверить логи PostgreSQL
docker compose logs postgres

# Проверить имя хоста (должно быть "postgres", не "localhost")
docker compose exec server env | grep DATABASE_HOST
```

### Seed контейнер не запускается

```bash
# Профиль "tools" требует явного запуска
docker compose run --rm seed

# Или временно убрать профиль из docker-compose.yml
```

### Образ слишком большой

```bash
# Проверить размер
docker images lenalink

# Оптимизировать:
# 1. Проверить .dockerignore
# 2. Использовать multi-stage build (уже используется)
# 3. Использовать alpine базовый образ (уже используется)
```

## Мониторинг

### Логи

```bash
# Все логи сервера
docker compose logs server

# Последние 100 строк
docker compose logs --tail=100 server

# Следить за новыми логами
docker compose logs -f server

# Логи синхронизации (если через cron)
tail -f logs/sync.log
```

### Метрики

```bash
# Использование ресурсов
docker stats

# Детали контейнера
docker inspect lenalink_server

# Процессы внутри контейнера
docker top lenalink_server
```

### Health checks

```bash
# HTTP endpoint
curl http://localhost:8080/health

# Docker health status
docker inspect --format='{{.State.Health.Status}}' lenalink_server

# Детальная информация
docker inspect --format='{{json .State.Health}}' lenalink_server | jq
```

## Сравнение: Docker vs Make

| Команда | Make | Docker Compose |
|---------|------|----------------|
| Запуск сервера | `go run ./cmd/server` | `docker compose up -d server` |
| Синхронизация | `make seed` | `docker compose run --rm seed` |
| Подключение к БД | `make psql` | `docker compose exec postgres psql -U lenalink lenalink_db` |
| Миграции | `make migrate-up` | Автоматически при старте |
| Логи | stdout | `docker compose logs -f` |

## Полная инструкция

Смотрите [DEPLOYMENT.md](./DEPLOYMENT.md) для подробной документации по production деплою.

## Ссылки

- [Dockerfile](../Dockerfile)
- [docker-compose.yml](../docker-compose.yml)
- [.dockerignore](../.dockerignore)
- [setup-cron.sh](../scripts/setup-cron.sh)
