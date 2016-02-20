package dragonscript

// Value represents a data type from within DragonScript as it is represented
// in Go. To get it's "go type" you would call GoValue().
type Value interface {
	GoValue() interface{}
}
