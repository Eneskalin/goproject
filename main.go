package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func init() {
	log.SetPrefix("INFO:")
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	logDosyam, _ := os.Create("log.txt")

	log.SetOutput(logDosyam)
}
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
			log.Println("Admin")

			if user == userinfo && Password == Passwordinfo {
				fmt.Println("Login successfully")
				log.Println("Login successfully")
				fmt.Println("0-Viewing Logs")
				fmt.Println("1-Log out")
				break

			} else {
				fmt.Println("Wrong Username or Password try again.")
				fmt.Printf("remaining trial right: ")
				fmt.Println(counter)

			}
			log.Println("unsuccessful")

		}

	} else {
		fmt.Println("Welcome Student")
		log.Println("Student")
		for counter := 4; counter > 0; counter-- {

			fmt.Println("Enter Your User Name")
			fmt.Scan(&user)
			fmt.Println("Enter Your Password")
			fmt.Scan(&Password)

			if user == userinfo && Password == Passwordinfo {
				fmt.Println("Login successfully")
				log.Println("Login successfully")
				break

			} else {
				fmt.Println("Wrong Username or Password try again.")
				fmt.Printf("remaining trial right: ")
				fmt.Println(counter)

			}
			log.Println("unsuccessful")

		}

	}

}
