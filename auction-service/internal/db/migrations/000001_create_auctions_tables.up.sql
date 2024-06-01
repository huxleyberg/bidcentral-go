-- up.sql

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE auctions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    reserve_price INT,
    seller VARCHAR(255),
    winner VARCHAR(255),
    sold_amount INT,
    current_high_bid INT,
    auction_end TIMESTAMP,
    status VARCHAR(20)
);

CREATE TABLE items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    make VARCHAR(255),
    model VARCHAR(255),
    year INT,
    color VARCHAR(255),
    mileage INT,
    image_url VARCHAR(255),
    auction_id UUID REFERENCES auctions(id)
);
