package userservice

import "injam/param"

func (s Service) Profile(request param.UserProfileRequest) (param.UserProfileResponse, error) {
	u, err := s.repo.GetUserByID(request.ID)
	if err != nil {
		return param.UserProfileResponse{}, err
	}
	return param.UserProfileResponse{
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}, nil
}
