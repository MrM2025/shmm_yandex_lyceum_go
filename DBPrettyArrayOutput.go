package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < 9; i++ {
		if i < 7 {
			fmt.Println(i+1, "я уже сделал:", array[i])
		} else {
			fmt.Println(i+1, "не успел сделать:", array[i])
		} // ? need if? or we it's better to cut all the rubbish?

	}
}
