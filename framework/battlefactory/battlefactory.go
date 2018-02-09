package battlefactory

/*
用于生成usrbattle的工厂信息
*/
type battlefactory interface {
	GetBattle() usrBattle // 每个用户的界面信息
}
type battlefactoryImpl struct {
}

var (
	BFactory battlefactoryImpl
)

func init() {
	// 初始化battlefactory
}
