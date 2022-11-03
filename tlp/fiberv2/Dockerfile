ARG APPNAME
FROM alpine:3.15

RUN apk -U --no-cache add dumb-init tzdata ca-certificates && update-ca-certificates


WORKDIR /app

ADD ./build/maapi /bin/main


# Set TimeZone
ENV TZ=Asia/Bangkok
ENV ENVOLOPMENT=production

EXPOSE 80

ENTRYPOINT ["/usr/bin/dumb-init", "--"]


ENV APPNAME=${APPNAME|main}
CMD ["sh", "-c", "/app/${APPNAME}"]

