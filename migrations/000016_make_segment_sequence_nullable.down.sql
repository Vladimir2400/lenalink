-- Revert sequence_order to NOT NULL
-- WARNING: This will fail if there are segments with NULL sequence_order

-- First, delete any segments with NULL sequence_order
DELETE FROM segments WHERE sequence_order IS NULL;

-- Drop the new constraint
ALTER TABLE segments DROP CONSTRAINT IF EXISTS ck_segment_sequence_positive;

-- Restore the NOT NULL constraint
ALTER TABLE segments
ALTER COLUMN sequence_order SET NOT NULL;

-- Restore the original check constraint
ALTER TABLE segments
ADD CONSTRAINT ck_segment_sequence_positive
CHECK (sequence_order > 0);

-- Remove the comment
COMMENT ON COLUMN segments.sequence_order IS NULL;
