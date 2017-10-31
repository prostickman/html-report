// Copyright 2015 ThoughtWorks, Inc.

// This file is part of getgauge/html-report.

// getgauge/html-report is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// getgauge/html-report is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with getgauge/html-report.  If not, see <http://www.gnu.org/licenses/>.

package testHelper

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

var re = regexp.MustCompile("[\\s]*[\n\t][\\s]*")

func RemoveNewline(s string) string {
	return re.ReplaceAllLiteralString(s, "")
}

func AssertEqual(expected, actual, testName string, t *testing.T) {
	if expected != actual {
		if err := CheckEqual(testName, expected, actual); err!=nil {
			t.Error(err)
		}
	}
}

func CheckEqual(test string, want, got interface{}) error {
	if diff := pretty.Compare(got, want); diff != "" {
		return fmt.Errorf("Test:%s\n diff: (-got +want)\n%s", test, diff)
	}
	return nil
}
