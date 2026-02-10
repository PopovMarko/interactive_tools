package pomodoro_test

import (
	"testing"

	"pragprog.com/rggo/interactive_tools/pomo/pomodoro"
	"pragprog.com/rggo/interactive_tools/pomo/pomodoro/repository"
)

func GetRepository(t *testing.T) (*repository.InMemoryRepository, func()) {
	t.Helper()
	r := repository.New()
	return r, func() {}
}

func TestNewConfig(t *testing.T) {

}
