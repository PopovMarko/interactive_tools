package cmd

import (
	"pragprog.com/rggo/interactive_tools/pomo/pomodoro"
	"pragprog.com/rggo/interactive_tools/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.New(), error(nil)
}
