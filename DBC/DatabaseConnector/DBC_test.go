package DatabaseConnector

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	errPing := Connect()
	if errPing != nil {
		t.Errorf("Could not do Connect()\n%+v", errPing)
	}

}
