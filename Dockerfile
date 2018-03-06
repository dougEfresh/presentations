from golang
WORKDIR /go/src/app
COPY . .

RUN go get -v  golang.org/x/tools/cmd/present
RUN go install golang.org/x/tools/cmd/present

CMD ["present"]