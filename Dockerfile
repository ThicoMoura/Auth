FROM golang:1.19.1@sha256:122f3484f844467ebe0674cf57272e61981770eb0bc7d316d1f0be281a88229f

WORKDIR /go/src
ENV PATH="go/bin:${PATH}"
ENV CGO_ENABLE=1

RUN apt-get update

CMD ["tail", "-f", "/dev/null"]