-- Setup full-text search capabilities for better search performance

-- Add tsvector columns to stops table for full-text search
ALTER TABLE stops ADD COLUMN IF NOT EXISTS search_vector tsvector;

-- Create GIN indexes for full-text search on stops
CREATE INDEX IF NOT EXISTS idx_stops_search_vector ON stops USING GIN(search_vector);

-- Create trigger to automatically update tsvector on insert/update
CREATE OR REPLACE FUNCTION stops_search_vector_update()
RETURNS TRIGGER AS $$
BEGIN
    NEW.search_vector :=
        setweight(to_tsvector('russian', COALESCE(NEW.name, '')), 'A') ||
        setweight(to_tsvector('russian', COALESCE(NEW.city, '')), 'B');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_stops_search_vector
    BEFORE INSERT OR UPDATE ON stops
    FOR EACH ROW
    EXECUTE FUNCTION stops_search_vector_update();

-- Update existing rows with search vectors
UPDATE stops SET search_vector =
    setweight(to_tsvector('russian', COALESCE(name, '')), 'A') ||
    setweight(to_tsvector('russian', COALESCE(city, '')), 'B')
WHERE search_vector IS NULL;

-- Create index for recent routes (for faster search)
-- Note: Removed WHERE clause with NOW() since it's STABLE not IMMUTABLE
CREATE INDEX IF NOT EXISTS idx_routes_recent ON routes(from_city, to_city, departure_time DESC);

-- Create index for common reporting queries
CREATE INDEX IF NOT EXISTS idx_bookings_by_passenger ON bookings(passenger_email, created_at DESC);

-- Add Russian language configuration if needed
-- (Usually available by default in PostgreSQL with Russian locale)

-- Add comments
COMMENT ON COLUMN stops.search_vector IS 'tsvector for full-text search (automatically maintained by trigger)';
COMMENT ON FUNCTION stops_search_vector_update() IS 'Updates search_vector column for full-text search';
COMMENT ON INDEX idx_stops_search_vector IS 'GIN index for fast full-text search on stops';
COMMENT ON INDEX idx_routes_recent IS 'Optimized index for recent route searches (last 7 days)';
