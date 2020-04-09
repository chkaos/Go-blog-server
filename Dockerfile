FROM scratch

WORKDIR $GOPATH/src/github.com/chkaos/Go-blog-server
COPY . $GOPATH/src/github.com/chkaos/Go-blog-server
RUN go build .

EXPOSE 3001
ENTRYPOINT ["./Go-blog-server"]