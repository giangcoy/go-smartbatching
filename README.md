# Golang SmartBatching
An implementation smartmatching in golang based on [Smart Batching](https://mechanical-sympathy.blogspot.com/2011/10/smart-batching.html)
    
## Quick Start
### Download and install
```bash
go get -u -v github.com/giangcoy/go-smartbatching
```
## Example
> A complete example of balance hotspot account with TPS(~100K), please check [example](/example)
### Build and run
```bash
go build balance.go

./balance -a="mysql adress"
```
### Benchmark
We do it with 10 request/thread and mysql(local)

#### TPS( higher is better)
![TPS](https://user-images.githubusercontent.com/17874394/191669973-55c08ee5-a925-456e-95d3-04e80893c6e1.png)
#### Latency(lower is better)
![Latency](https://user-images.githubusercontent.com/17874394/191670018-5b8c2158-89e1-4ab8-b5a8-63c8c6b40853.png)
## **Documentation**
- Blog:
    - [Smart Batching](https://mechanical-sympathy.blogspot.com/2011/10/smart-batching.html)
