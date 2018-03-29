package main

import (
	"fmt"

	"github.com/RafaMarquesLearn/go-bank/models"
)

func main() {
	vl1 := []string{"John Doe", "0001", "getin", "5000", "true"}
	var cl1 models.Client
	cl1 = cl1.AddClient(vl1)

	vl2 := []string{"Jane Doe", "0002", "iamgod", "12500", "true"}
	var cl2 models.Client
	cl2 = cl2.AddClient(vl2)

	vl3 := []string{"Jack", "0003", "broke", "0", "true"}
	var cl3 models.Client
	cl3 = cl3.AddClient(vl3)

	account := 0001
	password := "getin"

	lcli := []models.Client{}

	lcli = append(lcli, cl1, cl2, cl3)

	for _, lc := range lcli {
		if lc.CheckClient(account, password) {
			fmt.Println("OK!")
		}
	}

}
