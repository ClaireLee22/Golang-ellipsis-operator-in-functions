package main

import "fmt"

// variadic fuction
func userInfo(name string, emails ...string) {
	fmt.Println("userName: ", name, " emails: ", emails)
}

func main() {
	emails := []string{"claire123@test.com", "claire456@test.com"}

	// not pass the variadic argument
	userInfo("claire")

	// pass the variadic argument with a slice variable
	userInfo("claire", emails...)

	// pass the variadic argument with a slice
	userInfo("claire", []string{"claire123@test.com", "claire456@test.com"}...)

}

/* output
userName:  claire  emails:  []
userName:  claire  emails:  [claire123@test.com claire456@test.com]
userName:  claire  emails:  [claire123@test.com claire456@test.com]
*/

// error: can only use ... with final parameter in list
// func userInfo(emails ...string, name string) {
// 	fmt.Printf("emails type: %T\n", emails)
// 	fmt.Println("UserName: ", name, " emails: ", emails)
// }
