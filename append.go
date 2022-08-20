package main

import "fmt"

func main() {
	emails := []string{"claire123@test.com", "claire456@test.com"}
	fmt.Println("emails before append: ", emails)

	// append a slice to another with a slice variable
	newEmails := []string{"claire789@test.com", "claire10@test.com"}
	emails = append(emails, newEmails...)
	fmt.Println("emails after append newEmails: ", emails)

	// append a slice to another with a slice
	emails = append(emails, []string{"claireExtraEmail1@test.com", "claireExtraEmail1@test.com"}...)
	fmt.Println("emails after append extra emails: ", emails)
}

/* output
emails before append:  [claire123@test.com claire456@test.com]
emails after append newEmails:  [claire123@test.com claire456@test.com claire789@test.com claire10@test.com]
emails after append extra emails:  [claire123@test.com claire456@test.com claire789@test.com claire10@test.com claireExtraEmail1@test.com claireExtraEmail1@test.com]
*/
