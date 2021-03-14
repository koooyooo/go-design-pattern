package pc

import "fmt"

type (
	// pcCase はPCのケースです。外部からはPCとして認識されます。
	pcCase struct {
		isOpen          bool
		motherBoard     *motherBoard
		powerSupplyUnit *powerSupplyUnit
		ssd             *ssd
	}

	// powerSupplyUnit は電源ユニットです。
	powerSupplyUnit struct {
		powerCapacity int
	}

	// motherBord はマザーボードです。
	motherBoard struct {
		*cpu
		*memory
	}

	// cpu はCPUです。コア数を持ちます。
	cpu struct {
		core int
	}

	// memory はメモリです。
	memory struct {
		storageGb int
	}

	// ssd はSSDです。
	ssd struct {
		storageGB int
	}
)

// open はPCケースを開けます
func (c *pcCase) open() {
	c.isOpen = true
}

// close はPCケースを閉じます
func (c *pcCase) close() {
	c.isOpen = false
}

func (c pcCase) Spec() string {
	return fmt.Sprintf("Power: %d, CPU: %d, Mem: %d, SSD: %d", c.powerSupplyUnit.powerCapacity, c.motherBoard.cpu.core, c.motherBoard.memory.storageGb, c.ssd.storageGB)
}

// removeCPUCover はCPUカバーを開けます
func (m *motherBoard) removeCPUCover() {
	// Do Something
}

// openCPUSocket はCPUソケットを開けます
func (m *motherBoard) openCPUSocket() {
	// Do Something
}

// closeCPUSocket はCPUソケットを閉じます
func (m *motherBoard) closeCPUSocket() {
	// Do Something
}

// setCPUCooler はCPUクーラーを設置します
func (m *motherBoard) setCPUCooler() {
	// Do Something
}

// openMemorySlot はメモリースロットを開けます
func (m *motherBoard) openMemorySlot() {
	// Do Something
}

// closeMemorySlot はメモリースロットを閉じます
func (m *motherBoard) closeMemorySlot() {
	// Do Something
}
