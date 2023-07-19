# Set the base image, alias as the builder stage of docker file
FROM golang:1.17 AS builder 
# Create an app directory
RUN mkdir /app
# Add all the conetents of the current directory to this app directory
ADD . /app
# Specify the working directory as this app directory
WORKDIR /app
# Having access to the go cli from the base image, we can run go build
# Build the compiled binary of the go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# Define the second stage of our multi stage dockerfile
# Define a lightweight base image, alpine latest, aliased as production stage
FROM alpine:latest AS production
# Copy it from the builder stage, the app compiled binary=
COPY --from=builder /app .
# Specify the command the execute this app compiled binary
CMD ["./app"]


# Benefits of multi stage
# reduced resource usage from the first stage not necessary in second stage

# docker build -t go-rest-api .

# docker run go-rest-api   
# or 
# docker run -it -p 8080:8080 go-rest-api