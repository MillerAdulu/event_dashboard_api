package ally

import (
	"context"
)

// Usecase -
type Usecase interface {
	RegisterAlly(ctx context.Context, ally map[string]interface{})
}
