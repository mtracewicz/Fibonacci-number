FROM golang:1.16.3-alpine3.13 as build
WORKDIR /src
COPY . .
RUN ["go", "build", "-o", "main", "."]

FROM alpine:latest as execution
COPY --from=build /src /src
WORKDIR /src
CMD ["./main"]
