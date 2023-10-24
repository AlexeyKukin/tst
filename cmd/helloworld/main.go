package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(`
Alexey Kukin by Golang 2023
Этот исполняемый файл собирает данные с HW серверов указанных в списке. 
Так же можно запустить сканирование с ключем \s`)
	fl := make([]string, 0, 10)
	fl = flag.Args()
	fmt.Println(fl)

}
