# syntax=docker/dockerfile:1

# Build Stage
# pull golang image
FROM golang:1.16-buster as build

# create workdir
WORKDIR /build

# copy over mod file
COPY go.mod ./
RUN go mod download

# copy the .go files
COPY *.go ./

# run the build command to compile the code
RUN go build -o /docker-get-going


# Run Stage
# from debian10
FROM gcr.io/distroless/base-debian10

# set workdir
WORKDIR /

# copy only the binary
COPY --from=build /docker-get-going .

# expose the port
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-get-going"]