CREATE TABLE city_zones (
    id UUID DEFAULT gen_random_uuid(),
    zone_name VARCHAR(100) NOT NULL,
    area POLYGON NOT NULL,
    current_usage VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE planning_proposals (
    id UUID DEFAULT gen_random_uuid(),
    zone_id INTEGER REFERENCES city_zones(zone_id) ON DELETE SET NULL,
    proposal_type VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(20) NOT NULL,
    submitted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE proposal_feedback (
    id UUID DEFAULT gen_random_uuid(),
    proposal_id INTEGER REFERENCES planning_proposals(proposal_id) ON DELETE CASCADE,
    citizen_id INTEGER NOT NULL,
    feedback_text TEXT NOT NULL,
    submitted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE demographic_data (
    id UUID DEFAULT gen_random_uuid(),
    zone_id INTEGER REFERENCES city_zones(zone_id) ON DELETE CASCADE,
    population INTEGER NOT NULL,
    median_age NUMERIC(4, 1) NOT NULL,
    median_income NUMERIC(10, 2) NOT NULL,
    data_year INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE infrastructure_assets (
    id UUID DEFAULT gen_random_uuid(),
    asset_type VARCHAR(50) NOT NULL,
    location POINT NOT NULL,
    installation_date DATE NOT NULL,
    current_status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE development_scenarios (
    id UUID DEFAULT gen_random_uuid(),
    scenario_name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE simulation_results (
    id UUID DEFAULT gen_random_uuid(),
    scenario_id INTEGER REFERENCES development_scenarios(scenario_id) ON DELETE CASCADE,
    result_data JSONB NOT NULL,
    simulated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE building_permits (
    id UUID DEFAULT gen_random_uuid(),
    zone_id INTEGER REFERENCES city_zones(zone_id) ON DELETE SET NULL,
    permit_type VARCHAR(50) NOT NULL,
    applicant_name VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL,
    applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE land_parcels (
    id UUID DEFAULT gen_random_uuid(),
    zone_id INTEGER REFERENCES city_zones(zone_id) ON DELETE SET NULL,
    area NUMERIC(10, 2) NOT NULL,
    current_usage VARCHAR(50) NOT NULL,
    owner VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);
