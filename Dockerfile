FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build/cmd/tictactoe
RUN go build -o main .

FROM alpine
COPY --from=builder /build/cmd/tictactoe /app/
WORKDIR /app
CMD ["./main"]
