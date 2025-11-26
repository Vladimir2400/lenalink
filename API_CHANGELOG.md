# LenaLink API Changelog - Authentication & Booking Updates

## Дата обновления: 2025-11-27

Этот документ описывает все изменения в API для интеграции фронтенда.

---

## 🔐 Новая функциональность: Аутентификация

### 1. Регистрация пользователя

**POST** `/api/register`

**Request:**
```json
{
  "name": "Иван Иванов",
  "email": "ivan@example.com",
  "password": "password123"
}
```

**Response (201 Created):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Iml2YW5AZXhhbXBsZS5jb20iLCJleHAiOjE3MzI4MjQwMDAsImlhdCI6MTczMjczNzYwMCwibmFtZSI6ItCY0LLQsNC9INCY0LLQsNC90L7QsiIsInVzZXJfaWQiOiJhYmMxMjMifQ.xyz",
  "user": {
    "id": "abc123",
    "name": "Иван Иванов",
    "email": "ivan@example.com",
    "created_at": "2025-11-27T10:00:00Z",
    "last_login_at": null
  }
}
```

**Валидация:**
- `name` - обязательное, не пустое
- `email` - обязательное, валидный email формат
- `password` - обязательное, минимум 6 символов

**Errors:**
- `400 VALIDATION_ERROR` - неверные данные
- `400 REGISTRATION_FAILED` - email уже зарегистрирован

---

### 2. Вход пользователя

**POST** `/api/login`

**Request:**
```json
{
  "email": "ivan@example.com",
  "password": "password123"
}
```

**Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "abc123",
    "name": "Иван Иванов",
    "email": "ivan@example.com",
    "created_at": "2025-11-27T10:00:00Z",
    "last_login_at": "2025-11-27T12:30:00Z"
  }
}
```

**Errors:**
- `401 LOGIN_FAILED` - неверный email или пароль

---

### 3. Использование токена

JWT токен действителен **24 часа** с момента выдачи.

Для всех защищенных эндпоинтов добавляйте заголовок:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Пример:**
```javascript
fetch('http://localhost:8080/api/my_routes', {
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  }
})
```

**Errors при отсутствии/невалидном токене:**
- `401 UNAUTHORIZED` - "Authorization header required"
- `401 UNAUTHORIZED` - "Invalid Authorization header format. Use: Bearer <token>"
- `401 UNAUTHORIZED` - "Invalid or expired token"

---

## 📋 Обновленные эндпоинты бронирования

### 4. Создание бронирования (теперь требует авторизации)

**POST** `/api/v1/bookings`

⚠️ **Теперь требует авторизацию!** Добавьте заголовок `Authorization: Bearer <token>`

**Request:** (без изменений)
```json
{
  "route_id": "route_msk_yks_opt",
  "passenger": {
    "first_name": "Иван",
    "last_name": "Иванов",
    "middle_name": "Петрович",
    "date_of_birth": "1990-01-01",
    "passport_number": "1234 567890",
    "email": "ivan@example.com",
    "phone": "+79991234567"
  },
  "include_insurance": false,
  "payment_method": "card"
}
```

**Response (201 Created):**
```json
{
  "id": "booking123",
  "user_id": "abc123",
  "route_id": "route_msk_yks_opt",
  "status": "pending",
  "passenger": {
    "first_name": "Иван",
    "last_name": "Иванов",
    "email": "ivan@example.com",
    "phone": "+79991234567"
  },
  "segments": [...],
  "total_price": 32500.00,
  "grand_total": 34125.00,
  "created_at": "2025-11-27T12:00:00Z"
}
```

**Изменения:**
- Добавлено поле `user_id` в ответе (ID пользователя, создавшего бронирование)
- Бронирование автоматически привязывается к авторизованному пользователю

---

### 5. Получение своих бронирований (новый эндпоинт)

**GET** `/api/my_routes`

🆕 **Новый эндпоинт!** Возвращает все бронирования текущего пользователя.

⚠️ **Требует авторизацию!**

**Request:** Нет тела запроса, только заголовок Authorization

**Response (200 OK):**
```json
{
  "total": 2,
  "bookings": [
    {
      "id": "booking123",
      "user_id": "abc123",
      "route_id": "route_msk_yks_opt",
      "status": "in_progress",
      "passenger": {
        "first_name": "Иван",
        "last_name": "Иванов",
        "email": "ivan@example.com",
        "phone": "+79991234567"
      },
      "segments": [
        {
          "id": "seg1",
          "transport_type": "air",
          "provider": "S7 Airlines",
          "from": {
            "id": "moscow_dme",
            "name": "Domodedovo International Airport",
            "city": "Moscow",
            "latitude": 55.4088,
            "longitude": 37.9063
          },
          "to": {
            "id": "yakutsk_yks",
            "name": "Yakutsk Tolmachevo Airport",
            "city": "Yakutsk",
            "latitude": 62.0932,
            "longitude": 129.7708
          },
          "departure_time": "2025-11-20T08:00:00Z",
          "arrival_time": "2025-11-20T16:30:00Z",
          "duration": "8h 30m",
          "price": 30000.00,
          "distance": 4800
        }
      ],
      "total_price": 32500.00,
      "total_commission": 1625.00,
      "grand_total": 34125.00,
      "insurance_premium": 0,
      "include_insurance": false,
      "created_at": "2025-11-27T10:00:00Z",
      "confirmed_at": "2025-11-27T10:05:00Z",
      "cancelled_at": null,
      "cancellation_reason": null,
      "payment": {
        "id": "pay123",
        "amount": 34125.00,
        "currency": "RUB",
        "method": "card",
        "status": "completed"
      }
    },
    {
      "id": "booking456",
      "user_id": "abc123",
      "route_id": "route_msk_yks_cheap",
      "status": "cancelled",
      "cancellation_reason": "Изменились планы",
      "created_at": "2025-11-26T15:00:00Z",
      "cancelled_at": "2025-11-26T16:00:00Z",
      ...
    }
  ]
}
```

**Сортировка:** По дате создания (новые первыми)

---

### 6. Получение конкретного бронирования (обновлено)

**GET** `/api/v1/bookings/{id}`

⚠️ **Теперь требует авторизацию!**

Без изменений в структуре ответа, но теперь требуется токен.

---

### 7. Отмена бронирования (обновлено)

**POST** `/api/v1/bookings/{id}/cancel`

⚠️ **Теперь требует авторизацию!**

**Request:**
```json
{
  "reason": "Изменились планы"
}
```

**Response (200 OK):**
```json
{
  "id": "booking123",
  "status": "cancelled",
  "cancellation_reason": "Изменились планы",
  "cancelled_at": "2025-11-27T14:00:00Z",
  ...
}
```

---

## 📊 Обновленные статусы бронирований

### Старые статусы (удалены):
- ~~`confirmed`~~
- ~~`failed`~~
- ~~`refunded`~~

### Новые статусы:

| Статус | Описание | Когда используется |
|--------|----------|-------------------|
| `pending` | В ожидании | Ожидание оплаты |
| `in_progress` | В процессе | Подтверждено, поездка активна или предстоит |
| `completed` | Завершено | Поездка завершена |
| `cancelled` | Отменено | Пользователь отменил или произошла ошибка |

**Дополнительные поля для отмененных бронирований:**
- `cancelled_at` - дата и время отмены (ISO 8601)
- `cancellation_reason` - текстовая причина отмены

**Пример отмененного бронирования:**
```json
{
  "id": "booking123",
  "status": "cancelled",
  "cancelled_at": "2025-11-27T14:00:00Z",
  "cancellation_reason": "Не успеваю на рейс",
  ...
}
```

---

## 🆕 Новый эндпоинт: Поиск городов

**GET** `/api/v1/cities?name={prefix}`

🆕 **Новый эндпоинт для автокомплита!**

Поиск городов по началу названия (минимум 2 символа).

**Request:**
```
GET /api/v1/cities?name=Моск
```

**Response (200 OK):**
```json
{
  "cities": [
    {
      "name": "Москва",
      "latitude": 55.7558,
      "longitude": 37.6173
    }
  ]
}
```

**Примеры запросов:**
```
/api/v1/cities?name=Як    -> Якутск
/api/v1/cities?name=Мир   -> Мирный
/api/v1/cities?name=moscow -> (пусто, ищет по русскому названию)
```

**Errors:**
- `400 MISSING_PARAMETER` - параметр `name` обязателен
- `400 INVALID_PARAMETER` - минимум 2 символа

**Лимит:** Максимум 20 результатов

---

## 🔒 Защищенные эндпоинты (требуют Authorization)

| Метод | Путь | Описание |
|-------|------|----------|
| POST | `/api/v1/bookings` | Создание бронирования |
| GET | `/api/v1/bookings` | Список всех бронирований |
| GET | `/api/v1/bookings/{id}` | Получить бронирование |
| POST | `/api/v1/bookings/{id}/cancel` | Отменить бронирование |
| GET | `/api/my_routes` | **Новый!** Мои бронирования |

## 🌐 Открытые эндпоинты (без авторизации)

| Метод | Путь | Описание |
|-------|------|----------|
| POST | `/api/register` | **Новый!** Регистрация |
| POST | `/api/login` | **Новый!** Вход |
| POST | `/api/v1/routes/search` | Поиск маршрутов |
| GET | `/api/v1/routes/{id}` | Получить маршрут |
| GET | `/api/v1/cities` | **Новый!** Поиск городов |
| GET | `/api/v1/health` | Проверка работоспособности |

---

## 💡 Рекомендации для фронтенда

### 1. Хранение токена
```javascript
// После успешной регистрации/входа
const { token, user } = response.data;

// Сохранить в localStorage
localStorage.setItem('auth_token', token);
localStorage.setItem('user', JSON.stringify(user));

// Или в state management (Vuex/Pinia)
store.commit('setAuthToken', token);
store.commit('setUser', user);
```

### 2. Автоматическое добавление токена ко всем запросам
```javascript
// Axios interceptor
axios.interceptors.request.use(config => {
  const token = localStorage.getItem('auth_token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});
```

### 3. Обработка ошибок 401
```javascript
axios.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      // Токен истек или невалиден
      localStorage.removeItem('auth_token');
      localStorage.removeItem('user');
      router.push('/login');
    }
    return Promise.reject(error);
  }
);
```

### 4. Автокомплит городов
```vue
<template>
  <input
    v-model="cityQuery"
    @input="searchCities"
    placeholder="Введите город"
  />
  <ul v-if="cities.length">
    <li v-for="city in cities" :key="city.name" @click="selectCity(city)">
      {{ city.name }}
    </li>
  </ul>
</template>

<script>
export default {
  data() {
    return {
      cityQuery: '',
      cities: []
    }
  },
  methods: {
    async searchCities() {
      if (this.cityQuery.length < 2) {
        this.cities = [];
        return;
      }

      const response = await axios.get(`/api/v1/cities?name=${this.cityQuery}`);
      this.cities = response.data.cities;
    },
    selectCity(city) {
      this.cityQuery = city.name;
      // Сохранить координаты для карты
      this.selectedCity = city;
    }
  }
}
</script>
```

### 5. Отображение статусов бронирований
```javascript
const statusLabels = {
  'pending': 'Ожидает оплаты',
  'in_progress': 'Активно',
  'completed': 'Завершено',
  'cancelled': 'Отменено'
};

const statusColors = {
  'pending': 'orange',
  'in_progress': 'green',
  'completed': 'gray',
  'cancelled': 'red'
};
```

---

## 🧪 Примеры полного flow

### Регистрация и создание бронирования:

```javascript
// 1. Регистрация
const registerResponse = await axios.post('/api/register', {
  name: 'Иван Иванов',
  email: 'ivan@example.com',
  password: 'password123'
});

const { token } = registerResponse.data;
localStorage.setItem('auth_token', token);

// 2. Поиск маршрута (без токена)
const routesResponse = await axios.post('/api/v1/routes/search', {
  from: 'Москва',
  to: 'Якутск',
  departure_date: '2025-11-20',
  passengers: 1
});

const route = routesResponse.data.routes[0];

// 3. Создание бронирования (с токеном)
const bookingResponse = await axios.post('/api/v1/bookings', {
  route_id: route.id,
  passenger: {
    first_name: 'Иван',
    last_name: 'Иванов',
    date_of_birth: '1990-01-01',
    passport_number: '1234567890',
    email: 'ivan@example.com',
    phone: '+79991234567'
  },
  include_insurance: false,
  payment_method: 'card'
}, {
  headers: {
    Authorization: `Bearer ${token}`
  }
});

// 4. Получение своих бронирований
const myRoutesResponse = await axios.get('/api/my_routes', {
  headers: {
    Authorization: `Bearer ${token}`
  }
});

console.log('Мои бронирования:', myRoutesResponse.data.bookings);
```

---

## ⚠️ Breaking Changes

1. **POST /api/v1/bookings** - теперь требует авторизацию
2. **GET /api/v1/bookings** - теперь требует авторизацию
3. **GET /api/v1/bookings/{id}** - теперь требует авторизацию
4. **POST /api/v1/bookings/{id}/cancel** - теперь требует авторизацию
5. Статусы бронирований изменены (см. таблицу выше)
6. В ответах бронирований добавлено поле `user_id`

---

## 📞 Поддержка

При возникновении проблем с интеграцией:
- Проверьте формат токена: `Authorization: Bearer <token>`
- Проверьте срок действия токена (24 часа)
- Убедитесь что используете правильные статусы бронирований

**Дата последнего обновления:** 2025-11-27
