DO $$ 
BEGIN 
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'incidents_status') THEN
        CREATE TYPE incidents_status AS ENUM(
            'active', 
            'reported', 
            'responding',
            'resolved'
        );
    END IF; 
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'resources_status') THEN
        CREATE TYPE resources_status AS ENUM (
            'available', 
            'dispatched', 
            'out_of_service'
        );
    END IF; 
END $$;

CREATE TABLE IF NOT EXISTS emergency_incidents(
    id UUID PRIMARY KEY,
    type VARCHAR(64),
    location POINT,
    description VARCHAR(256),
    status incidents_status,
    reported_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS emergency_resources(
    id UUID PRIMARY KEY,
    type VARCHAR(64),
    current_location POINT,
    status resources_status,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0    
);

CREATE TABLE IF NOT EXISTS resource_dispatches(
    id UUID PRIMARY KEY,
    incident_id UUID REFERENCES emergency_incidents(id),
    resource_id UUID REFERENCES emergency_resources(id),
    dispatched_at TIMESTAMP DEFAULT NOW(),
    arrived_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0      
);

CREATE TABLE IF NOT EXISTS emergency_alerts(
    id UUID PRIMARY KEY,
    type VARCHAR(64),
    message VARCHAR(256),
    affected_area POINT,
    issued_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0      
);

CREATE TABLE IF NOT EXISTS evacuation_routes(
    id UUID PRIMARY KEY,
    start_point POINT,
    end_point POINT,
    description VARCHAR(256),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0 
);

CREATE TABLE IF NOT EXISTS emergency_drills(
    id UUID PRIMARY KEY,
    type VARCHAR(64),
    location POINT,
    scheduled_at TIMESTAMP DEFAULT NOW(),
    completed_at TIMESTAMP DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0 
);

CREATE TABLE IF NOT EXISTS emergency_facilities(
    id UUID PRIMARY KEY,
    type VARCHAR(64),
    name VARCHAR(64),
    address VARCHAR(128),
    capacity BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);