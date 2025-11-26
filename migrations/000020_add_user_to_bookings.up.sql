-- Add user_id to bookings table to link bookings with authenticated users

ALTER TABLE bookings ADD COLUMN user_id VARCHAR(36);

-- Add foreign key constraint
ALTER TABLE bookings ADD CONSTRAINT fk_bookings_user
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;

-- Create index for user bookings lookup
CREATE INDEX idx_bookings_user ON bookings(user_id);

-- Add comment
COMMENT ON COLUMN bookings.user_id IS 'User who created this booking (nullable for backwards compatibility)';
