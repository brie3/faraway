## Faraway's test assignment for server engineer

### Tasks
1. Design and implement “Word of Wisdom” tcp server.
1. TCP server should be protected from DDOS attacks with the [Proof of Work](https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
1. The choice of the POW algorithm should be explained.
1. After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
1. Docker file should be provided both for the server and for the client that solves the POW challenge

### Decisions
1. [go-pow](https://github.com/bwesterb/go-pow) as `Proof of Work` was chosen due to ease of implementation

### Examples
```sh
make up

Creating server ... done
Creating client ... done

docker ps

CONTAINER ID   IMAGE           COMMAND   CREATED         STATUS         PORTS                                       NAMES
4eb7d4fddb20   deploy_server   "./app"   8 seconds ago   Up 4 seconds   0.0.0.0:9095->9095/tcp, :::9095->9095/tcp   server

docker logs client
{"level":"info","body":"{\"text\":\"You can observe a lot just by watching.\",\"author\":\"Yogi Berra, type.fit\"}","time":"2024-01-09T21:29:14Z"}

make down
```