-- Make route_id nullable in segments table to support standalone segments from providers
-- This allows us to store individual transport segments (flights, buses, trains)
-- that are not yet part of a complete route

ALTER TABLE segments
ALTER COLUMN route_id DROP NOT NULL;

-- Add a comment to explain the change
COMMENT ON COLUMN segments.route_id IS 'Foreign key to routes table. Can be NULL for standalone segments from providers that are not yet part of a route.';
