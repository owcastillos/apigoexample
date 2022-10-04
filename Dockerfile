FROM golang

ENV APP_PORT="8080"

ENV DB_USER="admin"
ENV DB_PASSWORD="Qazwsx123."
ENV DB_HOST="api-test-database.cr7dqtlaxnco.us-east-2.rds.amazonaws.com"
ENV DB_PORT="3306"
ENV DB_NAME="api-test-database"

ADD . /
WORKDIR /
RUN go build -o main .
CMD ["/main"]