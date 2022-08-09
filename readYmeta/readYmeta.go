/*
readYmeta.go reading and converting Yoda metadata
Author: Brett G. Olivier
email: @bgoli
licence: BSD 3 Clause
version: 0.5
(C) Brett G. Olivier, Vrije Universiteit Amsterdam, Amsterdam, The Netherlands, 2022
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// Define global constants here?
const _VERSION_ = "0.5.alpha"

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
		} `json:"Persistent_Identifier"`
		RelationType string `json:"Relation_Type"`
		Title        string `json:"Title"`
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
		RelationType string `json:"Relation_Type,omitempty"`
		Title        string `json:"Title,omitempty"`
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

const DEBUG bool = true

func main() {

	msg := "readYmeta - (C)Brett G. Olivier, Vrije Universiteit Amsterdam, 2022"
	fmt.Println(msg)
	fmt.Println(" ")

	// define input and output files
	input_file_name, input_file_path, err1 := get_input_file_path_from_clargs()
	errcntrl(err1)
	output_file_name := input_file_name + ".pdf"
	if DEBUG {
		fmt.Println(input_file_name)
		fmt.Println(input_file_path)
		fmt.Println(output_file_name)
	}

	// read metadata file
	json_file, err1 := os.ReadFile(input_file_name)
	errcntrl(err1)

	// print the file cast as string
	if DEBUG {
		fmt.Print(string(json_file))
	}

	// create metadata struct and fill it with file data
	var json_dat Yoda18Metadata
	err2 := json.Unmarshal(json_file, &json_dat)
	errcntrl(err2)

	// print struct and explore new options
	/*
		fmt.Println(" ")
		fmt.Println(json_dat)
		fmt.Println(reflect.TypeOf(json_dat))
		fmt.Println(len(json_dat.Contributor))
	*/
	// lets do something more useful
	if DEBUG {
		fmt.Printf("\n\n----------------\n\n")
	}
	//var doc_text_basic []string
	// var basic_info_str []string

	doc_text_basic := get_basic_document_data(json_dat)
	// basic_info_str = get_basic_mutable_data(json_dat)

	// Random checks
	if DEBUG {
		fmt.Println(doc_text_basic)
		fmt.Println(reflect.TypeOf(doc_text_basic))
		fmt.Println(doc_text_basic[0])
	}

	// lets play with dumping to PDF
	doc_array := doc_text_basic
	// doc_array = append(simple_info_str, basic_info_str...)
	pdf_create_and_dump(output_file_name, doc_array)
}

func errcntrl(e error) {
	if e != nil {
		panic(e)
	}
}

func get_input_file_path_from_clargs() (string, string, error) {
	var cDir string = ""
	var err error = nil
	var fname string

	cDir, err = os.Getwd()
	errcntrl(err)

	if len(os.Args) > 1 {
		fname = os.Args[1]
	} else {
		fmt.Println("Filename argument not provided, using default: yoda-metadata.json")
		fname = "yoda-metadata.json"
	}

	input_file_path, err := filepath.Abs(filepath.Join(cDir, fname))
	errcntrl(err)

	//
	_, err = os.Stat(input_file_path)
	if os.IsNotExist(err) {
		fmt.Println("Input file path does not exist:", input_file_path)
	} else {
		fmt.Println("Input file path exists:", input_file_path)
		err = nil
	}
	return fname, input_file_path, err

}

func pdf_create_and_dump(fname string, sarr []string) {

	// Do things

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	//m.SetBorder(true)

	m.Row(10, func() {
		m.Col(4, func() {
			m.Text(fmt.Sprintf("Metadata report generated by readYmeta (v%s)", _VERSION_), props.Text{
				Top:         0,
				Size:        16,
				Extrapolate: true,
			})
		})
		m.ColSpace(4)
	})
	// m.Row(10, func() {})
	m.Line(10)

	//write the "document"
	for idx, ele := range sarr {
		if DEBUG {
			fmt.Println("Index :", idx, " Element :", ele)
		}
		// pdf_write_row(m, fmt.Sprintln("Index :", idx, " Element :", ele))
		pdf_write_row(m, fmt.Sprintln(ele))
	}

	run_time := time.Now()

	m.Line(4)
	m.Row(4, func() {
		m.Col(4, func() {
			m.Text(fmt.Sprintf("Generated by readYmeta (v%s) [ https://github.com/vu-rdm-tech/yoda-metadata-toolkit ] on %d-%d-%d (%d:%d)", _VERSION_,
				run_time.Year(), run_time.Month(), run_time.Day(), run_time.Hour(), run_time.Minute()),
				props.Text{
					Top:         0,
					Size:        6,
					Extrapolate: true,
				})
		})
		m.ColSpace(4)
	})
	// m.Row(10, func() {})

	err := m.OutputFileAndClose(fmt.Sprintf("%s", fname))
	errcntrl(err)
}

func pdf_write_row(m pdf.Maroto, line string) {
	m.Row(6, func() {
		m.Col(4, func() {
			m.Text(line, props.Text{
				Top:         0,
				Size:        10,
				Extrapolate: true,
			})
		})
		//m.ColSpace(12)
	})
}

func get_creators(doc Yoda18Metadata) []string {
	var output []string
	output = append(output, fmt.Sprintf("Creators"))
	for i := range doc.Creator {
		output = append(output, fmt.Sprintf("- Author: %s %s", doc.Creator[i].Name.GivenName, doc.Creator[i].Name.FamilyName))
		for j := range doc.Creator[i].PersonIdentifier {
			output = append(output, fmt.Sprintf("-- (%s) %s", doc.Creator[i].PersonIdentifier[j].NameIdentifierScheme, doc.Creator[i].PersonIdentifier[j].NameIdentifier))
		}
		// output = append(output, fmt.Sprintf("-- Affiliation:"))
		for k := range doc.Creator[i].Affiliation {
			output = append(output, fmt.Sprintf("-- %s", doc.Creator[i].Affiliation[k]))
		}
	}
	return output
}

func get_contributors(doc Yoda18Metadata) []string {
	var output []string
	output = append(output, fmt.Sprintf("Contributors"))
	for i := range doc.Contributor {
		output = append(output, fmt.Sprintf("- %s: %s %s", doc.Contributor[i].ContributorType, doc.Contributor[i].Name.GivenName, doc.Contributor[i].Name.FamilyName))
		for j := range doc.Contributor[i].PersonIdentifier {
			output = append(output, fmt.Sprintf("-- (%s) %s", doc.Contributor[i].PersonIdentifier[j].NameIdentifierScheme, doc.Contributor[i].PersonIdentifier[j].NameIdentifier))
		}
		// output = append(output, fmt.Sprintf("-- Affiliation:"))
		for k := range doc.Contributor[i].Affiliation {
			output = append(output, fmt.Sprintf("-- %s", doc.Contributor[i].Affiliation[k]))
		}
	}
	return output
}

func get_basic_document_data(doc Yoda18Metadata) []string {
	var output []string
	output = append(output, fmt.Sprintf("Title: %s", doc.Title))
	output = append(output, get_creators(doc)...)
	output = append(output, fmt.Sprintf("Description: %s", doc.Description))
	output = append(output, fmt.Sprintf("Disciplines"))
	for i := range doc.Discipline {
		output = append(output, fmt.Sprintf("- %s", doc.Discipline[i]))
	}
	output = append(output, fmt.Sprintf("Tags"))
	for i := range doc.Tag {
		output = append(output, fmt.Sprintf("- %s", doc.Tag[i]))
	}
	output = append(output, get_contributors(doc)...)

	output = append(output, fmt.Sprintf("Related_Datapackages"))
	for i := range doc.RelatedDatapackage {
		output = append(output, fmt.Sprintf("- %s", doc.RelatedDatapackage[i].Title))
		output = append(output, fmt.Sprintf("-- %s", doc.RelatedDatapackage[i].RelationType))
		if doc.RelatedDatapackage[i].PersistentIdentifier.IdentifierScheme == "" {
			output = append(output, fmt.Sprintf("--(string) %s", doc.RelatedDatapackage[i].PersistentIdentifier.Identifier))
		} else {
			output = append(output, fmt.Sprintf("--(%s) %s", doc.RelatedDatapackage[i].PersistentIdentifier.IdentifierScheme,
				doc.RelatedDatapackage[i].PersistentIdentifier.Identifier))
		}
	}

	output = append(output, fmt.Sprintf("Funding_reference"))
	for i := range doc.FundingReference {
		output = append(output, fmt.Sprintf("- %s: %s", doc.FundingReference[i].FunderName, doc.FundingReference[i].AwardNumber))
	}

	output = append(output, fmt.Sprintf("Collected: %s - %s", doc.Collected.StartDate, doc.Collected.EndDate))
	output = append(output, fmt.Sprintf("Covered_Period: %s - %s", doc.CoveredPeriod.StartDate, doc.CoveredPeriod.EndDate))
	output = append(output, fmt.Sprintf("Covered_Geolocation_Place"))
	for i := range doc.CoveredGeolocationPlace {
		output = append(output, fmt.Sprintf("- %s", doc.CoveredGeolocationPlace[i]))
	}
	output = append(output, fmt.Sprintf("Version: %s", doc.Version))
	output = append(output, fmt.Sprintf("Licence: %s", doc.License))
	output = append(output, fmt.Sprintf("Language: %s", doc.Language))
	output = append(output, fmt.Sprintf("Data_Type: %s", doc.DataType))
	output = append(output, fmt.Sprintf("Data_Classification: %s", doc.DataClassification))
	output = append(output, fmt.Sprintf("Data_Access_Restriction: %s", doc.DataAccessRestriction))
	output = append(output, fmt.Sprintf("Rentention_Period: %d", doc.RetentionPeriod))
	output = append(output, fmt.Sprintf("Retention_Information: %s", doc.RetentionInformation))
	output = append(output, fmt.Sprintf("Embargo_End_Date: %s", doc.EmbargoEndDate))
	output = append(output, fmt.Sprintf("Collection_Name: %s", doc.CollectionName))
	output = append(output, fmt.Sprintf("Remarks: %s", doc.Remarks))
	//links
	output = append(output, fmt.Sprintf("Links"))
	for i := range doc.Links {
		output = append(output, fmt.Sprintf("- %s %s", doc.Links[i].Rel, doc.Links[i].Href))
	}
	return output
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
