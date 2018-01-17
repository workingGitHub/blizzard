package BattleFrame

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
