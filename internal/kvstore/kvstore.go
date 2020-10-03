package kvstore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Elem represents the element being stored
type Elem interface {
	Print()
	GetStatus() int
	ChangeStatus(s int)
}

// Store represents the KV store struct
type Store struct {
	Data map[string]Elem
}

// DATAFILE contains the filename
var DATAFILE = "./datafile.json"

// New returns a new store
func New() *Store {
	return &Store{Data: make(map[string]Elem)}
}

// Save will write data to a file
func (s *Store) Save() error {
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(os.ExpandEnv(DATAFILE))
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()

	b, err := json.MarshalIndent(s.Data, "", " ")
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	fmt.Println(string(b))
	saveTo.Write(b)

	return nil
}

// Load will load data from a file
func (s *Store) Load() error {
	fmt.Println("Loading", DATAFILE)
	loadFrom, err := os.Open(DATAFILE)
	defer loadFrom.Close()
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}

	b, err := ioutil.ReadAll(loadFrom)
	if err != nil {
		fmt.Println("Unable to read from file!")
		return err
	}

	err = json.Unmarshal(b, &s.Data)
	if err != nil {
		fmt.Println("Unable unmarshal JSON file!")
		return err
	}

	fmt.Printf("%#v\n", s)
	return nil
}

// LOOKUP returns element by key
func (s *Store) LOOKUP(k string) *Elem {
	_, ok := s.Data[k]
	if ok {
		n := s.Data[k]
		return &n
	}
	return nil
}

// ADD elements to the store
func (s *Store) ADD(k string, n Elem) bool {
	if k == "" {
		return false
	}

	if s.LOOKUP(k) == nil {
		s.Data[k] = n
		return true
	}
	return false
}

// DELETE removes elem from store
func (s *Store) DELETE(k string) bool {
	if s.LOOKUP(k) != nil {
		delete(s.Data, k)
		return true
	}
	return false
}

// CHANGE modifies the content of store located in data[i]
func (s *Store) CHANGE(k string, n Elem) bool {
	s.Data[k] = n
	return true
}

// PRINT displays the content of store
func (s *Store) PRINT() {
	for k, v := range s.Data {
		fmt.Printf("key: %s\nval:", k)
		v.Print()
	}
}
