package model

var _ Interface = (*InterfaceImplEmbed)(nil)
var _ Interface = (*InterfaceImplEmbedRaw)(nil)

// InterfaceImplEmbed はEmbedを利用してインターフェイスを満足させる構造体
type InterfaceImplEmbed struct {
	InterfaceImpl
}

// InterfaceImplEmbedRaw はEmbedを利用してインターフェイスを満足させる構造体（値版）
type InterfaceImplEmbedRaw struct {
	InterfaceImplRaw
}
