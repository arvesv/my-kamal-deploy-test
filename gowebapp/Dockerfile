FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE  80
CMD [ "./main" ]

