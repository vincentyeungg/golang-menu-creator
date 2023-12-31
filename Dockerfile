# Build stage

# base image for golang
FROM golang:1.20-alpine3.17 AS builder
# setup working directory
WORKDIR /app
# copy files from current directory to working directory of image
COPY . .
# build the go executable file from main, set output file as main
RUN go build -o main main.go
# add curl command
RUN apk add curl
# download golang migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz


# Run stage

# base image as alpine
FROM alpine:3.17
# setup woring directory
WORKDIR /app
# copy files from the build stage above using alias name 
# and save the executable file /app/main into the current working directory
# . denotes WORKDIR we set which is /app
COPY --from=builder /app/main .
# save golang-migrate from builder
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
# copy current migrations
COPY db/migration ./migration

# which port the server will listen on
EXPOSE 8080
# command to run the executable
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]