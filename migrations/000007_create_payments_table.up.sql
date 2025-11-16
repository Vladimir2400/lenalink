-- Create PAYMENTS table
-- Payment transaction records

CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'RUB',
    method VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    provider_payment_id VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP,
    failure_reason TEXT,

    -- Foreign keys
    CONSTRAINT fk_payments_order FOREIGN KEY (order_id) REFERENCES bookings(id) ON DELETE CASCADE,

    -- Unique constraint (one payment per booking)
    CONSTRAINT unique_payment_per_order UNIQUE (order_id),

    -- Check constraints
    CONSTRAINT ck_payment_amount_positive CHECK (amount >= 0),
    CONSTRAINT ck_payment_method CHECK (
        method IN ('card', 'yookassa', 'cloudpay', 'sberpay')
    ),
    CONSTRAINT ck_payment_status CHECK (
        status IN ('pending', 'completed', 'failed', 'refunded')
    )
);

-- Create indexes
CREATE INDEX idx_payments_order ON payments(order_id);
CREATE INDEX idx_payments_status ON payments(status);
CREATE INDEX idx_payments_provider_id ON payments(provider_payment_id) WHERE provider_payment_id IS NOT NULL;
CREATE INDEX idx_payments_created ON payments(created_at DESC);

-- Add comments
COMMENT ON TABLE payments IS 'Payment transaction records';
COMMENT ON COLUMN payments.method IS 'Payment method: card, yookassa, cloudpay, sberpay';
COMMENT ON COLUMN payments.status IS 'Payment status: pending, completed, failed, refunded';
COMMENT ON COLUMN payments.provider_payment_id IS 'Payment gateway transaction ID';
COMMENT ON COLUMN payments.currency IS 'ISO 4217 currency code (default RUB for Russian rubles)';
