package domain

type AuditID string
type AuditSaveFunc func([]byte)

func (t AuditID) String() string {
	return string(t)
}
