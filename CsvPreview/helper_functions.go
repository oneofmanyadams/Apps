package main

// Converts a number (starting at 1) to the coresponding column in excel (letters)
// If the number 0 is provided, the character "." will be returned
func NumberToLetter(number int) (lettervalue string) {

	// the 27th column in excel is labeled AA, this section handles what to do with numbers larger than 26
	if number > 26 {
		// integer math in golang will return an integer as well.
		// So this first part will take the quotent and re-run it through
		// 	the NumberToLetter() function to get the first letter(s) of the column.
		first_num := int(number) / int(26)
		lettervalue = NumberToLetter(first_num)

		// reset the value of number to be just the remaineder of number / 26.
		number = number % 26

		// Example:
		// 28 / 26 is 1 with a remainder of 2.
		// The 1 is ran through NumberToLetter() again, which returns A
		// The number 2 is less than 26, so it's letter value is simply B
		// The Excel column of 28 is, therefore, AB
	}
	
	// We use a "." as the 0th character in the alphabet string for simplicity sake
	// 	(so that that the A = 1, B = 2, etc...)
	alphabet := ".ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lettervalue = lettervalue + string(alphabet[number])
	
	return
}