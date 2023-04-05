package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

// creating a function from a var
var (
	isValidUsername = regexp.MustCompile(`^[A-Za-z0-9_]+$`).MatchString
)

var (
	isValidTicker = regexp.MustCompile(`^[A-Z.]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d chars", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only letters, digits or underscores")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	_, err := mail.ParseAddress(value)
	if err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}

func ValidatePageID(value int32) error {
	if value < 1 {
		return fmt.Errorf("must be >= 1")
	}
	return nil
}

func ValidatePageSize(value int32) error {
	if value < 1 || value > 10 {
		return fmt.Errorf("val must be between (1 and 10)")
	}
	return nil
}

func ValidateTicker(value string, allow_empty bool) error {
	if value == "" {
		if allow_empty {
			return nil
		}
		return fmt.Errorf("ticker string was empty")
	}
	if !isValidTicker(value) {
		return fmt.Errorf("must contain only uppercase letters or a period")
	}
	return nil
}

func ValidateGuid(value string) error {
	if value == "" {
		return fmt.Errorf("guid was empty")
	}
	return nil
}

func ValidateTitle(value string) error {
	if value == "" {
		return fmt.Errorf("title was empty")
	}
	return nil
}

func ValidateSite(value string) error {
	if value == "" {
		return fmt.Errorf("invalid site url")
	}
	return nil
}

func ValidateHeadline(value string) error {
	if value == "" {
		return fmt.Errorf("headline was empty")
	}
	return nil
}
