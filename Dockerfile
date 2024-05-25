FROM golang

WORKDIR /app

COPY ./cmd .

ENV PORT=8000

EXPOSE 8000

CMD [ "go","run","main.go" ]