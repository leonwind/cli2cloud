package session_client

import (
	"flag"
	"github.com/leonwind/cli2cloud/service/serverpb"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	ip = flag.String("ip", ":8080", "")
)

func TestMain(m *testing.M)  {
	flag.Parse()
	ret := m.Run()
	os.Exit(ret)
}
func TestPubSub(t *testing.T) {
	pubC, err := NewPubClient(*ip)
	require.NoError(t, err)
	sessionID := pubC.GetSessionID()
	subC, err := NewSubClient(*ip)
	require.NoError(t, err)
	outputCh := make(chan *serverpb.Output, 256)
	stopCh := make(chan bool, 1)
	go func() {
		err := subC.Subscribe(sessionID, outputCh, stopCh)
		require.NoError(t, err)
	}()
	for i := 0; i < 20; i++ {
		err = pubC.Publish(&serverpb.Output{Num: uint32(i), Content: "Hello"})
		require.NoError(t, err)
	}
	for i := 0; i < 20; i++ {

	}
}