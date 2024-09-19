package service

import "log"

// singleton service

type IdServiceSingleton struct {
	idService *IdService
}

func (s *IdServiceSingleton) GetService() *IdService {
	if s.idService == nil {
		log.Print("NO ID SERVICE AVAILABLE, İNSTANTİATİON")
		s.idService = newIdService()
	}

	return s.idService
}

// ID Service
type IdService struct {
	counter int
}

func newIdService() *IdService {
	return &IdService{
		counter: 0,
	}
}

func (s *IdService) Next() int {
	s.counter++
	return s.counter
}
