package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var generatedValue = map[int]string{}

func main() {
	host := flag.String("host", "localhost", "redis server address")
	port := flag.Int("port", 6379, "redis server port")
	bytes := flag.Int("bytes", 10, "set value size")
	times := flag.String("times", "10000,100000,500000", "set value times")

	flag.Parse()

	client := NewAnalysisClient(*host, uint16(*port))

	batch := strings.Split(*times, ",")
	for _, t := range batch {
		i, err := strconv.Atoi(t)
		if err != nil {
			fmt.Printf("recv wrong format times: %s, skip!\n", t)
			continue
		}
		key := fmt.Sprintf("%vbytes_%v", *bytes, i)
		client.Insert(key, *bytes, i)
	}

	client.StartAnalysis()
}
