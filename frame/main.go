package BattleFrame

type usr interface {
	input // 用于获取输入信息
	output // 用于获取输出信息
}

type context  interface{
	
}

event()
/*
状态机包括：
1.init状态
2.游戏开始
3.usr1 回合开始
4.usr1 回合中
5.usr1 回合结束
3.usr2 回合开始
4.usr2 回合中
5.usr2 回合结束

init 状态：
#用于初始化上下文，打乱排序
preinit()
#[card]执行卡牌的init钩子
inithook()

#[]先手抽三张牌，后手抽四张牌
#等待双方选牌

*/
// usr 主要用于获取输入和输出
func NewBattleFrame(u1 usr,u2 usr)(context,error)

func Run(u1 usr,u2 usr) {
	// step1 创建对战平台上下文  context
	// 进入init状态
	context.gotostatue(INIT)
	context.run()	
}