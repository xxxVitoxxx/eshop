package fake

import (
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
)

// ConditionRepo _
type ConditionRepo struct {
	Condition []numberplate.Condition
}

// NewConditionRepo will return instance
func NewConditionRepo() *ConditionRepo {
	return &ConditionRepo{}
}

// PutConditionByStoreName 更新條件
func (repo *ConditionRepo) PutConditionByStoreName(storeName string, condition numberplate.PutCondition) error {
	for i := range repo.Condition {
		if repo.Condition[i].StoreName == storeName {
			repo.Condition[i].HowMany = condition.HowMany
			repo.Condition[i].HowLong = condition.HowLong
			repo.Condition[i].Remind = condition.Remind
		}
	}
	return nil
}

// CreateCondition _
func (repo *ConditionRepo) CreateCondition(condition numberplate.Condition) error {
	repo.Condition = append(repo.Condition, condition)
	return nil
}

// DeleteConditionByStoreName _
func (repo *ConditionRepo) DeleteConditionByStoreName(storeName string) error {
	for i := 0; i < len(repo.Condition); i++ {
		if repo.Condition[i].StoreName == storeName {
			repo.Condition = append(repo.Condition[:i], repo.Condition[i+1:]...)
		}
	}
	return nil
}
