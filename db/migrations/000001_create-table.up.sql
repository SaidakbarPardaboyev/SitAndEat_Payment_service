CREATE TYPE IF NOT EXISTS Stat AS ENUM ('paid', 'notpaid');

CREATE TYPE IF NOT EXISTS Method AS ENUM ('cash', 'card');

CREATE TABLE IF NOT EXISTS Payments(
    id uuid PRIMARY KEY default gen_random_uuid(), 
    reservation_id uuid not null,
    amount decimal not null,
    payment_method Method,
    payment_status Stat,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)


