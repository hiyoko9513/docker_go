FROM golang:1.16-alpine

WORKDIR /go/app

COPY . ../../hiyoko

RUN set -eux \
  && apk update \
  && apk add --no-cache git curl 
  # && go get github.com/go-sql-driver/mysql 
  # && go get github.com/lib/pq 

# realize
# RUN go get -u github.com/oxequa/realize 

# air
RUN go get -u github.com/cosmtrek/air

# delve
RUN go get -u github.com/go-delve/delve/cmd/dlv
