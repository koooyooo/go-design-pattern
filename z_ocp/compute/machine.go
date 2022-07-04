package compute

import (
	"log"
	"time"
)

// (変数・メソッドを privateにすることで閉鎖を実現)
type Machine struct {
	powered     bool
	cpu         CPU
	memory      Memory
	hd          HardDrive
	usbDevices  []USBDevice
	hdmiDevices []HDMIDevice

	startUp  func(*Machine) error
	shutDown func(*Machine) error
}

// マシンの電源スイッチを押下
// (閉鎖対象を守るためアクセス経路を限定)
func (m *Machine) PushPowerSwitch(d time.Duration) {
	if !m.powered {
		if err := m.startUp(m); err != nil {
			log.Fatal(err)
		}
	} else {
		if (5 * time.Second) < d {
			if err := m.shutDown(m); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (m *Machine) defaultStartUp() error {
	if err := m.cpu.init(); err != nil {
		return err
	}
	if err := m.memory.init(); err != nil {
		return err
	}
	if err := m.hd.init(); err != nil {
		return err
	}
	for _, u := range m.usbDevices {
		if err := u.init(); err != nil {
			return err
		}
	}
	for _, h := range m.hdmiDevices {
		if err := h.init(); err != nil {
			return err
		}
	}
	m.powered = true
	return nil
}

func (m *Machine) defaultShutDown() error {
	if err := m.cpu.finalize(); err != nil {
		return err
	}
	if err := m.memory.finalize(); err != nil {
		return err
	}
	if err := m.hd.finalize(); err != nil {
		return err
	}
	for _, u := range m.usbDevices {
		if err := u.finalize(); err != nil {
			return err
		}
	}
	for _, h := range m.hdmiDevices {
		if err := h.finalize(); err != nil {
			return err
		}
	}
	m.powered = false
	return nil
}

func (m *Machine) MountUSB(u USBDevice) error {
	if err := u.init(); err != nil {
		return err
	}
	m.usbDevices = append(m.usbDevices, u)
	return nil
}

func (m *Machine) MountHDMI(h HDMIDevice) error {
	if err := h.init(); err != nil {
		return err
	}
	m.hdmiDevices = append(m.hdmiDevices, h)
	return nil
}
