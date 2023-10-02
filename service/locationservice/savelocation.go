package locationservice

func (s Service) SaveLocationInMemory(userID string, od string) error {
	if err := s.inMemoryRepo.SaveUserLocation(userID, od); err != nil {
		return err
	}
	return nil
}
