package pubsub

// Config defines the configuration settings for the PubSub system.
// It allows customization of the number of buckets and the message buffer size for each bucket.
type Config struct {
	// BucketNum specifies the number of slices (buckets) used for message distribution.
	// This helps improve concurrency and load balancing.
	// Default value is 2 if not set or set to a non-positive number.
	BucketNum int

	// BucketMessageBuffer defines the buffer size for each bucket's publish message channel.
	// A larger buffer helps handle higher message throughput without blocking.
	// If the provided value is less than 256, it defaults to 256 to ensure efficient message handling.
	BucketMessageBuffer int
}

// init initializes the Config with default values if not explicitly set.
// Ensures the system operates with optimal performance settings.
func (c *Config) init() {
	if c.BucketNum <= 0 {
		c.BucketNum = 2
	}

	if c.BucketMessageBuffer <= 256 {
		c.BucketMessageBuffer = 256
	}
}
