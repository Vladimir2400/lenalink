-- Drop triggers and function
DROP TRIGGER IF EXISTS trigger_update_bookings_updated_at ON bookings CASCADE;
DROP FUNCTION IF EXISTS update_updated_at_column() CASCADE;
