package utils

import (
	"github.com/pkg/errors"
	"net"
)

func GetLocalIPs() ([]net.IP, error) {
	var ips []net.IP
	adders, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}
	for _, a := range adders {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP)
		}
	}
	return ips, nil
}

func SumIPs(ips []net.IP) uint16 {
	total := 0
	for _, ip := range ips {
		for i := range ip {
			total += int(ip[i])
		}
	}
	return uint16(total)
}

// LocalMac gets the first NIC's MAC address.
func LocalMac() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, inter := range interfaces {
		address, err := inter.Addrs()
		if err != nil {
			return "", err
		}

		for _, address := range address {
			// check the address type and if it is not a loopback the display it
			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					return inter.HardwareAddr.String(), nil
				}
			}
		}
	}

	return "", errors.New("can't get local mac")
}
