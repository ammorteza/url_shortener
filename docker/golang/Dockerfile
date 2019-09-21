#start from the latest golang image
FROM golang:latest

#define maintainer information
LABEL maintainer="Morteza MazdaGerd"

#set the current application work directory
WORKDIR /app

COPY ./go.mod ./
RUN go mod download

#copy source to app directory
COPY . .

#build the go app
RUN cd cmd/
RUN go build
RUN cd ..

#expose port 8080 to real world
EXPOSE 8080

#execute binary file
CMD ["./main"]