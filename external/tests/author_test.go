package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/helpers/errorcodehandling"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"
	"github.com/jinzhu/gorm"
	"testing"
)

type Suite struct {
	DB *gorm.DB

	mock      sqlmock.Sqlmock
	author    entity.Author
	codeError *errorcodehandling.CodeError
}

func TestCreateAuthor(t *testing.T) {
	//s := &Suite{}
	//var (
	//	db  *sql.DB
	//	err error
	//)
	//db, s.mock, err = sqlmock.New()
	//if err != nil {
	//	t.Errorf("Failed to open mock sql db, got error: %v", err)
	//}
	//if db == nil {
	//	t.Error("mock db is null")
	//}
	//if s.mock == nil {
	//	t.Error("sqlmock is null")
	//}
	//dialector := postgres.New(postgres.Config{
	//	DSN:                  "sqlmock_db_0",
	//	DriverName:           "postgres",
	//	Conn:                 db,
	//	PreferSimpleProtocol: true,
	//})
	//s.db, err = gorm.Open(dialector, &gorm.Config{})
	//if err != nil {
	//	t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	//}
	//if s.db == nil {
	//	t.Error("gorm db is null")
	//}
	//s.author = entity.Author{
	//	ID:        "",
	//	Name:      "",
	//	Email:     "",
	//	Username:  "",
	//	Password:  "",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//}
	//defer db.Close()
	//
	//s.mock.MatchExpectationsInOrder(false)
	//s.mock.ExpectBegin()
	//s.mock.ExpectQuery(regexp.QuoteMeta(
	//	`INSERT INTO "students" ("id","name")
	//                VALUES ($1,$2) RETURNING "students"."id"`)).
	//	WithArgs(s.student.ID, s.student.Name).
	//	WillReturnRows(sqlmock.NewRows([]string{"id"}).
	//		AddRow(s.student.ID))
	//s.mock.ExpectCommit()
	//if err = s.db.Create(&s.student).Error; err != nil {
	//	t.Errorf("Failed to insert to gorm db, got error: %v", err)
	//}
	//err = s.mock.ExpectationsWereMet()
	//if err != nil {
	//	t.Errorf("Failed to meet expectations, got error: %v", err)
	//}
}
