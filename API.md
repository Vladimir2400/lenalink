# LenaLink API Documentation

**Base URL:** `http://localhost:8080/api/v1`

**Version:** 0.5.0

**Content-Type:** `application/json`

---

## Overview

LenaLink - мультимодальный агрегатор транспорта для Якутии (MVP для HACK the ICE 2025). Позволяет искать и бронировать комбинированные маршруты (самолет, поезд, автобус, речной транспорт) с гарантией ACID транзакций.

**Основные возможности:**
- Граф-алгоритм поиска маршрутов (Dijkstra)
- Мультисегментное бронирование с ACID откатом
- Комиссионное ценообразование (5-15% наценка)
- Опциональное страхование
- Аутентификация пользователей (JWT)
- Mock платежный шлюз / YooKassa

---

## Аутентификация

API использует JWT (JSON Web Token) для аутентификации.

### Защищенные эндпоинты
Требуют заголовок `Authorization: Bearer <token>`:
- `POST /api/v1/bookings` - Создание бронирования
- `GET /api/v1/bookings` - Список всех бронирований
- `GET /api/v1/bookings/{id}` - Детали бронирования
- `POST /api/v1/bookings/{id}/cancel` - Отмена бронирования
- `GET /api/my_routes` - Мои маршруты (бронирования пользователя)

### Открытые эндпоинты
Не требуют аутентификации:
- `POST /api/register` - Регистрация
- `POST /api/login` - Вход
- `GET /api/v1/health` - Проверка здоровья
- `POST /api/v1/routes/search` - Поиск маршрутов
- `GET /api/v1/routes/{id}` - Детали маршрута
- `GET /api/v1/cities` - Поиск городов

---

## Endpoints

### 1. Регистрация пользователя

**POST** `/api/register`

Регистрация нового пользователя в системе.

#### Request Body

```json
{
  "name": "Иван Петров",
  "email": "ivan.petrov@example.com",
  "password": "SecurePassword123"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Полное имя пользователя |
| `email` | string | Yes | Email адрес (используется для входа) |
| `password` | string | Yes | Пароль (минимум 8 символов) |

#### Response

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "usr_abc123",
    "name": "Иван Петров",
    "email": "ivan.petrov@example.com",
    "created_at": "2025-06-15T10:30:00Z",
    "last_login_at": null
  }
}
```

#### Status Codes

- `201 Created` - Пользователь успешно зарегистрирован
- `400 Bad Request` - Невалидные данные или email уже используется
- `500 Internal Server Error` - Ошибка сервера

---

### 2. Вход в систему

**POST** `/api/login`

Аутентификация существующего пользователя.

#### Request Body

```json
{
  "email": "ivan.petrov@example.com",
  "password": "SecurePassword123"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `email` | string | Yes | Email адрес |
| `password` | string | Yes | Пароль |

#### Response

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "usr_abc123",
    "name": "Иван Петров",
    "email": "ivan.petrov@example.com",
    "created_at": "2025-06-15T10:30:00Z",
    "last_login_at": "2025-06-16T14:20:00Z"
  }
}
```

#### Status Codes

- `200 OK` - Вход выполнен успешно
- `401 Unauthorized` - Неверный email или пароль
- `400 Bad Request` - Невалидные данные
- `500 Internal Server Error` - Ошибка сервера

---

### 3. Health Check

**GET** `/api/v1/health`

Проверка работоспособности API сервера.

#### Response

```json
{
  "status": "healthy",
  "version": "0.5.0",
  "timestamp": "2025-06-15T10:30:00Z"
}
```

---

### 4. Readiness Check

**GET** `/api/v1/ready`

Проверка готовности сервера к обработке запросов (включая БД).

#### Response

```json
{
  "status": "ready",
  "database": "connected",
  "timestamp": "2025-06-15T10:30:00Z"
}
```

---

### 5. Поиск городов (автокомплит)

**GET** `/api/v1/cities?name={prefix}`

Поиск городов по префиксу названия (для автокомплита в UI).

#### Query Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `name` | string | Yes | Префикс названия города (минимум 2 символа) |

#### Examples

```bash
GET /api/v1/cities?name=як
GET /api/v1/cities?name=мос
GET /api/v1/cities?name=оле
```

#### Response

```json
{
  "cities": [
    {
      "name": "Якутск",
      "latitude": 62.0272,
      "longitude": 129.7322
    },
    {
      "name": "Якутск (аэропорт)",
      "latitude": 62.0932,
      "longitude": 129.7708
    }
  ]
}
```

#### Status Codes

- `200 OK` - Города найдены (может быть пустой массив)
- `400 Bad Request` - Невалидный параметр запроса
- `500 Internal Server Error` - Ошибка сервера

---

### 6. Поиск маршрутов

**POST** `/api/v1/routes/search`

Поиск доступных маршрутов между двумя городами с использованием граф-алгоритма.

#### Request Body

```json
{
  "from": "moscow",
  "to": "olyokminsk",
  "departure_date": "2025-06-20",
  "passengers": 1
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `from` | string | Yes | Город отправления (например, "moscow", "yakutsk") |
| `to` | string | Yes | Город назначения |
| `departure_date` | string | Yes | Дата отправления в формате ISO 8601 (YYYY-MM-DD) |
| `passengers` | integer | No | Количество пассажиров (по умолчанию: 1, максимум: 10) |

#### Response

```json
{
  "routes": [
    {
      "id": "route_abc123",
      "type": "optimal",
      "segments": [
        {
          "id": "seg_001",
          "transport_type": "air",
          "provider": "S7 Airlines",
          "from": {
            "id": "moscow_dme",
            "name": "Domodedovo Airport",
            "city": "Moscow",
            "latitude": 55.4088,
            "longitude": 37.9063
          },
          "to": {
            "id": "yakutsk_yks",
            "name": "Yakutsk Airport",
            "city": "Yakutsk",
            "latitude": 62.0932,
            "longitude": 129.7708
          },
          "departure_time": "2025-06-20T08:00:00Z",
          "arrival_time": "2025-06-20T14:30:00Z",
          "duration": "6h 30m",
          "price": 25000.0,
          "distance": 4884,
          "seat_count": 12,
          "reliability_rate": 0.95
        },
        {
          "id": "seg_002",
          "transport_type": "river",
          "provider": "Lenskie Zori",
          "from": {
            "id": "yakutsk_port",
            "name": "Yakutsk River Port",
            "city": "Yakutsk",
            "latitude": 62.0272,
            "longitude": 129.7322
          },
          "to": {
            "id": "olyokminsk_port",
            "name": "Olyokminsk Port",
            "city": "Olyokminsk",
            "latitude": 60.3733,
            "longitude": 120.4272
          },
          "departure_time": "2025-06-21T06:00:00Z",
          "arrival_time": "2025-06-21T14:00:00Z",
          "duration": "8h",
          "price": 3500.0,
          "distance": 612,
          "seat_count": 8,
          "reliability_rate": 0.85
        }
      ],
      "total_price": 28500.0,
      "total_distance": 5496,
      "total_duration": "30h",
      "reliability_score": 0.90,
      "geojson": {
        "type": "FeatureCollection",
        "features": [
          {
            "type": "Feature",
            "geometry": {
              "type": "LineString",
              "coordinates": [
                [37.9063, 55.4088],
                [129.7708, 62.0932],
                [120.4272, 60.3733]
              ]
            },
            "properties": {
              "type": "route",
              "distance": 5496,
              "price": 28500
            }
          }
        ]
      }
    }
  ],
  "search_criteria": {
    "from": "moscow",
    "to": "olyokminsk",
    "departure_date": "2025-06-20",
    "passengers": 1
  }
}
```

#### Route Types

- `optimal` - Оптимальный баланс цены, времени и надежности
- `fastest` - Самый быстрый маршрут
- `cheapest` - Самый дешевый маршрут

#### Status Codes

- `200 OK` - Маршруты найдены успешно
- `400 Bad Request` - Невалидные критерии поиска
- `404 Not Found` - Маршруты не найдены
- `500 Internal Server Error` - Ошибка сервера

---

### 7. Детали маршрута

**GET** `/api/v1/routes/{route_id}`

Получение детальной информации о конкретном маршруте.

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `route_id` | string | ID маршрута из результатов поиска |

#### Response

```json
{
  "id": "route_abc123",
  "type": "optimal",
  "segments": [...],
  "total_price": 28500.0,
  "commission_breakdown": {
    "base_price": 28500.0,
    "commission": 1995.0,
    "grand_total": 30495.0,
    "segments": [
      {
        "segment_id": "seg_001",
        "transport_type": "air",
        "base_price": 25000.0,
        "commission_rate": 0.07,
        "commission": 1750.0,
        "total": 26750.0
      },
      {
        "segment_id": "seg_002",
        "transport_type": "river",
        "base_price": 3500.0,
        "commission_rate": 0.10,
        "commission": 350.0,
        "total": 3850.0
      }
    ]
  },
  "insurance_available": true,
  "insurance_premium": 1524.75,
  "insurance_breakdown": {
    "base_premium": 1425.0,
    "tight_connection_surcharge": 0.0,
    "night_flight_surcharge": 0.0,
    "river_transport_surcharge": 285.0,
    "total": 1524.75
  }
}
```

#### Status Codes

- `200 OK` - Маршрут найден
- `404 Not Found` - Маршрут не найден
- `500 Internal Server Error` - Ошибка сервера

---

### 8. Создание бронирования

**POST** `/api/v1/bookings`

**⚠️ Требуется аутентификация:** `Authorization: Bearer <token>`

Бронирование всего мультисегментного маршрута в одной транзакции. Все сегменты бронируются у провайдеров, обрабатывается платеж. Если любой сегмент не удается забронировать, ВСЕ бронирования откатываются (ACID гарантия).

#### Request Body

```json
{
  "route_id": "route_abc123",
  "passenger": {
    "first_name": "Иван",
    "last_name": "Петров",
    "middle_name": "Сергеевич",
    "date_of_birth": "1990-05-15",
    "passport_number": "1234 567890",
    "email": "ivan.petrov@example.com",
    "phone": "+79001234567"
  },
  "include_insurance": true,
  "payment_method": "card"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `route_id` | string | Yes | ID маршрута из поиска |
| `passenger.first_name` | string | Yes | Имя пассажира |
| `passenger.last_name` | string | Yes | Фамилия пассажира |
| `passenger.middle_name` | string | No | Отчество пассажира |
| `passenger.date_of_birth` | string | Yes | Дата рождения (YYYY-MM-DD, 18+) |
| `passenger.passport_number` | string | Yes | Номер паспорта (формат: XXXX XXXXXX) |
| `passenger.email` | string | Yes | Email для контакта |
| `passenger.phone` | string | Yes | Телефон (формат: +7XXXXXXXXXX) |
| `include_insurance` | boolean | No | Включить страховку (default: false) |
| `payment_method` | string | Yes | Метод оплаты: `card`, `yookassa`, `cloudpay`, `sberpay` |

#### Response (Success)

```json
{
  "id": "booking_xyz789",
  "route_id": "route_abc123",
  "status": "in_progress",
  "passenger": {
    "first_name": "Иван",
    "last_name": "Петров",
    "email": "ivan.petrov@example.com",
    "phone": "+79001234567"
  },
  "segments": [
    {
      "id": "booked_seg_001",
      "segment_id": "seg_001",
      "provider": "S7 Airlines",
      "transport_type": "air",
      "from": {
        "name": "Domodedovo Airport",
        "city": "Moscow"
      },
      "to": {
        "name": "Yakutsk Airport",
        "city": "Yakutsk"
      },
      "departure_time": "2025-06-20T08:00:00Z",
      "arrival_time": "2025-06-20T14:30:00Z",
      "ticket_number": "TKT-S7A-abc12345",
      "price": 25000.0,
      "commission": 1750.0,
      "total_price": 26750.0,
      "booking_status": "in_progress",
      "provider_booking_ref": "BK-air-xyz78901"
    }
  ],
  "total_price": 28500.0,
  "total_commission": 2100.0,
  "insurance_premium": 1524.75,
  "grand_total": 32124.75,
  "include_insurance": true,
  "payment": {
    "id": "pay_123456",
    "order_id": "booking_xyz789",
    "amount": 32124.75,
    "currency": "RUB",
    "method": "card",
    "status": "completed",
    "provider_payment_id": "MOCK-PAY-abc12345",
    "created_at": "2025-06-15T10:30:00Z",
    "completed_at": "2025-06-15T10:30:05Z"
  },
  "created_at": "2025-06-15T10:30:00Z",
  "confirmed_at": "2025-06-15T10:30:05Z"
}
```

#### Booking Status Lifecycle

1. `pending` - Бронирование создано, сегменты бронируются
2. `pending_payment` - Сегменты забронированы, ожидается подтверждение оплаты
3. `in_progress` - Все сегменты забронированы, оплата успешна (подтверждено)
4. `completed` - Путешествие завершено
5. `failed` - Бронирование или оплата не удались, все откачено
6. `cancelled` - Пользователь отменил бронирование
7. `refunded` - Возврат средств обработан

#### Error Scenarios (ACID Rollback)

**Сценарий 1: Ошибка бронирования сегмента**
```json
{
  "error": {
    "code": "BOOKING_FAILED",
    "message": "Booking failed at segment 2 (Yakutsk -> Olyokminsk): no available seats",
    "details": "All previous segment bookings have been automatically cancelled"
  }
}
```

**Сценарий 2: Ошибка оплаты**
```json
{
  "error": {
    "code": "PAYMENT_FAILED",
    "message": "Payment processing failed: insufficient funds",
    "details": "All segment bookings have been automatically cancelled"
  }
}
```

#### Status Codes

- `201 Created` - Бронирование успешно
- `400 Bad Request` - Невалидные данные бронирования
- `401 Unauthorized` - Отсутствует или невалиден токен
- `404 Not Found` - Маршрут не найден
- `409 Conflict` - Бронирование не удалось (сегмент недоступен, оплата не прошла)
- `500 Internal Server Error` - Ошибка сервера

---

### 9. Получить бронирование

**GET** `/api/v1/bookings/{booking_id}`

**⚠️ Требуется аутентификация:** `Authorization: Bearer <token>`

Получение деталей конкретного бронирования.

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `booking_id` | string | ID бронирования (Order ID) |

#### Response

Тот же формат, что и при создании бронирования.

#### Status Codes

- `200 OK` - Бронирование найдено
- `401 Unauthorized` - Отсутствует или невалиден токен
- `404 Not Found` - Бронирование не найдено
- `500 Internal Server Error` - Ошибка сервера

---

### 10. Список всех бронирований

**GET** `/api/v1/bookings`

**⚠️ Требуется аутентификация:** `Authorization: Bearer <token>`

Получение всех бронирований (admin endpoint).

#### Query Parameters

| Parameter | Type | Description |
|-----------|------|----------|
| `status` | string | Фильтр по статусу: `pending`, `pending_payment`, `in_progress`, `completed`, `failed`, `cancelled`, `refunded` |
| `email` | string | Фильтр по email пассажира |

#### Response

```json
{
  "bookings": [
    {
      "id": "booking_xyz789",
      "route_id": "route_abc123",
      "status": "in_progress",
      "passenger_email": "ivan.petrov@example.com",
      "grand_total": 32124.75,
      "created_at": "2025-06-15T10:30:00Z",
      "confirmed_at": "2025-06-15T10:30:05Z"
    }
  ],
  "total": 1
}
```

#### Status Codes

- `200 OK` - Бронирования получены
- `401 Unauthorized` - Отсутствует или невалиден токен
- `500 Internal Server Error` - Ошибка сервера

---

### 11. Мои маршруты (бронирования пользователя)

**GET** `/api/my_routes`

**⚠️ Требуется аутентификация:** `Authorization: Bearer <token>`

Получение всех бронирований текущего пользователя.

#### Response

```json
{
  "bookings": [
    {
      "id": "booking_xyz789",
      "route_id": "route_abc123",
      "status": "in_progress",
      "passenger": {
        "first_name": "Иван",
        "last_name": "Петров",
        "email": "ivan.petrov@example.com"
      },
      "segments": [...],
      "grand_total": 32124.75,
      "created_at": "2025-06-15T10:30:00Z"
    }
  ],
  "total": 1
}
```

#### Status Codes

- `200 OK` - Бронирования получены
- `401 Unauthorized` - Отсутствует или невалиден токен
- `500 Internal Server Error` - Ошибка сервера

---

### 12. Отмена бронирования

**POST** `/api/v1/bookings/{booking_id}/cancel`

**⚠️ Требуется аутентификация:** `Authorization: Bearer <token>`

Отмена подтвержденного бронирования и обработка возврата средств.

#### Request Body

```json
{
  "reason": "User requested cancellation due to changed plans"
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `reason` | string | Yes | Причина отмены (10-500 символов) |

#### Response

```json
{
  "id": "booking_xyz789",
  "status": "cancelled",
  "cancelled_at": "2025-06-16T12:00:00Z",
  "cancellation_reason": "User requested cancellation due to changed plans",
  "payment": {
    "status": "refunded"
  }
}
```

#### Status Codes

- `200 OK` - Бронирование успешно отменено
- `400 Bad Request` - Невозможно отменить бронирование (уже отменено и т.д.)
- `401 Unauthorized` - Отсутствует или невалиден токен
- `404 Not Found` - Бронирование не найдено
- `500 Internal Server Error` - Ошибка сервера

---

### 13. Webhook YooKassa

**POST** `/api/v1/webhooks/yookassa`

Обработка webhook уведомлений от платежной системы YooKassa.

**⚠️ Этот эндпоинт используется только платежным провайдером, не вызывается напрямую клиентом.**

#### Events

- `payment.succeeded` - Оплата успешно завершена
- `payment.canceled` - Оплата отменена
- `refund.succeeded` - Возврат средств завершен

---

## Модели данных

### TransportType (Типы транспорта)

```
air    - Самолет
rail   - Поезд
bus    - Автобус
river  - Речной транспорт (теплоход/паром)
taxi   - Такси
walk   - Пешая пересадка
```

### Комиссионные ставки по типам транспорта

```
air:   7%
rail:  5%
bus:   8%
river: 10%
taxi:  15%
walk:  0%
```

### Расчет страховки

```
Базовый премиум: 5% от стоимости маршрута

Надбавки:
+ 1% за каждую тесную пересадку (< 2 часов между сегментами)
+ 0.5% за ночные рейсы (отправление 22:00-06:00)
+ 2% если маршрут включает речной транспорт
+ 1% если маршрут имеет 3+ сегмента
```

---

## Формат ошибок

Все ошибки возвращаются в следующем формате:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": "Additional details (optional)"
  }
}
```

### Основные коды ошибок

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `ROUTE_NOT_FOUND` | 404 | Маршрут не найден |
| `BOOKING_NOT_FOUND` | 404 | Бронирование не найдено |
| `INVALID_ROUTE` | 400 | Невалидные данные маршрута |
| `INVALID_BOOKING` | 400 | Невалидные данные бронирования |
| `BOOKING_FAILED` | 409 | Ошибка бронирования (сегмент недоступен) |
| `PAYMENT_FAILED` | 409 | Ошибка обработки платежа |
| `VALIDATION_ERROR` | 400 | Ошибка валидации запроса |
| `UNAUTHORIZED` | 401 | Неавторизованный доступ |
| `REGISTRATION_FAILED` | 400 | Ошибка регистрации |
| `LOGIN_FAILED` | 401 | Ошибка входа |
| `DATABASE_ERROR` | 500 | Ошибка базы данных |
| `MISSING_PARAMETER` | 400 | Отсутствует обязательный параметр |
| `INVALID_PARAMETER` | 400 | Невалидный параметр |

---

## Use Cases (Примеры использования)

### Use Case 1: Полный цикл бронирования

```bash
# 1. Регистрация пользователя
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Иван Петров",
    "email": "ivan@example.com",
    "password": "SecurePass123"
  }'

# Ответ: {"token": "eyJhbGc...", "user": {...}}

# 2. Поиск маршрута
curl -X POST http://localhost:8080/api/v1/routes/search \
  -H "Content-Type: application/json" \
  -d '{
    "from": "moscow",
    "to": "yakutsk",
    "departure_date": "2025-06-20",
    "passengers": 1
  }'

# Ответ: {"routes": [{"id": "route_abc123", ...}]}

# 3. Создание бронирования (с токеном из шага 1)
curl -X POST http://localhost:8080/api/v1/bookings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGc..." \
  -d '{
    "route_id": "route_abc123",
    "passenger": {
      "first_name": "Иван",
      "last_name": "Петров",
      "middle_name": "Сергеевич",
      "date_of_birth": "1990-05-15",
      "passport_number": "1234 567890",
      "email": "ivan@example.com",
      "phone": "+79001234567"
    },
    "include_insurance": true,
    "payment_method": "card"
  }'

# Ответ: {"id": "booking_xyz789", "status": "in_progress", ...}

# 4. Проверка статуса бронирования
curl -X GET http://localhost:8080/api/v1/bookings/booking_xyz789 \
  -H "Authorization: Bearer eyJhbGc..."

# 5. Просмотр моих маршрутов
curl -X GET http://localhost:8080/api/my_routes \
  -H "Authorization: Bearer eyJhbGc..."
```

### Use Case 2: Поиск города для автокомплита

```bash
# Пользователь вводит "як" в поле поиска
curl -X GET "http://localhost:8080/api/v1/cities?name=як"

# Ответ:
# {
#   "cities": [
#     {"name": "Якутск", "latitude": 62.0272, "longitude": 129.7322},
#     {"name": "Якутск (аэропорт)", "latitude": 62.0932, "longitude": 129.7708}
#   ]
# }
```

### Use Case 3: Отмена бронирования

```bash
curl -X POST http://localhost:8080/api/v1/bookings/booking_xyz789/cancel \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGc..." \
  -d '{
    "reason": "Изменились планы поездки, нужно перенести даты"
  }'

# Ответ: {"id": "booking_xyz789", "status": "cancelled", ...}
```

### Use Case 4: Вход существующего пользователя

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "ivan@example.com",
    "password": "SecurePass123"
  }'

# Ответ: {"token": "eyJhbGc...", "user": {...}}
```

---

## Rate Limiting

🚧 **Пока не реализовано** (MVP)

Для продакшена:
- 100 запросов в минуту на IP
- 429 Too Many Requests при превышении

---

## Версионирование

Версия API включена в URL путь: `/api/v1/...`

Изменения мажорной версии будут коммуницироваться через:
- Заголовок ответа `X-API-Version`
- Уведомления о deprecation в заголовках

---

## Техническая поддержка

**Issues:** https://github.com/lenalink/backend/issues

**Email:** support@lenalink.ru (для продакшена)

---

## Changelog

### v0.5.0 (2025-11-27)
- ✅ Добавлена аутентификация пользователей (JWT)
- ✅ Добавлены эндпоинты регистрации и входа
- ✅ Защищены эндпоинты бронирований
- ✅ Добавлен эндпоинт `/api/my_routes`
- ✅ Добавлен эндпоинт поиска городов `/api/v1/cities`
- ✅ Добавлена поддержка YooKassa webhook
- ✅ Исправлены все статусы бронирований
- ✅ Обновлена документация с актуальными use case'ами

### v0.4.0
- ✅ Граф-алгоритм поиска маршрутов
- ✅ Мультисегментное бронирование с ACID
- ✅ Комиссионная модель
- ✅ Опциональное страхование
- ✅ Mock платежный шлюз
