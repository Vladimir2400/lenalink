-- Revert booking statuses to original

-- Drop new constraint
ALTER TABLE bookings DROP CONSTRAINT IF EXISTS ck_booking_status;

-- Add old constraint
ALTER TABLE bookings ADD CONSTRAINT ck_booking_status CHECK (
    status IN ('pending', 'confirmed', 'failed', 'cancelled', 'refunded')
);

-- Revert comment
COMMENT ON COLUMN bookings.status IS 'Booking status: pending, confirmed, failed, cancelled, refunded';
