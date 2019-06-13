// Lute - A structural markdown engine.
// Copyright (C) 2019-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package lute

func (t *Tree) parseBlocks() {
	curNode := t.context.CurNode
	for token := t.peek(); itemEOF != token.typ; token = t.peek() {
		t.parseBlock()
		t.context.CurNode = curNode
	}
}

func (t *Tree) parseBlock() (ret Node) {
	curNode := t.context.CurNode
	token := t.peek()

	switch token.typ {
	case itemAsterisk, itemHyphen:
		ret = t.parseList()
	case itemCrosshatch:
		ret = t.parseHeading()
	case itemGreater:
		ret = t.parseBlockquote()
	case itemSpace, itemTab:
		spaces, tabs, tokens, firstNonWhitespace := t.nextNonWhitespace()
		if itemAsterisk == firstNonWhitespace.typ || itemHyphen == firstNonWhitespace.typ {
			t.backups(tokens)
			ret = t.parseList()
		} else {
			if 1 <= tabs || 4 <= spaces {
				t.backups(tokens)
				ret = t.parseIndentCode()
			} else {
				t.backup()
				ret = t.parseParagraph()
			}
		}
		curNode.Append(ret)
		return
	case itemNewline:
		t.nextToken()
		return
	default:
		ret = t.parseParagraph()
	}
	curNode.Append(ret)
	return
}
