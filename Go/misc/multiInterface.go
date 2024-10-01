package advanced

import "fmt"

type Warrior interface {
	GetWarriorSetting()
}
type Witcher interface {
	GetWitcherSetting()
}

func (t RoleSetting) GetWarriorSetting() string {
	return "Warrior"
}

func (t RoleSetting) GetWitcherSetting() string {
	return "Witcher"
}

type RoleSetting struct {
	Warrior
	Witcher
}

type GameData struct {
	Role RoleSetting
}

func RunMultiInterface() {
	G := GameData{}
	fmt.Println(G.Role.GetWarriorSetting())
}
