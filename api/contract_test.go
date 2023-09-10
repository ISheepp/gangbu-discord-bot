package api

import (
	"gangbu/pkg/util"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestGeneratePrivate(t *testing.T) {
	private, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	text := private.D.Text(16)
	// 将十六进制字符串转换为 big.Int
	t.Logf("%s", text)
	trans, err := util.StringToPrivateKey("0x" + text)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(trans.D.Text(16))
}

func TestKec(t *testing.T) {
	event := "RequestSent(uint256,address)"
	s := crypto.Keccak256Hash([]byte(event)).String()
	t.Log(s)
}
