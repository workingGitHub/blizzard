package repositories

import (
	"errors"

	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"

	_ "github.com/mattn/go-sqlite3"
)

type DbManagerIf interface {
	Close() error
	CreateUsrTable(name string, password string) (int64, error)
	GetUsrtableByID(id int64) (UsrTable, error)
	GetByUsernameAndPassword(username, userPassword string) (UsrTable, error)
}

/********************[数据表]*******************************/
type UsrTable struct {
	Id       int64
	Username string `xorm:"varchar(25) notnull unique "`
	PhoneNO  string `xorm:"varchar(25)  "`
	Email    string `xorm:"varchar(25)  "`
	Password string `xorm:"varchar(25) notnull "`
	Gold     uint   `xorm:"Int  default(0)"`
	Dusy     uint   `xorm:"Int  default(0)"`
	Card     uint   `xorm:"Int  default(0)"`
	IsJJCing bool   `xorm:"Bool  default(0)"`
	JJCdeck  int64  `xorm:"BigInt  default(0)"`
	Deck01   int64  `xorm:"BigInt  default(0)"`
	Deck02   int64  `xorm:"BigInt  default(0)"`
	Deck03   int64  `xorm:"BigInt  default(0)"`
	Deck04   int64  `xorm:"BigInt  default(0)"`
	Deck05   int64  `xorm:"BigInt  default(0)"`
	Deck06   int64  `xorm:"BigInt  default(0)"`
	Deck07   int64  `xorm:"BigInt  default(0)"`
	Deck08   int64  `xorm:"BigInt  default(0)"`
	Deck09   int64  `xorm:"BigInt  default(0)"`
	Deck10   int64  `xorm:"BigInt  default(0)"`
	Deck11   int64  `xorm:"BigInt  default(0)"`
	Deck12   int64  `xorm:"BigInt  default(0)"`
	Deck13   int64  `xorm:"BigInt  default(0)"`
	Deck14   int64  `xorm:"BigInt  default(0)"`
	Deck15   int64  `xorm:"BigInt  default(0)"`
	Deck16   int64  `xorm:"BigInt  default(0)"`
	Deck17   int64  `xorm:"BigInt  default(0)"`
	Deck18   int64  `xorm:"BigInt  default(0)"`
}
type DeckTable struct {
	Id       int64
	DeckType uint   `xorm:"Int  default(0)"`
	Deckname string `xorm:"varchar(25) "`
	Card01   int64  `xorm:"BigInt  default(0)"`
	Card02   int64  `xorm:"BigInt  default(0)"`
	Card03   int64  `xorm:"BigInt  default(0)"`
	Card04   int64  `xorm:"BigInt  default(0)"`
	Card05   int64  `xorm:"BigInt  default(0)"`
	Card06   int64  `xorm:"BigInt  default(0)"`
	Card07   int64  `xorm:"BigInt  default(0)"`
	Card08   int64  `xorm:"BigInt  default(0)"`
	Card09   int64  `xorm:"BigInt  default(0)"`
	Card10   int64  `xorm:"BigInt  default(0)"`
	Card11   int64  `xorm:"BigInt  default(0)"`
	Card12   int64  `xorm:"BigInt  default(0)"`
	Card13   int64  `xorm:"BigInt  default(0)"`
	Card14   int64  `xorm:"BigInt  default(0)"`
	Card15   int64  `xorm:"BigInt  default(0)"`
	Card16   int64  `xorm:"BigInt  default(0)"`
	Card17   int64  `xorm:"BigInt  default(0)"`
	Card18   int64  `xorm:"BigInt  default(0)"`
	Card19   int64  `xorm:"BigInt  default(0)"`
	Card20   int64  `xorm:"BigInt  default(0)"`
	Card21   int64  `xorm:"BigInt  default(0)"`
	Card22   int64  `xorm:"BigInt  default(0)"`
	Card23   int64  `xorm:"BigInt  default(0)"`
	Card24   int64  `xorm:"BigInt  default(0)"`
	Card25   int64  `xorm:"BigInt  default(0)"`
	Card26   int64  `xorm:"BigInt  default(0)"`
	Card27   int64  `xorm:"BigInt  default(0)"`
	Card28   int64  `xorm:"BigInt  default(0)"`
	Card29   int64  `xorm:"BigInt  default(0)"`
	Card30   int64  `xorm:"BigInt  default(0)"`
}

/*
卡牌表：

|字段		   | 字段类型 	|字段含义		|备注|
|--------	|--------	|----		|---|
|CardID	    |  String 	|卡牌ID	| 主键  	|
|CardName	|  String 	|卡牌名称 	|   	|
|CardSet	|  int 		|卡牌类组	|基础/经典/加基森|
|CardType	|  int 		|卡牌类型	|随从/法术/英雄卡/武器|
|Faction	|  int 		|卡牌派系	|通用/法..猎/污..玉莲帮|
|Rarity		|  int 		|稀有程度	|基本/基础/稀有/史诗/传说|
|Cost		|  int 		|卡牌费用   ||
|Atk		|  int 		|攻击力    |武器-攻击力/随从-攻击力|
|Health		|  int 		|血量	     |武器-耐久/随从-血量|
|AttackType |  int 		|攻击类型  |普通攻击/随机攻击/主动攻击|
|CardTextInHand|int|随从类型|嘲讽/圣盾/冲锋/隐藏/剧毒/法伤+1+2+5/免疫/吸血/|
|CardRace	| int|随从种族|普通/野兽/元素/鱼人/恶魔/机械/龙/海盗 |
|Actor-Card	| String|Actor处理函数|在卡组中处理方式|
|Actor-Mudi	| String|Actor处理函数|在墓地中处理方式|
|Actor-Hand	| String|Actor处理函数|在手牌中处理方式|
|Actor-Battle| String|Actor处理函数|在战场中处理方式|
*/
type CardTable struct {
	Id             int64
	CardID         string `xorm:"varchar(50) notnull unique"`
	CardName       string `xorm:"varchar(50) notnull unique"`
	CardSet        uint   `xorm:"Int  default(0)"`
	CardType       uint   `xorm:"Int  default(0)"`
	Faction        uint   `xorm:"Int  default(0)"`
	Rarity         uint   `xorm:"Int  default(0)"`
	Cost           uint   `xorm:"Int  default(0)"`
	Atk            uint   `xorm:"Int  default(0)"`
	Health         uint   `xorm:"Int  default(0)"`
	AttackType     uint   `xorm:"Int  default(0)"`
	CardTextInHand int64  `xorm:"BigInt  default(0)"`
	CardRace       uint   `xorm:"Int  default(0)"`
	Detail         string `xorm:"varchar(255) "`
	Actor_Card     string `xorm:"varchar(50) "`
	Actor_Mudi     string `xorm:"varchar(50) "`
	Actor_Hand     string `xorm:"varchar(50) "`
	Actor_Battle   string `xorm:"varchar(50) "`
}

/***************************************************/
// TODO：创建数据库管理类
func NewRepository(app *iris.Application) DbManagerIf {
	orm, err := xorm.NewEngine("sqlite3", "./blizzard.db")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}
	err = orm.Sync2(new(UsrTable))

	if err != nil {
		app.Logger().Fatalf("orm failed to initialized User table: %v", err)
	}
	err = orm.Sync2(new(DeckTable))

	if err != nil {
		app.Logger().Fatalf("orm failed to initialized Deck table: %v", err)
	}

	err = orm.Sync2(new(CardTable))

	if err != nil {
		app.Logger().Fatalf("orm failed to initialized Card table: %v", err)
	}

	return &DbManager{orm: orm}
}

type DbManager struct {
	orm *xorm.Engine
}

// 关闭链接数据库
func (u DbManager) Close() error {
	if u.orm == nil {
		return nil
	}

	return u.orm.Close()
}

// 创建用户信息
func (u DbManager) CreateUsrTable(name string, password string) (int64, error) {
	if u.orm == nil {
		return -1, errors.New("数据库 创建失败。")
	}
	if name == "" || password == "" {
		return -1, errors.New("用户名密码不能为空")
	}
	/*
		Username string  `xorm:"varchar(25) notnull unique "`
		PhoneNO  string  `xorm:"varchar(25) default("") unique "`
		Email	        string  `xorm:"varchar(25) default("") unique "`
	*/
	// 查找注册用户名是否已被占用
	usr := UsrTable{Username: name}
	if ok, _ := u.orm.Get(&usr); ok {
		return -1, errors.New("用户名已被占用")
	}
	usr = UsrTable{PhoneNO: name}
	if ok2, _ := u.orm.Get(&usr); ok2 {
		return -1, errors.New("用户名已被占用")
	}
	usr = UsrTable{Email: name}
	if ok3, _ := u.orm.Get(&usr); ok3 {
		return -1, errors.New("用户名已被占用")
	}

	// TODO: 后续需要吧密码做加密处理
	usr = UsrTable{Username: name, Password: password}
	_, err := u.orm.Insert(usr)
	if err != nil {
		return -1, err
	}

	usr = UsrTable{Username: name}
	if ok4, _ := u.orm.Get(&usr); ok4 {
		return usr.Id, nil
	} else {
		return -1, errors.New("用户创建失败")
	}
}

func (u DbManager) GetUsrtableByID(id int64) (UsrTable, error) {
	usr := UsrTable{Id: id}
	if ok, _ := u.orm.Get(&usr); !ok {
		return usr, errors.New("用户名不存在")
	}
	return usr, nil
}
func (u DbManager) GetByUsernameAndPassword(username, userPassword string) (UsrTable, error) {
	usr := UsrTable{Username: username, Password: userPassword}
	if ok, _ := u.orm.Get(&usr); !ok {
		return usr, errors.New("用户名密码错误")
	}
	return usr, nil
}
