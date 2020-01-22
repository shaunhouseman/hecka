# Hecka

A simple binary for piping to Splunk and a simple docker container to recieve them.

## [Download](https://github.com/shaunhouseman/hecka/releases)

### Arguments

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
  -ssl bool
        enables ssl verify, Example -ssl (default false)
  -t string
        HEC Token, Example -t="7129b26a-c177-4705-aa5d-0eavf3b09cdf" (default "lilbigdata")
  -v bool
        Turns on verbose mode, Example -v (default false)
```

### Hecka Setup

```bash
 go build hecka.go
 chmod +x hecka
 mv hecka /usr/local/bin/hecka
 ```

### Usage

Basic Example

```bash
ping google.com | hecka
```

Splunk Cloud Example
```bash
ping google.com | hecka -ssl -h="http-inputs-deployment.splunkcloud.com" -p="443" -s="ping" -skip=1 -t="HECTOKENGOESHERE"
```

### Lil Big Data Setup

You'll need Docker of course. Just run the below commands.

```bash
docker build -t lbd .
docker run -d -p 127.0.0.1:8000:8000 -p 8088:8088 lbd:latest
```

Access the container here:

[http://127.0.0.1:8000](http://127.0.0.1:8000)

```bash 
Username: admin
Password: heckadata
```

* Now you can pipe to splunk with a command like: ```ping localhost | hecka```

* View [live results](http://127.0.0.1:8000/en-US/app/search/search?q=search%20index%3Dmain&display.page.search.mode=smart&dispatch.sample_ratio=1&workload_pool=&earliest=rt-5m&latest=rt)

### Developement

* Test HEC input
```curl -k https://127.0.0.1:8088/services/collector -H 'Authorization: Splunk lilbigdata' -d '{"sourcetype": "mysourcetype", "event":"Hello, World!"}'```
