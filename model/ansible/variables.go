package ansible

type Variable struct {
	Name        string
	Type        string
	Description string
	Value       interface{}
	ValueType   string
	Check       struct {
		InUse       bool
		CorrectType bool
		Declared    bool
		Required    bool
	}
}
