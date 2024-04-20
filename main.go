package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	fullName string
	age      int
}

func (u User) saveToFile() {
	data := []byte(fmt.Sprintf("Name: %s\nAge: %d", u.fullName, u.age))

	// Create the directory if doesn't exist
	errFolder := os.MkdirAll("./users", 0755)
	if errFolder != nil {
		fmt.Println("Unable to create the directory!")
		panic(errFolder)
	}

	err := os.WriteFile("./users/"+"user_data"+".txt", data, 0755)
	if err != nil {
		fmt.Println("Unable to save user to file!")
		panic(err)
	}
}

func getUserInput(inputName string) string {
	println("Enter " + inputName + ": ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	return strings.TrimSpace(value)

}

func (u *User) updateAge() {
	userInput := getUserInput("Your Age")
	age, err := strconv.Atoi(userInput)

	if err != nil {
		return
	}
	u.age = age
}
func (u *User) updateFullName() {
	userInput := getUserInput("Your full name")

	u.fullName = userInput
}

func main() {

	newUser := User{}

	newUser.updateAge()
	newUser.updateFullName()
	newUser.saveToFile()
}
