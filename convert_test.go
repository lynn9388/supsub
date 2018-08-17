/*
 * Copyright © 2018 Lynn <lynn9388@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package superscriptsubscript

import (
	"fmt"
	"sort"
	"testing"
)

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func TestToSuperscript(t *testing.T) {
	fmt.Println("Superscripts:")
	for _, r := range getSortedKeys(superscripts) {
		s, err := ToSuperscript(r)
		if err != nil && s != r {
			t.Errorf("no corresponding superscript and returns different value: ToSuperscript(%c) = %c", r, s)
		}
		fmt.Printf("%c%c ", r, s)
	}
	fmt.Println()
}

func TestToSubscript(t *testing.T) {
	fmt.Println("Subscripts:")
	for _, r := range getSortedKeys(subscripts) {
		s, err := ToSubscript(r)
		if err != nil && s != r {
			t.Errorf("no corresponding subscript and returns different value: ToSubscript(%c) = %c", r, s)
		}
		fmt.Printf("%c%c ", r, s)
	}
	fmt.Println()
}

func TestCombine(t *testing.T) {
	set := make(map[rune]struct{})
	for k := range superscripts {
		set[k] = struct{}{}
	}
	for k := range subscripts {
		set[k] = struct{}{}
	}

	var keys []rune
	for k := range set {
		keys = append(keys, k)
	}
	sort.Sort(runes(keys))

	for _, k := range keys {
		superscript, ok := superscripts[k]
		if !ok {
			superscript = ' '
		}
		subscript, ok := subscripts[k]
		if !ok {
			subscript = ' '
		}
		fmt.Printf("%c %c %c\n", k, superscript, subscript)
	}
}

func getSortedKeys(m map[rune]rune) []rune {
	var keys []rune
	for key := range m {
		keys = append(keys, key)
	}
	sort.Sort(runes(keys))
	return keys
}
