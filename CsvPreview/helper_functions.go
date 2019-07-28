package main

// This works, but is ugly and requires numbering starting at 1 (0's are ignored)
func NumberToLetter(number int) (lettervalue string) {
	// The 0th character in the alphabet string is just a placeholder (.), 
	// so we just want to return empty string for any number 0 or less
	if number <= 0 {
		return
	}

	// I think Golang actually does integer math here by default, but using int() to force this just in case.
	quotent := int(number) / int(26)
	remainder := number % 26
	
	if remainder == 0 {
		quotent = quotent - 1
		remainder = 26
	}

	alphabet := ".ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lettervalue = NumberToLetter(quotent) + string(alphabet[remainder])

	return
}