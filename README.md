# Kafka Upload API
Provide an API to upload data to a Kafka topic.  

## Overview

- Language: Go v1.21.4
- Web FrameWork: Gin v1.10.0
- Tool: Kafka-Go v0.4.47


## ENV
copy .env.example and rename as .env  
```
KAFKA_BROKER=localhost:9092
```

## Run

### Update Modules
```
go get -u && go mod tidy -v
```


### Run
```
go run cmd/main.go
```



## API
- POST /kafka-upload/{topic}
