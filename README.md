# Description

This is implementation of simple web based API that steps through the Fibonacci sequence. For how-to run please see [README.md](./cmd/serverd/README.md).

## Functional description

Service exposes three endpoints:

### GET /current
Returns the current number in the sequence.

Send a request to the running service instance ( presuming its running on port 80 ):

```bash
curl 'http://localhost/current' -v
```

### GET /next
Returns the next number in the sequence. If the next term would overflow largest allowed value in the sequence error will be returned instead.

Send a request to the running service instance ( presuming its running on port 80 ):

```bash
curl 'http://localhost/next' -v
```

### GET /previous
Returns the previous number in the sequence. If the previous term would overflow smallest allowed value in the sequence error will be returned instead.

Send a request to the running service instance ( presuming its running on port 80 ):

```bash
curl 'http://localhost/previous' -v
```

## Requirements and Implementation

Solution was implemented having following presumptions in mind:

* API is supposed to handle `> 1k` requests ( see benchmark below ).
* If requested term is above `MaxThTerm` that is currently `92th` term in the sequence which maximum number fitting to `int64` error will be returned.
* The API will recover in event of panic.

Current Fibonacci term calculation implementation uses `O(n)` where `0 <= n < 92`, which is *not* ideal but within expected throughput range. Note: benchmarking was done on PC running Intel Core i7-3667U processor (2 cores, 2.0GHz, 4MB cache), where counter is reset each time reaching `MaxThTerm` to simulate `O(n)`.

```
wrk -t12 -c400 -d30s http://localhost:8000/next
Running 30s test @ http://localhost:8000/next
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    21.31ms   29.89ms 375.91ms   88.41%
    Req/Sec     2.73k     1.22k    9.92k    78.17%
  979697 requests in 30.07s, 119.48MB read
Requests/sec:  32580.73
Transfer/sec:      3.97MB
```

Considered/Alternative approaches: 

* [Binet's formula](https://en.wikipedia.org/wiki/Fibonacci_sequence#Binet's_formula) will not work using standard data types such as `float64` due loosing precision on the higher terms, for example `88th` term would result into not a valid sequence number. Alternative approach might be to leverage [Binet's formula](https://en.wikipedia.org/wiki/Fibonacci_sequence#Binet's_formula) using [big](https://pkg.go.dev/math/big) library.
* Implement algorithm in [Matrix_form](https://en.wikipedia.org/wiki/Fibonacci_sequence#Matrix_form).


## Improvements

This application can be improved in many different ways, several ideas are listed below.

* API should respond with empty JSON body in event of error.
* Use more suitable tool for documenting API spec.
