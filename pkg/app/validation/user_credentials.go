package validation

func UserCredentials(login, password string) error {
	if len(login) < MinLoginLength || len(login) > MaxLoginLength {
		return ErrIncorrectUsernameLength
	}

	if len(password) < MinPasswordLength || len(password) > MaxPasswordLength {
		return ErrIncorrectPasswordLength
	}
	return nil
}
