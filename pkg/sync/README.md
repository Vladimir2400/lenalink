# Lena Link Integration Service

Swagger API:
```
http://localhost:8080/api/docs
```

Сервис предназначен для доступа к данным 1С ГАРС через стандартный OData API
и предоставляет:

* HTTP-шлюз с единообразными REST-эндпоинтами для основных сущностей
  (маршруты, остановки, рейсы, расписания, тарифы, квоты мест и т. д.).
* Go-пакет `pkg/gars` с типизированным клиентом и удобными опциями фильтрации,
  сортировки и пагинации, который можно переиспользовать в других сервисах.

## Запуск HTTP-сервиса

```bash
make run # или go run ./cmd/gars-server
```

Переменные окружения:

| Переменная | Описание | Значение по умолчанию |
|------------|----------|------------------------|
| `GARS_BASE_URL` | Базовый URL OData | `https://avibus.gars-ykt.ru:4443/avitest/odata/standard.odata` |
| `GARS_USERNAME` | Логин basic auth | `ХАКАТОН` |
| `GARS_PASSWORD` | Пароль basic auth | `123456` |
| `GARS_TIMEOUT` | Таймаут HTTP-запросов (Go duration) | `30s` |
| `GARS_LISTEN_ADDR` | Адрес HTTP-сервера | `:8080` |

После запуска сервис предоставляет следующие эндпоинты:

```
GET /healthz
GET /api/routes
GET /api/stops
GET /api/route-stops
GET /api/trips
GET /api/trip-schedules
GET /api/trip-schedule-stops
GET /api/trip-schedule-regularity
GET /api/trip-schedule-seat-quotas
GET /api/actual-trips
GET /api/trip-sale-statuses
GET /api/active-fares
GET /api/trip-fares
GET /api/service-prices
GET /api/fees
GET /api/seat-availability
```

### Поддерживаемые параметры запроса

Каждый список поддерживает параметры фильтрации и пагинации, которые напрямую
преобразуются в параметры OData:

| Параметр     | Описание                      | Пример |
|--------------|-------------------------------|--------|
| `select`     | Список полей (`$select`)      | `select=Ref_Key,Description` |
| `expand`     | Связанные сущности (`$expand`)| `expand=Остановки` |
| `filter`     | Условие (`$filter`)           | `filter=startswith(Description,'Ленск')` |
| `order`      | Сортировка (`$orderby`)       | `order=Description asc` |
| `top`        | Ограничение (`$top`)          | `top=20` |
| `skip`       | Пропуск (`$skip`)             | `skip=40` |
| `page`       | Номер страницы (1..n)         | `page=2` |
| `page_size`  | Размер страницы               | `page_size=50` |
| `count`      | Возвращать общее количество   | `count=true` |

`page` и `page_size` автоматически переводятся в `$skip` и `$top`.

### Пример

```bash
curl "http://localhost:8080/api/routes?select=Ref_Key,Description,Расстояние&top=5&order=Description asc"
```

Ответ:

```json
{
  "data": [
    {"Description": "Витим — Ленск ", "Ref_Key": "...", "Расстояние": 211},
    {"Description": "Ленск — Мирный", "Ref_Key": "...", "Расстояние": 208}
  ]
}
```

Если был указан `count=true`, то в корне JSON появится поле `count`.

## Использование Go-пакета

```go
client, _ := gars.NewClient(gars.Config{BaseURL: baseURL, Username: user, Password: pass})
service := gars.NewService(client)

routes, meta, err := service.Routes(ctx,
    gars.WithFilter("startswith(Description,'Ленск')"),
    gars.WithOrderBy("Description asc"),
    gars.WithPagination(1, 20),
    gars.WithCount(true),
)
```

`meta.Count` содержит значение `@odata.count`, если на стороне сервера оно
доступно.

## Тесты

```bash
go test ./...
```
