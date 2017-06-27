FROM vulhub/nginx:heartbleed

COPY a-beer-a-day /
EXPOSE 8080
ENTRYPOINT ["/a-beer-a-day"]

