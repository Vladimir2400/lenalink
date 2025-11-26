-- Add city_display_name field for human-readable city names
-- This separates the internal city code from the display name

ALTER TABLE stops ADD COLUMN city_display_name VARCHAR(255);

-- Update existing data: copy city to city_display_name for now
UPDATE stops SET city_display_name = city;

-- Make it NOT NULL after populating
ALTER TABLE stops ALTER COLUMN city_display_name SET NOT NULL;

-- Add index for searching by display name
CREATE INDEX idx_stops_city_display ON stops(city_display_name);
CREATE INDEX idx_stops_city_display_gin ON stops USING GIN (to_tsvector('russian', city_display_name));

-- Add comment
COMMENT ON COLUMN stops.city_display_name IS 'Human-readable city name for display in UI';
