package repository

import (
	"sync"

	"pragprog.com/rggo/interactive_tools/pomo/pomodoro"
)

// in memory repository implementaton
type InMemoryRepository struct {
	mu        sync.RWMutex
	Intervals []pomodoro.Interval
}

// reporitory constructor
func New() *InMemoryRepository {
	return &InMemoryRepository{
		Intervals: make([]pomodoro.Interval, 0),
	}
}

func (r *InMemoryRepository) Create(i pomodoro.Interval) int64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := int64(len(r.Intervals) + 1)
	i.Id = id
	r.Intervals = append(r.Intervals, i)
	return id
}

func (r *InMemoryRepository) Update(i pomodoro.Interval) error {
	id := i.Id
	if id <= 0 {
		return pomodoro.ErrInvalidId
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if id > int64(len(r.Intervals)) {
		return pomodoro.ErrInvalidId
	}

	r.Intervals[id-1] = i
	return nil
}

func (r *InMemoryRepository) GetById(id int64) (pomodoro.Interval, error) {
	if id <= 0 {
		return pomodoro.Interval{}, pomodoro.ErrInvalidId
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	if id > int64(len(r.Intervals)) {
		return pomodoro.Interval{}, pomodoro.ErrInvalidId
	}
	i := r.Intervals[id-1]
	return i, nil
}

func (r *InMemoryRepository) Last() (pomodoro.Interval, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if len(r.Intervals) == 0 {
		return pomodoro.Interval{}, pomodoro.ErrNoIntervals
	}
	return r.Intervals[len(r.Intervals)-1], nil
}

func (r *InMemoryRepository) Breaks(n int) ([]pomodoro.Interval, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	breaks := make([]pomodoro.Interval, 0)
	for i := len(r.Intervals) - 1; i >= 0; i-- {
		if r.Intervals[i].Category == pomodoro.CategoryLongBreak || r.Intervals[i].Category == pomodoro.CategoryShortBreak {
			breaks = append(breaks, r.Intervals[i])
			if len(breaks) == n {
				return breaks, nil
			}
		}
	}
	return breaks, nil
}
