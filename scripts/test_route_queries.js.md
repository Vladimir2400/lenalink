# Тестовые примеры поиска маршрутов LenaLink (JavaScript/Frontend)

## Содержание
- [Быстрый старт](#быстрый-старт)
- [Конфигурация API](#конфигурация-api)
- [Примеры интеграции](#примеры-интеграции)
- [Тестовые запросы](#тестовые-запросы)

---

## Быстрый старт

### Установка зависимостей (опционально)

```bash
# Если используете axios
npm install axios

# Для TypeScript проектов
npm install -D @types/axios
```

---

## Конфигурация API

### 1. Базовая конфигурация (Vanilla JS + Fetch)

```javascript
// config/api.js
const API_BASE_URL = 'https://lena.linkpc.net/api/v1';

/**
 * Универсальная функция для выполнения POST запросов
 */
async function searchRoutes(searchParams) {
  try {
    const response = await fetch(`${API_BASE_URL}/routes/search`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(searchParams),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Ошибка поиска маршрутов:', error);
    throw error;
  }
}

// Экспорт
export { searchRoutes, API_BASE_URL };
```

### 2. Конфигурация с Axios

```javascript
// config/api.js
import axios from 'axios';

const API_BASE_URL = 'https://lena.linkpc.net/api/v1';

// Создание axios instance с настройками по умолчанию
const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 10000, // 10 секунд
});

// Interceptor для обработки ошибок
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error('API Error:', error.response?.data || error.message);
    return Promise.reject(error);
  }
);

/**
 * Поиск маршрутов
 */
async function searchRoutes(searchParams) {
  try {
    const response = await apiClient.post('/routes/search', searchParams);
    return response.data;
  } catch (error) {
    throw error;
  }
}

export { searchRoutes, apiClient, API_BASE_URL };
```

### 3. TypeScript версия

```typescript
// types/api.ts
export interface SearchRouteRequest {
  from: string;
  to: string;
  departure_date: string; // формат: YYYY-MM-DD
  passengers: number;
}

export interface RouteSegment {
  id: string;
  from: string;
  to: string;
  transport_type: 'air' | 'bus' | 'rail' | 'river' | 'taxi' | 'ferry';
  provider: string;
  departure_time: string;
  arrival_time: string;
  duration: number;
  price: number;
}

export interface Route {
  id: string;
  from_city: string;
  to_city: string;
  departure_time: string;
  arrival_time: string;
  total_duration: number;
  total_price: number;
  transport_types: string[];
  segments: RouteSegment[];
}

export interface SearchRouteResponse {
  routes: Route[];
  total_count: number;
}

// api/routes.ts
import axios, { AxiosInstance } from 'axios';
import type { SearchRouteRequest, SearchRouteResponse } from '../types/api';

const API_BASE_URL = 'https://lena.linkpc.net/api/v1';

class RoutesAPI {
  private client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: API_BASE_URL,
      headers: {
        'Content-Type': 'application/json',
      },
      timeout: 10000,
    });
  }

  async searchRoutes(params: SearchRouteRequest): Promise<SearchRouteResponse> {
    const response = await this.client.post<SearchRouteResponse>(
      '/routes/search',
      params
    );
    return response.data;
  }
}

export const routesAPI = new RoutesAPI();
```

---

## Примеры интеграции

### 1. React Hooks (Functional Component)

```jsx
// hooks/useRouteSearch.js
import { useState } from 'react';
import { searchRoutes } from '../config/api';

export function useRouteSearch() {
  const [routes, setRoutes] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const search = async (params) => {
    setLoading(true);
    setError(null);

    try {
      const data = await searchRoutes(params);
      setRoutes(data.routes || []);
      return data;
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  return { routes, loading, error, search };
}

// components/RouteSearchForm.jsx
import React, { useState } from 'react';
import { useRouteSearch } from '../hooks/useRouteSearch';

export function RouteSearchForm() {
  const { routes, loading, error, search } = useRouteSearch();
  const [formData, setFormData] = useState({
    from: 'Moscow',
    to: 'Yakutsk',
    departure_date: '2025-11-20',
    passengers: 1,
  });

  const handleSubmit = async (e) => {
    e.preventDefault();
    await search(formData);
  };

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          name="from"
          value={formData.from}
          onChange={handleChange}
          placeholder="Откуда"
        />
        <input
          name="to"
          value={formData.to}
          onChange={handleChange}
          placeholder="Куда"
        />
        <input
          name="departure_date"
          type="date"
          value={formData.departure_date}
          onChange={handleChange}
        />
        <input
          name="passengers"
          type="number"
          min="1"
          value={formData.passengers}
          onChange={handleChange}
        />
        <button type="submit" disabled={loading}>
          {loading ? 'Поиск...' : 'Найти маршруты'}
        </button>
      </form>

      {error && <div className="error">{error}</div>}

      {routes.length > 0 && (
        <div className="routes">
          <h3>Найдено маршрутов: {routes.length}</h3>
          {routes.map((route) => (
            <div key={route.id} className="route-card">
              <h4>{route.from_city} → {route.to_city}</h4>
              <p>Цена: {route.total_price}₽</p>
              <p>Длительность: {Math.round(route.total_duration / 3600000000000)}ч</p>
              <p>Транспорт: {route.transport_types.join(', ')}</p>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
```

### 2. Vue 3 Composition API

```vue
<!-- composables/useRouteSearch.js -->
<script setup>
import { ref } from 'vue';
import { searchRoutes } from '../config/api';

export function useRouteSearch() {
  const routes = ref([]);
  const loading = ref(false);
  const error = ref(null);

  const search = async (params) => {
    loading.value = true;
    error.value = null;

    try {
      const data = await searchRoutes(params);
      routes.value = data.routes || [];
      return data;
    } catch (err) {
      error.value = err.message;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  return { routes, loading, error, search };
}
</script>

<!-- components/RouteSearchForm.vue -->
<template>
  <div class="route-search">
    <form @submit.prevent="handleSearch">
      <input
        v-model="searchForm.from"
        placeholder="Откуда"
        required
      />
      <input
        v-model="searchForm.to"
        placeholder="Куда"
        required
      />
      <input
        v-model="searchForm.departure_date"
        type="date"
        required
      />
      <input
        v-model.number="searchForm.passengers"
        type="number"
        min="1"
        required
      />
      <button type="submit" :disabled="loading">
        {{ loading ? 'Поиск...' : 'Найти маршруты' }}
      </button>
    </form>

    <div v-if="error" class="error">{{ error }}</div>

    <div v-if="routes.length > 0" class="routes">
      <h3>Найдено маршрутов: {{ routes.length }}</h3>
      <div
        v-for="route in routes"
        :key="route.id"
        class="route-card"
      >
        <h4>{{ route.from_city }} → {{ route.to_city }}</h4>
        <p>Цена: {{ route.total_price }}₽</p>
        <p>Длительность: {{ formatDuration(route.total_duration) }}ч</p>
        <p>Транспорт: {{ route.transport_types.join(', ') }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import { useRouteSearch } from '../composables/useRouteSearch';

const { routes, loading, error, search } = useRouteSearch();

const searchForm = reactive({
  from: 'Moscow',
  to: 'Yakutsk',
  departure_date: '2025-11-20',
  passengers: 1,
});

const handleSearch = async () => {
  await search(searchForm);
};

const formatDuration = (nanoseconds) => {
  return Math.round(nanoseconds / 3600000000000);
};
</script>
```

### 3. Vanilla JavaScript (без фреймворков)

```html
<!-- index.html -->
<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="UTF-8">
  <title>LenaLink - Поиск маршрутов</title>
  <style>
    .loading { color: blue; }
    .error { color: red; }
    .route-card {
      border: 1px solid #ccc;
      padding: 15px;
      margin: 10px 0;
      border-radius: 8px;
    }
  </style>
</head>
<body>
  <h1>Поиск маршрутов LenaLink</h1>

  <form id="searchForm">
    <input id="from" placeholder="Откуда" value="Moscow" required>
    <input id="to" placeholder="Куда" value="Yakutsk" required>
    <input id="date" type="date" value="2025-11-20" required>
    <input id="passengers" type="number" min="1" value="1" required>
    <button type="submit">Найти маршруты</button>
  </form>

  <div id="status"></div>
  <div id="results"></div>

  <script>
    const API_BASE_URL = 'https://lena.linkpc.net/api/v1';

    async function searchRoutes(params) {
      const response = await fetch(`${API_BASE_URL}/routes/search`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params),
      });

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}`);
      }

      return response.json();
    }

    document.getElementById('searchForm').addEventListener('submit', async (e) => {
      e.preventDefault();

      const statusDiv = document.getElementById('status');
      const resultsDiv = document.getElementById('results');

      statusDiv.innerHTML = '<p class="loading">Поиск маршрутов...</p>';
      resultsDiv.innerHTML = '';

      try {
        const params = {
          from: document.getElementById('from').value,
          to: document.getElementById('to').value,
          departure_date: document.getElementById('date').value,
          passengers: parseInt(document.getElementById('passengers').value),
        };

        const data = await searchRoutes(params);

        statusDiv.innerHTML = `<p>Найдено маршрутов: ${data.routes?.length || 0}</p>`;

        if (data.routes && data.routes.length > 0) {
          resultsDiv.innerHTML = data.routes.map(route => `
            <div class="route-card">
              <h3>${route.from_city} → ${route.to_city}</h3>
              <p><strong>Цена:</strong> ${route.total_price}₽</p>
              <p><strong>Время:</strong> ${Math.round(route.total_duration / 3600000000000)}ч</p>
              <p><strong>Транспорт:</strong> ${route.transport_types.join(', ')}</p>
            </div>
          `).join('');
        }
      } catch (error) {
        statusDiv.innerHTML = `<p class="error">Ошибка: ${error.message}</p>`;
      }
    });
  </script>
</body>
</html>
```

---

## Тестовые запросы

### Вспомогательная функция для всех примеров

```javascript
// utils/testRoutes.js
import { searchRoutes } from '../config/api';

/**
 * Выполняет поиск маршрута и выводит результат в консоль
 */
async function testRoute(description, params) {
  console.log(`\n========== ${description} ==========`);
  console.log('Параметры:', params);

  try {
    const data = await searchRoutes(params);
    console.log(`✅ Найдено маршрутов: ${data.routes?.length || 0}`);
    console.log('Результат:', JSON.stringify(data, null, 2));
    return data;
  } catch (error) {
    console.error('❌ Ошибка:', error.message);
    throw error;
  }
}

export { testRoute };
```

---

## 1. ПОИСК АВИАМАРШРУТОВ

### 1.1 Москва → Якутск (все варианты)

```javascript
await testRoute('Москва → Якутск (все варианты)', {
  from: 'Moscow',
  to: 'Yakutsk',
  departure_date: '2025-11-20',
  passengers: 1
});

// Ожидаемый результат: 3 маршрута
// - Прямой рейс S7 Airlines (оптимальный, 32500₽)
// - Прямой рейс Yakutia Airlines (дешевый, 28000₽)
// - С пересадкой Ural Airlines (самый дешевый, 25000₽)
```

### 1.2 Якутск → Мирный (авиа vs автобус)

```javascript
await testRoute('Якутск → Мирный (авиа vs автобус)', {
  from: 'Yakutsk',
  to: 'Mirny',
  departure_date: '2025-11-20',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - ALROSA Air (1.5 часа, 12000₽)
// - Автобус ALROSA Transport (12 часов, 4500₽)
```

### 1.3 Якутск → Нерюнгри (авиа + автобус + комби)

```javascript
await testRoute('Якутск → Нерюнгри (авиа + автобус + комби)', {
  from: 'Yakutsk',
  to: 'Neryungri',
  departure_date: '2025-11-22',
  passengers: 1
});

// Ожидаемый результат: 3 маршрута
// - Прямой авиарейс Yakutia Airlines (1.5 часа, 13500₽)
// - Прямой автобус Amur Transport (10 часов, 3500₽)
// - Через БАМ с пересадками (13 часов, 4500₽)
```

### 1.4 Якутск → Тикси (север, летний рейс)

```javascript
await testRoute('Якутск → Тикси (север, летний рейс)', {
  from: 'Yakutsk',
  to: 'Tiksi',
  departure_date: '2025-06-15',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Polar Airlines (3 часа, 18000₽)
```

### 1.5 Якутск → Оймякон (полюс холода)

```javascript
await testRoute('Якутск → Оймякон (полюс холода)', {
  from: 'Yakutsk',
  to: 'Oymyakon',
  departure_date: '2025-11-22',
  passengers: 2
});

// Ожидаемый результат: 1 маршрут
// - Polar Airlines (2.5 часа, 19000₽ × 2 пассажира)
```

---

## 2. МУЛЬТИМОДАЛЬНЫЕ МАРШРУТЫ

### 2.1 Москва → Олекминск (авиа + такси + река)

```javascript
await testRoute('Москва → Олекминск (авиа + такси + река)', {
  from: 'Moscow',
  to: 'Olekminsk',
  departure_date: '2025-11-20',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - Оптимальный: Москва ✈️ Якутск → 🚕 Порт → ⛴️ Олекминск (41500₽)
// - Дешевый: Москва ✈️ Якутск → 🚕 Автовокзал → 🚌 Олекминск (32000₽)
//
// Особенности:
// - 3 вида транспорта
// - Пересадки в Якутске
// - Речной транспорт по Лене
```

### 2.2 Москва → Сангар (4 вида транспорта)

```javascript
await testRoute('Москва → Сангар (4 вида транспорта)', {
  from: 'Moscow',
  to: 'Sangur',
  departure_date: '2025-11-20',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Москва ✈️ Якутск → 🚕 Автовокзал → 🚌 Покровск → ⛴️ Сангар (43000₽)
//
// Сегменты:
// 1. S7 Airlines (авиа)
// 2. Yandex Taxi (такси)
// 3. Siberia Lines (автобус)
// 4. Sakha River Transport (река)
```

### 2.3 Тында → Якутск (ЖД БАМ + автобус)

```javascript
await testRoute('Тында → Якутск (ЖД БАМ + автобус)', {
  from: 'Tynda',
  to: 'Yakutsk',
  departure_date: '2025-11-22',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - Прямой ЖД до Нижнего Бестяха + переправа (9100₽, ~28 часов)
// - Только ЖД до Нижнего Бестяха (8500₽, остановка на другом берегу)
//
// Особенности:
// - Амуро-Якутская магистраль (АЯМ)
// - Переправа через Лену (паром летом, ледовая дорога зимой)
```

---

## 3. АВТОБУСНЫЕ МАРШРУТЫ ВНУТРИ ЯКУТИИ

### 3.1 Якутск → Покровск (пригород)

```javascript
await testRoute('Якутск → Покровск (пригород)', {
  from: 'Yakutsk',
  to: 'Pokrovsk',
  departure_date: '2025-11-22',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Avtotrans Yakutia (2.5 часа, 1200₽)
```

### 3.2 Мирный → Удачный (между алмазными городами)

```javascript
await testRoute('Мирный → Удачный (между алмазными городами)', {
  from: 'Mirny',
  to: 'Udachny',
  departure_date: '2025-11-22',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - ALROSA Transport (6 часов, 2800₽)
```

### 3.3 Нерюнгри → Алдан (короткий маршрут)

```javascript
await testRoute('Нерюнгри → Алдан (короткий маршрут)', {
  from: 'Neryungri',
  to: 'Aldan',
  departure_date: '2025-11-20',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Magistral (4 часа, 1200₽)
```

---

## 4. РЕЧНЫЕ МАРШРУТЫ (ЛЕТНЯЯ НАВИГАЦИЯ)

### 4.1 Якутск → Ленский (2 варианта)

```javascript
await testRoute('Якутск → Ленский (2 варианта)', {
  from: 'Yakutsk',
  to: 'Lensky',
  departure_date: '2025-11-20',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - Речной теплоход Lenskiye Zori (20 часов, 8000₽)
// - Автобус Siberia Lines (12 часов, 2500₽)
```

### 4.2 Якутск → Жиганск (дальний север по Лене)

```javascript
await testRoute('Якутск → Жиганск (дальний север по Лене)', {
  from: 'Yakutsk',
  to: 'Zhigansk',
  departure_date: '2025-06-20',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Lenskiye Zori (46 часов, 12000₽)
//
// Особенности:
// - Только летняя навигация (июнь-сентябрь)
// - Теплоход по реке Лена
// - 880 км вниз по течению
```

### 4.3 Якутск → Кюсюр (дельта Лены, 4+ дня)

```javascript
await testRoute('Якутск → Кюсюр (дельта Лены, 4+ дня)', {
  from: 'Yakutsk',
  to: 'Kyusyur',
  departure_date: '2025-06-21',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Lenskiye Zori (102 часа / 4.25 дня, 25000₽)
```

### 4.4 Покровск → Сангар (короткий речной рейс)

```javascript
await testRoute('Покровск → Сангар (короткий речной рейс)', {
  from: 'Pokrovsk',
  to: 'Sangur',
  departure_date: '2025-06-20',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Sakha River Transport (12 часов, 3500₽)
```

---

## 5. ЖЕЛЕЗНОДОРОЖНЫЕ МАРШРУТЫ (БАМ/АЯМ)

### 5.1 Томмот → Нижний Бестях (АЯМ)

```javascript
await testRoute('Томмот → Нижний Бестях (АЯМ)', {
  from: 'Tommot',
  to: 'Nizhny Bestyakh',
  departure_date: '2025-11-22',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - РЖД АЯМ (13 часов, 4200₽)
//
// Справка:
// - АЯМ - Амуро-Якутская железнодорожная магистраль
// - Конечная станция на правом берегу Лены
// - Открыта в 2019 году
```

### 5.2 Нижний Бестях → Якутск (переправа)

```javascript
await testRoute('Нижний Бестях → Якутск (переправа)', {
  from: 'Nizhny Bestyakh',
  to: 'Yakutsk',
  departure_date: '2025-11-22',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Автобус + паром/ледовая дорога (1.5 часа, 600₽)
//
// Особенности:
// - Летом: паромная переправа
// - Зимой: ледовая автодорога по Лене
// - Весна/осень: возможны задержки из-за ледохода
```

---

## 6. СЛОЖНЫЕ МУЛЬТИМОДАЛЬНЫЕ МАРШРУТЫ

### 6.1 Москва → Якутск через БАМ (экстрим)

```javascript
await testRoute('Москва → Якутск через БАМ (экстрим)', {
  from: 'Moscow',
  to: 'Yakutsk',
  departure_date: '2025-11-21',
  passengers: 1
});

// Ожидаемый результат: 4+ маршрута включая:
// - Прямые авиарейсы (3 варианта, 25000-32500₽)
// - Москва → Якутск → БАМ туда-обратно → Якутск (28000₽, 70+ часов)
//
// Описание сложного маршрута:
// 1. ✈️ Москва → Якутск (S7 Airlines)
// 2. 🚕 Аэропорт → Автовокзал (Yandex Taxi)
// 3. 🚌 Якутск → Нижний Бестях (паром)
// 4. 🚂 Нижний Бестях → Тында (БАМ)
// 5. 🚂 Тында → Нижний Бестях (обратно)
// 6. 🚌 Нижний Бестях → Якутск (паром)
```

---

## 7. СПЕЦИФИЧЕСКИЕ СЦЕНАРИИ

### 7.1 Множественные пассажиры

```javascript
await testRoute('Множественные пассажиры', {
  from: 'Yakutsk',
  to: 'Mirny',
  departure_date: '2025-11-20',
  passengers: 5
});

// Особенности:
// - Цена умножается на количество пассажиров
// - Проверка доступности мест (seat_count)
```

### 7.2 Дальний восток (Зырянка)

```javascript
await testRoute('Дальний восток (Зырянка)', {
  from: 'Yakutsk',
  to: 'Zyryanka',
  departure_date: '2025-11-22',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Polar Airlines (4 часа, 24000₽)
```

### 7.3 Алмазная провинция (Удачный)

```javascript
await testRoute('Алмазная провинция (Удачный)', {
  from: 'Yakutsk',
  to: 'Udachny',
  departure_date: '2025-11-22',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - ALROSA Air (2 часа 15 минут, 15000₽)
```

---

## 8. ТЕСТИРОВАНИЕ EDGE CASES

### 8.1 Несуществующий маршрут

```javascript
await testRoute('Несуществующий маршрут', {
  from: 'Yakutsk',
  to: 'Vladivostok',
  departure_date: '2025-11-20',
  passengers: 1
});

// Ожидаемый результат: Пустой массив маршрутов
```

### 8.2 Зимний/летний сезон (Тикси)

```javascript
// Летний рейс (работает)
await testRoute('Тикси летом (работает)', {
  from: 'Yakutsk',
  to: 'Tiksi',
  departure_date: '2025-06-15',
  passengers: 1
});

// Зимний период (может не быть рейсов)
await testRoute('Тикси зимой (может не быть рейсов)', {
  from: 'Yakutsk',
  to: 'Tiksi',
  departure_date: '2025-12-15',
  passengers: 1
});
```

---

## 9. НОВЫЕ МАРШРУТЫ - МОСКВА → РЕГИОНАЛЬНЫЕ ГОРОДА

### 9.1 Москва → Мирный (2 варианта)

```javascript
await testRoute('Москва → Мирный (2 варианта)', {
  from: 'Moscow',
  to: 'Mirny',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - Прямой ALROSA Air (9.5 часов, 38000₽)
// - Через Якутск (11.5 часов, 35000₽)
```

### 9.2 Москва → Нерюнгри (через Якутск)

```javascript
await testRoute('Москва → Нерюнгри (через Якутск)', {
  from: 'Moscow',
  to: 'Neryungri',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Москва → Якутск → Нерюнгри (11 часов, 36000₽)
```

### 9.3 Москва → Удачный (алмазный город)

```javascript
await testRoute('Москва → Удачный (алмазный город)', {
  from: 'Moscow',
  to: 'Udachny',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Москва → Якутск → Удачный (13 часов, 42000₽)
```

---

## 10. ОБРАТНЫЕ МАРШРУТЫ В МОСКВУ

### 10.1 Якутск → Москва (выбор времени)

```javascript
await testRoute('Якутск → Москва (выбор времени)', {
  from: 'Yakutsk',
  to: 'Moscow',
  departure_date: '2025-11-24',
  passengers: 1
});

// Ожидаемый результат: 3+ маршрута
// - Утренний S7 Airlines (8 часов, 33000₽)
// - Вечерний Yakutia Airlines (8 часов, 29000₽)
// - Ночной рейс (8.5 часов, 27000₽)
```

### 10.2 Мирный → Москва

```javascript
await testRoute('Мирный → Москва', {
  from: 'Mirny',
  to: 'Moscow',
  departure_date: '2025-11-24',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Прямой ALROSA Air (9 часов, 39000₽)
```

---

## 11. МАРШРУТЫ ЧЕРЕЗ СИБИРСКИЕ ХАБЫ

### 11.1 Москва → Якутск через Новосибирск (дешево!)

```javascript
await testRoute('Москва → Якутск через Новосибирск', {
  from: 'Moscow',
  to: 'Yakutsk',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 7+ маршрутов включая:
// - Через Новосибирск (16 часов, 24000₽) - **САМЫЙ ДЕШЕВЫЙ**
// - Через Иркутск (17.5 часов, 27000₽)
// - Через Красноярск (17 часов, 26000₽)
// - Прямые рейсы (8-9 часов, 25000-32500₽)
//
// Особенности:
// - Пересадки в крупных сибирских аэропортах
// - Больше времени, но дешевле
// - Возможность осмотреть другие города
```

---

## 12. МЕЖРЕГИОНАЛЬНЫЕ МАРШРУТЫ ВНУТРИ ЯКУТИИ

### 12.1 Мирный → Нерюнгри (прямой авиа)

```javascript
await testRoute('Мирный → Нерюнгри (прямой авиа)', {
  from: 'Mirny',
  to: 'Neryungri',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - Прямой авиа Yakutia Airlines (1ч 45мин, 14000₽)
// - Комбо через Якутск (автобус+авиа, 15 часов, 16000₽)
```

### 12.2 Удачный → Якутск (через Мирный)

```javascript
await testRoute('Удачный → Якутск (через Мирный)', {
  from: 'Udachny',
  to: 'Yakutsk',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - Прямой авиа (2ч 15мин, 15000₽)
// - Автобус → Мирный, затем самолет (7.5 часов, 18500₽)
```

### 12.3 Алдан → Якутск (автобус)

```javascript
await testRoute('Алдан → Якутск (автобус)', {
  from: 'Aldan',
  to: 'Yakutsk',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Magistral (10 часов, 3300₽)
```

### 12.4 Вилюйск → Нюрба (местный автобус)

```javascript
await testRoute('Вилюйск → Нюрба (местный автобус)', {
  from: 'Vilyuysk',
  to: 'Nyurba',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Avtotrans Yakutia (5 часов, 1800₽)
```

### 12.5 Вилюйск → Якутск (авиа vs автобус)

```javascript
await testRoute('Вилюйск → Якутск (авиа vs автобус)', {
  from: 'Vilyuysk',
  to: 'Yakutsk',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 3 маршрута
// - Утренний авиарейс Polar Airlines (1ч 40мин, 10000₽)
// - Вечерний авиарейс (1ч 40мин, 10500₽)
// - Автобус (10 часов, 2800₽)
```

---

## 13. СЛОЖНЫЕ МУЛЬТИМОДАЛЬНЫЕ МАРШРУТЫ

### 13.1 Алдан → Мирный (3 сегмента)

```javascript
await testRoute('Алдан → Мирный (3 сегмента)', {
  from: 'Aldan',
  to: 'Mirny',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Алдан 🚌 Нерюнгри 🚌 Якутск ✈️ Мирный (18 часов, 19000₽)
//
// Сегменты:
// 1. Автобус Magistral (4 часа)
// 2. Автобус Amur Transport (10 часов)
// 3. Авиа ALROSA Air (1 час)
```

### 13.2 Нерюнгри → Мирный (через Якутск)

```javascript
await testRoute('Нерюнгри → Мирный (через Якутск)', {
  from: 'Neryungri',
  to: 'Mirny',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - Прямой авиа (1ч 45мин, 14000₽)
// - Автобус + авиа через Якутск (15 часов, 16000₽)
```

---

## 14. ДОПОЛНИТЕЛЬНЫЕ РЕЧНЫЕ МАРШРУТЫ

### 14.1 Якутск → Батагай (дальний север по Лене)

```javascript
await testRoute('Якутск → Батагай (дальний север по Лене)', {
  from: 'Yakutsk',
  to: 'Batagay',
  departure_date: '2025-06-22',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Lenskiye Zori (54 часа / 2.25 дня, 16000₽)
//
// Особенности:
// - 1200 км по реке Лена
// - Только летняя навигация
// - Проходит через Жиганск
```

### 14.2 Ленский → Витим (речной маршрут)

```javascript
await testRoute('Ленский → Витим (речной маршрут)', {
  from: 'Lensky',
  to: 'Vitim',
  departure_date: '2025-06-23',
  passengers: 1
});

// Ожидаемый результат: 1 маршрут
// - Sakha River Transport (30 часов, 6500₽)
```

---

## 15. НОЧНЫЕ РЕЙСЫ

### 15.1 Москва → Якутск (ночной)

```javascript
await testRoute('Москва → Якутск (ночной)', {
  from: 'Moscow',
  to: 'Yakutsk',
  departure_date: '2025-11-23',
  passengers: 1
});

// Особенности:
// - Вылет в 23:00, прибытие в 07:30
// - Дешевле дневных рейсов (27000₽)
// - Можно поспать в полете
```

### 15.2 Якутск → Нерюнгри (ночной автобус)

```javascript
await testRoute('Якутск → Нерюнгри (ночной автобус)', {
  from: 'Yakutsk',
  to: 'Neryungri',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 4 маршрута
// - Дневной авиа (1.5 часа, 13500₽)
// - Дневной автобус (10 часов, 3500₽)
// - Ночной автобус (10 часов, 3200₽) - **ДЕШЕВЛЕ**
// - Комбо ЖД+автобус через БАМ (13 часов, 4500₽)
```

---

## 16. ВЕЧЕРНИЕ РЕЙСЫ (РАСШИРЕННОЕ РАСПИСАНИЕ)

### 16.1 Якутск → Мирный (вечерний рейс)

```javascript
await testRoute('Якутск → Мирный (вечерний рейс)', {
  from: 'Yakutsk',
  to: 'Mirny',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 3 маршрута
// - Утренний авиа 09:00 (1.5 часа, 12000₽)
// - Вечерний авиа 17:00 (1.5 часа, 12500₽)
// - Дневной автобус 07:00 (12 часов, 4500₽)
```

### 16.2 Якутск → Покровск (выбор времени)

```javascript
await testRoute('Якутск → Покровск (выбор времени)', {
  from: 'Yakutsk',
  to: 'Pokrovsk',
  departure_date: '2025-11-23',
  passengers: 1
});

// Ожидаемый результат: 2 маршрута
// - Утренний автобус 08:00 (2.5 часа, 1200₽)
// - Вечерний автобус 17:00 (2.5 часа, 1250₽)
```

---

## Запуск всех тестов

### Функция для массового тестирования

```javascript
// tests/runAllTests.js
import { testRoute } from '../utils/testRoutes';

async function runAllTests() {
  console.log('🚀 Запуск всех тестов маршрутов LenaLink\n');

  const tests = [
    // Авиамаршруты
    {
      name: 'Москва → Якутск',
      params: { from: 'Moscow', to: 'Yakutsk', departure_date: '2025-11-20', passengers: 1 }
    },
    {
      name: 'Якутск → Мирный',
      params: { from: 'Yakutsk', to: 'Mirny', departure_date: '2025-11-20', passengers: 1 }
    },
    // Добавьте остальные тесты...
  ];

  let passed = 0;
  let failed = 0;

  for (const test of tests) {
    try {
      await testRoute(test.name, test.params);
      passed++;
    } catch (error) {
      failed++;
      console.error(`❌ Тест "${test.name}" провален:`, error.message);
    }
  }

  console.log(`\n========== РЕЗУЛЬТАТЫ ==========`);
  console.log(`✅ Пройдено: ${passed}`);
  console.log(`❌ Провалено: ${failed}`);
  console.log(`📊 Всего: ${tests.length}`);
}

// Запуск
runAllTests().catch(console.error);
```

### Запуск в браузере (DevTools Console)

```javascript
// Скопируйте и вставьте в консоль браузера
(async function() {
  const API_BASE_URL = 'https://lena.linkpc.net/api/v1';

  async function search(params) {
    const response = await fetch(`${API_BASE_URL}/routes/search`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(params),
    });
    return response.json();
  }

  // Тест 1: Москва → Якутск
  const result1 = await search({
    from: 'Moscow',
    to: 'Yakutsk',
    departure_date: '2025-11-20',
    passengers: 1
  });
  console.log('Москва → Якутск:', result1);

  // Добавьте другие тесты...
})();
```

---

## Полезные утилиты

### Форматирование данных

```javascript
// utils/formatters.js

/**
 * Форматирует длительность из наносекунд в читаемый формат
 */
export function formatDuration(nanoseconds) {
  const hours = Math.floor(nanoseconds / 3600000000000);
  const minutes = Math.floor((nanoseconds % 3600000000000) / 60000000000);

  if (hours === 0) {
    return `${minutes}мин`;
  }
  if (minutes === 0) {
    return `${hours}ч`;
  }
  return `${hours}ч ${minutes}мин`;
}

/**
 * Форматирует цену с разделителями
 */
export function formatPrice(price) {
  return `${price.toLocaleString('ru-RU')}₽`;
}

/**
 * Форматирует дату и время
 */
export function formatDateTime(isoString) {
  const date = new Date(isoString);
  return date.toLocaleString('ru-RU', {
    day: '2-digit',
    month: '2-digit',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });
}

/**
 * Иконки транспорта
 */
export const TRANSPORT_ICONS = {
  air: '✈️',
  bus: '🚌',
  rail: '🚂',
  river: '⛴️',
  taxi: '🚕',
  ferry: '⛴️',
};

/**
 * Форматирует сегменты маршрута
 */
export function formatRouteSegments(segments) {
  return segments.map(segment =>
    `${TRANSPORT_ICONS[segment.transport_type] || '🚗'} ${segment.from} → ${segment.to} (${segment.provider})`
  ).join('\n');
}
```

---

## Обработка ошибок

```javascript
// utils/errorHandler.js

export class APIError extends Error {
  constructor(message, statusCode, details) {
    super(message);
    this.name = 'APIError';
    this.statusCode = statusCode;
    this.details = details;
  }
}

export async function handleAPIResponse(response) {
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    throw new APIError(
      errorData.message || `HTTP Error: ${response.status}`,
      response.status,
      errorData
    );
  }
  return response.json();
}

// Использование
try {
  const response = await fetch(`${API_BASE_URL}/routes/search`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(params),
  });

  const data = await handleAPIResponse(response);
  // Обработка данных...
} catch (error) {
  if (error instanceof APIError) {
    console.error(`API Error [${error.statusCode}]:`, error.message);
    console.error('Details:', error.details);
  } else {
    console.error('Network Error:', error.message);
  }
}
```

---

## СВОДНАЯ СТАТИСТИКА ТЕСТОВЫХ ДАННЫХ

### Транспортные узлы (stops)
- **Аэропорты:** 36 (включая Магадан, Хабаровск, сибирские хабы)
- **ЖД станции:** 5 (БАМ/АЯМ)
- **Речные порты:** 11
- **Автовокзалы:** 15

### Маршруты (routes)
- **Всего маршрутов:** 71+
- **Прямые авиа:** ~25
- **Автобусные:** ~18
- **ЖД:** 5
- **Речные:** 7
- **Мультимодальные:** ~16

### Провайдеры
- **Авиа:** S7 Airlines, Yakutia Airlines, ALROSA Air, Polar Airlines, Ural Airlines
- **ЖД:** РЖД (БАМ/АЯМ)
- **Автобус:** Avtotrans Yakutia, ALROSA Transport, Magistral, Siberia Lines, Amur Transport, City Transport
- **Река:** Lenskiye Zori, Sakha River Transport
- **Такси:** Yandex Taxi, Maxim
- **Паром:** Ferry/Ice Road

### География
- **Центр:** Якутск (главный хаб)
- **Алмазы:** Мирный, Удачный, Полярный
- **Север:** Тикси, Жиганск, Кюсюр, Верхоянск, Батагай, Саскылах
- **Восток:** Оймякон, Зырянка, Среднеколымск, Депутатский
- **Юг:** Нерюнгри, Алдан, Томмот, Чульман (БАМ)
- **Запад:** Вилюйск, Ньюрба, Ленский, Сунтар
- **Связи:** Москва, Новосибирск, Красноярск, Иркутск, Магадан, Хабаровск
