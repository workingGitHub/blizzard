# 对战局框架设计  

- 整体对局框架采用状态机模型  
- 处理均采用消息堆栈模型  
- 每张卡采用钩子进行处理，钩子函数中只能追加消息处理内容，不能对对局产生直接影响  


包含如下状态：
1. init
1. GameBegin
1. 回合开始
1. 回合中
1. 回合结束
1. 3～5 反复循环。。。  直到一方获得胜利


判断胜负的时机
1. 状态机切换时
1. 用户输入消息处理完成之后

用户交互消息处理函数格式如下：  
`func msgif(para,context)`

event消息处理函数：  
`func msgif(para,context)`

随从所受到的伤害以及buf均在一条消息处理完成之后统一结算判断该随从是否死亡

## 状态机-- init
1. [battle] 初始化对战上下文（用户信息，卡组信息，打乱双方排序,初始化卡牌信息）  
1. [card] 执行每张卡的init钩子  `card：init（）`
1. [event] 添加执行先手抽三张牌，后手抽四张牌（注意：这里是否可以算严格意义上的抽牌？先把这个和正常抽牌保持一致） `event：drawcard，para=usr` 
1. [event] 等待双方用户选择结束 `event：wait,para=“usr:usr`
2. [event] 在后手的手牌中添加一个水晶卡牌 `event: addcard,para="cardname:xxx,usr=usr1/usr2"`
3. 开始顺序执行event中的消息，消息空了就开始执行结算，结算后继续执行event，直到结算后消息依然为空
4. 【battle】 执行initend钩子

该状态下可以两个usr可以执行和输出的命令：  
	- `exchange para=“1，2，3，4”`  
    - `end`

## 状态机-- GameBegin
1. [card] 执行牌库中所有卡牌的GameBegin钩子 `card:GameBegin()`
2. [battle] 执行对战的Gameegin处理（需要跳转到回合开始  event goto,para = usr）
2. 

## 状态机 -- usr 回合开始
1. 【battle】 恢复水晶，preBegin 执行回合开始的操作
1. [card] 执行所有卡牌的回合开始钩子 `card：begin（indeck,谁的回合）`
2. [card] 执行所有手牌中卡牌的回合开始钩子 `card：begin（inhand,谁的回合）`
2. [card] 执行所有战场上卡牌的回合开始钩子 `card：begin（inbattle,谁的回合）`
3. 开始顺序执行event中的消息，消息空了就开始执行结算，结算后继续执行event，直到结算后消息依然为空
5. 【battle】  执行回合开始的end钩子

## 状态机 -- usr 回合中
1. 等待用户命令输入  
用户可以执行如下命令；
	- `putcard para=“1～10,指定目标【0～f】”`
	- `attack para="0~7,8~f"`
    - `英雄技能 para=指定目标【0～f】`
    - `end`

在后台已经是选择完成后的结果

抉择，升级，选择需要交互，需要进入一个交互选择界面，提供交互状态机  
	在交互界面中需要选择，只能执行一条命令
    - `select 1～3`

