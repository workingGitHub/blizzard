package usragent

type usragent interface {
	// 本身存在一个channel  输入的channel，info信息
	func GetChannel() (chan<- string,error)
	// 获取卡组信息
	func GetDeck()(uint64,error)
	// 获取英雄的信息
	func GetHero()(Hero,error)
}
