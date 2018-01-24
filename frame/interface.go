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
    Intitate UsrInfo = iota  //先手
    Gote                    
)
const (
    Round OutputCmd = iota  //回合开始
    GameBegin               // 游戏开始
    SelectSta      // 状态选择
    BattleInfo   //对战信息
    MsgInfo     //消息信息
)
const (
	fashi Hero = iota
)

type Card interface {
	// 输出卡牌信息
}



/*
	// 英雄（英雄技能等）
	// 奥秘
	// 任务
	// 斧子
	// 随从
	// 手牌
	// 卡组
	// 水晶
	// 战场光环效果
	// 通用参数
*/
// 用户状态机，int均是用户需要操作的状态
// End 为本用户不需要操作的状态
const (
	Init_Choseing UsrSta = iota
	Rounding
	Selecting
	RoundEnd
)
// 卡组类
type Deck interface {
	// 获取英雄信息
	func GetHero() (Hero,error)
	// 抽取一张牌(触发卡牌的抽取描述动作)
	func DrawCard()(Card,error)
	// 获取N张牌(不触发卡牌抽取描述动作)
	func GetCard(n int)(Card[],error)
}
// 每一个用户
type Usr interface {
	// 发送给前端
	func Send(sta UsrSta) error
	func GetUsrSta() UsrSta
	// 获取deck信息
	func GetDeck()(Deck,error)
	
	// 洗牌
	func RandDeck()
	
	// 获取卡牌信息
	func SetUsrInfo(usr UsrInfo) 
	func GetUsrInfo() UsrInfo	 
	// 设置Init_Choseing内容
	func SetChoseCard(card []Card)error
	
	// 获取Init_Choseing内容 
	func GetChoseCard()(card []Card,error)
}

type Cmd interface {
	// 检查命令合法性
	func Checkout() bool
	func Run()(error)
}

// 游戏引擎
type Engine interface {	
	// 运行游戏
	func Run()  error
	
	// 获取游戏状态
	func GetStatue() StatueInfo
	
	// 添加执行的命令
	func AddCmd(Cmd) error
	
	// 添加用户信息
	func AddAttacker(usr Usr) error
	
	// 添加裁判信息
	func AddWatcher(usr Usr) error
	
	// 收到用户的交互信息
	func Rev(usr UsrInfo,cmd InputCmd,para string) error
	
	// 获取对战用户信息
	func GetAttacker(usr UsrInfo) Usr
	
	// 随机选择用户先后手
	func RandAttacker() error	
}


