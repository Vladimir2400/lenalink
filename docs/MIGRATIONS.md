# Database Migrations Guide

Это руководство описывает, как работать с миграциями базы данных в LenaLink при деплое на сервере с использованием Docker.

## Оглавление

- [Быстрый старт](#быстрый-старт)
- [Команды миграций](#команды-миграций)
- [Создание новых миграций](#создание-новых-миграций)
- [Деплой на сервер](#деплой-на-сервер)
- [Откат миграций](#откат-миграций)
- [Troubleshooting](#troubleshooting)

## Быстрый старт

### 1. Запуск PostgreSQL

```bash
# Запустить PostgreSQL контейнер
docker compose up -d postgres

# Проверить, что контейнер работает
docker compose ps
```

### 2. Применить миграции

```bash
# Применить все pending миграции через Docker
make migrate-up

# Или напрямую через docker compose
docker compose run --rm migrate -path=/migrations -database "postgres://lenalink:password@postgres:5432/lenalink_db?sslmode=disable" up
```

### 3. Проверить версию миграции

```bash
make migrate-version
```

## Команды миграций

### Docker команды (рекомендуется для production)

```bash
# Применить все pending миграции
make migrate-up

# Откатить последнюю миграцию
make migrate-down

# Проверить текущую версию миграции
make migrate-version

# Принудительно установить версию (если миграция "грязная")
make migrate-force VERSION=16
```

### Локальные команды (требуют golang-migrate CLI)

Если у вас установлен [golang-migrate](https://github.com/golang-migrate/migrate) локально:

```bash
# Применить миграции локально
make migrate-up-local

# Откатить миграцию локально
make migrate-down-local

# Проверить версию локально
make migrate-version-local
```

## Создание новых миграций

### 1. Создать пару файлов миграции

```bash
# Создать новую миграцию
make migrate-create NAME=add_user_notifications

# Это создаст два файла:
# migrations/000017_add_user_notifications.up.sql
# migrations/000017_add_user_notifications.down.sql
```

### 2. Отредактировать файлы

**000017_add_user_notifications.up.sql:**
```sql
-- Добавить новую таблицу/колонку
ALTER TABLE users ADD COLUMN notifications_enabled BOOLEAN DEFAULT true;

-- Добавить индекс при необходимости
CREATE INDEX idx_users_notifications ON users(notifications_enabled);
```

**000017_add_user_notifications.down.sql:**
```sql
-- Откатить изменения
DROP INDEX IF EXISTS idx_users_notifications;
ALTER TABLE users DROP COLUMN IF EXISTS notifications_enabled;
```

### 3. Применить новую миграцию

```bash
make migrate-up
```

## Деплой на сервер

### Вариант 1: Через Docker Compose (рекомендуется)

1. **Скопировать файлы на сервер:**

```bash
# На локальной машине
rsync -avz migrations/ user@server:/path/to/lenalink/migrations/
rsync -avz docker-compose.yml user@server:/path/to/lenalink/
```

2. **На сервере:**

```bash
cd /path/to/lenalink

# Запустить PostgreSQL если еще не запущен
docker compose up -d postgres

# Применить миграции
docker compose run --rm migrate \
  -path=/migrations \
  -database "postgres://lenalink:password@postgres:5432/lenalink_db?sslmode=disable" \
  up

# Проверить версию
docker compose run --rm migrate \
  -path=/migrations \
  -database "postgres://lenalink:password@postgres:5432/lenalink_db?sslmode=disable" \
  version
```

### Вариант 2: Через Makefile

Если на сервере установлен `make`:

```bash
cd /path/to/lenalink
make migrate-up
make migrate-version
```

### Вариант 3: CI/CD Pipeline

Добавить в `.github/workflows/deploy.yml` или аналогичный CI/CD конфиг:

```yaml
- name: Run Database Migrations
  run: |
    docker compose run --rm migrate \
      -path=/migrations \
      -database "${{ secrets.DATABASE_URL }}" \
      up
```

## Откат миграций

### Откатить одну миграцию

```bash
# Откатить последнюю миграцию
make migrate-down

# Или через docker compose
docker compose run --rm migrate \
  -path=/migrations \
  -database "postgres://lenalink:password@postgres:5432/lenalink_db?sslmode=disable" \
  down 1
```

### Откатить несколько миграций

```bash
# Откатить последние 3 миграции
docker compose run --rm migrate \
  -path=/migrations \
  -database "postgres://lenalink:password@postgres:5432/lenalink_db?sslmode=disable" \
  down 3
```

### Откатить до конкретной версии

```bash
# Откатить до версии 10
docker compose run --rm migrate \
  -path=/migrations \
  -database "postgres://lenalink:password@postgres:5432/lenalink_db?sslmode=disable" \
  goto 10
```

## Troubleshooting

### Проблема: "Dirty database version"

Если миграция была прервана, база данных может быть в "грязном" состоянии:

```bash
# Проверить статус
make migrate-version

# Вывод: "16 (dirty)"
```

**Решение 1: Принудительно установить версию**

```bash
# Если миграция 16 уже применена
make migrate-force VERSION=16

# Если миграция 16 не была применена
make migrate-force VERSION=15
```

**Решение 2: Вручную исправить в базе**

```bash
# Подключиться к БД
docker compose exec postgres psql -U lenalink -d lenalink_db

# Проверить таблицу миграций
SELECT * FROM schema_migrations;

# Исправить флаг dirty
UPDATE schema_migrations SET dirty = false WHERE version = 16;
```

### Проблема: Контейнер postgres недоступен

```bash
# Проверить статус контейнеров
docker compose ps

# Проверить логи PostgreSQL
docker compose logs postgres

# Перезапустить PostgreSQL
docker compose restart postgres

# Дождаться готовности
docker compose exec postgres pg_isready -U lenalink
```

### Проблема: Миграция не может подключиться к БД

Убедитесь, что используете правильный URL:

- **Внутри Docker сети**: `postgres://lenalink:password@postgres:5432/lenalink_db?sslmode=disable`
- **С хост-машины**: `postgres://lenalink:password@localhost:15432/lenalink_db?sslmode=disable`

### Проблема: Файлы миграций не найдены

```bash
# Проверить, что файлы существуют
ls -la migrations/

# Убедиться, что volume примонтирован
docker compose run --rm migrate ls -la /migrations
```

## Backup перед миграциями

**ВАЖНО:** Всегда делайте backup перед применением миграций на production!

```bash
# Создать backup перед миграцией
docker compose exec postgres pg_dump -U lenalink lenalink_db > backup_$(date +%Y%m%d_%H%M%S).sql

# Применить миграции
make migrate-up

# В случае проблем - восстановить из backup
docker compose exec -T postgres psql -U lenalink -d lenalink_db < backup_20251121_210000.sql
```

## Проверка миграций

### Проверить применены ли все миграции

```bash
# Текущая версия
make migrate-version

# Должно быть: 16 (на момент написания документации)

# Список всех миграций
ls -1 migrations/*.up.sql | wc -l
```

### Проверить состояние таблиц

```bash
# Подключиться к БД
docker compose exec postgres psql -U lenalink -d lenalink_db

# Список всех таблиц
\dt

# Описание таблицы segments
\d+ segments

# Проверить, что route_id и sequence_order nullable
SELECT
  column_name,
  is_nullable,
  data_type
FROM information_schema.columns
WHERE table_name = 'segments'
  AND column_name IN ('route_id', 'sequence_order');
```

## Best Practices

1. **Всегда тестируйте миграции локально** перед деплоем на production
2. **Делайте backup** перед применением миграций на production
3. **Пишите обратные миграции (down)** для возможности отката
4. **Используйте транзакции** в миграциях где возможно
5. **Избегайте изменения данных** в миграциях схемы (делайте отдельные data migrations)
6. **Документируйте сложные миграции** в комментариях SQL
7. **Проверяйте зависимости** между миграциями (foreign keys, indexes)

## Дополнительные ресурсы

- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [PostgreSQL Migration Best Practices](https://www.postgresql.org/docs/current/ddl-alter.html)
- [Zero-downtime Migrations](https://www.braintreepayments.com/blog/safe-operations-for-high-volume-postgresql/)
