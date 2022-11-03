package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

func main() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid numbers")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	ret, err := prompt.Run()
	if err != nil {
		fmt.Println("Promt Run fail, ", err)
		return
	}

	fmt.Println("You typed ", ret)
}
