package config

import (
    "os"
)

var (
    AuthServiceURL       = getEnv("AUTH_SERVICE_URL", "http://localhost:2121")
    CitizenServiceURL    = getEnv("CITIZEN_SERVICE_URL", "http://localhost:2123")
    CityPlanningServiceURL = getEnv("CITY_PLANNING_SERVICE_URL", "http://localhost:2125")
    EmergencyResponseServiceURL = getEnv("EMERGENCY_RESPONSE_SERVICE_URL", "http://localhost:2127")
    EnergyManagmentServiceURL = getEnv("EMERGENCY_MANAGMENT_SERVICE_URL", "http://localhost:2129")
    EnvironmentalMonitoringServiceURL = getEnv("ENVIRONMENTAL_MONITORING_SERVICE_URL", "http://localhost:2131")
    TransportationServiceURL = getEnv("TRANSPORTATION_SERVICE", "http://localhost:2133")
)

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
