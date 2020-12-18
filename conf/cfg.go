package conf

// LogTransfer 全局配置
type LogTransfer struct {
	ESCfg    `ini:"es"`
	KafkaCfg `ini:"kafka"`
}

// KafkaCfg is a struct
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// ESCfg is a struct
type ESCfg struct {
	Address string `ini:"address"`
}
