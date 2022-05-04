package matcher

import (
	"encoding/hex"
	"fmt"
	"net"
	"net/netip"
	"testing"

	"github.com/discover255/syslog-go/pkg/config"
)

func TestChangeFirst(t *testing.T) {
	filename := "./sample.yaml"
	config.ReadConfig(&filename)
	fmt.Println(config.GlobaConfigMap.Default.MQ)
}

func TestDecode(t *testing.T) {
	test_case, _ := hex.DecodeString("c8d5d6beb2e2cad4")
	fmt.Println(Decode(test_case, "GBK"))
}

func TestInit(t *testing.T) {
	filename := "./sample.yaml"
	config.ReadConfig(&filename)
	InitFromConfig()
	for ip, pairs := range KwBindings {
		fmt.Printf("KwBindings: %v: %v\n", ip, *pairs)
	}
	for _, test_ip := range []net.IP{
		net.ParseIP("10.0.0.1"),
		net.ParseIP("1.1.1.1"),
	} {
		fmt.Println(MatchEncoding(OldIPtoNewIP(test_ip)))
	}
}

func TestChooseProducer(t *testing.T) {
	tests := []struct {
		Msg  string
		IP   netip.Addr
		Want string
	}{
		{
			Msg:  "<123>slfksjdf sldfjslkd USG",
			IP:   netip.MustParseAddr("127.0.0.1"),
			Want: "default",
		},
		{
			Msg:  "<123>slfksjdf sldfjslkd USG",
			IP:   netip.MustParseAddr("10.0.1.2"),
			Want: "2",
		},
		{
			Msg:  "<12> TEST",
			IP:   netip.MustParseAddr("127.0.0.1"),
			Want: "default",
		},
		{
			Msg:  "<12> blabla",
			IP:   netip.MustParseAddr("127.0.0.1"),
			Want: "1",
		},
	}
	for _, test := range tests {
		fmt.Printf("%v: %v\n", test.IP, *KwBindings[test.IP])
		got := ChooseProducer(test.Msg, &test.IP)
		if got != test.Want {
			t.Fatalf("unexpected error, expected: %s, but got: %s", test.Want, got)
		}
	}
}
