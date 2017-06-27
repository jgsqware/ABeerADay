FROM ubuntu:16.10
RUN apt-get update && apt-get install -y \
    ca-certificates \
    juju-core=1.25.6-0ubuntu2.16.10.2 \
    && rm -rf /var/lib/apt/lists/*
COPY a-beer-a-day /
EXPOSE 8080
ENTRYPOINT ["/a-beer-a-day"]

