package main

import (
	"_lab/data-pump/common"
	"_lab/data-pump/db"
	"_lab/data-pump/models"
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	registry = common.NewRegistry()
)

// isSubset checks if struct B is a subset of struct A using generics
// All non-empty/non-zero fields in B must match the corresponding fields in A
func isSubset[T any](a, b T) bool {
	return compareSubset(reflect.ValueOf(a), reflect.ValueOf(b))
}

func compareSubset(a, b reflect.Value) bool {
	if a.Type() != b.Type() {
		return false
	}

	switch b.Kind() {
	case reflect.Struct:
		// Special handling for time.Time
		if b.Type().String() == "time.Time" {
			return compareTimeValues(a, b)
		}

		for i := 0; i < b.NumField(); i++ {
			fieldA := a.Field(i)
			fieldB := b.Field(i)

			if isZeroValue(fieldB) {
				continue
			}

			if !compareSubset(fieldA, fieldB) {
				return false
			}
		}
		return true

	case reflect.Slice, reflect.Array:
		if b.Len() == 0 {
			return true
		}
		if a.Len() != b.Len() {
			return false
		}
		for i := 0; i < b.Len(); i++ {
			if !compareSubset(a.Index(i), b.Index(i)) {
				return false
			}
		}
		return true

	case reflect.String:
		return b.String() == "" || a.String() == b.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return b.Int() == 0 || a.Int() == b.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return b.Uint() == 0 || a.Uint() == b.Uint()

	case reflect.Float32, reflect.Float64:
		return b.Float() == 0.0 || a.Float() == b.Float()

	case reflect.Bool:
		return !b.Bool() || a.Bool() == b.Bool()

	case reflect.Interface, reflect.Ptr:
		if b.IsNil() {
			return true
		}
		if a.IsNil() {
			return false
		}
		return compareSubset(a.Elem(), b.Elem())

	default:
		return reflect.DeepEqual(a.Interface(), b.Interface())
	}
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Slice, reflect.Array:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0.0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Struct:
		if v.Type().String() == "time.Time" {
			return v.Interface().(time.Time).IsZero()
		}
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(v.Interface(), zero.Interface())
	default:
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(v.Interface(), zero.Interface())
	}
}

func compareTimeValues(a, b reflect.Value) bool {
	timeA := a.Interface().(time.Time)
	timeB := b.Interface().(time.Time)

	if timeB.IsZero() {
		return true
	}

	return timeA.Truncate(time.Microsecond).Equal(timeB.Truncate(time.Microsecond))
}

func main() {
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to connect client: %v", err)
	}

	instance := registry.Instance(client, common.ContractAddress)

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("\033[31m[ERR]\033[0m Failed to connect the database")
	}

	var dbEntry models.Entry
	if err := dbConn.Order("RANDOM()").Preload("Ownership").Preload("Properties").First(&dbEntry).Error; err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to fetch an entry: %v", err)
	}

	log.Printf("DB Data: \033[1;36m%v\033[0m\n", dbEntry)

	data, err := bind.Call(instance, nil, registry.PackGetLatestProperty(dbEntry.ID), registry.UnpackGetLatestProperty)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to unpack latest property: %v", err)
	}

	var bcEntry models.Entry
	if err := json.Unmarshal([]byte(data), &bcEntry); err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to unmarshal data: %v", err)
	}

	log.Printf("BC Data: \033[1;36m%v\033[0m\n", bcEntry)

	log.Printf("Success: \033[1;32m%t\033[0m\n", isSubset(dbEntry, bcEntry))
}
