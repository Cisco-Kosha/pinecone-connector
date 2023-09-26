FROM docker.io/golang:alpine as builder
RUN apk add git
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN rm /build/go.sum
RUN go mod tidy
RUN go build -o pinecone-connector .
FROM docker.io/alpine
RUN apk --no-cache add curl
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/pinecone-connector /app/
WORKDIR /app
CMD ["./pinecone-connector"]
