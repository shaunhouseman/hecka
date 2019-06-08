# HECka

A simple binary for piping to Splunk.

## Arguments

```bash
Usage of hecka:
  -h string
        host IP, Example -h="192.168.0.33" (default "127.0.0.1")
  -i string
        Sets index, Example -p="main" (default "main")
  -p string
        Sets port, Example -p="443" (default "8088")
  -s string
        Sets source, Example -p="hax" (default "lilbigdata")
  -skip int
        Skips specified lines, Example -skip=4
  -st string
        Manually sets sourcetype, Example -c="hax" (default "mahdata")
  -t string
        HEC Token, Example -t="7129b26a-c177-4705-aa5d-0eavf3b09cdf" (default "lilbigdata")
  -v string
        Turns on verbose mode, Example -v="true" (default "false")
```

## Setup

```bash
 go build hecka.go
 chmod +x hecka
 mv hecka /usr/local/bin/hecka
 ```

# lil Big Data

Simple Splunk container.

## Setup

You'll need Docker

```docker build -t lbd .```

```docker run -d -p 8000:8000 -p 8088:8088 lbd:latest```

Access the container here:

http://127.0.0.1:8000
Username: admin
Password: lilbigdata

* Now you can pipe to splunk with a command like: ```ping localhost | hecka```

## Developement

* Test HEC input
 
```curl -k https://127.0.0.1:8088/services/collector -H 'Authorization: Splunk lilbigdata' -d '{"sourcetype": "mysourcetype", "event":"Hello, World!"}'```
