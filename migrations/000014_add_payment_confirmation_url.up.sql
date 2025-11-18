-- Add confirmation_url column to payments table
ALTER TABLE payments ADD COLUMN confirmation_url TEXT;

-- Add pending_payment status to bookings
ALTER TABLE bookings
    DROP CONSTRAINT IF EXISTS ck_booking_status;

ALTER TABLE bookings
    ADD CONSTRAINT ck_booking_status CHECK (
        status IN ('pending', 'pending_payment', 'confirmed', 'failed', 'cancelled', 'refunded')
    );

-- Add index for webhook lookups
CREATE INDEX idx_payments_provider_id_status ON payments(provider_payment_id, status);

COMMENT ON COLUMN payments.confirmation_url IS 'URL for redirecting user to payment provider (YooKassa)';
