package router

type (
	UserData struct {
		LOGIN             string
		MOVING            string
		LEAVE             string
		JOIN              string
		ATTACK            string
		ATTACK_SUCCESS    string
		OUT_RANGE         string
		FLAG              string
		JUMP              string
		SKILL_SHIELD      string
		SKILL_THORNS      string
		SKILL_DEFORMATION string
		SKILL_IMMUNE      string
		RELIVE            string
		FLAG_SUCCESS      string
		USE_ITEM          string
		ITEM_ADD_HP       string
		ITEM_1003         string
		UPDATE_DEVT       string
		LEADERBOARD       string
		REWAED            string
	}

	GameData struct {
		Start string
		End   string
	}

	GlobalData struct {
		INIT_USER  string
		GAME_START string
		GAME_END   string
		TIME       string
	}

	RoomData struct {
		ROOM_ALL          string
		ROOM_CTEATE       string
		ROOM_JOIN         string
		ROOM_JOIN_QUICK   string
		ROOM_LEAVE        string
		ROOM_MATCH        string
		ROOM_RENAME       string
		ROOM_CHANGE_PWD   string
		ROOM_KICK_OFF     string
		ROOM_DEL          string
		ROOM_OWNER_CHANGE string
	}

	ErrorData struct {
		ERROR string
	}

	SUserEvent struct {
		USER   UserData
		GAME   GameData
		GLOBAL GlobalData
		ROOM   RoomData
		ERROR  ErrorData
	}
	ISUserEvent interface {
	}
)

var USER_EVENT = &SUserEvent{
	USER: UserData{
		LOGIN:             "login",
		MOVING:            "moving",
		LEAVE:             "leave",
		JOIN:              "join",
		ATTACK:            "attack",
		ATTACK_SUCCESS:    "attack_success",
		OUT_RANGE:         "out_range",
		FLAG:              "flag",
		JUMP:              "jump",
		SKILL_SHIELD:      "skill_shield",
		SKILL_THORNS:      "skill_thorns",
		SKILL_DEFORMATION: "skill_deformation",
		SKILL_IMMUNE:      "skill_immune",
		RELIVE:            "relive",
		FLAG_SUCCESS:      "flag_success",
		USE_ITEM:          "useitem",
		ITEM_ADD_HP:       "addHp",
		ITEM_1003:         "item_1003",
		UPDATE_DEVT:       "update_dvt",
		LEADERBOARD:       "leaderboard",
		REWAED:            "reward",
	},

	ROOM: RoomData{
		ROOM_ALL:          "Room_all",        //获取房间数据
		ROOM_CTEATE:       "room_create",     //创建房间
		ROOM_JOIN:         "room_join",       //加入房间
		ROOM_JOIN_QUICK:   "room_join_quick", //加入房间
		ROOM_LEAVE:        "room_leave",      //离开房间
		ROOM_MATCH:        "room_match",      //匹配，
		ROOM_RENAME:       "room_rename",     //房间重命名
		ROOM_CHANGE_PWD:   "room_change_pwd", //房间修改密码
		ROOM_KICK_OFF:     "room_kick_off",   //踢人
		ROOM_DEL:          "room_del",
		ROOM_OWNER_CHANGE: "room_owner_change",
	},

	GAME: GameData{
		Start: "game_start",
		End:   "game_end",
	},
	GLOBAL: GlobalData{
		INIT_USER:  "init_user",
		GAME_START: "game_start",
		GAME_END:   "game_end",
		TIME:       "time",
	},

	ERROR: ErrorData{
		ERROR: "error",
	},
}
