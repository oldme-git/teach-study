package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		host string
		name int
	)
	flag.StringVar(&host, "host", "", "地址")
	flag.IntVar(&name, "name", 12, "名称")
	port := flag.String("port", "80", "端口")
	flag.Parse()
	fmt.Printf("host: %s\n", host)
	fmt.Printf("name: %d\n", name)
	fmt.Printf("port: %s\n", *port)
}
