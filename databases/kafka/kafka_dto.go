package kafka

// Validate method for validate STR values
func (ks *STR) Validate() {
	if len(ks.Brokers) == 0 {
		panic("kafka need Brokers, please set the Brokers")
	}
	if len(ks.Version) == 0 {
		panic("kafka need Version, please set the Version")
	}
	if len(ks.Group) == 0 {
		panic("kafka need Group, please set the Group")
	}
	if len(ks.Topics) == 0 {
		panic("kafka need Topics, please set the Topics")
	}
	if len(ks.Assignor) == 0 {
		panic("kafka need Assignor, please set the Assignor")
	}
	if ks.Oldest == nil {
		panic("kafka need Oldest, please set the Oldest")
	}
	if ks.Verbose == nil {
		panic("kafka need Verbose, please set the Verbose")
	}
	if ks.Config == nil {
		panic("kafka need Config, please set the Config")
	}
}
