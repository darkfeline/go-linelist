package linelist

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestStore(t *testing.T) {
	t.Parallel()
	t.Run("edit new list", func(t *testing.T) {
		t.Parallel()
		d, err := ioutil.TempDir("", "test")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		var s Store
		f := filepath.Join(d, "list")
		ls := s.Open(f)
		if err := s.Err(); err != nil {
			t.Fatal(err)
		}
		*ls = []string{"lacia", "prachta"}
		if err := s.Flush(); err != nil {
			t.Fatal(err)
		}
		got, err := ioutil.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		want := []byte("lacia\nprachta\n")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("File contained %v; want %v", got, want)
		}
	})
	t.Run("edit existing list", func(t *testing.T) {
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
	})
}
