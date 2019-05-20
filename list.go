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
	"io"
	"sort"
)

// List is a list of lines of text.
type List []string

// Load loads a line list from a reader.
func Load(r io.Reader) (List, error) {
	var ls List
	s := bufio.NewScanner(r)
	for s.Scan() {
		ls = append(ls, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return ls, nil
}

// WriteTo implements io.WriterTo.
func (ls List) WriteTo(w io.Writer) (n int64, err error) {
	bw := bufio.NewWriter(w)
	for _, s := range ls {
		n2, _ := io.WriteString(bw, s)
		n += int64(n2)
		n2, _ = bw.Write([]byte{'\n'})
		n += int64(n2)
	}
	return n, bw.Flush()
}

// lineSet returns a set of the lines in the list.
func (ls List) lineSet() map[string]bool {
	m := make(map[string]bool)
	for _, line := range ls {
		m[line] = true
	}
	return m
}

// Exclude returns a list excluding the lines in the argument list.
func (ls List) Exclude(a List) List {
	m := a.lineSet()
	result := make(List, 0, len(ls))
	for _, a := range ls {
		if !m[a] {
			result = append(result, a)
		}
	}
	return result
}

// Keep returns a list keeping only the lines in the argument.
func (ls List) Keep(a List) List {
	m := a.lineSet()
	result := make(List, 0, len(ls))
	for _, a := range ls {
		if m[a] {
			result = append(result, a)
		}
	}
	return result
}

// Unique returns a list with only unique lines, preserving order.
func (ls List) Unique() List {
	m := make(map[string]bool)
	result := make(List, 0, len(ls))
	for _, line := range ls {
		if !m[line] {
			result = append(result, line)
			m[line] = true
		}
	}
	return result
}

// Sort sorts the list in place.
func (ls List) Sort() {
	sort.Strings(ls)
}
