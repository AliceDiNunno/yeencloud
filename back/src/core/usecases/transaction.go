package usecases

import (
	"time"
)

type TransactionRequest interface {
	StartRequest() Usecases
	EndRequest(success bool) error
}

func (self UCs) StartRequest() Usecases {
	usecases := NewUsecases(self.i.Cluster, self.i.Localize, self.i.Validator, self.i.Trace, self.i.Persistence.Begin())

	// This helps to prevent the transaction from being open forever and then hanging the database.
	// TODO: Change this to an environment variable.
	usecases.requestTimer = time.NewTimer(2 * time.Second)

	go func() {
		<-usecases.requestTimer.C
		// self.i.Log.Log(domain.LogLevelError).Msg("Transaction request timed out. Forcing rollback.")
		usecases.EndRequest(false)
	}()

	return usecases
}

func (self UCs) EndRequest(success bool) error {
	if success {
		err := self.i.Persistence.Commit()
		if err != nil {
			// self.i.Log.Log(domain.LogLevelError).Msg("Error committing transaction")
			return err
		}
	} else {
		err := self.i.Persistence.Rollback()
		if err != nil {
			// self.i.Log.Log(domain.LogLevelError).Msg("Error rolling back transaction")
			return err
		}
	}

	return nil
}
