# golang-prometheus-example
A sample code snippet to leanr how can connect with prometheus and monitor logs 


    Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud. 
    Prometheus collects and stores its metrics as time series data, i.e. metrics information is stored with the timestamp at which it was recorded, alongside optional key-value pairs called labels.


Build docker image
```
docker build -t go-prometheus-app .
```

You can simply run the following command to up the system as docker containers.

```go
   docker-compose up -d
```


### Find links here

[Prometheus official site](https://prometheus.io/)

[INSTRUMENTING A GO APPLICATION FOR PROMETHEUS](https://prometheus.io/docs/guides/go-application/)