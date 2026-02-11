package pomodoro_test

import (
	"fmt"
	"testing"
	"time"

	"pragprog.com/rggo/interactive_tools/pomo/pomodoro"
	"pragprog.com/rggo/interactive_tools/pomo/pomodoro/repository"
)

func GetRepository(t *testing.T) (*repository.InMemoryRepository, func()) {
	t.Helper()
	r := repository.New()
	return r, func() {}
}

func TestNewConfig(t *testing.T) {
	testCases := []struct {
		name          string
		timeDurations [3]time.Duration
		expected      pomodoro.IntervalConfig
	}{
		{name: "Default durations",
			timeDurations: [3]time.Duration{
				0, 0, 0,
			},
			expected: pomodoro.IntervalConfig{
				WorkDuration:       25 * time.Minute,
				LongBreakDuration:  15 * time.Minute,
				ShortBreakDuration: 5 * time.Minute,
			},
		},
		{name: "One parameter set",
			timeDurations: [3]time.Duration{
				20 * time.Minute,
			},
			expected: pomodoro.IntervalConfig{
				WorkDuration:       20 * time.Minute,
				LongBreakDuration:  15 * time.Minute,
				ShortBreakDuration: 5 * time.Minute,
			},
		},
		{name: "All parameters set",
			timeDurations: [3]time.Duration{
				10 * time.Minute,
				10 * time.Minute,
				10 * time.Minute,
			},
			expected: pomodoro.IntervalConfig{
				WorkDuration:       10 * time.Minute,
				LongBreakDuration:  10 * time.Minute,
				ShortBreakDuration: 10 * time.Minute,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, cleanUp := GetRepository(t)
			defer cleanUp()
			cfg := pomodoro.NewConfig(r,
				tc.timeDurations[0],
				tc.timeDurations[1],
				tc.timeDurations[2])

			if cfg.WorkDuration != tc.expected.WorkDuration {
				t.Errorf("Expected WorkDuration %v, got %v", tc.expected.WorkDuration, cfg.WorkDuration)
			}
			if cfg.LongBreakDuration != tc.expected.LongBreakDuration {
				t.Errorf("Expected LongBreakDuration %v, got %v", tc.expected.LongBreakDuration, cfg.LongBreakDuration)
			}
			if cfg.ShortBreakDuration != tc.expected.ShortBreakDuration {
				t.Errorf("Expected ShortBreakDuration %v, got %v", tc.expected.ShortBreakDuration, cfg.ShortBreakDuration)
			}
		})
	}
}

func TestNewInterval(t *testing.T) {
	const duration = 1 * time.Millisecond
	r, cleanUp := GetRepository(t)
	defer cleanUp()

	cfg := pomodoro.NewConfig(r, 3*duration, 2*duration, duration)

	for j := 1; j <= 16; j++ {
		var (
			expCategory string
			expDuration time.Duration
		)
		switch {
		case j%2 != 0:
			expCategory = pomodoro.CategoryWork
			expDuration = 3 * duration
		case j%8 == 0:
			expCategory = pomodoro.CategoryLongBreak
			expDuration = 2 * duration
		case j%2 == 0:
			expCategory = pomodoro.CategoryShortBreak
			expDuration = duration
		}
		name := fmt.Sprintf("%s_%d", expCategory, j)
		t.Run(name, func(t *testing.T) {
			res, err := pomodoro.NewInterval(cfg)
			if err != nil {
				t.Fatalf("Expect no error. got %v", err)
			}
			cfg.Repo.Create(res)
			if res.Category != expCategory {
				t.Errorf("Expected category %s, got %s", expCategory, res.Category)
			}
			if res.PlannedDuration != expDuration {
				t.Errorf("Expected planed duration %v, got %v", expDuration, res.PlannedDuration)
			}
		})
	}
}
