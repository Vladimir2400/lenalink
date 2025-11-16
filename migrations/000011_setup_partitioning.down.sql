-- Disable row level security (partitions will be automatically dropped when parent table constraints change)
ALTER TABLE routes DISABLE ROW LEVEL SECURITY;
ALTER TABLE bookings DISABLE ROW LEVEL SECURITY;
ALTER TABLE payments DISABLE ROW LEVEL SECURITY;

-- Note: Individual partitions don't need to be explicitly dropped
-- They will be cleaned up when the partitioning is reconfigured
