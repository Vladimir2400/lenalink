# LenaLink - Быстрый старт деплоя

Краткое руководство по деплою LenaLink на сервере с использованием Docker.

## 📋 Предварительные требования

- Docker и Docker Compose установлены
- Git (для клонирования репозитория)
- Порты 8080, 15432, 15050 свободны

## 🚀 Деплой за 5 шагов

### 1. Клонировать репозиторий

```bash
git clone <repository-url> lenalink
cd lenalink
```

### 2. Настроить переменные окружения

```bash
# Скопировать пример конфига
cp .env.example .env

# Отредактировать .env (установить пароли, API токены)
nano .env
```

**Важные переменные:**
- `DATABASE_PASSWORD` - пароль для PostgreSQL
- `GARS_BASE_URL` - должен быть `https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata`
- `AVIASALES_TOKEN` - токен от Travelpayouts

### 3. Запустить PostgreSQL

```bash
docker compose up -d postgres

# Дождаться готовности БД (проверить здоровье)
docker compose ps
```

### 4. Применить миграции

```bash
# Применить все миграции через Docker
make migrate-up

# Или напрямую:
docker compose run --rm migrate \
  -path=/migrations \
  -database "postgres://lenalink:password@postgres:5432/lenalink_db?sslmode=disable" \
  up
```

### 5. Запустить сервер и загрузить данные

```bash
# Запустить HTTP сервер
docker compose up -d server

# Загрузить данные из внешних провайдеров (опционально)
docker compose run --rm seed
```

## ✅ Проверка работоспособности

```bash
# Проверить статус всех контейнеров
docker compose ps

# Проверить логи сервера
docker compose logs -f server

# Проверить здоровье API
curl http://localhost:8080/health

# Проверить версию миграций
make migrate-version
```

## 🔄 Обновление на новую версию

### Вариант 1: С сохранением данных

```bash
# 1. Сделать backup БД
docker compose exec postgres pg_dump -U lenalink lenalink_db > backup_$(date +%Y%m%d_%H%M%S).sql

# 2. Остановить сервер
docker compose stop server

# 3. Получить обновления
git pull origin main

# 4. Применить новые миграции
make migrate-up

# 5. Пересобрать и перезапустить сервер
docker compose build server
docker compose up -d server

# 6. Обновить данные из провайдеров
docker compose run --rm seed
```

### Вариант 2: С полным сбросом (для dev/staging)

```bash
# 1. Остановить все контейнеры и удалить volumes
docker compose down -v

# 2. Получить обновления
git pull origin main

# 3. Запустить заново
docker compose up -d postgres
make migrate-up
docker compose up -d server
docker compose run --rm seed
```

## 📊 Управление данными

### Синхронизация данных с провайдерами

```bash
# Синхронизировать все провайдеры
docker compose run --rm seed

# Синхронизировать только GARS (автобусы)
SYNC_PROVIDER=gars docker compose run --rm seed

# Синхронизировать только Aviasales (самолеты)
SYNC_PROVIDER=aviasales docker compose run --rm seed
```

### Проверка загруженных данных

```bash
# Подключиться к БД
docker compose exec postgres psql -U lenalink -d lenalink_db

# Проверить количество остановок
SELECT COUNT(*) FROM stops;

# Проверить количество сегментов по типу транспорта
SELECT transport_type, COUNT(*), AVG(price)::NUMERIC(10,2) as avg_price
FROM segments
WHERE route_id IS NULL
GROUP BY transport_type;

# Топ-10 маршрутов
SELECT
  s1.name || ' → ' || s2.name as route,
  COUNT(*),
  MIN(seg.price)::NUMERIC(10,2) as min_price
FROM segments seg
JOIN stops s1 ON seg.start_stop_id = s1.id
JOIN stops s2 ON seg.end_stop_id = s2.id
WHERE seg.route_id IS NULL
GROUP BY s1.name, s2.name
ORDER BY COUNT(*) DESC
LIMIT 10;
```

## 🔧 Управление миграциями

### Проверить текущую версию

```bash
make migrate-version
```

### Применить pending миграции

```bash
make migrate-up
```

### Откатить последнюю миграцию

```bash
make migrate-down
```

### Исправить "грязное" состояние миграции

```bash
# Если миграция прервалась
make migrate-force VERSION=16
```

Подробная документация: [docs/MIGRATIONS.md](docs/MIGRATIONS.md)

## 🐛 Troubleshooting

### Проблема: Сервер не запускается

```bash
# Проверить логи
docker compose logs server

# Проверить, что БД доступна
docker compose exec postgres pg_isready -U lenalink
```

### Проблема: GARS API возвращает 404

Проверьте, что в `.env` установлен полный URL:
```
GARS_BASE_URL=https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata
```

### Проблема: Aviasales не возвращает данные

- Проверьте токен в `.env`
- API может не иметь данных для некоторых маршрутов
- Проверьте логи: `docker compose logs seed`

### Проблема: Ошибка "route_id violates not-null constraint"

Убедитесь, что применена миграция 15:
```bash
make migrate-version  # Должно быть >= 15
```

## 📚 Дополнительные ресурсы

- [CLAUDE.md](CLAUDE.md) - Общая документация проекта
- [docs/MIGRATIONS.md](docs/MIGRATIONS.md) - Детальное руководство по миграциям
- [docs/DOCKER.md](docs/DOCKER.md) - Docker конфигурация
- [docs/DEPLOYMENT.md](docs/DEPLOYMENT.md) - Production деплой

## 🔗 Полезные команды

```bash
# Показать все доступные команды
make help

# Просмотреть логи всех сервисов
docker compose logs -f

# Остановить все
docker compose down

# Полностью удалить (включая volumes)
docker compose down -v

# Перезапустить конкретный сервис
docker compose restart server

# Пересобрать образ
docker compose build server
```

## 📞 Контакты и поддержка

Если возникли проблемы:
1. Проверьте логи: `docker compose logs`
2. Проверьте документацию в `/docs`
3. Создайте issue в репозитории
