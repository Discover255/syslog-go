package matcher

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/netip"
	"strings"

	"github.com/discover255/syslog-go/pkg/config"
	"github.com/discover255/syslog-go/pkg/mqclient"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var DecoderMap map[string]*encoding.Decoder

func init() {
	DecoderMap = make(map[string]*encoding.Decoder)
	DecoderMap["gbk"] = simplifiedchinese.GBK.NewDecoder()
	DecoderMap["gb2312"] = simplifiedchinese.HZGB2312.NewDecoder()
	DecoderMap["gb18030"] = simplifiedchinese.GB18030.NewDecoder()
	IPEncodingMap = make(map[netip.Addr]string)
	GlobalKwBindings = []KwProdPair{}
	KwBindings = make(map[netip.Addr]*[]KwProdPair)
	DefaultProducer = "default"
	DefaultEncoding = "utf-8"
}

func Decode(b []byte, encoding string) string {
	encoding = strings.ToLower(encoding)
	if encoding == "utf8" || encoding == "utf-8" {
		return string(b)
	} else {
		reader := transform.NewReader(bytes.NewReader(b), DecoderMap[encoding])
		utf8Bytes, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("Decode Failed %v, Encoding: %s", err, encoding)
		}
		return string(utf8Bytes)
	}
}

var IPEncodingMap map[netip.Addr]string

func OldIPtoNewIP(ip net.IP) *netip.Addr {
	ipv4 := ip.To4()
	if ipv4 != nil {
		ip = ipv4
	}
	if newip, ok := netip.AddrFromSlice(ip); ok {
		return &newip
	} else {
		fmt.Printf("IP Conversion is failed")
		return nil
	}
}

func InitFromConfig() {
	configEncoding := config.GlobaConfigMap.Default.Encoding
	if configEncoding != "" {
		DefaultEncoding = configEncoding
	}
	configMQ := config.GlobaConfigMap.Default.MQ
	if configMQ != "" {
		DefaultProducer = configMQ
	}
	for _, rule := range config.GlobaConfigMap.Rules {
		fmt.Println(rule)
		if len(rule.Hosts) == 0 && len(rule.Keywords) > 0 {
			for _, keyword := range rule.Keywords {
				AddBinding(nil, keyword, rule.MQs[0])
			}
		} else if len(rule.Hosts) > 0 {
			for _, host := range rule.Hosts {
				ip := netip.MustParseAddr(host)
				if rule.Encoding != "" {
					IPEncodingMap[ip] = strings.ToLower(rule.Encoding)
				}
				if len(rule.MQs) == 0 {
					continue
				}
				if len(rule.Keywords) == 0 {
					AddBinding(&ip, "", rule.MQs[0])
				} else {
					for _, keyword := range rule.Keywords {
						AddBinding(&ip, keyword, rule.MQs[0])
					}
				}
			}
		}
	}
}

func AddBinding(ip *netip.Addr, keyword string, mq string) {
	// fmt.Printf("AddBinding: %v, %s, %s\n", ip, keyword, mq)
	if ip == nil {
		GlobalKwBindings = append(GlobalKwBindings, KwProdPair{Keyword: keyword, ProducerName: mq})
		return
	}
	if _, ok := KwBindings[*ip]; !ok {
		KwBindings[*ip] = &[]KwProdPair{}
	}
	pairs := KwBindings[*ip]
	new_pairs := append(*pairs, KwProdPair{Keyword: keyword, ProducerName: mq})
	KwBindings[*ip] = &new_pairs
}

func MatchEncoding(ip *netip.Addr) string {
	if coding, ok := IPEncodingMap[*ip]; ok {
		return coding
	} else {
		return DefaultEncoding
	}
}

type KwProdPair struct {
	Keyword      string
	ProducerName string
}

var KwBindings map[netip.Addr]*[]KwProdPair
var GlobalKwBindings []KwProdPair
var DefaultEncoding string
var DefaultProducer string

func ChooseProducer(msg string, ip *netip.Addr) string {
	if pairs, ok := KwBindings[*ip]; ok {
		for _, pair := range *pairs {
			if pair.Keyword == "" {
				return pair.ProducerName
			}
			if strings.Contains(msg, pair.Keyword) {
				return pair.ProducerName
			}
		}
	}
	for _, pair := range GlobalKwBindings {
		if strings.Contains(msg, pair.Keyword) {
			return pair.ProducerName
		}
	}
	return DefaultProducer
}

func MatchAndSend(rawData []byte, remoteAddr *net.UDPAddr) {
	ip := OldIPtoNewIP(remoteAddr.IP)
	coding := MatchEncoding(ip)
	msg := Decode(rawData, coding)
	fmt.Println(msg)
	producer := ChooseProducer(msg, ip)
	mqclient.SendToPulsar(producer, msg)
}
