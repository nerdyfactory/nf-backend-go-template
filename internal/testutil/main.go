package testutil

import (
	"os"
)

// InitDb : InitDb
func InitDb() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "1521")
	os.Setenv("DB_SID", "XEPDB1")
	os.Setenv("DB_USER", "KEDDBP")
	os.Setenv("DB_PASSWORD", "yourpassword")
	os.Setenv("DB_LIB_PATH", "/opt/oracle")
	os.Setenv("ENC_KEY", "TestSecretKeyKey")
	os.Setenv("ENC_IV", "TestSecretKeyIv2")
}
