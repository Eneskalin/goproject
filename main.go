package main

import "fmt"

func main() {
	var userinfo int = 2210205010
	var Passwordinfo int = 6868

	fmt.Println("Select user type(admin:0,student:1)")
	var tip int
	var user, Password int
	fmt.Scan(&tip)
	if tip == 0 {
		fmt.Println("Welcome Admin")
		for counter := 4; counter >= 0; counter-- {

			fmt.Println("Enter Your User Name")
			fmt.Scan(&user)
			fmt.Println("Enter Your Password")
			fmt.Scan(&Password)

			if user == userinfo && Password == Passwordinfo {
				fmt.Println("Login successfully")
				fmt.Println("0-Viewing Logs")
				fmt.Println("1-Log out")
				break

			} else {
				fmt.Println("Wrong Username or Password try again.")
				fmt.Printf("remaining trial right: ")
				fmt.Println(counter)

			}

		}

	} else {
		fmt.Println("Welcome Student")
		for counter := 4; counter > 0; counter-- {

			fmt.Println("Enter Your User Name")
			fmt.Scan(&user)
			fmt.Println("Enter Your Password")
			fmt.Scan(&Password)

			if user == userinfo && Password == Passwordinfo {
				fmt.Println("Login successfully")
				break

			} else {
				fmt.Println("Wrong Username or Password try again.")
				fmt.Printf("remaining trial right: ")
				fmt.Println(counter)

			}

		}

	}

}
