package model

type AvailableSpots struct {
	main      uint16
	male      uint16
	female    uint16
	formal    uint16
	inter     uint16
	vocat     uint16
	nonFormal uint16
}

func NewAvailableSpots(main, male, female, formal, inter, vocat, nonFormal uint16) *AvailableSpots {
	return &AvailableSpots{main, male, female, formal, inter, vocat, nonFormal}
}

func (as *AvailableSpots) Main() uint16 {
	return as.main
}

func (as *AvailableSpots) Male() uint16 {
	return as.male
}

func (as *AvailableSpots) Female() uint16 {
	return as.female
}

func (as *AvailableSpots) Formal() uint16 {
	return as.formal
}

func (as *AvailableSpots) Inter() uint16 {
	return as.inter
}

func (as *AvailableSpots) Vocat() uint16 {
	return as.vocat
}

func (as *AvailableSpots) NonFormal() uint16 {
	return as.nonFormal
}
