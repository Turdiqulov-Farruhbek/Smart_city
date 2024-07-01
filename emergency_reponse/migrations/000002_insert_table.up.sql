-- Mock Data for `emergency_incidents`
INSERT INTO emergency_incidents (id, type, location, description, status, reported_at, created_at, updated_at, deleted_at) VALUES
('11111111-1111-1111-1111-111111111111', 'fire', '(40.7128,-74.0060)', 'Fire at building', 'active', NOW(), NOW(), NOW(), 0),
('22222222-2222-2222-2222-222222222222', 'flood', '(34.0522,-118.2437)', 'Flood in downtown', 'reported', NOW(), NOW(), NOW(), 0);

-- Mock Data for `emergency_resources`
INSERT INTO emergency_resources (id, type, current_location, status, created_at, updated_at, deleted_at) VALUES
('33333333-3333-3333-3333-333333333333', 'firetruck', '(40.7128,-74.0059)', 'available', NOW(), NOW(), 0),
('44444444-4444-4444-4444-444444444444', 'ambulance', '(34.0522,-118.2436)', 'dispatched', NOW(), NOW(), 0);

-- Mock Data for `resource_dispatches`
INSERT INTO resource_dispatches (id, incident_id, resource_id, dispatched_at, arrived_at, created_at, updated_at, deleted_at) VALUES
('55555555-5555-5555-5555-555555555555', '11111111-1111-1111-1111-111111111111', '33333333-3333-3333-3333-333333333333', NOW(), NULL, NOW(), NOW(), 0),
('66666666-6666-6666-6666-666666666666', '22222222-2222-2222-2222-222222222222', '44444444-4444-4444-4444-444444444444', NOW(), NOW(), NOW(), NOW(), 0);

-- Mock Data for `emergency_alerts`
INSERT INTO emergency_alerts (id, type, message, affected_area, issued_at, created_at, updated_at, deleted_at) VALUES
('77777777-7777-7777-7777-777777777777', 'weather', 'Severe thunderstorm warning', '(40.7128,-74.0061)', NOW(), NOW(), NOW(), 0),
('88888888-8888-8888-8888-888888888888', 'public health', 'COVID-19 outbreak', '(34.0522,-118.2438)', NOW(), NOW(), NOW(), 0);

-- Mock Data for `evacuation_routes`
INSERT INTO evacuation_routes (id, start_point, end_point, description, created_at, updated_at, deleted_at) VALUES
('99999999-9999-9999-9999-999999999999', '(40.7127,-74.0058)', '(40.7129,-74.0062)', 'Main street to Riverside park', NOW(), NOW(), 0),
('10101010-1010-1010-1010-101010101010', '(34.0521,-118.2435)', '(34.0523,-118.2439)', 'Downtown to Hill park', NOW(), NOW(), 0);

-- Mock Data for `emergency_drills`
INSERT INTO emergency_drills (id, type, location, scheduled_at, completed_at, created_at, updated_at, deleted_at) VALUES
('11111111-2222-3333-4444-555555555555', 'fire drill', '(40.7126,-74.0057)', NOW(), NULL, NOW(), NOW(), 0),
('66666666-7777-8888-9999-000000000000', 'earthquake drill', '(34.0520,-118.2434)', NOW(), NOW(), NOW(), NOW(), 0);

-- Mock Data for `emergency_facilities`
INSERT INTO emergency_facilities (id, type, name, address, capacity, created_at, updated_at, deleted_at) VALUES
('12121212-1212-1212-1212-121212121212', 'shelter', 'Downtown Shelter', '123 Main St, New York, NY', 100, NOW(), NOW(), 0),
('13131313-1313-1313-1313-131313131313', 'hospital', 'Central Hospital', '456 Elm St, Los Angeles, CA', 200, NOW(), NOW(), 0);
