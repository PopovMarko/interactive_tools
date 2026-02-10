package pomodoro

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Category constats
const (
	CategoryWork       = "work"
	CategoryLongBreak  = "long_break"
	CategoryShortBreak = "short_break"
)

// State constats
const (
	StateNotStarted = iota
	StateRunning
	StatePaused
	StateCompleted
	StateCancelled
)

// Interval type
type Interval struct {
	Id              int64
	Category        string
	State           int
	StartTime       time.Time
	PlannedDuration time.Duration
	ActualDuration  time.Duration
}

// Repository interface
type Repository interface {
	Create(i Interval) int64
	Update(i Interval) error
	GetById(id int64) (Interval, error)
	Last() (Interval, error)
	Breaks(n int) ([]Interval, error)
}

// Custom errors
var (
	ErrNoIntervals        = errors.New("no interval found")
	ErrIntervalNotRunning = errors.New("interval is not running")
	ErrIntervalCompleted  = errors.New("interval is already completed")
	ErrInvaldState        = errors.New("invalid interval state")
	ErrInvalidId          = errors.New("invalid interval Id")
	ErrInvalidCategory    = errors.New("invalid category")
)

// Interval configuration
type IntervalConfig struct {
	Repo               Repository
	WorkDuration       time.Duration
	LongBreakDuration  time.Duration
	ShortBreakDuration time.Duration
}

// interval configuration constructor
func NewConfig(
	repo Repository,
	workDuration,
	longBreakDuration,
	ShortBreakDuration time.Duration) *IntervalConfig {
	c := &IntervalConfig{
		Repo:               repo,
		WorkDuration:       25 * time.Minute,
		LongBreakDuration:  15 * time.Minute,
		ShortBreakDuration: 5 * time.Minute,
	}
	if workDuration > 0 {
		c.WorkDuration = workDuration
	}
	if longBreakDuration > 0 {
		c.LongBreakDuration = longBreakDuration
	}
	if ShortBreakDuration > 0 {
		c.ShortBreakDuration = ShortBreakDuration
	}
	return c
}

// Interval constructor
func NewInterval(cfg *IntervalConfig) (Interval, error) {
	i := Interval{}
	c, err := getNextCategory(cfg)
	if err == nil {
		i.Category = c
	}
	if err != nil && !errors.Is(err, ErrNoIntervals) {
		return i, err
	}
	if err != nil && errors.Is(err, ErrNoIntervals) {
		i.Category = CategoryWork
	}

	switch c {
	case CategoryWork:
		i.PlannedDuration = cfg.WorkDuration
	case CategoryLongBreak:
		i.PlannedDuration = cfg.LongBreakDuration
	case CategoryShortBreak:
		i.PlannedDuration = cfg.ShortBreakDuration
	default:
		return i, fmt.Errorf("invalid category: %s", c)
	}
	return i, nil
}

// Next Category
func getNextCategory(cfg *IntervalConfig) (string, error) {
	last, err := GetLast(cfg)
	if err != nil {
		if errors.Is(err, ErrNoIntervals) {
			return CategoryWork, nil
		}
		return "", err
	}
	switch last.Category {

	case CategoryShortBreak, CategoryLongBreak:
		return CategoryWork, nil

	case CategoryWork:
		breaks, err := cfg.Repo.Breaks(3)
		if err != nil {
			return "", err
		}
		if len(breaks) < 3 {
			return CategoryShortBreak, nil
		}
		for _, b := range breaks {
			if b.Category == CategoryLongBreak {
				return CategoryShortBreak, nil
			}
		}
		return CategoryLongBreak, nil
	}
	return "", ErrInvalidCategory
}

// Callback type
type Callback func(Interval)

// Ticker
func tick(ctx context.Context, id int64, cfg *IntervalConfig, start, periodic, end Callback) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	i, err := cfg.Repo.GetById(id)
	if err != nil {
		return err
	}
	expire := time.After(i.PlannedDuration - i.ActualDuration)
	start(i)
	for {
		select {
		case <-ticker.C:
			i, err = cfg.Repo.GetById(id)
			if err != nil {
				return err
			}
			if i.State == StatePaused {
				return nil
			}
			i.ActualDuration += time.Second
			periodic(i)
		case <-expire:
			i, err = cfg.Repo.GetById(id)
			if err != nil {
				return err
			}
			i.State = StateCompleted
			end(i)
			return cfg.Repo.Update(i)
		case <-ctx.Done():
			i, err = cfg.Repo.GetById(id)
			if err != nil {
				return err
			}
			i.State = StateCancelled
			return cfg.Repo.Update(i)
		}
	}
}

// Get last interval or new one
func GetLast(cfg *IntervalConfig) (Interval, error) {
	var (
		i   Interval
		err error
	)
	i, err = cfg.Repo.Last()
	if err != nil && !errors.Is(err, ErrNoIntervals) {
		return i, err
	}
	if err == nil && i.State != StateCompleted && i.State != StateCancelled {
		return i, nil
	}
	return i, err
}

// Start interval method
func (i *Interval) Start(ctx context.Context, id int64, cfg *IntervalConfig, start, periodic, end Callback) error {
	switch i.State {
	case StateRunning:
		return nil
	case StateNotStarted:
		i.StartTime = time.Now()
		fallthrough
	case StatePaused:
		i.State = StateRunning
		if err := cfg.Repo.Update(*i); err != nil {
			return err
		}
		return tick(ctx, id, cfg, start, periodic, end)
	case StateCancelled, StateCompleted:
		return ErrIntervalCompleted
	default:
		return ErrInvaldState
	}
}

// Pause interval method
func (i *Interval) Pause(cfg *IntervalConfig) error {
	i.State = StatePaused
	return cfg.Repo.Update(*i)

}
