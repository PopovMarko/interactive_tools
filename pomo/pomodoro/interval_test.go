package pomodoro_test

import (
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
