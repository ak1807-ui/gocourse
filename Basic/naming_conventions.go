package basic

import (
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Pascal Case
	// Eg. CalculationArea, UserInfo, NewHTTPRequest
	// Used In -> Structs, enums, Interfaces

	// Snake Case
	// Eg. user_id, first_name, last_name

	// UPPERCASE
	// Used In -> Constants

	// mixedCase
	// Eg. javaScript, htmlDocument, isValid

	const MAXTRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID : ", employeeID)

}
