package syslog

// import (
// 	"fmt"
// 	"testing"

// 	// "github.com/influxdata/go-syslog/v3/rfc3164"
// 	"github.com/influxdata/go-syslog/v3/rfc3164"
// 	"github.com/influxdata/go-syslog/v3/rfc5424"
// )

// // func TestSendSyslog(t *testing.T) {
// // 	sysLog, err := syslog.Dial("udp", "localhost:30514",
// // 		syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	fmt.Fprintf(sysLog, "This is a daemon warning with demotag.")
// // 	sysLog.Emerg("And this is a daemon emergency with demotag.")
// // }

// func TestBuilder(t *testing.T) {
// 	msg := &rfc5424.SyslogMessage{}
// 	msg.SetTimestamp("not a RFC3339MICRO timestamp")
// 	fmt.Println("Valid?", msg.Valid())
// 	msg.SetPriority(191)
// 	msg.SetVersion(1)
// 	fmt.Println("Valid?", msg.Valid())
// 	str, _ := msg.String()
// 	fmt.Println(str)
// }

// func TestParse(t *testing.T) {
// 	i := []byte(`<190>May  1 2022 03:22:22 HZ_916_SWITCH_USG9560_FW_02 %%01SECIF/6/STREAM(l)[342]:In Last Five Minutes Stream Statistic is :IF1-GE0/0/0,STATE-U,IN-1,OUT-0,IF2-GE1/1/0,STATE-U,IN-0,OUT-0,IF3-GE1/1/1,STATE-U,IN-0,OUT-0,IF4-GE1/1/2,STATE-U,IN-0,OUT-0,IF5-GE1/1/3,STATE-U,IN-0,OUT-0,IF6-GE1/1/4,STATE-U,IN-1181,OUT-5,IF7-GE1/1/5,STATE-U,IN-1178,OUT-0.`)
// 	p := rfc3164.NewParser(rfc3164.WithBestEffort())
// 	m, _ := p.Parse(i)
// 	fmt.Println(m.(*rfc3164.SyslogMessage))
// 	msg := m.(*rfc3164.SyslogMessage)
// 	fmt.Println(*msg.Message)
// }
