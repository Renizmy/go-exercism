package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var line = strings.Repeat("*", numStarsPerLine)
	return line + "\n" + welcomeMsg + "\n" + line
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	r := strings.NewReplacer("*", "")
	sanitizedMsg := strings.Trim(r.Replace(oldMsg), " \r\n")
	return sanitizedMsg
}
