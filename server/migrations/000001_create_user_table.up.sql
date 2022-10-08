CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    username text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);