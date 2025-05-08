package constant

type VerificationType int

const (
	AppTokenValue          VerificationType = 1
	InternalToolTokenValue VerificationType = 2
)

var VerificationTypeConstants = struct {
	AppToken          VerificationType
	InternalToolToken VerificationType
}{
	AppToken:          AppTokenValue,
	InternalToolToken: InternalToolTokenValue,
}
