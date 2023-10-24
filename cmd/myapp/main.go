package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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
	rscanip("192.168.1.1-192.168.2.1")
}

func rscanip(irang string) map[*net.IP]bool {

	rn := strings.Split(irang, "-")
	if len(rn) != 2 {
		log.Fatal("WRONG IP range")
	}

	minmax := make([]net.IP, 0, 2)

	for _, ipAddr := range rn {
		ipv4Addr := net.ParseIP(ipAddr)
		if ipv4Addr == nil {
			log.Fatal("Can NOT parse ipv4 address ")
		}

		fmt.Println(ipv4Addr)

		minmax = append(minmax, net.IP.To4(ipv4Addr))

	}
	v := minmax[0]
	for {
		fmt.Println(v)
		con(v)
		if v.Equal(minmax[1]) == false {
			if v[3] == 255 {
				v[2]++
				v[3] = 0
			} else {
				v[3]++
			}
		} else {
			break
		}
	}

	return map[*net.IP]bool{}

}

func con(ip net.IP) {
	c, er := net.DialTimeout("tcp", ip.String()+":80", time.Second)
	if er != nil {
		log.Println(er)
	} else {
		//		b := make([]byte, 400, 1000)
		//		fmt.Fprintf(c, "GET / HTTP/1.0\r\n\r\n")
		//		n, er := c.Read(b)
		//		fmt.Println(n, er)
		//		fmt.Println(b)

		c.Close()
	}

}
