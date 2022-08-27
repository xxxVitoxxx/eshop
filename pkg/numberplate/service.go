package numberplate

// Service _
type Service struct {
	r Repository
}

// NewService will return instance
func NewService(r Repository) *Service {
	return &Service{r}
}

// PutCondition 更新條件
func (s *Service) PutCondition(condition PutCondition) error {
	return s.r.PutCondition(condition)
}
