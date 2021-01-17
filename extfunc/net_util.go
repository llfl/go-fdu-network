package extfunc

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"time"
	"os/exec"
)

var icmp ICMP

//ICMP packages
type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

//Ping ping
func Ping(ip string) bool {
	//开始填充数据包
	icmp.Type = 8 //8->echo message  0->reply message
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0

	recvBuf := make([]byte, 32)
	var buffer bytes.Buffer

	//先在buffer中写入icmp数据报求去校验和
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.Checksum = checkSum(buffer.Bytes())
	//然后清空buffer并把求完校验和的icmp数据报写入其中准备发送
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)

	Time, _ := time.ParseDuration("2s")
	conn, err := net.DialTimeout("ip4:icmp", ip, Time)
	
	if err != nil {
		return false
	}
	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		log.Println("conn.Write error:", err)
		return false
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	num, err := conn.Read(recvBuf)
	if err != nil {
		log.Println("conn.Read error:", err)
		return false
	}

	conn.SetReadDeadline(time.Time{})
	if string(recvBuf[0:num]) != "" {
		return true
	}
	return false

}

func checkSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}

//CheckServer check online server
func CheckServer(ip string, port string ) bool {
	timeout := time.Duration(2 * time.Second)
	addr := ip + ":" + port
	_, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		return false
	}
	return true
}

//CheckDNS CheckDNS txtrecords
func CheckDNS(addr string) bool {
	txtrecords, _ := net.LookupTXT(addr)
	if txtrecords[0] == "ok" {
		return true
	}
	return false

}

//SystemPing call system ping
func SystemPing(addr string) bool {
	cmd := exec.Command("ping", addr, "-c", "1", "-W", "5")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
