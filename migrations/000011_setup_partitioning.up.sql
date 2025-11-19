-- Setup table security and performance features
-- Note: Partitioning requires tables to be created with PARTITION BY clause
-- Since our tables are already created as regular tables, we'll skip partitioning
-- and only enable Row Level Security for future use

-- Enable Row Level Security on ROUTES
-- (Policies can be added later for multi-tenancy)
ALTER TABLE routes
    ENABLE ROW LEVEL SECURITY;

-- Enable Row Level Security on BOOKINGS
ALTER TABLE bookings
    ENABLE ROW LEVEL SECURITY;

-- Enable Row Level Security on PAYMENTS
ALTER TABLE payments
    ENABLE ROW LEVEL SECURITY;

-- Create indexes for better query performance (if not exist)
CREATE INDEX IF NOT EXISTS idx_routes_saved_at ON routes(saved_at DESC);
CREATE INDEX IF NOT EXISTS idx_bookings_created_at ON bookings(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_payments_created_at ON payments(created_at DESC);

-- Add comments
COMMENT ON TABLE routes IS 'Complete multi-segment journeys with pricing and reliability (RLS enabled)';
COMMENT ON TABLE bookings IS 'Customer orders with embedded passenger data (RLS enabled)';
COMMENT ON TABLE payments IS 'Payment transactions 1:1 with bookings (RLS enabled)';

-- Note: For production with high volume, consider recreating tables with native partitioning:
-- CREATE TABLE routes (...) PARTITION BY RANGE (saved_at);
-- This would require data migration and is beyond the scope of this migration
