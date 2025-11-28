-- Rollback tariff and seat selection changes

-- Drop seat selection table
DROP TABLE IF EXISTS booked_segment_seats;

-- Drop tariff columns from bookings
ALTER TABLE bookings
    DROP CONSTRAINT IF EXISTS ck_booking_tariff,
    DROP CONSTRAINT IF EXISTS ck_booking_tariff_price,
    DROP COLUMN IF EXISTS tariff,
    DROP COLUMN IF EXISTS tariff_price;

-- Drop indexes
DROP INDEX IF EXISTS idx_bookings_tariff;
