package model

var _ Interface = (*InterfaceImplEmbed)(nil)
var _ Interface = (*InterfaceImplEmbedVal)(nil)

// InterfaceImplEmbed はEmbedを利用してインターフェイスを満足させる構造体
type InterfaceImplEmbed struct {
	InterfaceImpl
}

// InterfaceImplEmbedVal はEmbedを利用してインターフェイスを満足させる構造体（値版）
type InterfaceImplEmbedVal struct {
	InterfaceImplVal
}
