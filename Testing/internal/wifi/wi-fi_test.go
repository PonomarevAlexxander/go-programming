package wifi_test

import (
	"errors"
	myWifi "go/course/example_mock/internal/wifi"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .
type rawWifi struct {
	addr string
	name string
}
type rowTestSysInfo struct {
	wifis []rawWifi
	errExpected error
}

var testTable = []rowTestSysInfo{
		{
		wifis: []rawWifi{
			{addr: "00:11:22:33:44:55", name: "first"},
			{addr: "aa:bb:cc:dd:ee:ff", name: "second"},
		},
	},
	{
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetAdresses(t *testing.T) {
	for i, test := range testTable {
		mockWifi := NewWiFi(t)
		wifiService := myWifi.Service{WiFi: mockWifi}

		mockWifi.On("Interfaces").Return(mockIfaces(test.wifis), test.errExpected)
		actualAddrs, err := wifiService.GetAddresses()

		if test.errExpected != nil {
			require.ErrorIs(t, err, test.errExpected, "row: %d, expected error: %w, actual error: %w", i, test.errExpected, err)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		expectedMacs := parseMACs(test.wifis)
		require.Equal(t, expectedMacs, actualAddrs, "row: %d, expected addrs: %s, actual addrs: %s", i, expectedMacs, actualAddrs)
	}
}

func TestGetNames(t *testing.T) {
	for i, test := range testTable {
		mockWifi := NewWiFi(t)
		wifiService := myWifi.Service{WiFi: mockWifi}

		mockWifi.On("Interfaces").Return(mockIfaces(test.wifis), test.errExpected)
		actualNames, err := wifiService.GetNames()

		if test.errExpected != nil {
			require.ErrorIs(t, err, test.errExpected, "row: %d, expected error: %w, actual error: %w", i, test.errExpected, err)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		expectedNames := getNamesFromWifis(test.wifis)
		require.Equal(t, expectedNames, actualNames, "row: %d, expected name: %s, actual name: %s", i, expectedNames, actualNames)
	}
}

func mockIfaces(wifis []rawWifi) []*wifi.Interface {
	var interfaces []*wifi.Interface
	for i, element := range wifis {
		hwAddr := parseMAC(element.addr)

		if hwAddr == nil {
			continue
		}

		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         element.name,
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		interfaces = append(interfaces, iface)
	}

	return interfaces
}

func getNamesFromWifis(wifis []rawWifi) []string {
	nameList := make([]string, len(wifis))

	for index, wifi := range wifis {
		nameList[index] = wifi.name
	}

	return nameList
}

func parseMACs(wifis []rawWifi) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, wifi := range wifis {
		addrs = append(addrs, parseMAC(wifi.addr))
	}

	return addrs
}

func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}

	return hwAddr
}
