FROM golang:1.20 as builder


RUN mkdir /app
ADD . /app
WORKDIR /app

COPY . .

# Download all the dependencies
RUN go mod download
RUN go mod verify
RUN go build -o main

# GO Repo base repo
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

RUN mkdir /app

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# USER employeemgmt

# Expose port 8080
EXPOSE 8080

# Run Executable
CMD ["./main"]

