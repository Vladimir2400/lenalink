-- Remove user_id from bookings

DROP INDEX IF EXISTS idx_bookings_user;
ALTER TABLE bookings DROP CONSTRAINT IF EXISTS fk_bookings_user;
ALTER TABLE bookings DROP COLUMN IF EXISTS user_id;
