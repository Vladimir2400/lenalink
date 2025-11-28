-- Add tariff and seat selection to bookings
-- Tariffs: tarif1 (0), tarif2 (2244), tarif3 (4244), tarif4 (6244)
-- Seat types: random (0), window (3000), aisle (2300), extra_legroom (7900)

-- Add tariff columns to bookings table
ALTER TABLE bookings
    ADD COLUMN tariff VARCHAR(20) NOT NULL DEFAULT 'tarif1',
    ADD COLUMN tariff_price DECIMAL(10, 2) NOT NULL DEFAULT 0,
    ADD CONSTRAINT ck_booking_tariff CHECK (
        tariff IN ('tarif1', 'tarif2', 'tarif3', 'tarif4')
    ),
    ADD CONSTRAINT ck_booking_tariff_price CHECK (tariff_price >= 0);

-- Create seat selection table for air segments
CREATE TABLE IF NOT EXISTS booked_segment_seats (
    id VARCHAR(36) PRIMARY KEY,
    booked_segment_id VARCHAR(36) NOT NULL,
    seat_type VARCHAR(20) NOT NULL,
    seat_price DECIMAL(10, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- Foreign key
    CONSTRAINT fk_segment_seats_booked_segment
        FOREIGN KEY (booked_segment_id)
        REFERENCES booked_segments(id)
        ON DELETE CASCADE,

    -- Check constraints
    CONSTRAINT ck_seat_type CHECK (
        seat_type IN ('random', 'window', 'aisle', 'extra_legroom')
    ),
    CONSTRAINT ck_seat_price_positive CHECK (seat_price >= 0)
);

-- Create indexes
CREATE INDEX idx_segment_seats_booked_segment ON booked_segment_seats(booked_segment_id);
CREATE INDEX idx_bookings_tariff ON bookings(tariff);

-- Add comments
COMMENT ON COLUMN bookings.tariff IS 'Tariff type: tarif1 (0₽), tarif2 (2244₽), tarif3 (4244₽), tarif4 (6244₽)';
COMMENT ON COLUMN bookings.tariff_price IS 'Price of selected tariff';
COMMENT ON TABLE booked_segment_seats IS 'Seat selections for air segments in bookings';
COMMENT ON COLUMN booked_segment_seats.seat_type IS 'Seat type: random (0₽), window (3000₽), aisle (2300₽), extra_legroom (7900₽)';
COMMENT ON COLUMN booked_segment_seats.seat_price IS 'Price of selected seat type';
