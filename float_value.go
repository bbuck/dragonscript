package dragonscript

// Float represent a floating point value inside of the DragonScript langauge.
type Float float64

// GoValue returns a float64 value associated with this node for use by Go.
func (f Float) GoValue() interface{} {
	return float64(f)
}
