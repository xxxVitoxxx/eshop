package numberplate

// PutCondition _
type PutCondition struct {
	HowMany int `json:"how_many"`
	HowLong int `json:"how_long"`
	Remind  int `json:"remind"`
}

// Condition _
type Condition struct {
	StoreName string `json:"store_name"`
	HowMany   int    `json:"how_many"`
	HowLong   int    `json:"how_long"`
	Remind    int    `json:"remind"`
}
