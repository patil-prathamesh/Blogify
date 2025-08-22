FROM golang:1.24-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV PORT=8000
ENV MONGODB_URI=mongodb+srv://prathameshpatil2906:qIF0jZAGf4xCMfvr@cluster0.bj4sw0l.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
ENV DATABASE=sample_supplies
ENV USERS_COLLECTION=users
ENV POSTS_COLLECTION=posts
ENV JWT_SECRET=a*r2QW104bdhf()Q>+^{}

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 5000
CMD ["./app"]