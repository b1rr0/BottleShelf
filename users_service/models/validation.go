package models

func ValidateUsername(username string) bool {
	// TODO: Define more accurate validness criteria
	// Check special characters e.t.c.
	if len(username) == 0 {
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	// TODO: Define more accurate validness criteria
	if len(password) < 8 {
		// Short password = weak password
		return false
	}
	return true
}

func ValidateOrgname(orgname string) bool {
	// TODO: Define more accurate validness criteria
	// Check special characters e.t.c.
	if len(orgname) == 0 {
		return false
	}
	return true
}
