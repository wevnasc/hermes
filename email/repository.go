package email

type repo interface {
	save(schedule *schedule)
}

type EmailMemoRepo struct {
	schedules []schedule
}

func NewEmailMemoRepo() *EmailMemoRepo {
	return &EmailMemoRepo{schedules: make([]schedule, 100)}
}

func (r *EmailMemoRepo) save(s *schedule) {
	r.schedules = append(r.schedules, *s)
}
