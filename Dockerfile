FROM golang

ENV APP_PORT="8080"

ENV DB_USER="local"
ENV DB_PASSWORD="local"
ENV DB_HOST="localhost"
ENV DB_PORT="3306"
ENV DB_NAME="local-db"

ADD . /
WORKDIR /
RUN go build -o main .
CMD ["/main"]