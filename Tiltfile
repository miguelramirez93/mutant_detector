docker_compose(['./docker-compose.yml'])
docker_build('local/mutant_detector', '.', dockerfile = 'Dockerfile',
    live_update=[
        sync('.', '/go/src/github.com/miguelramirez93/mutant_detector'),
    ],
    target="dev"
)