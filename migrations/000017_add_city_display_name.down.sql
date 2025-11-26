-- Remove city_display_name field

DROP INDEX IF EXISTS idx_stops_city_display_gin;
DROP INDEX IF EXISTS idx_stops_city_display;

ALTER TABLE stops DROP COLUMN IF EXISTS city_display_name;
