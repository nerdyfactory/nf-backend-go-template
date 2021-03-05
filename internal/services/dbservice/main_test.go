package dbservice

import (
	"testing"

	"github.com/nerdyfactory/nf-backend-go-template/internal/testutil"
)

func TestConnect(t *testing.T) {
	testutil.InitDb()

	db, err := Connect()
	if err != nil {
		t.Errorf("Error opening db connection: %v", err)
		return
	}

	err = db.Ping()
	if err != nil {
		t.Errorf("Error connecting to db: %v", err)
		return
	}
}
