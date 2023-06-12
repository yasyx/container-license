
# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR /app

COPY . /app

RUN sed -i "s|default_project_uuid|shdssss-asdfs-jjj|g" /app/pkg/constants/constants.go && \
    CGO_ENABLED=0 GOOS=linux go build  -o /app/generate /app/cmd/generate && \
    CGO_ENABLED=0 GOOS=linux go build  -o /app/checker /app/cmd/checker

RUN /app/generate --month=1 --duration="2m"

# Deploy the application binary into a lean image

FROM nginx AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/checker /app/license /app/

EXPOSE 80 443

CMD ["/app/checker","--cmd=nginx" ,"--args=-g daemon off;"]
