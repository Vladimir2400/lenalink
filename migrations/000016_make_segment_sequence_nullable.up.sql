-- Make sequence_order nullable in segments table to support standalone segments from providers
-- sequence_order only makes sense in the context of a route

-- First, drop the check constraint that requires sequence_order > 0
ALTER TABLE segments DROP CONSTRAINT IF EXISTS ck_segment_sequence_positive;

-- Make sequence_order nullable
ALTER TABLE segments
ALTER COLUMN sequence_order DROP NOT NULL;

-- Add a new check constraint that only validates when sequence_order is not null
ALTER TABLE segments
ADD CONSTRAINT ck_segment_sequence_positive
CHECK (sequence_order IS NULL OR sequence_order > 0);

-- Add a comment to explain the change
COMMENT ON COLUMN segments.sequence_order IS 'Order of segment within a route (1-based). NULL for standalone segments not part of a route.';
