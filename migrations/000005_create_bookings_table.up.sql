-- Create BOOKINGS table
-- Complete customer orders for multi-segment journeys

CREATE TABLE IF NOT EXISTS bookings (
    id VARCHAR(36) PRIMARY KEY,
    route_id VARCHAR(36) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    total_price DECIMAL(10, 2) NOT NULL,
    total_commission DECIMAL(10, 2) NOT NULL DEFAULT 0,
    grand_total DECIMAL(10, 2) NOT NULL,
    insurance_premium DECIMAL(10, 2) DEFAULT 0,
    include_insurance BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    confirmed_at TIMESTAMP,
    cancelled_at TIMESTAMP,
    cancellation_reason TEXT,

    -- Embedded PASSENGER fields
    passenger_first_name VARCHAR(255) NOT NULL,
    passenger_last_name VARCHAR(255) NOT NULL,
    passenger_middle_name VARCHAR(255),
    passenger_date_of_birth DATE NOT NULL,
    passenger_passport_number VARCHAR(50) NOT NULL,
    passenger_email VARCHAR(255) NOT NULL,
    passenger_phone VARCHAR(50) NOT NULL,

    -- Foreign keys
    CONSTRAINT fk_bookings_route FOREIGN KEY (route_id) REFERENCES routes(id) ON DELETE RESTRICT,

    -- Check constraints
    CONSTRAINT ck_booking_status CHECK (
        status IN ('pending', 'confirmed', 'failed', 'cancelled', 'refunded')
    ),
    CONSTRAINT ck_booking_total_price CHECK (total_price >= 0),
    CONSTRAINT ck_booking_total_commission CHECK (total_commission >= 0),
    CONSTRAINT ck_booking_grand_total CHECK (grand_total >= 0),
    CONSTRAINT ck_booking_insurance_premium CHECK (insurance_premium >= 0),
    CONSTRAINT ck_booking_passenger_name_not_empty CHECK (
        LENGTH(TRIM(passenger_first_name)) > 0 AND LENGTH(TRIM(passenger_last_name)) > 0
    )
);

-- Create indexes
CREATE INDEX idx_bookings_status ON bookings(status);
CREATE INDEX idx_bookings_created ON bookings(created_at DESC);
CREATE INDEX idx_bookings_email ON bookings(passenger_email);
CREATE INDEX idx_bookings_phone ON bookings(passenger_phone);
CREATE INDEX idx_bookings_route ON bookings(route_id);

-- Partial index for confirmed bookings
CREATE INDEX idx_bookings_confirmed ON bookings(id) WHERE status = 'confirmed';

-- Add comments
COMMENT ON TABLE bookings IS 'Complete customer orders for multi-segment journeys';
COMMENT ON COLUMN bookings.status IS 'Booking status: pending, confirmed, failed, cancelled, refunded';
COMMENT ON COLUMN bookings.grand_total IS 'Total with commission and optional insurance';
COMMENT ON COLUMN bookings.passenger_email IS 'Used for customer lookup and confirmation';
COMMENT ON COLUMN bookings.passenger_phone IS 'Contact phone in Russian format (+79XX or 89XX)';
