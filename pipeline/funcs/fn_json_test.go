// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package funcs

import (
	"testing"

	tu "gitlab.jiagouyun.com/cloudcare-tools/cliutils/testutil"
)

func TestJSON(t *testing.T) {
	testCase := []*funcCase{
		{
			data: `{
			  "name": {"first": "Tom", "last": "Anderson"},
			  "age":37,
			  "children": ["Sara","Alex","Jack"],
			  "fav.movie": "Deer Hunter",
			  "friends": [
			    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
			    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
			    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
			  ]
			}`,
			script:   `json(_, name) json(name, first)`,
			expected: "Tom",
			key:      "first",
		},
		{
			data: `[
				    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
				    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
				    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
				]`,
			script:   `json(_, .[0].nets[-1])`,
			expected: "tw",
			key:      "[0].nets[-1]",
		},
	}

	for idx, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			runner, err := NewTestingRunner(tc.script)
			tu.Equals(t, nil, err)

			err = runner.Run(tc.data)
			tu.Equals(t, nil, err)

			ret := runner.Result()

			r, ok := ret.Fields[tc.key]
			tu.Equals(t, true, ok)
			tu.Equals(t, tc.expected, r)

			t.Logf("[%d] PASS", idx)
		})
	}
}
