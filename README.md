# Go Project Template

A modern template for Go projects following clean architecture and design pattern best practices. This template provides a robust and scalable foundation for Go application development.

## 🚀 Features

- ✨ Clean and modular architecture
- 🔒 Clear separation of public and private code
- 📝 Ready-to-use logging and metrics setup
- 🔄 CI/CD configuration
- 🧪 Complete testing structure
- 📚 Automatic Swagger documentation
- 🐳 Docker containerization
- 🛠️ Development-friendly Makefile

## 📁 Project Structure

```
.
├── cmd/                    # Application executables
├── internal/              # Private application code
├── pkg/                   # Public libraries
├── configs/              # Configuration files
├── scripts/              # Automation scripts
├── test/                 # Integration tests
└── api/                  # API documentation
```

## 🔧 Prerequisites

- Go 1.21 or higher
- Docker (optional)
- Make (optional)

## 🚀 Getting Started

1. Clone this template:
```bash
git clone https://github.com/obrunogonzaga/go-template.git new-project
cd new-project
```

2. Initialize Go module:
```bash
go mod init github.com/your-username/new-project
go mod tidy
```

3. Run the application:
```bash
make run
```

## 📚 Available Commands

```bash
make build      # Compile the project
make test       # Run tests
make lint       # Run linter
make swagger    # Generate Swagger documentation
make docker     # Build Docker image
```

## 📝 Development Guidelines

### Package Structure
- `cmd/`: Contains application entry points
- `internal/`: Application-specific code
- `pkg/`: Libraries that can be used by other projects
- `test/`: Integration and end-to-end tests

### Code Standards
- Use small, focused interfaces
- Follow SOLID principles
- Keep functions short and single-purpose
- Document public APIs

## 🧪 Testing

The project uses the following testing tools:
- `testing`: Standard library for unit tests
- `testify`: Assertions and mocks
- `gomock`: Mock generation
- `httptest`: HTTP endpoint testing

## 🔒 Security

- All dependencies are verified with `go mod verify`
- Static code analysis with `golangci-lint`
- Vulnerability scanning with `gosec`

## 📦 Deployment

The project includes configurations for:
- Docker
- Kubernetes (manifests in `./deploy`)
- CI/CD (GitHub Actions)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## ✨ Best Practices Implemented

- Dependency injection
- Interface-based design
- Configuration management
- Error handling
- Middleware support
- Graceful shutdown
- Structured logging
- Metrics and monitoring
- Health checks
- Rate limiting

## 🛠️ Project Setup

### Environment Variables
```env
APP_ENV=development
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
```

### Configuration
The project uses a hierarchical configuration system:
1. Default values
2. Configuration files
3. Environment variables
4. Command line flags

### Error Handling
- Custom error types
- Error wrapping
- Structured error responses
- Centralized error handling

### Logging
- Structured logging with levels
- Request ID tracking
- Performance metrics
- Audit logging

## 📫 Contact

Bruno Gonzaga Santos - [@brunogsantos](https://www.linkedin.com/in/brunogsantos/) - brunog86@gmail.com

Project Link: [https://github.com/obrunogonzaga/go-template](https://github.com/your-username/go-template)

## 🌟 Acknowledgments

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Twelve-Factor App](https://12factor.net/)
