FROM golang:1.13.15-alpine as build
WORKDIR /src
COPY . .
RUN ["go", "build", "-o", "main", "."]

FROM alpine:latest as execution
COPY --from=build /src /src
WORKDIR /src
CMD ["./main"]
