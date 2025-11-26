-- Update booking statuses to include in_progress and completed
-- New statuses: pending, in_progress, completed, cancelled

-- First, drop the old constraint
ALTER TABLE bookings DROP CONSTRAINT IF EXISTS ck_booking_status;

-- Add new constraint with updated statuses
ALTER TABLE bookings ADD CONSTRAINT ck_booking_status CHECK (
    status IN ('pending', 'in_progress', 'completed', 'cancelled')
);

-- Update existing 'confirmed' to 'in_progress' (if any exist)
UPDATE bookings SET status = 'in_progress' WHERE status = 'confirmed';

-- Update existing 'failed' to 'cancelled' (if any exist)
UPDATE bookings SET status = 'cancelled' WHERE status = 'failed';

-- Update existing 'refunded' to 'cancelled' (if any exist)
UPDATE bookings SET status = 'cancelled' WHERE status = 'refunded';

-- Update comment
COMMENT ON COLUMN bookings.status IS 'Booking status: pending (awaiting payment), in_progress (confirmed and active), completed (finished), cancelled (user cancelled or failed)';
