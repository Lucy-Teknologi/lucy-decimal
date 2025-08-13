package decimal

import "fmt"

var (
	zero     = "0"
	One      = "1"
	Ten      = "10"
	Hundred  = "100"
	Thousand = "1000"
	Million  = "1000000"
	Billion  = "1000000000"
	Trillion = "1000000000000"
)

func IsZero(value string) bool {
	return value == zero
}

func ParseAny(value any) (string, error) {
	if value == nil {
		return zero, nil
	}
	switch v := value.(type) {
	case string:
		return v, nil
	case int:
		return NewFromInt(int64(v)).String(), nil
	case float64:
		return NewFromFloat(v).String(), nil
	default:
		return "", fmt.Errorf("unsupported type: %T", v)
	}
}
