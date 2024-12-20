package security_token

import "time"

type SecurityTokenLife struct{}

const (
	ActivationDuration          = 48 * time.Hour
	ResetDuration               = 48 * time.Hour
	RefreshDefaultDuration      = 48 * time.Hour
	RefreshDefaultHalfDuration  = 24 * time.Hour
	RefreshRememberDuration     = 30 * 24 * time.Hour
	RefreshRememberHalfDuration = 15 * 24 * time.Hour
)
