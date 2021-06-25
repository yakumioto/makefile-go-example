FROM golang:1.16 AS builder

ARG command

WORKDIR /src

COPY . /src

RUN bash -c "${command}"

FROM scratch

LABEL maintainer="yakumioto <yaku.mioto@gmail.com>"

COPY --from=builder /app /app

CMD ["/app"]