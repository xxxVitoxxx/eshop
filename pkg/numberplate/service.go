package numberplate

// Service _
type Service struct {
	cr ConditionRepository
}

// NewService will return instance
func NewService(cr ConditionRepository) *Service {
	return &Service{cr}
}

// PutConditionByStoreName 更新條件
func (s *Service) PutConditionByStoreName(storeName string, condition PutCondition) error {
	return s.cr.PutConditionByStoreName(storeName, condition)
}

// CreateCondition _
func (s *Service) CreateCondition(condition Condition) error {
	return s.cr.CreateCondition(condition)
}
