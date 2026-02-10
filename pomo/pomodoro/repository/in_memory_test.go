package repository_test

import (
	"testing"

	"pragprog.com/rggo/interactive_tools/pomo/pomodoro"
	"pragprog.com/rggo/interactive_tools/pomo/pomodoro/repository"
)

func getData(t *testing.T) (pomodoro.Interval, *repository.InMemoryRepository) {
	t.Helper()

	i := pomodoro.Interval{
		Category: pomodoro.CategoryWork,
		State:    pomodoro.StateRunning,
	}
	r := repository.New()
	return i, r
}

func TestCreate(t *testing.T) {
	i, r := getData(t)
	id := r.Create(i)
	if id != 1 {
		t.Errorf("Expected id: 1, got: %d", id)
	}
	if len(r.Intervals) != 1 {
		t.Errorf("Expected intervals length: 1, got %d", len(r.Intervals))
	}
}

func TestGetById(t *testing.T) {
	i, r := getData(t)
	id := r.Create(i)
	iSaved, err := r.GetById(id)
	if err != nil {
		t.Fatalf("Failed to get interval by id: %v", err)
	}
	if iSaved.Id != id {
		t.Errorf("Expected id: %d, got: %d", id, iSaved.Id)
	}
}

func TestUpdate(t *testing.T) {
	i, r := getData(t)
	id := r.Create(i)

	iUpdated, err := r.GetById(id)
	if err != nil {
		t.Fatalf("Failed to get interval by id: %v", err)
	}

	iUpdated.Category = pomodoro.CategoryLongBreak

	if err := r.Update(iUpdated); err != nil {
		t.Fatalf("Failed to update interval: %v", err)
	}
	i, err = r.GetById(id)
	if err != nil {
		t.Fatalf("Failed to get inderval by id: %v", err)
	}
	if i.Category != pomodoro.CategoryLongBreak {
		t.Errorf("Expected category: %s, got: %s", pomodoro.CategoryLongBreak, i.Category)
	}
}

func TestLast(t *testing.T) {
	i, r := getData(t)
	id := r.Create(i)

	iLast, err := r.Last()
	if err != nil {
		t.Fatalf("Failed to get last interval: %v", err)
	}
	if iLast.Id != id {
		t.Errorf("Expected id: %d, got %d", i.Id, iLast.Id)
	}
}

func TestBreaks(t *testing.T) {

}
