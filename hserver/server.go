package hserver

import (
	"context"
)

// Server is a hserver interface.
type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}
