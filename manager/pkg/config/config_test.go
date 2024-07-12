package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Setup
	originalCfgPath := *cfgPath
	defer func() {
		*cfgPath = originalCfgPath // Restore original path after test
	}()

	// Test cases
	tests := []struct {
		name           string
		configFilePath string
		envVars        map[string]string
		expectError    bool
		expectedConfig Schema
	}{
		{
			name:           "Valid configuration",
			configFilePath: "../../config/config.yaml",
			envVars: map[string]string{
				"http_port": "8080",
			},
			expectError: false,
			expectedConfig: Schema{
				HttpPort:                8080,
				Environment:             "production",
				OtlpEnabled:             false,
				LogLevel:                "debug",
				OtlpEndpoint:            "http://localhost:4318/v1/traces",
				AppVersion:              "",
				ReadTimeout:             10,
				WriteTimeout:            10,
				ServerName:              "",
				RedisURI:                "localhost:6379",
				RedisPassword:           "",
				RedisDB:                 0,
				RedisStreamGlobalPrefix: "tyrscale:stream",
				RedisDBGlobalPrefix:     "tyrscale:db",
				GatewayUrl:              "localhost:7778",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set environment variables
			for key, value := range tc.envVars {
				os.Setenv(key, value)
				defer os.Unsetenv(key) // Clean up environment variables
			}

			// Set configuration path
			*cfgPath = tc.configFilePath

			// Call LoadConfig
			config, err := LoadConfig()

			// Assert results
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedConfig, *config)
			}
		})
	}
}

// test GetConfig
func TestGetConfig(t *testing.T) {
	*cfgPath = "../../config/config.yaml"
	config, err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, config, GetConfig())
}
