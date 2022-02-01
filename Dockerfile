FROM golang as build

WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o example .

FROM debian:10-slim
RUN apt-get update && apt-get install -y ca-certificates
RUN rm /etc/localtime && echo Asia/Tokyo > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata
WORKDIR /workspace
COPY --from=build /workspace/example /bin/example

ENTRYPOINT ["/bin/gatekeeper"]