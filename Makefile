
prom:
	docker run \
    --network "host" \
    -p 9090:9090 \
    -v ./prometheus.yml \
    prom/prometheus


