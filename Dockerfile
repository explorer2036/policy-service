FROM alpine:latest
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.8/main" > /etc/apk/repositories
RUN apk add --update curl bash && rm -rf /var/cache/apk/*
WORKDIR /app
COPY dist/alpine/policy-service /app
ENTRYPOINT ["/app/policy-service", "serve"]