-- Add additional performance indexes and composite indexes

-- Additional composite indexes for common query patterns
CREATE INDEX idx_routes_city_date ON routes(from_city, to_city, departure_time DESC);

-- Index for fast lookups by status groups
CREATE INDEX idx_bookings_status_created ON bookings(status, created_at DESC);

-- Indexes for analytics queries
CREATE INDEX idx_bookings_passenger_created ON bookings(passenger_email, created_at DESC);
CREATE INDEX idx_payments_completed ON payments(status, completed_at DESC);

-- BRIN (Block Range Index) for large tables with time-series data
-- More efficient for large tables that are mostly sorted by timestamp
CREATE INDEX idx_bookings_created_brin ON bookings USING BRIN (created_at)
    WITH (pages_per_range = 128);
CREATE INDEX idx_payments_created_brin ON payments USING BRIN (created_at)
    WITH (pages_per_range = 128);

-- Functional index for case-insensitive email search
CREATE INDEX idx_bookings_email_lower ON bookings(LOWER(passenger_email));

-- Multi-column index for common reporting queries
CREATE INDEX idx_bookings_report ON bookings(status, created_at, route_id);

-- Add comments
COMMENT ON INDEX idx_routes_city_date IS 'Optimized for route searches by city and date';
COMMENT ON INDEX idx_bookings_status_created IS 'Fast status filtering with time ordering';
COMMENT ON INDEX idx_bookings_email_lower IS 'Case-insensitive email search';
