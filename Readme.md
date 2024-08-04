# Report

The report is in file crawl-report.csv

# run tests and create a coverprofile

```sh
go test ./... -coverprofile=cover.out
```

# open the interactive UI to check the Coverage Repor

```sh
go tool cover -html=cover.out
```

# benchmark & profiling

```sh
go test -bench=. -run=x -benchmem -memprofile mem.prof -cpuprofile cpu.prof -benchtime=10s > 0.bench
go tool pprof cpu.prof

```

# Run the binary

```sh
go build -o webcrawler

./webcrawler
```

- cover.out : Contains the test coverage

1. Used a read-write mutex to allow concurrent reads but exclusive writes
2. Consider using more efficient data structures for tracking visited URLs, such as a concurrent map.
3. HTTP client is reusing connections efficiently by using http.Transport for connection pooling.
4. Use of concurrency to speed up the crawl:
   - created a struct so goroutines can share access
   - map keeps track of the pages crawled (Improvement - Consider using a more efficient data structure, such as a concurrent map. This also has a loadOrstore method to check for existing keys which is helpful in this case)
   - Created a buffered channel as a worker pool to ensure not too many goroutines are created.
   - used waitgroup to wait until all the in-flight goroutines are done
