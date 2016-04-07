package hal

/*
 * Copyright 2016 Albert P. Tobey <atobey@netflix.com>
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

import (
	"fmt"
	"testing"
)

func TestAsciiTable(t *testing.T) {
	samples := [][][]string{
		{
			{"hdr"},
			{"one"},
		},
		{
			{"hdr"},
			{"one"},
			{"two"},
		},
		{
			{"left", "right"},
			{"one", "three"},
			{"two"},
		},
		{
			{"HEADER 1", "HDR 2", "LOL WUT"},
			{"one", "two", "three"},
			{"four", "five", "six"},
		},
		{
			{"Col 1", "Col 2", "3rd Column", "4th", "FIFTH"},
			{"one", "two", "three"},
			{"four", "five", "six"},
			{"hi"},
			{"", "", "", "-", "+"},
		},
	}

	for _, sample := range samples {
		// first row is the header, the rest is data rows
		out := AsciiTable(sample[0], sample[1:])
		// not a very useful test ... yet
		if len(out) == 0 {
			t.Fail()
		}

		fmt.Println(out)
	}
}
