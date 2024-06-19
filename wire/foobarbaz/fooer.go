package foobarbaz

import "github.com/google/wire"

type Fooer interface {
	Foo() string
}

type MyFooer string

type BBar string

// MyFooer 實作 Fooer
func (b *MyFooer) Foo() string {
	return string(*b)
}

// MyFooer Provider
func provideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

// BBar Provider , BBar 依賴 Fooer
func provideBBar(f Fooer) string {
	// f will be a *MyFooer.
	return f.Foo()
}

var Set = wire.NewSet(
	provideMyFooer,                       // MyFooer
	wire.Bind(new(Fooer), new(*MyFooer)), // Fooer介面綁MyFooer
	provideBBar)                          // 注入
