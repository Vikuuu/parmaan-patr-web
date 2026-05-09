ALTER TABLE items ADD CONSTRAINT items_hsn_check CHECK (hsn >= 0);
ALTER TABLE items ADD CONSTRAINT items_price_check CHECK (price >= 0);
ALTER TABLE items ADD CONSTRAINT items_gst_check CHECK (gst BETWEEN 0 and 100);
