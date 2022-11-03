package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	promt := promptui.Select{
		Label: "Select Day",
		Items: []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
	}

	_, ret, err := promt.Run()
	if err != nil {
		fmt.Println("prompt fail, ", err)
	}
	fmt.Println("u chossed, ", ret)
}
