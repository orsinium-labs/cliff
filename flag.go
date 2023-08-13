package cliff

import (
	"net"
	"time"

	"github.com/spf13/pflag"
)

type Constraint interface {
	[]bool |
		[]byte |
		[]float32 |
		[]float64 |
		[]int |
		[]int32 |
		[]int64 |
		[]net.IP |
		[]string |
		[]time.Duration |
		[]uint |
		bool |
		float32 |
		float64 |
		int |
		int16 |
		int32 |
		int64 |
		int8 |
		map[string]int |
		map[string]int64 |
		map[string]string |
		net.IP |
		net.IPMask |
		net.IPNet |
		string |
		time.Duration |
		uint |
		uint16 |
		uint32 |
		uint64 |
		uint8
}

type Name string
type Short rune
type Help string

type tFlag struct {
	T       any
	Default any
	Short   string
	Help    string
}

func Fn[T Constraint](val *T, short Short, def T, help Help) tFlag {
	shortStr := ""
	if short != 0 {
		shortStr = string(short)
	}
	return tFlag{
		T:       val,
		Default: def,
		Short:   shortStr,
		Help:    string(help),
	}
}

func (f tFlag) pflagAdd(name string, fs *pflag.FlagSet) {
	switch def := any(f.Default).(type) {
	case []bool:
		v := any(f.T).(*[]bool)
		fs.BoolSliceVarP(v, name, f.Short, def, f.Help)
	case bool:
		v := any(f.T).(*bool)
		fs.BoolVarP(v, name, f.Short, def, f.Help)
	case []byte:
		v := any(f.T).(*[]byte)
		fs.BytesHexVarP(v, name, f.Short, def, f.Help)
	case []time.Duration:
		v := any(f.T).(*[]time.Duration)
		fs.DurationSliceVarP(v, name, f.Short, def, f.Help)
	case time.Duration:
		v := any(f.T).(*time.Duration)
		fs.DurationVarP(v, name, f.Short, def, f.Help)
	case []float32:
		v := any(f.T).(*[]float32)
		fs.Float32SliceVarP(v, name, f.Short, def, f.Help)
	case float32:
		v := any(f.T).(*float32)
		fs.Float32VarP(v, name, f.Short, def, f.Help)
	case []float64:
		v := any(f.T).(*[]float64)
		fs.Float64SliceVarP(v, name, f.Short, def, f.Help)
	case float64:
		v := any(f.T).(*float64)
		fs.Float64VarP(v, name, f.Short, def, f.Help)
	case net.IPMask:
		v := any(f.T).(*net.IPMask)
		fs.IPMaskVarP(v, name, f.Short, def, f.Help)
	case net.IPNet:
		v := any(f.T).(*net.IPNet)
		fs.IPNetVarP(v, name, f.Short, def, f.Help)
	case []net.IP:
		v := any(f.T).(*[]net.IP)
		fs.IPSliceVarP(v, name, f.Short, def, f.Help)
	case net.IP:
		v := any(f.T).(*net.IP)
		fs.IPVarP(v, name, f.Short, def, f.Help)
	case int16:
		v := any(f.T).(*int16)
		fs.Int16VarP(v, name, f.Short, def, f.Help)
	case []int32:
		v := any(f.T).(*[]int32)
		fs.Int32SliceVarP(v, name, f.Short, def, f.Help)
	case int32:
		v := any(f.T).(*int32)
		fs.Int32VarP(v, name, f.Short, def, f.Help)
	case []int64:
		v := any(f.T).(*[]int64)
		fs.Int64SliceVarP(v, name, f.Short, def, f.Help)
	case int64:
		v := any(f.T).(*int64)
		fs.Int64VarP(v, name, f.Short, def, f.Help)
	case int8:
		v := any(f.T).(*int8)
		fs.Int8VarP(v, name, f.Short, def, f.Help)
	case []int:
		v := any(f.T).(*[]int)
		fs.IntSliceVarP(v, name, f.Short, def, f.Help)
	case int:
		v := any(f.T).(*int)
		fs.IntVarP(v, name, f.Short, def, f.Help)
	case []string:
		v := any(f.T).(*[]string)
		fs.StringSliceVarP(v, name, f.Short, def, f.Help)
	case map[string]int64:
		v := any(f.T).(*map[string]int64)
		fs.StringToInt64VarP(v, name, f.Short, def, f.Help)
	case map[string]int:
		v := any(f.T).(*map[string]int)
		fs.StringToIntVarP(v, name, f.Short, def, f.Help)
	case map[string]string:
		v := any(f.T).(*map[string]string)
		fs.StringToStringVarP(v, name, f.Short, def, f.Help)
	case string:
		v := any(f.T).(*string)
		fs.StringVarP(v, name, f.Short, def, f.Help)
	case uint16:
		v := any(f.T).(*uint16)
		fs.Uint16VarP(v, name, f.Short, def, f.Help)
	case uint32:
		v := any(f.T).(*uint32)
		fs.Uint32VarP(v, name, f.Short, def, f.Help)
	case uint64:
		v := any(f.T).(*uint64)
		fs.Uint64VarP(v, name, f.Short, def, f.Help)
	case uint8:
		v := any(f.T).(*uint8)
		fs.Uint8VarP(v, name, f.Short, def, f.Help)
	case []uint:
		v := any(f.T).(*[]uint)
		fs.UintSliceVarP(v, name, f.Short, def, f.Help)
	case uint:
		v := any(f.T).(*uint)
		fs.UintVarP(v, name, f.Short, def, f.Help)
	}
}
