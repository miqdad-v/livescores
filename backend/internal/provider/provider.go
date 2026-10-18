package provider

import (
	"context"

	"github.com/miqdad-v/livescores/backend/internal/model"
)

type Provider interface {
	LiveGames(ctx context.Context, leauge string) ([]model.Game, error)
}
