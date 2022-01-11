package gamedata

type (
	SFlagCtrl struct {
	}

	ISFlagCtrl interface {
	}
)

var FlagCtrl *SFlagCtrl = &SFlagCtrl{}

func (this *SFlagCtrl) Flag(round int, fromuser string, part int) {
	GameCtrl.UpDateFlagOwner(round, fromuser, part)
}
