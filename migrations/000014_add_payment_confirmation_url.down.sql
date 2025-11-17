-- Remove confirmation_url column
ALTER TABLE payments DROP COLUMN IF EXISTS confirmation_url;

-- Revert booking status constraint
ALTER TABLE bookings
    DROP CONSTRAINT IF EXISTS ck_booking_status;

ALTER TABLE bookings
    ADD CONSTRAINT ck_booking_status CHECK (
        status IN ('pending', 'confirmed', 'failed', 'cancelled', 'refunded')
    );

-- Remove index
DROP INDEX IF EXISTS idx_payments_provider_id_status;
