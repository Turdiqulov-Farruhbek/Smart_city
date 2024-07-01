CREATE TABLE air_quality_stations (
    station_id uuid PRIMARY KEY default gen_random_uuid(),
    location POINT NOT NULL,
    installation_date   TIMESTAMP WITH TIME ZONE
);
CREATE TABLE air_quality_readings (
    reading_id uuid PRIMARY KEY default gen_random_uuid(),
    station_id INTEGER REFERENCES air_quality_stations(station_id),
    pm25_level NUMERIC(5, 2) NOT NULL,
    pm10_level NUMERIC(5, 2) NOT NULL,
    ozone_level NUMERIC(5, 2) NOT NULL,
    time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE noise_monitoring_zones (
    zone_id SERIAL PRIMARY KEY,
    zone_name VARCHAR(100) NOT NULL,
    area_covered POLYGON NOT NULL
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMEST,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMEST,
    deleted_at int default 0 
);
CREATE TABLE noise_level_readings (
    reading_id SERIAL PRIMARY KEY,
    zone_id INTEGER REFERENCES noise_monitoring_zones(zone_id) ON DELETE CASCADE,
    decibel_level NUMERIC(5, 2) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);