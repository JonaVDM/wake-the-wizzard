package wol

import (
	"encoding/hex"
	"errors"
	"net"
	"strings"
)

func SendWol(mac string) error {
	if ok := ValidMac(mac); !ok {
		return errors.New("invalid mac addres")
	}

	msg, err := buildMessage(strings.ReplaceAll(mac, ":", ""))
	if err != nil {
		return err
	}

	pc, err := net.ListenPacket("udp4", ":9")
	if err != nil {
		return err
	}
	defer pc.Close()

	addr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:9")
	if err != nil {
		return nil
	}

	_, err = pc.WriteTo(msg, addr)
	return err
}

func buildMessage(mac string) ([]byte, error) {
	msg := make([]byte, 0)

	// Fill the first 6 bytes with FF
	for i := 0; i < 6; i++ {
		msg = append(msg, byte(255))
	}

	// Convert the mac address to bytes
	macBytes, err := hex.DecodeString(strings.ToLower(mac))
	if err != nil {
		return []byte{}, err
	}

	// Fill in the rest with the bytes of the mac
	for i := 0; i < 16; i++ {
		msg = append(msg, macBytes...)
	}

	return msg, nil
}

func ValidMac(mac string) bool {
	if len(mac) != 17 {
		return false
	}

	if mac[2] != ':' || mac[5] != ':' || mac[8] != ':' || mac[11] != ':' || mac[14] != ':' {
		return false
	}

	for _, c := range mac {
		if c != ':' && (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return false
		}
	}

	return true
}
