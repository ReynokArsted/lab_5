package main

import (
	"fmt"
	"strconv"
	"sync"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		select {
		case val := <-firstChan:
			ch <- val * val
		case val := <-secondChan:
			ch <- val * 3
		case <-stopChan:
		}
		close(ch)
	}()
	return ch
}

func main() {
	var chanNumb int
	fmt.Println("В какой chan запишем значение?\n1. firstChan (значение будет возведено в квадрат)\n2. secondChan (значение будет умножено на 3)\n3. stopChan (никакое значение не будет возращено, и функция calculator() завершится)")
repeat:
	fmt.Print("Вберите номер chan'а (без точки): ")
	if _, err := fmt.Scan(&chanNumb); err != nil {
		fmt.Println("Было введено нечисловое значение. Пожалуйста, повторите ввод!")
		goto repeat
	}

	syncer := new(sync.WaitGroup)
	fChan := make(chan int, 1)
	SChan := make(chan int, 1)
	STChan := make(chan struct{}, 1)

	if chanNumb == 1 {
		var n string
	rp1:
		fmt.Print("Введите числовое значение: ")
		fmt.Scan(&n)
		if val, err := strconv.Atoi(n); err != nil {
			fmt.Println("Было введено нечисловое значение. Пожалуйста, повторите ввод!")
			goto rp1
		} else {
			syncer.Add(1)
			go func() {
				defer syncer.Done()
				fChan <- val
			}()
		}
	} else if chanNumb == 2 {
		var n string
	rp2:
		fmt.Print("Введите числовое значение: ")
		fmt.Scan(&n)
		if val, err := strconv.Atoi(n); err != nil {
			fmt.Println("Было введено нечисловое значение. Пожалуйста, повторите ввод!")
			goto rp2
		} else {
			syncer.Add(1)
			go func() {
				defer syncer.Done()
				SChan <- val
			}()
		}
	} else if chanNumb == 3 {
		syncer.Add(1)
		go func() {
			syncer.Done()
			STChan <- struct{}{}
			fmt.Println("Функция calculate() завершила своё выполнение!")
		}()
	}
	syncer.Wait()
	fmt.Printf("Результат: %d\n", <-calculator(fChan, SChan, STChan))
}
