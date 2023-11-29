package metrics

import "github.com/giovannicuriel/gocache/lib/v4/codec"

// MetricsInterface represents the metrics interface for all available providers
type MetricsInterface interface {
	RecordFromCodec(codec codec.CodecInterface)
}
