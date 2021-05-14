package serverh

import (
	"context"
)

// Server is a serverh interface.
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}
