# Use official Golang image as the base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /BookCrud_App

# Copy go.mod and go.sum first to use layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Make sure logs directory exists inside the cmd folder 
RUN mkdir -p /BookCrud_App/cmd

# Optional: download MySQL driver if not in go.mod already
RUN go get -v github.com/go-sql-driver/mysql

# Build the Go app from the cmd directory and place the binary at project root
WORKDIR /BookCrud_App/cmd
RUN go build -o /BookCrud_App/main .

# Default command to run the app
CMD ["/BookCrud_App/main"]
