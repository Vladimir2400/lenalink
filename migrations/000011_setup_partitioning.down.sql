-- Rollback security and performance features

-- Disable row level security
ALTER TABLE routes DISABLE ROW LEVEL SECURITY;
ALTER TABLE bookings DISABLE ROW LEVEL SECURITY;
ALTER TABLE payments DISABLE ROW LEVEL SECURITY;

-- Drop the additional indexes (performance indexes from migration 8 remain)
DROP INDEX IF EXISTS idx_routes_saved_at;
DROP INDEX IF EXISTS idx_bookings_created_at;
DROP INDEX IF EXISTS idx_payments_created_at;

-- Restore original comments
COMMENT ON TABLE routes IS 'Complete multi-segment journeys with pricing and reliability';
COMMENT ON TABLE bookings IS 'Customer orders with embedded passenger data';
COMMENT ON TABLE payments IS 'Payment transactions 1:1 with bookings';
