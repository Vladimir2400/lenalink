-- ============================================================================
-- SEED DATA FOR NOVEMBER 27, 2025
-- LenaLink - Multi-modal Transport Aggregator for Yakutia
-- ============================================================================

BEGIN;

-- ============================================================================
-- 1. STOPS (Transport hubs)
-- ============================================================================

INSERT INTO stops (id, name, city, city_display_name, latitude, longitude, stop_type) VALUES
-- Moscow airports
('moscow_dme', 'Domodedovo International Airport', 'moscow', 'Москва', 55.4088, 37.9063, 'airport'),
('moscow_svo', 'Sheremetyevo International Airport', 'moscow', 'Москва', 55.9726, 37.4147, 'airport'),

-- Yakutsk (main hub)
('yakutsk_yks', 'Yakutsk Airport (Platon Oyunsky)', 'yakutsk', 'Якутск', 62.0932, 129.7708, 'airport'),
('yakutsk_bus', 'Yakutsk Bus Terminal', 'yakutsk', 'Якутск', 62.0339, 129.7331, 'station'),
('yakutsk_port', 'Yakutsk River Port', 'yakutsk', 'Якутск', 62.0272, 129.7322, 'port'),

-- Diamond cities
('mirny_aprt', 'Mirny Airport', 'mirny', 'Мирный', 62.5347, 114.0389, 'airport'),
('udachny_aprt', 'Udachny Airport', 'udachny', 'Удачный', 66.4000, 112.0333, 'airport'),

-- Southern Yakutia
('neryungri_aprt', 'Neryungri Airport', 'neryungri', 'Нерюнгри', 56.9139, 124.9144, 'airport'),
('neryungri_bus', 'Neryungri Bus Terminal', 'neryungri', 'Нерюнгри', 56.6600, 124.7200, 'station'),

-- Olekminsk
('olekminsk_port', 'Olekminsk River Port', 'olekminsk', 'Олекминск', 60.3733, 120.4272, 'port'),

-- Vilyuysk group
('vilyuysk_aprt', 'Vilyuysk Airport', 'vilyuysk', 'Вилюйск', 63.7486, 121.5675, 'airport'),

-- Northern cities
('tiksi_aprt', 'Tiksi Airport', 'tiksi', 'Тикси', 71.6977, 128.9031, 'airport'),

-- Pokrovsk
('pokrovsk_bus', 'Pokrovsk Bus Terminal', 'pokrovsk', 'Покровск', 61.4833, 129.1500, 'station');

-- ============================================================================
-- 2. ROUTES - Moscow ↔ Yakutsk (November 27, 2025)
-- ============================================================================

-- Route 1: Moscow → Yakutsk (morning flight, DME)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1127_morning', 'moscow', 'yakutsk',
 '2025-11-27 08:00:00', '2025-11-27 23:00:00',
 28800000000000, 32000.00, 92.00, 1600.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1127_m', 'msk_yks_1127_morning', 'air', 'S7 Airlines',
 'moscow_dme', 'yakutsk_yks',
 '2025-11-27 08:00:00', '2025-11-27 23:00:00',
 32000.00, 28800000000000, 180, 92.00, 4884, 1);

-- Route 2: Moscow → Yakutsk (evening flight, SVO)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_yks_1127_evening', 'moscow', 'yakutsk',
 '2025-11-27 18:00:00', '2025-11-28 09:00:00',
 28800000000000, 28000.00, 90.00, 1400.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1127_e', 'msk_yks_1127_evening', 'air', 'Yakutia Airlines',
 'moscow_svo', 'yakutsk_yks',
 '2025-11-27 18:00:00', '2025-11-28 09:00:00',
 28000.00, 28800000000000, 200, 90.00, 4884, 1);

-- Route 3: Yakutsk → Moscow (morning return)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_msk_1127_morning', 'yakutsk', 'moscow',
 '2025-11-27 06:00:00', '2025-11-27 14:00:00',
 28800000000000, 34000.00, 91.00, 1700.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_msk_1127_m', 'yks_msk_1127_morning', 'air', 'S7 Airlines',
 'yakutsk_yks', 'moscow_dme',
 '2025-11-27 06:00:00', '2025-11-27 14:00:00',
 34000.00, 28800000000000, 180, 91.00, 4884, 1);

-- ============================================================================
-- 3. ROUTES - Within Yakutia (November 27, 2025)
-- ============================================================================

-- Route 4: Yakutsk → Mirny (morning)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_mirny_1127_morning', 'yakutsk', 'mirny',
 '2025-11-27 09:00:00', '2025-11-27 10:30:00',
 5400000000000, 12000.00, 88.00, 600.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_1127', 'yks_mirny_1127_morning', 'air', 'ALROSA Air',
 'yakutsk_yks', 'mirny_aprt',
 '2025-11-27 09:00:00', '2025-11-27 10:30:00',
 12000.00, 5400000000000, 100, 88.00, 520, 1);

-- Route 5: Yakutsk → Mirny (evening)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_mirny_1127_evening', 'yakutsk', 'mirny',
 '2025-11-27 17:00:00', '2025-11-27 18:30:00',
 5400000000000, 13500.00, 90.00, 675.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_1127_e', 'yks_mirny_1127_evening', 'air', 'ALROSA Air',
 'yakutsk_yks', 'mirny_aprt',
 '2025-11-27 17:00:00', '2025-11-27 18:30:00',
 13500.00, 5400000000000, 100, 90.00, 520, 1);

-- Route 6: Yakutsk → Neryungri
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_ner_1127', 'yakutsk', 'neryungri',
 '2025-11-27 10:00:00', '2025-11-27 11:30:00',
 5400000000000, 14000.00, 87.00, 700.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_ner_1127', 'yks_ner_1127', 'air', 'Yakutia Airlines',
 'yakutsk_yks', 'neryungri_aprt',
 '2025-11-27 10:00:00', '2025-11-27 11:30:00',
 14000.00, 5400000000000, 120, 87.00, 560, 1);

-- Route 7: Yakutsk → Udachny
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_udachny_1127', 'yakutsk', 'udachny',
 '2025-11-27 11:30:00', '2025-11-27 13:45:00',
 8100000000000, 16000.00, 85.00, 800.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_udachny_1127', 'yks_udachny_1127', 'air', 'ALROSA Air',
 'yakutsk_yks', 'udachny_aprt',
 '2025-11-27 11:30:00', '2025-11-27 13:45:00',
 16000.00, 8100000000000, 80, 85.00, 630, 1);

-- Route 8: Yakutsk → Vilyuysk
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_vilyuysk_1127', 'yakutsk', 'vilyuysk',
 '2025-11-27 14:00:00', '2025-11-27 15:40:00',
 6000000000000, 11000.00, 82.00, 550.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_vilyuysk_1127', 'yks_vilyuysk_1127', 'air', 'Polar Airlines',
 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-27 14:00:00', '2025-11-27 15:40:00',
 11000.00, 6000000000000, 70, 82.00, 460, 1);

-- Route 9: Yakutsk → Tiksi (far north)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_tiksi_1127', 'yakutsk', 'tiksi',
 '2025-11-27 08:30:00', '2025-11-27 11:00:00',
 9000000000000, 22000.00, 78.00, 1100.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_tiksi_1127', 'yks_tiksi_1127', 'air', 'Polar Airlines',
 'yakutsk_yks', 'tiksi_aprt',
 '2025-11-27 08:30:00', '2025-11-27 11:00:00',
 22000.00, 9000000000000, 60, 78.00, 1670, 1);

-- Route 10: Yakutsk → Pokrovsk (bus, suburban)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_pokrovsk_1127', 'yakutsk', 'pokrovsk',
 '2025-11-27 08:00:00', '2025-11-27 10:30:00',
 9000000000000, 1200.00, 86.00, 60.00, false,
 ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_pokrovsk_1127', 'yks_pokrovsk_1127', 'bus', 'Avtotrans Yakutia',
 'yakutsk_bus', 'pokrovsk_bus',
 '2025-11-27 08:00:00', '2025-11-27 10:30:00',
 1200.00, 9000000000000, 45, 86.00, 80, 1);

-- Route 11: Yakutsk → Neryungri (bus, cheaper option)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('yks_ner_bus_1127', 'yakutsk', 'neryungri',
 '2025-11-27 09:00:00', '2025-11-27 19:00:00',
 36000000000000, 3500.00, 75.00, 175.00, false,
 ARRAY['bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_ner_bus_1127', 'yks_ner_bus_1127', 'bus', 'Amur Transport',
 'yakutsk_bus', 'neryungri_bus',
 '2025-11-27 09:00:00', '2025-11-27 19:00:00',
 3500.00, 36000000000000, 50, 75.00, 350, 1);

-- ============================================================================
-- 4. COMBINED ROUTES - Moscow → Yakutia cities (November 27-28, 2025)
-- ============================================================================

-- Route 12: Moscow → Olekminsk (via Yakutsk + bus)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_olek_1127', 'moscow', 'olekminsk',
 '2025-11-27 08:00:00', '2025-11-28 08:00:00',
 64800000000000, 35000.00, 85.00, 1750.00, false,
 ARRAY['air','bus']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_olek_1', 'msk_olek_1127', 'air', 'S7 Airlines',
 'moscow_dme', 'yakutsk_yks',
 '2025-11-27 08:00:00', '2025-11-27 23:00:00',
 32000.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_olek_2', 'msk_olek_1127', 'bus', 'Avtotrans Yakutia',
 'yakutsk_bus', 'olekminsk_port',
 '2025-11-28 01:00:00', '2025-11-28 08:00:00',
 3000.00, 25200000000000, 40, 78.00, 610, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id, transfer_duration,
                        transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('msk_olek_1127', 'seg_msk_olek_1', 'seg_msk_olek_2',
 7200000000000, 12, true, true, 7200000000000, 1);

-- Route 13: Moscow → Mirny (via Yakutsk)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('msk_mirny_1127', 'moscow', 'mirny',
 '2025-11-27 08:00:00', '2025-11-28 01:30:00',
 42300000000000, 42000.00, 88.00, 2100.00, false,
 ARRAY['air']::text[], NOW());

INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_mirny_1', 'msk_mirny_1127', 'air', 'S7 Airlines',
 'moscow_dme', 'yakutsk_yks',
 '2025-11-27 08:00:00', '2025-11-27 23:00:00',
 32000.00, 28800000000000, 180, 92.00, 4884, 1),
('seg_msk_mirny_2', 'msk_mirny_1127', 'air', 'ALROSA Air',
 'yakutsk_yks', 'mirny_aprt',
 '2025-11-28 00:00:00', '2025-11-28 01:30:00',
 10000.00, 5400000000000, 100, 84.00, 520, 2);

INSERT INTO connections (route_id, from_segment_id, to_segment_id, transfer_duration,
                        transfer_distance, requires_transport, is_valid, gap, sequence_order) VALUES
('msk_mirny_1127', 'seg_msk_mirny_1', 'seg_msk_mirny_2',
 3600000000000, 0, false, true, 3600000000000, 1);

COMMIT;

-- ============================================================================
-- SUMMARY:
-- - 14 stops added (Moscow, Yakutsk, regional cities)
-- - 13 routes added (3 Moscow↔Yakutsk, 8 within Yakutia, 2 combined)
-- - 15 segments added
-- - 2 connections added (for multi-segment routes)
-- - All data for November 27, 2025
-- ============================================================================
