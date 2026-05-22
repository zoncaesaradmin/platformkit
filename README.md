# Shared Module

This module provides common functionality shared between platform repos, with a focus on message bus implementations, logging, datastore integrations, and small reusable helper packages.

## Core Packages

### Message Bus (`messagebus/`)
- **Dual Implementation**: 
  - Local development: File-based message bus (`localbus.go`)
  - Production: Kafka-based message bus (`kafkabus.go`)
- **Common Interfaces**: Producer/Consumer patterns in `mbinterfaces.go`
- **Build Tags**: Use `local` tag for development mode
- **Offset Management**: Sequential message processing support
- **Error Handling**: Transient vs. fatal error distinction

### Logging (`logging/`)
- **Structured Logging**: JSON-formatted logs via zerolog
- **Log Levels**: DEBUG, INFO, WARN, ERROR support
- **Context Support**: Trace ID and metadata handling
- **File/Console Output**: Configurable output targets
- **Performance Optimized**: Zero-allocation logging

### Config Helpers (`configutil/`)
- **Configuration Files**: Resolve config locations and load YAML maps

### Environment Helpers (`envutil/`)
- **Environment Access**: Typed helpers for string, int, and bool env values

### Filesystem Helpers (`fsutil/`)
- **Filesystem Operations**: Safe directory setup utilities

### Time Helpers (`timeutil/`)
- **Timestamp Formatting**: UTC millisecond timestamp helpers

## Usage

### Importing
```go
import (
    "corekit/messagebus"
    "corekit/logging"
    "corekit/configutil"
)
```

### Message Bus Example
```go
// Local development (file-based)
producer := messagebus.NewProducer(configMap, "local-producer")
consumer := messagebus.NewConsumer(configMap, "local-consumer")

// Production (Kafka)
producer := messagebus.NewProducer(kafkaConfig, "kafka-producer")
consumer := messagebus.NewConsumer(kafkaConfig, "kafka-consumer")
```

### Logging Example
```go
logger := logging.NewLogger(logging.Config{
    Level:      "info",
    Format:     "json",
    Output:     "service.log",
    ServiceName: "cratos",
})

logger.Info("Processing message", "msgId", msg.ID)
```

## Configuration

### Message Bus Settings
- Local mode uses `/tmp/cratos-messagebus/`
- Production mode requires Kafka configuration:
  ```yaml
  kafka:
    bootstrap.servers: "localhost:9092"
    group.id: "cratos-group"
    client.id: "cratos-client"
    security.protocol: "PLAINTEXT"
  ```

### Logging Configuration
```yaml
logging:
  level: "info"
  format: "json"
  filename: "service.log"
  service_name: "cratos"
```

## Development

### Build Tags
- Use `-tags local` for local development
- Omit tags for production build

### Testing
```bash
# Run all tests
make test

# Generate coverage report
make coverage-html

# Enforce coverage (85% threshold)
make coverage-enforce
```

### Quality Standards
- **Coverage Target**: 85% minimum
- **Documentation**: All exported functions documented
- **Error Handling**: All errors properly handled
- **Thread Safety**: Concurrent access supported
- **Performance**: Zero-allocation where possible

## Package Layout

- Keep public, reusable packages at the repository root.
- Move implementation-only helpers into `internal/` when they should not be imported by other repos.
- Prefer focused helper packages like `configutil`, `envutil`, `fsutil`, and `timeutil` over growing a broad `utils` package.
- The existing `utils` package remains as a compatibility layer and should not be the default choice for new code.

## Directory Structure
```
corekit/
├── configstore/         # Config persistence interfaces and implementations
├── configutil/          # Config file helpers
├── ctxutil/             # Context metadata helpers
├── datastore/           # OpenSearch and datastore abstractions
├── envutil/             # Typed environment helpers
├── fsutil/              # Filesystem helpers
├── internal/            # Non-public implementation helpers
├── logging/             # Logging framework
├── messagebus/          # Message bus implementations
├── timeutil/            # Time formatting helpers
└── utils/               # Backward-compatible helper facade
```
