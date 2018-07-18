package params

import "fmt"
import "regexp"

var (
	customValidation = make(map[string]func(string) bool)

	numberRegex  = regexp.MustCompile(`^\d+$`)
	visaRegex    = regexp.MustCompile(`^\d{4}-\d{4}-\d{4}-\d{4}$`)
	emailRegex   = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	zipCodeRegex = regexp.MustCompile(`^\d{5}(-\d{4})?$`)
)

// AddCustomValidation registers a tag option an associated validator function
func AddCustomValidation(tagName string, isValid func(string) bool) {
	customValidation[tagName] = isValid
}

func valueIsValid(value string, opts tagOptions) error {
	for _, tag := range opts.Options() {
		if isValid, ok := customValidation[tag]; ok {
			if !isValid(value) {
				return fmt.Errorf("value is not a valid %s", tag)
			}
		}
	}
	return nil
}

func validNumber(s string) bool {
	return numberRegex.MatchString(s)
}

func validVisaNumber(s string) bool {
	return visaRegex.MatchString(s)
}

func validEmailAddress(s string) bool {
	return emailRegex.MatchString(s)
}

func validZipCode(s string) bool {
	return zipCodeRegex.MatchString(s)
}
