-- Rollback booked_segments status constraint

-- Drop new constraint
ALTER TABLE booked_segments DROP CONSTRAINT IF EXISTS ck_booked_segment_status;

-- Restore old constraint
ALTER TABLE booked_segments ADD CONSTRAINT ck_booked_segment_status CHECK (
    booking_status IN ('pending', 'confirmed', 'failed', 'cancelled', 'refunded')
);
