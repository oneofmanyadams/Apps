package main

// This works, but is ugly and requires numbering starting at 1 (0's are ignored)
func NumberToLetter(number int) (lettervalue string) {
	if number <= 0 {
		return
	}

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