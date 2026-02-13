package app

import (
	"context"
	"fmt"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/button"
	"pragprog.com/rggo/interactive_tools/pomo/pomodoro"
)

type buttonSet struct {
	btnStart *button.Button
	btnPause *button.Button
}

func newButtonSet(
	ctx context.Context,
	cfg *pomodoro.IntervalConfig,
	w *widgets,
	redrawCh chan<- bool,
	errorCh chan<- error) (*buttonSet, error) {

	//Callback to Start interval. Start button action
	startLinterval := func() {
		i, err := pomodoro.GetLast(cfg)
		errorCh <- err

		//Callback start to use in tick func
		start := func(i pomodoro.Interval) {
			message := "Take a break"
			if i.Category == pomodoro.CategoryWork {
				message = "Focus on your task"
			}
			w.update([]int{}, i.Category, message, "", redrawCh)
		}

		//Callback end to use in tick func
		end := func(pomodoro.Interval) {
			w.update([]int{}, "", "Nothing running", "", redrawCh)
		}

		//Callback periodic to use in tick func
		periodic := func(pomodoro.Interval) {
			w.update([]int{int(i.ActualDuration), int(i.PlannedDuration)},
				"", "",
				fmt.Sprint(i.PlannedDuration-i.ActualDuration),
				redrawCh)
		}

		// Call Start func of the last interval instance
		errorCh <- i.Start(ctx, i.Id, cfg, start, periodic, end)
	}

	//Callback to pause interval. Pause button action.
	pauseInterval := func() {
		i, err := pomodoro.GetLast(cfg)
		if err != nil {
			errorCh <- err
		}
		if err := i.Pause(cfg); err != nil {
			if err == pomodoro.ErrIntervalNotRunning {
				return
			}
			w.update([]int{}, "", "Paused... Press Start to continue", "", redrawCh)
		}
	}

	//Button Start instance
	btnStart, err := button.New("(s)tart", func() error {
		go startLinterval()
		return nil
	},
		button.GlobalKey('s'),
		button.WidthFor("(p)ause"),
		button.Height(2),
	)
	if err != nil {
		return nil, err
	}

	//Button Pause instance
	btnPause, err := button.New("(p)ause", func() error {
		go pauseInterval()
		return nil
	},
		button.FillColor(cell.ColorNumber(220)),
		button.GlobalKey('p'),
		button.Height(2))
	if err != nil {
		return nil, err
	}

	return &buttonSet{btnStart, btnPause}, nil

}
