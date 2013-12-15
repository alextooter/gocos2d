package gocos2d

type Scheduler interface {
	Schedule() error
	Unschedule() error
}
type scheduler struct {
}

func NewScheduler() Scheduler {
	return new(scheduler)
}

func (s *scheduler) Schedule() error {
	return nil
}

func (s *scheduler) Unschedule() error {
	return nil
}
