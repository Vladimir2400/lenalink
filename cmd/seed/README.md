# LenaLink Data Seeder

Утилита для синхронизации данных из внешних транспортных провайдеров в базу данных LenaLink.

## Описание

Seed команда загружает актуальные данные о маршрутах, остановках и расписаниях из следующих источников:

- **GARS (АвиБус)** - автобусные маршруты Якутии
- **Aviasales** - данные об авиарейсах и аэропортах
- **RZD** - данные о поездах (пока mock)

## Использование

### Синхронизация всех провайдеров

```bash
make seed
```

или

```bash
go run ./cmd/seed
```

### Синхронизация конкретного провайдера

```bash
# Только GARS
make seed-gars

# Только Aviasales
make seed-aviasales

# Только RZD
make seed-rzd
```

### Переменные окружения

#### GARS (обязательно)
```bash
GARS_BASE_URL=https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata
GARS_USERNAME=ХАКАТОН
GARS_PASSWORD=123456
GARS_TIMEOUT=30s
```

#### Aviasales (опционально)
```bash
AVIASALES_TOKEN=your_token_here
AVIASALES_MARKER=your_marker_here
```

#### Aviasales Token

Чтобы получить токен:
1. Зарегистрируйтесь на https://www.travelpayouts.com/
2. Получите API токен в личном кабинете
3. Установите в `.env` файл

#### Фильтрация провайдеров
```bash
# Синхронизировать только один провайдер
SYNC_PROVIDER=gars go run ./cmd/seed
SYNC_PROVIDER=aviasales go run ./cmd/seed
SYNC_PROVIDER=rzd go run ./cmd/seed
```

## Что загружается

### GARS (АвиБус)
- ✅ Автобусные остановки (stops)
- 🚧 Расписание автобусов (в разработке)

### Aviasales
- ✅ Аэропорты России (stops с типом airport)
- ✅ Цены на авиабилеты (segments) для популярных маршрутов:
  - Москва → Санкт-Петербург, Екатеринбург, Красноярск, Иркутск, Якутск
  - Санкт-Петербург → Якутск

### RZD (Mock)
- ✅ Железнодорожные станции (stops с типом station)
- ✅ Расписание поездов (segments)

## Примеры

### Полная синхронизация
```bash
# Запустить PostgreSQL
make docker-up

# Применить миграции
make migrate-up

# Загрузить данные
make seed
```

### Проверка результата
```bash
# Подключиться к БД
make psql

# Посмотреть статистику
SELECT
  stop_type,
  COUNT(*) as count
FROM stops
GROUP BY stop_type;

SELECT
  transport_type,
  COUNT(*) as count
FROM segments
GROUP BY transport_type;
```

### Сброс и повторная загрузка
```bash
# Сбросить БД и загрузить данные заново
make db-reset
make seed
```

## Логирование

Seed выводит подробную информацию о процессе:

```
🌱 LenaLink Data Seeder v0.1.0
Data synchronization tool for LenaLink
========================================

Database: postgres
📦 Connecting to PostgreSQL...
✓ PostgreSQL connected
🔄 Running database migrations...
✓ Migrations completed

⚙️  Loading sync configuration...
GARS BaseURL: https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata
GARS Username: ХАКАТОН
Aviasales Token: abc1****
RZD Enabled: true

🔌 Initializing provider clients...
  ✓ GARS client created
  ✓ Aviasales client created
  ✓ RZD client created

🗄️  Initializing repositories...
✓ Repositories initialized

📊 Checking current data in database...
  Current stops in database: 0
  Current segments in database: 0

🚀 Starting data synchronization...
========================================
Syncing GARS (АвиБус) data...
Fetched 150 stops from GARS
Saved 150 stops from GARS
...
========================================
✅ Synchronization completed in 45s

📈 Final statistics:
  Total stops: 350
  Total segments: 1200

✓ Seeding completed successfully!
```

## Troubleshooting

### Ошибка подключения к GARS
```
Failed to create GARS client: invalid credentials
```
**Решение:** Проверьте GARS_USERNAME и GARS_PASSWORD в `.env`

### Aviasales возвращает пустые данные
```
Fetched 0 flights for MOW-LED
```
**Решение:** Проверьте AVIASALES_TOKEN или используйте только GARS:
```bash
make seed-gars
```

### Таймаут при синхронизации
```
context deadline exceeded
```
**Решение:** Увеличьте таймаут:
```bash
GARS_TIMEOUT=60s make seed
```

## Архитектура

```
cmd/seed/main.go
    ↓
pkg/sync/
    ├── sync.go          (публичный API)
    ├── service.go       (реализация Syncer)
    ├── api/
    │   ├── gars/        (GARS OData client)
    │   ├── aviasales/   (Aviasales REST client)
    │   └── rzd/         (RZD mock client)
    └── internal/
        └── mapper/      (преобразование DTO → domain)
            ↓
internal/repository/postgres/
    ├── stop_repository.go     (сохранение остановок)
    └── segment_repository.go  (сохранение сегментов)
```

## См. также

- [CLAUDE.md](../../CLAUDE.md) - документация проекта
- [API.md](../../API.md) - описание REST API
- [pkg/sync/README.md](../../pkg/sync/README.md) - документация sync пакета
