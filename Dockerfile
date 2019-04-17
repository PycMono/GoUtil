#源镜像
#FROM golang:latest
#WORKDIR $GOPATH
#COPY . .
#CMD ["/",, "build.bat"]

FROM golang:latest
WORKDIR $GOPATH/src/moqikaka.com/Test
COPY . $GOPATH/src/moqikaka.com/Test
RUN go build .
#expose会暴露端口出去
EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]