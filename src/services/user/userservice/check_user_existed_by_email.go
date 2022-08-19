package userservice

import "context"

func (s *service) CheckUserExistedByEmail(ctx context.Context, email string) bool {
	_, err := s.repository.FindByEmail(ctx, email)
	if err != nil {
		return false
	}

	return true
}
