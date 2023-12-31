FROM golang:1.20 AS builder
ADD . /src
WORKDIR /src
# Fetch dependencies
# RUN go mod vendor


# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.cn make build


FROM scratch
COPY --from=builder /src/config.yaml /app/config.yaml
COPY --from=builder /src/bin /app
WORKDIR /app

EXPOSE 5040
EXPOSE 5050
EXPOSE 5060
ENTRYPOINT ["./erp"]
#CMD ["-conf", ""]



