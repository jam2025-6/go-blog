package config

// 高德地图配置
type Gaode struct {
	Enabled bool   `json:"enabled" yaml:"enabled"` // 是否启用高德地图
	Key     string `json:"key" yaml:"key"`         // 高德地图 Key
}
