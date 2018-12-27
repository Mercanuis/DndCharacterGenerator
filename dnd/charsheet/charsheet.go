package charsheet

import (
	"fmt"
	"math/rand"
	"time"
)

const ageLimit   = 100
const maxWeight  = 999

/*
 *	Gender Enumeration
 */
type Gender int

const genderTypes  = 3
const (
	Male   Gender = 0
	Female Gender = 1
	Other  Gender = 2
)

func (gender Gender) String() string {
	genders := make(map[Gender]string)
	genders[Male] 	= "Male"
	genders[Female] = "Female"
	genders[Other] 	= "Other"
	return genders[gender]
}

func genderFromInt(i int) Gender {
	genders := make(map[int]Gender)
	genders[0] = Male
	genders[1] = Female
	genders[2] = Other
	return genders[i]
}

/*
 * Hair Color Enumeration
 */
type HairColor int

const hairColorTypes = 8
const (
	BlueHair 		HairColor = 0
	BlondeHair 		HairColor = 1
	BlackHair		HairColor = 2
	BrunetteHair	HairColor = 3
	BrownHair		HairColor = 4
	RedHair			HairColor = 5
	WhiteHair		HairColor = 6
	GreenHair		HairColor = 7
)

func (hc HairColor) String() string {
	colors := make(map[HairColor]string)
	colors[BlueHair] 		= "Blue"
	colors[BlondeHair] 		= "Blonde"
	colors[BlackHair] 		= "Black"
	colors[BrunetteHair] 	= "Brunette"
	colors[BrownHair] 		= "Brown"
	colors[RedHair] 		= "Red"
	colors[WhiteHair] 		= "White"
	colors[GreenHair] 		= "Green"
	return colors[hc]
}

func hairFromInt(i int) HairColor {
	colors := make(map[int]HairColor)
	colors[0] = BlueHair
	colors[1] = BlondeHair
	colors[2] = BlackHair
	colors[3] = BrunetteHair
	colors[4] = BrownHair
	colors[5] = RedHair
	colors[6] = WhiteHair
	colors[7] = GreenHair
	return colors[i]
}

/*
 * Eye Color Enumeration
 */
type EyeColor int

const eyeColorTypes = 4
const (
	BlueEye 		= 0
	BrownEye 		= 1
	GreenEye 		= 2
	ColorlessEye 	= 3
)

func (ec EyeColor) String() string {
	colors := make(map[EyeColor]string)
	colors[BlueEye] 		= "Blue"
	colors[BrownEye] 		= "Brown"
	colors[GreenEye] 		= "Green"
	colors[ColorlessEye]	= "Colorless"
	return colors[ec]
}

func eyeFromInt(i int) EyeColor {
	colors := make(map[int]EyeColor)
	colors[0] = BlueEye
	colors[1] = BrownEye
	colors[2] = GreenEye
	colors[3] = ColorlessEye
	return colors[i]
}

/*
 *	Struct details for Attributes
 */
type Attributes struct {
	str  int
	dex  int
	con  int
	intl int
	wis  int
	chr  int
}

func (a *Attributes) generateAttributes() {
	var ats [6]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		n := rand.Intn(20)
		ats[i] = n
	}

	a.str 	= ats[0]
	a.dex 	= ats[1]
	a.con 	= ats[2]
	a.intl 	= ats[3]
	a.wis 	= ats[4]
	a.chr 	= ats[5]
}

func (a Attributes) getStr() int {
	return a.str
}

func (a Attributes) getDex() int {
	return a.dex
}

func (a Attributes) getCon() int {
	return a.con
}

func (a Attributes) getInt() int {
	return a.intl
}

func (a Attributes) getWis() int {
	return a.wis
}

func (a Attributes) getChr() int {
	return a.chr
}

/*
 *	Struct Details for CharDetails Struct
 */
type CharDetails struct {
	charName     string
	charAge      int
	charReligion string
	charHair     HairColor
	charEyes     EyeColor
	charGender   Gender
	charRace     string
}

func (d *CharDetails) generateDetails(name string) {
	rand.Seed(time.Now().UnixNano())

	d.charName = name
	d.charRace = "Hooman"
	d.charReligion = "Tyr"

	//Set Hair Color
	d.charHair = hairFromInt(rand.Intn(hairColorTypes))
	//Set Eye Color
	d.charEyes = eyeFromInt(rand.Intn(eyeColorTypes))
	//Set Age
	d.charAge = rand.Intn(ageLimit)
	//Set Gender
	d.charGender = genderFromInt(rand.Intn(genderTypes))
}

func (d CharDetails) getName() string {
	return d.charName
}

func (d CharDetails) getAge() int {
	return d.charAge
}

func (d CharDetails) getReligion() string {
	return d.charReligion
}

func (d CharDetails) getHair() string {
	return d.charHair.String()
}

func (d CharDetails) getEyes() string {
	return d.charEyes.String()
}

func (d CharDetails) getGender() string {
	return d.charGender.String()
}

func (d CharDetails) getRace() string {
	return d.charRace
}

/*
 *	Struct Details for Inventory Struct
 */
type Inventory struct {
	maxWeight  int
	currWeight int
}

func (i *Inventory) generateInventory() {
	//Set the Seed for rand
	rand.Seed(time.Now().UnixNano())
	i.currWeight = rand.Intn(maxWeight)
	i.maxWeight = rand.Intn(maxWeight)

}
func (i Inventory) getMaxWeight() int {
	return i.maxWeight
}

func (i Inventory) getCurrWeight() int {
	return i.currWeight
}

func (i Inventory) isEncumbered() bool {
	return i.currWeight > i.maxWeight
}

/*
 *	Struct Details for Character Sheet Struct
 */
type CharacterSheet struct {
	attributes *Attributes
	details    *CharDetails
	inventory  *Inventory
}

func CreateNewCharacter(name string) *CharacterSheet {
	details := CharDetails{}
	details.generateDetails(name)

	attributes := Attributes{}
	attributes.generateAttributes()

	inventory := Inventory{}
	inventory.generateInventory()

	sheet := CharacterSheet{details:&details, attributes:&attributes, inventory:&inventory}
	return &sheet
}

func PrintSheet(sheet *CharacterSheet) {
	fmt.Printf("\t\t\t\t=========== %s ===========\n", "CHARACTER SHEET")
	printDetails(sheet.details)
	printAttributes(sheet.attributes)
	printInventory(sheet.inventory)

}

func printDetails(d *CharDetails) {
	fmt.Println("===== Character Details =====")
	fmt.Printf("Name: %s\tAge: %d\tGender: %s\tRace: %s\tHair: %s\tEyes: %s\tReligion: %s\n\n",
		d.getName(), d.getAge(), d.getGender(), d.getRace(), d.getHair(), d.getEyes(), d.getReligion())
}

func printAttributes(a *Attributes) {
	fmt.Println("===== Attributes =====")
	fmt.Printf(
		"Strength %12d\nDexterity %11d\nConstitution %8d\nIntelligence %8d\nWisdom %14d\nCharisma %12d\n\n",
		a.getStr(), a.getDex(), a.getCon(), a.getInt(), a.getWis(), a.getWis())
}

func printInventory(i *Inventory) {
	fmt.Println("===== Inventory =====")
	fmt.Printf("Current Weight: %d\nMax Weight: %d\nEncumbered: %t\n",
		i.getCurrWeight(), i.getMaxWeight(), i.isEncumbered())
	fmt.Println("=====================")
}
