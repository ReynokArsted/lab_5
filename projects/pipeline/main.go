package main

import (
	"fmt"
	"sync"
	"unicode"
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

	var numQueue string
	errorKey := true
	for errorKey == true {
		fmt.Print("Введите последовальность чисел: ")
		fmt.Scan(&numQueue)
		for i := 0; i < len(numQueue); i++ {
			if unicode.IsDigit(rune(numQueue[i])) == true {
				errorKey = false
			} else {
				errorKey = true
			}
		}
		if errorKey == true {
			fmt.Println("Было введено то, что не является последовательностью чисел. Пожалуйста, повторите ввод!")
		}
	}

	inChan := make(chan string, len(numQueue))
	outChan := make(chan string, len(numQueue))

	syncerFirst := new(sync.WaitGroup)
	syncerFirst.Add(1)
	go func() {
		defer close(inChan)
		defer syncerFirst.Done()
		for _, num := range numQueue {
			inChan <- string(num)
		}
	}()
	syncerFirst.Wait()

	go removeDuplicates(inChan, outChan)

	fmt.Print("Результат: ")
	for num := range outChan {
		fmt.Print(num)
	}
	fmt.Println()
}
