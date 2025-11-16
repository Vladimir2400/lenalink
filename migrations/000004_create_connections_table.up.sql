-- Create CONNECTIONS table
-- Tracks transfer/connection information between segments within a route

CREATE TABLE IF NOT EXISTS connections (
    id SERIAL PRIMARY KEY,
    route_id VARCHAR(36) NOT NULL,
    from_segment_id VARCHAR(36),
    to_segment_id VARCHAR(36),
    transfer_duration BIGINT NOT NULL,
    transfer_distance INTEGER NOT NULL DEFAULT 0,
    requires_transport BOOLEAN NOT NULL DEFAULT FALSE,
    is_valid BOOLEAN NOT NULL,
    gap BIGINT NOT NULL,
    sequence_order INTEGER NOT NULL,

    -- Foreign keys
    CONSTRAINT fk_connections_route FOREIGN KEY (route_id) REFERENCES routes(id) ON DELETE CASCADE,
    CONSTRAINT fk_connections_from_segment FOREIGN KEY (from_segment_id) REFERENCES segments(id) ON DELETE CASCADE,
    CONSTRAINT fk_connections_to_segment FOREIGN KEY (to_segment_id) REFERENCES segments(id) ON DELETE CASCADE,

    -- Check constraints
    CONSTRAINT ck_connection_transfer_distance CHECK (transfer_distance >= 0),
    CONSTRAINT ck_connection_sequence_positive CHECK (sequence_order >= 0)
);

-- Create indexes
CREATE INDEX idx_connections_route ON connections(route_id, sequence_order);
CREATE INDEX idx_connections_validity ON connections(is_valid);
CREATE INDEX idx_connections_from_segment ON connections(from_segment_id);
CREATE INDEX idx_connections_to_segment ON connections(to_segment_id);

-- Add comments
COMMENT ON TABLE connections IS 'Transfer/connection metadata between route segments';
COMMENT ON COLUMN connections.from_segment_id IS 'Nullable: first segment has no from_segment';
COMMENT ON COLUMN connections.to_segment_id IS 'Nullable: last segment has no to_segment';
COMMENT ON COLUMN connections.requires_transport IS 'Whether taxi/shuttle transfer needed';
COMMENT ON COLUMN connections.is_valid IS 'Connection validity (satisfies 1-24 hour rule)';
COMMENT ON COLUMN connections.gap IS 'Time gap in nanoseconds between arrival and next departure';
