package mysqldb

import (
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"gorm.io/gorm"
)

// NumberPlateRepo _
type NumberPlateRepo struct {
	db *gorm.DB
}

// NewNumberPlateRepo will return instance
func NewNumberPlateRepo(db *gorm.DB) *NumberPlateRepo {
	return &NumberPlateRepo{db}
}

// PutCondition 更新條件
func (repo *NumberPlateRepo) PutCondition(condition numberplate.PutCondition) error {
	// TODO: PutCondition
	return nil
}
