CREATE TABLE routes (
    id UUID DEFAULT gen_random_uuid(),
    route_name VARCHAR(100) NOT NULL,
    route_type VARCHAR(50) NOT NULL,
    start_point POINT NOT NULL,
    end_point POINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE route_schedules (
    id UUID DEFAULT gen_random_uuid(),
    route_id UUID REFERENCES routes(route_id) ON DELETE CASCADE,
    departure_time TIME NOT NULL,
    arrival_time TIME NOT NULL,
    day_of_week SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE vehicles (
    id UUID DEFAULT gen_random_uuid(),
    vehicle_type VARCHAR(50) NOT NULL,
    capacity INTEGER NOT NULL,
    current_location POINT,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE parking_lots (
    id UUID DEFAULT gen_random_uuid(),
    lot_name VARCHAR(100) NOT NULL,
    total_spaces INTEGER NOT NULL,
    free_space INTEGER NOT NULL,
    occupied_space INTEGER NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0

);

CREATE TABLE parking_spaces (
    id UUID DEFAULT gen_random_uuid(),
    lot_id INTEGER REFERENCES parking_lots(lot_id) ON DELETE CASCADE,
    space_number VARCHAR(10) NOT NULL,
    is_occupied BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE(lot_id, space_number),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS reservation_parking (
    id UUID DEFAULT gen_random_uuid(),
    parking_space_id UUID REFERENCES parking_spaces(id) ON DELETE CASCADE,
    citizen_id UUID NOT NULL,
    time_from TIMESTAMP DEFAULT NOW(),
    time_to TIMESTAMP DEFAULT NOW() + INTERVAL '1 hour',
    status VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE roads (
    id UUID DEFAULT gen_random_uuid(),
    road_name VARCHAR(100) NOT NULL,
    start_point POINT NOT NULL,
    end_point POINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE traffic_conditions (
    id UUID DEFAULT gen_random_uuid(),
    road_id INTEGER REFERENCES roads(road_id) ON DELETE CASCADE,
    traffic_level SMALLINT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE incidents (
    id UUID DEFAULT gen_random_uuid(),
    road_id UUID REFERENCES roads(road_id) ON DELETE SET NULL,
    incident_type VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(20) NOT NULL,
    reported_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS maintances (
    id UUID DEFAULT gen_random_uuid(),
    road_id UUID REFERENCES roads(road_id) ON DELETE CASCADE,
    start_point POINT NOT NULL,
    end_point POINT NOT NULL,
    start_time TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    end_time TIMESTAMP DEFAULT NOW() + INTERVAL '1 day',
    organizer VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    total_area POLYGON NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
) 

