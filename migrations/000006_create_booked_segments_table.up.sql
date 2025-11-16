-- Create BOOKED_SEGMENTS table
-- Individual segment bookings within a multi-segment order

CREATE TABLE IF NOT EXISTS booked_segments (
    id VARCHAR(36) PRIMARY KEY,
    booking_id VARCHAR(36) NOT NULL,
    segment_id VARCHAR(36) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    transport_type VARCHAR(20) NOT NULL,
    from_stop_id VARCHAR(36) NOT NULL,
    to_stop_id VARCHAR(36) NOT NULL,
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    ticket_number VARCHAR(255),
    price DECIMAL(10, 2) NOT NULL,
    commission DECIMAL(10, 2) NOT NULL DEFAULT 0,
    total_price DECIMAL(10, 2) NOT NULL,
    booking_status VARCHAR(20) NOT NULL DEFAULT 'pending',
    provider_booking_ref VARCHAR(255),
    sequence_order INTEGER NOT NULL,

    -- Foreign keys
    CONSTRAINT fk_booked_segments_booking FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE CASCADE,
    CONSTRAINT fk_booked_segments_segment FOREIGN KEY (segment_id) REFERENCES segments(id) ON DELETE RESTRICT,
    CONSTRAINT fk_booked_segments_from_stop FOREIGN KEY (from_stop_id) REFERENCES stops(id) ON DELETE RESTRICT,
    CONSTRAINT fk_booked_segments_to_stop FOREIGN KEY (to_stop_id) REFERENCES stops(id) ON DELETE RESTRICT,

    -- Check constraints
    CONSTRAINT ck_booked_segment_price_positive CHECK (price >= 0),
    CONSTRAINT ck_booked_segment_commission_positive CHECK (commission >= 0),
    CONSTRAINT ck_booked_segment_total_positive CHECK (total_price >= 0),
    CONSTRAINT ck_booked_segment_status CHECK (
        booking_status IN ('pending', 'confirmed', 'failed', 'cancelled', 'refunded')
    ),
    CONSTRAINT ck_booked_segment_times_logical CHECK (departure_time < arrival_time),
    CONSTRAINT ck_booked_segment_sequence_positive CHECK (sequence_order >= 1)
);

-- Create indexes
CREATE INDEX idx_booked_segments_booking ON booked_segments(booking_id, sequence_order);
CREATE INDEX idx_booked_segments_ticket ON booked_segments(ticket_number) WHERE ticket_number IS NOT NULL;
CREATE INDEX idx_booked_segments_status ON booked_segments(booking_status);
CREATE INDEX idx_booked_segments_provider_ref ON booked_segments(provider_booking_ref) WHERE provider_booking_ref IS NOT NULL;
CREATE INDEX idx_booked_segments_segment ON booked_segments(segment_id);

-- Add comments
COMMENT ON TABLE booked_segments IS 'Individual segment bookings within a multi-segment order';
COMMENT ON COLUMN booked_segments.ticket_number IS 'Ticket number issued by provider';
COMMENT ON COLUMN booked_segments.provider_booking_ref IS 'Provider''s PNR or booking reference';
COMMENT ON COLUMN booked_segments.sequence_order IS 'Order of segment within the booking';
