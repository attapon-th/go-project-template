ARG APPNAME="AppFiberV2"

FROM alpine:3.15

ARG APPNAME

WORKDIR /app

RUN apk -U --no-cache add tzdata tini \
    && chown 1000:1000 -R /app \
    && rm -rf /root /tmp/* /var/cache/apk/* && mkdir /root

COPY ./build/${APPNAME} /app/${APPNAME}


# Set TimeZone
ENV TZ=Asia/Bangkok
ENV ENVOLOPMENT=production
ENV APPNAME="${APPNAME}"

USER "1000:1000"

ENTRYPOINT ["tini", "--"]
CMD /app/$APPNAME serv

EXPOSE 80