package mysqldb

import (
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"gorm.io/gorm"
)

// ConditionRepo _
type ConditionRepo struct {
	db *gorm.DB
}

// NewConditionRepo will return instance
func NewConditionRepo(db *gorm.DB) *ConditionRepo {
	return &ConditionRepo{db}
}

// Condition _
type Condition struct {
	StoreName string `gorm:"size:100; not null"`
	HowMany   int    `gorm:"size:5; null"`
	HowLong   int    `gorm:"size:4; null"`
	Remind    int    `gorm:"size:4; null"`
}

// PutCondition 更新條件
func (repo *ConditionRepo) PutCondition(condition numberplate.PutCondition) error {
	return repo.db.Model(&Condition{}).Where("store_name = ?", "eshop").Updates(map[string]interface{}{
		"how_many": condition.HowMany,
		"how_long": condition.HowLong,
		"remind":   condition.Remind,
	}).Error
}
