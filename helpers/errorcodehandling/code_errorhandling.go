package errorcodehandling

import (
	"os"

	"github.com/akhidnukhlis/implement-gRpc-server-author-service/config/app"
)

type CodeError struct{}

// ParseSQLError parses driver specific error into known errors.
func (*CodeError) ParseSQLError(err error) error {

	// return nil
	driver := os.Getenv("DB_DRIVER")

	switch driver {
	case "mysql":
		return app.ParseMysqlSQLError(err)
	case "postgres":
		return app.ParsePostgresSQLError(err)
	}
	return err
}
