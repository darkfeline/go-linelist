// Copyright (C) 2019  Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linelist

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Store opens lists and tracks them so they can be written back.
// The zero value is ready to use.  Store should not be copied
// after use.
type Store struct {
	err     error
	toWrite []trackedList
}

type trackedList struct {
	path string
	ls   *List
}

// trackList adds a list to write out later.
func (s *Store) trackList(path string, ls *List) {
	s.toWrite = append(s.toWrite, trackedList{
		path: path,
		ls:   ls,
	})
}

// getList gets a previously tracked list and returns the same slice
// pointer.  If the list isn't tracked, return nil.
func (s *Store) getList(path string) *List {
	for _, tl := range s.toWrite {
		if tl.path == path {
			return tl.ls
		}
	}
	return nil
}

// Open opens a list with the given path.  Any errors will be reported
// by Err.  If the Store already encountered an error, this function
// does nothing.
//
// The returned List pointer can be modified and will be written back
// to the path when Flush is called.
func (s *Store) Open(path string) *List {
	if s.err != nil {
		return nil
	}
	if list := s.getList(path); list != nil {
		return list
	}
	f, err := os.Open(path)
	if err != nil {
		s.err = err
		return nil
	}
	defer f.Close()
	ls, err := Load(f)
	if err != nil {
		s.err = err
		return nil
	}
	s.trackList(path, &ls)
	return &ls
}

// Err returns any error that the Store has encountered.
func (s *Store) Err() error {
	return s.err
}

// Flush writes any opened lists back to the files they were read
// from.  If the Store already encountered an error, this function
// does nothing.  If no errors occur, this function can be called
// multiple times.
func (s *Store) Flush() error {
	if s.err != nil {
		return s.err
	}
	for _, t := range s.toWrite {
		if err := s.flushOne(t); err != nil {
			s.err = err
			return s.err
		}
	}
	return nil
}

func (s *Store) flushOne(t trackedList) error {
	f, err := ioutil.TempFile(filepath.Dir(t.path), "tmp*~")
	if err != nil {
		return err
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	t.ls.WriteTo(bw)
	if err := bw.Flush(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	os.Rename(f.Name(), t.path)
	return nil
}
