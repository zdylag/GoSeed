FROM scratch
COPY TestGoRepo /
ENTRYPOINT ["/TestGoRepo"]
