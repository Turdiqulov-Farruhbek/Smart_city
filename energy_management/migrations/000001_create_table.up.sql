CREATE TABLE buildings (
    building_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    address TEXT NOT NULL,
    building_type VARCHAR(50) NOT NULL,
    total_area string NOT NULL,
    created_at timestamp default now(),
    updated_at timestamp,
    deleted_at bigint default 0
);

CREATE TABLE energy_meters (
    meter_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    building_id UUID REFERENCES buildings(building_id) ON DELETE CASCADE,
    meter_type VARCHAR(50) NOT NULL,
    created_at timestamp default now(),
    updated_at timestamp,
    deleted_at bigint default 0
);

CREATE TABLE meter_readings (
    reading_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    meter_id UUID REFERENCES energy_meters(meter_id) ON DELETE CASCADE,
    reading_value NUMERIC(10, 2) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);