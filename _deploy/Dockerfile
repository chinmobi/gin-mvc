## Copyright 2020 Zhaoping Yu.  All rights reserved.
## Use of this source code is governed by a MIT style
## license that can be found in the LICENSE file.

FROM alpine:3.12

# If not building with CGO_ENABLED=0
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# If need SSL certificates for HTTPS
#RUN apk add --no-cache ca-certificates


ENV APP_HOME=/opt/chinmobi/app \
    APP_USER=chinmobi

RUN set -x \
    && addgroup -g 1010 -S $APP_USER \
    && adduser -S -D -H -u 1010 -h $APP_HOME -s /sbin/nologin -G $APP_USER -g $APP_USER $APP_USER \
    && mkdir -p $APP_HOME \
    && chown -R $APP_USER:$APP_USER $APP_HOME

USER $APP_USER


WORKDIR $APP_HOME


COPY --chown=$APP_USER:$APP_USER bin $APP_HOME/bin/


COPY --chown=$APP_USER:$APP_USER prod.env $APP_HOME/app.env
COPY --chown=$APP_USER:$APP_USER docker-entrypoint.sh $APP_HOME/
ENTRYPOINT ["./docker-entrypoint.sh"]


EXPOSE 8080

CMD ["start-app"]
