FROM daocloud.io/library/golang:1.6.0

WORKDIR /app/gopath/wall
ENV GOPATH /app/gopath

RUN git clone --depth 1 git://github.com/toukii/wall.git . && go get -u github.com/toukii/wall && mv $GOPATH/bin/wall /bin/wall

EXPOSE 80

CMD ["/bin/wall"]


