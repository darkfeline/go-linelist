package linelist

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestStore(t *testing.T) {
	t.Parallel()
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	_ = f.Close()
	var s Store
	ls := s.Open(f.Name())
	if err := s.Err(); err != nil {
		t.Fatal(err)
	}
	*ls = []string{"lacia", "prachta"}
	if err := s.Flush(); err != nil {
		t.Fatal(err)
	}
	got, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("lacia\nprachta\n")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("File contained %v; want %v", got, want)
	}
}
