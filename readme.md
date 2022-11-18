How to run :

- go to each service

```
cd microservices/{service_name}
```

- copy .env.example to .env and then fill with the correct credentials

```
cp .env.example .env
```

- go to misc

```
cd misc
```

- run docker compose

```
docker-compose up -d --build
```

Kafka UI :
[Kafka UI](http://localhost:8080)

Import postman collection from [Postman Collection](jojonomic-test.postman_collection.json)
