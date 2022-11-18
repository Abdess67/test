package user

import (
    "fmt"
	"time"
)

type User struct {
    Name        string
	FirstName   string
	Email       string
	Password    string
	Phone       string
	Date		time.Time
    CGU        	bool
    Premium  	int
}

/*
Affiche des informations sur un personnage

@return: void
*/
func (p User) Affichage() { // déclaration de ma méthode Affichage() liée à ma structure Personnage
    fmt.Println("--------------------------------------------------")
    fmt.Println("Vie du personnage", p.Name, ":", p.FirstName)
    fmt.Println("Puissance du personnage", p.Email, ":", p.Password)

    if p.Mort {
        fmt.Println("Vie du personnage", p.Phone, "est mort")
    } else {
        fmt.Println("Vie du personnage", p.Date, "est vivant")
    }

    fmt.Println("\nLe personnage", p.CGU, "possède dans son inventaire :", p.Premium)

    // for _, item := range p.Inventaire {
    //     fmt.Println("-", item)
    // }
}
