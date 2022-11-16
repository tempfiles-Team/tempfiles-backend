FROM ubuntu:latest

ARG TARGETARCH

WORKDIR /app

COPY dist/tempfiles-backend_linux_$TARGETARCH/tempfiles-backend /

EXPOSE 5000

CMD ["/app/tempfiles-backend"]