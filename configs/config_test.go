package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitEnv(t *testing.T) {
	t.Setenv("DB_PASSWORD", "password")

	err := LoadEnv()
	assert.Nil(t, err)

	assert.Equal(t, "development", Config.Env)
	assert.Equal(t, "0.0.0.0", Config.DBHost)
	assert.Equal(t, 3306, Config.DBPort)
	assert.Equal(t, "mysql", Config.DBDriver)
	assert.Equal(t, "api_database", Config.DBName)
	assert.Equal(t, "app", Config.DBUser)
	assert.Equal(t, "password", Config.DBPassword)
	assert.Equal(t, true, Config.IsDevelopment())
}
