# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

FROM  golang AS builder
ADD . /src
WORKDIR /src
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
RUN go mod tidy
RUN go build -o /lear_log .

# final stage
FROM ubuntu
COPY --from=builder /lear_log ./
RUN chmod +x ./lear_log

ENTRYPOINT ["./lear_log"]
