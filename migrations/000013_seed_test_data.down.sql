-- Rollback seed test data
-- This migration removes all test data inserted by 000013_seed_test_data.up.sql

-- Delete connections first (due to foreign keys)
DELETE FROM connections WHERE route_id LIKE 'route_%';

-- Delete segments
DELETE FROM segments WHERE route_id LIKE 'route_%';

-- Delete routes
DELETE FROM routes WHERE id LIKE 'route_%';

-- Delete stops (in reverse order of dependencies if needed)
DELETE FROM stops WHERE id IN (
  -- Moscow
  'moscow_dme', 'moscow_svo',
  -- Yakutia main
  'yakutsk_yks', 'yakutsk_port', 'yakutsk_bus', 'yakutsk_railway',
  -- Diamond cities
  'mirny_aprt', 'mirny_bus', 'udachny_aprt',
  -- Coal and mining
  'nerungri_aprt', 'nerungri_bus', 'aldan_aprt', 'aldan_bus',
  -- River ports
  'lensky_port', 'olekminsk_port', 'sangur_port', 'pokrovsk_port', 'batagay_port',
  -- Far north
  'tiksi_aprt', 'tiksi_port', 'verkhoyansk_aprt', 'oymyakon_aprt',
  -- Other cities
  'vilyuysk_aprt', 'nyurba_aprt', 'zhataay_bus',
  -- BAM railway
  'tynda_railway', 'tommot_railway'
);
