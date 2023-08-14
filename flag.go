package cliff

import (
	"errors"
	"net"
	"time"
	"unsafe"

	"github.com/spf13/pflag"
)

// Constraint makes sure that the given type is one of the supported.
type Constraint interface {
	[]bool |
		[]byte |
		[]float32 | []float64 |
		[]int | []int32 | []int64 |
		[]net.IP |
		[]string |
		[]time.Duration |
		[]uint |
		bool |
		float32 | float64 |
		int | int16 | int32 | int64 | int8 |
		map[string]int | map[string]int64 | map[string]string |
		net.IP | net.IPMask | net.IPNet |
		string |
		time.Duration |
		uint | uint16 | uint32 | uint64 | uint8 |
		Count | BytesHex | BytesBase64
}

// Short is a literal character representing shortcut for a flag.
type Short rune

// Help is a literal string describing the flag usage.
type Help string

// Count is an int represented in CLI by repeating the argument N times.
// For example, "-vvv" will be parsed as 3.
type Count int

// BytesHex is a slice of bytes represented in CLI as a hexadecimal-encoded string.
type BytesHex []byte

// BytesBase64 is a slice of bytes represented in CLI as a base64-encoded string.
type BytesBase64 []byte

// tPFlag represents all info about a CLI flag except its name.
type tPFlag struct {
	baseFlag
	tar   any    // target where to put the parsed result
	def   any    // default value to use if flag not specified
	short string // short alias for the flag
	help  string // usage message
}

// F creates a new flag.
func F[T Constraint](val *T, short Short, def T, help Help) Flag {
	shortStr := ""
	if short != 0 {
		shortStr = string(short)
	}
	return tPFlag{
		tar:   val,
		def:   def,
		short: shortStr,
		help:  string(help),
	}
}

func (f tPFlag) Deprecated(message string) Flag {
	f.depr = message
	return f
}

func (f tPFlag) ShortDeprecated(message string) Flag {
	f.shortDepr = message
	return f
}

func (f tPFlag) Hidden() Flag {
	f.hidden = true
	return f
}

func (f tPFlag) AddTo(fs *pflag.FlagSet, name string) error {
	if f.short != "" && !isAlNum(f.short) {
		return errors.New("flag short name must be an alpha-numeric ASCII character")
	}
	err := f.pflagAddFlag(name, fs)
	if err != nil {
		return err
	}
	return f.setProperties(fs, name)
}

func (f tPFlag) pflagAddFlag(name string, fs *pflag.FlagSet) error {
	switch def := any(f.def).(type) {
	case []bool:
		v := any(f.tar).(*[]bool)
		fs.BoolSliceVarP(v, name, f.short, def, f.help)
	case bool:
		v := any(f.tar).(*bool)
		fs.BoolVarP(v, name, f.short, def, f.help)
	case []byte:
		v := any(f.tar).(*[]byte)
		fs.BytesHexVarP(v, name, f.short, def, f.help)
	case BytesHex:
		v := any(f.tar).(*BytesHex)
		p := (*[]byte)(unsafe.Pointer(v))
		fs.BytesHexVarP(p, name, f.short, def, f.help)
	case BytesBase64:
		v := any(f.tar).(*BytesBase64)
		p := (*[]byte)(unsafe.Pointer(v))
		fs.BytesBase64VarP(p, name, f.short, def, f.help)
	case []time.Duration:
		v := any(f.tar).(*[]time.Duration)
		fs.DurationSliceVarP(v, name, f.short, def, f.help)
	case time.Duration:
		v := any(f.tar).(*time.Duration)
		fs.DurationVarP(v, name, f.short, def, f.help)
	case []float32:
		v := any(f.tar).(*[]float32)
		fs.Float32SliceVarP(v, name, f.short, def, f.help)
	case float32:
		v := any(f.tar).(*float32)
		fs.Float32VarP(v, name, f.short, def, f.help)
	case []float64:
		v := any(f.tar).(*[]float64)
		fs.Float64SliceVarP(v, name, f.short, def, f.help)
	case float64:
		v := any(f.tar).(*float64)
		fs.Float64VarP(v, name, f.short, def, f.help)
	case net.IPMask:
		v := any(f.tar).(*net.IPMask)
		fs.IPMaskVarP(v, name, f.short, def, f.help)
	case net.IPNet:
		v := any(f.tar).(*net.IPNet)
		fs.IPNetVarP(v, name, f.short, def, f.help)
	case []net.IP:
		v := any(f.tar).(*[]net.IP)
		fs.IPSliceVarP(v, name, f.short, def, f.help)
	case net.IP:
		v := any(f.tar).(*net.IP)
		fs.IPVarP(v, name, f.short, def, f.help)
	case int16:
		v := any(f.tar).(*int16)
		fs.Int16VarP(v, name, f.short, def, f.help)
	case []int32:
		v := any(f.tar).(*[]int32)
		fs.Int32SliceVarP(v, name, f.short, def, f.help)
	case int32:
		v := any(f.tar).(*int32)
		fs.Int32VarP(v, name, f.short, def, f.help)
	case []int64:
		v := any(f.tar).(*[]int64)
		fs.Int64SliceVarP(v, name, f.short, def, f.help)
	case int64:
		v := any(f.tar).(*int64)
		fs.Int64VarP(v, name, f.short, def, f.help)
	case int8:
		v := any(f.tar).(*int8)
		fs.Int8VarP(v, name, f.short, def, f.help)
	case []int:
		v := any(f.tar).(*[]int)
		fs.IntSliceVarP(v, name, f.short, def, f.help)
	case int:
		v := any(f.tar).(*int)
		fs.IntVarP(v, name, f.short, def, f.help)
	case Count:
		v := any(f.tar).(*Count)
		p := (*int)(unsafe.Pointer(v))
		fs.CountVarP(p, name, f.short, f.help)
	case []string:
		v := any(f.tar).(*[]string)
		fs.StringSliceVarP(v, name, f.short, def, f.help)
	case map[string]int64:
		v := any(f.tar).(*map[string]int64)
		fs.StringToInt64VarP(v, name, f.short, def, f.help)
	case map[string]int:
		v := any(f.tar).(*map[string]int)
		fs.StringToIntVarP(v, name, f.short, def, f.help)
	case map[string]string:
		v := any(f.tar).(*map[string]string)
		fs.StringToStringVarP(v, name, f.short, def, f.help)
	case string:
		v := any(f.tar).(*string)
		fs.StringVarP(v, name, f.short, def, f.help)
	case uint16:
		v := any(f.tar).(*uint16)
		fs.Uint16VarP(v, name, f.short, def, f.help)
	case uint32:
		v := any(f.tar).(*uint32)
		fs.Uint32VarP(v, name, f.short, def, f.help)
	case uint64:
		v := any(f.tar).(*uint64)
		fs.Uint64VarP(v, name, f.short, def, f.help)
	case uint8:
		v := any(f.tar).(*uint8)
		fs.Uint8VarP(v, name, f.short, def, f.help)
	case []uint:
		v := any(f.tar).(*[]uint)
		fs.UintSliceVarP(v, name, f.short, def, f.help)
	case uint:
		v := any(f.tar).(*uint)
		fs.UintVarP(v, name, f.short, def, f.help)
	default:
		return errors.New("unsupported type")
	}
	return nil
}
