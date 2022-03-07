package common

//功能模块
type SMode struct {
	none  int
	Game  int
	Map   int
	User  int
	Skill int
	Item  int
	Room  int
	Timer int
}

var Mode = &SMode{
	none:  0,
	Game:  1,
	Map:   2,
	User:  3,
	Skill: 4,
	Item:  5,
	Room:  6,
	Timer: 999,
}

type DRoomSatus struct {
	Room_wait  int32
	Game_start int32
}

var RoomSatus = &DRoomSatus{
	Room_wait:  0,
	Game_start: 1,
}

type SErrorCode struct {
	No_player      int
	Already_in     int
	Pwd_error      int
	Already_start  int
	dvt_not_enough int
	No_room        int
	Not_owner      int
	Not_hero       int
	Play_max_limit int
	No_empty_room  int
	Error_name     int
}

var ErrorCode = &SErrorCode{
	No_player:      5001,
	Already_in:     5002,
	Pwd_error:      5003,
	Already_start:  5004,
	dvt_not_enough: 5005,
	No_room:        5006,
	Not_owner:      5007,
	Not_hero:       5008,
	Play_max_limit: 5009,
	No_empty_room:  5010,
	Error_name:     5011,
}

type SItem struct {
	consume   int
	equip     int
	box       int
	allAttr   int
	box_piece int
}

var Item = &SItem{
	consume:   1,
	equip:     2,
	box:       1001, //宝箱
	allAttr:   1004, //全属性
	box_piece: 3001, //宝箱碎片
}

type SItemAttr struct {
	Type_1  int
	Type_2  int
	Type_3  int
	Type_4  int
	Type_5  int
	Type_6  int
	Type_7  int
	Type_8  int
	Type_9  int
	Type_10 int
}

var ItemAttr = &SItemAttr{
	Type_1:  1,
	Type_2:  2,
	Type_3:  3,
	Type_4:  4,
	Type_5:  5,
	Type_6:  6,
	Type_7:  7,
	Type_8:  8,
	Type_9:  9,
	Type_10: 10,
}

type SEquipAttr struct {
	atk         int
	def         int
	hp          int
	crit        int
	crit_damage int
	dod         int
	all         int
}

var EquipAttr = &SEquipAttr{
	atk:         1,
	def:         2,
	hp:          3,
	crit:        4,
	crit_damage: 5,
	dod:         6,
	all:         9,
}

//========================================================================
//职业枚举
type SRole struct {
	All        int
	Ranger     int
	Alchemist  int
	Warrior    int
	Adventurer int
	Rogue      int
}

var Role = &SRole{
	All:        0,
	Ranger:     1,
	Alchemist:  2,
	Warrior:    3,
	Adventurer: 4,
	Rogue:      5,
}

type SAttr struct {
	Crit_damage_base float64
}

var Attr = &SAttr{
	Crit_damage_base: 2000.0,
}

type SBattle_mode struct {
	Reduce  float64
	Defence float64
}

var Battle_mode = &SBattle_mode{
	Reduce:  1,
	Defence: 2,
}

type SPart struct {
	Part_1 int32
	Part_2 int32
}

var Part = &SPart{
	Part_1: 1,
	Part_2: 2,
}

//=========================================================================
//排行榜
type SLeaderBoard struct {
	maxRank int
	BaseNum int
}

var LeaderBoard = &SLeaderBoard{
	maxRank: 1000,
	BaseNum: 10000000000000,
}

type SLeaderBoardRefresh struct {
	week    int
	history int
}

var LeaderBoardRefresh = &SLeaderBoardRefresh{
	week:    1,
	history: 2,
}

type SLeaderBoardType struct {
	Week    int
	History int
}

var LeaderBoardType = &SLeaderBoardType{
	Week:    1,
	History: 2,
}

type SLeaderBoardTypePrefix struct {
	week    string
	history string
}

var LeaderBoardTypePrefix = &SLeaderBoardTypePrefix{
	week:    "week",
	history: "history",
}

type SLeaderBoardMode struct {
	Kill int
	Win  int
}

var LeaderBoardMode = &SLeaderBoardMode{
	Kill: 1,
	Win:  2,
}

type SLeaderBoardWeek struct {
}

var LeaderBoardWeek = &SLeaderBoardWeek{}

type SLeaderBoardModePrefix struct {
	Kill string
	Win  string
}

var LeaderBoardModePrefix = &SLeaderBoardModePrefix{
	Kill: "kill",
	Win:  "win",
}

type SLeaderBoardRole struct {
	All        int
	Ranger     int
	Alchemist  int
	Warrior    int
	Adventurer int
	Rogue      int
}

var LeaderBoardRole = &SLeaderBoardRole{
	All:        0,
	Ranger:     1,
	Alchemist:  2,
	Warrior:    3,
	Adventurer: 4,
	Rogue:      5,
}

type SLeaderBoardRolePrefix struct {
	All        string
	Ranger     string
	Alchemist  string
	Warrior    string
	Adventurer string
	Rogue      string
}

var LeaderBoardRolePrefix = &SLeaderBoardRolePrefix{
	All:        "All",
	Ranger:     "Ranger",
	Alchemist:  "Alchemist",
	Warrior:    "Warrior",
	Adventurer: "Adventurer",
	Rogue:      "Rogue",
}

type SLeaderBoardHistoryPrefix struct {
	h_kill   int
	h_win    int
	h_mining int
}

var LeaderBoardHistoryPrefix = SLeaderBoardHistoryPrefix{
	h_kill:   1,
	h_win:    2,
	h_mining: 3,
}

type SLogMode struct {
	Login        int
	Battle_round int
	Match        int
	Attack       int
}

var LogMode = &SLogMode{
	Login:        1,
	Battle_round: 2,
	Match:        3,
	Attack:       4,
}
