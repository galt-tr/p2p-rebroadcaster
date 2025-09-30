package types

import (
	"time"
)

// MessageType represents the type of P2P message
type MessageType string

const (
	TypeBlock      MessageType = "block"
	TypeSubtree    MessageType = "subtree"
	TypeHandshake  MessageType = "handshake"
	TypeNodeStatus MessageType = "node_status"
	TypeMiningOn   MessageType = "mining-on"
	TypeRejectedTx MessageType = "rejected-tx"
)

// BaseMessage represents a generic P2P message that can be relayed
type BaseMessage struct {
	Type      string    `json:"type"`
	PeerID    string    `json:"peer_id"`
	Timestamp time.Time `json:"timestamp"`
	Data      []byte    `json:"data,omitempty"`
}

// HandshakeMessage represents a P2P handshake message
type HandshakeMessage struct {
	Type        string `json:"type"`
	PeerID      string `json:"peer_id"`
	BestHeight  int32  `json:"best_height"`
	BestHash    string `json:"best_hash"`
	DataHubURL  string `json:"data_hub_url,omitempty"`
	TopicPrefix string `json:"topic_prefix,omitempty"`
	UserAgent   string `json:"user_agent,omitempty"`
	Services    uint64 `json:"services,omitempty"`
}

// NodeStatusMessage represents a node status update
type NodeStatusMessage struct {
	Type          string    `json:"type"`
	PeerID        string    `json:"peer_id"`
	IsSelf        bool      `json:"is_self"`
	Version       string    `json:"version,omitempty"`
	BestHeight    int32     `json:"best_height"`
	BestBlockHash string    `json:"best_block_hash"`
	Timestamp     time.Time `json:"timestamp"`
}

// BlockMessage represents a block propagation message
type BlockMessage struct {
	Type      string `json:"type"`
	BlockHash string `json:"block_hash"`
	Height    int32  `json:"height"`
	Size      int    `json:"size"`
	Data      []byte `json:"data,omitempty"`
}

// SubtreeMessage represents a subtree validation message
type SubtreeMessage struct {
	Type        string `json:"type"`
	SubtreeID   string `json:"subtree_id"`
	BlockHash   string `json:"block_hash"`
	MerkleRoot  string `json:"merkle_root"`
	Data        []byte `json:"data,omitempty"`
}

// RelayStats tracks relay performance metrics
type RelayStats struct {
	MessagesRelayed        uint64
	BytesTransferred       uint64
	ErrorCount             uint64
	LastMessageTime        time.Time
	ConnectionsActive      int
	StartTime              time.Time
	// Bidirectional stats
	ReverseMessagesRelayed uint64
	ReverseBytesTransferred uint64
	ReverseErrorCount      uint64
	ReverseLastMessageTime time.Time
}

// Config represents the complete configuration for the relay service
type Config struct {
	Service     ServiceConfig     `yaml:"service"`
	PrivateDHT  DHTConfig        `yaml:"private_dht"`
	PublicDHT   DHTConfig        `yaml:"public_dht"`
	Relay       RelayConfig      `yaml:"relay"`
	Monitoring  MonitoringConfig `yaml:"monitoring"`
}

// ServiceConfig contains general service settings
type ServiceConfig struct {
	LogLevel    string `yaml:"log_level"`
	MetricsPort int    `yaml:"metrics_port"`
	HealthPort  int    `yaml:"health_port"`
	KeyDir      string `yaml:"key_dir"`  // Directory to store persistent keys
}

// DHTConfig contains configuration for a DHT node
type DHTConfig struct {
	SharedKey      string   `yaml:"shared_key,omitempty"`
	TopicPrefix    string   `yaml:"topic_prefix"`
	BootstrapPeers []string `yaml:"bootstrap_peers"`
	ListenPort     int      `yaml:"listen_port"`
	Topics         []string `yaml:"topics"`
	DHTProtocolID  string   `yaml:"dht_protocol_id,omitempty"`
}

// RelayConfig contains relay-specific settings
type RelayConfig struct {
	BufferSize          int               `yaml:"buffer_size"`
	DedupCacheSize      int               `yaml:"dedup_cache_size"`
	DedupCacheTTL       time.Duration     `yaml:"dedup_cache_ttl"`
	MaxRetries          int               `yaml:"max_retries"`
	RetryDelay          time.Duration     `yaml:"retry_delay"`
	RateLimit           RateLimitConfig   `yaml:"rate_limit"`
	Filter              FilterConfig      `yaml:"filter"`
	Bidirectional       bool              `yaml:"bidirectional"`
	ReverseBufferSize   int               `yaml:"reverse_buffer_size,omitempty"`
}

// RateLimitConfig contains rate limiting settings
type RateLimitConfig struct {
	Enabled              bool `yaml:"enabled"`
	MaxMessagesPerSecond int  `yaml:"max_messages_per_second"`
}

// FilterConfig contains message filtering settings
type FilterConfig struct {
	Enabled       bool     `yaml:"enabled"`
	AllowedPeers  []string `yaml:"allowed_peers"`
	BlockedPeers  []string `yaml:"blocked_peers"`
}

// MonitoringConfig contains monitoring and alerting settings
type MonitoringConfig struct {
	MetricsEnabled bool          `yaml:"metrics_enabled"`
	StatsInterval  time.Duration `yaml:"stats_interval"`
	Alerts         AlertConfig   `yaml:"alerts"`
}

// AlertConfig contains alert threshold settings
type AlertConfig struct {
	MaxLagMessages int     `yaml:"max_lag_messages"`
	MaxErrorRate   float64 `yaml:"max_error_rate"`
}