package postik

type Validator interface {
	IsValid(field *Field) bool
	Error() string
}
