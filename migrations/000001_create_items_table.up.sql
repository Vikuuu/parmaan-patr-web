CREATE TABLE IF NOT EXISTS items (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    hsn integer NOT NULL,
    price integer NOT NULL,
    gst integer NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

