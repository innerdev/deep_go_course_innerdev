package main

import (
	"fmt"
	"math"
	"unsafe"
)

type Option func(*GamePerson)

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

const (
	minX = math.MinInt32
	maxX = math.MaxInt32

	minY = math.MinInt32
	maxY = math.MaxInt32

	minZ = math.MinInt32
	maxZ = math.MaxInt32

	minGold = 0
	maxGold = math.MaxUint32

	minMana = 0
	maxMana = 1000

	minHealth = 0
	maxHealth = 1000

	minRespect = 0
	maxRespect = 10

	minStrength = 0
	maxStrength = 10

	minExperience = 0
	maxExperience = 10

	minLevel = 0
	maxLevel = 10

	nameMaxChars = 42
)

const (
	manaBitsShift       = 0
	healthBitsShift     = 10
	respectBitsShift    = 20
	strengthBitsShift   = 24
	houseBitsShift      = 28
	gunBitsShift        = 29
	familyBitsShift     = 30

	experienceBitsShift = 0
	levelBitsShift      = 4
)

const (
	manaMask       = (1 << 10) - 1
	healthMask     = (1 << 10) - 1
	respectMask    = (1 << 4) - 1
	strengthMask   = (1 << 4) - 1
	experienceMask = (1 << 4) - 1
	levelMask      = (1 << 4) - 1
)

const (
	bitTrue = 1
)

type GamePerson struct { // 64 bytes exectly
	x    int32
	y    int32
	z    int32
	gold uint32 // 1 bit could be used, but no need to

	// order:
	// 10 bit mana, shift 0
	// 10 bits health, shift 10
	// 4 bits respect, shift 20
	// 4 bits strength, shift 24
	// 1 bit house, shift 28
	// 1 bit gun, shift 29
	// 1 bit family, shift 30
	// size: 31 bits
	set1 uint32

	// order:
	// 4 bits experience, shift 0
	// 4 bits level, shift 4
	// sum size: 8 bits
	set2 uint8

	personType uint8

	name [42]byte
}

func main() {
	options := []Option{
		WithName("__Terminator3000__"),
		WithCoordinates(math.MinInt32, math.MaxInt32, 0),
		WithGold(math.MaxInt32),
		WithMana(1000),
		WithHealth(999),
		WithRespect(10),
		WithStrength(10),
		WithExperience(10),
		WithLevel(10),
		WithHouse(),
		WithFamily(),
		WithGun(),
		WithType(BuilderGamePersonType),
	}

	gp := NewGamePerson(options...)

	fmt.Println(gp)
	fmt.Printf("Name: |%s|\n", gp.Name())
	fmt.Println("Size: ", unsafe.Sizeof(gp))
}

func NewGamePerson(options ...Option) GamePerson {
	gp := GamePerson{}
	for _, f := range options {
		f(&gp)
	}
	return gp
}

func WithName(name string) func(*GamePerson) {
	if len(name) <= 0 {
		panic("Person name can not be empty.")
	}

	for _, s := range name {
		if s > 'z' {
			panic("Non-latin symbols are not allowed in name.")
		}
	}

	if len(name) > nameMaxChars {
		panic("Person name's length can't be greater then 42 symbols.")
	}

	return func(person *GamePerson) {
		for i, s := range name {
			person.name[i] = byte(s)
		}
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	if x < minX || x > maxX ||
		y < minY || y > maxY ||
		z < minZ || z > maxZ {
		panic("X, Y or Z does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	if gold < minGold || gold > maxGold {
		panic("Gold does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	if mana < minMana || mana > maxMana {
		panic("Mana does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(mana) << manaBitsShift)
	}
}
func WithHealth(health int) func(*GamePerson) {
	if health < minHealth || health > maxHealth {
		panic("Health does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(health) << healthBitsShift)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	if respect < minRespect || respect > maxRespect {
		panic("Respect does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(respect) << respectBitsShift)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	if strength < minStrength || strength > maxStrength {
		panic("Strength does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(strength) << strengthBitsShift)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	if experience < minExperience || experience > maxExperience {
		panic("Experience does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set2 = person.set2 | (uint8(experience) << experienceBitsShift)
	}
}

func WithLevel(level int) func(*GamePerson) {
	if level < minLevel || level > maxLevel {
		panic("Level does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set2 = person.set2 | (uint8(level) << levelBitsShift)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(bitTrue) << houseBitsShift)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(bitTrue) << gunBitsShift)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(bitTrue) << familyBitsShift)
	}
}

func WithType(personType int) func(*GamePerson) {
	if personType != BuilderGamePersonType &&
		personType != BlacksmithGamePersonType &&
		personType != WarriorGamePersonType {
		panic("Person Type does not meet the requirements")
	}
	return func(person *GamePerson) {
		person.personType = uint8(personType)
	}
}

func (p *GamePerson) Name() string {
	name := make([]rune, 0, len(p.name))
	for _, s := range p.name {
		name = append(name, rune(s))
	}
	return string(name)
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int((p.set1 >> manaBitsShift) & uint32(manaMask))
}

func (p *GamePerson) Health() int {
	return int((p.set1 >> healthBitsShift) & uint32(healthMask))
}

func (p *GamePerson) Respect() int {
	return int((p.set1 >> respectBitsShift) & uint32(respectMask))
}

func (p *GamePerson) Strength() int {
	return int((p.set1 >> strengthBitsShift) & uint32(strengthMask))
}

func (p *GamePerson) Experience() int {
	return int((p.set2 >> experienceBitsShift) & uint8(experienceMask))
}

func (p *GamePerson) Level() int {
	return int((p.set2 >> levelBitsShift) & uint8(levelMask))
}

func (p *GamePerson) HasHouse() bool {
	has := int(((p.set1 >> houseBitsShift) & uint32(bitTrue)))
	if has == bitTrue {
		return true
	}
	return false
}

func (p *GamePerson) HasGun() bool {
	has := int(((p.set1 >> gunBitsShift) & uint32(bitTrue)))
	if has == bitTrue {
		return true
	}
	return false
}

func (p *GamePerson) HasFamilty() bool {
	has := int(((p.set1 >> familyBitsShift) & uint32(bitTrue)))
	if has == bitTrue {
		return true
	}
	return false
}

func (p *GamePerson) Type() int {
	return int(p.personType)
}
