package models

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type (
	Entry struct {
		gorm.Model
		CardNumber string `json:"card_number" gorm:"uniqueIndex;size:15"`
		Lands      []Land `json:"lands" gorm:"foreignKey:EntryID"`
	}

	Owner struct {
		gorm.Model
		LandID             uint
		Name               string `json:"name"`
		Surname            string `json:"surname"`
		Address            string `json:"address"`
		ThandaperNumber    string `json:"thandaper_number"`
		ThandaperSubNumber string `json:"thandaper_sub_number"`
		OwnershipType      string `json:"ownership_type"`
		IsAadharLinked     bool   `json:"is_aadhar_linked"`
	}

	District struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		RevenueID      int    `json:"revenue_id"`
		RevenueCode    string `json:"revenue_code"`
		RevenueName    string `json:"revenue_name"`
		RevenueNameMal string `json:"revenue_name_mal"`
		LsgCode        string `json:"lsg_code"`
	}

	Taluk struct {
		ID             int      `json:"id"`
		Name           string   `json:"name"`
		RevenueID      int      `json:"revenue_id"`
		RevenueCode    string   `json:"revenue_code"`
		RevenueName    string   `json:"revenue_name"`
		RevenueNameMal string   `json:"revenue_name_mal"`
		LsgCode        string   `json:"lsg_code"`
		District       District `json:"district" gorm:"embedded"`
	}

	Village struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		RevenueID      int    `json:"revenue_id"`
		RevenueCode    string `json:"revenue_code"`
		RevenueName    string `json:"revenue_name"`
		RevenueNameMal string `json:"revenue_name_mal"`
		LsgCode        string `json:"lsg_code"`
		Taluk          Taluk  `json:"taluk" gorm:"embedded"`
	}

	Block struct {
		ID      string `json:"id"`
		Code    string `json:"code"`
		Village int    `json:"village"`
	}

	Building struct {
		gorm.Model
		LandID                uint
		LrdReferenceID        string   `json:"lrd_reference_id"`
		Name                  string   `json:"name"`
		ThandaperNumber       string   `json:"thandaper_number"`
		AssessmentOrderNumber string   `json:"assessment_order_number"`
		BuildingType          string   `json:"building_type"`
		Area                  *float64 `json:"area"`
		AreaUnit              string   `json:"area_unit"`
	}

	Land struct {
		gorm.Model
		EntryID            uint
		ReferenceID        string     `json:"reference_id"`
		Owners             []Owner    `json:"owners" gorm:"foreignKey:LandID"`
		Village            Village    `json:"village" gorm:"embedded"`
		Block              Block      `json:"block" gorm:"embedded"`
		SurveyNumber       string     `json:"survey_number"`
		SubDivisionNumber  string     `json:"sub_division_number"`
		Area               string     `json:"area"`
		AreaUnit           string     `json:"area_unit"`
		LandType           string     `json:"land_type"`
		LandTypeCode       string     `json:"land_type_code"`
		LastTaxPaymentDate string     `json:"last_tax_payment_date"`
		LsgiCode           string     `json:"lsgi_code"`
		LsgName            string     `json:"lsg_name"`
		IsStLand           bool       `json:"is_st_land"`
		IsEmrLinked        bool       `json:"is_emr_linked"`
		FairValue          string     `json:"fair_value"`
		FairValueLandType  string     `json:"fair_value_land_type"`
		Buildings          []Building `json:"buildings" gorm:"foreignKey:LandID"`
	}
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateRandomOwner() Owner {
	names := []string{"ALICE", "BOB", "CAROL", "DAVID", "EVE"}
	surnames := []string{"J", "K", "L", "M", "N"}
	ownershipTypes := []string{"Single", "Joint"}
	addresses := []string{"123 MAIN ST", "456 NEW ST", "789 OLD ST", "369 CENTRAL"}

	return Owner{
		Name:               names[rand.Intn(len(names))],
		Surname:            surnames[rand.Intn(len(surnames))],
		Address:            addresses[rand.Intn(len(addresses))],
		ThandaperNumber:    randomIntString(4),
		ThandaperSubNumber: randomIntString(rand.Intn(2)),
		OwnershipType:      ownershipTypes[rand.Intn(len(ownershipTypes))],
		IsAadharLinked:     rand.Intn(2) == 1,
	}
}

func generateRandomDistrict() District {
	return District{
		ID:             rand.Intn(30) + 1,
		Name:           "District_" + randomString(3),
		RevenueID:      rand.Intn(2000) + 100,
		RevenueCode:    randomIntString(2),
		RevenueName:    "REVENUE_DISTRICT_" + randomString(4),
		RevenueNameMal: "ജില്ല", // Mal: District
		LsgCode:        randomIntString(3),
	}
}

func generateRandomTaluk() Taluk {
	return Taluk{
		ID:             rand.Intn(200) + 1,
		Name:           "Taluk_" + randomString(3),
		RevenueID:      rand.Intn(100) + 1,
		RevenueCode:    randomIntString(2),
		RevenueName:    "REVENUE_TALUK_" + randomString(4),
		RevenueNameMal: "താലൂക്ക്", // Mal: Taluk
		LsgCode:        randomIntString(4),
		District:       generateRandomDistrict(),
	}
}

func generateRandomVillage() Village {
	return Village{
		ID:             rand.Intn(3000) + 1,
		Name:           "Village_" + randomString(4),
		RevenueID:      rand.Intn(2000) + 1,
		RevenueCode:    randomIntString(2),
		RevenueName:    "REVENUE_VILLAGE_" + randomString(5),
		RevenueNameMal: "ഗ്രാമം", // Mal: Village
		LsgCode:        randomIntString(6),
		Taluk:          generateRandomTaluk(),
	}
}

func generateRandomBlock(villageID int) Block {
	return Block{
		ID:      randomString(8) + "-" + randomString(4) + "-" + randomString(4) + "-" + randomString(4) + "-" + randomString(12),
		Code:    randomIntString(3),
		Village: villageID,
	}
}

func generateRandomBuilding() Building {
	var area *float64
	if rand.Intn(3) != 0 { // 2/3 chance of having an area
		a := rand.Float64() * 1000.0
		area = &a
	}

	buildingTypes := []string{"R.B", "C.B", "S.B"}
	return Building{
		LrdReferenceID:        randomIntString(5),
		Name:                  "Name " + randomString(5),
		ThandaperNumber:       randomIntString(5),
		AssessmentOrderNumber: randomString(2) + "/" + randomIntString(6) + "/" + randomIntString(4),
		BuildingType:          buildingTypes[rand.Intn(len(buildingTypes))],
		Area:                  area,
		AreaUnit:              "m²",
	}
}

func generateRandomLand() Land {
	village := generateRandomVillage()

	land := Land{
		ReferenceID:        randomIntString(2) + "/" + randomIntString(2) + "/" + randomIntString(2) + "/" + randomIntString(3) + "/" + randomIntString(2) + "/" + randomIntString(1) + "/" + randomIntString(4) + "//" + randomIntString(3),
		Owners:             []Owner{generateRandomOwner()},
		Village:            village,
		Block:              generateRandomBlock(village.ID),
		SurveyNumber:       randomIntString(2),
		SubDivisionNumber:  randomIntString(1),
		Area:               randomFloatString(100, 500),
		AreaUnit:           []string{"sqm", "cent"}[rand.Intn(2)],
		LandType:           "പുരയിടം",
		LandTypeCode:       randomIntString(1),
		LastTaxPaymentDate: randomDate(),
		LsgiCode:           randomIntString(5),
		LsgName:            "Lsg_" + randomString(5),
		IsStLand:           rand.Intn(2) == 1,
		IsEmrLinked:        rand.Intn(2) == 1,
		FairValue:          randomFloatString(10000, 100000),
		FairValueLandType:  []string{"Wet land", "Dry land"}[rand.Intn(2)],
		Buildings:          []Building{generateRandomBuilding()},
	}

	return land
}

func RandomEntry() Entry {
	return Entry{
		CardNumber: randomIntString(10),
		Lands:      []Land{generateRandomLand()},
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func randomIntString(length int) string {
	const charset = "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func randomDate() string {
	min := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}

func randomFloatString(min, max float64) string {
	return fmt.Sprintf("%.2f", min+(rand.Float64()*(max-min)))
}
