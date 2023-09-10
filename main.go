package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Id      int
	Name    string
	Address string
	Work    string
	Reason  func() string
}

var friends = []Friend{
	{
		1,
		"John",
		"Jakarta",
		"Backend Developer",
		func() string {
			return "Ini adalah alasan "
		},
	},
	{
		2,
		"Jane",
		"Surabaya",
		"Frontend Developer",
		func() string {
			return "Ini adalah alasan "
		},
	},
	{
		3,
		"Jack",
		"Semarang",
		"Devops Engineer",
		func() string {
			return "Ini adalah alasan "
		},
	},

	{4,
		"Jill",
		"Bandung",
		"Mobile Developer",
		func() string {
			return "Ini adalah alasan "
		},
	},

	{5,
		"James",
		"Yogyakarta",
		"QA Engineer",
		func() string {
			return "Ini adalah alasan "
		},
	},

	{6,
		"Judy",
		"Padang",
		"Project Manager",
		func() string {
			return "Ini adalah alasan "
		},
	},

	{7,
		"Josh",
		"Palembang",
		"Data Scientist",
		func() string {
			return "Ini adalah alasan "
		},
	},

	{8,
		"Jasmine",
		"Makassar",
		"Data Engineer",
		func() string {
			return "Ini adalah alasan "
		},
	},

	{9,
		"Jules",
		"Bali",
		"UI/UX Designer",
		func() string {
			return "Ini adalah alasan "
		},
	},
}

func main() {
	if len(os.Args) < 2 {
		panic("Missing argument: name")
	}

	name := os.Args[1]
	index := -1
	for i, friend := range friends {
		if friend.Name == name {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Sorry, %s is not my friend.\n", name)
	} else {
		fmt.Println("ID : ", friends[index].Id)
		fmt.Println("Name : ", friends[index].Name)
		fmt.Println("Address : ", friends[index].Address)
		fmt.Println("Work : ", friends[index].Work)
		fmt.Println("Reason : ", friends[index].Reason()+friends[index].Name)
	}
}
