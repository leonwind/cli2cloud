FROM envoyproxy/envoy:v1.20-latest

COPY envoy.yaml /etc/envoy/envoy.yaml

EXPOSE 8000
RUN chmod go+r /etc/envoy/envoy.yaml
