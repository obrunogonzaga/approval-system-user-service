#!/bin/bash
# scripts/setup.sh
echo "Setting up development environment..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go first."
    exit 1
fi

# Install development tools
go install github.com/golang/mock/mockgen@latest
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Initialize go modules
go mod tidy

echo "Setup complete!"

#!/bin/bash
# scripts/deploy.sh
set -e

ENV=$1
if [ -z "$ENV" ]; then
    echo "Usage: ./deploy.sh <environment>"
    exit 1
fi

echo "Deploying to $ENV environment..."

# Build the application
echo "Building application..."
go build -o bin/app cmd/api/main.go

# Run database migrations
echo "Running database migrations..."
go run cmd/migrate/main.go

# Deploy based on environment
case $ENV in
    "prod")
        echo "Deploying to production..."
        # Add production deployment steps
        ;;
    "staging")
        echo "Deploying to staging..."
        # Add staging deployment steps
        ;;
    *)
        echo "Unknown environment: $ENV"
        exit 1
        ;;
esac

echo "Deployment complete!"

#!/bin/bash
# scripts/test.sh
echo "Running tests..."

# Run unit tests with coverage
go test -v -cover ./...

# Run integration tests
go test -v -tags=integration ./test/integration/...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

echo "Tests complete!"

#!/bin/bash
# scripts/db/migrate.sh
echo "Running database migrations..."

# Check if database exists
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c '\q' 2>/dev/null
if [ $? -ne 0 ]; then
    echo "Creating database..."
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -c "CREATE DATABASE $DB_NAME;"
fi

# Run migrations
go run cmd/migrate/main.go

echo "Migrations complete!"

#!/bin/bash
# scripts/docker/build-and-push.sh
VERSION=$1
if [ -z "$VERSION" ]; then
    VERSION=latest
fi

echo "Building and pushing Docker image version: $VERSION"

# Build the image
docker build -t yourregistry.com/yourapp:$VERSION .

# Push to registry
docker push yourregistry.com/yourapp:$VERSION

echo "Build and push complete!"

#!/bin/bash
# scripts/generate-mocks.sh
echo "Generating mocks..."

# Generate mocks for interfaces
mockgen -source=internal/domain/user.go -destination=internal/mocks/user_mock.go
mockgen -source=internal/usecase/user_usecase.go -destination=internal/mocks/user_usecase_mock.go

echo "Mock generation complete!"

#!/bin/bash
# scripts/local-dev.sh
echo "Starting local development environment..."

# Start Docker Compose
docker-compose up -d

# Watch for file changes and rebuild
go run github.com/cosmtrek/air

echo "Local development environment stopped."