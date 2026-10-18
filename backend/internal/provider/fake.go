package provider

import (
	"context"
	"time"

	"github.com/miqdad-v/livescores/backend/internal/model"
)

type Fake struct {
	start time.Time
}

func NewFake() *Fake {
	return &Fake{start: time.Now()}
}

func (f *Fake) LiveGames(ctx context.Context, league string) ([]model.Game, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	elapsed := int(time.Since(f.start).Seconds())

	return []model.Game{
		{
			ID:        "fake-1",
			League:    league,
			Home:      model.Team{ID: "tor", Name: "Toronto Raptors", Abbreviation: "TOR"},
			Away:      model.Team{ID: "bos", Name: "Boston Celtics", Abbreviation: "BOS"},
			HomeScore: (elapsed / 10) * 2, // +2 every 10 seconds
			AwayScore: (elapsed / 15) * 3, // +3 every 15 seconds
			Status:    model.StatusLive,
			Clock:     "Q2 5:30",
			StartTime: f.start,
		},
		{
			ID:        "fake-2",
			League:    league,
			Home:      model.Team{ID: "lal", Name: "Los Angeles Lakers", Abbreviation: "LAL"},
			Away:      model.Team{ID: "gsw", Name: "Golden State Warriors", Abbreviation: "GSW"},
			HomeScore: 112,
			AwayScore: 108,
			Status:    model.StatusFinal,
			Clock:     "Final",
			StartTime: f.start.Add(-3 * time.Hour),
		},
	}, nil
}
