-- Drop performance indexes
DROP INDEX IF EXISTS idx_routes_city_date CASCADE;
DROP INDEX IF EXISTS idx_bookings_status_created CASCADE;
DROP INDEX IF EXISTS idx_bookings_passenger_created CASCADE;
DROP INDEX IF EXISTS idx_payments_completed CASCADE;
DROP INDEX IF EXISTS idx_bookings_created_brin CASCADE;
DROP INDEX IF EXISTS idx_payments_created_brin CASCADE;
DROP INDEX IF EXISTS idx_bookings_email_lower CASCADE;
DROP INDEX IF EXISTS idx_bookings_report CASCADE;
