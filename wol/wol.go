package wol

import (
	"encoding/hex"
	"net"
	"strings"
)

func SendWol(mac string) error {
	msg, err := buildMessage(mac)
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
