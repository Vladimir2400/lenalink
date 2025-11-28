-- ============================================================================
-- WEEKLY SEED DATA: NOVEMBER 28 - DECEMBER 4, 2025
-- LenaLink - Multi-modal Transport Aggregator for Yakutia
-- ============================================================================
-- This script generates routes for 7 days to enable realistic search testing
-- Based on seed_2025_11_27.sql pattern
-- ============================================================================

BEGIN;

-- ============================================================================
-- DAY 1: NOVEMBER 28, 2025 (Thursday)
-- ============================================================================

-- Moscow → Yakutsk (2 flights daily)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1128_morning', 'moscow', 'yakutsk', '2025-11-28 08:00:00', '2025-11-28 23:00:00',
 28800000000000, 32500.00, 92.00, 1625.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1128_evening', 'moscow', 'yakutsk', '2025-11-28 18:00:00', '2025-11-29 09:00:00',
 28800000000000, 28500.00, 90.00, 1425.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1128_m', 'msk_yks_1128_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-28 08:00:00', '2025-11-28 23:00:00', 32500.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1128_e', 'msk_yks_1128_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-28 18:00:00', '2025-11-29 09:00:00', 28500.00, 28800000000000, 200, 90.00, 4884, 1);

-- Yakutsk → Moscow (1 morning return)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_msk_1128_morning', 'yakutsk', 'moscow', '2025-11-28 06:00:00', '2025-11-28 14:00:00',
 28800000000000, 34500.00, 91.00, 1725.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_msk_1128_m', 'yks_msk_1128_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-28 06:00:00', '2025-11-28 14:00:00', 34500.00, 28800000000000, 180, 91.00, 4884, 1);

-- Yakutsk → Mirny (2 flights)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_mirny_1128_morning', 'yakutsk', 'mirny', '2025-11-28 09:00:00', '2025-11-28 10:30:00',
 5400000000000, 12200.00, 88.00, 610.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1128_evening', 'yakutsk', 'mirny', '2025-11-28 17:00:00', '2025-11-28 18:30:00',
 5400000000000, 13700.00, 90.00, 685.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_1128_m', 'yks_mirny_1128_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-28 09:00:00', '2025-11-28 10:30:00', 12200.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1128_e', 'yks_mirny_1128_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-28 17:00:00', '2025-11-28 18:30:00', 13700.00, 5400000000000, 100, 90.00, 520, 1);

-- Yakutsk → Neryungri
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_ner_1128', 'yakutsk', 'neryungri', '2025-11-28 10:00:00', '2025-11-28 11:30:00',
 5400000000000, 14200.00, 87.00, 710.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_ner_1128', 'yks_ner_1128', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-28 10:00:00', '2025-11-28 11:30:00', 14200.00, 5400000000000, 120, 87.00, 560, 1);

-- Yakutsk → Udachny
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_udachny_1128', 'yakutsk', 'udachny', '2025-11-28 11:30:00', '2025-11-28 13:45:00',
 8100000000000, 16200.00, 85.00, 810.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_udachny_1128', 'yks_udachny_1128', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-11-28 11:30:00', '2025-11-28 13:45:00', 16200.00, 8100000000000, 80, 85.00, 630, 1);

-- Yakutsk → Vilyuysk
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_vilyuysk_1128', 'yakutsk', 'vilyuysk', '2025-11-28 14:00:00', '2025-11-28 15:40:00',
 6000000000000, 11200.00, 82.00, 560.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_vilyuysk_1128', 'yks_vilyuysk_1128', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-28 14:00:00', '2025-11-28 15:40:00', 11200.00, 6000000000000, 70, 82.00, 460, 1);

-- Yakutsk → Tiksi
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_tiksi_1128', 'yakutsk', 'tiksi', '2025-11-28 08:30:00', '2025-11-28 11:00:00',
 9000000000000, 22200.00, 78.00, 1110.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_tiksi_1128', 'yks_tiksi_1128', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-11-28 08:30:00', '2025-11-28 11:00:00', 22200.00, 9000000000000, 60, 78.00, 1670, 1);

-- Yakutsk → Pokrovsk (bus)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_pokrovsk_1128', 'yakutsk', 'pokrovsk', '2025-11-28 08:00:00', '2025-11-28 10:30:00',
 9000000000000, 1250.00, 86.00, 62.50, false, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_pokrovsk_1128', 'yks_pokrovsk_1128', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-28 08:00:00', '2025-11-28 10:30:00', 1250.00, 9000000000000, 45, 86.00, 80, 1);

-- ============================================================================
-- DAY 2: NOVEMBER 29, 2025 (Friday)
-- ============================================================================

INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1129_morning', 'moscow', 'yakutsk', '2025-11-29 08:00:00', '2025-11-29 23:00:00',
 28800000000000, 33000.00, 92.00, 1650.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1129_evening', 'moscow', 'yakutsk', '2025-11-29 18:00:00', '2025-11-30 09:00:00',
 28800000000000, 29000.00, 90.00, 1450.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1129_morning', 'yakutsk', 'moscow', '2025-11-29 06:00:00', '2025-11-29 14:00:00',
 28800000000000, 35000.00, 91.00, 1750.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1129_morning', 'yakutsk', 'mirny', '2025-11-29 09:00:00', '2025-11-29 10:30:00',
 5400000000000, 12400.00, 88.00, 620.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1129_evening', 'yakutsk', 'mirny', '2025-11-29 17:00:00', '2025-11-29 18:30:00',
 5400000000000, 13900.00, 90.00, 695.00, false, ARRAY['air']::text[], NOW()),
('yks_ner_1129', 'yakutsk', 'neryungri', '2025-11-29 10:00:00', '2025-11-29 11:30:00',
 5400000000000, 14400.00, 87.00, 720.00, false, ARRAY['air']::text[], NOW()),
('yks_udachny_1129', 'yakutsk', 'udachny', '2025-11-29 11:30:00', '2025-11-29 13:45:00',
 8100000000000, 16400.00, 85.00, 820.00, false, ARRAY['air']::text[], NOW()),
('yks_vilyuysk_1129', 'yakutsk', 'vilyuysk', '2025-11-29 14:00:00', '2025-11-29 15:40:00',
 6000000000000, 11400.00, 82.00, 570.00, false, ARRAY['air']::text[], NOW()),
('yks_tiksi_1129', 'yakutsk', 'tiksi', '2025-11-29 08:30:00', '2025-11-29 11:00:00',
 9000000000000, 22400.00, 78.00, 1120.00, false, ARRAY['air']::text[], NOW()),
('yks_pokrovsk_1129', 'yakutsk', 'pokrovsk', '2025-11-29 08:00:00', '2025-11-29 10:30:00',
 9000000000000, 1300.00, 86.00, 65.00, false, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1129_m', 'msk_yks_1129_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-29 08:00:00', '2025-11-29 23:00:00', 33000.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1129_e', 'msk_yks_1129_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-29 18:00:00', '2025-11-30 09:00:00', 29000.00, 28800000000000, 200, 90.00, 4884, 1),
('seg_yks_msk_1129_m', 'yks_msk_1129_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-29 06:00:00', '2025-11-29 14:00:00', 35000.00, 28800000000000, 180, 91.00, 4884, 1),
('seg_yks_mirny_1129_m', 'yks_mirny_1129_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-29 09:00:00', '2025-11-29 10:30:00', 12400.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1129_e', 'yks_mirny_1129_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-29 17:00:00', '2025-11-29 18:30:00', 13900.00, 5400000000000, 100, 90.00, 520, 1),
('seg_yks_ner_1129', 'yks_ner_1129', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-29 10:00:00', '2025-11-29 11:30:00', 14400.00, 5400000000000, 120, 87.00, 560, 1),
('seg_yks_udachny_1129', 'yks_udachny_1129', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-11-29 11:30:00', '2025-11-29 13:45:00', 16400.00, 8100000000000, 80, 85.00, 630, 1),
('seg_yks_vilyuysk_1129', 'yks_vilyuysk_1129', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-29 14:00:00', '2025-11-29 15:40:00', 11400.00, 6000000000000, 70, 82.00, 460, 1),
('seg_yks_tiksi_1129', 'yks_tiksi_1129', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-11-29 08:30:00', '2025-11-29 11:00:00', 22400.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_yks_pokrovsk_1129', 'yks_pokrovsk_1129', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-29 08:00:00', '2025-11-29 10:30:00', 1300.00, 9000000000000, 45, 86.00, 80, 1);

-- ============================================================================
-- DAY 3: NOVEMBER 30, 2025 (Saturday)
-- ============================================================================

INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1130_morning', 'moscow', 'yakutsk', '2025-11-30 08:00:00', '2025-11-30 23:00:00',
 28800000000000, 33500.00, 92.00, 1675.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1130_evening', 'moscow', 'yakutsk', '2025-11-30 18:00:00', '2025-12-01 09:00:00',
 28800000000000, 29500.00, 90.00, 1475.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1130_morning', 'yakutsk', 'moscow', '2025-11-30 06:00:00', '2025-11-30 14:00:00',
 28800000000000, 35500.00, 91.00, 1775.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1130_morning', 'yakutsk', 'mirny', '2025-11-30 09:00:00', '2025-11-30 10:30:00',
 5400000000000, 12600.00, 88.00, 630.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1130_evening', 'yakutsk', 'mirny', '2025-11-30 17:00:00', '2025-11-30 18:30:00',
 5400000000000, 14100.00, 90.00, 705.00, false, ARRAY['air']::text[], NOW()),
('yks_ner_1130', 'yakutsk', 'neryungri', '2025-11-30 10:00:00', '2025-11-30 11:30:00',
 5400000000000, 14600.00, 87.00, 730.00, false, ARRAY['air']::text[], NOW()),
('yks_udachny_1130', 'yakutsk', 'udachny', '2025-11-30 11:30:00', '2025-11-30 13:45:00',
 8100000000000, 16600.00, 85.00, 830.00, false, ARRAY['air']::text[], NOW()),
('yks_vilyuysk_1130', 'yakutsk', 'vilyuysk', '2025-11-30 14:00:00', '2025-11-30 15:40:00',
 6000000000000, 11600.00, 82.00, 580.00, false, ARRAY['air']::text[], NOW()),
('yks_tiksi_1130', 'yakutsk', 'tiksi', '2025-11-30 08:30:00', '2025-11-30 11:00:00',
 9000000000000, 22600.00, 78.00, 1130.00, false, ARRAY['air']::text[], NOW()),
('yks_pokrovsk_1130', 'yakutsk', 'pokrovsk', '2025-11-30 08:00:00', '2025-11-30 10:30:00',
 9000000000000, 1350.00, 86.00, 67.50, false, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1130_m', 'msk_yks_1130_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-30 08:00:00', '2025-11-30 23:00:00', 33500.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1130_e', 'msk_yks_1130_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-30 18:00:00', '2025-12-01 09:00:00', 29500.00, 28800000000000, 200, 90.00, 4884, 1),
('seg_yks_msk_1130_m', 'yks_msk_1130_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-30 06:00:00', '2025-11-30 14:00:00', 35500.00, 28800000000000, 180, 91.00, 4884, 1),
('seg_yks_mirny_1130_m', 'yks_mirny_1130_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-30 09:00:00', '2025-11-30 10:30:00', 12600.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1130_e', 'yks_mirny_1130_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-30 17:00:00', '2025-11-30 18:30:00', 14100.00, 5400000000000, 100, 90.00, 520, 1),
('seg_yks_ner_1130', 'yks_ner_1130', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-30 10:00:00', '2025-11-30 11:30:00', 14600.00, 5400000000000, 120, 87.00, 560, 1),
('seg_yks_udachny_1130', 'yks_udachny_1130', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-11-30 11:30:00', '2025-11-30 13:45:00', 16600.00, 8100000000000, 80, 85.00, 630, 1),
('seg_yks_vilyuysk_1130', 'yks_vilyuysk_1130', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-30 14:00:00', '2025-11-30 15:40:00', 11600.00, 6000000000000, 70, 82.00, 460, 1),
('seg_yks_tiksi_1130', 'yks_tiksi_1130', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-11-30 08:30:00', '2025-11-30 11:00:00', 22600.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_yks_pokrovsk_1130', 'yks_pokrovsk_1130', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-30 08:00:00', '2025-11-30 10:30:00', 1350.00, 9000000000000, 45, 86.00, 80, 1);

-- ============================================================================
-- DAY 4: DECEMBER 1, 2025 (Sunday)
-- ============================================================================

INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1201_morning', 'moscow', 'yakutsk', '2025-12-01 08:00:00', '2025-12-01 23:00:00',
 28800000000000, 34000.00, 92.00, 1700.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1201_evening', 'moscow', 'yakutsk', '2025-12-01 18:00:00', '2025-12-02 09:00:00',
 28800000000000, 30000.00, 90.00, 1500.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1201_morning', 'yakutsk', 'moscow', '2025-12-01 06:00:00', '2025-12-01 14:00:00',
 28800000000000, 36000.00, 91.00, 1800.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1201_morning', 'yakutsk', 'mirny', '2025-12-01 09:00:00', '2025-12-01 10:30:00',
 5400000000000, 12800.00, 88.00, 640.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1201_evening', 'yakutsk', 'mirny', '2025-12-01 17:00:00', '2025-12-01 18:30:00',
 5400000000000, 14300.00, 90.00, 715.00, false, ARRAY['air']::text[], NOW()),
('yks_ner_1201', 'yakutsk', 'neryungri', '2025-12-01 10:00:00', '2025-12-01 11:30:00',
 5400000000000, 14800.00, 87.00, 740.00, false, ARRAY['air']::text[], NOW()),
('yks_udachny_1201', 'yakutsk', 'udachny', '2025-12-01 11:30:00', '2025-12-01 13:45:00',
 8100000000000, 16800.00, 85.00, 840.00, false, ARRAY['air']::text[], NOW()),
('yks_vilyuysk_1201', 'yakutsk', 'vilyuysk', '2025-12-01 14:00:00', '2025-12-01 15:40:00',
 6000000000000, 11800.00, 82.00, 590.00, false, ARRAY['air']::text[], NOW()),
('yks_tiksi_1201', 'yakutsk', 'tiksi', '2025-12-01 08:30:00', '2025-12-01 11:00:00',
 9000000000000, 22800.00, 78.00, 1140.00, false, ARRAY['air']::text[], NOW()),
('yks_pokrovsk_1201', 'yakutsk', 'pokrovsk', '2025-12-01 08:00:00', '2025-12-01 10:30:00',
 9000000000000, 1400.00, 86.00, 70.00, false, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1201_m', 'msk_yks_1201_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-12-01 08:00:00', '2025-12-01 23:00:00', 34000.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1201_e', 'msk_yks_1201_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-12-01 18:00:00', '2025-12-02 09:00:00', 30000.00, 28800000000000, 200, 90.00, 4884, 1),
('seg_yks_msk_1201_m', 'yks_msk_1201_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-12-01 06:00:00', '2025-12-01 14:00:00', 36000.00, 28800000000000, 180, 91.00, 4884, 1),
('seg_yks_mirny_1201_m', 'yks_mirny_1201_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-12-01 09:00:00', '2025-12-01 10:30:00', 12800.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1201_e', 'yks_mirny_1201_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-12-01 17:00:00', '2025-12-01 18:30:00', 14300.00, 5400000000000, 100, 90.00, 520, 1),
('seg_yks_ner_1201', 'yks_ner_1201', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-12-01 10:00:00', '2025-12-01 11:30:00', 14800.00, 5400000000000, 120, 87.00, 560, 1),
('seg_yks_udachny_1201', 'yks_udachny_1201', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-12-01 11:30:00', '2025-12-01 13:45:00', 16800.00, 8100000000000, 80, 85.00, 630, 1),
('seg_yks_vilyuysk_1201', 'yks_vilyuysk_1201', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-12-01 14:00:00', '2025-12-01 15:40:00', 11800.00, 6000000000000, 70, 82.00, 460, 1),
('seg_yks_tiksi_1201', 'yks_tiksi_1201', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-12-01 08:30:00', '2025-12-01 11:00:00', 22800.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_yks_pokrovsk_1201', 'yks_pokrovsk_1201', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-12-01 08:00:00', '2025-12-01 10:30:00', 1400.00, 9000000000000, 45, 86.00, 80, 1);

-- ============================================================================
-- DAY 5: DECEMBER 2, 2025 (Monday)
-- ============================================================================

INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1202_morning', 'moscow', 'yakutsk', '2025-12-02 08:00:00', '2025-12-02 23:00:00',
 28800000000000, 31500.00, 92.00, 1575.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1202_evening', 'moscow', 'yakutsk', '2025-12-02 18:00:00', '2025-12-03 09:00:00',
 28800000000000, 27500.00, 90.00, 1375.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1202_morning', 'yakutsk', 'moscow', '2025-12-02 06:00:00', '2025-12-02 14:00:00',
 28800000000000, 33500.00, 91.00, 1675.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1202_morning', 'yakutsk', 'mirny', '2025-12-02 09:00:00', '2025-12-02 10:30:00',
 5400000000000, 11800.00, 88.00, 590.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1202_evening', 'yakutsk', 'mirny', '2025-12-02 17:00:00', '2025-12-02 18:30:00',
 5400000000000, 13300.00, 90.00, 665.00, false, ARRAY['air']::text[], NOW()),
('yks_ner_1202', 'yakutsk', 'neryungri', '2025-12-02 10:00:00', '2025-12-02 11:30:00',
 5400000000000, 13800.00, 87.00, 690.00, false, ARRAY['air']::text[], NOW()),
('yks_udachny_1202', 'yakutsk', 'udachny', '2025-12-02 11:30:00', '2025-12-02 13:45:00',
 8100000000000, 15800.00, 85.00, 790.00, false, ARRAY['air']::text[], NOW()),
('yks_vilyuysk_1202', 'yakutsk', 'vilyuysk', '2025-12-02 14:00:00', '2025-12-02 15:40:00',
 6000000000000, 10800.00, 82.00, 540.00, false, ARRAY['air']::text[], NOW()),
('yks_tiksi_1202', 'yakutsk', 'tiksi', '2025-12-02 08:30:00', '2025-12-02 11:00:00',
 9000000000000, 21800.00, 78.00, 1090.00, false, ARRAY['air']::text[], NOW()),
('yks_pokrovsk_1202', 'yakutsk', 'pokrovsk', '2025-12-02 08:00:00', '2025-12-02 10:30:00',
 9000000000000, 1150.00, 86.00, 57.50, false, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1202_m', 'msk_yks_1202_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-12-02 08:00:00', '2025-12-02 23:00:00', 31500.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1202_e', 'msk_yks_1202_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-12-02 18:00:00', '2025-12-03 09:00:00', 27500.00, 28800000000000, 200, 90.00, 4884, 1),
('seg_yks_msk_1202_m', 'yks_msk_1202_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-12-02 06:00:00', '2025-12-02 14:00:00', 33500.00, 28800000000000, 180, 91.00, 4884, 1),
('seg_yks_mirny_1202_m', 'yks_mirny_1202_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-12-02 09:00:00', '2025-12-02 10:30:00', 11800.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1202_e', 'yks_mirny_1202_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-12-02 17:00:00', '2025-12-02 18:30:00', 13300.00, 5400000000000, 100, 90.00, 520, 1),
('seg_yks_ner_1202', 'yks_ner_1202', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-12-02 10:00:00', '2025-12-02 11:30:00', 13800.00, 5400000000000, 120, 87.00, 560, 1),
('seg_yks_udachny_1202', 'yks_udachny_1202', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-12-02 11:30:00', '2025-12-02 13:45:00', 15800.00, 8100000000000, 80, 85.00, 630, 1),
('seg_yks_vilyuysk_1202', 'yks_vilyuysk_1202', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-12-02 14:00:00', '2025-12-02 15:40:00', 10800.00, 6000000000000, 70, 82.00, 460, 1),
('seg_yks_tiksi_1202', 'yks_tiksi_1202', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-12-02 08:30:00', '2025-12-02 11:00:00', 21800.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_yks_pokrovsk_1202', 'yks_pokrovsk_1202', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-12-02 08:00:00', '2025-12-02 10:30:00', 1150.00, 9000000000000, 45, 86.00, 80, 1);

-- ============================================================================
-- DAY 6: DECEMBER 3, 2025 (Tuesday)
-- ============================================================================

INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1203_morning', 'moscow', 'yakutsk', '2025-12-03 08:00:00', '2025-12-03 23:00:00',
 28800000000000, 32200.00, 92.00, 1610.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1203_evening', 'moscow', 'yakutsk', '2025-12-03 18:00:00', '2025-12-04 09:00:00',
 28800000000000, 28200.00, 90.00, 1410.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1203_morning', 'yakutsk', 'moscow', '2025-12-03 06:00:00', '2025-12-03 14:00:00',
 28800000000000, 34200.00, 91.00, 1710.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1203_morning', 'yakutsk', 'mirny', '2025-12-03 09:00:00', '2025-12-03 10:30:00',
 5400000000000, 12100.00, 88.00, 605.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1203_evening', 'yakutsk', 'mirny', '2025-12-03 17:00:00', '2025-12-03 18:30:00',
 5400000000000, 13600.00, 90.00, 680.00, false, ARRAY['air']::text[], NOW()),
('yks_ner_1203', 'yakutsk', 'neryungri', '2025-12-03 10:00:00', '2025-12-03 11:30:00',
 5400000000000, 14100.00, 87.00, 705.00, false, ARRAY['air']::text[], NOW()),
('yks_udachny_1203', 'yakutsk', 'udachny', '2025-12-03 11:30:00', '2025-12-03 13:45:00',
 8100000000000, 16100.00, 85.00, 805.00, false, ARRAY['air']::text[], NOW()),
('yks_vilyuysk_1203', 'yakutsk', 'vilyuysk', '2025-12-03 14:00:00', '2025-12-03 15:40:00',
 6000000000000, 11100.00, 82.00, 555.00, false, ARRAY['air']::text[], NOW()),
('yks_tiksi_1203', 'yakutsk', 'tiksi', '2025-12-03 08:30:00', '2025-12-03 11:00:00',
 9000000000000, 22100.00, 78.00, 1105.00, false, ARRAY['air']::text[], NOW()),
('yks_pokrovsk_1203', 'yakutsk', 'pokrovsk', '2025-12-03 08:00:00', '2025-12-03 10:30:00',
 9000000000000, 1220.00, 86.00, 61.00, false, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1203_m', 'msk_yks_1203_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-12-03 08:00:00', '2025-12-03 23:00:00', 32200.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1203_e', 'msk_yks_1203_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-12-03 18:00:00', '2025-12-04 09:00:00', 28200.00, 28800000000000, 200, 90.00, 4884, 1),
('seg_yks_msk_1203_m', 'yks_msk_1203_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-12-03 06:00:00', '2025-12-03 14:00:00', 34200.00, 28800000000000, 180, 91.00, 4884, 1),
('seg_yks_mirny_1203_m', 'yks_mirny_1203_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-12-03 09:00:00', '2025-12-03 10:30:00', 12100.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1203_e', 'yks_mirny_1203_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-12-03 17:00:00', '2025-12-03 18:30:00', 13600.00, 5400000000000, 100, 90.00, 520, 1),
('seg_yks_ner_1203', 'yks_ner_1203', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-12-03 10:00:00', '2025-12-03 11:30:00', 14100.00, 5400000000000, 120, 87.00, 560, 1),
('seg_yks_udachny_1203', 'yks_udachny_1203', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-12-03 11:30:00', '2025-12-03 13:45:00', 16100.00, 8100000000000, 80, 85.00, 630, 1),
('seg_yks_vilyuysk_1203', 'yks_vilyuysk_1203', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-12-03 14:00:00', '2025-12-03 15:40:00', 11100.00, 6000000000000, 70, 82.00, 460, 1),
('seg_yks_tiksi_1203', 'yks_tiksi_1203', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-12-03 08:30:00', '2025-12-03 11:00:00', 22100.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_yks_pokrovsk_1203', 'yks_pokrovsk_1203', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-12-03 08:00:00', '2025-12-03 10:30:00', 1220.00, 9000000000000, 45, 86.00, 80, 1);

-- ============================================================================
-- DAY 7: DECEMBER 4, 2025 (Wednesday)
-- ============================================================================

INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1204_morning', 'moscow', 'yakutsk', '2025-12-04 08:00:00', '2025-12-04 23:00:00',
 28800000000000, 32800.00, 92.00, 1640.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1204_evening', 'moscow', 'yakutsk', '2025-12-04 18:00:00', '2025-12-05 09:00:00',
 28800000000000, 28800.00, 90.00, 1440.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1204_morning', 'yakutsk', 'moscow', '2025-12-04 06:00:00', '2025-12-04 14:00:00',
 28800000000000, 34800.00, 91.00, 1740.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1204_morning', 'yakutsk', 'mirny', '2025-12-04 09:00:00', '2025-12-04 10:30:00',
 5400000000000, 12300.00, 88.00, 615.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1204_evening', 'yakutsk', 'mirny', '2025-12-04 17:00:00', '2025-12-04 18:30:00',
 5400000000000, 13800.00, 90.00, 690.00, false, ARRAY['air']::text[], NOW()),
('yks_ner_1204', 'yakutsk', 'neryungri', '2025-12-04 10:00:00', '2025-12-04 11:30:00',
 5400000000000, 14300.00, 87.00, 715.00, false, ARRAY['air']::text[], NOW()),
('yks_udachny_1204', 'yakutsk', 'udachny', '2025-12-04 11:30:00', '2025-12-04 13:45:00',
 8100000000000, 16300.00, 85.00, 815.00, false, ARRAY['air']::text[], NOW()),
('yks_vilyuysk_1204', 'yakutsk', 'vilyuysk', '2025-12-04 14:00:00', '2025-12-04 15:40:00',
 6000000000000, 11300.00, 82.00, 565.00, false, ARRAY['air']::text[], NOW()),
('yks_tiksi_1204', 'yakutsk', 'tiksi', '2025-12-04 08:30:00', '2025-12-04 11:00:00',
 9000000000000, 22300.00, 78.00, 1115.00, false, ARRAY['air']::text[], NOW()),
('yks_pokrovsk_1204', 'yakutsk', 'pokrovsk', '2025-12-04 08:00:00', '2025-12-04 10:30:00',
 9000000000000, 1270.00, 86.00, 63.50, false, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1204_m', 'msk_yks_1204_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-12-04 08:00:00', '2025-12-04 23:00:00', 32800.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1204_e', 'msk_yks_1204_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-12-04 18:00:00', '2025-12-05 09:00:00', 28800.00, 28800000000000, 200, 90.00, 4884, 1),
('seg_yks_msk_1204_m', 'yks_msk_1204_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-12-04 06:00:00', '2025-12-04 14:00:00', 34800.00, 28800000000000, 180, 91.00, 4884, 1),
('seg_yks_mirny_1204_m', 'yks_mirny_1204_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-12-04 09:00:00', '2025-12-04 10:30:00', 12300.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1204_e', 'yks_mirny_1204_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-12-04 17:00:00', '2025-12-04 18:30:00', 13800.00, 5400000000000, 100, 90.00, 520, 1),
('seg_yks_ner_1204', 'yks_ner_1204', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-12-04 10:00:00', '2025-12-04 11:30:00', 14300.00, 5400000000000, 120, 87.00, 560, 1),
('seg_yks_udachny_1204', 'yks_udachny_1204', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-12-04 11:30:00', '2025-12-04 13:45:00', 16300.00, 8100000000000, 80, 85.00, 630, 1),
('seg_yks_vilyuysk_1204', 'yks_vilyuysk_1204', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-12-04 14:00:00', '2025-12-04 15:40:00', 11300.00, 6000000000000, 70, 82.00, 460, 1),
('seg_yks_tiksi_1204', 'yks_tiksi_1204', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-12-04 08:30:00', '2025-12-04 11:00:00', 22300.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_yks_pokrovsk_1204', 'yks_pokrovsk_1204', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-12-04 08:00:00', '2025-12-04 10:30:00', 1270.00, 9000000000000, 45, 86.00, 80, 1);

COMMIT;

-- ============================================================================
-- SUMMARY:
-- - 7 days of routes (November 28 - December 4, 2025)
-- - Each day includes:
--   * 2-3 Moscow ↔ Yakutsk flights (morning/evening)
--   * 8-9 Yakutsk → Yakutia cities routes (air + bus)
-- - Total: 70 routes, 70 segments
-- - Prices vary slightly by day (±5-10% variation for realism)
-- - Same schedule pattern as seed_2025_11_27.sql
-- ============================================================================
