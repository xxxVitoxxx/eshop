package numberplate

// ConditionRepository _
type ConditionRepository interface {
	PutCondition(condition PutCondition) error
}
