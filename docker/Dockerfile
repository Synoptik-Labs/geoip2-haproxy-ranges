FROM golang:alpine3.20

RUN apk add curl git
RUN mkdir /output

WORKDIR /tmp
RUN curl -O -J -L https://github.com/Synoptik-Labs/geoip2-haproxy-ranges/archive/refs/heads/master.zip && \
    unzip geoip2-haproxy-ranges-master.zip

COPY ./generate_list.sh /tmp/geoip2-haproxy-ranges-master/

WORKDIR /tmp/geoip2-haproxy-ranges-master
CMD [ "/bin/sh", "generate_list.sh" ]
