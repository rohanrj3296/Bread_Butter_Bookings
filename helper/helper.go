package helper
import (
	"strings"
	"fmt"
)
//creating a function for user input validation
func UserInputValidation(firstName string,lastName string,email string,userTickets uint,remainingTickets uint)(bool,bool,bool){
	var validName bool = len(firstName)<2 || len(lastName)<2
	var validMail bool = strings.Contains(email,"@")
	var validUsertickets bool = userTickets>0 && userTickets<=remainingTickets
	//printing what invalid info user has given
	if validName{
		fmt.Println("Invalid Name!!")
	}
	if !validMail{
		fmt.Println("Invalid e-mail entered!!")
	}
	if !validUsertickets{
		fmt.Println("Invalid number of tickets entered")
	}
	return validName, validMail, validUsertickets
}
