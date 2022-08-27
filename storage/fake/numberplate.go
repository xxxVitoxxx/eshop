package fake

import "github.com/xxxVitoxxx/eshop/pkg/numberplate"

// ConditionRepo _
type ConditionRepo struct {
	HowMany, HowLong, Remind int
}

// NewConditionRepo will return instance
func NewConditionRepo() *ConditionRepo {
	return &ConditionRepo{0, 0, 0}
}

// PutCondition 更新條件
func (repo *ConditionRepo) PutCondition(condition numberplate.PutCondition) error {
	repo.HowMany = condition.HowMany
	repo.HowLong = condition.HowLong
	repo.Remind = condition.Remind
	return nil
}
