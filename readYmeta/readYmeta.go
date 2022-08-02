// readYmeta.go

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// Vanilla Yoda metadata struct
type Yoda18Metadata struct {
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Discipline []string `json:"Discipline"`
	Language   string   `json:"Language"`
	Collected  struct {
		StartDate string `json:"Start_Date"`
		EndDate   string `json:"End_Date"`
	} `json:"Collected"`
	CoveredGeolocationPlace []string `json:"Covered_Geolocation_Place"`
	CoveredPeriod           struct {
		StartDate string `json:"Start_Date"`
		EndDate   string `json:"End_Date"`
	} `json:"Covered_Period"`
	Tag                []string `json:"Tag"`
	RelatedDatapackage []struct {
		PersistentIdentifier struct {
			IdentifierScheme string `json:"Identifier_Scheme"`
			Identifier       string `json:"Identifier"`
		} `json:"Persistent_Identifier,omitempty"`
		RelationType          string `json:"Relation_Type"`
		Title                 string `json:"Title"`
		PersistentIdentifier0 struct {
			Identifier string `json:"Identifier"`
		} `json:"Persistent_Identifier,omitempty"`
	} `json:"Related_Datapackage"`
	RetentionPeriod  int    `json:"Retention_Period"`
	DataType         string `json:"Data_Type"`
	FundingReference []struct {
		FunderName  string `json:"Funder_Name"`
		AwardNumber string `json:"Award_Number"`
	} `json:"Funding_Reference"`
	Creator []struct {
		Name struct {
			GivenName  string `json:"Given_Name"`
			FamilyName string `json:"Family_Name"`
		} `json:"Name"`
		Affiliation      []string `json:"Affiliation"`
		PersonIdentifier []struct {
			NameIdentifierScheme string `json:"Name_Identifier_Scheme"`
			NameIdentifier       string `json:"Name_Identifier"`
		} `json:"Person_Identifier"`
	} `json:"Creator"`
	Contributor []struct {
		Name struct {
			GivenName  string `json:"Given_Name"`
			FamilyName string `json:"Family_Name"`
		} `json:"Name"`
		Affiliation      []string `json:"Affiliation"`
		PersonIdentifier []struct {
			NameIdentifierScheme string `json:"Name_Identifier_Scheme"`
			NameIdentifier       string `json:"Name_Identifier"`
		} `json:"Person_Identifier"`
		ContributorType string `json:"Contributor_Type"`
	} `json:"Contributor"`
	DataAccessRestriction string `json:"Data_Access_Restriction"`
	Title                 string `json:"Title"`
	Description           string `json:"Description"`
	Version               string `json:"Version"`
	RetentionInformation  string `json:"Retention_Information"`
	EmbargoEndDate        string `json:"Embargo_End_Date"`
	DataClassification    string `json:"Data_Classification"`
	CollectionName        string `json:"Collection_Name"`
	Remarks               string `json:"Remarks"`
	License               string `json:"License"`
}

// Yoda metadata struct with advanced options
type Yoda18MetadataV2 struct {
	Links []struct {
		Rel  string `json:"rel,omitempty"`
		Href string `json:"href,omitempty"`
	} `json:"links,omitempty"`
	Discipline []string `json:"Discipline,omitempty"`
	Language   string   `json:"Language,omitempty"`
	Collected  struct {
		StartDate string `json:"Start_Date,omitempty"`
		EndDate   string `json:"End_Date,omitempty"`
	} `json:"Collected,omitempty"`
	CoveredGeolocationPlace []string `json:"Covered_Geolocation_Place,omitempty"`
	CoveredPeriod           struct {
		StartDate string `json:"Start_Date,omitempty"`
		EndDate   string `json:"End_Date,omitempty"`
	} `json:"Covered_Period,omitempty"`
	Tag                []string `json:"Tag,omitempty"`
	RelatedDatapackage []struct {
		PersistentIdentifier struct {
			IdentifierScheme string `json:"Identifier_Scheme,omitempty"`
			Identifier       string `json:"Identifier,omitempty"`
		} `json:"Persistent_Identifier,omitempty"`
		RelationType          string `json:"Relation_Type,omitempty"`
		Title                 string `json:"Title,omitempty"`
		PersistentIdentifier0 struct {
			Identifier string `json:"Identifier,omitempty"`
		} `json:"Persistent_Identifier,omitempty"`
	} `json:"Related_Datapackage,omitempty"`
	RetentionPeriod  int    `json:"Retention_Period,omitempty"`
	DataType         string `json:"Data_Type,omitempty"`
	FundingReference []struct {
		FunderName  string `json:"Funder_Name,omitempty"`
		AwardNumber string `json:"Award_Number,omitempty"`
	} `json:"Funding_Reference,omitempty"`
	Creator []struct {
		Name struct {
			GivenName  string `json:"Given_Name,omitempty"`
			FamilyName string `json:"Family_Name,omitempty"`
		} `json:"Name,omitempty"`
		Affiliation      []string `json:"Affiliation,omitempty"`
		PersonIdentifier []struct {
			NameIdentifierScheme string `json:"Name_Identifier_Scheme,omitempty"`
			NameIdentifier       string `json:"Name_Identifier,omitempty"`
		} `json:"Person_Identifier,omitempty"`
	} `json:"Creator,omitempty"`
	Contributor []struct {
		Name struct {
			GivenName  string `json:"Given_Name,omitempty"`
			FamilyName string `json:"Family_Name,omitempty"`
		} `json:"Name,omitempty"`
		Affiliation      []string `json:"Affiliation,omitempty"`
		PersonIdentifier []struct {
			NameIdentifierScheme string `json:"Name_Identifier_Scheme,omitempty"`
			NameIdentifier       string `json:"Name_Identifier,omitempty"`
		} `json:"Person_Identifier,omitempty"`
		ContributorType string `json:"Contributor_Type,omitempty"`
	} `json:"Contributor,omitempty"`
	DataAccessRestriction string `json:"Data_Access_Restriction,omitempty"`
	Title                 string `json:"Title,omitempty"`
	Description           string `json:"Description,omitempty"`
	Version               string `json:"Version,omitempty"`
	RetentionInformation  string `json:"Retention_Information,omitempty"`
	EmbargoEndDate        string `json:"Embargo_End_Date,omitempty"`
	DataClassification    string `json:"Data_Classification,omitempty"`
	CollectionName        string `json:"Collection_Name,omitempty"`
	Remarks               string `json:"Remarks,omitempty"`
	License               string `json:"License,omitempty"`
}

func errcntrl(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	msg := "Welcome to the Yoda metadata translator"
	fmt.Println(msg)

	// read metadata file
	json_file, err1 := os.ReadFile("yoda-metadata.json")
	errcntrl(err1)

	// print the file cast as string
	fmt.Print(string(json_file))

	// create metadata struct and fill it with file data
	var json_dat Yoda18Metadata
	err2 := json.Unmarshal(json_file, &json_dat)
	errcntrl(err2)

	// print struct and explore new options
	fmt.Println(" ")
	fmt.Println(json_dat)
	fmt.Println(reflect.TypeOf(json_dat))
	fmt.Println(len(json_dat.Contributor))
}

/*
Seems to work!
Funky GO template builder: https://mholt.github.io/json-to-go/
*/

// learning how not to do stuff
/*
func read_json_branch(json_map map[string]interface{}) {
	fmt.Println(" ")
	fmt.Println("json_map")
	fmt.Println(json_map)
	fmt.Println(reflect.TypeOf(json_map))

	for k, v := range json_map {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			fmt.Println(reflect.TypeOf(k))
			fmt.Println(reflect.TypeOf(v))
			fmt.Println(reflect.TypeOf(vv))
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}


*/
