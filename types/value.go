package types

// Value is an interface that defines a holder for any DragonScript value type.
// Value types should have one method that return the GoValue associated with
// that type.
type Value interface {
	GoValue() interface{}
}
