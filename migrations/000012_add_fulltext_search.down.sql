-- Drop full-text search indexes and functions
DROP TRIGGER IF EXISTS trigger_stops_search_vector ON stops CASCADE;
DROP FUNCTION IF EXISTS stops_search_vector_update() CASCADE;
DROP INDEX IF EXISTS idx_stops_search_vector CASCADE;
DROP INDEX IF EXISTS idx_routes_recent CASCADE;
DROP INDEX IF EXISTS idx_bookings_by_passenger CASCADE;
ALTER TABLE stops DROP COLUMN IF EXISTS search_vector;
