FROM alpine:3.6

RUN apk add --update ca-certificates

ADD ./bysykkelMain /
ADD ./translation /

CMD /bysykkelMain
