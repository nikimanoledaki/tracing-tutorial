# Tutorial: How to implement Jaeger and Opentracing as tracing middleware.

This repo contains the code (which is subject to changes!) for an upcoming OpenTracing & Jaeger step-by-step tutorial.

## Run
```
go build main.go
go run main.go

TODO: curl request
```

Navigate to http://localhost:16686/ to see your service's traces show up on the Jaeger UI.

## Todo
- Add link to written blog post - this is being drafted.
- Endpoint should make an external service call to be able to demonstrate how child spans work. Trying to find the easiest external service call that the endpoint could make in order to demonstrate this!