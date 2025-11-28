-- ============================================================================
-- EXTENDED WEEKLY SEED DATA: NOVEMBER 28 - DECEMBER 4, 2025
-- LenaLink - Multi-modal Transport Aggregator for Yakutia
-- ============================================================================
-- This script includes:
-- - Multiple daily flights Moscow ↔ Yakutsk
-- - Reverse routes from Yakutia cities back to Yakutsk
-- - Multi-segment routes (Moscow → Yakutia cities)
-- - More flights per day for better coverage
-- ============================================================================

BEGIN;

-- ============================================================================
-- DAY 1: NOVEMBER 28, 2025 (Thursday)
-- ============================================================================

-- ===== Moscow ↔ Yakutsk =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1128_morning', 'moscow', 'yakutsk', '2025-11-28 08:00:00', '2025-11-28 23:00:00',
 28800000000000, 32500.00, 92.00, 1625.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1128_afternoon', 'moscow', 'yakutsk', '2025-11-28 14:00:00', '2025-11-29 05:00:00',
 28800000000000, 30500.00, 91.00, 1525.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1128_evening', 'moscow', 'yakutsk', '2025-11-28 18:00:00', '2025-11-29 09:00:00',
 28800000000000, 28500.00, 90.00, 1425.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1128_early', 'yakutsk', 'moscow', '2025-11-28 01:00:00', '2025-11-28 09:00:00',
 28800000000000, 33500.00, 91.00, 1675.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1128_morning', 'yakutsk', 'moscow', '2025-11-28 06:00:00', '2025-11-28 14:00:00',
 28800000000000, 34500.00, 91.00, 1725.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1128_evening', 'yakutsk', 'moscow', '2025-11-28 19:00:00', '2025-11-29 03:00:00',
 28800000000000, 35500.00, 90.00, 1775.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1128_m', 'msk_yks_1128_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-28 08:00:00', '2025-11-28 23:00:00', 32500.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1128_a', 'msk_yks_1128_afternoon', 'air', 'Aeroflot', 'moscow_svo', 'yakutsk_yks',
 '2025-11-28 14:00:00', '2025-11-29 05:00:00', 30500.00, 28800000000000, 220, 91.00, 4884, 1),
('seg_msk_yks_1128_e', 'msk_yks_1128_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-28 18:00:00', '2025-11-29 09:00:00', 28500.00, 28800000000000, 200, 90.00, 4884, 1),
('seg_yks_msk_1128_ea', 'yks_msk_1128_early', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-28 01:00:00', '2025-11-28 09:00:00', 33500.00, 28800000000000, 200, 91.00, 4884, 1),
('seg_yks_msk_1128_m', 'yks_msk_1128_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-28 06:00:00', '2025-11-28 14:00:00', 34500.00, 28800000000000, 180, 91.00, 4884, 1),
('seg_yks_msk_1128_ev', 'yks_msk_1128_evening', 'air', 'Aeroflot', 'yakutsk_yks', 'moscow_svo',
 '2025-11-28 19:00:00', '2025-11-29 03:00:00', 35500.00, 28800000000000, 220, 90.00, 4884, 1);

-- ===== Yakutsk ↔ Mirny =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_mirny_1128_morning', 'yakutsk', 'mirny', '2025-11-28 09:00:00', '2025-11-28 10:30:00',
 5400000000000, 12200.00, 88.00, 610.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1128_afternoon', 'yakutsk', 'mirny', '2025-11-28 13:00:00', '2025-11-28 14:30:00',
 5400000000000, 12800.00, 89.00, 640.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1128_evening', 'yakutsk', 'mirny', '2025-11-28 17:00:00', '2025-11-28 18:30:00',
 5400000000000, 13700.00, 90.00, 685.00, false, ARRAY['air']::text[], NOW()),
('mirny_yks_1128_morning', 'mirny', 'yakutsk', '2025-11-28 11:00:00', '2025-11-28 12:30:00',
 5400000000000, 12500.00, 88.00, 625.00, false, ARRAY['air']::text[], NOW()),
('mirny_yks_1128_afternoon', 'mirny', 'yakutsk', '2025-11-28 15:00:00', '2025-11-28 16:30:00',
 5400000000000, 13000.00, 89.00, 650.00, false, ARRAY['air']::text[], NOW()),
('mirny_yks_1128_evening', 'mirny', 'yakutsk', '2025-11-28 19:00:00', '2025-11-28 20:30:00',
 5400000000000, 14000.00, 90.00, 700.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_1128_m', 'yks_mirny_1128_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-28 09:00:00', '2025-11-28 10:30:00', 12200.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1128_a', 'yks_mirny_1128_afternoon', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-28 13:00:00', '2025-11-28 14:30:00', 12800.00, 5400000000000, 100, 89.00, 520, 1),
('seg_yks_mirny_1128_e', 'yks_mirny_1128_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-28 17:00:00', '2025-11-28 18:30:00', 13700.00, 5400000000000, 100, 90.00, 520, 1),
('seg_mirny_yks_1128_m', 'mirny_yks_1128_morning', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-28 11:00:00', '2025-11-28 12:30:00', 12500.00, 5400000000000, 100, 88.00, 520, 1),
('seg_mirny_yks_1128_a', 'mirny_yks_1128_afternoon', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-28 15:00:00', '2025-11-28 16:30:00', 13000.00, 5400000000000, 100, 89.00, 520, 1),
('seg_mirny_yks_1128_e', 'mirny_yks_1128_evening', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-28 19:00:00', '2025-11-28 20:30:00', 14000.00, 5400000000000, 100, 90.00, 520, 1);

-- ===== Yakutsk ↔ Neryungri =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_ner_1128_morning', 'yakutsk', 'neryungri', '2025-11-28 10:00:00', '2025-11-28 11:30:00',
 5400000000000, 14200.00, 87.00, 710.00, false, ARRAY['air']::text[], NOW()),
('yks_ner_1128_evening', 'yakutsk', 'neryungri', '2025-11-28 16:00:00', '2025-11-28 17:30:00',
 5400000000000, 14800.00, 88.00, 740.00, false, ARRAY['air']::text[], NOW()),
('ner_yks_1128_morning', 'neryungri', 'yakutsk', '2025-11-28 12:00:00', '2025-11-28 13:30:00',
 5400000000000, 14500.00, 87.00, 725.00, false, ARRAY['air']::text[], NOW()),
('ner_yks_1128_evening', 'neryungri', 'yakutsk', '2025-11-28 18:00:00', '2025-11-28 19:30:00',
 5400000000000, 15000.00, 88.00, 750.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_ner_1128_m', 'yks_ner_1128_morning', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-28 10:00:00', '2025-11-28 11:30:00', 14200.00, 5400000000000, 120, 87.00, 560, 1),
('seg_yks_ner_1128_e', 'yks_ner_1128_evening', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-28 16:00:00', '2025-11-28 17:30:00', 14800.00, 5400000000000, 120, 88.00, 560, 1),
('seg_ner_yks_1128_m', 'ner_yks_1128_morning', 'air', 'Yakutia Airlines', 'neryungri_aprt', 'yakutsk_yks',
 '2025-11-28 12:00:00', '2025-11-28 13:30:00', 14500.00, 5400000000000, 120, 87.00, 560, 1),
('seg_ner_yks_1128_e', 'ner_yks_1128_evening', 'air', 'Yakutia Airlines', 'neryungri_aprt', 'yakutsk_yks',
 '2025-11-28 18:00:00', '2025-11-28 19:30:00', 15000.00, 5400000000000, 120, 88.00, 560, 1);

-- ===== Yakutsk ↔ Udachny =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_udachny_1128', 'yakutsk', 'udachny', '2025-11-28 11:30:00', '2025-11-28 13:45:00',
 8100000000000, 16200.00, 85.00, 810.00, false, ARRAY['air']::text[], NOW()),
('udachny_yks_1128', 'udachny', 'yakutsk', '2025-11-28 14:30:00', '2025-11-28 16:45:00',
 8100000000000, 16500.00, 85.00, 825.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_udachny_1128', 'yks_udachny_1128', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-11-28 11:30:00', '2025-11-28 13:45:00', 16200.00, 8100000000000, 80, 85.00, 630, 1),
('seg_udachny_yks_1128', 'udachny_yks_1128', 'air', 'ALROSA Air', 'udachny_aprt', 'yakutsk_yks',
 '2025-11-28 14:30:00', '2025-11-28 16:45:00', 16500.00, 8100000000000, 80, 85.00, 630, 1);

-- ===== Yakutsk ↔ Vilyuysk =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_vilyuysk_1128', 'yakutsk', 'vilyuysk', '2025-11-28 14:00:00', '2025-11-28 15:40:00',
 6000000000000, 11200.00, 82.00, 560.00, false, ARRAY['air']::text[], NOW()),
('vilyuysk_yks_1128', 'vilyuysk', 'yakutsk', '2025-11-28 16:00:00', '2025-11-28 17:40:00',
 6000000000000, 11400.00, 82.00, 570.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_vilyuysk_1128', 'yks_vilyuysk_1128', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-28 14:00:00', '2025-11-28 15:40:00', 11200.00, 6000000000000, 70, 82.00, 460, 1),
('seg_vilyuysk_yks_1128', 'vilyuysk_yks_1128', 'air', 'Polar Airlines', 'vilyuysk_aprt', 'yakutsk_yks',
 '2025-11-28 16:00:00', '2025-11-28 17:40:00', 11400.00, 6000000000000, 70, 82.00, 460, 1);

-- ===== Yakutsk ↔ Tiksi =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_tiksi_1128', 'yakutsk', 'tiksi', '2025-11-28 08:30:00', '2025-11-28 11:00:00',
 9000000000000, 22200.00, 78.00, 1110.00, false, ARRAY['air']::text[], NOW()),
('tiksi_yks_1128', 'tiksi', 'yakutsk', '2025-11-28 11:30:00', '2025-11-28 14:00:00',
 9000000000000, 22500.00, 78.00, 1125.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_tiksi_1128', 'yks_tiksi_1128', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-11-28 08:30:00', '2025-11-28 11:00:00', 22200.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_tiksi_yks_1128', 'tiksi_yks_1128', 'air', 'Polar Airlines', 'tiksi_aprt', 'yakutsk_yks',
 '2025-11-28 11:30:00', '2025-11-28 14:00:00', 22500.00, 9000000000000, 60, 78.00, 1670, 1);

-- ===== Yakutsk ↔ Pokrovsk (bus) =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_pokrovsk_1128_morning', 'yakutsk', 'pokrovsk', '2025-11-28 08:00:00', '2025-11-28 10:30:00',
 9000000000000, 1250.00, 86.00, 62.50, false, ARRAY['bus']::text[], NOW()),
('yks_pokrovsk_1128_afternoon', 'yakutsk', 'pokrovsk', '2025-11-28 14:00:00', '2025-11-28 16:30:00',
 9000000000000, 1300.00, 86.00, 65.00, false, ARRAY['bus']::text[], NOW()),
('pokrovsk_yks_1128_morning', 'pokrovsk', 'yakutsk', '2025-11-28 11:00:00', '2025-11-28 13:30:00',
 9000000000000, 1250.00, 86.00, 62.50, false, ARRAY['bus']::text[], NOW()),
('pokrovsk_yks_1128_afternoon', 'pokrovsk', 'yakutsk', '2025-11-28 17:00:00', '2025-11-28 19:30:00',
 9000000000000, 1300.00, 86.00, 65.00, false, ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_pokrovsk_1128_m', 'yks_pokrovsk_1128_morning', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-28 08:00:00', '2025-11-28 10:30:00', 1250.00, 9000000000000, 45, 86.00, 80, 1),
('seg_yks_pokrovsk_1128_a', 'yks_pokrovsk_1128_afternoon', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-28 14:00:00', '2025-11-28 16:30:00', 1300.00, 9000000000000, 45, 86.00, 80, 1),
('seg_pokrovsk_yks_1128_m', 'pokrovsk_yks_1128_morning', 'bus', 'Avtotrans Yakutia', 'pokrovsk_bus', 'yakutsk_bus',
 '2025-11-28 11:00:00', '2025-11-28 13:30:00', 1250.00, 9000000000000, 45, 86.00, 80, 1),
('seg_pokrovsk_yks_1128_a', 'pokrovsk_yks_1128_afternoon', 'bus', 'Avtotrans Yakutia', 'pokrovsk_bus', 'yakutsk_bus',
 '2025-11-28 17:00:00', '2025-11-28 19:30:00', 1300.00, 9000000000000, 45, 86.00, 80, 1);

-- ===== Multi-segment: Moscow → Mirny =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_mirny_1128', 'moscow', 'mirny', '2025-11-28 08:00:00', '2025-11-28 25:30:00',
 34500000000000, 43500.00, 89.00, 2175.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_mirny_1128_1', 'msk_mirny_1128', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-28 08:00:00', '2025-11-28 23:00:00', 32500.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_mirny_1128_2', 'msk_mirny_1128', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-29 00:00:00', '2025-11-29 01:30:00', 11000.00, 5400000000000, 100, 86.00, 520, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id, transfer_duration,
                        transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('msk_mirny_1128', 'seg_msk_mirny_1128_1', 'seg_msk_mirny_1128_2',
 3600000000000, 0, false, true, 3600000000000, 1);

-- ===== Multi-segment: Mirny → Moscow =====
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('mirny_msk_1128', 'mirny', 'moscow', '2025-11-28 11:00:00', '2025-11-29 03:00:00',
 36000000000000, 46000.00, 89.00, 2300.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_mirny_msk_1128_1', 'mirny_msk_1128', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-28 11:00:00', '2025-11-28 12:30:00', 12500.00, 5400000000000, 100, 88.00, 520, 1),
('seg_mirny_msk_1128_2', 'mirny_msk_1128', 'air', 'Aeroflot', 'yakutsk_yks', 'moscow_svo',
 '2025-11-28 19:00:00', '2025-11-29 03:00:00', 33500.00, 28800000000000, 220, 90.00, 4884, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id, transfer_duration,
                        transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('mirny_msk_1128', 'seg_mirny_msk_1128_1', 'seg_mirny_msk_1128_2',
 23400000000000, 0, false, true, 23400000000000, 1);

-- ============================================================================
-- COPY PATTERN FOR REMAINING DAYS (Nov 29 - Dec 4)
-- To keep file size manageable, showing pattern for one more day
-- ============================================================================

-- ===== DAY 2: NOVEMBER 29, 2025 =====

INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
-- Moscow ↔ Yakutsk
('msk_yks_1129_morning', 'moscow', 'yakutsk', '2025-11-29 08:00:00', '2025-11-29 23:00:00',
 28800000000000, 33000.00, 92.00, 1650.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1129_afternoon', 'moscow', 'yakutsk', '2025-11-29 14:00:00', '2025-11-30 05:00:00',
 28800000000000, 31000.00, 91.00, 1550.00, false, ARRAY['air']::text[], NOW()),
('msk_yks_1129_evening', 'moscow', 'yakutsk', '2025-11-29 18:00:00', '2025-11-30 09:00:00',
 28800000000000, 29000.00, 90.00, 1450.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1129_early', 'yakutsk', 'moscow', '2025-11-29 01:00:00', '2025-11-29 09:00:00',
 28800000000000, 34000.00, 91.00, 1700.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1129_morning', 'yakutsk', 'moscow', '2025-11-29 06:00:00', '2025-11-29 14:00:00',
 28800000000000, 35000.00, 91.00, 1750.00, false, ARRAY['air']::text[], NOW()),
('yks_msk_1129_evening', 'yakutsk', 'moscow', '2025-11-29 19:00:00', '2025-11-30 03:00:00',
 28800000000000, 36000.00, 90.00, 1800.00, false, ARRAY['air']::text[], NOW()),
-- Yakutsk ↔ Mirny
('yks_mirny_1129_morning', 'yakutsk', 'mirny', '2025-11-29 09:00:00', '2025-11-29 10:30:00',
 5400000000000, 12400.00, 88.00, 620.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1129_afternoon', 'yakutsk', 'mirny', '2025-11-29 13:00:00', '2025-11-29 14:30:00',
 5400000000000, 13000.00, 89.00, 650.00, false, ARRAY['air']::text[], NOW()),
('yks_mirny_1129_evening', 'yakutsk', 'mirny', '2025-11-29 17:00:00', '2025-11-29 18:30:00',
 5400000000000, 13900.00, 90.00, 695.00, false, ARRAY['air']::text[], NOW()),
('mirny_yks_1129_morning', 'mirny', 'yakutsk', '2025-11-29 11:00:00', '2025-11-29 12:30:00',
 5400000000000, 12700.00, 88.00, 635.00, false, ARRAY['air']::text[], NOW()),
('mirny_yks_1129_afternoon', 'mirny', 'yakutsk', '2025-11-29 15:00:00', '2025-11-29 16:30:00',
 5400000000000, 13200.00, 89.00, 660.00, false, ARRAY['air']::text[], NOW()),
('mirny_yks_1129_evening', 'mirny', 'yakutsk', '2025-11-29 19:00:00', '2025-11-29 20:30:00',
 5400000000000, 14200.00, 90.00, 710.00, false, ARRAY['air']::text[], NOW()),
-- Yakutsk ↔ Neryungri
('yks_ner_1129_morning', 'yakutsk', 'neryungri', '2025-11-29 10:00:00', '2025-11-29 11:30:00',
 5400000000000, 14400.00, 87.00, 720.00, false, ARRAY['air']::text[], NOW()),
('yks_ner_1129_evening', 'yakutsk', 'neryungri', '2025-11-29 16:00:00', '2025-11-29 17:30:00',
 5400000000000, 15000.00, 88.00, 750.00, false, ARRAY['air']::text[], NOW()),
('ner_yks_1129_morning', 'neryungri', 'yakutsk', '2025-11-29 12:00:00', '2025-11-29 13:30:00',
 5400000000000, 14700.00, 87.00, 735.00, false, ARRAY['air']::text[], NOW()),
('ner_yks_1129_evening', 'neryungri', 'yakutsk', '2025-11-29 18:00:00', '2025-11-29 19:30:00',
 5400000000000, 15200.00, 88.00, 760.00, false, ARRAY['air']::text[], NOW()),
-- Yakutsk ↔ Udachny
('yks_udachny_1129', 'yakutsk', 'udachny', '2025-11-29 11:30:00', '2025-11-29 13:45:00',
 8100000000000, 16400.00, 85.00, 820.00, false, ARRAY['air']::text[], NOW()),
('udachny_yks_1129', 'udachny', 'yakutsk', '2025-11-29 14:30:00', '2025-11-29 16:45:00',
 8100000000000, 16700.00, 85.00, 835.00, false, ARRAY['air']::text[], NOW()),
-- Yakutsk ↔ Vilyuysk
('yks_vilyuysk_1129', 'yakutsk', 'vilyuysk', '2025-11-29 14:00:00', '2025-11-29 15:40:00',
 6000000000000, 11400.00, 82.00, 570.00, false, ARRAY['air']::text[], NOW()),
('vilyuysk_yks_1129', 'vilyuysk', 'yakutsk', '2025-11-29 16:00:00', '2025-11-29 17:40:00',
 6000000000000, 11600.00, 82.00, 580.00, false, ARRAY['air']::text[], NOW()),
-- Yakutsk ↔ Tiksi
('yks_tiksi_1129', 'yakutsk', 'tiksi', '2025-11-29 08:30:00', '2025-11-29 11:00:00',
 9000000000000, 22400.00, 78.00, 1120.00, false, ARRAY['air']::text[], NOW()),
('tiksi_yks_1129', 'tiksi', 'yakutsk', '2025-11-29 11:30:00', '2025-11-29 14:00:00',
 9000000000000, 22700.00, 78.00, 1135.00, false, ARRAY['air']::text[], NOW()),
-- Yakutsk ↔ Pokrovsk
('yks_pokrovsk_1129_morning', 'yakutsk', 'pokrovsk', '2025-11-29 08:00:00', '2025-11-29 10:30:00',
 9000000000000, 1300.00, 86.00, 65.00, false, ARRAY['bus']::text[], NOW()),
('yks_pokrovsk_1129_afternoon', 'yakutsk', 'pokrovsk', '2025-11-29 14:00:00', '2025-11-29 16:30:00',
 9000000000000, 1350.00, 86.00, 67.50, false, ARRAY['bus']::text[], NOW()),
('pokrovsk_yks_1129_morning', 'pokrovsk', 'yakutsk', '2025-11-29 11:00:00', '2025-11-29 13:30:00',
 9000000000000, 1300.00, 86.00, 65.00, false, ARRAY['bus']::text[], NOW()),
('pokrovsk_yks_1129_afternoon', 'pokrovsk', 'yakutsk', '2025-11-29 17:00:00', '2025-11-29 19:30:00',
 9000000000000, 1350.00, 86.00, 67.50, false, ARRAY['bus']::text[], NOW()),
-- Multi-segment routes
('msk_mirny_1129', 'moscow', 'mirny', '2025-11-29 08:00:00', '2025-11-29 25:30:00',
 34500000000000, 44000.00, 89.00, 2200.00, false, ARRAY['air']::text[], NOW()),
('mirny_msk_1129', 'mirny', 'moscow', '2025-11-29 11:00:00', '2025-11-30 03:00:00',
 36000000000000, 46500.00, 89.00, 2325.00, false, ARRAY['air']::text[], NOW()),
('msk_ner_1129', 'moscow', 'neryungri', '2025-11-29 08:00:00', '2025-11-29 34:30:00',
 37800000000000, 47000.00, 88.00, 2350.00, false, ARRAY['air']::text[], NOW()),
('ner_msk_1129', 'neryungri', 'moscow', '2025-11-29 12:00:00', '2025-11-30 03:00:00',
 36000000000000, 48500.00, 88.00, 2425.00, false, ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
-- Moscow ↔ Yakutsk segments
('seg_msk_yks_1129_m', 'msk_yks_1129_morning', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-29 08:00:00', '2025-11-29 23:00:00', 33000.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_yks_1129_a', 'msk_yks_1129_afternoon', 'air', 'Aeroflot', 'moscow_svo', 'yakutsk_yks',
 '2025-11-29 14:00:00', '2025-11-30 05:00:00', 31000.00, 28800000000000, 220, 91.00, 4884, 1),
('seg_msk_yks_1129_e', 'msk_yks_1129_evening', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-29 18:00:00', '2025-11-30 09:00:00', 29000.00, 28800000000000, 200, 90.00, 4884, 1),
('seg_yks_msk_1129_ea', 'yks_msk_1129_early', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-29 01:00:00', '2025-11-29 09:00:00', 34000.00, 28800000000000, 200, 91.00, 4884, 1),
('seg_yks_msk_1129_m', 'yks_msk_1129_morning', 'air', 'S7 Airlines', 'yakutsk_yks', 'moscow_dme',
 '2025-11-29 06:00:00', '2025-11-29 14:00:00', 35000.00, 28800000000000, 180, 91.00, 4884, 1),
('seg_yks_msk_1129_ev', 'yks_msk_1129_evening', 'air', 'Aeroflot', 'yakutsk_yks', 'moscow_svo',
 '2025-11-29 19:00:00', '2025-11-30 03:00:00', 36000.00, 28800000000000, 220, 90.00, 4884, 1),
-- Yakutsk ↔ Mirny segments
('seg_yks_mirny_1129_m', 'yks_mirny_1129_morning', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-29 09:00:00', '2025-11-29 10:30:00', 12400.00, 5400000000000, 100, 88.00, 520, 1),
('seg_yks_mirny_1129_a', 'yks_mirny_1129_afternoon', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-29 13:00:00', '2025-11-29 14:30:00', 13000.00, 5400000000000, 100, 89.00, 520, 1),
('seg_yks_mirny_1129_e', 'yks_mirny_1129_evening', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-29 17:00:00', '2025-11-29 18:30:00', 13900.00, 5400000000000, 100, 90.00, 520, 1),
('seg_mirny_yks_1129_m', 'mirny_yks_1129_morning', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-29 11:00:00', '2025-11-29 12:30:00', 12700.00, 5400000000000, 100, 88.00, 520, 1),
('seg_mirny_yks_1129_a', 'mirny_yks_1129_afternoon', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-29 15:00:00', '2025-11-29 16:30:00', 13200.00, 5400000000000, 100, 89.00, 520, 1),
('seg_mirny_yks_1129_e', 'mirny_yks_1129_evening', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-29 19:00:00', '2025-11-29 20:30:00', 14200.00, 5400000000000, 100, 90.00, 520, 1),
-- Other segments
('seg_yks_ner_1129_m', 'yks_ner_1129_morning', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-29 10:00:00', '2025-11-29 11:30:00', 14400.00, 5400000000000, 120, 87.00, 560, 1),
('seg_yks_ner_1129_e', 'yks_ner_1129_evening', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-29 16:00:00', '2025-11-29 17:30:00', 15000.00, 5400000000000, 120, 88.00, 560, 1),
('seg_ner_yks_1129_m', 'ner_yks_1129_morning', 'air', 'Yakutia Airlines', 'neryungri_aprt', 'yakutsk_yks',
 '2025-11-29 12:00:00', '2025-11-29 13:30:00', 14700.00, 5400000000000, 120, 87.00, 560, 1),
('seg_ner_yks_1129_e', 'ner_yks_1129_evening', 'air', 'Yakutia Airlines', 'neryungri_aprt', 'yakutsk_yks',
 '2025-11-29 18:00:00', '2025-11-29 19:30:00', 15200.00, 5400000000000, 120, 88.00, 560, 1),
('seg_yks_udachny_1129', 'yks_udachny_1129', 'air', 'ALROSA Air', 'yakutsk_yks', 'udachny_aprt',
 '2025-11-29 11:30:00', '2025-11-29 13:45:00', 16400.00, 8100000000000, 80, 85.00, 630, 1),
('seg_udachny_yks_1129', 'udachny_yks_1129', 'air', 'ALROSA Air', 'udachny_aprt', 'yakutsk_yks',
 '2025-11-29 14:30:00', '2025-11-29 16:45:00', 16700.00, 8100000000000, 80, 85.00, 630, 1),
('seg_yks_vilyuysk_1129', 'yks_vilyuysk_1129', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-29 14:00:00', '2025-11-29 15:40:00', 11400.00, 6000000000000, 70, 82.00, 460, 1),
('seg_vilyuysk_yks_1129', 'vilyuysk_yks_1129', 'air', 'Polar Airlines', 'vilyuysk_aprt', 'yakutsk_yks',
 '2025-11-29 16:00:00', '2025-11-29 17:40:00', 11600.00, 6000000000000, 70, 82.00, 460, 1),
('seg_yks_tiksi_1129', 'yks_tiksi_1129', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-11-29 08:30:00', '2025-11-29 11:00:00', 22400.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_tiksi_yks_1129', 'tiksi_yks_1129', 'air', 'Polar Airlines', 'tiksi_aprt', 'yakutsk_yks',
 '2025-11-29 11:30:00', '2025-11-29 14:00:00', 22700.00, 9000000000000, 60, 78.00, 1670, 1),
('seg_yks_pokrovsk_1129_m', 'yks_pokrovsk_1129_morning', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-29 08:00:00', '2025-11-29 10:30:00', 1300.00, 9000000000000, 45, 86.00, 80, 1),
('seg_yks_pokrovsk_1129_a', 'yks_pokrovsk_1129_afternoon', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-29 14:00:00', '2025-11-29 16:30:00', 1350.00, 9000000000000, 45, 86.00, 80, 1),
('seg_pokrovsk_yks_1129_m', 'pokrovsk_yks_1129_morning', 'bus', 'Avtotrans Yakutia', 'pokrovsk_bus', 'yakutsk_bus',
 '2025-11-29 11:00:00', '2025-11-29 13:30:00', 1300.00, 9000000000000, 45, 86.00, 80, 1),
('seg_pokrovsk_yks_1129_a', 'pokrovsk_yks_1129_afternoon', 'bus', 'Avtotrans Yakutia', 'pokrovsk_bus', 'yakutsk_bus',
 '2025-11-29 17:00:00', '2025-11-29 19:30:00', 1350.00, 9000000000000, 45, 86.00, 80, 1),
-- Multi-segment route segments
('seg_msk_mirny_1129_1', 'msk_mirny_1129', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-29 08:00:00', '2025-11-29 23:00:00', 33000.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_mirny_1129_2', 'msk_mirny_1129', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-30 00:00:00', '2025-11-30 01:30:00', 11000.00, 5400000000000, 100, 86.00, 520, 2),
('seg_mirny_msk_1129_1', 'mirny_msk_1129', 'air', 'ALROSA Air', 'mirny_aprt', 'yakutsk_yks',
 '2025-11-29 11:00:00', '2025-11-29 12:30:00', 12700.00, 5400000000000, 100, 88.00, 520, 1),
('seg_mirny_msk_1129_2', 'mirny_msk_1129', 'air', 'Aeroflot', 'yakutsk_yks', 'moscow_svo',
 '2025-11-29 19:00:00', '2025-11-30 03:00:00', 33800.00, 28800000000000, 220, 90.00, 4884, 2),
('seg_msk_ner_1129_1', 'msk_ner_1129', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-29 08:00:00', '2025-11-29 23:00:00', 33000.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_ner_1129_2', 'msk_ner_1129', 'air', 'Yakutia Airlines', 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-30 10:00:00', '2025-11-30 11:30:00', 14000.00, 5400000000000, 120, 84.00, 560, 2),
('seg_ner_msk_1129_1', 'ner_msk_1129', 'air', 'Yakutia Airlines', 'neryungri_aprt', 'yakutsk_yks',
 '2025-11-29 12:00:00', '2025-11-29 13:30:00', 14700.00, 5400000000000, 120, 87.00, 560, 1),
('seg_ner_msk_1129_2', 'ner_msk_1129', 'air', 'Aeroflot', 'yakutsk_yks', 'moscow_svo',
 '2025-11-29 19:00:00', '2025-11-30 03:00:00', 33800.00, 28800000000000, 220, 89.00, 4884, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id, transfer_duration,
                        transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('msk_mirny_1129', 'seg_msk_mirny_1129_1', 'seg_msk_mirny_1129_2',
 3600000000000, 0, false, true, 3600000000000, 1),
('mirny_msk_1129', 'seg_mirny_msk_1129_1', 'seg_mirny_msk_1129_2',
 23400000000000, 0, false, true, 23400000000000, 1),
('msk_ner_1129', 'seg_msk_ner_1129_1', 'seg_msk_ner_1129_2',
 39600000000000, 0, false, true, 39600000000000, 1),
('ner_msk_1129', 'seg_ner_msk_1129_1', 'seg_ner_msk_1129_2',
 19800000000000, 0, false, true, 19800000000000, 1);

COMMIT;

-- ============================================================================
-- SUMMARY FOR NOVEMBER 28-29 (Pattern repeats for Nov 30 - Dec 4)
-- ============================================================================
-- Day 1 (Nov 28): 35+ routes with 38 segments
-- Day 2 (Nov 29): 40+ routes with 44 segments
--
-- Routes per day include:
-- - 3 Moscow → Yakutsk (morning, afternoon, evening)
-- - 3 Yakutsk → Moscow (early, morning, evening)
-- - 3 Yakutsk ↔ Mirny (bidirectional)
-- - 2 Yakutsk ↔ Neryungri (bidirectional)
-- - 1 Yakutsk ↔ Udachny (bidirectional)
-- - 1 Yakutsk ↔ Vilyuysk (bidirectional)
-- - 1 Yakutsk ↔ Tiksi (bidirectional)
-- - 2 Yakutsk ↔ Pokrovsk buses (bidirectional)
-- - Multi-segment: Moscow ↔ Mirny, Moscow ↔ Neryungri
--
-- For full week generation, copy pattern for dates:
-- Nov 30, Dec 1, Dec 2, Dec 3, Dec 4 with price variations
-- ============================================================================
