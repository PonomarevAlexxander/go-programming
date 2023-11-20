package wifi

import (
	"fmt"
	"net"

	"github.com/mdlayher/wifi"
)

type WiFi interface {
	Interfaces() ([]*wifi.Interface, error)
}

type Service struct {
	WiFi WiFi
}

func New(wifi WiFi) Service {
	return Service{WiFi: wifi}
}

func (service Service) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get Interfaecs: %w", err)
	}

	addrs := make([]net.HardwareAddr, len(interfaces))

	for index, iface := range interfaces {
		addrs[index] = iface.HardwareAddr
	}

	return addrs, nil
}

func (service Service) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get Interfaecs: %w", err)
	}

	nameList := make([]string, len(interfaces))

	for index, iface := range interfaces {
		nameList[index] = iface.Name
	}

	return nameList, nil
}
