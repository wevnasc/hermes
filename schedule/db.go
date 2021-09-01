package schedule

type memoDB struct {
	schedules []schedule
}

func newMemoDB() *memoDB {
	return &memoDB{schedules: make([]schedule, 100)}
}

func (r *memoDB) save(s *schedule) {
	r.schedules = append(r.schedules, *s)
}
