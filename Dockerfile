# Backend
FROM golang:latest AS go-builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./backend/ ./backend
WORKDIR /app/backend
RUN go build -o main .

# Frontend
FROM node:20.11.1-bullseye-slim AS react-builder
WORKDIR /app
COPY ./client/package.json ./client/package-lock.json ./
RUN npm install --ignore-scripts
COPY ./client/ .
RUN npm run build

# Final image
FROM golang:latest
WORKDIR /app

COPY --from=go-builder /app/backend/main ./backend/main
COPY --from=react-builder /app/dist ./client/dist

# Expose port 5000 and 80 to the outside world
EXPOSE 5000
EXPOSE 80
WORKDIR /app/backend

# Command to run the executable
CMD ["./main"]

