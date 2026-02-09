package repository

import (
	"sync"

	"pragprog.com/rggo/interactive_tools/pomo/pomodoro"
)

// in memory repository implementaton
type InMemoryRepository struct {
	mu        sync.RWMutex
	intervals []pomodoro.Interval
}

// reporitory constructor
func New() *InMemoryRepository {
	return &InMemoryRepository{
		intervals: make([]pomodoro.Interval, 0),
	}
}

func (r *InMemoryRepository) Create(i pomodoro.Interval) int64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := int64(len(r.intervals) + 1)
	i.Id = id
	r.intervals = append(r.intervals, i)
	return id
}

func (r *InMemoryRepository) Update(i pomodoro.Interval) error {
	id := i.Id
	if id <= 0 {
		return pomodoro.ErrInvalidId
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if id > int64(len(r.intervals)) {
		return pomodoro.ErrInvalidId
	}

	r.intervals[id-1] = i
	return nil
}

func (r *InMemoryRepository) GetById(id int64) (pomodoro.Interval, error) {
	if id <= 0 {
		return pomodoro.Interval{}, pomodoro.ErrInvalidId
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	if id > int64(len(r.intervals)) {
		return pomodoro.Interval{}, pomodoro.ErrInvalidId
	}
	i := r.intervals[id-1]
	return i, nil
}

func (r *InMemoryRepository) Last() (pomodoro.Interval, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if len(r.intervals) == 0 {
		return pomodoro.Interval{}, pomodoro.ErrNoIntervals
	}
	return r.intervals[len(r.intervals)-1], nil
}

func (r *InMemoryRepository) Breaks(n int) ([]pomodoro.Interval, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	breaks := make([]pomodoro.Interval, 0)
	for i := len(r.intervals) - 1; i >= 0; i-- {
		if r.intervals[i].Category == pomodoro.CategoryLongBreak || r.intervals[i].Category == pomodoro.CategoryShortBreak {
			breaks = append(breaks, r.intervals[i])
			if len(breaks) == n {
				return breaks, nil
			}
		}
	}
	return breaks, nil
}
