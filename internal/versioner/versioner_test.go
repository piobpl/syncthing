// Copyright (C) 2014 The Syncthing Authors.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along
// with this program. If not, see <http://www.gnu.org/licenses/>.

package versioner

import "testing"

func TestTaggedFilename(t *testing.T) {
	cases := [][3]string{
		{"foo/bar.baz", "tag", "foo/bar~tag.baz"},
		{"bar.baz", "tag", "bar~tag.baz"},
		{"bar", "tag", "bar~tag"},

		// Parsing test only
		{"", "tag-only", "foo/bar.baz~tag-only"},
		{"", "tag-only", "bar.baz~tag-only"},
	}

	for _, tc := range cases {
		if tc[0] != "" {
			// Test tagger
			tf := taggedFilename(tc[0], tc[1])
			if tf != tc[2] {
				t.Errorf("%s != %s", tf, tc[2])
			}
		}

		// Test parser
		tag := filenameTag(tc[2])
		if tag != tc[1] {
			t.Errorf("%s != %s", tag, tc[1])
		}
	}
}
