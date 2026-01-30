package main

import (
	"fmt"
	"log"
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
	MaxCharsInName = 42
)

type Name = [MaxCharsInName]byte

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

	name Name
}

func main() {
	const x, y, z = math.MinInt32, math.MaxInt32, 0
	const name = "aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = math.MaxInt32
	const mana = 1000
	const health = 1000
	const respect = 10
	const strength = 10
	const experience = 10
	const level = 10

	options := []Option{
		WithName("Вася Terminator3000 Ватрушкин"),
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
		WithType(personType),
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

	return func(person *GamePerson) {
		// only latin symbols and '_' are allowed
		var j int
		for i, s := range name {
			if (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') || s == '_' {
				person.name[j] = byte(s)
				j++
				if j >= MaxCharsInName {
					break // no panic, just trim
				}
			} else {
				log.Println("Warning: name symbol", s, "on position", i, "is not a latin symbol and skipped.")
			}
		}

	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	if x < math.MinInt32 || x > math.MaxInt32 ||
		y < math.MinInt32 || y > math.MaxInt32 ||
		z < math.MinInt32 || z > math.MaxInt32 {
		panic("X, Y or Z does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	if gold < 0 || gold > math.MaxUint32 {
		panic("Gold does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	if mana < 0 || mana > 1000 {
		panic("Mana does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(mana) << 0)
	}
}
func WithHealth(health int) func(*GamePerson) {
	if health < 0 || health > 1000 {
		panic("Health does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(health) << 10)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	if respect < 0 || respect > 10 {
		panic("Respect does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(respect) << 20)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	if strength < 0 || strength > 10 {
		panic("Strength does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(strength) << 24)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	if experience < 0 || experience > 10 {
		panic("Experience does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set2 = person.set2 | (uint8(experience) << 0)
	}
}

func WithLevel(level int) func(*GamePerson) {
	if level < 0 || level > 10 {
		panic("Level does not meet the requirements")
	}

	return func(person *GamePerson) {
		person.set2 = person.set2 | (uint8(level) << 4)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(1) << 28)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(1) << 29)
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.set1 = person.set1 | (uint32(1) << 30)
	}
}

func WithType(personType int) func(*GamePerson) {
	if personType < 0 || personType > 3 {
		panic("Person Type does not meet the requirements")
	}
	return func(person *GamePerson) {
		person.personType = uint8(personType)
	}
}

func (p *GamePerson) Name() string {
	name := make([]rune, len(p.name))
	for i := 0; i < len(p.name); i++ {
		name[i] = rune(p.name[i])
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
	return int((p.set1 >> 0) & uint32((1<<10)-1))
}

func (p *GamePerson) Health() int {
	return int((p.set1 >> 10) & uint32((1<<10)-1))
}

func (p *GamePerson) Respect() int {
	return int((p.set1 >> 20) & uint32((1<<4)-1))
}

func (p *GamePerson) Strength() int {
	return int((p.set1 >> 24) & uint32((1<<4)-1))
}

func (p *GamePerson) Experience() int {
	return int((p.set2 >> 0) & uint8((1<<4)-1))
}

func (p *GamePerson) Level() int {
	return int((p.set2 >> 4) & uint8((1<<4)-1))
}

func (p *GamePerson) HasHouse() bool {
	has := int(((p.set1 >> 28) & uint32(1)))
	if has == 1 {
		return true
	}
	return false
}

func (p *GamePerson) HasGun() bool {
	has := int(((p.set1 >> 29) & uint32(1)))
	if has == 1 {
		return true
	}
	return false
}

func (p *GamePerson) HasFamilty() bool {
	has := int(((p.set1 >> 30) & uint32(1)))
	if has == 1 {
		return true
	}
	return false
}

func (p *GamePerson) Type() int {
	return int(p.personType)
}
