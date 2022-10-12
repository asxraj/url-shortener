CREATE TABLE IF NOT EXISTS urls (
id bigserial PRIMARY KEY,
long_url text NOT NULL,
short_url text NOT NULL,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
expires timestamp with time zone NOT NULL
);

ALTER TABLE urls ADD COLUMN user_id bigint REFERENCES users ON DELETE CASCADE;
