package dragonscript

// Nil represents the DragonsScript nil value which means 'nothing'
type Nil struct{}

// GoValue returns nil as it's Go value.
func (n Nil) GoValue() interface{} {
	return nil
}
