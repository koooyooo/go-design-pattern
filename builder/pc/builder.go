package pc

import "errors"

// pcBuilder はPCのBuilderです
type pcBuilder struct {
	PC  *pcCase
	Err error
}

// NewPCBuilder はpcBuilderのコンストラクタ関数です
func NewPCBuilder() *pcBuilder {
	return &pcBuilder{
		PC:  &pcCase{},
		Err: nil,
	}
}

// SetUpBaseUnit はPCの基本構成を構築します
func (b *pcBuilder) SetUpBaseUnit(powerCapacity int) *pcBuilder {
	// PCケースを用意
	b.PC = &pcCase{
		isOpen: false,
	}
	// PCケースを開け
	b.PC.open()
	// マザーボードを設置
	b.PC.motherBoard = &motherBoard{}
	// 電源ユニットを設置
	b.PC.powerSupplyUnit = &powerSupplyUnit{
		powerCapacity: powerCapacity,
	}
	return b
}

// SetCPU はCPUを設定します
func (b *pcBuilder) SetCPU(core int) *pcBuilder {
	if b.PC == nil || b.PC.motherBoard == nil {
		b.Err = errors.New("case or motherboard is not set")
		return b
	}
	if !b.PC.isOpen {
		b.PC.open()
	}
	mb := b.PC.motherBoard
	// CPU設置作業
	mb.removeCPUCover()
	mb.openCPUSocket()
	mb.cpu = &cpu{core: core}
	mb.closeCPUSocket()
	mb.setCPUCooler()
	return b
}

// SetMemory はMemoryを設定します
func (b *pcBuilder) SetMemory(storageGB int) *pcBuilder {
	if b.PC == nil || b.PC.motherBoard == nil {
		b.Err = errors.New("case or motherboard is not set")
		return b
	}
	if !b.PC.isOpen {
		b.PC.open()
	}
	mb := b.PC.motherBoard
	// Memory設置作業
	mb.openMemorySlot()
	mb.memory = &memory{storageGb: storageGB}
	mb.closeMemorySlot()
	return b
}

// SetSSD はSSDを設定します
func (b *pcBuilder) SetSSD(storageGB int) *pcBuilder {
	if b.PC == nil || b.PC.motherBoard == nil {
		b.Err = errors.New("case or motherboard is not set")
		return b
	}
	if !b.PC.isOpen {
		b.PC.open()
	}
	// SSDを設置
	b.PC.ssd = &ssd{
		storageGB: storageGB,
	}
	return b
}

// Buildは構築済みのPCを返します
func (b *pcBuilder) Build() *pcCase {
	return b.PC
}
