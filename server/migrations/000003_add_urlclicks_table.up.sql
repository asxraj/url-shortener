CREATE TABLE IF NOT EXISTS urlclicks (
urls_short_url text NOT NULL REFERENCES urls ON DELETE CASCADE,
ip_address text NOT NULL,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
