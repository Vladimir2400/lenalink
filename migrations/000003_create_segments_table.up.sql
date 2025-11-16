-- Create SEGMENTS table
-- Individual transport legs within a route (flights, trains, buses, boats, etc.)

CREATE TABLE IF NOT EXISTS segments (
    id VARCHAR(36) PRIMARY KEY,
    route_id VARCHAR(36) NOT NULL,
    transport_type VARCHAR(20) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    start_stop_id VARCHAR(36) NOT NULL,
    end_stop_id VARCHAR(36) NOT NULL,
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    duration BIGINT NOT NULL,
    seat_count INTEGER NOT NULL,
    reliability_rate DECIMAL(5, 2) NOT NULL,
    distance INTEGER NOT NULL,
    sequence_order INTEGER NOT NULL,

    -- Foreign keys
    CONSTRAINT fk_segments_route FOREIGN KEY (route_id) REFERENCES routes(id) ON DELETE CASCADE,
    CONSTRAINT fk_segments_start_stop FOREIGN KEY (start_stop_id) REFERENCES stops(id) ON DELETE RESTRICT,
    CONSTRAINT fk_segments_end_stop FOREIGN KEY (end_stop_id) REFERENCES stops(id) ON DELETE RESTRICT,

    -- Check constraints
    CONSTRAINT ck_transport_type CHECK (
        transport_type IN ('air', 'rail', 'bus', 'river', 'taxi', 'walk')
    ),
    CONSTRAINT ck_segment_price_positive CHECK (price >= 0),
    CONSTRAINT ck_segment_seat_count_positive CHECK (seat_count >= 0),
    CONSTRAINT ck_segment_reliability_range CHECK (reliability_rate >= 0 AND reliability_rate <= 100),
    CONSTRAINT ck_segment_distance_positive CHECK (distance >= 0),
    CONSTRAINT ck_segment_times_logical CHECK (departure_time < arrival_time),
    CONSTRAINT ck_segment_sequence_positive CHECK (sequence_order > 0)
);

-- Create indexes
CREATE INDEX idx_segments_route ON segments(route_id, sequence_order);
CREATE INDEX idx_segments_transport ON segments(transport_type);
CREATE INDEX idx_segments_provider ON segments(provider);
CREATE INDEX idx_segments_times ON segments(departure_time, arrival_time);
CREATE INDEX idx_segments_start_stop ON segments(start_stop_id);
CREATE INDEX idx_segments_end_stop ON segments(end_stop_id);

-- Add comments
COMMENT ON TABLE segments IS 'Individual transport legs within routes';
COMMENT ON COLUMN segments.sequence_order IS 'Order of segment within the route (1, 2, 3, ...)';
COMMENT ON COLUMN segments.reliability_rate IS 'Provider reliability rating (0-100)';
