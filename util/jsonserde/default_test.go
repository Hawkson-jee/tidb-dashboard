// Copyright 2022 PingCAP, Inc. Licensed under Apache-2.0.

package jsonserde

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	FooBar  string
	SQLTime time.Time
}

func TestMarshal(t *testing.T) {
	data := testStruct{
		FooBar:  "zoo",
		SQLTime: time.UnixMicro(1641733771580123),
	}
	val, err := Default.Marshal(data)
	require.NoError(t, err)
	require.Equal(t, `{"foo_bar":"zoo","sql_time":1641733771580}`, string(val))
}

func TestUnmarshal(t *testing.T) {
	var data testStruct
	err := Default.Unmarshal([]byte(`{"foo_bar":"zoo","sql_time":1641733771580}`), &data)
	require.NoError(t, err)
	require.Equal(t, "zoo", data.FooBar)
	require.Equal(t, time.UnixMilli(1641733771580), data.SQLTime)
}