package servlet

import (
	"time"

	"github.com/spf13/viper"
)

// Config
// Configuration interface
type Config interface {
	Get(k string, v ...interface{}) interface{}
	GetString(k string, v ...string) string
	GetBool(k string, v ...bool) bool
	GetInt(k string, v ...int) int
	GetUInt(k string, v ...uint) uint
	GetInt16(k string, v ...int16) int16
	GetUInt16(k string, v ...uint16) uint16
	GetInt32(k string, v ...int32) int32
	GetUInt32(k string, v ...uint32) uint32
	GetInt64(k string, v ...int64) int64
	GetUInt64(k string, v ...uint64) uint64
	GetFloat32(k string, v ...float32) float32
	GetFloat64(k string, v ...float64) float64
	GetTime(k string, v ...time.Time) time.Time
	GetDuration(k string, v ...time.Duration) time.Duration

	GetStringSlice(k string, v ...[]string) []string
	GetStringMap(k string, v ...map[string]interface{}) map[string]interface{}
	GetStringMapString(k string, v ...map[string]string) map[string]string
	GetStringMapStringSlice(k string, v ...map[string][]string) map[string][]string

	Unmarshal(v interface{}) error
	UnmarshalKey(k string, v interface{}) error
	Sub(k string) Config
	Has(k string) bool
}

type vConfig struct {
	v *viper.Viper
}

// MakeConfig
// Make a Config instance from Viper.
func MakeConfig(v *viper.Viper) Config {
	if nil == v {
		v = viper.New()
	}
	_v := &vConfig{v: v}
	return _v
}

func (c vConfig) Sub(k string) Config {
	v := c.v.Sub(k)
	if nil == v {
		return nil
	}
	return MakeConfig(v)
}

func (c vConfig) Has(k string) bool {
	return c.v.InConfig(k)
}

func (c vConfig) Get(k string, v ...interface{}) interface{} {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.Get(k)
}
func (c vConfig) GetString(k string, v ...string) string {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetString(k)
}
func (c vConfig) GetBool(k string, v ...bool) bool {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetBool(k)
}
func (c vConfig) GetInt(k string, v ...int) int {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetInt(k)
}
func (c vConfig) GetUInt(k string, v ...uint) uint {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return uint(c.v.GetInt(k))
}
func (c vConfig) GetInt16(k string, v ...int16) int16 {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return int16(c.v.GetInt(k))
}
func (c vConfig) GetUInt16(k string, v ...uint16) uint16 {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return uint16(c.v.GetInt(k))
}
func (c vConfig) GetInt32(k string, v ...int32) int32 {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetInt32(k)
}
func (c vConfig) GetUInt32(k string, v ...uint32) uint32 {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return uint32(c.v.GetInt32(k))
}
func (c vConfig) GetInt64(k string, v ...int64) int64 {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetInt64(k)
}
func (c vConfig) GetUInt64(k string, v ...uint64) uint64 {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return uint64(c.v.GetInt64(k))
}
func (c vConfig) GetFloat32(k string, v ...float32) float32 {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return float32(c.v.GetFloat64(k))
}
func (c vConfig) GetFloat64(k string, v ...float64) float64 {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetFloat64(k)
}
func (c vConfig) GetTime(k string, v ...time.Time) time.Time {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetTime(k)
}
func (c vConfig) GetDuration(k string, v ...time.Duration) time.Duration {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetDuration(k)
}

func (c vConfig) GetStringSlice(k string, v ...[]string) []string {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetStringSlice(k)
}
func (c vConfig) GetStringMap(k string, v ...map[string]interface{}) map[string]interface{} {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetStringMap(k)
}
func (c vConfig) GetStringMapString(k string, v ...map[string]string) map[string]string {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetStringMapString(k)
}
func (c vConfig) GetStringMapStringSlice(k string, v ...map[string][]string) map[string][]string {
	if len(v) > 0 && !(c.v.InConfig(k) || c.v.IsSet(k)) {
		return v[0]
	}
	return c.v.GetStringMapStringSlice(k)
}

func (c vConfig) Unmarshal(v interface{}) error {
	return c.v.Unmarshal(v)
}

func (c vConfig) UnmarshalKey(k string, v interface{}) error {
	return c.v.UnmarshalKey(k, v)
}
