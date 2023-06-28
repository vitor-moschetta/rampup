package config

// Configs Keys.
const (
	KvsEnabledKey       = "KEY_VALUE_STORE_ENABLED"
	KvsContainerNameKey = "KEY_VALUE_STORE_CONTAINER_NAME"
	KvsContainerKey     = "KEY_VALUE_STORE_SBOXM7EZ8X9VHN_CONTAINER_NAME"
	KvsHostReadKey      = "KEY_VALUE_STORE_SBOXM7EZ8X9VHN_END_POINT_READ"
	KvsHostWriteKey     = "KEY_VALUE_STORE_SBOXM7EZ8X9VHN_END_POINT_WRITE"
)

var StringEmpty = ""

// ConfigMap is a map of all configs.
var ConfigMap = map[string]string{
	KvsEnabledKey:       StringEmpty,
	KvsContainerNameKey: StringEmpty,
}
