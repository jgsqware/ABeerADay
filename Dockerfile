FROM ubuntu:16.10
RUN apt-get update && apt-get install -y \
    ca-certificates \
    sudo=1.8.16-0ubuntu3.2 \
    && rm -rf /var/lib/apt/lists/*
COPY a-beer-a-day /
EXPOSE 8080
ENTRYPOINT ["/a-beer-a-day"]

