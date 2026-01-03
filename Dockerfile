FROM golang

WORKDIR /app

RUN useradd -m appuser

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -buildvcs=false -o main .

RUN chown -R appuser:appuser /app

USER appuser

EXPOSE 8080
ENTRYPOINT ["./main"]
CMD ["-port", "8080"]
