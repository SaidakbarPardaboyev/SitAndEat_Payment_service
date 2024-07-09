CREATE TYPE status AS ENUM ('COMPLETED', 'FAILED', 'PENDING');

CREATE TABLE IF NOT EXISTS payments (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    reservation_id uuid REFERENCES reservations(id),
    amount decimal,
    payment_method VARCHAR,
    payment_status status,
    created_at timestamp DEFAULT current_timestamp,
    update_at timestamp DEFAULT current_timestamp,
    deleted_at timestamp
)

