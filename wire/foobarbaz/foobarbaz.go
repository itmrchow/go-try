//go:build wireinject
// +build wireinject

package foobarbaz

import (
	"context"
	"errors"

	"github.com/google/wire"
)

type Foo struct {
	X int
}

type Bar struct {
	X int
}

type Baz struct {
	X int
}

func ProvideFoo() Foo {
	return Foo{X: 42}
}

func ProvideBar(foo Foo) Bar {
	return Bar{X: -foo.X}
}

func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X < 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}

	return Baz{X: bar.X}, nil
}

var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

// var MergeSet =  wire.Newset(SuperSet,pkg.OtherSet)
