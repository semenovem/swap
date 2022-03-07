ARG BASE_IMAGE
FROM $BASE_IMAGE
WORKDIR /app

COPY . .
ARG BUILD_VERSION

RUN go build  \
  -mod vendor \
  -ldflags="-X 'main.buildVersion=$BUILD_VERSION'" \
  -o /app/app-clearing-service \
  clearing/*.go


ARG BASE_IMAGE
FROM $BASE_IMAGE
WORKDIR /app
COPY --from=0 /app/app-clearing-service .
CMD ./app-clearing-service
