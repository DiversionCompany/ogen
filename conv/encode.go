package conv

import (
	"net/netip"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func IntToString(v int) string     { return strconv.Itoa(v) }
func Int8ToString(v int8) string   { return strconv.FormatInt(int64(v), 10) }
func Int16ToString(v int16) string { return strconv.FormatInt(int64(v), 10) }
func Int32ToString(v int32) string { return strconv.FormatInt(int64(v), 10) }
func Int64ToString(v int64) string { return strconv.FormatInt(v, 10) }

func UintToString(v uint) string     { return strconv.FormatUint(uint64(v), 10) }
func Uint8ToString(v uint8) string   { return strconv.FormatUint(uint64(v), 10) }
func Uint16ToString(v uint16) string { return strconv.FormatUint(uint64(v), 10) }
func Uint32ToString(v uint32) string { return strconv.FormatUint(uint64(v), 10) }
func Uint64ToString(v uint64) string { return strconv.FormatUint(v, 10) }

func Float32ToString(v float32) string { return strconv.FormatFloat(float64(v), 'f', 10, 64) }

func Float64ToString(v float64) string { return strconv.FormatFloat(v, 'f', 10, 64) }

func StringToString(v string) string { return v }

func BytesToString(v []byte) string { return string(v) }

func TimeToString(v time.Time) string { return v.Format(timeLayout) }

func DateToString(v time.Time) string { return v.Format(dateLayout) }

func DateTimeToString(v time.Time) string { return v.Format(time.RFC3339) }

func UnixSecondsToString(v time.Time) string {
	return StringInt64ToString(v.Unix())
}

func UnixNanoToString(v time.Time) string {
	return StringInt64ToString(v.UnixNano())
}

func UnixMicroToString(v time.Time) string {
	return StringInt64ToString(v.UnixMicro())
}

func UnixMilliToString(v time.Time) string {
	return StringInt64ToString(v.UnixMilli())
}

func BoolToString(v bool) string { return strconv.FormatBool(v) }

func UUIDToString(v uuid.UUID) string { return v.String() }

func AddrToString(v netip.Addr) string { return v.String() }

func URLToString(v url.URL) string { return v.String() }

func DurationToString(v time.Duration) string { return v.String() }

func StringInt32ToString(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

func StringInt64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Int32ArrayToString(vs []int32) []string {
	strs := make([]string, 0, len(vs))
	for _, v := range vs {
		strs = append(strs, Int32ToString(v))
	}
	return strs
}

func Int64ArrayToString(vs []int64) []string {
	strs := make([]string, 0, len(vs))
	for _, v := range vs {
		strs = append(strs, Int64ToString(v))
	}
	return strs
}

func Float32ArrayToString(vs []float32) []string {
	strs := make([]string, 0, len(vs))
	for _, v := range vs {
		strs = append(strs, Float32ToString(v))
	}
	return strs
}

func Float64ArrayToString(vs []float64) []string {
	strs := make([]string, 0, len(vs))
	for _, v := range vs {
		strs = append(strs, Float64ToString(v))
	}
	return strs
}

func StringArrayToString(vs []string) []string {
	return vs
}

func BytesArrayToString(vs [][]byte) []string {
	strs := make([]string, 0, len(vs))
	for _, v := range vs {
		strs = append(strs, BytesToString(v))
	}
	return strs
}

func TimeArrayToString(vs []time.Time) []string {
	strs := make([]string, 0, len(vs))
	for _, v := range vs {
		strs = append(strs, TimeToString(v))
	}
	return strs
}

func BoolArrayToString(vs []bool) []string {
	strs := make([]string, 0, len(vs))
	for _, v := range vs {
		strs = append(strs, BoolToString(v))
	}
	return strs
}

func UUIDArrayToString(vs []uuid.UUID) []string {
	strs := make([]string, 0, len(vs))
	for _, v := range vs {
		strs = append(strs, UUIDToString(v))
	}
	return strs
}
