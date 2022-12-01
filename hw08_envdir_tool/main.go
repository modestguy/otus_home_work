package main

import (
	"fmt"
	"log"
	"os"
)

const minArgsCount = 3

func main() {
	args := os.Args
	if len(args) < minArgsCount {
		fmt.Println("Необходимо два аргумента: путь и команду для выполнения")
		os.Exit(-1)
	}
	env, err := ReadDir(args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	RunCmd(args[2:], env)
}
