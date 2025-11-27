-- ============================================================================
-- DROP ALL DATA FROM DATABASE
-- WARNING: This will delete ALL data from all tables
-- ============================================================================

BEGIN;

-- Disable triggers temporarily
SET session_replication_role = 'replica';

-- Truncate all tables in correct order (respecting foreign keys)
TRUNCATE TABLE
    payment_status_audit,
    booking_status_audit,
    payments,
    booked_segments,
    bookings,
    connections,
    segments,
    routes,
    stops,
    users
RESTART IDENTITY CASCADE;

-- Re-enable triggers
SET session_replication_role = 'origin';

COMMIT;

SELECT 'All data dropped successfully!' AS status;
