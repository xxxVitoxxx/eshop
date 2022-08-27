package numberplate

// Service _
type Service struct {
	cr ConditionRepository
}

// NewService will return instance
func NewService(cr ConditionRepository) *Service {
	return &Service{cr}
}

// PutCondition 更新條件
func (s *Service) PutCondition(condition PutCondition) error {
	return s.cr.PutCondition(condition)
}
