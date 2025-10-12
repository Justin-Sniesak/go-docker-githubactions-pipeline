FROM golang:alpine3.22
WORKDIR /app
COPY gcpcrcb.go .
CMD ["go", "run", "gcpcrcb.go"]
