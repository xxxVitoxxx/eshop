package fake

import "github.com/xxxVitoxxx/eshop/pkg/numberplate"

// NumberPlateRepo _
type NumberPlateRepo struct {
	HowMany, HowLong, Remind int
}

// NewNumberPlateRepo will return instance
func NewNumberPlateRepo() *NumberPlateRepo {
	return &NumberPlateRepo{0, 0, 0}
}

// PutCondition 更新條件
func (repo *NumberPlateRepo) PutCondition(condition numberplate.PutCondition) error {
	repo.HowMany = condition.HowMany
	repo.HowLong = condition.HowLong
	repo.Remind = condition.Remind
	return nil
}
