package seeders

import (
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/external/faker"
	"github.com/akhidnukhlis/implement-gRpc-server-author-service/internal/entity"
	"github.com/jinzhu/gorm"
)

// SeedAuthor is seeder for testing database
func SeedAuthor(db *gorm.DB) (*entity.Author, error) {
	author := faker.FakeAuthor()
	err := db.Create(&author).Error
	if err != nil {
		return nil, err
	}

	return author, nil
}
