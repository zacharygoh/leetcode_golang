# Dockerfile References: https://docs.docker.com/engine/reference/builder/
 
# Start from the latest golang base image
FROM golang:latest as builder
 
# Add Maintainer Info
LABEL maintainer="Zachary <zacharygoh@pingspace.co>"
 
# Set the Current Working Directory inside the container
WORKDIR /go/src/space-ims-api
 
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
 
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest  
 
RUN apk --no-cache add ca-certificates
 
WORKDIR /go/src/space-ims-api
 
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/space-ims-api .
 
# Expose port 3333 to the outside world
EXPOSE 3333
 
# Command to run the executable
CMD ["./main"]