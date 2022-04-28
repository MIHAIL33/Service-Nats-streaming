# REST server with NATS-streaming
## Dev
This server uses Nats-streaming and PostgresSQL.
```
./.devcontainer/docker-compose up -d
```
## Executable files
### producer
```
go build ./cmd/producer.go
```
The script reads the `./task/model.json` file and writes it to the Nats-streaming channel and then exits.

### main
```
go build ./cmd/main.go
```
The script starts the server, reads data from the database, and writes the data to the cache. Starts the Nats-streaming channel listener, reads data from there, writing it to the cache and database.

#### Routes
- **GET /api/models** (get all models)
- **GET /api/models/:id** (get model by id)
- **POST /api/models** (create new model)
- **DELETE /api/models/:id** (delete model by id)


- **GET /api/models/cache** (get all models from cache)
- **GET /api/models/cache/:id** (get model by id from cache)

## Swagger

Route: **/swagger/index.html**

## Stress Testing

### WRK
From cache:
```
wrk -t2 -c5 -d30s --timeout 2s http://localhost:8000/api/models/cache/1
Running 30s test @ http://localhost:8000/api/models/cache/1
2 threads and 5 connections
Thread Stats   Avg      Stdev     Max   +/- Stdev
Latency   130.73us  398.61us  19.65ms   99.02%
Req/Sec    17.95k   711.94    18.92k    87.17%
1071620 requests in 30.01s, 0.94GB read
Requests/sec:  35708.85
Transfer/sec:     32.05MB
```
From database:
```
wrk -t2 -c5 -d30s --timeout 2s http://localhost:8000/api/models/1
Running 30s test @ http://localhost:8000/api/models/1
2 threads and 5 connections
Thread Stats   Avg      Stdev     Max   +/- Stdev
Latency     2.81ms    3.78ms  47.82ms   86.01%
Req/Sec     1.17k   216.97     1.83k    68.67%
69941 requests in 30.02s, 62.77MB read
Requests/sec:   2329.65
Transfer/sec:      2.09MB
```

### Vegeta
From cache:
```
echo "GET http://localhost:8000/api/models/cache/1" | vegeta attack -duration=30s -rate=11000 | vegeta report
Requests      [total, rate, throughput]         330001, 10999.62, 10955.16
Duration      [total, attack, wait]             30.003s, 30.001s, 1.615ms
Latencies     [min, mean, 50, 90, 95, 99, max]  51.394µs, 4.403ms, 454.535µs, 8.195ms, 14.886ms, 81.342ms, 9.29s
Bytes In      [total, mean]                     268535645, 813.74
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           99.60%
Status Codes  [code:count]                      0:1316  200:328685  
Error Set: Get "http://localhost:8000/api/models/cache/1": dial tcp 0.0.0.0:0->127.0.0.1:8000: socket: too many open files
```
Shell file descriptor file system limits: `ulimit -n 4096`

From database:
```
echo "GET http://localhost:8000/api/models/1" | vegeta attack -duration=30s -rate=480 | vegeta report
Requests      [total, rate, throughput]         14400, 480.03, 475.78
Duration      [total, attack, wait]             30.066s, 29.998s, 68.284ms
Latencies     [min, mean, 50, 90, 95, 99, max]  581.348µs, 17.731ms, 1.326ms, 1.889ms, 120.584ms, 437.912ms, 774.87ms
Bytes In      [total, mean]                     11695140, 812.16
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           99.34%
Status Codes  [code:count]                      200:14305  500:95  
Error Set: 500 Internal Server Error
```