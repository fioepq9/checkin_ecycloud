FROM golang:1.18.3

WORKDIR /workspace
ADD . .

RUN go env -w GO111MODULE="on" \
&& go env -w GOPROXY="https://goproxy.cn,direct" \
&& go mod tidy

ENV TZ="Asia/Shanghai"

CMD ["go", "run", "main.go"]
