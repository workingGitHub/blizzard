package BattleFrame

/*
对战需要几个功能
1.用户输入,输入需要能够解析
	func input(string) error
   主要输入有如下几个：
	#1.获取实时对战情况
	1.选择命令 参数（1，2，3，4）
	2.单选  参数（1～3）
	2.用户操作信息，打出某一张牌（先手：1～10  后手：11～20）（参数）
	3.操作某个随从或者自己进行攻击（先手：1～8 后手：11～18）（参数）
	4.使用用户技能（先手：1  后手：11）

2.输出到用户中。
	1.状态输出：
		告诉用户你的回合开始
		游戏开始时选择状态，参数是（先手/后手，卡牌列表） 状态为选牌阶段
		游戏过程中的选择状态，参数是（最多三种状态）
		游戏过程中对决（显示两张卡），确定输赢
		游戏过程中显示一张牌
	2.输出对战信息（）
		对战信息包括如下内容：
		本方英雄，本方技能，本方手牌数组[1~10],本方奥秘数组[1~7]，
		本方任务，本方随从数组[1~7]，本方武器，本方水晶（过载水晶，下回合过载水晶，剩余水晶，总水晶等信息）
		对方英雄，对方技能，对方手牌个数，对方奥秘个数
		对方任务，对方随从数组[1~7]，对方武器，对方水晶（过载水晶，下回合过载水晶，剩余水晶，总水晶等信息）

	3.消息接收器，发送给两个用户状态信息，
		比如 随从A攻击随从B,随从B死亡触发亡语
		

对外的界面基本上就这样
*/
const (
    SelectCard InputCmd = iota 
    SelectOnce              
    PutCard            
    Attack
    Skill       
)

const (
    Round OutputCmd = iota  //回合开始
    GameBegin               // 游戏开始
    SelectSta      // 状态选择
    BattleInfo   //对战信息
    MsgInfo     //消息信息
    DisplayCard // 显示某一张卡 （暂时方消息类里面得了）
    BattleCard      // 对战卡（暂时放消息里面得了）
)
const (
	fashi Hero = iota
)

type Card interface {
	// 输出卡牌信息
}

type BattleView interface {
	func input(cmd InputCmd,para string) error
	func  Run()  error
}
// 卡组类
type Deck interface {
	// 获取英雄信息
	func GetHero() (Hero,error)
	func GetCard() (Card[],error)
}
type Usr interface {
	func Send(cmd OutputCmd,info string) error
	// 获取用户的卡组信息
	func GetDeckInfo() (Deck,error)
}

func NewBattleFrame(u1 Usr, u2 Usr) (BattleView, error)

