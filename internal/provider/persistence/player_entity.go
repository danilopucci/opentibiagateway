package persistence

type PlayerEntity struct {
	ID                  int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name                string `gorm:"column:name;unique;not null"`
	GroupID             int    `gorm:"column:group_id;default:1"`
	AccountID           int    `gorm:"column:account_id;default:0"`
	Level               int    `gorm:"column:level;default:1"`
	Vocation            int    `gorm:"column:vocation;default:0"`
	Health              int    `gorm:"column:health;default:150"`
	HealthMax           int    `gorm:"column:healthmax;default:150"`
	Experience          uint64 `gorm:"column:experience;default:0"`
	LookBody            int    `gorm:"column:lookbody;default:0"`
	LookFeet            int    `gorm:"column:lookfeet;default:0"`
	LookHead            int    `gorm:"column:lookhead;default:0"`
	LookLegs            int    `gorm:"column:looklegs;default:0"`
	LookType            int    `gorm:"column:looktype;default:136"`
	LookAddons          int    `gorm:"column:lookaddons;default:0"`
	MagLevel            int    `gorm:"column:maglevel;default:0"`
	Mana                int    `gorm:"column:mana;default:0"`
	ManaMax             int    `gorm:"column:manamax;default:0"`
	ManaSpent           uint64 `gorm:"column:manaspent;default:0"`
	Soul                uint   `gorm:"column:soul;default:0"`
	TownID              int    `gorm:"column:town_id;default:1"`
	PosX                int    `gorm:"column:posx;default:0"`
	PosY                int    `gorm:"column:posy;default:0"`
	PosZ                int    `gorm:"column:posz;default:0"`
	Conditions          []byte `gorm:"column:conditions"`
	Cap                 int    `gorm:"column:cap;default:400"`
	Sex                 int    `gorm:"column:sex;default:0"`
	LastLogin           uint64 `gorm:"column:lastlogin;default:0"`
	LastIP              uint   `gorm:"column:lastip;default:0"`
	Save                bool   `gorm:"column:save;default:1"`
	Skull               int8   `gorm:"column:skull;default:0"`
	SkullTime           int64  `gorm:"column:skulltime;default:0"`
	LastLogout          uint64 `gorm:"column:lastlogout;default:0"`
	Blessings           int8   `gorm:"column:blessings;default:0"`
	OnlineTime          int64  `gorm:"column:onlinetime;default:0"`
	Deletion            int64  `gorm:"column:deletion;default:0"`
	Balance             uint64 `gorm:"column:balance;default:0"`
	Stamina             uint16 `gorm:"column:stamina;default:1560"`
	SkillFist           uint   `gorm:"column:skill_fist;default:10"`
	SkillFistTries      uint64 `gorm:"column:skill_fist_tries;default:0"`
	SkillClub           uint   `gorm:"column:skill_club;default:10"`
	SkillClubTries      uint64 `gorm:"column:skill_club_tries;default:0"`
	SkillSword          uint   `gorm:"column:skill_sword;default:10"`
	SkillSwordTries     uint64 `gorm:"column:skill_sword_tries;default:0"`
	SkillAxe            uint   `gorm:"column:skill_axe;default:10"`
	SkillAxeTries       uint64 `gorm:"column:skill_axe_tries;default:0"`
	SkillDist           uint   `gorm:"column:skill_dist;default:10"`
	SkillDistTries      uint64 `gorm:"column:skill_dist_tries;default:0"`
	SkillShielding      uint   `gorm:"column:skill_shielding;default:10"`
	SkillShieldingTries uint64 `gorm:"column:skill_shielding_tries;default:0"`
	SkillFishing        uint   `gorm:"column:skill_fishing;default:10"`
	SkillFishingTries   uint64 `gorm:"column:skill_fishing_tries;default:0"`
	Created             int    `gorm:"column:created;default:0"`
	Hidden              bool   `gorm:"column:hidden;default:0"`
	Comment             string `gorm:"column:comment"`
}

func (PlayerEntity) TableName() string {
	return "players"
}
