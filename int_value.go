package dragonscript

// Integer represents a DragonScript integer value.
type Integer int64

// GoValue returns an int64 value that has the same value as the DragonScript
// Integer
func (i Integer) GoValue() interface{} {
	return int64(i)
}
