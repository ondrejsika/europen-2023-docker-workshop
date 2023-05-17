FROM golang:1.20 as build
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN go build

FROM scratch
WORKDIR /app
COPY --from=build /build/example-server .
CMD ["./example-server"]
EXPOSE 8000