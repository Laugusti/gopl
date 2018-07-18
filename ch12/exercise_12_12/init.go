package params

func init() {
	AddCustomValidation("number", validNumber)
	AddCustomValidation("visa", validVisaNumber)
	AddCustomValidation("email", validEmailAddress)
	AddCustomValidation("zipcode", validZipCode)
}
