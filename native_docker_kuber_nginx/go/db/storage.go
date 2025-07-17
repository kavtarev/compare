package db

type Storage struct{}

func (s *Storage) GetMe() string {
	return "me"
}
