CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    phone VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL CHECK (role IN ('user', 'driver', 'admin')),
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE drivers (
    id UUID PRIMARY KEY REFERENCES users(id),
    license_number VARCHAR,
    ktp_number VARCHAR,
    status VARCHAR DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'rejected')),
    vehicle_info TEXT,
    current_lat DOUBLE PRECISION,
    current_lng DOUBLE PRECISION,
    is_active BOOLEAN DEFAULT FALSE
);

CREATE TABLE trips (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    driver_id UUID REFERENCES users(id),
    pickup_lat DOUBLE PRECISION,
    pickup_lng DOUBLE PRECISION,
    drop_lat DOUBLE PRECISION,
    drop_lng DOUBLE PRECISION,
    status VARCHAR NOT NULL CHECK (status IN ('requested', 'accepted', 'ongoing', 'completed', 'cancelled')),
    distance_km DOUBLE PRECISION,
    price VARCHAR,
    is_paid BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP,
    completed_at TIMESTAMP
);


CREATE TABLE payments (
    id UUID PRIMARY KEY,
    trip_id UUID REFERENCES trips(id),
    user_id UUID REFERENCES users(id),
    payment_method VARCHAR,
    amount NUMERIC,
    status VARCHAR CHECK (status IN ('pending', 'success', 'failed')),
    paid_at TIMESTAMP
);


CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) UNIQUE,
    balance NUMERIC DEFAULT 0
);

CREATE TABLE trip_ratings (
    id UUID PRIMARY KEY,
    trip_id UUID REFERENCES trips(id),
    from_user_id UUID REFERENCES users(id),
    to_user_id UUID REFERENCES users(id),
    rating INT CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE driver_documents (
    id UUID PRIMARY KEY,
    driver_id UUID REFERENCES drivers(id),
    doc_type VARCHAR, -- e.g. KTP, SIM
    file_url TEXT,
    is_verified BOOLEAN DEFAULT FALSE,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
