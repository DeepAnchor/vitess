package eventtoken

import (
	"testing"

	"github.com/golang/protobuf/proto"
	querypb "github.com/youtube/vitess/go/vt/proto/query"
)

func TestMinimum(t *testing.T) {
	testcases := []struct {
		ev1      *querypb.EventToken
		ev2      *querypb.EventToken
		expected *querypb.EventToken
	}{{
		ev1:      nil,
		ev2:      nil,
		expected: nil,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 123,
		},
		ev2:      nil,
		expected: nil,
	}, {
		ev1: nil,
		ev2: &querypb.EventToken{
			Timestamp: 123,
		},
		expected: nil,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 123,
		},
		ev2: &querypb.EventToken{
			Timestamp: 456,
		},
		expected: &querypb.EventToken{
			Timestamp: 123,
		},
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 456,
		},
		ev2: &querypb.EventToken{
			Timestamp: 123,
		},
		expected: &querypb.EventToken{
			Timestamp: 123,
		},
	}}

	for _, tcase := range testcases {
		got := Minimum(tcase.ev1, tcase.ev2)
		if tcase.expected == nil && got != nil {
			t.Errorf("expected nil result for Minimum(%v, %v) but got: %v", tcase.ev1, tcase.ev2, got)
			continue
		}
		if !proto.Equal(got, tcase.expected) {
			t.Errorf("got %v but expected %v for Minimum(%v, %v)", got, tcase.expected, tcase.ev1, tcase.ev2)
		}
	}
}

func TestFresher(t *testing.T) {
	testcases := []struct {
		ev1      *querypb.EventToken
		ev2      *querypb.EventToken
		expected int
	}{{
		// Test cases with no shards.
		ev1:      nil,
		ev2:      nil,
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 123,
		},
		ev2:      nil,
		expected: -1,
	}, {
		ev1: nil,
		ev2: &querypb.EventToken{
			Timestamp: 123,
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 123,
		},
		ev2: &querypb.EventToken{
			Timestamp: 123,
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 200,
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
		},
		expected: 100,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
		},
		ev2: &querypb.EventToken{
			Timestamp: 200,
		},
		expected: -100,
	}, {
		// Test cases with not enough information to compare.
		ev1: &querypb.EventToken{
			Timestamp: 100,
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s2",
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "pos1",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "pos2",
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "pos1", // invalid on purpose
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "pos2", // invalid on purpose
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MariaDB/0-1-123", // valid but different
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MySQL56/33333333-3333-3333-3333-333333333333:456-789",
		},
		expected: -1,
	}, {
		// MariaDB test cases.
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MariaDB/0-1-200",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MariaDB/0-1-100",
		},
		expected: 1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MariaDB/0-1-100",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MariaDB/0-1-200",
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MariaDB/0-1-100",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MariaDB/0-1-100",
		},
		expected: 0,
	}, {
		// MySQL56 test cases.
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MySQL56/33333333-3333-3333-3333-333333333333:1-200",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MySQL56/33333333-3333-3333-3333-333333333333:1-100",
		},
		expected: 1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MySQL56/33333333-3333-3333-3333-333333333333:1-100",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MySQL56/33333333-3333-3333-3333-333333333333:1-200",
		},
		expected: -1,
	}, {
		ev1: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MySQL56/33333333-3333-3333-3333-333333333333:1-100",
		},
		ev2: &querypb.EventToken{
			Timestamp: 100,
			Shard:     "s1",
			Position:  "MySQL56/33333333-3333-3333-3333-333333333333:1-100",
		},
		expected: 0,
	}}

	for _, tcase := range testcases {
		got := Fresher(tcase.ev1, tcase.ev2)
		if got != tcase.expected {
			t.Errorf("got %v but expected %v for Fresher(%v, %v)", got, tcase.expected, tcase.ev1, tcase.ev2)
		}
	}
}
