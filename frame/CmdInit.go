package BattleFrame
import {
	"time"
}
/*
type Cmd interface {
	// 检查命令合法性
	func Checkout() bool
	func Run() error
}

*/
type CmdInit struct {
	para   string
	engine Engine
}

func (n CmdInit) Checkout() bool {
	return n.engine != nil && n.engine.GetStatue() == BattleFrame.GameInit
}

func (n CmdInit) Run() error {

	// 决定先后手
	n.engine.RandAttacker()

	// 随机化卡牌信息
	n.engine.GetAttacker(Intitate).RandDeck()
	n.engine.GetAttacker(Gote).RandDeck()

	// 添加抽牌命令发送给各个
	n.engine.AddCmd(new CmdInitDrawCard(u:n.engine.GetAttacker(Intitate)))
	n.engine.AddCmd(new CmdInitDrawCard(u:n.engine.GetAttacker(Gote)))
	
	// 等待用户选择结束  1s一次查询双用户的状态是否均变为选择完成
	n.engine.AddCmd(new CmdInitWaiteChoose(1,n.engine))
	
}
// 先手抽三张牌，后手抽四张牌
type CmdInitDrawCard struct {
	u Usr
}

func (n CmdInitDrawCard) Checkout() bool {
	return true
}
// 安全性健壮性后面再说
func (n CmdInitDrawCard) Run() error {
	usr := n.u
	deck,_ := usr.GetDeck()
	cards,_ := deck.GetCard((usr.GetUsrInfo() == Intitate)? 3:4)
	usr.SetChoseCard(cards)
	usr.Send(Init_Choseing)
}

// 选择完成之后  添加GameBegin的命令
type CmdInitWaiteChoose struct {
	delaytime int16
	engine Engine
}
func (n CmdInitWaiteChoose) Checkout() bool {
	return true
}
func (n CmdInitWaiteChoose) Run() error {
	// TODO:延迟个1s
	time.Sleep(time.Second * delaytime)
	
	// 如果双用户均选择完成，则进入游戏开始命令中，否则添加自己
	u_Intitate := n.engine.GetAttacker(Intitate)
	u_Gote := n.engine.GetAttacker(Gote)
	if((u_Intitate.GetUsrSta() == Init_Choseing)
		||(u_Gote.GetUsrSta() == Init_Choseing)){
			n.engine.AddCmd(n)
			return nil
		}
	// TODO: 进入游戏开始阶段， 这个阶段后手将执行添加水晶币操作
	// 执行各个卡牌的游戏开始阶段的处理内容
	n.engine.AddCmd(new CmdGameBegin(n))
}
