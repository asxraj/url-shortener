CREATE TABLE IF NOT EXISTS urls (
short_url text PRIMARY KEY NOT NULL,
long_url text NOT NULL,
user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
expires timestamp with time zone NOT NULL
);
