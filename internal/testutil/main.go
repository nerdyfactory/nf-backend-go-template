package testutil

import (
	"os"
)

// InitDb : InitDb
func InitDb() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5434")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_DATABASE", "local")
	os.Setenv("DB_USERNAME", "testuser")
	os.Setenv("SECRET_KEY", "secretkey")
	os.Setenv("ENC_KEY", "TestSecretKeyKey")
	os.Setenv("ENC_IV", "TestSecretKeyIv2")
}
