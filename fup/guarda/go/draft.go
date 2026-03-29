package main

import "fmt"

func main() {
	var wifi, login, admin bool
	fmt.Scan(&wifi, &login, &admin)

	if wifi == false {
		fmt.Println("you must connect to wifi")
	} else if login == false {
		fmt.Println("you need to login first")
	} else if admin == false {
		fmt.Println("you must to login as admin")
	} else {
		fmt.Println("done")
	}

}
