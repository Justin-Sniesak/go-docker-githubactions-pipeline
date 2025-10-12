FROM golang:alpine3.22
WORKDIR /app
COPY utc-time-date.go .
CMD ["go", "run", "utc-time-date.go"]
