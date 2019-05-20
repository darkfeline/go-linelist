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
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Parallel()
	ls, err := Load(strings.NewReader("lacia\nkouka\n"))
	if err != nil {
		t.Fatalf("Error reading: %s", err)
	}
	exp := List{"lacia", "kouka"}
	if !reflect.DeepEqual(ls, exp) {
		t.Errorf("Got %#v, expected %#v", ls, exp)
	}
}

func TestList_WriteTo(t *testing.T) {
	t.Parallel()
	var b bytes.Buffer
	ls := List{
		"lacia",
		"sophie",
		"lacia",
		"clarion",
	}
	if _, err := ls.WriteTo(&b); err != nil {
		t.Fatalf("Write returned error: %s", err)
	}
	got := b.String()
	exp := "lacia\nsophie\nlacia\nclarion\n"
	if got != exp {
		t.Errorf("Got %#v, expected %#v", got, exp)
	}
}

func TestList_Exclude(t *testing.T) {
	t.Parallel()
	ls := List{
		"lacia",
		"sophie",
		"lacia",
		"clarion",
		"firis",
	}
	m := List{"lacia", "clarion"}
	got := ls.Exclude(m)
	exp := List{"sophie", "firis"}
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("Exclude(%#v, %#v) = %#v (expected %#v)", ls, m, got, exp)
	}
}

func TestList_Exclude_side_effect_free(t *testing.T) {
	t.Parallel()
	ls := List{
		"lacia",
		"sophie",
		"lacia",
		"clarion",
		"firis",
	}
	m := List{"lacia", "clarion"}
	_ = ls.Exclude(m)
	exp := List{
		"lacia",
		"sophie",
		"lacia",
		"clarion",
		"firis",
	}
	if !reflect.DeepEqual(ls, exp) {
		t.Errorf("List was modified from %#v to %#v", exp, ls)
	}
}

func TestList_Keep(t *testing.T) {
	t.Parallel()
	ls := List{
		"lacia",
		"sophie",
		"lacia",
		"clarion",
		"firis",
	}
	m := List{"lacia", "clarion", "nagato"}
	got := ls.Keep(m)
	exp := List{"lacia", "lacia", "clarion"}
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("Keep(%#v, %#v) = %#v (expected %#v)", ls, m, got, exp)
	}
}

func TestList_Keep_side_effect_free(t *testing.T) {
	t.Parallel()
	ls := List{
		"lacia",
		"sophie",
		"lacia",
		"clarion",
		"firis",
	}
	m := List{"lacia", "clarion"}
	_ = ls.Keep(m)
	exp := List{
		"lacia",
		"sophie",
		"lacia",
		"clarion",
		"firis",
	}
	if !reflect.DeepEqual(ls, exp) {
		t.Errorf("List was modified from %#v to %#v", exp, ls)
	}
}
