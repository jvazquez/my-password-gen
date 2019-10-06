FROM golang:1.13.1-buster
RUN go get -u github.com/sethvargo/go-diceware/diceware
COPY . /my-password-gen
WORKDIR /my-password-gen/cmd
EXPOSE 8000
CMD ["go", "run", "xkcd.go"]