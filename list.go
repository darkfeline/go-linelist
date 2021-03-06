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

// Package linelist provides tools for working with lists of things
// stored in text files, one thing per line.
//
// This makes it easy to dump captured text from various sources and
// automatically sort, process, and deduplicate them for a personal
// workflow.
package linelist

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// List is a list of lines of text.
type List []string

// Load loads a line list from a reader.
func Load(r io.Reader) (List, error) {
	var l List
	s := bufio.NewScanner(r)
	for s.Scan() {
		l = append(l, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("linelist: load: %w", err)
	}
	return l, nil
}

// WriteTo implements io.WriterTo.
func (l List) WriteTo(w io.Writer) (n int64, err error) {
	bw := bufio.NewWriter(w)
	for _, s := range l {
		n2, _ := io.WriteString(bw, s)
		n += int64(n2)
		n2, _ = bw.Write([]byte{'\n'})
		n += int64(n2)
	}
	err = bw.Flush()
	if err != nil {
		err = fmt.Errorf("linelist: write: %w", err)
	}
	return n, err
}

// Set returns a set of the lines in the list.
func (l List) Set() Set {
	s := make(Set, len(l))
	for _, line := range l {
		s.Add(line)
	}
	return s
}

// Exclude returns a list excluding the lines in the argument list.
func (l List) Exclude(a List) List {
	return l.ExcludeSet(a.Set())
}

// ExcludeSet returns a list excluding the lines in the argument set.
func (l List) ExcludeSet(s Set) List {
	result := make(List, 0, len(l))
	for _, line := range l {
		if !s.Contains(line) {
			result = append(result, line)
		}
	}
	return result
}

// Keep returns a list keeping only the lines in the argument list.
func (l List) Keep(a List) List {
	return l.KeepSet(a.Set())
}

// KeepSet returns a list keeping only the lines in the argument set.
func (l List) KeepSet(s Set) List {
	result := make(List, 0, len(l))
	for _, line := range l {
		if s.Contains(line) {
			result = append(result, line)
		}
	}
	return result
}

// Unique returns a list with only unique lines, preserving order.
func (l List) Unique() List {
	m := make(map[string]bool)
	result := make(List, 0, len(l))
	for _, line := range l {
		if !m[line] {
			result = append(result, line)
			m[line] = true
		}
	}
	return result
}

// Sort sorts the list in place.
func (l List) Sort() {
	sort.Strings(l)
}
