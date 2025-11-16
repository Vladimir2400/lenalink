-- Create ROUTES table
-- Stores complete multi-segment journeys with pricing and reliability information

CREATE TABLE IF NOT EXISTS routes (
    id VARCHAR(36) PRIMARY KEY,
    from_city VARCHAR(255) NOT NULL,
    to_city VARCHAR(255) NOT NULL,
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    total_duration BIGINT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    reliability_score DECIMAL(5, 2) NOT NULL DEFAULT 0,
    insurance_premium DECIMAL(10, 2) NOT NULL DEFAULT 0,
    insurance_included BOOLEAN NOT NULL DEFAULT FALSE,
    transport_types TEXT[] NOT NULL DEFAULT '{}',
    saved_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- Add CHECK constraints for data integrity
    CONSTRAINT ck_total_price_positive CHECK (total_price >= 0),
    CONSTRAINT ck_insurance_premium_positive CHECK (insurance_premium >= 0),
    CONSTRAINT ck_reliability_score_range CHECK (reliability_score >= 0 AND reliability_score <= 100),
    CONSTRAINT ck_times_logical CHECK (departure_time < arrival_time),
    CONSTRAINT ck_from_to_different CHECK (from_city != to_city)
);

-- Create indexes for common search patterns
-- Primary search use case: from_city, to_city, departure_date
CREATE INDEX idx_routes_from_to ON routes(from_city, to_city);
CREATE INDEX idx_routes_departure ON routes(departure_time DESC);
CREATE INDEX idx_routes_price ON routes(total_price);
CREATE INDEX idx_routes_saved_at ON routes(saved_at DESC);

-- Composite index for most common search pattern
CREATE INDEX idx_routes_search ON routes(from_city, to_city, departure_time DESC);

-- Index for reliability scoring
CREATE INDEX idx_routes_reliability ON routes(reliability_score DESC);

-- Add comments
COMMENT ON TABLE routes IS 'Complete multi-segment journeys with pricing and reliability';
COMMENT ON COLUMN routes.transport_types IS 'Array of transport types used in route (air, rail, bus, river, taxi, walk)';
COMMENT ON COLUMN routes.saved_at IS 'When route was cached (for cache expiration)';
