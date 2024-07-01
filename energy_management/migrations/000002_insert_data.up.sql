INSERT INTO buildings (building_id, address, building_type, total_area) VALUES
('123e4567-e89b-12d3-a456-426614174000', '123 Main St', 'Residential', 2000.00);
INSERT INTO energy_meters (meter_id, building_id, meter_type) VALUES
('223e4567-e89b-12d3-a456-426614174000', '123e4567-e89b-12d3-a456-426614174000', 'Electric');
INSERT INTO meter_readings (reading_id, meter_id, reading_value) VALUES
('323e4567-e89b-12d3-a456-426614174000', '223e4567-e89b-12d3-a456-426614174000', 150.75);
