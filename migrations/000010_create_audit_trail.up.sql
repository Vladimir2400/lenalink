-- Create audit trail tables for tracking status changes

-- Booking status audit table
CREATE TABLE IF NOT EXISTS booking_status_audit (
    id SERIAL PRIMARY KEY,
    booking_id VARCHAR(36) NOT NULL,
    old_status VARCHAR(20),
    new_status VARCHAR(20) NOT NULL,
    changed_by VARCHAR(255) DEFAULT 'system',
    changed_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_audit_booking FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE CASCADE
);

-- Payment status audit table
CREATE TABLE IF NOT EXISTS payment_status_audit (
    id SERIAL PRIMARY KEY,
    payment_id VARCHAR(36) NOT NULL,
    order_id VARCHAR(36) NOT NULL,
    old_status VARCHAR(20),
    new_status VARCHAR(20) NOT NULL,
    changed_by VARCHAR(255) DEFAULT 'system',
    changed_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_audit_payment FOREIGN KEY (payment_id) REFERENCES payments(id) ON DELETE CASCADE
);

-- Indexes on audit tables for efficient querying
CREATE INDEX idx_booking_audit_booking_id ON booking_status_audit(booking_id);
CREATE INDEX idx_booking_audit_changed_at ON booking_status_audit(changed_at DESC);
CREATE INDEX idx_payment_audit_payment_id ON payment_status_audit(payment_id);
CREATE INDEX idx_payment_audit_order_id ON payment_status_audit(order_id);
CREATE INDEX idx_payment_audit_changed_at ON payment_status_audit(changed_at DESC);

-- Create function to log booking status changes
CREATE OR REPLACE FUNCTION audit_booking_status_change()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status IS DISTINCT FROM OLD.status THEN
        INSERT INTO booking_status_audit(booking_id, old_status, new_status, changed_at)
        VALUES(NEW.id, OLD.status, NEW.status, NOW());
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create function to log payment status changes
CREATE OR REPLACE FUNCTION audit_payment_status_change()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status IS DISTINCT FROM OLD.status THEN
        INSERT INTO payment_status_audit(payment_id, order_id, old_status, new_status, changed_at)
        VALUES(NEW.id, NEW.order_id, OLD.status, NEW.status, NOW());
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create triggers for audit logging
CREATE TRIGGER trigger_audit_booking_status
    AFTER UPDATE ON bookings
    FOR EACH ROW
    EXECUTE FUNCTION audit_booking_status_change();

CREATE TRIGGER trigger_audit_payment_status
    AFTER UPDATE ON payments
    FOR EACH ROW
    EXECUTE FUNCTION audit_payment_status_change();

-- Add comments
COMMENT ON TABLE booking_status_audit IS 'Audit trail for booking status changes';
COMMENT ON TABLE payment_status_audit IS 'Audit trail for payment status changes';
COMMENT ON FUNCTION audit_booking_status_change() IS 'Logs all booking status changes to audit table';
COMMENT ON FUNCTION audit_payment_status_change() IS 'Logs all payment status changes to audit table';
