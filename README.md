# Golang SmartBatching
An implementation smartmatching in golang based on [Smart Batching](https://mechanical-sympathy.blogspot.com/2011/10/smart-batching.html)
    
## Quick Start
### Download and install
```bash
go get -u -v github.com/giangcoy/go-smartbatching
```
## Example
> A complete example of balance hotspot account with TPS(>5K), please check [example](/example)
### Build and run
```bash
go build balance.go

./balance -a="mysql adress"
```
### Benchmark
We do it with 10 request/thread and mysql(local)

#### TPS( higher is better)
![alt text](https://github.com/giangcoy/go-smartbatching/tree/main/examples/images/TPS.png?raw=true)
#### Latency(lower is better)
![alt text](https://github.com/giangcoy/go-smartbatching/tree/main/examples/images/Latency.png?raw=true)
## **Documentation**
- Blog:
    - [Smart Batching](https://mechanical-sympathy.blogspot.com/2011/10/smart-batching.html)
