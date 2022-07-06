package log

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	api "github.com/travisjeffery/proglog/api/v1"
)

func TestSegment(t *testing.T) {
	dir, _ := ioutil.TempDir("", "segmenht-test")
	defer os.RemoveAll(dir)

	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = entWidth * 3
	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), s.nextOffset, s.nextOffset)
	require.False(t, s.IsMaxed())
	want := &api.Record{Value: []byte("hello world")}
	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		require.NoError(t, err)
		require.Equal(t, 16+i, off)

		got, err := s.Read(off)
		require.NoError(t, err)
		require.Equal(t, want.Value, got.Value)
	}
}
