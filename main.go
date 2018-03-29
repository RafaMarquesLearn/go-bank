package main

import (
	"fmt"

	"github.com/RafaMarquesLearn/go-bank/models"
)

func createClients() ([5]string, [5]string) {
	var cl1 [5]string
	cl1[0] = "John Doe"
	cl1[1] = "0001"
	cl1[2] = "getin"
	cl1[3] = "5000"
	cl1[4] = "1"

	var cl2 [5]string
	cl2[0] = "Jane Doe"
	cl2[1] = "0002"
	cl2[2] = "iamgod"
	cl2[3] = "12500"
	cl2[4] = "1"

	return cl1, cl2
}
func main() {
	cli1, cli2 := createClients()
	var client1 = models.Client{}
	var client2 = models.Client{}

	client1.AddClient(cli1)
	client2.AddClient(cli2)

	account := "0002"
	password := "getin"

	if client1.CheckClient(account, password) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nope")
	}

}
