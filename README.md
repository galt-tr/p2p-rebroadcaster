# P2P DHT Relay Service

A Go-based relay service that bridges messages between a private libp2p DHT network and a public DHT network. This enables seamless migration from private to public DHT networks by ensuring all messages are available on both networks during the transition period.

## Overview

This service subscribes to topics on a private DHT network (similar to Teranode's implementation) and republishes received messages to corresponding topics on a public DHT network. It's designed to facilitate the migration of nodes from private to public DHT infrastructure.

## Features

- **Dual DHT Support**: Simultaneously connects to both private and public DHT networks
- **Topic-based Message Relay**: Subscribes to and forwards messages from configurable topics
- **Message Deduplication**: Prevents forwarding duplicate messages using SHA256 hashing
- **Retry Logic**: Automatic retry for failed message deliveries
- **Rate Limiting**: Optional rate limiting to control message flow
- **Monitoring**: Built-in health checks and Prometheus metrics
- **Docker Support**: Ready-to-deploy Docker containers

## Supported Topics

The relay service handles the following message types from Teranode:

- `blocks` - Block propagation messages
- `subtrees` - Subtree validation data
- `handshake` - P2P handshake messages
- `node_status` - Node status updates
- `mining-on` - Mining status messages
- `rejected-tx` - Rejected transaction notifications

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (for containerized deployment)
- Access to both private and public DHT bootstrap peers

### Configuration

1. Copy and modify the configuration file:

```bash
cp config.yaml config.local.yaml
```

2. Update the configuration with your network details:

```yaml
private_dht:
  shared_key: "your-private-network-key"
  bootstrap_peers:
    - "/ip4/10.0.0.1/tcp/4001/p2p/YourPrivatePeer1"
    - "/ip4/10.0.0.2/tcp/4001/p2p/YourPrivatePeer2"

public_dht:
  bootstrap_peers:
    - "/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN"
```

### Building

```bash
# Install dependencies
go mod download

# Build the binary
go build -o relay ./cmd/relay

# Or use Docker
docker build -t p2p-relay:latest .
```

### Running

#### Local Development

```bash
# Run with default config
./relay

# Run with custom config
./relay -config config.local.yaml
```

#### Docker Deployment

```bash
# Start with docker-compose
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the service
docker-compose down
```

## Architecture

The relay service consists of three main components:

### 1. Private DHT Node (`relay/private_node.go`)
- Connects to the private DHT network using a shared key
- Subscribes to configured topics
- Forwards received messages to the relay service

### 2. Public DHT Node (`relay/public_node.go`)
- Connects to the public DHT network
- Publishes messages to configured topics
- Manages connections to public bootstrap peers

### 3. Relay Service (`relay/relay.go`)
- Orchestrates message flow between nodes
- Handles deduplication and retry logic
- Provides monitoring and statistics

## Monitoring

### Health Endpoints

- `http://localhost:8081/health` - Overall service health
- `http://localhost:8081/ready` - Readiness check (verifies connections)

### Metrics Endpoint

- `http://localhost:8080/metrics` - Prometheus-compatible metrics

Available metrics:
- `relay_messages_total` - Total messages relayed
- `relay_bytes_total` - Total bytes transferred
- `relay_errors_total` - Total relay errors
- `relay_connections_active` - Active peer connections

## Configuration Options

### Service Configuration

- `log_level` - Logging verbosity (debug, info, warn, error)
- `metrics_port` - Port for metrics endpoint
- `health_port` - Port for health checks

### DHT Configuration

- `shared_key` - Private network key (private DHT only)
- `topic_prefix` - Prefix for topic names
- `bootstrap_peers` - List of bootstrap peer addresses
- `listen_port` - P2P listening port
- `topics` - List of topics to relay

### Relay Configuration

- `buffer_size` - Message buffer size
- `dedup_cache_size` - Deduplication cache size
- `dedup_cache_ttl` - Cache entry TTL
- `max_retries` - Maximum retry attempts
- `retry_delay` - Delay between retries
- `rate_limit` - Rate limiting configuration
- `filter` - Peer filtering options

## Development

### Project Structure

```
p2p-rebroadcaster/
├── cmd/relay/          # Main application entry point
├── relay/              # Core relay service implementation
│   ├── private_node.go # Private DHT node
│   ├── public_node.go  # Public DHT node
│   └── relay.go        # Relay orchestration
├── types/              # Type definitions
│   └── messages.go     # Message and config types
├── config.yaml         # Default configuration
├── docker-compose.yml  # Docker Compose configuration
└── Dockerfile          # Container build file
```

### Testing

```bash
# Run unit tests
go test ./...

# Run with race detection
go test -race ./...

# Run with coverage
go test -cover ./...
```

## Troubleshooting

### No Messages Being Relayed

1. Check bootstrap peer connectivity:
   - Verify private DHT peers are reachable
   - Ensure public DHT bootstrap peers are accessible

2. Verify topic configuration:
   - Ensure topic names match between source and relay config
   - Check topic prefixes are correct

3. Review logs for errors:
   ```bash
   docker-compose logs -f | grep ERROR
   ```

### High Memory Usage

- Reduce `dedup_cache_size` in configuration
- Lower `buffer_size` to limit queued messages
- Enable rate limiting to control message flow

### Connection Issues

- Verify firewall rules allow P2P ports (4001, 4002)
- Check NAT configuration if behind a router
- Ensure private network key matches other nodes

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.