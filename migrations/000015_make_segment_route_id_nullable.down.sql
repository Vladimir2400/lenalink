-- Revert route_id to NOT NULL
-- WARNING: This will fail if there are segments with NULL route_id

-- First, delete any segments with NULL route_id
DELETE FROM segments WHERE route_id IS NULL;

-- Then restore the NOT NULL constraint
ALTER TABLE segments
ALTER COLUMN route_id SET NOT NULL;

-- Remove the comment
COMMENT ON COLUMN segments.route_id IS NULL;
