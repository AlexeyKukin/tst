package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(`Alexey Kukin by Golang 2023
Этот исполняемый файл собирает данные с HW серверов указанных в списке. 
Так же можно запустить сканирование с ключем -s`)

	flag.Int("t", 10, "Scan timeout")
	flag.String("f", "", "Scan list filename")
	flag.Parse()

	fmt.Println(flag.Parsed())
}
