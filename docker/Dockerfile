FROM golang:alpine3.20

RUN apk add curl git
RUN mkdir /output

WORKDIR /tmp
RUN curl -O -J -L https://github.com/Synoptik-Labs/geoip2-haproxy-ranges/archive/refs/heads/main.zip && \
    unzip geoip2-haproxy-ranges-main.zip

COPY ./generate_list.sh /tmp/geoip2-haproxy-ranges-main/

WORKDIR /tmp/geoip2-haproxy-ranges-main
RUN go mod download && go mod verify
CMD [ "/bin/sh", "generate_list.sh" ]
