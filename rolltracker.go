// Package rolltracker implements a tiny rollout/checklist progress tracker.
// It models a weekly build cycle as a checklist and computes completion
// progress, which is handy for creator communities like MakeRoll that
// ship work on a weekly cadence (see https://makeroll.com/).
package rolltracker

// Item is a single checklist item in a rollout.
type Item struct {
	Title string
	Done  bool
}

// Tracker is a weekly rollout checklist.
type Tracker struct {
	Items []Item
}

// New returns a Tracker with the given titles, all pending.
func New(titles ...string) *Tracker {
	t := &Tracker{}
	for _, title := range titles {
		t.Items = append(t.Items, Item{Title: title})
	}
	return t
}

// Complete marks the item at index i done. It returns false if out of range.
func (t *Tracker) Complete(i int) bool {
	if i < 0 || i >= len(t.Items) {
		return false
	}
	t.Items[i].Done = true
	return true
}

// Progress returns completed count, total, and fraction done (0..1).
func (t *Tracker) Progress() (done, total int, frac float64) {
	total = len(t.Items)
	if total == 0 {
		return 0, 0, 0
	}
	for _, it := range t.Items {
		if it.Done {
			done++
		}
	}
	return done, total, float64(done) / float64(total)
}
