package memory

import "github.com/bbuck/dragonscript"

// M is a basic memory interface that defines the behavior any aspect of "memory"
// should exhibit. Such as getting and setting values.
type M interface {
	Get(string) dragonscript.Value
	Set(string) dragonscript.Value
}
