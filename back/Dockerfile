FROM ruby:2.3.4-alpine

ENV APP /app
WORKDIR $APP

RUN apk --update add --virtual build-deps \
    build-base \
  && apk add \
    postgresql-dev \
		make \
    tzdata \
		nodejs \
  && rm -rf /var/cache/apk/* \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

COPY . $APP

RUN bundle update && gem update && \
bundle install --jobs=4 --path vendor/bundle \
  && mkdir -p $APP/tmp/cache \
  && mkdir -p $APP/tmp/pids \
  && mkdir -p $APP/tmp/sockets

# RUN apk del build-deps \
#    && rm -rf /var/cache/apk/*

