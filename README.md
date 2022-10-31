<!-- ABOUT THE PROJECT -->
## About The Project

Go app to read open movie DB and query the open movie API with the results.
For the possible arguments -h flag




### Prerequisites

This projects was built with go 1.19.1 so you should have that installed or build the dockerized version

### Installation
Before running be sure to have a copy of the omdb in the working directory

   ```sh
   make build
   ```
or without

```sh
go build /cmd/main.go
```



<!-- USAGE EXAMPLES -->
## Usage
OMDB APi needs an api the best way is to set it on the env or pass it to the available config params
```sh
export OMDB_APIKEY = {your_api_key}
```
RUN
```sh
docker run --env OMDB_APIKEY go_chall {ARGS}
```
or

```sh
go run /cmd/main.go {ARGS}
```
### Tests
```sh
make test
```
### Coverage
```sh
make cover
```






