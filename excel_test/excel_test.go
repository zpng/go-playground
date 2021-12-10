package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

//模拟多个goroutine相关
func main() {
	data := make(chan string)
	go GenData(data)
	go GenExcelData(data)
	wg.Add(2)
	wg.Wait()
}

func GenData(data chan string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		data <- strconv.Itoa(i + 100)
	}
}

func GenExcelData(data chan string) {
	defer wg.Done()
	file := excelize.NewFile()
	row := 0
	for line := range data {
		// set data
		row = row + 1
		file.SetCellStr("Sheet1", ExcelAxis(0, row), line)

	}

	file.DeleteSheet("Sheet1")
	_ = file.Save()
}

func ExcelAxis(col, row int) string {
	var s string
	for {
		if col < 26 {
			return fmt.Sprintf("%c%s%d", col+'A', s, row+1)
		}
		s = fmt.Sprintf("%c%s", col%26+'A', s)
		col = col/26 - 1
	}
}
