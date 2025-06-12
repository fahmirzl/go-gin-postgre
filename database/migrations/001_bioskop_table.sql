-- +migrate Up
CREATE TABLE IF NOT EXISTS bioskop (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    lokasi VARCHAR(150) NOT NULL,
    rating FLOAT
);

-- +migrate Down
DROP TABLE IF EXISTS bioskop;