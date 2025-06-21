package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	d := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d) // Вычисляем время до указанной даты
	return int64(dur.Hours() / 24)
}
