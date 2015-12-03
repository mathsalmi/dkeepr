package test

import (
	"fmt"
	"log"
)

var (
	orm *Dkeepr
)

func init() {
	orm, err := NewDkeepr("postgres", "postgres://salmi:@localhost/orm?sslmode=disable")
	if err != nil {
		log.Fatalln("Error on getting a new ORM instance. ", err)
	}

	err = orm.Open()
	if err != nil {
		log.Fatalln("Error on connecting to DB", err)
	}
}

func main() {
	fmt.Println("The main function is being used for tests purposes only. It'll be deleted")

	u := new(user)
	u.id = 1
	u.name = "Matheus Salmi"

	orm.Save(u)
}

type user struct {
	id   int
	name string
}
