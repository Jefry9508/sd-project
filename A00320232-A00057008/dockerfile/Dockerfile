FROM golang:alpine3.6 AS binary
ADD . /app
WORKDIR /app
RUN go build -o configurations

FROM alpine:3.6
WORKDIR /app
ENV PORT 8000
EXPOSE 8000
COPY --from=binary /app/configurations /app
CMD ["/app/configurations"]