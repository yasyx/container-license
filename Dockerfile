
# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20.3 AS build-stage

WORKDIR /app

COPY . /app

RUN sed -i "s|default_project_uuid|PROJECT_UUID|g" /app/pkg/constants/constants.go && \
    CGO_ENABLED=0 GOOS=linux go build  -o /app/generate /app/cmd/generate && \
    CGO_ENABLED=0 GOOS=linux go build  -o /app/checker /app/cmd/checker && \
    echo `/app/generate --month=PROJECT_LICENSE_MONTH --duration=""`

# Deploy the application binary into a lean image

FROM APP_CONTAINER AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/checker /app/license /app/

EXPOSE 80 443

ENTRYPOINT ["/app/checker","--cmd=APP_CMD" ,"--args=APP_ARGS"]
