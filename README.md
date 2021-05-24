# echo-server

This mini echo server needs for debug source IP

## Usage

### Local
For local usage just start
```shell
# Default port is 8080
go run main.go 
```
If you want to specify port, just add as environment variable `PORT`
```shell
PORT=8888 go run main.go
```

### As Docker Container

From source
```shell
docker build -t echo-server:dev .
docker run -p 8080:8080 echo-server:dev
```

By Docker Compose
```shell
docker-compose up
```

From docker hub
```shell
docker run -p 8080:8080 afrofunkylover/echo-server
```

### As Helm Chart

Using the default helm chart.
```shell
helm repo add afrofunkylover https://
helm install stable afrofunkylover/echo-server
```
