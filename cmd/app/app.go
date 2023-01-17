package main

import (
	"bufio"
	"fmt"
	"module34/module346/pkg/regcalc"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите адрес файла с данными: ")
	inputFile, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputFile = strings.TrimSpace(inputFile)
	fmt.Print("Введите адрес файла для записи данных: ")
	outputFile, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	outputFile = strings.TrimSpace(outputFile)
	err = regcalc.RegFile(inputFile, outputFile)
	if err != nil {
		panic(err)
	}
}
