# Go Chat Application with NATS.io

This is a simple chat application implemented in Go, using NATS for message brokering and Docker for running the NATS server.

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed on your machine
- [Go](https://golang.org/dl/) installed on your machine

## Run the App
```sh
docker compose up -d
```
This command will download the latest NATS image if you don't already have it, which starts the NATS server in the background.
<br>
<br>
Open a terminal, pass your username as flag to the run command, to run the application:

```sh
go run chat.go -username=Alice
```
<br>
Open another terminal, set a different username, and run the application again:

```sh
go run chat.go -username=Bob
```

<br>
You can hit CTRL+C or close the terminal to stop the app.
<br>
<br>
Let's be good citizens by cleaning after ourselves when we finish using it:

```sh
docker compose down
```
