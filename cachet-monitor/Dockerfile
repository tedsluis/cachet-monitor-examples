FROM registry.access.redhat.com/openshift3/ose

MAINTAINER ted.sluis@gmail.com

RUN mkdir /var/lib/cachet-monitor; \
    mkdir /var/lib/cachet-monitor/config

ADD cachet-monitor-linux-amd64 /var/lib/cachet-monitor/
ADD cachet-monitor-config.yml  /var/lib/cachet-monitor/config/

RUN chmod -R 777  /var/lib/cachet-monitor

ENTRYPOINT ["/var/lib/cachet-monitor/cachet-monitor-linux-amd64","-c","/var/lib/cachet-monitor/config/cachet-monitor-config.yml"]
