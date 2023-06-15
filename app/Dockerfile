FROM golang:alpine as development

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go build -o main main.go

FROM alpine as production

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=development /app/main /app/

WORKDIR /app

EXPOSE 8080

CMD ["./main"]