FROM ubuntu:bionic

COPY choiny /app/

ENTRYPOINT ["/app/choiny"]

EXPOSE 8080
