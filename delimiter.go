// Lute - A structured markdown engine.
// Copyright (C) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lute

type delimiter struct {
	node           Node      // the text node point to
	typ            itemType   // the type of delimiter ([, ![, *, _)
	num            int        // the number of delimiters
	originalNum    int        // the original number of delimiters
	active         bool       // whether the delimiter is "active" (all are active to start)
	canOpen        bool       // whether the delimiter is a potential opener
	canClose       bool       // whether the delimiter is a potential closer
	previous, next *delimiter // doubly linked list
}

func (t *Tree) removeDelimiter(delim *delimiter) (ret *delimiter) {
	if delim.previous != nil {
		delim.previous.next = delim.next
	}
	if delim.next == nil {
		// top of stack
		t.context.Delimiters = delim.previous
	} else {
		delim.next.previous = delim.previous
	}

	return
}