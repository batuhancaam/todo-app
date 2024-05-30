# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .
RUN go mod download


# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-api



# Run
CMD [ "/todo-api" ]