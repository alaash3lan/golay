
# GOLAY

A simple Golang app to be used as intro layout for golang app , it's folowing the DDD architecture , using Gin as router and GORM library



## Run Locally

Clone the project

```bash
  git clone https://github.com/alaash3lan/golay.git
```

Go to the project directory

```bash
  cd golay
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run main
```

To add new domain

```bash
  go run cmd/helper/main.go helper make:domain product 
```
