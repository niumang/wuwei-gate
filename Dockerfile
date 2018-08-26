ARG GOLANG_VERSION=1.10.1
FROM golang:${GOLANG_VERSION}

ARG GOTOOLS="github.com/go-chi/chi github.com/go-chi/render"

RUN go get -u -v ${GOTOOLS} && mkdir -p ${GOPATH}/src/wuwei-gate

WORKDIR $GOPATH/src/wuwei-gate
COPY .  $GOPATH/src/wuwei-gate
RUN go build

