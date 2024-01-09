FROM golang:1.21 as builder

WORKDIR /go/src/faraway
COPY . .
RUN go mod download
RUN useradd -u 10001 app_user
RUN CGO_ENABLED=0 go build -v -o /out/app ./cmd/server

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
USER app_user

WORKDIR /root/
COPY --from=builder /out/app .
CMD ["./app"]
