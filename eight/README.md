# Week8

### 1. Using redis benchmark tool, test 10 20 50 100 200 1k 5k bytes value size, redis get set performance

```bash
$ cd example/redis-benchmark
# Set the size you want
$ SIZE=100 docker-compose up
```

* output
```bash
====== SET ======
client_1  |   100000 requests completed in 0.87 seconds
client_1  |   50 parallel clients
client_1  |   100 bytes payload
client_1  |   keep alive: 1
client_1  |   host configuration "save": 3600 1 300 100 60 10000
client_1  |   host configuration "appendonly": no
client_1  |   multi-thread: no
client_1  |
client_1  | Latency by percentile distribution:
client_1  | 0.000% <= 0.087 milliseconds (cumulative count 1)
client_1  | 50.000% <= 0.223 milliseconds (cumulative count 51785)
client_1  | 75.000% <= 0.247 milliseconds (cumulative count 76671)
client_1  | 87.500% <= 0.271 milliseconds (cumulative count 88689)
client_1  | 93.750% <= 0.295 milliseconds (cumulative count 94555)
client_1  | 96.875% <= 0.319 milliseconds (cumulative count 97439)
client_1  | 98.438% <= 0.343 milliseconds (cumulative count 98573)
client_1  | 99.219% <= 0.391 milliseconds (cumulative count 99272)
client_1  | 99.609% <= 0.575 milliseconds (cumulative count 99610)
client_1  | 99.805% <= 1.015 milliseconds (cumulative count 99806)
client_1  | 99.902% <= 3.543 milliseconds (cumulative count 99903)
client_1  | 99.951% <= 6.919 milliseconds (cumulative count 99952)
client_1  | 99.976% <= 7.071 milliseconds (cumulative count 99976)
client_1  | 99.988% <= 7.255 milliseconds (cumulative count 99988)
client_1  | 99.994% <= 7.295 milliseconds (cumulative count 99994)
client_1  | 99.997% <= 7.327 milliseconds (cumulative count 99997)
client_1  | 99.998% <= 7.343 milliseconds (cumulative count 99999)
client_1  | 99.999% <= 7.351 milliseconds (cumulative count 100000)
client_1  | 100.000% <= 7.351 milliseconds (cumulative count 100000)
client_1  |
client_1  | Cumulative distribution of latencies:
client_1  | 0.004% <= 0.103 milliseconds (cumulative count 4)
client_1  | 36.576% <= 0.207 milliseconds (cumulative count 36576)
client_1  | 95.847% <= 0.303 milliseconds (cumulative count 95847)
client_1  | 99.347% <= 0.407 milliseconds (cumulative count 99347)
client_1  | 99.549% <= 0.503 milliseconds (cumulative count 99549)
client_1  | 99.625% <= 0.607 milliseconds (cumulative count 99625)
client_1  | 99.682% <= 0.703 milliseconds (cumulative count 99682)
client_1  | 99.716% <= 0.807 milliseconds (cumulative count 99716)
client_1  | 99.768% <= 0.903 milliseconds (cumulative count 99768)
client_1  | 99.804% <= 1.007 milliseconds (cumulative count 99804)
client_1  | 99.819% <= 1.103 milliseconds (cumulative count 99819)
client_1  | 99.828% <= 1.207 milliseconds (cumulative count 99828)
client_1  | 99.837% <= 1.303 milliseconds (cumulative count 99837)
client_1  | 99.843% <= 1.407 milliseconds (cumulative count 99843)
client_1  | 99.847% <= 1.503 milliseconds (cumulative count 99847)
client_1  | 99.850% <= 1.607 milliseconds (cumulative count 99850)
client_1  | 99.851% <= 2.103 milliseconds (cumulative count 99851)
client_1  | 99.896% <= 3.103 milliseconds (cumulative count 99896)
client_1  | 99.930% <= 4.103 milliseconds (cumulative count 99930)
client_1  | 99.950% <= 5.103 milliseconds (cumulative count 99950)
client_1  | 99.977% <= 7.103 milliseconds (cumulative count 99977)
client_1  | 100.000% <= 8.103 milliseconds (cumulative count 100000)
client_1  |
client_1  | Summary:
client_1  |   throughput summary: 114416.48 requests per second
client_1  |   latency summary (msec):
client_1  |           avg       min       p50       p95       p99       max
client_1  |         0.232     0.080     0.223     0.303     0.367     7.351
====== GET ======                                                     l: 0.267)
client_1  |   100000 requests completed in 0.87 seconds
client_1  |   50 parallel clients
client_1  |   100 bytes payload
client_1  |   keep alive: 1
client_1  |   host configuration "save": 3600 1 300 100 60 10000
client_1  |   host configuration "appendonly": no
client_1  |   multi-thread: no
client_1  |
client_1  | Latency by percentile distribution:
client_1  | 0.000% <= 0.071 milliseconds (cumulative count 1)
client_1  | 50.000% <= 0.223 milliseconds (cumulative count 50571)
client_1  | 75.000% <= 0.255 milliseconds (cumulative count 80600)
client_1  | 87.500% <= 0.271 milliseconds (cumulative count 87802)
client_1  | 93.750% <= 0.295 milliseconds (cumulative count 93991)
client_1  | 96.875% <= 0.319 milliseconds (cumulative count 97670)
client_1  | 98.438% <= 0.335 milliseconds (cumulative count 98756)
client_1  | 99.219% <= 0.351 milliseconds (cumulative count 99242)
client_1  | 99.609% <= 0.383 milliseconds (cumulative count 99683)
client_1  | 99.805% <= 0.407 milliseconds (cumulative count 99825)
client_1  | 99.902% <= 0.455 milliseconds (cumulative count 99904)
client_1  | 99.951% <= 0.559 milliseconds (cumulative count 99955)
client_1  | 99.976% <= 0.623 milliseconds (cumulative count 99976)
client_1  | 99.988% <= 0.687 milliseconds (cumulative count 99988)
client_1  | 99.994% <= 0.735 milliseconds (cumulative count 99995)
client_1  | 99.997% <= 0.767 milliseconds (cumulative count 99997)
client_1  | 99.998% <= 0.775 milliseconds (cumulative count 99999)
client_1  | 99.999% <= 0.863 milliseconds (cumulative count 100000)
client_1  | 100.000% <= 0.863 milliseconds (cumulative count 100000)
client_1  |
client_1  | Cumulative distribution of latencies:
client_1  | 0.003% <= 0.103 milliseconds (cumulative count 3)
client_1  | 35.973% <= 0.207 milliseconds (cumulative count 35973)
client_1  | 95.481% <= 0.303 milliseconds (cumulative count 95481)
client_1  | 99.825% <= 0.407 milliseconds (cumulative count 99825)
client_1  | 99.927% <= 0.503 milliseconds (cumulative count 99927)
client_1  | 99.971% <= 0.607 milliseconds (cumulative count 99971)
client_1  | 99.990% <= 0.703 milliseconds (cumulative count 99990)
client_1  | 99.999% <= 0.807 milliseconds (cumulative count 99999)
client_1  | 100.000% <= 0.903 milliseconds (cumulative count 100000)
client_1  |
client_1  | Summary:
client_1  |   throughput summary: 115606.94 requests per second
client_1  |   latency summary (msec):
client_1  |           avg       min       p50       p95       p99       max
client_1  |         0.225     0.064     0.223     0.303     0.343     0.863
client_1  |
redis-benchmark_client_1 exited with code 0
```

### 2. Write a certain amount of k-v data, evaluate the data size from 1w to 50w, combine the info memory information before and after writing, and analyze the average memory space occupied by each key for different value sizes mentioned above.

* execution
```bash
$ cd ./example/go-client-analysis
# repeat {TIMES} set {SIZE} bytes data into redis 
$ SIZE=500000 TIMES=5,5000,50000 docker-compose up
$ docker-compose down
```

* view analysis report
```bash
$ cat ./reports/redis-analysis-server-6379-0{timestamp}.csv
Key,Count,Size,NeverExpire,AvgTtl(excluded never expire)
500000bytes_2000:*,2000,10.864 MB,2000,0
500000bytes_10:*,10,55.625 KB,10,0
```