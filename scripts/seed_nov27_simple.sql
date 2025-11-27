-- Seed data for November 27, 2025
BEGIN;

-- Delete old data
DELETE FROM connections;
DELETE FROM segments;
DELETE FROM routes;
DELETE FROM stops WHERE city IN ('moscow', 'yakutsk', 'mirny', 'neryungri', 'udachny');

-- Add stops
INSERT INTO stops (id, name, city, city_display_name, latitude, longitude, stop_type) VALUES
('moscow_dme', 'Domodedovo Airport', 'moscow', 'Москва', 55.4088, 37.9063, 'airport'),
('moscow_svo', 'Sheremetyevo Airport', 'moscow', 'Москва', 55.9726, 37.4147, 'airport'),
('yakutsk_yks', 'Yakutsk Airport', 'yakutsk', 'Якутск', 62.0932, 129.7708, 'airport'),
('yakutsk_bus', 'Yakutsk Bus Terminal', 'yakutsk', 'Якутск', 62.0339, 129.7331, 'station'),
('mirny_aprt', 'Mirny Airport', 'mirny', 'Мирный', 62.5347, 114.0389, 'airport'),
('nerungri_aprt', 'Neryungri Airport', 'neryungri', 'Нерюнгри', 56.9139, 124.9144, 'airport'),
('udachny_aprt', 'Udachny Airport', 'udachny', 'Удачный', 66.4000, 112.0333, 'airport');

-- Route 1: Moscow -> Yakutsk (morning flight)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration, total_price, reliability_score, transport_types, saved_at) VALUES
('msk_yks_morning_1127', 'moscow', 'yakutsk', '2025-11-27 08:00:00+03', '2025-11-27 23:00:00+09', 28800000000000, 32000, 92, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id, departure_time, arrival_time, price, duration, seat_count, reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_morning_1127', 'msk_yks_morning_1127', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks', '2025-11-27 08:00:00+03', '2025-11-27 23:00:00+09', 32000, 28800000000000, 180, 92, 4884, 1);

-- Route 2: Moscow -> Yakutsk (evening flight)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration, total_price, reliability_score, transport_types, saved_at) VALUES
('msk_yks_evening_1127', 'moscow', 'yakutsk', '2025-11-27 18:00:00+03', '2025-11-28 09:00:00+09', 28800000000000, 28000, 90, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id, departure_time, arrival_time, price, duration, seat_count, reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_evening_1127', 'msk_yks_evening_1127', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks', '2025-11-27 18:00:00+03', '2025-11-28 09:00:00+09', 28000, 28800000000000, 200, 90, 4884, 1);

-- Route 3: Yakutsk -> Moscow (morning flight)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration, total_price, reliability_score, transport_types, saved_at) VALUES
('yks_msk_morning_1127', 'yakutsk', 'moscow', '2025-11-27 06:00:00+09', '2025-11-27 14:00:00+03', 28800000000000, 34000, 91, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id, departure_time, arrival_time, price, duration, seat_count, reliability_rate, distance, sequence_order) VALUES
('seg_yks_msk_morning_1127', 'yks_msk_morning_1127', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme', '2025-11-27 06:00:00+09', '2025-11-27 14:00:00+03', 34000, 28800000000000, 180, 91, 4884, 1);

-- Route 4: Yakutsk -> Mirny
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration, total_price, reliability_score, transport_types, saved_at) VALUES
('yks_mirny_1127', 'yakutsk', 'mirny', '2025-11-27 09:00:00+09', '2025-11-27 10:30:00+09', 5400000000000, 12000, 88, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id, departure_time, arrival_time, price, duration, seat_count, reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_1127', 'yks_mirny_1127', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt', '2025-11-27 09:00:00+09', '2025-11-27 10:30:00+09', 12000, 5400000000000, 100, 88, 520, 1);

-- Route 5: Yakutsk -> Neryungri
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration, total_price, reliability_score, transport_types, saved_at) VALUES
('yks_ner_1127', 'yakutsk', 'neryungri', '2025-11-27 10:00:00+09', '2025-11-27 11:30:00+09', 5400000000000, 14000, 87, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id, departure_time, arrival_time, price, duration, seat_count, reliability_rate, distance, sequence_order) VALUES
('seg_yks_ner_1127', 'yks_ner_1127', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'nerungri_aprt', '2025-11-27 10:00:00+09', '2025-11-27 11:30:00+09', 14000, 5400000000000, 120, 87, 560, 1);

-- Route 6: Yakutsk -> Udachny
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration, total_price, reliability_score, transport_types, saved_at) VALUES
('yks_udachny_1127', 'yakutsk', 'udachny', '2025-11-27 11:30:00+09', '2025-11-27 13:45:00+09', 8100000000000, 16000, 85, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id, departure_time, arrival_time, price, duration, seat_count, reliability_rate, distance, sequence_order) VALUES
('seg_yks_udachny_1127', 'yks_udachny_1127', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt', '2025-11-27 11:30:00+09', '2025-11-27 13:45:00+09', 16000, 8100000000000, 80, 85, 630, 1);

COMMIT;
