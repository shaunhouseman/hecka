FROM splunk/splunk:7.2

EXPOSE 8000:8000
EXPOSE 8088:8088
ENV SPLUNK_START_ARGS=--accept-license
ENV SPLUNK_PASSWORD=hecka
ENV SPLUNK_HEC_TOKEN=lilbigdata
