CREATE TYPE status AS ENUM ('PAID', 'NOTPAID');

CREATE TABLE IF NOT EXISTS payments (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    reservation_id uuid NOT NULL,
    amount decimal not null,
    payment_method VARCHAR not null,
    payment_status status not null,
    created_at timestamp DEFAULT current_timestamp,
    update_at timestamp DEFAULT current_timestamp,
    deleted_at timestamp
)

