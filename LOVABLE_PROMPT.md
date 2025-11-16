# LenaLink Frontend - Lovable Prompt

## 🎯 Что создаем

Мультимодальный транспортный агрегатор в стиле **Rome2Rio** для России. Поиск маршрутов с комбинацией авиа, ЖД, автобусов, речного транспорта и такси.

**Backend API:** `http://localhost:8080/api/v1`

---

## 📐 Стек

- **Vue 3** (Composition API) + **TypeScript**
- **Vite**
- **Tailwind CSS**
- **Leaflet** + `@vue-leaflet/vue-leaflet` для карты
- **TanStack Query (VueQuery)** для API и кэширования
- **Pinia** для state management
- **Vue Router** для навигации
- **VeeValidate + Yup** для валидации форм
- **Axios** для HTTP запросов
- **vue-toastification** для уведомлений

---

## 🎨 Дизайн и цвета

### Основные цвета:
- Primary Blue: `#1e40af`
- Primary Green: `#059669`
- Text Dark: `#1f2937`
- Text Gray: `#6b7280`
- Background: `#f9fafb`

### Цвета транспорта:
```
✈️  Авиа:      #3b82f6 (синий)
🚂 ЖД:        #ef4444 (красный)
🚌 Автобус:   #f59e0b (оранжевый)
🚢 Речной:    #06b6d4 (бирюзовый)
🚕 Такси:     #eab308 (жёлтый)
🚶 Пешком:    #8b5cf6 (фиолетовый)
```

### Ответственность:
- Мобильный-first подход
- Breakpoints: 320px / 768px / 1024px

---

## 📱 5 Основных страниц

### 1. Главная страница `/`

**Компоненты:**
- `AppHeader` - логотип, навигация
- `SearchForm` - основная форма поиска

**SearchForm требования:**
```vue
<script setup>
// Поля
- from: string (город отправления)
- to: string (город назначения)
- departureDate: Date (дата вылета)
- passengers: number (1-9, по умолчанию 1)

// Функция
- Автокомплит городов (debounce 300ms, минимум 2 символа)
- Кнопка "Swap" для обмена городов
- Валидация: from !== to, date >= сегодня
- VeeValidate + Yup для валидации
- При submit → router.push('/search?from=X&to=Y&date=Z&passengers=N')
</script>
```

**Популярные маршруты:** Показать 4-5 популярных маршрутов (Москва→Якутск, СПб→Мирный и т.д.)

---

### 2. Страница результатов `/search?from=...&to=...&date=...&passengers=...`

**Layout:**
- Десктоп: Карта 40% слева + Список маршрутов 60% справа
- Мобиль: Карта в collapsed режиме (кнопка для раскрытия)

**API запрос:**
```
POST /routes/search
{
  "from": "moscow",
  "to": "olyokminsk",
  "departure_date": "2025-06-20",
  "passengers": 1
}
```

**RouteCard компонент:**
```vue
Props: route (Route объект)
Emits: mouseenter, mouseleave (для highlight на карте)

Показывает:
- Бейдж типа маршрута (Оптимальный/Быстрый/Дешёвый) - разный стиль
- Рейтинг надёжности (⭐ 90%)
- Список сегментов (SegmentItem x N):
  - Иконка транспорта (✈️ 🚂 🚌 🚢 🚕)
  - "Город → Город" (провайдер)
  - Длительность, цена
- Итого строка:
  - 💰 Цена
  - ⏱️  Время в пути
  - 🔄 Кол-во пересадок (segments.length - 1)
- Кнопка "Посмотреть детали →" → route-details/:id
```

**RouteMap компонент (Leaflet):**
```vue
- Отображать первый/выбранный маршрут как GeoJSON
- Разные цвета линий для типов транспорта
- Маркеры для городов (отправление, пересадки, прибытие)
- Popup при клике на маркер
- Legend с расшифровкой цветов
- При hover на RouteCard → highlight маршрута на карте
- Center: [62.0, 129.7], Zoom: 5
```

**Loading/Error:**
- Loading: показать 3 скелетона RouteCard
- Error: "❌ Не удалось найти маршруты" + кнопка вернуться
- Empty: "🔍 Маршруты не найдены"

---

### 3. Страница деталей маршрута `/routes/:id`

**API запрос:**
```
GET /routes/:routeId
Returns: Route (с commission_breakdown, insurance_breakdown)
```

**Содержимое:**
- Заголовок:
  ```
  Москва → Олёкминск
  💰 30,495₽ • ⏱️  30ч • ⭐ 90%
  ```

- Развёрнутые сегменты (SegmentDetailCard x N):
  ```
  ✈️  Авиа: S7 Airlines
  Москва (Домодедово) → Якутск ( Якутск)
  20 июня 2025, 08:00 → 14:30 (6ч 30м)

  Расстояние: 4,884 км
  Мест: 12
  Надёжность: 95%

  Цена: 25,000₽ + комиссия (7%) 1,750₽ = 26,750₽
  ```

- Разбивка цены (PricingBreakdown):
  ```
  Стоимость сегментов:      28,500₽
  Комиссия сервиса:       + 2,100₽
                           ────────
  Итого:                    30,600₽

  ☑️ Добавить страховку (+1,524₽)

  ВСЕГО:                    32,124₽
  ```

- Кнопка "Забронировать маршрут" → /booking/:routeId?insurance=0/1

---

### 4. Страница бронирования `/booking/:routeId`

**3-step форма с прогресс-баром:**

**Шаг 1: Пассажир (PassengerForm)**
```
Имя *                [________]
Фамилия *            [________]
Отчество             [________]
Дата рождения *      [DatePicker]
Паспорт * (1234 567890) [__ ______]
Email *              [_________@___.___]
Телефон * (+79001234567) [+7_____________]

VeeValidate + Yup валидация
Маски для паспорта и телефона
```

**Шаг 2: Опции (InsuranceInfo)**
```
Информация о страховке (readonly, выбрано на предыдущей странице)
- Покрытие: 100,000₽
- Стоимость: +1,524₽
- Включает: отмена, задержка, утеря багажа
```

**Шаг 3: Оплата (PaymentForm)**
```
Способ оплаты:
○ Банковская карта
○ ЮKassa
○ CloudPayments
○ SberPay

OrderSummary (readonly):
Маршрут: Москва → Олёкминск
Пассажир: Иван Петров
Сегменты: 2
Страховка: Да
ИТОГО: 32,124₽
```

**Прогресс-бар вверху:**
```
1 ─── 2 ─── 3
○     ○     ○
```

**API запрос на финал:**
```
POST /bookings
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

**После успеха:** router.push('/booking-confirmation/:bookingId')
**На ошибку:** toast.error() + оставить на шаге 3

---

### 5. Страница подтверждения `/booking-confirmation/:bookingId`

**API запрос:**
```
GET /bookings/:bookingId
```

**Содержимое:**
```
✅ Бронирование подтверждено!

Номер заказа: #booking_xyz789
Статус: Подтверждено

────────────────────────

Данные пассажира:
Иван Петров (Сергеевич)
ivan.petrov@example.com
+79001234567
Дата бронирования: 15 июня 2025, 10:30

────────────────────────

Билеты:
1. TKT-S7A-abc12345 (Авиа, S7 Airlines)
   20 июня 2025, 08:00 → 14:30

2. TKT-Len-def45678 (Речной, Ленские Зори)
   21 июня 2025, 06:00 → 14:00

────────────────────────

Платёж:
Стоимость сегментов: 28,500₽
Комиссия:         + 2,100₽
Страховка:        + 1,524₽
                  ─────────
ОПЛАЧЕНО:          32,124₽

Метод: card
ID платежа: MOCK-PAY-abc12345

────────────────────────

[📧 Отправить на email]
[📱 Скачать PDF]
[🏠 На главную]
```

---

## 🧩 Компоненты (список)

### Common
- `AppHeader.vue` - шапка с логотипом
- `BaseButton.vue` - универсальная кнопка
- `BaseInput.vue` - универсальный инпут
- `LoadingSpinner.vue` - спиннер

### Search
- `SearchForm.vue` - основная форма
- `CityAutocomplete.vue` - автокомплит городов
- `DatePicker.vue` - выбор даты

### Routes
- `RouteCard.vue` - карточка маршрута
- `RouteCardSkeleton.vue` - скелетон загрузки
- `RoutesList.vue` - контейнер списка
- `SegmentItem.vue` - элемент сегмента (компактный)
- `SegmentDetailCard.vue` - развёрнутый сегмент
- `RouteMap.vue` - Leaflet карта

### Booking
- `PassengerForm.vue` - форма пассажира
- `InsuranceInfo.vue` - информация о страховке
- `PaymentMethodSelector.vue` - выбор способа оплаты
- `OrderSummary.vue` - итоговый заказ
- `StepsProgress.vue` - прогресс-бар

---

## 📂 Структура проекта

```
src/
├── components/
│   ├── common/
│   │   ├── AppHeader.vue
│   │   ├── BaseButton.vue
│   │   ├── BaseInput.vue
│   │   └── LoadingSpinner.vue
│   ├── search/
│   │   ├── SearchForm.vue
│   │   ├── CityAutocomplete.vue
│   │   └── DatePicker.vue
│   ├── routes/
│   │   ├── RouteCard.vue
│   │   ├── RouteCardSkeleton.vue
│   │   ├── RoutesList.vue
│   │   ├── SegmentItem.vue
│   │   ├── SegmentDetailCard.vue
│   │   └── RouteMap.vue
│   └── booking/
│       ├── PassengerForm.vue
│       ├── InsuranceInfo.vue
│       ├── PaymentMethodSelector.vue
│       ├── OrderSummary.vue
│       └── StepsProgress.vue
├── pages/
│   ├── HomePage.vue
│   ├── SearchResultsPage.vue
│   ├── RouteDetailsPage.vue
│   ├── BookingPage.vue
│   └── BookingConfirmationPage.vue
├── services/
│   └── api.ts
├── stores/
│   ├── searchStore.ts
│   └── bookingStore.ts
├── composables/
│   ├── useRouteSearch.ts
│   ├── useBooking.ts
│   └── useSearchHistory.ts
├── types/
│   └── index.ts
├── router/
│   └── index.ts
├── utils/
│   ├── formatters.ts
│   ├── constants.ts
│   └── validators.ts
├── assets/
│   └── styles/
│       └── main.css
├── App.vue
└── main.ts
```

---

## 🎯 TypeScript типы

```typescript
// Основные типы для всех компонентов

interface Route {
  id: string
  type: 'optimal' | 'fastest' | 'cheapest'
  segments: Segment[]
  total_price: number
  total_distance: number
  total_duration: string
  reliability_score: number
  commission_breakdown: CommissionBreakdown
  insurance_available: boolean
  insurance_premium: number
  insurance_breakdown: InsuranceBreakdown
  geojson: GeoJSON
}

interface Segment {
  id: string
  transport_type: 'air' | 'rail' | 'bus' | 'river' | 'taxi' | 'walk'
  provider: string
  from: Location
  to: Location
  departure_time: string
  arrival_time: string
  duration: string
  price: number
  distance: number
  seat_count: number
  reliability_rate: number
}

interface Location {
  id: string
  name: string
  city: string
  latitude: number
  longitude: number
}

interface CommissionBreakdown {
  base_price: number
  commission: number
  grand_total: number
  segments: SegmentCommission[]
}

interface SegmentCommission {
  segment_id: string
  transport_type: string
  base_price: number
  commission_rate: number
  commission: number
  total: number
}

interface InsuranceBreakdown {
  base_premium: number
  tight_connection_surcharge: number
  night_flight_surcharge: number
  river_transport_surcharge: number
  total: number
}

interface Passenger {
  first_name: string
  last_name: string
  middle_name?: string
  date_of_birth: string
  passport_number: string
  email: string
  phone: string
}

interface Booking {
  id: string
  route_id: string
  status: 'pending' | 'confirmed' | 'failed' | 'cancelled' | 'refunded'
  passenger: Passenger
  segments: BookedSegment[]
  total_price: number
  total_commission: number
  insurance_premium?: number
  grand_total: number
  include_insurance: boolean
  payment: Payment
  created_at: string
  confirmed_at?: string
}

interface BookedSegment {
  id: string
  segment_id: string
  provider: string
  transport_type: string
  from: { name: string; city: string }
  to: { name: string; city: string }
  departure_time: string
  arrival_time: string
  ticket_number: string
  price: number
  commission: number
  total_price: number
  booking_status: string
  provider_booking_ref: string
}

interface Payment {
  id: string
  order_id: string
  amount: number
  currency: string
  method: 'card' | 'yookassa' | 'cloudpay' | 'sberpay'
  status: string
  provider_payment_id: string
  created_at: string
  completed_at?: string
}

interface SearchParams {
  from: string
  to: string
  departure_date: string
  passengers: number
}
```

---

## 🔌 API Client (services/api.ts)

```typescript
import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  headers: { 'Content-Type': 'application/json' }
})

// Поиск маршрутов
export const searchRoutes = (params: SearchParams) =>
  api.post('/routes/search', params).then(r => r.data)

// Детали маршрута
export const getRouteDetails = (routeId: string) =>
  api.get(`/routes/${routeId}`).then(r => r.data)

// Создание бронирования
export const createBooking = (bookingData: {
  route_id: string
  passenger: Passenger
  include_insurance: boolean
  payment_method: string
}) => api.post('/bookings', bookingData).then(r => r.data)

// Получение бронирования
export const getBooking = (bookingId: string) =>
  api.get(`/bookings/${bookingId}`).then(r => r.data)

export default api
```

---

## 🎯 Маршруты (router/index.ts)

```typescript
[
  { path: '/', name: 'home', component: HomePage },
  { path: '/search', name: 'search', component: SearchResultsPage },
  { path: '/routes/:id', name: 'route-details', component: RouteDetailsPage },
  { path: '/booking/:routeId', name: 'booking', component: BookingPage },
  { path: '/booking-confirmation/:bookingId', name: 'booking-confirmation', component: BookingConfirmationPage }
]
```

---

## 📋 Требования

### Функционал
- ✅ Поиск маршрутов с автокомплитом
- ✅ Карта Leaflet с GeoJSON
- ✅ 3-step форма бронирования
- ✅ Валидация форм (VeeValidate + Yup)
- ✅ Loading states (скелетоны)
- ✅ Error handling (toast уведомления)
- ✅ Мобильная адаптивность
- ✅ TanStack Query для кэширования API
- ✅ Pinia для state management
- ✅ TypeScript для всех типов

### Дизайн
- ✅ Tailwind CSS
- ✅ Mobile-first
- ✅ Разные цвета для типов транспорта
- ✅ Иконки (текстовые эмодзи или простые SVG)

---

## 💡 Дополнительно

- История поисков в localStorage (опционально)
- Маски для паспорта и телефона (openmoji или pattern)
- Smooth transitions между шагами
- Adaptive images (если добавляются)

---

**Версия:** 1.0
**Дата:** 2025-11-15
**Статус:** Ready for Lovable
