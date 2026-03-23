# redis, k6, grafana with golang

## Document

- Redis https://redis.io/docs/latest/operate/oss_and_stack/management/config/
- K6 https://k6.io
- Influx https://influxdata.com
- Grafana https://grafana.com

## Docker Image
- redis https://hub.docker.com/_/redis
- loadimpact/k6 https://hub.docker.com/r/loadimpact/k6
- influxdb https://hub.docker.com/_/influxdb
- mariadb https://hub.docker.com/_/mariadb

## 🔴 Redis

### What it is
- In-memory key-value store (super fast)
- Used for:
  - caching
  - session storage
  - rate limiting
  - queues

### Key concepts
- Data types: string, list, set, hash, zset
- Runs in RAM → fast but not primary DB
- Persistence options:
  - RDB (snapshot)
  - AOF (append-only log)

### Important config (your redis.conf)
```
- bind 0.0.0.0 → allow external connection (Docker)
- protected-mode no → required in container sometimes

- appendonly yes → don’t lose data

What it means
Enables AOF (Append Only File) persistence
How it works

Every write operation is logged:

SET key value

→ saved to a file

Why it matters
Redis is in-memory → data can be lost
AOF makes it durable

- SAVE "" → disable snapshots

What it means
Disables RDB snapshotting

Normally Redis does:

SAVE 900 1
SAVE 300 10
SAVE 60 10000

→ Save snapshot based on changes/time

But:

SAVE ""

= disable ALL snapshots

```
### Useful commands
```
# connect inside container
docker exec -it redis redis-cli
OR
docker exec -it redis redis-cli -a <password>

# basic usage
SET key value
GET key
KEYS *
FLUSHALL
```

## ⚡ k6

### What it is
- Load testing tool (like JMeter but modern + JS-based)
- You write test scripts in JavaScript

### What it measures
- http_req_duration → total request time 
- http_req_waiting → server processing time
- http_req_connecting → TCP connect time
- vus → virtual users

### Example script
```
import http from 'k6/http'

export let options = {
    vus: 5,
    duration: '5s'
}

export default function () { 
    http.get('http://host.docker.internal:8000/hello')
}
```
### Run with Docker
```
docker run --rm -i loadimpact/k6 run - <script.js
```
or with your compose:
```
docker compose run --rm k6 run /scripts/test.js
```
### Output to InfluxDB

I set in docker compose environment :
```
K6_OUT=influxdb=http://influxdb:8086/k6
```
That’s how metrics go → InfluxDB

## 📊 InfluxDB

### What it is
- Database for metrics over time
- Perfect for k6 results
### Key concepts
- Measurement = table (e.g. http_req_duration)
- Tags = indexed fields (fast filtering)
- Fields = actual values
- Time = main dimension
### Check data
```
docker exec -it influxdb influx

# inside CLI
USE k6
SHOW MEASUREMENTS
SELECT * FROM http_req_duration LIMIT 10
```

## 🐬 MariaDB

### What it is
- Traditional SQL database (like MySQL)
### Use in Go
- Store persistent data (users, products)
Access
```
docker compose exec -it mariadb mariadb -u root -p
```
Then:
```
SHOW DATABASES;
USE infinitas;
SHOW TABLES;
SELECT * FROM products;
```

## 🐳 Docker Commands You Should Know
### Start everything
```
docker compose up -d
```
### Stop
```
docker compose down
```
### View logs
```
docker compose logs -f k6
docker compose logs -f influxdb
```
### Run k6 manually
```
docker compose run --rm k6 run /scripts/test.js
```
