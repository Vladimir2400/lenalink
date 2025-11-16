# LenaLink API Documentation

**Base URL:** `http://localhost:8080/api/v1`

**Version:** 0.4.0

**Content-Type:** `application/json`

---

## Overview

LenaLink is a multi-modal transport aggregator that allows users to search for routes and book entire multi-segment journeys (air, rail, bus, river) in a single transaction.

**Key Features:**
- Graph-based route finding with Dijkstra pathfinding
- Multi-segment booking with ACID rollback
- Commission-based pricing (5-15% markup)
- Optional travel insurance
- Mock payment gateway (MVP)

---

## Authentication

🚧 **Not implemented yet** (MVP uses open endpoints)

For production, all endpoints except `/health` will require Bearer token authentication.

---

## Endpoints

### 1. Health Check

**GET** `/health`

Check if the API server is running.

#### Response

```json
{
  "status": "healthy",
  "version": "0.4.0",
  "timestamp": "2025-06-15T10:30:00Z"
}
```

---

### 2. Search Routes

**POST** `/api/v1/routes/search`

Search for available routes between two cities using graph-based pathfinding.

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
| `from` | string | Yes | Departure city (e.g., "moscow", "yakutsk") |
| `to` | string | Yes | Destination city |
| `departure_date` | string | Yes | Departure date in ISO 8601 format (YYYY-MM-DD) |
| `passengers` | integer | No | Number of passengers (default: 1) |

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

- `optimal` - Best balance of price, time, and reliability
- `fastest` - Shortest total duration
- `cheapest` - Lowest total price

#### Status Codes

- `200 OK` - Routes found successfully
- `400 Bad Request` - Invalid search criteria
- `404 Not Found` - No routes available
- `500 Internal Server Error` - Server error

---

### 3. Get Route Details

**GET** `/api/v1/routes/{route_id}`

Get detailed information about a specific route.

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `route_id` | string | Route ID from search results |

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

- `200 OK` - Route found
- `404 Not Found` - Route not found
- `500 Internal Server Error` - Server error

---

### 4. Create Booking

**POST** `/api/v1/bookings`

Book an entire multi-segment route in a single transaction. All segments are booked with providers, and payment is processed. If any segment fails, ALL bookings are rolled back (ACID guarantee).

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
| `route_id` | string | Yes | Route ID from search |
| `passenger.first_name` | string | Yes | Passenger first name |
| `passenger.last_name` | string | Yes | Passenger last name |
| `passenger.middle_name` | string | No | Passenger middle name (отчество) |
| `passenger.date_of_birth` | string | Yes | Date of birth (YYYY-MM-DD) |
| `passenger.passport_number` | string | Yes | Passport number |
| `passenger.email` | string | Yes | Contact email |
| `passenger.phone` | string | Yes | Contact phone |
| `include_insurance` | boolean | No | Include travel insurance (default: false) |
| `payment_method` | string | Yes | Payment method: `card`, `yookassa`, `cloudpay`, `sberpay` |

#### Response

```json
{
  "id": "booking_xyz789",
  "route_id": "route_abc123",
  "status": "confirmed",
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
      "booking_status": "confirmed",
      "provider_booking_ref": "BK-air-xyz78901"
    },
    {
      "id": "booked_seg_002",
      "segment_id": "seg_002",
      "provider": "Lenskie Zori",
      "transport_type": "river",
      "from": {
        "name": "Yakutsk River Port",
        "city": "Yakutsk"
      },
      "to": {
        "name": "Olyokminsk Port",
        "city": "Olyokminsk"
      },
      "departure_time": "2025-06-21T06:00:00Z",
      "arrival_time": "2025-06-21T14:00:00Z",
      "ticket_number": "TKT-Len-def45678",
      "price": 3500.0,
      "commission": 350.0,
      "total_price": 3850.0,
      "booking_status": "confirmed",
      "provider_booking_ref": "BK-river-abc12345"
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

#### Booking Lifecycle

1. `pending` - Booking created, segments being booked
2. `confirmed` - All segments booked, payment successful
3. `failed` - Booking or payment failed, all rolled back
4. `cancelled` - User cancelled booking
5. `refunded` - Refund processed

#### Error Scenarios with ACID Rollback

**Scenario 1: Segment booking fails**
```json
{
  "error": {
    "code": "BOOKING_FAILED",
    "message": "Booking failed at segment 2 (Yakutsk -> Olyokminsk): no available seats",
    "details": "All previous segment bookings have been automatically cancelled"
  }
}
```

**Scenario 2: Payment fails**
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

- `201 Created` - Booking successful
- `400 Bad Request` - Invalid booking data
- `404 Not Found` - Route not found
- `409 Conflict` - Booking failed (segment unavailable, payment failed)
- `500 Internal Server Error` - Server error

---

### 5. Get Booking

**GET** `/api/v1/bookings/{booking_id}`

Retrieve booking details.

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `booking_id` | string | Booking ID (Order ID) |

#### Response

Same format as Create Booking response.

#### Status Codes

- `200 OK` - Booking found
- `404 Not Found` - Booking not found
- `500 Internal Server Error` - Server error

---

### 6. Cancel Booking

**POST** `/api/v1/bookings/{booking_id}/cancel`

Cancel a confirmed booking and process refund.

#### Request Body

```json
{
  "reason": "User requested cancellation"
}
```

#### Response

```json
{
  "id": "booking_xyz789",
  "status": "cancelled",
  "cancelled_at": "2025-06-16T12:00:00Z",
  "cancellation_reason": "User requested cancellation",
  "payment": {
    "status": "refunded"
  }
}
```

#### Status Codes

- `200 OK` - Booking cancelled successfully
- `400 Bad Request` - Cannot cancel booking (already cancelled, etc.)
- `404 Not Found` - Booking not found
- `500 Internal Server Error` - Server error

---

### 7. List Bookings

**GET** `/api/v1/bookings`

Get all bookings (admin endpoint).

#### Query Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `status` | string | Filter by status: `pending`, `confirmed`, `failed`, `cancelled`, `refunded` |
| `email` | string | Filter by passenger email |

#### Response

```json
{
  "bookings": [
    {
      "id": "booking_xyz789",
      "route_id": "route_abc123",
      "status": "confirmed",
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

- `200 OK` - Bookings retrieved
- `500 Internal Server Error` - Server error

---

## Data Models

### TransportType

```
air    - Airplane
rail   - Train
bus    - Bus
river  - River boat/ferry
taxi   - Taxi
walk   - Walking transfer
```

### Commission Rates by Transport Type

```
air:   7%
rail:  5%
bus:   8%
river: 10%
taxi:  15%
walk:  0%
```

### Insurance Calculation

```
Base premium: 5% of route cost

Surcharges:
+ 1% per tight connection (< 2 hours between segments)
+ 0.5% for night flights (departure 22:00-06:00)
+ 2% if route includes river transport
+ 1% if route has 3+ segments
```

---

## Error Response Format

All errors return the following format:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Human-readable error message",
    "details": "Additional details (optional)"
  }
}
```

### Common Error Codes

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `ROUTE_NOT_FOUND` | 404 | Route not found |
| `BOOKING_NOT_FOUND` | 404 | Booking not found |
| `INVALID_ROUTE` | 400 | Invalid route data |
| `INVALID_BOOKING` | 400 | Invalid booking data |
| `BOOKING_FAILED` | 409 | Booking failed (segment unavailable) |
| `PAYMENT_FAILED` | 409 | Payment processing failed |
| `VALIDATION_FAILED` | 400 | Request validation failed |
| `DATABASE_ERROR` | 500 | Database error |

---

## Rate Limiting

🚧 **Not implemented yet** (MVP)

For production:
- 100 requests per minute per IP
- 429 Too Many Requests response when exceeded

---

## Versioning

API version is included in the URL path: `/api/v1/...`

Major version changes will be communicated via:
- `X-API-Version` response header
- Deprecation notices in response headers

---

## Support

**Issues:** https://github.com/lenalink/backend/issues

**Email:** support@lenalink.ru (for production)
