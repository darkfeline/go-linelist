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

// Set is a set of lines of text.
type Set map[string]struct{}

// List returns the lines in the set as a list in arbitrary order.
func (s Set) List() List {
	list := make(List, len(s))
	for l := range s {
		list = append(list, l)
	}
	return list
}

// Contains returns true if the set contains the line.
func (s Set) Contains(l string) bool {
	_, ok := s[l]
	return ok
}

// Add adds a line to the set.
func (s Set) Add(l string) {
	s[l] = struct{}{}
}
