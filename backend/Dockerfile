# Go image
FROM golang:1.24

# Жұмыс директориясын орнату
WORKDIR /app

# Go модульдерді көшіру
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Қалған кодты көшіру
COPY . .

# Билд жасау
RUN go build -o main .

# 8080 портын ашу
EXPOSE 9090

# Қосу
#CMD ["./main"]

CMD ["sh", "-c", "sleep 5 && ./main"]

