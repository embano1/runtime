Simple Go binary useful to print basic runtime information (CPU) and optionally stress CPU, e.g. in Docker containers.

# Build
```
make build # for Linux
go build . # for your OS
```

# Build binary, Docker container and push to a registry
Modify `"LABEL MAINTAINER"` (Dockerfile) and `"REPOTAG"` (Makefile), then:

```
make all
``` 

# Use (binary)
To print just runtime information (CPUs, Goroutines):
```
./runtime
``` 

To print runtime information and stress 2 CPUs:
```
./runtime -c 2
``` 

# Example (Docker)
Burn 4 CPUs and limit the container CPU resources (2 CPUs). Use a tool like ctop or top to see what´s going on:
```
docker run --rm -i --cpus 2 embano1/runtime:1.0 -c 4
```