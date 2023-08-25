package app

import (
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/entity"
)

// SetMigrationTable is used to register entity model which want to be migrated
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&entity.Author{},
	}

	return migrationData
}
