# This repo is created for final project of CS157C taught by Prof. Ching-seh Wu in San Jose State University.

Project Name:
Secret Box

Team Members: 
1. Yunlin Xie
2. Shuang Pan

### Dockerize Golang APIs and Cassaandra

Make sure you have docker installed and `make` command on your laptop.
Run the following command from `$HOME/go/src/CS157C-team8`
```
make docker
```
Then run `docker ps`, you will see something like 
```
CONTAINER ID   IMAGE              COMMAND                  CREATED          STATUS          PORTS                                                                          NAMES
47bd82b19097   secretbox-api:v1   "./CS157C-TEAM8"         3 seconds ago    Up 3 seconds    0.0.0.0:3000->3000/tcp, :::3000->3000/tcp                                      relaxed_blackwell
74655547da58   cassandra          "docker-entrypoint.s…"   28 minutes ago   Up 34 seconds   7000-7001/tcp, 7199/tcp, 9160/tcp, 0.0.0.0:9042->9042/tcp, :::9042->9042/tcp   cassandra
```
If you are not able to see `secretbox-api:v1`, run the following command and run `docker ps` again. You will see that.
```
docker run -p 3000:3000 -d -e CASSANDRA_URL=cassandra:9042 --link=cassandra secretbox-api:v1
```
If it still doesn't work. Run the following commands one by one.
```
1. docker run -p 9042:9042 -d --name cassandra cassandra
2. docker restart cassandra
3. docker build -t secretbox-api:v1 .
4. docker run -p 3000:3000 -d -e CASSANDRA_URL=cassandra:9042 --link=cassandra secretbox-api:v1
```

### Try APIS from Postman
