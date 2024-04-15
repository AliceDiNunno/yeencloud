package postgres

import "fmt"

func sqlerr(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("postgresql: %w", err)
}
