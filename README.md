# LenaLink - Мультимодальный транспортный агрегатор для Якутии

**Агрегатор транспортных билетов с единой транзакцией**: авиа, железная дорога, автобусы и речной транспорт в одном заказе.

Пользователь видит весь маршрут (например, Москва → Якутск → Олёкминск) и **покупает все билеты одним платежом** на нашем сайте. Мы бронируем билеты у поставщиков и добавляем небольшую наценку (5-15%).

**Статус:** MVP для HACK the ICE 2025 с ACID транзакциями и graph-based routing

## 🚀 Быстрый старт

### Требования

- Go 1.22+
- PostgreSQL 14+ (для production, в разработке используется in-memory)
- git (опционально)

### Установка и запуск

```bash
# Клонируйте репо
git clone <repo-url>
cd backend

# Установите зависимости
go mod download

# Запустите сервер
go run ./cmd/server

# Или соберите исполняемый файл
go build -o bin/lenalink-server ./cmd/server
./bin/lenalink-server
```

Сервер будет доступен на `http://localhost:8080`

### Проверка здоровья сервера

```bash
curl http://localhost:8080/health
```

## 🏗️ Архитектура - SOLID принципы

LenaLink построен с использованием **SOLID принципов** для обеспечения maintainability, scalability и testability.

### Слои архитектуры

```
┌──────────────────────────────────────────────────┐
│         HTTP Handlers (API Layer)                │
├──────────────────────────────────────────────────┤
│         Service Layer (Business Logic)           │
│  - RouteService (graph-based pathfinding)        │
│  - BookingService (multi-segment + ACID)         │
│  - CommissionService (markup calculation)        │
│  - InsuranceService (premium calculation)        │
│  - PaymentService (payment processing)           │
├──────────────────────────────────────────────────┤
│         Graph Engine (Route Finding)             │
│  - Graph (nodes + edges)                         │
│  - DijkstraPathfinder (shortest path)            │
│  - GeoJSON Visualization                         │
├──────────────────────────────────────────────────┤
│    Repository Pattern (Data Access)              │
│  - RouteRepository (interface + memory impl)     │
│  - BookingRepository (interface + memory impl)   │
│  - Transaction (ACID guarantees)                 │
├──────────────────────────────────────────────────┤
│    Infrastructure (Persistence, Logging)         │
│  - In-Memory Storage (MVP)                       │
│  - PostgreSQL (planned for production)           │
│  - Structured Logger                             │
├──────────────────────────────────────────────────┤
│         Domain Layer (Entities)                  │
│  - Route, Segment, Stop, Connection              │
│  - Booking, BookedSegment, Payment, Passenger    │
│  - Node, Edge, Graph                             │
└──────────────────────────────────────────────────┘
```

### SOLID применение

#### 1. Single Responsibility (SRP)
Каждый компонент имеет одну ответственность:
- **Handler** - только HTTP обработка запросов
- **Service** - только бизнес-логика (поиск, оптимизация, расчеты)
- **Repository** - только доступ к данным

#### 2. Open/Closed (OCP)
Архитектура открыта для расширения:
- Новые транспорты добавляются без изменения существующего кода
- Новые провайдеры добавляются через новые segment типы

#### 3. Liskov Substitution (LSP)
Repository интерфейсы имеют разные реализации:
- **In-Memory** - для разработки и unit тестов
- **PostgreSQL** - для production (в разработке)
- Обе реализации полностью взаимозаменяемы

#### 4. Interface Segregation (ISP)
Отдельные интерфейсы для разных сущностей:
```go
RouteRepository         // только операции с маршрутами
BookingRepository       // только операции с бронированиями
Transaction             // ACID транзакции
TransactionManager      // управление транзакциями
PaymentGateway          // обработка платежей
ProviderBookingService  // бронирование у поставщиков
```

#### 5. Dependency Inversion (DIP)
Services зависят от интерфейсов, а не конкретных реализаций:
```go
// ✅ Правильно: зависит от интерфейса
type BookingService struct {
    repo repository.BookingRepository
}

// ❌ Неправильно: зависит от конкретной реализации
type BookingService struct {
    repo *postgres.BookingRepository
}
```

## 📁 Структура проекта

```
backend/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point с DI контейнером
│
├── internal/                       # Приватный код приложения
│   ├── config/
│   │   └── config.go              # Config management (env vars)
│   │
│   ├── domain/
│   │   ├── route.go               # Route, Segment, Stop entities
│   │   ├── booking.go             # Booking, BookedSegment, Payment, Passenger
│   │   └── errors.go              # Domain errors
│   │
│   ├── graph/                     # 🆕 Graph Engine
│   │   ├── node.go                # Node (города/станции)
│   │   ├── edge.go                # Edge (транспортные соединения)
│   │   ├── graph.go               # Graph structure (thread-safe)
│   │   └── builder.go             # Builder для создания графа из routes
│   │
│   ├── routing/                   # 🆕 Routing Algorithms
│   │   └── dijkstra.go            # Dijkstra pathfinding с priority queue
│   │
│   ├── visualization/             # 🆕 Map Visualization
│   │   └── geojson.go             # GeoJSON generation для карт
│   │
│   ├── repository/
│   │   ├── interface.go           # Repository + Transaction interfaces
│   │   ├── memory/
│   │   │   ├── route_repository.go
│   │   │   └── booking_repository.go
│   │   └── postgres/              # (готово для реализации)
│   │       ├── connection.go
│   │       ├── route_repository.go
│   │       └── transaction.go      # ACID транзакции
│   │
│   ├── service/
│   │   ├── route_service.go       # Поиск маршрутов (graph-based)
│   │   ├── booking_service.go     # 🔄 Multi-segment booking + ACID
│   │   ├── commission_service.go  # 🆕 Расчет наценки (5-15%)
│   │   ├── insurance_service.go   # Расчет страховки
│   │   └── payment_service.go     # 🆕 Payment processing (mock)
│   │
│   ├── handler/
│   │   └── http/                  # (готово для реализации)
│   │       ├── route_handler.go   # POST /api/v1/routes/search
│   │       ├── booking_handler.go # POST /api/v1/bookings
│   │       └── middleware.go      # CORS, логирование
│   │
│   └── infrastructure/
│       ├── database/
│       │   └── postgres.go        # PostgreSQL connection pool
│       └── logger/
│           └── logger.go          # Structured logging
│
├── pkg/                           # Публичные утилиты
│   └── utils/
│       ├── helpers.go             # ID generation, date parsing
│       ├── cache.go               # In-memory cache с TTL
│       └── logger.go              # Логирование
│
├── migrations/                    # SQL миграции
│   ├── 000001_init_schema.up.sql
│   └── 000001_init_schema.down.sql
│
├── configs/
│   └── config.yaml               # Конфигурация приложения
│
├── .env.example                  # Пример переменных окружения
├── go.mod
├── go.sum
└── README.md
```

## 🔄 Основные потоки приложения

### 1. Поиск маршрутов

```
POST /api/v1/routes/search
    ↓
RouteHandler.SearchRoutes()
    ↓
RouteService.SearchRoutes()
    ├─→ Валидация критериев поиска
    ├─→ RouteRepository.FindByCriteria()
    ├─→ Выбор 3 лучших маршрутов (optimal, fast, cheap)
    └─→ Возврат RouteSearchResult

Критерии отбора:
• Optimal  → максимальный reliability score
• Fast     → минимальная общая продолжительность
• Cheap    → минимальная общая стоимость
```

### 2. Бронирование маршрута (Multi-Segment + ACID)

```
POST /api/v1/bookings
    ↓
BookingHandler.CreateBooking()
    ↓
BookingService.CreateBooking()
    ├─→ Получение Route из RouteRepository
    ├─→ Создание Booking с пустыми segments
    │
    ├─→ ДЛЯ КАЖДОГО СЕГМЕНТА (ACID loop):
    │   ├─→ CommissionService.CalculateCommission()  // Наценка 5-15%
    │   ├─→ ProviderBookingService.BookSegment()     // Бронь у поставщика
    │   │   └─→ Если FAILED → ROLLBACK всех предыдущих сегментов!
    │   ├─→ Создание BookedSegment (ticket number, booking ref)
    │   └─→ Добавление в Booking
    │
    ├─→ InsuranceService.CalculatePremium() (опционально)
    │   • Базовая премия: 5% от стоимости
    │   • +1% за каждое tight connection (< 2 часа)
    │   • +0.5% за ночные рейсы
    │   • +2% за речной транспорт (погодные риски)
    │
    ├─→ PaymentService.CreatePayment()
    ├─→ PaymentService.ProcessPayment()
    │   └─→ Если FAILED → ROLLBACK всех бронирований!
    │
    ├─→ Booking.MarkAsConfirmed()
    └─→ BookingRepository.Save()

КЛЮЧЕВАЯ ОСОБЕННОСТЬ:
Если хотя бы ОДИН сегмент или платёж failed →
отменяются ВСЕ бронирования (ACID гарантия)
```

### 3. Dependency Injection в main.go

```go
// Инициализация слоев от нижнего к верхнему
config := config.Load()         // ← Загрузка конфигурации
logger := logger.New()          // ← Logger
routeRepo := memory.New()       // ← Repository layer
routeService := service.New()   // ← Service layer (зависит от repo)
handler := handler.New()        // ← Handler layer (зависит от service)
```

## 🗂️ Configuration

Конфигурация управляется через environment variables:

```bash
# Сервер
SERVER_HOST=localhost
SERVER_PORT=8080
SERVER_READ_TIMEOUT=15s
SERVER_WRITE_TIMEOUT=15s

# База данных
DB_DRIVER=memory              # memory, postgres, sqlite
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=lenalink
DB_SSLMODE=disable

# Логирование
LOG_LEVEL=INFO               # DEBUG, INFO, WARN, ERROR
LOG_JSON_FORMAT=false
```

Для разработки используйте `.env` файл или копируйте из `.env.example`.

## 🗄️ База данных и миграции

### PostgreSQL (Production)

```bash
# Создание базы данных
createdb lenalink

# Запуск миграций
migrate -path migrations -database "postgres://user:pass@localhost/lenalink" up
```

### Миграции

Миграции находятся в папке `migrations/`:

```
migrations/
├── 000001_init_schema.up.sql       # Создание таблиц
└── 000001_init_schema.down.sql     # Откат
```

### Структура таблиц

```sql
-- Routes table
CREATE TABLE routes (
    id UUID PRIMARY KEY,
    from_city VARCHAR(100),
    to_city VARCHAR(100),
    total_price DECIMAL(10, 2),
    reliability_score FLOAT,
    created_at TIMESTAMP
);

-- Bookings table
CREATE TABLE bookings (
    id UUID PRIMARY KEY,
    route_id UUID REFERENCES routes,
    passenger_count INT,
    total_price DECIMAL(10, 2),
    status VARCHAR(50),
    created_at TIMESTAMP
);

-- Tickets table
CREATE TABLE tickets (
    id UUID PRIMARY KEY,
    booking_id UUID REFERENCES bookings,
    segment_id UUID,
    ticket_number VARCHAR(50),
    status VARCHAR(50),
    created_at TIMESTAMP
);
```

## 👨‍💻 Инструкции для разработчиков

### Запуск сервера

```bash
# Development (with in-memory storage)
DB_DRIVER=memory go run ./cmd/server

# Production (with PostgreSQL)
DB_DRIVER=postgres \
DB_HOST=localhost \
DB_PORT=5432 \
DB_USER=postgres \
DB_NAME=lenalink \
go run ./cmd/server
```

### Сборка

```bash
# Debug build
go build -o bin/lenalink-server ./cmd/server

# Release build с оптимизацией
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" \
  -o bin/lenalink-server ./cmd/server
```

### Тестирование

```bash
# Все тесты
go test ./...

# С покрытием
go test -cover ./...

# Конкретный пакет
go test ./internal/service

# Конкретный тест
go test -run TestRouteService ./internal/service
```

### Форматирование и анализ кода

```bash
# Форматирование
go fmt ./...

# Линтер (установить предварительно)
golangci-lint run ./...

# Проверка зависимостей
go mod tidy
go mod verify
```

### Работа с миграциями

```bash
# Создание новой миграции
migrate create -ext sql -dir migrations -seq create_users_table

# Применение миграций
migrate -path migrations -database "postgres://..." up

# Откат на одну версию
migrate -path migrations -database "postgres://..." down 1

# Откат всех миграций
migrate -path migrations -database "postgres://..." down
```

## 🧪 Unit Testing

Проект использует in-memory repositories для unit тестов:

```go
// Test setup
routeRepo := memory.NewRouteRepository()
routeService := service.NewRouteService(routeRepo)

// Test
result, err := routeService.SearchRoutes(ctx, criteria)
assert.NoError(t, err)
assert.Len(t, result, 3)
```

## 📊 API Endpoints

### 1. Поиск маршрутов

```bash
POST /api/v1/routes/search
Content-Type: application/json

{
  "from_city": "Москва",
  "to_city": "Олекминск",
  "departure_date": "2025-11-20",
  "passenger_count": 1
}

# Ответ: RouteSearchResult с 3 лучшими маршрутами
```

### 2. Детали маршрута

```bash
GET /api/v1/routes/{route_id}

# Ответ: полная информация о маршруте
```

### 3. Создание бронирования

```bash
POST /api/v1/bookings
Content-Type: application/json

{
  "route_id": "route-123",
  "passengers": [
    {
      "first_name": "Иван",
      "last_name": "Петров",
      "date_of_birth": "1990-01-15",
      "passport_number": "123456789",
      "passport_series": "12 AB"
    }
  ],
  "include_insurance": true
}

# Ответ: подтвержденное бронирование с ticket numbers
```

## 🚨 Обработка ошибок

Все API ответы следуют стандартному формату:

**Успех:**
```json
{
  "success": true,
  "data": {
    "id": "route-123",
    "from_city": "Москва",
    "to_city": "Олекминск",
    ...
  },
  "error": null
}
```

**Ошибка:**
```json
{
  "success": false,
  "data": null,
  "error": {
    "code": "ROUTE_NOT_FOUND",
    "message": "Route not found with ID: route-123",
    "timestamp": "2025-11-15T14:30:00Z"
  }
}
```

## 📦 Зависимости

```
github.com/gorilla/mux v1.8.1    # HTTP router
```

Остальное из стандартной библиотеки Go:
- `context` - Context management
- `crypto/rand` - Secure random generation
- `sync` - Concurrency primitives
- `time` - Date/time operations
- `encoding/json` - JSON marshalling

## 🔒 Безопасность

### Текущие ограничения (нормально для MVP)
- ❌ Нет аутентификации (JWT/OAuth)
- ❌ Нет authorization (RBAC)
- ❌ Нет rate limiting
- ❌ Нет HTTPS в разработке
- ❌ Нет CORS по умолчанию

### Production рекомендации
- ✅ Добавить JWT аутентификацию
- ✅ Использовать HTTPS/TLS
- ✅ Добавить rate limiting
- ✅ Настроить CORS политику
- ✅ Валидировать все входные данные
- ✅ Использовать prepared statements для БД

## 🐛 Отладка

### Проверка сервера

```bash
# Health check
curl http://localhost:8080/health

# Проверить что сервер слушает на порту
lsof -i :8080

# Смотреть логи
tail -f app.log
```

### Debug режим

```bash
# С debug логированием
LOG_LEVEL=DEBUG go run ./cmd/server

# С трассировкой всех запросов
TRACE=1 go run ./cmd/server
```

## 📈 Performance

### Оптимизации в коде

1. **Connection pooling** - PostgreSQL pool размер configurable
2. **Caching** - In-memory кэш с TTL для маршрутов
3. **Goroutines** - Параллельная обработка запросов
4. **Database indexes** - На часто запрашиваемых полях

### Мониторинг

Рекомендуется добавить:
- Prometheus метрики
- OpenTelemetry трейсинг
- Structured logging (JSON format)

## 🚀 Планы развития

- [ ] Реальная интеграция с провайдерами (WEB-API 2.4.1, 1С ГАРС, 1С АвиБус)
- [ ] PostgreSQL persistence
- [ ] Redis кэширование
- [ ] JWT аутентификация
- [ ] Email/SMS уведомления
- [ ] Webhook интеграция
- [ ] GraphQL API
- [ ] WebSocket для live updates
- [ ] Kubernetes поддержка
- [ ] Мониторинг и алертинг

## 📞 Поддержка

Вопросы или проблемы? Создавайте issues в репозитории.

## 📄 Лицензия

Проект создан для хакатона HACK the ICE 2025.

---

**Версия:** 0.2.0 (SOLID архитектура)
**Статус:** В активной разработке
**Последнее обновление:** 2025-11-15
