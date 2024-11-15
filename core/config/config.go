package config

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/river-build/river/core/node/base"
	. "github.com/river-build/river/core/node/protocol"
)

func GetDefaultConfig() *Config {
	return &Config{
		Port: 443,
		Database: DatabaseConfig{
			StartupDelay:  2 * time.Second,
			NumPartitions: 256,
		},
		StorageType:  "postgres",
		DisableHttps: false,
		BaseChain: ChainConfig{
			// TODO: ChainId:
			BlockTimeMs: 2000,
		},
		RiverChain: ChainConfig{
			// TODO: ChainId:
			BlockTimeMs: 2000,
			TransactionPool: TransactionPoolConfig{
				TransactionTimeout:               6 * time.Second,
				GasFeeCap:                        150_000_000_000, // 150 Gwei
				MinerTipFeeReplacementPercentage: 10,
				GasFeeIncreasePercentage:         10,
			},
		},
		// TODO: ArchitectContract: ContractConfig{},
		// TODO: RegistryContract:  ContractConfig{},
		StreamReconciliation: StreamReconciliationConfig{
			WorkerPoolSize: 8,
		},
		Log: LogConfig{
			Level:   "info", // NOTE: this default is replaced by flag value
			Console: true,   // NOTE: this default is replaced by flag value
			File:    "",     // NOTE: this default is replaced by flag value
			NoColor: false,  // NOTE: this default is replaced by flag value
			Format:  "json",
		},
		Metrics: MetricsConfig{
			Enabled: true,
		},
		// TODO: Network: NetworkConfig{},
		StandByOnStart:    true,
		StandByPollPeriod: 500 * time.Millisecond,
		ShutdownTimeout:   1 * time.Second,
		History:           30 * time.Second,
		Archive: ArchiveConfig{
			PrintStatsPeriod: 10 * time.Second,
		},
		DebugEndpoints: DebugEndpointsConfig{
			Cache:                 true,
			Memory:                true,
			PProf:                 false,
			Stacks:                true,
			StacksMaxSizeKb:       5 * 1024,
			TxPool:                true,
			EnableStorageEndpoint: true,
		},
		Scrubbing: ScrubbingConfig{
			ScrubEligibleDuration: time.Hour,
		},
		RiverRegistry: RiverRegistryConfig{
			PageSize:               1000,
			MaxRetries:             100,
			MaxRetryElapsedTime:    5 * time.Minute,
			SingleCallTimeout:      30 * time.Second, // geth internal timeout is 30 seconds
			ProgressReportInterval: 10 * time.Second,
		},
	}
}

// Viper uses mapstructure module to marshal settings into config struct.
type Config struct {
	// Network
	// 0 can be used in tests to elect a free available port.
	Port int
	// DNS name of the node. Used to select interface to listen on. Can be empty.
	Address string

	DisableHttps bool // If FALSE TLSConfig must be set.
	TLSConfig    TLSConfig

	// Storage
	Database    DatabaseConfig
	StorageType string

	// Blockchain configuration
	BaseChain  ChainConfig
	RiverChain ChainConfig

	// Base chain contract configuration
	ArchitectContract ContractConfig

	// Contract configuration
	RegistryContract ContractConfig

	// Logging
	Log LogConfig

	// Metrics
	Metrics             MetricsConfig
	PerformanceTracking PerformanceTrackingConfig

	// Scrubbing
	Scrubbing ScrubbingConfig

	// Stream reconciliation
	StreamReconciliation StreamReconciliationConfig

	// Network configuration
	Network NetworkConfig

	// Go in stand-by mode on start checking if public address resolves to this node instance.
	// This allows to reduce downtime when new version of the node is deployed in the new container or VM.
	// Depending on the network routing configuration this approach may not work.
	StandByOnStart    bool
	StandByPollPeriod time.Duration

	// ShutdownTimeout is the time the node waits for the graceful shutdown of the server.
	// Then all active connections are closed and the node exits.
	// If StandByOnStart is true, it's recommended to set it to the half of DatabaseConfig.StartupDelay.
	// If set to 0, then default value is used. To disable the timeout set to 1ms or less.
	ShutdownTimeout time.Duration

	// Graffiti is returned in status and info requests.
	Graffiti string

	// Should be set if node is run in archive mode.
	Archive ArchiveConfig

	// Feature flags
	// Used to disable functionality for some testing setups.

	// Disable base chain contract usage.
	DisableBaseChain bool

	// Chains provides a map of chain IDs to their provider URLs as
	// a comma-serparated list of chainID:URL pairs.
	// It is parsed into ChainsString variable.
	Chains string `dlog:"omit" json:"-" yaml:"-"`

	// ChainsString is an another alias for Chains kept for backward compatibility.
	ChainsString string `dlog:"omit" json:"-" yaml:"-"`

	// This is comma-separated list chaidID:blockTimeDuration pairs.
	// GetDefaultBlockchainInfo() provides default values for known chains so there is no
	// need to set block time is it's in the GetDefaultBlockchainInfo().
	// I.e. 1:12s,84532:2s,6524490:2s
	ChainBlocktimes string

	ChainConfigs map[uint64]*ChainConfig `mapstructure:"-"` // This is a derived field from Chains.

	// extra xChain configuration
	EntitlementContract     ContractConfig `mapstructure:"entitlement_contract"`
	TestEntitlementContract ContractConfig `mapstructure:"test_contract"`

	// History indicates how far back xchain must look for entitlement check requests after start
	History time.Duration

	// EnableTestAPIs enables additional APIs used for testing.
	EnableTestAPIs bool

	// EnableDebugEndpoints is a legacy setting, enables all debug endpoints.
	// Per endpoint configuration is in DebugEndpoints.
	EnableDebugEndpoints bool

	DebugEndpoints DebugEndpointsConfig

	// RiverRegistry contains settings for calling registry contract on River chain.
	RiverRegistry RiverRegistryConfig
}

type TLSConfig struct {
	Cert   string // Path to certificate file or BASE64 encoded certificate
	Key    string `dlog:"omit" json:"-" yaml:"-"` // Path to key file or BASE64 encoded key. Sensitive data, omitted from logging.
	TestCA string // Path to CA certificate file or BASE64 encoded CA certificate
}

type NetworkConfig struct {
	NumRetries int
	// RequestTimeout only applies to unary requests.
	RequestTimeout time.Duration

	// If unset or <= 0, 5 seconds is used.
	HttpRequestTimeout time.Duration
}

func (nc *NetworkConfig) GetHttpRequestTimeout() time.Duration {
	if nc.HttpRequestTimeout <= 0 {
		return 5 * time.Second
	}
	return nc.HttpRequestTimeout
}

type DatabaseConfig struct {
	Url      string `dlog:"omit" json:"-" yaml:"-"` // Sensitive data, omitted from logging.
	Host     string
	Port     int
	User     string
	Password string `dlog:"omit" json:"-" yaml:"-"` // Sensitive data, omitted from logging.
	Database string
	Extra    string

	// StartupDelay is the time the node waits between taking control of the database and starting the server
	// if other nodes' records are found in the database.
	// If StandByOnStart is true, it's recommended to set it to the double of Config.ShutdownTimeout.
	// If set to 0, then default value is used. To disable the delay set to 1ms or less.
	StartupDelay time.Duration

	// IsolationLevel is the transaction isolation level to use for the database operations.
	// Allowed values: "serializable", "repeatable read", "read committed".
	// If not set or value can't be parsed, defaults to "serializable".
	// Intention is to migrate to "read committed" for performance reasons after testing is complete.
	IsolationLevel string

	// MigrateStreamCreation indicates to the database that all new streams should be allocated
	// in one of the 256 pre-allocated stream partitions for either regular or media streams instead
	// of allocating new tables per stream according to the legacy schema. If this flag is unset, a
	// node will continue to allocate new tables for each stream as it is created.
	MigrateStreamCreation bool

	// NumPartitions specifies the number of partitions to use when creating the schema for stream
	// data storage. If <= 0, a default value of 256 will be used. No more than 256 partitions is
	// supported at this time.
	NumPartitions int
}

func (c DatabaseConfig) GetUrl() string {
	if c.Host != "" {
		return fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s%s",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			c.Database,
			c.Extra,
		)
	}

	return c.Url
}

// TransactionPoolConfig specifies when it is time for a replacement transaction and its gas fee costs.
type TransactionPoolConfig struct {
	// TransactionTimeout is the duration in which a transaction must be included in the chain before it is marked
	// eligible for replacement. It is advisable to set the timeout as a multiple of the block period. If not set it
	// estimates the chains block period and sets Timeout to 3x block period.
	TransactionTimeout time.Duration

	// GasFeeCap determines for EIP-1559 transaction the maximum amount fee per gas the node operator is willing to
	// pay. If set to 0 the node will use 2 * chain.BaseFee by default. The base fee + miner tip must be below this
	// cap, if not the transaction could not be made.
	GasFeeCap int

	// MinerTipFeeReplacementPercentage is the percentage the miner tip for EIP-1559 transactions is incremented when
	// replaced. Nodes accept replacements only when the miner tip is at least 10% higher than the original transaction.
	// The node will add 1 Wei to the miner tip and therefore 10% is the least recommended value. Default is 10.
	MinerTipFeeReplacementPercentage int

	// GasFeeIncreasePercentage is the percentage by which the gas fee for legacy transaction is incremented when it is
	// replaced. Recommended is >= 10% since nodes typically only accept replacements transactions with at least 10%
	// higher gas price. The node will add 1 Wei, therefore 10% will also work. Default is 10.
	GasFeeIncreasePercentage int
}

type ChainConfig struct {
	NetworkUrl  string `dlog:"omit" json:"-" yaml:"-"` // Sensitive data, omitted from logging.
	ChainId     uint64
	BlockTimeMs uint64

	TransactionPool TransactionPoolConfig

	// TODO: these need to be removed from here
	LinkedWalletsLimit                        int
	ContractCallsTimeoutMs                    int
	PositiveEntitlementCacheSize              int
	PositiveEntitlementCacheTTLSeconds        int
	NegativeEntitlementCacheSize              int
	NegativeEntitlementCacheTTLSeconds        int
	PositiveEntitlementManagerCacheSize       int
	PositiveEntitlementManagerCacheTTLSeconds int
	NegativeEntitlementManagerCacheSize       int
	NegativeEntitlementManagerCacheTTLSeconds int
}

func (c ChainConfig) BlockTime() time.Duration {
	return time.Duration(c.BlockTimeMs) * time.Millisecond
}

type PerformanceTrackingConfig struct {
	ProfilingEnabled bool

	// If true, write trace data to one of the exporters configured below
	TracingEnabled bool
	// If set, write trace data to this jsonl file
	OtlpFile string
	// If set, send trace data to using OTLP HTTP
	// Exporter is configured with OTLP env variables as described here:
	// go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp
	OtlpEnableHttp bool
	// If set, send trace data to using OTLP gRRC
	// Exporter is configured with OTLP env variables as described here:
	// go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
	OtlpEnableGrpc bool
	// If set, connet to OTLP endpoint using http instead of https
	// Also can be configured by env var from the links above
	OtlpInsecure bool

	// If set, send trace spans to this Zipkin endpoint
	ZipkinUrl string
}

type ContractConfig struct {
	// Address of the contract
	Address common.Address
}

type ArchiveConfig struct {
	// ArchiveId is the unique identifier of the archive node. Must be set for nodes in archive mode.
	ArchiveId string

	Filter FilterConfig

	// Number of miniblocks to read at once from the remote node.
	ReadMiniblocksSize uint64

	DisablePrintStats bool
	PrintStatsPeriod  time.Duration // If 0, default to 1 minute.

	TaskQueueSize int // If 0, default to 100000.

	WorkerPoolSize int // If 0, default to 20.

	StreamsContractCallPageSize int64 // If 0, default to 5000.
}

type LogConfig struct {
	Level        string // Used for both file and console if their levels not set explicitly
	File         string // Path to log file
	FileLevel    string // If not set, use Level
	Console      bool   // Log to console if true
	ConsoleLevel string // If not set, use Level
	NoColor      bool   // If true, disable color text output to console
	Format       string // "json" or "text"

	// Intended for dev use with text logs, do not output instance attributes with each log entry,
	// drop some large messages.
	Simplify bool
}

type MetricsConfig struct {
	// Enable metrics collection, publish on /metrics endpoint on public port unless DisablePublic is set.
	Enabled bool

	// If set, do not publish /metrics on public port.
	DisablePublic bool

	// If not 0, also publish /metrics on this port.
	Port int

	// Interface to use with the port above. Usually left empty to bind to all interfaces.
	Interface string
}

type DebugEndpointsConfig struct {
	Cache           bool
	Memory          bool
	PProf           bool
	Stacks          bool
	StacksMaxSizeKb int
	TxPool          bool

	// Make storage statistics available via debug endpoints. This may involve running queries
	// on the underlying database.
	EnableStorageEndpoint bool
}

type RiverRegistryConfig struct {
	// PageSize is the number of streams to read from the contract at once using GetPaginatedStreams.
	PageSize int

	// If not 0, stop retrying failed GetPaginatedStreams calls after this number of retries.
	MaxRetries int

	// Stop retrying failed GetPaginatedStreams calls after this duration.
	MaxRetryElapsedTime time.Duration

	// Timeout for a singe call to GetPaginatedStreams.
	SingleCallTimeout time.Duration

	// ProgressReportInterval is the interval at which to report progress of the GetPaginatedStreams calls.
	ProgressReportInterval time.Duration
}

func (ac *ArchiveConfig) GetReadMiniblocksSize() uint64 {
	if ac.ReadMiniblocksSize <= 0 {
		return 100
	}
	return ac.ReadMiniblocksSize
}

func (ac *ArchiveConfig) GetPrintStatsPeriod() time.Duration {
	if ac.DisablePrintStats {
		return 0
	}
	if ac.PrintStatsPeriod <= 0 {
		return time.Minute
	}
	return ac.PrintStatsPeriod
}

func (ac *ArchiveConfig) GetTaskQueueSize() int {
	if ac.TaskQueueSize <= 0 {
		return 100000
	}
	return ac.TaskQueueSize
}

func (ac *ArchiveConfig) GetWorkerPoolSize() int {
	if ac.WorkerPoolSize <= 0 {
		return 20
	}
	return ac.WorkerPoolSize
}

func (ac *ArchiveConfig) GetStreamsContractCallPageSize() int64 {
	if ac.StreamsContractCallPageSize <= 0 {
		return 5000
	}
	return ac.StreamsContractCallPageSize
}

type ScrubbingConfig struct {
	// ScrubEligibleDuration is the minimum length of time that must pass before a stream is eligible
	// to be re-scrubbed.
	// If unset, it defaults to 1 hour.
	ScrubEligibleDuration time.Duration
}

type StreamReconciliationConfig struct {
	WorkerPoolSize int // If 0, default to 8.
}

type FilterConfig struct {
	// If set, only archive streams hosted on the nodes with the specified addresses.
	Nodes []string

	// If set, only archive stream if Nodes list contains first hosting node for the stream.
	// This may be used to archive only once copy of replicated stream
	// if multiple archival nodes are used in conjunction.
	FirstOnly bool

	// If set, partition all stream names using hash into specified number of shards and
	// archive only listed shards.
	NumShards uint64
	Shards    []uint64
}

func (c *Config) GetGraffiti() string {
	if c.Graffiti == "" {
		return "River Node welcomes you!"
	}
	return c.Graffiti
}

func (c *Config) GetEntitlementContractAddress() common.Address {
	return c.EntitlementContract.Address
}

func (c *Config) GetWalletLinkContractAddress() common.Address {
	return c.ArchitectContract.Address
}

func (c *Config) GetTestEntitlementContractAddress() common.Address {
	return c.TestEntitlementContract.Address
}

func (c *Config) Init() error {
	return c.parseChains()
}

// Return the schema to use for accessing the node.
func (c *Config) UrlSchema() string {
	s := "https"
	if c != nil && c.DisableHttps {
		s = "http"
	}
	return s
}

func parseBlockchainDurations(str string, result map[uint64]BlockchainInfo) error {
	pairs := strings.Split(str, ",")
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			chainID, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 64)
			if err != nil {
				return WrapRiverError(Err_BAD_CONFIG, err).Message("Failed to parse chain Id").Tag("value", str)
			}
			duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
			if err != nil {
				return WrapRiverError(Err_BAD_CONFIG, err).Message("Failed to parse block time").Tag("value", str)
			}
			result[chainID] = BlockchainInfo{
				ChainId:   chainID,
				Blocktime: duration,
				Name:      parts[0],
			}
		} else {
			return RiverError(Err_BAD_CONFIG, "Failed to parse chain blocktimes").Tag("value", str)
		}
	}
	return nil
}

func (c *Config) parseChains() error {
	defaultChainInfo := GetDefaultBlockchainInfo()
	err := parseBlockchainDurations(c.ChainBlocktimes, defaultChainInfo)
	if err != nil {
		return err
	}

	// If Chains is empty, fallback to ChainsString.
	if c.Chains == "" {
		c.Chains = strings.TrimSpace(c.ChainsString)
	}
	chains := strings.TrimSpace(c.Chains)

	chainConfigs := make(map[uint64]*ChainConfig)
	chainPairs := strings.Split(chains, ",")
	for _, pair := range chainPairs {
		if pair == "" {
			continue
		}
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			chainID, err := strconv.ParseUint(strings.TrimSpace(parts[0]), 10, 64)
			if err != nil {
				return WrapRiverError(Err_BAD_CONFIG, err).Message("Failed to pase chain Id").Tag("chainId", parts[0])
			}

			info, ok := defaultChainInfo[chainID]
			if !ok {
				return RiverError(Err_BAD_CONFIG, "Chain blocktime not set").Tag("chainId", chainID)
			}

			chainConfigs[chainID] = &ChainConfig{
				NetworkUrl:  strings.TrimSpace(parts[1]),
				ChainId:     chainID,
				BlockTimeMs: uint64(info.Blocktime / time.Millisecond),
			}
		} else {
			return RiverError(Err_BAD_CONFIG, "Failed to parse chain config").Tag("value", pair)
		}
	}
	c.ChainConfigs = chainConfigs

	return nil
}

type confifCtxKeyType struct{}

var configCtxKey = confifCtxKeyType{}

func CtxWithConfig(ctx context.Context, c *Config) context.Context {
	return context.WithValue(ctx, configCtxKey, c)
}

func FromCtx(ctx context.Context) *Config {
	if c, ok := ctx.Value(configCtxKey).(*Config); ok {
		return c
	}
	return nil
}

func UseDetailedLog(ctx context.Context) bool {
	c := FromCtx(ctx)
	if c != nil {
		return !c.Log.Simplify
	} else {
		return true
	}
}
