FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/server .

FROM alpine:3.7 AS bin
WORKDIR /app
COPY --from=build /out/server .

EXPOSE 8080
CMD ["./server"]