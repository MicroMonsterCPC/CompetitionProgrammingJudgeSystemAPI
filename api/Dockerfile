FROM golang 

ENV PATH $PATH:$GOPATH/bin:$GOROOT/bin

COPY . /echo-server

WORKDIR /echo-server
RUN sh ./getfile

EXPOSE 1323

CMD ["go", "run", "Main.go"]
