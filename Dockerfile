FROM renlu/telize
MAINTAINER renlu<xurenlu@gmail.com>


RUN apt-get update && apt install -y golang-1.18-go
RUN mkdir /www/dev/gmm/
COPY ./src/go.mod /www/dev/gmm/
COPY ./src/go.sum /www/dev/gmm/
COPY ./src/main.go /www/dev/gmm/
RUN cd /www/dev/gmm &&  /usr/lib/go-1.17/bin/go mod download && /usr/lib/go-1.17/bin/go mod download github.com/gin-gonic/gin &&  /usr/lib/go-1.17/bin/go build -o ./gmm ./main.go
EXPOSE 7654
CMD ["/www/dev/gmm/gmm"]

