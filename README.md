# Tutorial: How to implement Jaeger and Opentracing as tracing middleware.

This repo contains the code for the OpenTracing & Jaeger step-by-step tutorial that you can find [here](https://medium.com/from-the-edge/tutorial-how-to-implement-jaeger-and-opentracing-as-tracing-middleware-e3e693ee0802).

## Run your API, send a curl request, and check your span in Jaeger.
```
// First, run the Jaeger container.
make jaeger-up

// Apply a SERVICE_NAME environment variable.
export SERVICE_NAME=jaeger-test

// Apply a DEV_ENV environment variable.
export DEV_ENV=true

// Alternatively, you can also set both env vars by applying the .passrc file:
source .passrc

// Then, build and run the API.
make build
make run

// In a new terminal, send a curl request to your API.
curl http://localhost:9092/data

// To nuke the Jaeger container:
make jaeger-down
```

Finally, navigate to http://localhost:16686/ to see your service's traces displayed on the Jaeger UI!