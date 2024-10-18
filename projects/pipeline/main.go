package main

import (
	"fmt"
	"strconv"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	pre := ""
	for val := range inputStream {
		if pre != "" {
			if val != pre {
				outputStream <- val
			}
			pre = val
		} else {
			pre = val
			outputStream <- pre
		}
	}
}

func main() {
	fmt.Println("/* Эта программа убирает повторения цифр подряд из введённой последовательнсти цифр */")
	inChan := make(chan string)
	outChan := make(chan string)
	go removeDuplicates(inChan, outChan)

	var numQueue string
repeat:
	fmt.Print("Введите последовальность чисел: ")
	fmt.Scan(&numQueue)
	if _, err := strconv.Atoi(numQueue); err != nil {
		fmt.Println("Было введено то, что не является последовательностью чисел. Пожалуйста, повторите ввод!")
		goto repeat
	} else {
		go func() {
			defer close(inChan)
			for _, num := range numQueue {
				inChan <- string(num)
			}
		}()
	}

	fmt.Print("Результат: ")
	for num := range outChan {
		fmt.Print(num)
	}
	fmt.Println()
}
