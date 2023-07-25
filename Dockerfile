FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

ENV OMETRIA_APIKEY <your_api_key>
ENV MAILCHIMP_APIKEY <your_api_key>

ADD go.* /owlery/
WORKDIR /owlery
COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main.go

FROM scratch
# Copy our static executable.
COPY --from=builder /owlery /go/bin/owlery

ENTRYPOINT ["/go/bin/owlery/owlery"]