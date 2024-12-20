package security_token

type SecurityTokenType string

const (
	ACTIVATION SecurityTokenType = "ACTIVATION"
	REFRESH    SecurityTokenType = "REFRESH"
	RESET      SecurityTokenType = "RESET"
)

func (t SecurityTokenType) IsValid() bool {
	switch t {
	case ACTIVATION, REFRESH, RESET:
		return true
	default:
		return false
	}
}
