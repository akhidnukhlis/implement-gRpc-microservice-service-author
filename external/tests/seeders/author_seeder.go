package tests

import (
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/external/tests/faker"
	"github.com/akhidnukhlis/implement-gRpc-microservice-service-author/internal/entity"
	"github.com/jinzhu/gorm"
)

// SeedUser is seeder for testing database
func SeedUser(db *gorm.DB) (*entity.Author, error) {
	fakeAuthor := tests.FakeAuthor()
	err := db.Create(&fakeAuthor).Error
	if err != nil {
		return nil, err
	}

	return fakeAuthor, nil
}
