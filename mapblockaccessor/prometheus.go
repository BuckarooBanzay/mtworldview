package mapblockaccessor

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	getCacheHitCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "dbcache_hit_count",
			Help: "Count of db cache hits",
		},
	)
	getCacheMissCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "dbcache_miss_count",
			Help: "Count of db cache miss",
		},
	)
	cacheBlockCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "dbcache_block_count",
			Help: "Count of db blocks inserted",
		},
	)
	cacheBlocks = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "dbcache_blocks",
			Help: "Block count currently in the cache",
		},
	)
	dbGetDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "db_get_duration",
		Help:    "Histogram for db mapblock get durations",
		Buckets: prometheus.LinearBuckets(0.001, 0.005, 10),
	})
)

func init() {
	prometheus.MustRegister(getCacheHitCount)
	prometheus.MustRegister(getCacheMissCount)

	prometheus.MustRegister(cacheBlockCount)
	prometheus.MustRegister(cacheBlocks)

	prometheus.MustRegister(dbGetDuration)
}
