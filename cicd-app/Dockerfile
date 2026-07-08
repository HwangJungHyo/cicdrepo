# 1단계: 빌드
FROM golang:1.21 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /app .

# 2단계: 실행 (작은 이미지)
FROM gcr.io/distroless/static
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
