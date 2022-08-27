package numberplate

// Repository _
type Repository interface {
	PutCondition(condition PutCondition) error
}
