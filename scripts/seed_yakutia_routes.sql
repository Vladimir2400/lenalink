-- ============================================================================
-- РАСШИРЕННОЕ ЗАПОЛНЕНИЕ МАРШРУТОВ ПО ЯКУТИИ
-- ============================================================================
-- Этот скрипт добавляет дополнительные маршруты для тестирования
-- Покрывает: авиа, ЖД (БАМ), автобусы, речной транспорт
-- ============================================================================

-- ============================================================================
-- ДОПОЛНИТЕЛЬНЫЕ ОСТАНОВКИ
-- ============================================================================

INSERT INTO stops (id, name, city, latitude, longitude, stop_type) VALUES
-- ЖД станции БАМ (расширенные)
('skovorodino_railway', 'Skovorodino Railway Station (BAM)', 'Skovorodino', 53.9867, 123.9494, 'station'),
('nizhny_bestyakh_railway', 'Nizhny Bestyakh Railway Station', 'Nizhny Bestyakh', 61.8950, 129.9289, 'station'),

-- Дополнительные аэропорты
('chulman_aprt', 'Chulman Airport', 'Chulman', 56.9333, 124.9167, 'airport'),
('polyarny_aprt', 'Polyarny Airport', 'Polyarny', 66.4000, 112.0333, 'airport'),
('deputatsky_aprt', 'Deputatsky Airport', 'Deputatsky', 69.3958, 139.8986, 'airport'),
('zyryanka_aprt', 'Zyryanka Airport', 'Zyryanka', 65.7444, 150.8861, 'airport'),
('srednekolymsk_aprt', 'Srednekolymsk Airport', 'Srednekolymsk', 67.4805, 153.7364, 'airport'),
('saskylakh_aprt', 'Saskylakh Airport', 'Saskylakh', 71.9278, 114.0833, 'airport'),

-- Автовокзалы
('pokrovsk_bus', 'Pokrovsk Bus Terminal', 'Pokrovsk', 61.4833, 129.1500, 'station'),
('khangalas_bus', 'Khangalas Bus Terminal', 'Khangalas', 61.3667, 128.9833, 'station'),
('bestyakh_bus', 'Bestyakh Bus Terminal', 'Bestyakh', 61.9000, 129.9333, 'station'),
('kangalassy_bus', 'Kangalassy Bus Terminal', 'Kangalassy', 61.7500, 129.6667, 'station'),

-- Речные порты на притоках Лены
('vitim_port', 'Vitim River Port', 'Vitim', 59.4500, 112.5833, 'port'),
('peleduy_port', 'Peleduy River Port', 'Peleduy', 59.6333, 112.2500, 'port'),
('kyusyur_port', 'Kyusyur River Port', 'Kyusyur', 70.6833, 127.3833, 'port'),
('zhigansk_port', 'Zhigansk River Port', 'Zhigansk', 66.7667, 123.3667, 'port')

ON CONFLICT (name, city) DO NOTHING;

-- ============================================================================
-- ДОПОЛНИТЕЛЬНЫЕ МАРШРУТЫ (всего добавим 25+ новых маршрутов)
-- ============================================================================

-- ============================================================================
-- 1. ВНУТРИ ЯКУТИИ - АВИАМАРШРУТЫ
-- ============================================================================

-- Якутск → Нерюнгри (авиа)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_ner_air', 'Yakutsk', 'Neryungri',
 '2025-11-22 07:00:00', '2025-11-22 08:30:00', 5400000000000,
 13500.00, 90.00, 675.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_ner_air', 'route_yks_ner_air', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'nerungri_aprt',
 '2025-11-22 07:00:00', '2025-11-22 08:30:00', 13500.00, 5400000000000, 120, 90.0, 560, 1);

-- Якутск → Удачный (алмазы)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_udachny', 'Yakutsk', 'Udachny',
 '2025-11-22 09:30:00', '2025-11-22 11:45:00', 8100000000000,
 15000.00, 88.00, 750.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_udachny', 'route_yks_udachny', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-11-22 09:30:00', '2025-11-22 11:45:00', 15000.00, 8100000000000, 100, 88.0, 630, 1);

-- Якутск → Депутатский (север)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_deputatsky', 'Yakutsk', 'Deputatsky',
 '2025-11-22 08:00:00', '2025-11-22 11:30:00', 12600000000000,
 22000.00, 72.00, 1100.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_deputatsky', 'route_yks_deputatsky', 'air', 'Polar Airlines', 'yakutsk_yks', 'deputatsky_aprt',
 '2025-11-22 08:00:00', '2025-11-22 11:30:00', 22000.00, 12600000000000, 60, 72.0, 1420, 1);

-- Якутск → Зырянка (восток)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_zyryanka', 'Yakutsk', 'Zyryanka',
 '2025-11-22 10:00:00', '2025-11-22 14:00:00', 14400000000000,
 24000.00, 70.00, 1200.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_zyryanka', 'route_yks_zyryanka', 'air', 'Polar Airlines', 'yakutsk_yks', 'zyryanka_aprt',
 '2025-11-22 10:00:00', '2025-11-22 14:00:00', 24000.00, 14400000000000, 50, 70.0, 1650, 1);

-- Якутск → Оймякон (полюс холода)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_oymyakon', 'Yakutsk', 'Oymyakon',
 '2025-11-22 07:30:00', '2025-11-22 10:00:00', 9000000000000,
 19000.00, 75.00, 950.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_oymyakon', 'route_yks_oymyakon', 'air', 'Polar Airlines', 'yakutsk_yks', 'oymyakon_aprt',
 '2025-11-22 07:30:00', '2025-11-22 10:00:00', 19000.00, 9000000000000, 50, 75.0, 930, 1);

-- Якутск → Верхоянск (север)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_verkhoyansk', 'Yakutsk', 'Verkhoyansk',
 '2025-11-22 09:00:00', '2025-11-22 11:45:00', 9900000000000,
 20000.00, 73.00, 1000.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_verkhoyansk', 'route_yks_verkhoyansk', 'air', 'Polar Airlines', 'yakutsk_yks', 'verkhoyansk_aprt',
 '2025-11-22 09:00:00', '2025-11-22 11:45:00', 20000.00, 9900000000000, 50, 73.0, 675, 1);

-- Якутск → Ньюрба
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_nyurba', 'Yakutsk', 'Nyurba',
 '2025-11-22 11:00:00', '2025-11-22 12:30:00', 5400000000000,
 11000.00, 76.00, 550.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_nyurba', 'route_yks_nyurba', 'air', 'Polar Airlines', 'yakutsk_yks', 'nyurba_aprt',
 '2025-11-22 11:00:00', '2025-11-22 12:30:00', 11000.00, 5400000000000, 60, 76.0, 420, 1);

-- ============================================================================
-- 2. АВТОБУСНЫЕ МАРШРУТЫ ВНУТРИ ЯКУТИИ
-- ============================================================================

-- Якутск → Покровск (автобус)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_pokrovsk_bus', 'Yakutsk', 'Pokrovsk',
 '2025-11-22 08:00:00', '2025-11-22 10:30:00', 9000000000000,
 1200.00, 85.00, 60.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_pokrovsk_bus', 'route_yks_pokrovsk_bus', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-22 08:00:00', '2025-11-22 10:30:00', 1200.00, 9000000000000, 45, 85.0, 80, 1);

-- Якутск → Бестях (пригород, паром зимой мост)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_bestyakh_bus', 'Yakutsk', 'Bestyakh',
 '2025-11-22 07:00:00', '2025-11-22 07:45:00', 2700000000000,
 350.00, 90.00, 18.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_bestyakh_bus', 'route_yks_bestyakh_bus', 'bus', 'City Transport', 'yakutsk_bus', 'bestyakh_bus',
 '2025-11-22 07:00:00', '2025-11-22 07:45:00', 350.00, 2700000000000, 50, 90.0, 15, 1);

-- Мирный → Удачный (автобус между алмазными городами)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_mirny_udachny_bus', 'Mirny', 'Udachny',
 '2025-11-22 09:00:00', '2025-11-22 15:00:00', 21600000000000,
 2800.00, 72.00, 140.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_mirny_udachny_bus', 'route_mirny_udachny_bus', 'bus', 'ALROSA Transport', 'mirny_bus', 'udachny_aprt',
 '2025-11-22 09:00:00', '2025-11-22 15:00:00', 2800.00, 21600000000000, 40, 72.0, 260, 1);

-- Якутск → Кангалассы (пригород)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_kangalassy', 'Yakutsk', 'Kangalassy',
 '2025-11-22 09:00:00', '2025-11-22 10:15:00', 4500000000000,
 450.00, 88.00, 23.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_kangalassy', 'route_yks_kangalassy', 'bus', 'City Transport', 'yakutsk_bus', 'kangalassy_bus',
 '2025-11-22 09:00:00', '2025-11-22 10:15:00', 450.00, 4500000000000, 50, 88.0, 40, 1);

-- ============================================================================
-- 3. ЖЕЛЕЗНОДОРОЖНЫЕ МАРШРУТЫ (БАМ)
-- ============================================================================

-- Тында → Нижний Бестях (конец АЯМ - Амуро-Якутская магистраль)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_tynda_bestyakh_rail', 'Tynda', 'Nizhny Bestyakh',
 '2025-11-22 10:00:00', '2025-11-23 14:00:00', 100800000000000,
 8500.00, 82.00, 425.00, true, ARRAY['rail']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_tynda_bestyakh_rail', 'route_tynda_bestyakh_rail', 'rail', 'RZD (AYaM)', 'tynda_railway', 'nizhny_bestyakh_railway',
 '2025-11-22 10:00:00', '2025-11-23 14:00:00', 8500.00, 100800000000000, 300, 82.0, 1250, 1);

-- Томмот → Нижний Бестях
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_tommot_bestyakh_rail', 'Tommot', 'Nizhny Bestyakh',
 '2025-11-22 08:00:00', '2025-11-22 21:00:00', 46800000000000,
 4200.00, 88.00, 210.00, true, ARRAY['rail']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_tommot_bestyakh_rail', 'route_tommot_bestyakh_rail', 'rail', 'RZD (AYaM)', 'tommot_railway', 'nizhny_bestyakh_railway',
 '2025-11-22 08:00:00', '2025-11-22 21:00:00', 4200.00, 46800000000000, 300, 88.0, 450, 1);

-- Нижний Бестях → Якутск (паромная переправа/ледовая дорога + автобус)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_bestyakh_yks_combined', 'Nizhny Bestyakh', 'Yakutsk',
 '2025-11-22 22:00:00', '2025-11-22 23:30:00', 5400000000000,
 600.00, 85.00, 30.00, true, ARRAY['bus','walk']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_bestyakh_yks_1', 'route_bestyakh_yks_combined', 'bus', 'Avtotrans Yakutia', 'nizhny_bestyakh_railway', 'bestyakh_bus',
 '2025-11-22 22:00:00', '2025-11-22 22:20:00', 100.00, 1200000000000, 50, 90.0, 3, 1),
('seg_bestyakh_yks_2', 'route_bestyakh_yks_combined', 'bus', 'Ferry/Ice Road', 'bestyakh_bus', 'yakutsk_bus',
 '2025-11-22 22:30:00', '2025-11-22 23:30:00', 500.00, 3600000000000, 45, 80.0, 20, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_bestyakh_yks_combined', 'seg_bestyakh_yks_1', 'seg_bestyakh_yks_2',
 600000000000, 200, false, true, 600000000000, 1);

-- ============================================================================
-- 4. КОМБИНИРОВАННЫЕ МАРШРУТЫ (ЖД + АВТОБУС + АВИА)
-- ============================================================================

-- Тында → Якутск (ЖД до Бестяха + автобус)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_tynda_yks_combined', 'Tynda', 'Yakutsk',
 '2025-11-22 10:00:00', '2025-11-23 15:30:00', 105300000000000,
 9100.00, 80.00, 455.00, true, ARRAY['rail','bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_tynda_yks_1', 'route_tynda_yks_combined', 'rail', 'RZD (AYaM)', 'tynda_railway', 'nizhny_bestyakh_railway',
 '2025-11-22 10:00:00', '2025-11-23 14:00:00', 8500.00, 100800000000000, 300, 82.0, 1250, 1),
('seg_tynda_yks_2', 'route_tynda_yks_combined', 'bus', 'Ferry/Ice Road', 'nizhny_bestyakh_railway', 'yakutsk_bus',
 '2025-11-23 14:30:00', '2025-11-23 15:30:00', 600.00, 3600000000000, 45, 78.0, 22, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_tynda_yks_combined', 'seg_tynda_yks_1', 'seg_tynda_yks_2',
 1800000000000, 500, false, true, 1800000000000, 1);

-- Нерюнгри → Якутск (автобус через БАМ)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_ner_yks_via_bam', 'Neryungri', 'Yakutsk',
 '2025-11-22 09:00:00', '2025-11-22 22:00:00', 46800000000000,
 4500.00, 70.00, 225.00, true, ARRAY['bus','rail']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_ner_yks_bam_1', 'route_ner_yks_via_bam', 'bus', 'Magistral', 'nerungri_bus', 'tommot_railway',
 '2025-11-22 09:00:00', '2025-11-22 11:00:00', 800.00, 7200000000000, 40, 85.0, 70, 1),
('seg_ner_yks_bam_2', 'route_ner_yks_via_bam', 'rail', 'RZD (AYaM)', 'tommot_railway', 'nizhny_bestyakh_railway',
 '2025-11-22 12:00:00', '2025-11-22 20:00:00', 3200.00, 28800000000000, 300, 88.0, 450, 2),
('seg_ner_yks_bam_3', 'route_ner_yks_via_bam', 'bus', 'Ferry/Ice Road', 'nizhny_bestyakh_railway', 'yakutsk_bus',
 '2025-11-22 21:00:00', '2025-11-22 22:00:00', 500.00, 3600000000000, 45, 80.0, 22, 3);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_ner_yks_via_bam', 'seg_ner_yks_bam_1', 'seg_ner_yks_bam_2',
 3600000000000, 300, false, true, 3600000000000, 1),
('route_ner_yks_via_bam', 'seg_ner_yks_bam_2', 'seg_ner_yks_bam_3',
 3600000000000, 200, false, true, 3600000000000, 2);

-- ============================================================================
-- 5. РЕЧНЫЕ МАРШРУТЫ (летняя навигация по Лене и притокам)
-- ============================================================================

-- Якутск → Жиганск (по Лене на север)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_zhigansk_river', 'Yakutsk', 'Zhigansk',
 '2025-06-20 10:00:00', '2025-06-22 08:00:00', 165600000000000,
 12000.00, 68.00, 600.00, true, ARRAY['river']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_zhigansk_river', 'route_yks_zhigansk_river', 'river', 'Lenskiye Zori', 'yakutsk_port', 'zhigansk_port',
 '2025-06-20 10:00:00', '2025-06-22 08:00:00', 12000.00, 165600000000000, 150, 68.0, 880, 1);

-- Якутск → Кюсюр (дальний север по Лене)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_kyusyur_river', 'Yakutsk', 'Kyusyur',
 '2025-06-21 09:00:00', '2025-06-25 15:00:00', 367200000000000,
 25000.00, 62.00, 1250.00, true, ARRAY['river']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_kyusyur_river', 'route_yks_kyusyur_river', 'river', 'Lenskiye Zori', 'yakutsk_port', 'kyusyur_port',
 '2025-06-21 09:00:00', '2025-06-25 15:00:00', 25000.00, 367200000000000, 120, 62.0, 1780, 1);

-- Олекминск → Витим (по притоку)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_olek_vitim_river', 'Olekminsk', 'Vitim',
 '2025-06-22 08:00:00', '2025-06-23 18:00:00', 122400000000000,
 7500.00, 70.00, 375.00, true, ARRAY['river']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_olek_vitim_river', 'route_olek_vitim_river', 'river', 'Sakha River Transport', 'olekminsk_port', 'vitim_port',
 '2025-06-22 08:00:00', '2025-06-23 18:00:00', 7500.00, 122400000000000, 80, 70.0, 420, 1);

-- Покровск → Сангар (короткий речной маршрут)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_pokrovsk_sangur_river', 'Pokrovsk', 'Sangur',
 '2025-06-20 14:00:00', '2025-06-21 02:00:00', 43200000000000,
 3500.00, 75.00, 175.00, true, ARRAY['river']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_pokrovsk_sangur_river', 'route_pokrovsk_sangur_river', 'river', 'Sakha River Transport', 'pokrovsk_port', 'sangur_port',
 '2025-06-20 14:00:00', '2025-06-21 02:00:00', 3500.00, 43200000000000, 100, 75.0, 240, 1);

-- ============================================================================
-- 6. КОМПЛЕКСНЫЕ МУЛЬТИМОДАЛЬНЫЕ МАРШРУТЫ
-- ============================================================================

-- Москва → Тында → Якутск (авиа + ЖД + автобус)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_yks_via_tynda', 'Moscow', 'Yakutsk',
 '2025-11-21 10:00:00', '2025-11-24 08:30:00', 253800000000000,
 28000.00, 75.00, 1400.00, true, ARRAY['air','rail','bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_tynda_1', 'route_msk_yks_via_tynda', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-21 10:00:00', '2025-11-21 18:30:00', 15000.00, 30600000000000, 150, 92.0, 2100, 1),
('seg_msk_tynda_2', 'route_msk_yks_via_tynda', 'taxi', 'Yandex Taxi', 'yakutsk_yks', 'yakutsk_bus',
 '2025-11-21 19:00:00', '2025-11-21 19:30:00', 400.00, 1800000000000, 4, 95.0, 8, 2),
('seg_msk_tynda_3', 'route_msk_yks_via_tynda', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'nizhny_bestyakh_railway',
 '2025-11-21 20:00:00', '2025-11-21 21:30:00', 600.00, 5400000000000, 45, 80.0, 22, 3),
('seg_msk_tynda_4', 'route_msk_yks_via_tynda', 'rail', 'RZD (AYaM)', 'nizhny_bestyakh_railway', 'tynda_railway',
 '2025-11-22 08:00:00', '2025-11-23 14:00:00', 8500.00, 108000000000000, 300, 82.0, 1250, 4),
('seg_msk_tynda_5', 'route_msk_yks_via_tynda', 'rail', 'RZD (AYaM)', 'tynda_railway', 'nizhny_bestyakh_railway',
 '2025-11-23 16:00:00', '2025-11-24 06:00:00', 8500.00, 50400000000000, 300, 82.0, 1250, 5),
('seg_msk_tynda_6', 'route_msk_yks_via_tynda', 'bus', 'Ferry/Ice Road', 'nizhny_bestyakh_railway', 'yakutsk_bus',
 '2025-11-24 07:00:00', '2025-11-24 08:30:00', 600.00, 5400000000000, 45, 78.0, 22, 6);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_yks_via_tynda', 'seg_msk_tynda_1', 'seg_msk_tynda_2',
 1800000000000, 15000, true, true, 1800000000000, 1),
('route_msk_yks_via_tynda', 'seg_msk_tynda_2', 'seg_msk_tynda_3',
 1800000000000, 200, false, true, 1800000000000, 2),
('route_msk_yks_via_tynda', 'seg_msk_tynda_3', 'seg_msk_tynda_4',
 37800000000000, 100, false, true, 37800000000000, 3),
('route_msk_yks_via_tynda', 'seg_msk_tynda_4', 'seg_msk_tynda_5',
 7200000000000, 50, false, true, 7200000000000, 4),
('route_msk_yks_via_tynda', 'seg_msk_tynda_5', 'seg_msk_tynda_6',
 3600000000000, 200, false, true, 3600000000000, 5);

-- ============================================================================
-- 7. ДОПОЛНИТЕЛЬНЫЕ ОСТАНОВКИ ДЛЯ РАСШИРЕНИЯ СЕТИ
-- ============================================================================

INSERT INTO stops (id, name, city, latitude, longitude, stop_type) VALUES
-- Дальневосточные города
('magadan_aprt', 'Magadan Sokol Airport', 'Magadan', 59.9103, 150.7203, 'airport'),
('khabarovsk_aprt', 'Khabarovsk Novy Airport', 'Khabarovsk', 48.5280, 135.1883, 'airport'),

-- Дополнительные аэропорты Якутии
('suntar_aprt', 'Suntar Airport', 'Suntar', 62.1731, 117.6356, 'airport'),
('olekminsk_aprt', 'Olekminsk Airport', 'Olekminsk', 60.3733, 120.4272, 'airport'),
('lensk_aprt', 'Lensk Airport', 'Lensky', 60.7458, 114.8833, 'airport'),

-- Московские вокзалы и аэропорты
('moscow_vko', 'Vnukovo International Airport', 'Moscow', 55.5914, 37.2615, 'airport'),

-- Сибирские хабы
('krasnoyarsk_aprt', 'Krasnoyarsk Yemelyanovo Airport', 'Krasnoyarsk', 56.1729, 92.4933, 'airport'),
('irkutsk_aprt', 'Irkutsk International Airport', 'Irkutsk', 52.2680, 104.3889, 'airport'),
('novosibirsk_aprt', 'Novosibirsk Tolmachevo Airport', 'Novosibirsk', 55.0125, 82.6506, 'airport'),

-- Дополнительные автовокзалы
('vilyuysk_bus', 'Vilyuysk Bus Terminal', 'Vilyuysk', 63.7486, 121.5675, 'station'),
('nyurba_bus', 'Nyurba Bus Terminal', 'Nyurba', 65.3875, 118.4778, 'station'),
('suntar_bus', 'Suntar Bus Terminal', 'Suntar', 62.1731, 117.6356, 'station')

ON CONFLICT (name, city) DO NOTHING;

-- ============================================================================
-- 7. МОСКВА → РЕГИОНАЛЬНЫЕ ГОРОДА ЯКУТИИ (прямые и с пересадками)
-- ============================================================================

-- Москва → Мирный (прямой рейс)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_mirny_direct', 'Moscow', 'Mirny',
 '2025-11-23 09:00:00', '2025-11-23 18:30:00', 34200000000000,
 38000.00, 88.00, 1900.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_mirny_direct', 'route_msk_mirny_direct', 'air', 'ALROSA Air', 'moscow_dme', 'mirny_aprt',
 '2025-11-23 09:00:00', '2025-11-23 18:30:00', 38000.00, 34200000000000, 120, 88.0, 4420, 1);

-- Москва → Мирный (через Якутск - дешевле)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_mirny_via_yks', 'Moscow', 'Mirny',
 '2025-11-23 10:30:00', '2025-11-23 22:00:00', 41400000000000,
 35000.00, 85.00, 1750.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_mirny_yks_1', 'route_msk_mirny_via_yks', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-23 10:30:00', '2025-11-23 19:15:00', 28000.00, 31500000000000, 180, 88.0, 4100, 1),
('seg_msk_mirny_yks_2', 'route_msk_mirny_via_yks', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-23 20:30:00', '2025-11-23 22:00:00', 7000.00, 5400000000000, 100, 82.0, 520, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_mirny_via_yks', 'seg_msk_mirny_yks_1', 'seg_msk_mirny_yks_2',
 4500000000000, 500, false, true, 4500000000000, 1);

-- Москва → Нерюнгри (через Якутск)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_ner_via_yks', 'Moscow', 'Neryungri',
 '2025-11-23 08:00:00', '2025-11-23 19:00:00', 39600000000000,
 36000.00, 87.00, 1800.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_ner_yks_1', 'route_msk_ner_via_yks', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-23 08:00:00', '2025-11-23 16:30:00', 32500.00, 31800000000000, 150, 92.0, 4100, 1),
('seg_msk_ner_yks_2', 'route_msk_ner_via_yks', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'nerungri_aprt',
 '2025-11-23 17:30:00', '2025-11-23 19:00:00', 3500.00, 5400000000000, 120, 82.0, 560, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_ner_via_yks', 'seg_msk_ner_yks_1', 'seg_msk_ner_yks_2',
 3600000000000, 300, false, true, 3600000000000, 1);

-- Москва → Удачный (через Якутск)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_udachny', 'Moscow', 'Udachny',
 '2025-11-23 08:00:00', '2025-11-23 21:00:00', 46800000000000,
 42000.00, 84.00, 2100.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_udachny_1', 'route_msk_udachny', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-23 08:00:00', '2025-11-23 16:30:00', 32500.00, 31800000000000, 150, 92.0, 4100, 1),
('seg_msk_udachny_2', 'route_msk_udachny', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-11-23 18:45:00', '2025-11-23 21:00:00', 9500.00, 8100000000000, 100, 76.0, 630, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_udachny', 'seg_msk_udachny_1', 'seg_msk_udachny_2',
 8100000000000, 300, false, true, 8100000000000, 1);

-- ============================================================================
-- 8. ОБРАТНЫЕ МАРШРУТЫ (Якутия → Москва)
-- ============================================================================

-- Якутск → Москва (ранний рейс)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_msk_morning', 'Yakutsk', 'Moscow',
 '2025-11-24 06:00:00', '2025-11-24 14:00:00', 28800000000000,
 33000.00, 92.00, 1650.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_msk_morning', 'route_yks_msk_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-24 06:00:00', '2025-11-24 14:00:00', 33000.00, 28800000000000, 150, 92.0, 4100, 1);

-- Якутск → Москва (вечерний рейс)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_msk_evening', 'Yakutsk', 'Moscow',
 '2025-11-24 18:00:00', '2025-11-25 02:00:00', 28800000000000,
 29000.00, 90.00, 1450.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_msk_evening', 'route_yks_msk_evening', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'moscow_svo',
 '2025-11-24 18:00:00', '2025-11-25 02:00:00', 29000.00, 28800000000000, 180, 90.0, 4100, 1);

-- Мирный → Москва (прямой)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_mirny_msk_direct', 'Mirny', 'Moscow',
 '2025-11-24 10:00:00', '2025-11-24 19:00:00', 32400000000000,
 39000.00, 87.00, 1950.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_mirny_msk_direct', 'route_mirny_msk_direct', 'air', 'ALROSA Air', 'mirny_aprt', 'moscow_dme',
 '2025-11-24 10:00:00', '2025-11-24 19:00:00', 39000.00, 32400000000000, 120, 87.0, 4420, 1);

-- ============================================================================
-- 9. МЕЖРЕГИОНАЛЬНЫЕ МАРШРУТЫ ВНУТРИ ЯКУТИИ
-- ============================================================================

-- Мирный → Нерюнгри (авиа)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_mirny_ner_air', 'Mirny', 'Neryungri',
 '2025-11-23 12:00:00', '2025-11-23 13:45:00', 6300000000000,
 14000.00, 85.00, 700.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_mirny_ner_air', 'route_mirny_ner_air', 'air', 'Yakutia Airlines', 'mirny_aprt', 'nerungri_aprt',
 '2025-11-23 12:00:00', '2025-11-23 13:45:00', 14000.00, 6300000000000, 100, 85.0, 580, 1);

-- Удачный → Мирный → Якутск (комбо)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_udachny_yks_via_mirny', 'Udachny', 'Yakutsk',
 '2025-11-23 08:00:00', '2025-11-23 15:30:00', 27000000000000,
 18500.00, 82.00, 925.00, true, ARRAY['bus','air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_udachny_yks_1', 'route_udachny_yks_via_mirny', 'bus', 'ALROSA Transport', 'udachny_aprt', 'mirny_bus',
 '2025-11-23 08:00:00', '2025-11-23 14:00:00', 2800.00, 21600000000000, 40, 72.0, 260, 1),
('seg_udachny_yks_2', 'route_udachny_yks_via_mirny', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-23 14:30:00', '2025-11-23 15:30:00', 15700.00, 3600000000000, 100, 92.0, 520, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_udachny_yks_via_mirny', 'seg_udachny_yks_1', 'seg_udachny_yks_2',
 1800000000000, 1500, true, true, 1800000000000, 1);

-- Алдан → Якутск (автобус)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_aldan_yks_bus', 'Aldan', 'Yakutsk',
 '2025-11-23 07:00:00', '2025-11-23 17:00:00', 36000000000000,
 3300.00, 78.00, 165.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_aldan_yks_bus', 'route_aldan_yks_bus', 'bus', 'Magistral', 'aldan_bus', 'yakutsk_bus',
 '2025-11-23 07:00:00', '2025-11-23 17:00:00', 3300.00, 36000000000000, 45, 78.0, 380, 1);

-- Вилюйск → Нюрба (автобус, между алмазными районами)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_vilyuysk_nyurba_bus', 'Vilyuysk', 'Nyurba',
 '2025-11-23 09:00:00', '2025-11-23 14:00:00', 18000000000000,
 1800.00, 74.00, 90.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_vilyuysk_nyurba_bus', 'route_vilyuysk_nyurba_bus', 'bus', 'Avtotrans Yakutia', 'vilyuysk_bus', 'nyurba_bus',
 '2025-11-23 09:00:00', '2025-11-23 14:00:00', 1800.00, 18000000000000, 35, 74.0, 180, 1);

-- Вилюйск → Якутск (авиа + автобус варианты)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_vilyuysk_yks_bus', 'Vilyuysk', 'Yakutsk',
 '2025-11-23 08:00:00', '2025-11-23 18:00:00', 36000000000000,
 2800.00, 70.00, 140.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_vilyuysk_yks_bus', 'route_vilyuysk_yks_bus', 'bus', 'Avtotrans Yakutia', 'vilyuysk_bus', 'yakutsk_bus',
 '2025-11-23 08:00:00', '2025-11-23 18:00:00', 2800.00, 36000000000000, 40, 70.0, 350, 1);

-- ============================================================================
-- 10. МАРШРУТЫ ЧЕРЕЗ СИБИРСКИЕ ХАБЫ
-- ============================================================================

-- Москва → Якутск через Новосибирск (дешевый вариант)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_yks_via_nsk', 'Moscow', 'Yakutsk',
 '2025-11-23 07:00:00', '2025-11-23 23:00:00', 57600000000000,
 24000.00, 82.00, 1200.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_nsk_1', 'route_msk_yks_via_nsk', 'air', 'S7 Airlines', 'moscow_dme', 'novosibirsk_aprt',
 '2025-11-23 07:00:00', '2025-11-23 12:00:00', 12000.00, 18000000000000, 180, 90.0, 2800, 1),
('seg_msk_yks_nsk_2', 'route_msk_yks_via_nsk', 'air', 'S7 Airlines', 'novosibirsk_aprt', 'yakutsk_yks',
 '2025-11-23 14:00:00', '2025-11-23 23:00:00', 12000.00, 32400000000000, 180, 85.0, 3200, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_yks_via_nsk', 'seg_msk_yks_nsk_1', 'seg_msk_yks_nsk_2',
 7200000000000, 500, false, true, 7200000000000, 1);

-- Москва → Якутск через Красноярск
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_yks_via_krsk', 'Moscow', 'Yakutsk',
 '2025-11-23 08:00:00', '2025-11-24 01:00:00', 61200000000000,
 26000.00, 80.00, 1300.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_krsk_1', 'route_msk_yks_via_krsk', 'air', 'S7 Airlines', 'moscow_dme', 'krasnoyarsk_aprt',
 '2025-11-23 08:00:00', '2025-11-23 13:30:00', 13000.00, 19800000000000, 180, 88.0, 3350, 1),
('seg_msk_yks_krsk_2', 'route_msk_yks_via_krsk', 'air', 'Yakutia Airlines', 'krasnoyarsk_aprt', 'yakutsk_yks',
 '2025-11-23 16:00:00', '2025-11-24 01:00:00', 13000.00, 32400000000000, 150, 82.0, 2450, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_yks_via_krsk', 'seg_msk_yks_krsk_1', 'seg_msk_yks_krsk_2',
 9000000000000, 600, false, true, 9000000000000, 1);

-- Москва → Якутск через Иркутск
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_yks_via_irkutsk', 'Moscow', 'Yakutsk',
 '2025-11-23 09:00:00', '2025-11-24 02:30:00', 63000000000000,
 27000.00, 83.00, 1350.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_irk_1', 'route_msk_yks_via_irkutsk', 'air', 'S7 Airlines', 'moscow_dme', 'irkutsk_aprt',
 '2025-11-23 09:00:00', '2025-11-23 15:00:00', 14000.00, 21600000000000, 180, 90.0, 4200, 1),
('seg_msk_yks_irk_2', 'route_msk_yks_via_irkutsk', 'air', 'Yakutia Airlines', 'irkutsk_aprt', 'yakutsk_yks',
 '2025-11-23 18:00:00', '2025-11-24 02:30:00', 13000.00, 30600000000000, 150, 85.0, 1950, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_msk_yks_via_irkutsk', 'seg_msk_yks_irk_1', 'seg_msk_yks_irk_2',
 10800000000000, 700, false, true, 10800000000000, 1);

-- ============================================================================
-- 11. ДОПОЛНИТЕЛЬНЫЕ ВНУТРЕННИЕ МАРШРУТЫ С РАЗНЫМ ВРЕМЕНЕМ
-- ============================================================================

-- Якутск → Мирный (вечерний рейс)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_mirny_evening', 'Yakutsk', 'Mirny',
 '2025-11-23 17:00:00', '2025-11-23 18:30:00', 5400000000000,
 12500.00, 90.00, 625.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_evening', 'route_yks_mirny_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-23 17:00:00', '2025-11-23 18:30:00', 12500.00, 5400000000000, 100, 90.0, 520, 1);

-- Якутск → Вилюйск (вечерний)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_vilyuysk_evening', 'Yakutsk', 'Vilyuysk',
 '2025-11-23 16:00:00', '2025-11-23 17:40:00', 6000000000000,
 10500.00, 76.00, 525.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_vilyuysk_evening', 'route_yks_vilyuysk_evening', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-23 16:00:00', '2025-11-23 17:40:00', 10500.00, 6000000000000, 70, 76.0, 460, 1);

-- Якутск → Покровск (утренний и вечерний автобусы)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_pokrovsk_evening', 'Yakutsk', 'Pokrovsk',
 '2025-11-23 17:00:00', '2025-11-23 19:30:00', 9000000000000,
 1250.00, 83.00, 63.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_pokrovsk_evening', 'route_yks_pokrovsk_evening', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-23 17:00:00', '2025-11-23 19:30:00', 1250.00, 9000000000000, 45, 83.0, 80, 1);

-- ============================================================================
-- 12. ЛЕТНИЕ РЕЧНЫЕ МАРШРУТЫ (расширенные)
-- ============================================================================

-- Якутск → Баtagay (по Лене, дальний север)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_batagay_river', 'Yakutsk', 'Batagay',
 '2025-06-22 09:00:00', '2025-06-24 15:00:00', 194400000000000,
 16000.00, 65.00, 800.00, true, ARRAY['river']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_batagay_river', 'route_yks_batagay_river', 'river', 'Lenskiye Zori', 'yakutsk_port', 'batagay_port',
 '2025-06-22 09:00:00', '2025-06-24 15:00:00', 16000.00, 194400000000000, 100, 65.0, 1200, 1);

-- Ленский → Витим (речной маршрут по притоку)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_lensky_vitim_river', 'Lensky', 'Vitim',
 '2025-06-23 08:00:00', '2025-06-24 14:00:00', 108000000000000,
 6500.00, 68.00, 325.00, true, ARRAY['river']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_lensky_vitim_river', 'route_lensky_vitim_river', 'river', 'Sakha River Transport', 'lensky_port', 'vitim_port',
 '2025-06-23 08:00:00', '2025-06-24 14:00:00', 6500.00, 108000000000000, 80, 68.0, 380, 1);

-- ============================================================================
-- 13. КОМБО-МАРШРУТЫ АВТОБУС + АВИА
-- ============================================================================

-- Нерюнгри → Мирный (автобус + авиа через Якутск)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_ner_mirny_combo', 'Neryungri', 'Mirny',
 '2025-11-23 10:00:00', '2025-11-24 01:00:00', 54000000000000,
 16000.00, 75.00, 800.00, true, ARRAY['bus','air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_ner_mirny_1', 'route_ner_mirny_combo', 'bus', 'Amur Transport', 'nerungri_bus', 'yakutsk_bus',
 '2025-11-23 10:00:00', '2025-11-23 20:00:00', 3500.00, 36000000000000, 50, 75.0, 350, 1),
('seg_ner_mirny_2', 'route_ner_mirny_combo', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-23 23:00:00', '2025-11-24 01:00:00', 12500.00, 7200000000000, 100, 90.0, 520, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_ner_mirny_combo', 'seg_ner_mirny_1', 'seg_ner_mirny_2',
 10800000000000, 8000, true, true, 10800000000000, 1);

-- Алдан → Мирный (через Нерюнгри и Якутск)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_aldan_mirny_complex', 'Aldan', 'Mirny',
 '2025-11-23 09:00:00', '2025-11-24 03:00:00', 64800000000000,
 19000.00, 72.00, 950.00, true, ARRAY['bus','air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_aldan_mirny_1', 'route_aldan_mirny_complex', 'bus', 'Magistral', 'aldan_bus', 'nerungri_bus',
 '2025-11-23 09:00:00', '2025-11-23 13:00:00', 1200.00, 14400000000000, 45, 88.0, 130, 1),
('seg_aldan_mirny_2', 'route_aldan_mirny_complex', 'bus', 'Amur Transport', 'nerungri_bus', 'yakutsk_bus',
 '2025-11-23 14:00:00', '2025-11-24 00:00:00', 3500.00, 36000000000000, 50, 75.0, 350, 2),
('seg_aldan_mirny_3', 'route_aldan_mirny_complex', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-24 02:00:00', '2025-11-24 03:00:00', 14300.00, 3600000000000, 100, 90.0, 520, 3);

INSERT INTO connections (route_id, from_segment_id, to_segment_id,
                        transfer_duration, transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('route_aldan_mirny_complex', 'seg_aldan_mirny_1', 'seg_aldan_mirny_2',
 3600000000000, 200, false, true, 3600000000000, 1),
('route_aldan_mirny_complex', 'seg_aldan_mirny_2', 'seg_aldan_mirny_3',
 7200000000000, 8000, true, true, 7200000000000, 2);

-- ============================================================================
-- 14. НОЧНЫЕ МАРШРУТЫ
-- ============================================================================

-- Москва → Якутск (ночной рейс)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_yks_night', 'Moscow', 'Yakutsk',
 '2025-11-23 23:00:00', '2025-11-24 07:30:00', 30600000000000,
 27000.00, 89.00, 1350.00, true, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_night', 'route_msk_yks_night', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-23 23:00:00', '2025-11-24 07:30:00', 27000.00, 30600000000000, 180, 89.0, 4100, 1);

-- Якутск → Нерюнгри (ночной автобус)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_nerungri_night', 'Yakutsk', 'Neryungri',
 '2025-11-23 22:00:00', '2025-11-24 08:00:00', 36000000000000,
 3200.00, 73.00, 160.00, true, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_nerungri_night', 'route_yks_nerungri_night', 'bus', 'Amur Transport', 'yakutsk_bus', 'nerungri_bus',
 '2025-11-23 22:00:00', '2025-11-24 08:00:00', 3200.00, 36000000000000, 50, 73.0, 350, 1);

-- ============================================================================
-- ИТОГО ДОБАВЛЕНО В ЭТОМ ОБНОВЛЕНИИ:
-- - Остановки: 13 новых (Магадан, Хабаровск, Сунтар, Олекминск, Ленск, + сибирские хабы)
-- - Маршруты: 35+ новых
-- - Сегменты: ~50+
-- - Коннекции: для всех мультимодальных маршрутов
--
-- ВСЕГО В БАЗЕ (с учетом предыдущих):
-- - Остановки: 59+
-- - Маршруты: 71+
-- - Покрытие: вся Якутия + связи с крупными городами России
-- ============================================================================
