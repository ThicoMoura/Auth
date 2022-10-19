FROM golang:1.19.1@sha256:122f3484f844467ebe0674cf57272e61981770eb0bc7d316d1f0be281a88229f

WORKDIR /go/src/github.com/ThicoMoura/Auth
ENV PATH="go/bin:${PATH}"
ENV CGO_ENABLE=1

RUN apt-get update \
    && go install github.com/go-task/task/v3/cmd/task@latest \
    && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
    && go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

CMD ["tail", "-f", "/dev/null"]