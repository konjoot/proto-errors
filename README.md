# proto-errors
Experiments with structured errors for services based on protobuf IDL and micro rpc-framework.

# Install and run

In first terminal

```
docker pull consul
docker run -p 8500:8500  consul
```
In second terminal

```
go get -u github.com/konjoot/proto-errors
cd $GOPATH/src/github.com/konjoot/proto-errors
go run main.go
```
