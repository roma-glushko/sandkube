docker_build(
    'sandkube/urlshortener:dev',
    context='services/urlshortener',
    dockerfile='services/urlshortener/Dockerfile',
    live_update=[
        sync('./services/urlshortener', '/service/'),
    ],
)

k8s_yaml('charts/namespace.yaml')
k8s_yaml('services/urlshortener/deployment.yaml')

k8s_resource(
    'sandkube-urlshortener',
    labels=['services'],
    port_forwards="8000:8000",
)
