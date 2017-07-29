FROM alpine:3.6

RUN apk add --update ca-certificates

ADD ./paupau.jpg /
ADD ./bysykkelMain /

CMD /bysykkelMain
