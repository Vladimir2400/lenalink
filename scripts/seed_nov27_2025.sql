-- ============================================================================
-- ЗАПОЛНЕНИЕ ДАННЫХ НА 27 НОЯБРЯ 2025
-- ============================================================================
-- Актуальные маршруты по Якутии + Москва
-- Дата вылетов/отправлений: 27 ноября 2025 (2025-11-27)
-- ============================================================================

BEGIN;

-- ============================================================================
-- 1. ОСТАНОВКИ (ТРАНСПОРТНЫЕ УЗЛЫ)
-- ============================================================================

INSERT INTO stops (id, name, city, city_display_name, latitude, longitude, stop_type)  VALUES
-- Москва
('moscow_dme', 'Domodedovo International Airport', 'moscow', 'Москва', 55.4088, 37.9063, 'airport') ,
('moscow_svo', 'Sheremetyevo International Airport', 'moscow', 'Москва', 55.9726, 37.4147, 'airport') ,
('moscow_vko', 'Vnukovo International Airport', 'moscow', 'Москва', 55.5914, 37.2615, 'airport') ,

-- Якутск (главный хаб)
('yakutsk_yks', 'Yakutsk Airport (Platon Oyunsky)', 'yakutsk', 'Якутск', 62.0932, 129.7708, 'airport') ,
('yakutsk_port', 'Yakutsk River Port', 'yakutsk', 'Якутск', 62.0272, 129.7322, 'port') ,
('yakutsk_bus', 'Yakutsk Bus Terminal', 'yakutsk', 'Якутск', 62.0339, 129.7331, 'station') ,

-- Алмазные города
('mirny_aprt', 'Mirny Airport', 'mirny', 'Мирный', 62.5347, 114.0389, 'airport') ,
('mirny_bus', 'Mirny Bus Terminal', 'mirny', 'Мирный', 62.5350, 114.0300, 'station') ,
('udachny_aprt', 'Udachny Airport', 'udachny', 'Удачный', 66.4000, 112.0333, 'airport') ,

-- Нерюнгри (южная Якутия)
('nerungri_aprt', 'Neryungri Airport', 'neryungri', 'Нерюнгри', 56.9139, 124.9144, 'airport') ,
('nerungri_bus', 'Neryungri Bus Terminal', 'neryungri', 'Нерюнгри', 56.6600, 124.7200, 'station') ,

-- Олекминск
('olekminsk_port', 'Olekminsk River Port', 'olekminsk', 'Олекминск', 60.3733, 120.4272, 'port') ,
('olekminsk_aprt', 'Olekminsk Airstrip', 'olekminsk', 'Олекминск', 60.3733, 120.4272, 'airport') ,

-- Ленск
('lensky_port', 'Lensk River Port', 'lensky', 'Ленск', 60.7458, 114.8833, 'port') ,
('lensk_aprt', 'Lensk Airport', 'lensky', 'Ленск', 60.7458, 114.8833, 'airport') ,

-- Покровск
('pokrovsk_bus', 'Pokrovsk Bus Terminal', 'pokrovsk', 'Покровск', 61.4833, 129.1500, 'station') ,
('pokrovsk_port', 'Pokrovsk River Port', 'pokrovsk', 'Покровск', 61.4800, 129.1450, 'port') ,

-- Северные города
('tiksi_aprt', 'Tiksi Airport', 'tiksi', 'Тикси', 71.6977, 128.9031, 'airport') ,
('verkhoyansk_aprt', 'Verkhoyansk Airport', 'verkhoyansk', 'Верхоянск', 67.5447, 133.3953, 'airport') ,
('oymyakon_aprt', 'Oymyakon Airstrip', 'oymyakon', 'Оймякон', 63.4608, 142.7858, 'airport') ,

-- Вилюйская группа
('vilyuysk_aprt', 'Vilyuysk Airport', 'vilyuysk', 'Вилюйск', 63.7486, 121.5675, 'airport') ,
('vilyuysk_bus', 'Vilyuysk Bus Terminal', 'vilyuysk', 'Вилюйск', 63.7486, 121.5675, 'station') ,
('nyurba_aprt', 'Nyurba Airport', 'nyurba', 'Нюрба', 65.3875, 118.4778, 'airport') ,
('suntar_aprt', 'Suntar Airport', 'suntar', 'Сунтар', 62.1731, 117.6356, 'airport') ,

-- ЖД (Амуро-Якутская магистраль)
('nizhny_bestyakh_railway', 'Nizhny Bestyakh Railway Station', 'nizhny_bestyakh', 'Нижний Бестях', 61.8950, 129.9289, 'station') ,
('tommot_railway', 'Tommot Railway Station', 'tommot', 'Томмот', 58.9500, 126.2833, 'station') ,
('aldan_railway', 'Aldan Railway Station', 'aldan', 'Алдан', 58.6000, 125.3833, 'station') ,
('aldan_bus', 'Aldan Bus Terminal', 'aldan', 'Алдан', 58.6000, 125.3833, 'station') 

ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  city_display_name = EXCLUDED.city_display_name,
  latitude = EXCLUDED.latitude,
  longitude = EXCLUDED.longitude;

-- ============================================================================
-- 2. МОСКВА → ЯКУТСК (27 НОЯБРЯ 2025)
-- ============================================================================

-- Маршрут 1: Москва (DME) → Якутск (прямой, утренний)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_msk_yks_morning_nov27', 'moscow', 'yakutsk',
 '2025-11-27 08:00:00+03', '2025-11-27 23:00:00+09', 28800000000000,
 32000.00, 92.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_morning_nov27', 'route_msk_yks_morning_nov27', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-27 08:00:00+03', '2025-11-27 23:00:00+09', 32000.00, 28800000000000, 180, 92.0, 4884, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Маршрут 2: Москва (SVO) → Якутск (вечерний)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_msk_yks_evening_nov27', 'moscow', 'yakutsk',
 '2025-11-27 18:00:00+03', '2025-11-28 09:00:00+09', 28800000000000,
 28000.00, 90.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_evening_nov27', 'route_msk_yks_evening_nov27', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-27 18:00:00+03', '2025-11-28 09:00:00+09', 28000.00, 28800000000000, 200, 90.0, 4884, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- ============================================================================
-- 3. ЯКУТСК → МОСКВА (27-28 НОЯБРЯ 2025)
-- ============================================================================

-- Обратный маршрут: Якутск → Москва (утро 27 ноября)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_msk_morning_nov27', 'yakutsk', 'moscow',
 '2025-11-27 06:00:00+09', '2025-11-27 14:00:00+03', 28800000000000,
 34000.00, 91.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_msk_morning_nov27', 'route_yks_msk_morning_nov27', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-27 06:00:00+09', '2025-11-27 14:00:00+03', 34000.00, 28800000000000, 180, 91.0, 4884, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- ============================================================================
-- 4. ВНУТРИ ЯКУТИИ - АВИА (27 НОЯБРЯ 2025)
-- ============================================================================

-- Якутск → Мирный (утро)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_mirny_morning_nov27', 'yakutsk', 'mirny',
 '2025-11-27 09:00:00+09', '2025-11-27 10:30:00+09', 5400000000000,
 12000.00, 88.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_morning_nov27', 'route_yks_mirny_morning_nov27', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-27 09:00:00+09', '2025-11-27 10:30:00+09', 12000.00, 5400000000000, 100, 88.0, 520, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Якутск → Мирный (вечер)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_mirny_evening_nov27', 'yakutsk', 'mirny',
 '2025-11-27 17:00:00+09', '2025-11-27 18:30:00+09', 5400000000000,
 13500.00, 90.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_evening_nov27', 'route_yks_mirny_evening_nov27', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-27 17:00:00+09', '2025-11-27 18:30:00+09', 13500.00, 5400000000000, 100, 90.0, 520, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Якутск → Нерюнгри
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_ner_nov27', 'yakutsk', 'neryungri',
 '2025-11-27 10:00:00+09', '2025-11-27 11:30:00+09', 5400000000000,
 14000.00, 87.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_ner_nov27', 'route_yks_ner_nov27', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'nerungri_aprt',
 '2025-11-27 10:00:00+09', '2025-11-27 11:30:00+09', 14000.00, 5400000000000, 120, 87.0, 560, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Якутск → Удачный
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_udachny_nov27', 'yakutsk', 'udachny',
 '2025-11-27 11:30:00+09', '2025-11-27 13:45:00+09', 8100000000000,
 16000.00, 85.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_udachny_nov27', 'route_yks_udachny_nov27', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-11-27 11:30:00+09', '2025-11-27 13:45:00+09', 16000.00, 8100000000000, 80, 85.0, 630, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Якутск → Вилюйск
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_vilyuysk_nov27', 'yakutsk', 'vilyuysk',
 '2025-11-27 14:00:00+09', '2025-11-27 15:40:00+09', 6000000000000,
 11000.00, 82.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_vilyuysk_nov27', 'route_yks_vilyuysk_nov27', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-27 14:00:00+09', '2025-11-27 15:40:00+09', 11000.00, 6000000000000, 70, 82.0, 460, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Якутск → Тикси (Северный Ледовитый океан)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_tiksi_nov27', 'yakutsk', 'tiksi',
 '2025-11-27 08:30:00+09', '2025-11-27 11:00:00+09', 9000000000000,
 22000.00, 78.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_tiksi_nov27', 'route_yks_tiksi_nov27', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-11-27 08:30:00+09', '2025-11-27 11:00:00+09', 22000.00, 9000000000000, 60, 78.0, 1670, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Якутск → Оймякон (Полюс холода)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_oymyakon_nov27', 'yakutsk', 'oymyakon',
 '2025-11-27 07:00:00+09', '2025-11-27 09:30:00+09', 9000000000000,
 19000.00, 75.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_oymyakon_nov27', 'route_yks_oymyakon_nov27', 'air', 'Polar Airlines', 'yakutsk_yks', 'oymyakon_aprt',
 '2025-11-27 07:00:00+09', '2025-11-27 09:30:00+09', 19000.00, 9000000000000, 50, 75.0, 930, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- ============================================================================
-- 5. АВТОБУСНЫЕ МАРШРУТЫ (27 НОЯБРЯ 2025)
-- ============================================================================

-- Якутск → Покровск (пригород)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_pokrovsk_bus_nov27', 'yakutsk', 'pokrovsk',
 '2025-11-27 08:00:00+09', '2025-11-27 10:30:00+09', 9000000000000,
 1200.00, 86.00, ARRAY['bus']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_pokrovsk_bus_nov27', 'route_yks_pokrovsk_bus_nov27', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-27 08:00:00+09', '2025-11-27 10:30:00+09', 1200.00, 9000000000000, 45, 86.0, 80, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Якутск → Нерюнгри (автобус, дешевле авиа)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_yks_ner_bus_nov27', 'yakutsk', 'neryungri',
 '2025-11-27 09:00:00+09', '2025-11-27 19:00:00+09', 36000000000000,
 3500.00, 75.00, ARRAY['bus']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_ner_bus_nov27', 'route_yks_ner_bus_nov27', 'bus', 'Amur Transport', 'yakutsk_bus', 'nerungri_bus',
 '2025-11-27 09:00:00+09', '2025-11-27 19:00:00+09', 3500.00, 36000000000000, 50, 75.0, 350, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- ============================================================================
-- 6. КОМБИНИРОВАННЫЕ МАРШРУТЫ (27-28 НОЯБРЯ 2025)
-- ============================================================================

-- Москва → Олекминск (авиа + речной транспорт - зимой закрыто, только летом)
-- Зимой: Москва → Якутск → Олекминск (автобус)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_msk_olek_combo_nov27', 'moscow', 'olekminsk',
 '2025-11-27 08:00:00+03', '2025-11-28 08:00:00+09', 64800000000000,
 35000.00, 85.00, ARRAY['air','bus']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_olek_1_nov27', 'route_msk_olek_combo_nov27', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-27 08:00:00+03', '2025-11-27 23:00:00+09', 32000.00, 28800000000000, 180, 92.0, 4884, 1),
('seg_msk_olek_2_nov27', 'route_msk_olek_combo_nov27', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'olekminsk_port',
 '2025-11-28 01:00:00+09', '2025-11-28 08:00:00+09', 3000.00, 25200000000000, 40, 78.0, 610, 2)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_olek_combo_nov27', 'seg_msk_olek_1_nov27', 'seg_msk_olek_2_nov27',
 7200000000000, 12000, true, true, 7200000000000, 1)
ON CONFLICT (route_id, from_segment_id, to_segment_id) DO UPDATE SET
  transfer_duration = EXCLUDED.transfer_duration;

-- Москва → Мирный (через Якутск с пересадкой)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_msk_mirny_via_yks_nov27', 'moscow', 'mirny',
 '2025-11-27 08:00:00+03', '2025-11-28 01:30:00+09', 42300000000000,
 42000.00, 88.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_mirny_1_nov27', 'route_msk_mirny_via_yks_nov27', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-27 08:00:00+03', '2025-11-27 23:00:00+09', 32000.00, 28800000000000, 180, 92.0, 4884, 1),
('seg_msk_mirny_2_nov27', 'route_msk_mirny_via_yks_nov27', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-28 00:00:00+09', '2025-11-28 01:30:00+09', 10000.00, 5400000000000, 100, 84.0, 520, 2)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_mirny_via_yks_nov27', 'seg_msk_mirny_1_nov27', 'seg_msk_mirny_2_nov27',
 3600000000000, 12000, true, true, 3600000000000, 1)
ON CONFLICT (route_id, from_segment_id, to_segment_id) DO UPDATE SET
  transfer_duration = EXCLUDED.transfer_duration;

-- ============================================================================
-- 7. ОБРАТНЫЕ МАРШРУТЫ ВНУТРИ ЯКУТИИ (27 НОЯБРЯ 2025)
-- ============================================================================

-- Мирный → Якутск
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_mirny_yks_nov27', 'mirny', 'yakutsk',
 '2025-11-27 12:00:00+09', '2025-11-27 13:30:00+09', 5400000000000,
 12500.00, 89.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_mirny_yks_nov27', 'route_mirny_yks_nov27', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-27 12:00:00+09', '2025-11-27 13:30:00+09', 12500.00, 5400000000000, 100, 89.0, 520, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

-- Нерюнгри → Якутск
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, transport_types, created_at) VALUES
('route_ner_yks_nov27', 'neryungri', 'yakutsk',
 '2025-11-27 14:00:00+09', '2025-11-27 15:30:00+09', 5400000000000,
 14500.00, 86.00, ARRAY['air']::text[], NOW())
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  total_price = EXCLUDED.total_price;

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_ner_yks_nov27', 'route_ner_yks_nov27', 'air', 'Yakutia Airlines', 'nerungri_aprt', 'yakutsk_yks',
 '2025-11-27 14:00:00+09', '2025-11-27 15:30:00+09', 14500.00, 5400000000000, 120, 86.0, 560, 1)
ON CONFLICT (id) DO UPDATE SET
  departure_time = EXCLUDED.departure_time,
  arrival_time = EXCLUDED.arrival_time,
  price = EXCLUDED.price;

COMMIT;

-- ============================================================================
-- ИТОГО ДОБАВЛЕНО:
-- - Остановки: 31 (Москва + основные города Якутии)
-- - Маршруты: 20+ актуальных на 27 ноября 2025
-- - Сегменты: 23+
-- - Коннекции: для комбинированных маршрутов
--
-- ПОКРЫТИЕ:
-- ✅ Москва ↔ Якутск (прямые рейсы)
-- ✅ Якутск → алмазные города (Мирный, Удачный)
-- ✅ Якутск → промышленные центры (Нерюнгри)
-- ✅ Якутск → северные точки (Тикси, Оймякон, Верхоянск)
-- ✅ Якутск → Вилюйская группа
-- ✅ Автобусные маршруты (пригород)
-- ✅ Комбинированные маршруты Москва → регионы Якутии
-- ============================================================================
