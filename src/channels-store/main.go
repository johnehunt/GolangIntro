package main

import "fmt"

type entry struct {
	key   string
	value string
}

type query struct {
	key   string
	reply chan string
}

type store struct {
	data map[string]string
	put  chan entry
	get  chan query
}

func main() {
	s := NewStore()
	s.Put("k", "v")
	fmt.Println(s.Get("k"))
}

func NewStore() *store {
	s := &store{
		data: map[string]string{},
		put:  make(chan entry),
		get:  make(chan query),
	}

	go s.listen()

	return s
}

func (s *store) Put(key, value string) {
	s.put <- entry{key: key, value: value}
}

func (s *store) Get(key string) string {
	c := make(chan string)

	s.get <- query{key: key, reply: c}

	return <-c
}

func (s *store) listen() string {
	for {
		select {
		case e := <-s.put:
			s.data[e.key] = e.value
		case r := <-s.get:
			r.reply <- s.data[r.key]
		}
	}
}
