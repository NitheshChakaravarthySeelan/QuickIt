# To run the build with the test 
docker compose -f .github/workflows/docker-compose.yml up --build --abort-on-container-exit --exit-code-from test


# To remove the contaners, volumes, networks, images for this project
docker compose -f .github/workflows/docker-compose.yml down --volumes --remove-orphans --rmi all
