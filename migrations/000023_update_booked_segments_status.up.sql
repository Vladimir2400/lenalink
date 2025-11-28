-- Update booked_segments booking_status constraint to match new statuses

-- Drop old constraint
ALTER TABLE booked_segments DROP CONSTRAINT IF EXISTS ck_booked_segment_status;

-- Add new constraint with updated statuses matching domain.BookingStatus
ALTER TABLE booked_segments ADD CONSTRAINT ck_booked_segment_status CHECK (
    booking_status IN ('pending', 'pending_payment', 'in_progress', 'completed', 'cancelled', 'failed', 'refunded')
);

-- Update comment
COMMENT ON COLUMN booked_segments.booking_status IS 'Segment booking status: pending, pending_payment, in_progress (confirmed), completed, cancelled, failed, refunded';
