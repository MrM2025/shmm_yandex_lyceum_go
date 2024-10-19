package students

import "time"

type Task struct {
	Summary     string
	Description string
	Deadline    time.Time
	Priority    int
}

func (t Task) IsOverdue() bool {
	return t.Deadline.After(time.Now())
}

func (t Task) IsTopPriority() bool {
	if t.Priority > 3 {
		return true
	}
	return false
}
