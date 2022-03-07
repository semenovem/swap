ARG BASE_IMAGE=golang:1.15.7

# Сборка
FROM $BASE_IMAGE
WORKDIR /app
ARG _PKG_="bhive/pkg/app/build"
COPY . .
ARG BUILD_VERSION
ARG BUILD_HASH_COMMIT
ARG BUILD_BRANCH
ARG BUILD_TIME
RUN go build -o app-sim-sftp-client \
      -mod vendor \
      -ldflags="-X '$_PKG_.version=$BUILD_VERSION' -X '$_PKG_.hashCommit=$BUILD_HASH_COMMIT' -X '$_PKG_.branch=$BUILD_BRANCH' -X '$_PKG_.time=$BUILD_TIME'"


# Образ с приложением
ARG BASE_IMAGE=golang:1.15.7
FROM $BASE_IMAGE
WORKDIR /app
COPY --from=0 /app/app-sim-sftp-client .
CMD ./app-sim-sftp-client
