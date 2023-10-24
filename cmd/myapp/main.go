package main

import (
	"flag"
	"fmt"
)

var (
	t    *int
	f, r *string
)

func init() {
	fmt.Println(`Alexey Kukin by Golang 2023
Этот исполняемый файл собирает данные с HW серверов указанных в списке. 
Так же можно запустить сканирование с ключем -s`)

	t = flag.Int("t", 10, "Scan timeout")
	f = flag.String("f", "", "Scan list filename")
	r = flag.String("r", "", "Scan range in StartIP-StopIP format")
	fmt.Println(flag.Args())
	flag.Parse()
}

func main() {

	fmt.Println(*t, *f, *r)

}
