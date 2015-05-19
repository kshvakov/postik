package postik

var hash Hasher

type Hasher func(string) string

func init() {

	hash = func(name string) string {

		return name
	}
}

func ReplaceHasher(hasher Hasher) {

	hash = hasher
}
