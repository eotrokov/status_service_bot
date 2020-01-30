FROM golang:alpine as build
RUN mkdir /app
WORKDIR /app
COPY  . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -cgo -o main main.go

FROM scratch
COPY --from=build /app/main .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT [ "/main" ]
