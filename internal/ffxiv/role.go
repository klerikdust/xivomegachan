package ffxiv

import (
	"reflect"
)

type Role int

const (
	DPS Role = iota
	Healer
	Tank
	Empty
)

type Roles struct {
	Roles []Role
}

func (rs Roles) Emoji() string {
	if reflect.DeepEqual(rs.Roles, []Role{DPS}) {
		return "<:DPS:1434116707902820393>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{Healer}) {
		return "<:HEALER:1434116710105092176>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{Tank}) {
		return "<:TANK:1434116711988068433>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{DPS, Healer}) {
		return "<:ALL:1434116705679970304>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{DPS, Tank}) {
		return "<:ALL:1434116705679970304>"
	}
	if reflect.DeepEqual(rs.Roles, []Role{Healer, Tank}) {
		return "<:ALL:1434116705679970304>"
	}

	if reflect.DeepEqual(rs.Roles, []Role{Healer, Tank, DPS}) {
		return "<:ALL:1434116705679970304>"
	}

	return "<:ALL:1434116705679970304>"
}
