package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID                        string     `json:"id" gorm:"primaryKey"`
	TpNo                      string     `json:"tp_no"`
	Village                   int        `json:"village"`
	VillageName               string     `json:"village__name"`
	VillageParentName         string     `json:"village__parent__name"`
	VillageParentDistrictName string     `json:"village__parent__district__name"`
	Ownership                 []Owner    `json:"ownership"`
	CreatedAt                 time.Time  `json:"created_at"`
	Properties                []Property `json:"properties"`
}

type Owner struct {
	ID             string `json:"id" gorm:"primaryKey"`
	EntryID        string `json:"-"`
	Owner          string `json:"owner"`
	OwnerFirstName string `json:"owner__first_name"`
}

type Property struct {
	ID             string     `json:"id" gorm:"primaryKey"`
	EntryID        string     `json:"-"`
	PropertyID     string     `json:"property_id"`
	Parent         Relative   `json:"parent" `
	Children       []Relative `json:"children"`
	BlockNoCode    string     `json:"block_no__code"`
	ResurveyNo     string     `json:"resurvey_no"`
	SurveyType     string     `json:"survey_type"`
	SubDivisionNo  string     `json:"sub_division_no"`
	Area           float64    `json:"area"`
	Type           string     `json:"type"`
	Classification string     `json:"classification"`
	BasicTaxRate   float64    `json:"basic_tax_rate"`
	OwnedDate      string     `json:"owned_date"`
	ForfeitedDate  string     `json:"forfeited_date"`
	IsFreezed      bool       `json:"is_freezed"`
}

type Relative struct {
	ID         string  `json:"id" gorm:"primaryKey"`
	PropertyID string  `json:"property_id"`
	TpEntry    TpEntry `json:"tp_entry"`
}

type TpEntry struct {
	ID         string `json:"id" gorm:"primaryKey"`
	RelativeID string `json:"-"`
	TpNo       string `json:"tp_no"`
}

func RandomEntry() Entry {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return Entry{
		ID:                        uuid.NewString(),
		TpNo:                      randomString("TP-", 4),
		Village:                   rand.Intn(10000),
		VillageName:               randomString("Village-", 2),
		VillageParentName:         randomString("Parent-", 2),
		VillageParentDistrictName: randomString("District-", 2),
		Ownership:                 generateOwners(),
		Properties:                generateProperties(),
	}
}

func generateOwners() []Owner {
	var owners []Owner
	firstName := randomString("John-", 2)
	fullName := firstName + " " + randomString("Doe-", 2)
	owners = append(owners, Owner{
		ID:             uuid.NewString(),
		Owner:          fullName,
		OwnerFirstName: firstName,
	})

	return owners
}

func generateProperties() []Property {
	var props []Property
	props = append(props, Property{
		ID:             uuid.NewString(),
		PropertyID:     randomString("PID-", 5),
		Parent:         Relative{},
		Children:       []Relative{},
		BlockNoCode:    randomString("BL-", 3),
		ResurveyNo:     randomString("RS-", 3),
		SurveyType:     randomPick([]string{"Private", "Government"}),
		SubDivisionNo:  randomString("SD-", 2),
		Area:           float64(rand.Intn(1000)) + rand.Float64(),
		Type:           randomPick([]string{"Residential", "Agricultural", "Commercial"}),
		Classification: randomPick([]string{"Urban", "Rural"}),
		BasicTaxRate:   rand.Float64() * 5,
		OwnedDate:      time.Now().AddDate(-rand.Intn(10), 0, 0).Format("2006-01-02"),
		ForfeitedDate:  "",
		IsFreezed:      rand.Intn(2) == 1,
	})
	return props
}

func randomString(prefix string, digits int) string {
	return prefix + randomDigits(digits)
}

func randomDigits(n int) string {
	d := ""
	for range n {
		d += string('0' + rune(rand.Intn(10)))
	}
	return d
}

func randomPick(options []string) string {
	return options[rand.Intn(len(options))]
}
