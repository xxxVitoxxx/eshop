package numberplate

// ConditionRepository _
type ConditionRepository interface {
	PutConditionByStoreName(storeName string, condition PutCondition) error
	CreateCondition(condition Condition) error
}
