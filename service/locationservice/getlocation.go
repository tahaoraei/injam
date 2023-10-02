package locationservice

func (s Service) GetLocationInMemory(userID string) (string, error) {
	msg, err := s.inMemoryRepo.GetUserLocation(userID)
	if err != nil {

		return "", err
	}
	return msg, nil
}
