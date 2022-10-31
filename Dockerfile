FROM ubuntu:18.04 AS build
# update ubuntu packages
RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get install -y wget lsof

# install go 1.19
RUN wget https://dl.google.com/go/go1.19.2.linux-amd64.tar.gz && \
    tar -xf go1.19.2.linux-amd64.tar.gz && \
    mv go /usr/local && \
    rm go1.19.2.linux-amd64.tar.gz

ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH
ENV GOPROXY="https://proxy.golang.org"

RUN mkdir -p /app/
WORKDIR /app/

COPY . .
RUN go mod tidy
RUN go build cmd/main.go
RUN mv main ../

FROM ubuntu:18.04 AS release
WORKDIR /usr/local/bin/
COPY --from=build /main .

ENTRYPOINT ["/usr/local/bin/main"]

