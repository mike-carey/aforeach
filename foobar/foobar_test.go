package foobar_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/async/foobar"
)

var _ = Describe("Foobar", func() {
	Context("ForEach", func() {
		It("Should cycle through every Foo and return proper Bars", func() {
			foos := make([]Foo, 0)
			bars := make([]Bar, 0)

			for _, i := range []string{"one", "two", "three"} {
				foo, bar := NewFooBarPair(i)
				foos = append(foos, foo)
				bars = append(bars, bar)
			}

			_bars, errs := ForEachFooToBar(foos, func(foo Foo) (Bar, error) {
				return *foo.Bar, nil
			})

			Expect(errs).To(BeEmpty())
			Expect(_bars).Should(ConsistOf(bars))
		})

		It("Should cycle through every Foo and return errors", func() {
			foos := make([]Foo, 0)
			bars := make([]Bar, 0)

			for _, i := range []string{"one", "two", "three"} {
				foo, bar := NewFooBarPair(i)
				foos = append(foos, foo)
				if foo.Name != "one" {
					bars = append(bars, bar)
				}
			}

			err := errors.New("Name is one")

			_bars, errs := ForEachFooToBar(foos, func(foo Foo) (Bar, error) {
				if foo.Name == "one" {
					return Bar{}, err
				}
				return *foo.Bar, nil
			})

			Expect(errs).Should(ConsistOf(err))
			Expect(_bars).Should(ConsistOf(bars))
		})
	})

	Context("MapOneToMany", func() {
		It("Should cycle though every Baz and return many Bars", func() {
			foos := make([]Foo, 0)
			bars := make([]Bar, 0)
			bazs := make([]Baz, 0)

			// var prevBaz *Baz

			for _, i := range []string{"one", "two", "three"} {
				foo, bar, baz := NewFooBarBazGroup(i)
				foos = append(foos, foo)
				bars = append(bars, bar)

				baz.Foo = foos
				bazs = append(bazs, baz)
			}

			_bars, errs := MapOneBazToManyBar(bazs, func(baz Baz) ([]Bar, error) {
				bs := make([]Bar, 0)
				for _, foo := range baz.Foo {
					bs = append(bs, *foo.Bar)
				}
				return bs, nil
			})

			Expect(errs).To(BeEmpty())
			Expect(_bars).Should(HaveKeyWithValue(&bazs[0], []Bar{bars[0]}))
			Expect(_bars).Should(HaveKeyWithValue(&bazs[1], []Bar{bars[0], bars[1]}))
			Expect(_bars).Should(HaveKeyWithValue(&bazs[2], bars))
		})
	})

})
