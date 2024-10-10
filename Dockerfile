FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

COPY scripts/wait-for-it.sh /usr/local/bin/wait-for-it.sh

RUN chmod +x /usr/local/bin/wait-for-it.sh

RUN go build -v -o /usr/local/bin/app ./cmd/carline

EXPOSE 9990

CMD /usr/local/bin/wait-for-it.sh $DB_HOST:$DB_PORT \
    && /usr/local/bin/wait-for-it.sh $RABBITMQ_HOST:$RABBITMQ_PORT \
    && /usr/local/bin/app
