package app

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"
)

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&entity.Author{},
	}

	return migrationData
}
