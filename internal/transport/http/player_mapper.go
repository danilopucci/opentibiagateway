package http

import (
	playerpb "github.com/danilopucci/opentibiagateway/internal/protogen/v1"
)

func MapPlayerToResponse(p *playerpb.Player) *PlayerResponse {
	if p == nil {
		return nil
	}

	return &PlayerResponse{
		ID:                  p.Id,
		Name:                p.Name,
		GroupID:             p.GroupId,
		AccountID:           p.AccountId,
		Level:               p.Level,
		Vocation:            p.Vocation,
		Health:              p.Health,
		HealthMax:           p.Healthmax,
		Experience:          p.Experience,
		LookBody:            p.Lookbody,
		LookFeet:            p.Lookfeet,
		LookHead:            p.Lookhead,
		LookLegs:            p.Looklegs,
		LookType:            p.Looktype,
		LookAddons:          p.Lookaddons,
		MagicLevel:          p.Maglevel,
		Mana:                p.Mana,
		ManaMax:             p.Manamax,
		ManaSpent:           p.Manaspent,
		Soul:                p.Soul,
		TownID:              p.TownId,
		PosX:                p.Posx,
		PosY:                p.Posy,
		PosZ:                p.Posz,
		Conditions:          p.Conditions,
		Cap:                 p.Cap,
		Sex:                 p.Sex,
		LastLogin:           p.Lastlogin,
		LastIP:              p.Lastip,
		Save:                p.Save,
		Skull:               p.Skull,
		SkullTime:           p.Skulltime,
		LastLogout:          p.Lastlogout,
		Blessings:           p.Blessings,
		OnlineTime:          p.Onlinetime,
		Deletion:            p.Deletion,
		Balance:             p.Balance,
		Stamina:             p.Stamina,
		SkillFist:           p.SkillFist,
		SkillFistTries:      p.SkillFistTries,
		SkillClub:           p.SkillClub,
		SkillClubTries:      p.SkillClubTries,
		SkillSword:          p.SkillSword,
		SkillSwordTries:     p.SkillSwordTries,
		SkillAxe:            p.SkillAxe,
		SkillAxeTries:       p.SkillAxeTries,
		SkillDist:           p.SkillDist,
		SkillDistTries:      p.SkillDistTries,
		SkillShielding:      p.SkillShielding,
		SkillShieldingTries: p.SkillShieldingTries,
		SkillFishing:        p.SkillFishing,
		SkillFishingTries:   p.SkillFishingTries,
		Created:             p.Created,
		Hidden:              p.Hidden,
		Comment:             p.Comment,
	}
}
