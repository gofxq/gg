package main

import (
	"log"
	"reflect"

	"github.com/gofxq/gg/set"
)

func main() {
	s2 := set.NewSet("a")
	s2.Add("b")
	log.Println(s2.ToList())

	s3 := set.NewSet("a", "b")
	log.Println(reflect.TypeOf(s3))

	s4 := set.NewSet[int]()
	s4.Add(1)
	log.Println(reflect.TypeOf(s4))
}
