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
### Performance Loadtest 
#### NThread = 100, Request/Thread=10
TPS: 500
P99: 283(ms)
P95: 283(ms)
P90: 275(ms)
P75: 262(ms)
P50: 253(ms)
#### NThread = 200, Request/Thread=10
TPS: 666
P99: 387(ms)
P95: 387(ms)
P90: 381(ms)
P75: 335(ms)
P50: 278(ms)
#### NThread = 400, Request/Thread=10
TPS: 1333
P99: 468(ms)
P95: 468(ms)
P90: 455(ms)
P75: 316(ms)
P50: 287(ms)
#### NThread = 800, Request/Thread=10
TPS: 4000
P99: 466(ms)
P95: 465(ms)
P90: 462(ms)
P75: 345(ms)
P50: 253(ms)
## **Documentation**
- Blog:
    - [Smart Batching](https://mechanical-sympathy.blogspot.com/2011/10/smart-batching.html)
