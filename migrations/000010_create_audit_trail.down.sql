-- Drop audit trail tables and functions
DROP TRIGGER IF EXISTS trigger_audit_booking_status ON bookings CASCADE;
DROP TRIGGER IF EXISTS trigger_audit_payment_status ON payments CASCADE;
DROP FUNCTION IF EXISTS audit_booking_status_change() CASCADE;
DROP FUNCTION IF EXISTS audit_payment_status_change() CASCADE;
DROP TABLE IF EXISTS booking_status_audit CASCADE;
DROP TABLE IF EXISTS payment_status_audit CASCADE;
