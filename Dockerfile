FROM golang:1.21 as dev

WORKDIR /work

# Define ARG for UID and GID
ARG UID=1000
ARG GID=1000

# Create a new user and group using the provided UID and GID, and set permissions
RUN groupadd -g $GID usergroup && \
    useradd -u $UID -g $GID -ms /bin/bash user && \
    chown -R user:usergroup /work/ && \
    chown -R user:usergroup /go/

USER user

RUN go install golang.org/x/tools/cmd/godoc@latest
RUN go install github.com/kisielk/errcheck@latest

FROM golang:1.21 as build

WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# We don't want the entire go sdk in here
FROM alpine as runtime 

COPY --from=build /app/app /
CMD ./app
