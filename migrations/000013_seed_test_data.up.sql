-- Seed test data for LenaLink - Yakutia transport routes
-- This migration populates realistic routes, stops, segments, and connections

-- ============================================================================
-- STOPS (Airports, Ports, Stations, Bus Terminals)
-- ============================================================================

INSERT INTO stops (id, name, city, latitude, longitude, stop_type) VALUES
-- Moscow airports
('moscow_dme', 'Domodedovo International Airport', 'Moscow', 55.4088, 37.9063, 'airport'),
('moscow_svo', 'Sheremetyevo International Airport', 'Moscow', 55.9725, 37.4146, 'airport'),

-- Yakutia
('yakutsk_yks', 'Yakutsk Tolmachevo Airport', 'Yakutsk', 62.0932, 129.7708, 'airport'),
('yakutsk_port', 'Yakutsk River Port', 'Yakutsk', 62.0272, 129.7322, 'port'),
('yakutsk_bus', 'Yakutsk Bus Terminal', 'Yakutsk', 62.0332, 129.7442, 'station'),
('yakutsk_railway', 'Yakutsk Railway Station', 'Yakutsk', 62.0083, 129.7192, 'station'),

-- Diamond cities
('mirny_aprt', 'Mirny Airport', 'Mirny', 62.5408, 114.0397, 'airport'),
('mirny_bus', 'Mirny Bus Terminal', 'Mirny', 62.5358, 114.0447, 'station'),
('udachny_aprt', 'Udachny Airport', 'Udachny', 66.4142, 112.3869, 'airport'),

-- Coal and mining
('nerungri_aprt', 'Neryungri Airport', 'Neryungri', 56.0068, 124.6756, 'airport'),
('nerungri_bus', 'Neryungri Bus Terminal', 'Neryungri', 56.0068, 124.6806, 'station'),
('aldan_aprt', 'Aldan Airport', 'Aldan', 58.6042, 125.3975, 'airport'),
('aldan_bus', 'Aldan Bus Terminal', 'Aldan', 58.6042, 125.4025, 'station'),

-- River ports on Lena river
('lensky_port', 'Lensky River Port', 'Lensky', 60.7458, 114.9417, 'port'),
('olekminsk_port', 'Olekminsk River Port', 'Olekminsk', 60.3733, 120.4272, 'port'),
('sangur_port', 'Sangur River Port', 'Sangur', 61.5806, 121.6511, 'port'),
('pokrovsk_port', 'Pokrovsk River Port', 'Pokrovsk', 61.5692, 118.5289, 'port'),
('batagay_port', 'Batagay River Port', 'Batagay', 63.2836, 129.5417, 'port'),

-- Far north
('tiksi_aprt', 'Tiksi Airport', 'Tiksi', 71.5900, 128.8889, 'airport'),
('tiksi_port', 'Tiksi Seaport', 'Tiksi', 71.6333, 128.8667, 'port'),
('verkhoyansk_aprt', 'Verkhoyansk Airport', 'Verkhoyansk', 67.5558, 133.3917, 'airport'),
('oymyakon_aprt', 'Oymyakon Airport', 'Oymyakon', 63.4667, 142.7833, 'airport'),

-- Other cities
('vilyuysk_aprt', 'Vilyuysk Airport', 'Vilyuysk', 63.7486, 121.5675, 'airport'),
('nyurba_aprt', 'Nyurba Airport', 'Nyurba', 65.3875, 118.4778, 'airport'),
('zhataay_bus', 'Zhataay Bus Terminal (Yakutsk suburb)', 'Zhataay', 62.1500, 129.5000, 'station'),

-- BAM railway
('tynda_railway', 'Tynda Railway Station (BAM)', 'Tynda', 62.0075, 124.2603, 'station'),
('tommot_railway', 'Tommot Railway Station (BAM)', 'Tommot', 56.3636, 121.6303, 'station');

-- ============================================================================
-- ROUTES (Main journeys)
-- ============================================================================

-- Route 1: Moscow to Yakutsk - Option A (S7 Airlines - optimal)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_yks_opt', 'Moscow', 'Yakutsk',
 '2025-11-20 08:00:00', '2025-11-20 16:30:00', 31800000000000,
 32500.00, 92.00, 1625.00, true, ARRAY['air']::text[], NOW());

-- Route 2: Moscow to Yakutsk - Option B (Yakutia Airlines - cheap)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_yks_cheap', 'Moscow', 'Yakutsk',
 '2025-11-20 10:30:00', '2025-11-20 19:15:00', 31500000000000,
 28000.00, 88.00, 1400.00, true, ARRAY['air']::text[], NOW());

-- Route 3: Moscow to Yakutsk - Option C (with transfer)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_yks_transfer', 'Moscow', 'Yakutsk',
 '2025-11-20 06:00:00', '2025-11-21 02:00:00', 72000000000000,
 25000.00, 80.00, 1250.00, true, ARRAY['air']::text[], NOW());

-- Route 4: Moscow to Olekminsk (multimodal - optimal)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_olek_opt', 'Moscow', 'Olekminsk',
 '2025-11-20 08:00:00', '2025-11-21 05:40:00', 95400000000000,
 41500.00, 85.00, 2075.00, true, ARRAY['air','taxi','river']::text[], NOW());

-- Route 5: Moscow to Olekminsk (cheap variant)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_olek_cheap', 'Moscow', 'Olekminsk',
 '2025-11-20 10:30:00', '2025-11-21 12:00:00', 99000000000000,
 32000.00, 78.00, 1600.00, true, ARRAY['air','taxi','bus']::text[], NOW());

-- Route 6: Moscow to Sangur (complex multimodal)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_msk_sangur', 'Moscow', 'Sangur',
 '2025-11-20 08:00:00', '2025-11-21 07:00:00', 97200000000000,
 43000.00, 80.00, 2150.00, true, ARRAY['air','taxi','bus','river']::text[], NOW());

-- Route 7: Yakutsk to Mirny (air)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_mirny_air', 'Yakutsk', 'Mirny',
 '2025-11-20 09:00:00', '2025-11-20 10:30:00', 5400000000000,
 12000.00, 92.00, 600.00, true, ARRAY['air']::text[], NOW());

-- Route 8: Yakutsk to Mirny (road)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_mirny_bus', 'Yakutsk', 'Mirny',
 '2025-11-20 07:00:00', '2025-11-20 19:00:00', 43200000000000,
 4500.00, 70.00, 225.00, true, ARRAY['bus']::text[], NOW());

-- Route 9: Yakutsk to Lensky (river)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_lensky_river', 'Yakutsk', 'Lensky',
 '2025-11-20 08:00:00', '2025-11-21 04:00:00', 72000000000000,
 8000.00, 78.00, 400.00, true, ARRAY['river']::text[], NOW());

-- Route 10: Yakutsk to Lensky (bus)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_lensky_bus', 'Yakutsk', 'Lensky',
 '2025-11-20 06:00:00', '2025-11-20 18:00:00', 43200000000000,
 2500.00, 65.00, 125.00, true, ARRAY['bus']::text[], NOW());

-- Route 11: Yakutsk to Neryungri
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_nerungri', 'Yakutsk', 'Neryungri',
 '2025-11-20 10:00:00', '2025-11-20 20:00:00', 36000000000000,
 3500.00, 75.00, 175.00, true, ARRAY['bus']::text[], NOW());

-- Route 12: Yakutsk to Aldan
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_aldan', 'Yakutsk', 'Aldan',
 '2025-11-20 08:00:00', '2025-11-20 18:00:00', 36000000000000,
 3200.00, 80.00, 160.00, true, ARRAY['bus']::text[], NOW());

-- Route 13: Yakutsk to Tiksi (summer air)
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_tiksi_air', 'Yakutsk', 'Tiksi',
 '2025-06-15 08:00:00', '2025-06-15 11:00:00', 10800000000000,
 18000.00, 70.00, 900.00, true, ARRAY['air']::text[], NOW());

-- Route 14: Neryungri to Aldan
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_nerungri_aldan', 'Neryungri', 'Aldan',
 '2025-11-20 09:00:00', '2025-11-20 13:00:00', 14400000000000,
 1200.00, 88.00, 60.00, true, ARRAY['bus']::text[], NOW());

-- Route 15: Yakutsk to Vilyuysk
INSERT INTO routes (id, from_city, to_city, departure_time, arrival_time, total_duration,
                   total_price, reliability_score, insurance_premium, insurance_included,
                   transport_types, saved_at) VALUES
('route_yks_vilyuysk', 'Yakutsk', 'Vilyuysk',
 '2025-11-20 09:00:00', '2025-11-20 10:40:00', 6000000000000,
 10000.00, 75.00, 500.00, true, ARRAY['air']::text[], NOW());

-- ============================================================================
-- SEGMENTS (Individual transport legs)
-- ============================================================================

-- Route 1 (Moscow to Yakutsk - S7)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_1', 'route_msk_yks_opt', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-20 08:00:00', '2025-11-20 16:30:00', 32500.00, 31800000000000, 150, 92.0, 4100, 1);

-- Route 2 (Moscow to Yakutsk - Yakutia)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_2', 'route_msk_yks_cheap', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-20 10:30:00', '2025-11-20 19:15:00', 28000.00, 31500000000000, 180, 88.0, 4100, 1);

-- Route 3 (Moscow to Yakutsk with transfer)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_yks_3a', 'route_msk_yks_transfer', 'air', 'Ural Airlines', 'moscow_dme', 'moscow_svo',
 '2025-11-20 06:00:00', '2025-11-20 13:30:00', 12000.00, 27000000000000, 200, 80.0, 1000, 1),
('seg_msk_yks_3b', 'route_msk_yks_transfer', 'air', 'Ural Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-20 14:30:00', '2025-11-21 02:00:00', 13000.00, 42300000000000, 200, 80.0, 3900, 2);

-- Route 4 (Moscow to Olekminsk - optimal)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_olek_1', 'route_msk_olek_opt', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-20 08:00:00', '2025-11-20 16:30:00', 32500.00, 31800000000000, 150, 92.0, 4100, 1),
('seg_msk_olek_2', 'route_msk_olek_opt', 'taxi', 'Yandex Taxi', 'yakutsk_yks', 'yakutsk_port',
 '2025-11-20 17:00:00', '2025-11-20 17:40:00', 500.00, 2400000000000, 4, 95.0, 10, 2),
('seg_msk_olek_3', 'route_msk_olek_opt', 'river', 'Lenskiye Zori', 'yakutsk_port', 'olekminsk_port',
 '2025-11-20 19:00:00', '2025-11-21 05:40:00', 8500.00, 58800000000000, 100, 75.0, 580, 3);

-- Route 5 (Moscow to Olekminsk - cheap)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_msk_olek_cheap_1', 'route_msk_olek_cheap', 'air', 'Yakutia Airlines', 'moscow_svo', 'yakutsk_yks',
 '2025-11-20 10:30:00', '2025-11-20 19:15:00', 28000.00, 31500000000000, 180, 88.0, 4100, 1),
('seg_msk_olek_cheap_2', 'route_msk_olek_cheap', 'taxi', 'Maxim', 'yakutsk_yks', 'yakutsk_bus',
 '2025-11-20 19:45:00', '2025-11-20 20:00:00', 400.00, 900000000000, 4, 95.0, 5, 2),
('seg_msk_olek_cheap_3', 'route_msk_olek_cheap', 'bus', 'Avtotrans Yakutia', 'yakutsk_bus', 'olekminsk_port',
 '2025-11-20 22:00:00', '2025-11-21 12:00:00', 3200.00, 36000000000000, 50, 72.0, 400, 3),
('seg_msk_olek_cheap_4', 'route_msk_olek_cheap', 'walk', 'On foot', 'olekminsk_port', 'olekminsk_port',
 '2025-11-21 12:00:00', '2025-11-21 12:00:00', 0.00, 0, 1, 100.0, 0, 4);

-- Route 6 (Moscow to Sangur)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_sangur_1', 'route_msk_sangur', 'air', 'S7 Airlines', 'moscow_dme', 'yakutsk_yks',
 '2025-11-20 08:00:00', '2025-11-20 16:30:00', 32500.00, 31800000000000, 150, 92.0, 4100, 1),
('seg_sangur_2', 'route_msk_sangur', 'taxi', 'Yandex Taxi', 'yakutsk_yks', 'yakutsk_bus',
 '2025-11-20 17:00:00', '2025-11-20 17:20:00', 400.00, 1200000000000, 4, 95.0, 8, 2),
('seg_sangur_3', 'route_msk_sangur', 'bus', 'Siberia Lines', 'yakutsk_bus', 'pokrovsk_port',
 '2025-11-20 18:00:00', '2025-11-21 00:00:00', 1800.00, 21600000000000, 45, 75.0, 150, 3),
('seg_sangur_4', 'route_msk_sangur', 'river', 'Sakha River Transport', 'pokrovsk_port', 'sangur_port',
 '2025-11-21 01:00:00', '2025-11-21 07:00:00', 5200.00, 21600000000000, 80, 72.0, 240, 4);

-- Route 7 (Yakutsk to Mirny - air)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_air', 'route_yks_mirny_air', 'air', 'ALROSA Air', 'yakutsk_yks', 'mirny_aprt',
 '2025-11-20 09:00:00', '2025-11-20 10:30:00', 12000.00, 5400000000000, 100, 92.0, 520, 1);

-- Route 8 (Yakutsk to Mirny - bus)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_mirny_bus', 'route_yks_mirny_bus', 'bus', 'ALROSA Transport', 'yakutsk_bus', 'mirny_bus',
 '2025-11-20 07:00:00', '2025-11-20 19:00:00', 4500.00, 43200000000000, 60, 70.0, 450, 1);

-- Route 9 (Yakutsk to Lensky - river)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_lensky_river', 'route_yks_lensky_river', 'river', 'Lenskiye Zori', 'yakutsk_port', 'lensky_port',
 '2025-11-20 08:00:00', '2025-11-21 04:00:00', 8000.00, 72000000000000, 120, 78.0, 640, 1);

-- Route 10 (Yakutsk to Lensky - bus)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_lensky_bus', 'route_yks_lensky_bus', 'bus', 'Siberia Lines', 'yakutsk_bus', 'lensky_port',
 '2025-11-20 06:00:00', '2025-11-20 18:00:00', 2500.00, 43200000000000, 55, 65.0, 550, 1);

-- Route 11 (Yakutsk to Neryungri)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_nerungri', 'route_yks_nerungri', 'bus', 'Amur Transport', 'yakutsk_bus', 'nerungri_bus',
 '2025-11-20 10:00:00', '2025-11-20 20:00:00', 3500.00, 36000000000000, 50, 75.0, 350, 1);

-- Route 12 (Yakutsk to Aldan)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_aldan', 'route_yks_aldan', 'bus', 'Magistral', 'yakutsk_bus', 'aldan_bus',
 '2025-11-20 08:00:00', '2025-11-20 18:00:00', 3200.00, 36000000000000, 45, 80.0, 380, 1);

-- Route 13 (Yakutsk to Tiksi - summer)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_tiksi', 'route_yks_tiksi_air', 'air', 'Polar Airlines', 'yakutsk_yks', 'tiksi_aprt',
 '2025-06-15 08:00:00', '2025-06-15 11:00:00', 18000.00, 10800000000000, 80, 70.0, 1100, 1);

-- Route 14 (Neryungri to Aldan)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_nerungri_aldan', 'route_nerungri_aldan', 'bus', 'Magistral', 'nerungri_bus', 'aldan_bus',
 '2025-11-20 09:00:00', '2025-11-20 13:00:00', 1200.00, 14400000000000, 45, 88.0, 130, 1);

-- Route 15 (Yakutsk to Vilyuysk)
INSERT INTO segments (id, route_id, transport_type, provider, start_stop_id, end_stop_id,
                     departure_time, arrival_time, price, duration, seat_count,
                     reliability_rate, distance, sequence_order) VALUES
('seg_yks_vilyuysk', 'route_yks_vilyuysk', 'air', 'Polar Airlines', 'yakutsk_yks', 'vilyuysk_aprt',
 '2025-11-20 09:00:00', '2025-11-20 10:40:00', 10000.00, 6000000000000, 70, 75.0, 460, 1);

-- ============================================================================
-- CONNECTIONS (Transfers between segments)
-- ============================================================================

-- Route 3 connections (Transfer in Moscow)
INSERT INTO connections (id, route_id, from_segment_id, to_segment_id, connection_type,
                        waiting_time, transfer_distance, is_valid) VALUES
('conn_msk_yks_3_1', 'route_msk_yks_transfer', 'seg_msk_yks_3a', 'seg_msk_yks_3b',
 'airport_transfer', 3600000000000, 35000, true);

-- Route 4 connections (Taxi and then river)
INSERT INTO connections (id, route_id, from_segment_id, to_segment_id, connection_type,
                        waiting_time, transfer_distance, is_valid) VALUES
('conn_msk_olek_1', 'route_msk_olek_opt', 'seg_msk_olek_1', 'seg_msk_olek_2',
 'airport_transfer', 1800000000000, 20000, true),
('conn_msk_olek_2', 'route_msk_olek_opt', 'seg_msk_olek_2', 'seg_msk_olek_3',
 'city_transfer', 4800000000000, 500, true);

-- Route 5 connections (complex transfers)
INSERT INTO connections (id, route_id, from_segment_id, to_segment_id, connection_type,
                        waiting_time, transfer_distance, is_valid) VALUES
('conn_msk_olek_cheap_1', 'route_msk_olek_cheap', 'seg_msk_olek_cheap_1', 'seg_msk_olek_cheap_2',
 'airport_transfer', 1800000000000, 20000, true),
('conn_msk_olek_cheap_2', 'route_msk_olek_cheap', 'seg_msk_olek_cheap_2', 'seg_msk_olek_cheap_3',
 'city_transfer', 7200000000000, 500, true),
('conn_msk_olek_cheap_3', 'route_msk_olek_cheap', 'seg_msk_olek_cheap_3', 'seg_msk_olek_cheap_4',
 'terminal_transfer', 0, 100, true);

-- Route 6 connections (Sangur route)
INSERT INTO connections (id, route_id, from_segment_id, to_segment_id, connection_type,
                        waiting_time, transfer_distance, is_valid) VALUES
('conn_sangur_1', 'route_msk_sangur', 'seg_sangur_1', 'seg_sangur_2',
 'airport_transfer', 1800000000000, 20000, true),
('conn_sangur_2', 'route_msk_sangur', 'seg_sangur_2', 'seg_sangur_3',
 'city_transfer', 3600000000000, 500, true),
('conn_sangur_3', 'route_msk_sangur', 'seg_sangur_3', 'seg_sangur_4',
 'terminal_transfer', 3600000000000, 100, true);
