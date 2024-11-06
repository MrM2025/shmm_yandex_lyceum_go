package main

import (
	"fmt"
	"os"
	"strconv"
)

func run() error {

	const numbers = "1234567890"

	args := os.Args[1:]
	if args == nil || len(args) < 3 {
		return fmt.Errorf("not enough args")
	}
	for i, _ := range args {
		_, converr := strconv.ParseInt(args[i], 0, 64)
		if converr != nil {
			return fmt.Errorf("args contain non-integer values")
		}
	}

	file, err := os.Create("config.txt")
	if file == nil && err != nil {
		return fmt.Errorf("can't create file")
	}
	defer file.Close()
	file.WriteString(args[0] + "x" + args[1] + " " + args[2] + "%")

	return nil
}

func main() {

	err := run()
	if err != nil {
		os.Exit(1)
	}
}
