package error

type Error interface {
	WrapError(domainCode string) JSONWrapError
	Error() string
}
