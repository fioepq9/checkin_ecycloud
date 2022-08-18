FROM golang:1.18.3 as builder

WORKDIR /workspace
ADD . .

RUN go env -w GO111MODULE="on" \
&& go env -w GOPROXY="https://goproxy.cn,direct" \
&& go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# ===== ===== ===== ===== ===== ===== ===== ===== ===== ===== ===== ===== =====
FROM alpine:3.15.5 as prod

WORKDIR /workspace
COPY --from=builder /workspace/app .
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

CMD ["./app"]
