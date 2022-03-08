FROM golang

WORKDIR /backend

COPY . ./
RUN go mod download

EXPOSE 50051
CMD ["./start_service.sh"]

