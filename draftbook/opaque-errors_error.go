package draftbook

// import "github.com/quux/bar"

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

// func fn() error {
// 	x, err := bar.Foo()
// 	if err != nil {
// 		return err
// 	}
// }
