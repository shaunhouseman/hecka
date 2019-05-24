# HECka

Pipe to Splunk

## lilbigdata

Quick Splunk

## Setup

* Build image

```bash
docker build -t lbd .
```

* Run container

```bash
docker run -d -p 8000:8000 -p 8088:8088 lbd:latest
```

```bash
docker stop <container id>
```

```bash
docker start <container id>
```

* echo "test message" | main

## Developement

* Test HEC input

```bash
curl -k https://127.0.0.1:8088/services/collector -H 'Authorization: Splunk lilbigdata' -d '{"sourcetype": "mysourcetype", "event":"Hello, World!"}'
```
