FROM golang:1.19-alpine as build
WORKDIR /
COPY go.mod ./
COPY *.go ./
RUN go build -o /gomyip

FROM alpine:latest
COPY --from=build /gomyip /
EXPOSE 8080
CMD [ "/gomyip" ]