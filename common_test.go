package btree

func assert(x bool) {
	if !x {
		panic("assert failed")
	}
}

// testKind is the item type.
// It's important to use the equal symbol, which tells Go to create an alias of
// the type, rather than creating an entirely new type.
type testKind = int

func testLess(a, b testKind) bool {
	return a < b
}

func intLess(a, b any) bool {
	return a.(int) < b.(int)
}

func kindsAreEqual[T ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
