package interactor

import (
	"github.com/AliceDiNunno/yeencloud/src/core/interactor/persistence"
)

type Interactor struct {
	// Log Logger

	Cluster   ClusterAdapter
	Validator Validator
	Localize  Localize
	Trace     Audit
	Mailer    Mailer

	Persistence persistence.Persistence
}
