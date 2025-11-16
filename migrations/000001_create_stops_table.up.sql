-- Create STOPS table
-- Stores all transportation stops/stations (airports, train stations, bus terminals, river ports)

CREATE TABLE IF NOT EXISTS stops (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    latitude DECIMAL(10, 7) NOT NULL,
    longitude DECIMAL(10, 7) NOT NULL,
    stop_type VARCHAR(20) DEFAULT NULL
);

-- Create UNIQUE constraint on (name, city) to prevent duplicate stops in same city
ALTER TABLE stops ADD CONSTRAINT unique_stop_per_city UNIQUE (name, city);

-- Create indexes for common queries
CREATE INDEX idx_stops_city ON stops(city);
CREATE INDEX idx_stops_name ON stops(name);
CREATE INDEX idx_stops_coords ON stops(latitude, longitude);

-- Create GIN index for full-text search (tsvector will be added in migration 12)
CREATE INDEX idx_stops_name_gin ON stops USING GIN (to_tsvector('russian', name));
CREATE INDEX idx_stops_city_gin ON stops USING GIN (to_tsvector('russian', city));

-- Add comment
COMMENT ON TABLE stops IS 'Transportation stops/stations (airports, train stations, bus terminals, river ports)';
COMMENT ON COLUMN stops.id IS 'Unique identifier (UUID format)';
COMMENT ON COLUMN stops.stop_type IS 'Type of stop: airport, station, port, terminal';
