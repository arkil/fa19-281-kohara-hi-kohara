FROM golang:latest 
EXPOSE 3000
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
ENV PORT 3000
ENV MONGO mongodb://admin:admin@10.4.1.57:27017/?authSource=admin
RUN cd /app ;
RUN export GOPATH=$(pwd);export PATH=$PATH:$GOPATH/bin
RUN go get -v github.com/codegangsta/negroni
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/unrolled/render
RUN go get -v github.com/satori/go.uuid
RUN go get gopkg.in/mgo.v2
RUN go install like-read
CMD ["/app/bin/like-read"]
