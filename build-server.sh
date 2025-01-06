SERVICES=("server-api" "server-grpc")

echo "Step 1: Stopping and removing containers for the services: ${SERVICES[*]}..."
for SERVICE in "${SERVICES[@]}"; do
  echo "Stopping containers for service: $SERVICE..."
  docker ps -a --filter "ancestor=${SERVICE}:latest" -q | xargs -r docker stop
  
  echo "Removing containers for service: $SERVICE..."
  docker ps -a --filter "ancestor=${SERVICE}:latest" -q | xargs -r docker rm
done

echo "Step 2: Removing old Docker images for the services: ${SERVICES[*]}..."
for SERVICE in "${SERVICES[@]}"; do
  echo "Removing image for service: $SERVICE..."
  docker images "${SERVICE}:latest" -q | xargs -r docker rmi
done

echo "Step 3: Rebuilding Docker images for all services in docker-compose.yml..."
docker-compose build --no-cache

echo "Build process completed. New images for ${SERVICES[*]} have been created successfully."
docker-compose up -d

echo "Cleaning up unused Docker images"
docker image prune -a -f

echo "Cleanup completed."