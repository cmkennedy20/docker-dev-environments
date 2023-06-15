FROM golang:latest
ENV NAME "HELLO WORLD"
ENTRYPOINT [ "bash", "/go/src/start.sh" ]