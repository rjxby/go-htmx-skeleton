# Frontend build stage
FROM node:14-alpine AS build-frontend

WORKDIR /frontend

COPY ./frontend ./
RUN npm install

# Backend build stage
FROM golang:1.21.1-alpine as build-backend

WORKDIR /backend

COPY ./backend ./
RUN go mod download

COPY --from=build-frontend ./frontend ./app/public

RUN go build -tags embed -o service

# Final stage
FROM alpine:3.14

WORKDIR /srv

COPY --from=build-backend /backend/service /srv

EXPOSE 8080

CMD ["/srv/service"]
