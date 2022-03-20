package main

import (
	"context"
	"fmt"
	lorem "github.com/drhodes/golorem"
	"github.com/go-redis/redis/v8"
	gorma "github.com/hhxsv5/go-redis-memory-analysis"
	"log"
	"os"
	"strconv"
	"time"
)

type analysisClient struct {
	host string
	port uint16
	ctx  context.Context
	rdb  *redis.Client
}

func NewAnalysisClient(host string, port uint16) *analysisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &analysisClient{
		host: host,
		port: port,
		ctx:  context.Background(),
		rdb:  rdb,
	}
}

func genValue(length int) string {
	value := ""
	min := 10
	for {
		min *= 2
		value += lorem.Paragraph(min, min*2)
		if len(value) >= length {
			break
		}
	}
	return value[0:length]
}

func (c *analysisClient) Insert(key string, len, times int) {
	for i := 0; i < times; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		var err error
		for retry := 3; retry > 0; retry-- {
			value := genValue(len)
			//log.Printf("Set %s=%s\n", key, value)
			err = c.rdb.Set(c.ctx, k, value, -1).Err()
			if err == nil {
				break
			}
		}

		if err != nil {
			log.Printf("failed(%s:%d/%d) error: %v\n", key, i, times, err)
			panic(err)
		}

		if i%1000 == 0 {
			time.Sleep(time.Nanosecond * 100)
		}
	}
}

func (c *analysisClient) StartAnalysis() {
	analysis, err := gorma.NewAnalysisConnection(c.host, c.port, "")
	if err != nil {
		log.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	folder := "./reports"
	err = analysis.SaveReports(folder)
	if err == nil {
		log.Println("done")
	} else {
		log.Println("error:", err)
	}

	filename := folder + fmt.Sprintf("/redis-analysis-%s-%d-%d", c.host, c.port, 0)
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	err = os.Rename(filename+".csv", filename+timestamp+".csv")

	if err != nil {
		log.Println("error:", err)
	}
}
