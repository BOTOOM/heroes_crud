FROM botom/beego:v2

RUN mkdir -p $GOPATH/src/github.com/BOTOOM/heroes_crud

WORKDIR /usr/src/app

COPY . .

RUN go get

# CMD [ "bee", "run"]

EXPOSE 8080