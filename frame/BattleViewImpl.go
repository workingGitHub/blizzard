package BattleFrame

const (
	System Role = iota
	First
	Secend
)

// 时机类型和消息类型还是不宜样的
/* 时机类型包括了各种各样的处理时机
比如：
游戏初始化
游戏开始
回合开始
回合中
回合结束
*/
/*  这个状态及不要了
GameNotReady ->  GameInit:两个对战用户添加完成
GameInit -> GameBegin: 用户选牌结束
GameBegin -> F_RoundBegin: 自动的，执行完Begin执行这个
const (
	GameError StatueInfo = iota
	GameNotReady
	GameInit
	GameBegin
	F_RoundBegin
	F_Rounding
	F_RoundEnd
	S_RoundBegin
	S_Rounding
	S_RoundEnd
	Selecting
)
*/

/*
消息处理是动作
比如：
抽一张牌
随从A受到了攻击
随从A受伤
随从A死亡
随从A被沉默
随从A攻击随从B（前处理）
水晶临时+1，永久+1等
额外一个回合
*/

type BattleViewImpl struct {
	// 用户名信息
	usr1 Usr
	usr2 Usr
	//消息堆栈
	cmd []Cmd
}

func NewBattleFrame(u1 Usr, u2 Usr) (Engine, error) {
	// TODO:实例化BattleiewImpl
}

// TODO：消息处理方式要考虑清楚
func (u BattleViewImpl) Run() error {
	// 创建初始化命令
	cmd = new
}

/*
卡牌定义：
触发条件：
[onbattle/ondeck/onhand][msgtype][para]
触发动作：
[影响范围]
[影响动作]
*/
