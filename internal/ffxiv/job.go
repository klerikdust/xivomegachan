package ffxiv

type Job int

const (
	GNB Job = iota
	PLD
	GLD
	DRK
	WAR
	MRD
	SCH
	ACN // Arcanist
	SGE
	AST
	WHM
	CNJ
	SAM
	DRG
	NIN
	MNK
	RPR
	VPR
	BRD
	MCH
	DNC
	BLM
	BLU
	SMN
	PCT
	RDM
	LNC
	PUG
	ROG
	THM
	ARC // Archer
	Unknown
)

func JobFromAbbreviation(abbreviation string) Job {
	switch abbreviation {
	case "GNB":
		return GNB
	case "PLD":
		return PLD
	case "GLD":
		return GLD
	case "DRK":
		return DRK
	case "WAR":
		return WAR
	case "MRD":
		return MRD
	case "SCH":
		return SCH
	case "ACN":
		return ACN
	case "SGE":
		return SGE
	case "AST":
		return AST
	case "WHM":
		return WHM
	case "CNJ":
		return CNJ
	case "SAM":
		return SAM
	case "DRG":
		return DRG
	case "NIN":
		return NIN
	case "MNK":
		return MNK
	case "RPR":
		return RPR
	case "VPR":
		return VPR
	case "BRD":
		return BRD
	case "MCH":
		return MCH
	case "DNC":
		return DNC
	case "BLM":
		return BLM
	case "BLU":
		return BLU
	case "SMN":
		return SMN
	case "PCT":
		return PCT
	case "RDM":
		return RDM
	case "LNC":
		return LNC
	case "PUG":
		return PUG
	case "ROG":
		return ROG
	case "THM":
		return THM
	case "ARC":
		return ARC
	}
	return Unknown
}

func (j Job) Emoji() string {
	switch j {
	case GNB:
		return "<:GNB:1434110976529268786>"
	case PLD:
		return "<:PLD:1434110969851805736>"
	case GLD:
		return "<:GLA:1434111643985379480>"
	case DRK:
		return "<:DRK:1434110973819490336>"
	case WAR:
		return "<:WAR:1434110972028522567>"
	case MRD:
		return "<:MRD:1434111648280608910>"
	case SCH:
		return "<:SCH:1434110963820400650>"
	case ACN:
		return "<:ACN:1434111174017945630>"
	case SGE:
		return "<:SGE:1434110965783203900>"
	case AST:
		return "<:AST:1434110967578628157>"
	case WHM:
		return "<:WHM:1434110955691704340>"
	case CNJ:
		return "<:CNJ:1434111652600614953>"
	case SAM:
		return "<:SAM:1434110935336878213>"
	case DRG:
		return "<:DRG:1434110951111528458>"
	case NIN:
		return "<:NIN:1434110942437834752>"
	case MNK:
		return "<:MNK:1434110948653797399>"
	case RPR:
		return "<:RPR:1434110920312881303>"
	case VPR:
		return "<:VPR:1434110917548834877>"
	case BRD:
		return "<:BRD:1434110953204613220>"
	case MCH:
		return "<:MCH:1434110940206600223>"
	case DNC:
		return "<:DNC:1434110923097899169>"
	case BLM:
		return "<:BLM:1434110946342862940>"
	case BLU:
		return "<:BLU:1434110932497207367>"
	case SMN:
		return "<:SMN:1434110944358695023>"
	case RDM:
		return "<:RDM:1434110937723437187>"
	case PCT:
		return "<:PCT:1434110914755559505>"
	case LNC:
		return "<:LNC:1434111646002843699>"
	case PUG:
		return "<:PGL:1434111650574635049>"
	case ROG:
		return "<:ROG:1434111657033990154>"
	case THM:
		return "<:THM:1434111654328537119>"
	case ARC:
		return "<:ARC:1434111641850613840>"
	}
	return "<:DOH:1434112269108641964>"
}
