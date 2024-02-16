package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Tagged(t *testing.T) {
	set := TaggedValueSet[int]{
		NewTaggedValue(1, "one"),
		NewTaggedValue(2, "two", "second"),
		NewTaggedValue(3, "three", "third"),
		NewTaggedValue(23, "twenty-three", "twenty", "third"),
		NewTaggedValue(4, "four", ""),
		NewTaggedValue(9, "nine"),
	}

	tests := []struct {
		name     string
		keep     []string
		remove   []string
		expected []int
	}{
		{
			name:     "by tag",
			keep:     arr("two"),
			expected: arr(2),
		},
		{
			name:     "by multiple",
			keep:     arr("one", "third"),
			expected: arr(1, 3, 23),
		},
		{
			name:     "nil keep",
			keep:     nil,
			expected: nil,
		},
		{
			name:     "empty keep",
			keep:     []string{},
			expected: nil,
		},
		{
			name:     "remove by tag",
			keep:     arr("one", "twenty-three"),
			remove:   arr("third"),
			expected: arr(1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := set.Select(test.keep...).Remove(test.remove...)
			if test.expected == nil {
				require.Empty(t, got)
				return
			}
			require.ElementsMatch(t, test.expected, got.Values())
		})
	}
}

func Test_Tags(t *testing.T) {
	values := TaggedValueSet[int]{
		NewTaggedValue(10, "one", "zero", "ten"),
		NewTaggedValue(11, "one", "eleven"),
		NewTaggedValue(20, "two", "zero", "twenty"),
		NewTaggedValue(22, "two", "twenty two"),
	}

	expected := []string{"one", "zero", "ten", "eleven", "two", "twenty", "twenty two"}
	require.Equal(t, expected, values.Tags())
}

func Test_twoStepFilter(t *testing.T) {
	tag := func(name string, tags ...string) TaggedValue[string] {
		return NewTaggedValue(name, append([]string{name}, tags...)...)
	}
	all := TaggedValueSet[string]{
		tag("js-1", "i", "js"),
		tag("js-2", "d", "js"),
		tag("js-3", "i", "d", "js"),
		tag("jv-1", "i", "d", "jv"),
		tag("jv-2", "i", "d", "jv"),
		tag("py-1", "i", "d", "py"),
		tag("py-2", "d", "py"),
		tag("py-3", "i", "d", "py"),
		tag("py-4", "i", "py"),
		tag("sc-1"),
	}

	tests := []struct {
		name     string
		req      filterRequest
		expected []string
	}{
		{
			name: "--override-default-catalogers i",
			req: filterRequest{
				Base: str("i"),
			},
			expected: str("js-1", "js-3", "jv-1", "jv-2", "py-1", "py-3", "py-4"),
		},
		{
			name: "--override-default-catalogers i,js",
			req: filterRequest{
				Base: str("i", "js"),
			},
			expected: str("js-1", "js-2", "js-3", "jv-1", "jv-2", "py-1", "py-3", "py-4"),
		},
		{
			name: "--select-catalogers “+javascript”  [ERROR]",
			req: filterRequest{
				Base: str("i"),
				Add:  str("js"),
			},
			expected: str("js-1", "js-3", "jv-1", "jv-2", "py-1", "py-3", "py-4", "js-2"),
		},
		{
			name: "--select-catalogers +sc-1",
			req: filterRequest{
				Add: str("sc-1"),
			},
			expected: str("js-1", "js-2", "js-3", "jv-1", "jv-2", "py-1", "py-2", "py-3", "py-4", "sc-1"),
		},
		{
			name: "--select-catalogers -py-1",
			req: filterRequest{
				Remove: str("py-1"),
			},
			expected: str("js-1", "js-2", "js-3", "jv-1", "jv-2", "py-2", "py-3", "py-4", "sc-1"),
		},
		{
			name: "--select-catalogers js",
			req: filterRequest{
				Select: str("js"),
			},
			expected: str("js-1", "js-2", "js-3"),
		},
		{
			name: "--override-default-catalogers d --select-catalogers py,js",
			req: filterRequest{
				Base:   str("d"),
				Select: str("py", "js"),
			},
			expected: str("js-2", "js-3", "py-1", "py-2", "py-3"),
		},
		{
			name: "--select-catalogers -py-1,-py-2,+sc-1,+js-2",
			req: filterRequest{
				Base:   str("i"),
				Add:    str("sc-1", "js-2"),
				Remove: str("py-1", "py-2"),
			},
			expected: str("js-1", "js-3", "jv-1", "jv-2", "py-3", "py-4", "js-2", "sc-1"),
		},
		{
			name: "--select-catalogers js,py",
			req: filterRequest{
				Base:   str("i"),
				Select: str("js", "py"),
			},
			expected: str("js-1", "js-3", "py-1", "py-3", "py-4"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := twoLevelFilter(all, test.req)
			require.ElementsMatch(t, test.expected, got)
		})
	}
}

type filterRequest struct {
	Base   []string
	Select []string
	Remove []string
	Add    []string
}

func twoLevelFilter[T comparable](allValues TaggedValueSet[T], r filterRequest) []T {
	values := allValues
	if len(r.Base) > 0 {
		values = values.Select(r.Base...)
	}
	if len(r.Select) > 0 {
		values = values.Select(r.Select...)
	}
	if len(r.Remove) > 0 {
		values = values.Remove(r.Remove...)
	}
	if len(r.Add) > 0 {
		values = values.Join(allValues.Select(r.Add...)...)
	}
	return values.Values()
}

func str(s ...string) []string {
	return s
}

func arr[T any](v ...T) []T {
	return v
}
