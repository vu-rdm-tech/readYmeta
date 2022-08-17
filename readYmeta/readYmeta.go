/*
readYmeta.go reading and converting Yoda metadata
Author: Brett G. Olivier
email: @bgoli
licence: BSD 3 Clause
version: 0.6
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

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// Define global constants here?
const _VERSION_ = "0.6.5.alpha"

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

const DEBUG bool = false

const fontsize float64 = 10
const indentsymb string = " "
const minRGB8Bytes = 0
const maxRGB8Bytes = 255

func main() {

	msg := "readYmeta - (C)Brett G. Olivier, Vrije Universiteit Amsterdam, 2022"
	fmt.Println(msg)
	// fmt.Println()
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

	//// old way of doing this
	doc_text_basic := get_basic_document_data(json_dat)

	// Random checks
	if DEBUG {
		fmt.Println(doc_text_basic)
		fmt.Println(reflect.TypeOf(doc_text_basic))
		fmt.Println(doc_text_basic[0])
	}

	// dump to PDF
	doc_array := doc_text_basic
	pdf_create_and_dump(output_file_name, doc_array)

	//// old way of doing things

	fmt.Printf("\n\n-------***-------\n\n")
	//// New way of doing things where we write the document directly
	doc := pdf.NewMaroto(consts.Portrait, consts.A4)
	//m.SetBorder(true)
	doc.SetPageMargins(10, 10, 10)

	doc = generate_pdf_report_basic(json_dat, doc, input_file_name)

	pdf_write_row(doc, "I am some BLACK coloured TEXT!", 12, 4, consts.Normal, pdfBlack())
	pdf_write_row(doc, "I am some BLUE coloured TEXT!", 12, 4, consts.Normal, pdfBlue())
	pdf_write_row(doc, "I am some GREEN coloured TEXT!", 12, 4, consts.Normal, pdfGreen())
	pdf_write_row(doc, "I am some RED coloured TEXT!", 12, 4, consts.Normal, pdfRed())

	err := doc.OutputFileAndClose(fmt.Sprintf("%s", "new_form.pdf"))
	errcntrl(err)

}

func errcntrl(e error) {
	if e != nil {
		panic(e)
	}
}

// Maroto PDF color defintions
func pdfRed() color.Color {
	return color.Color{
		Red:   maxRGB8Bytes,
		Green: minRGB8Bytes,
		Blue:  minRGB8Bytes,
	}
}

// Maroto PDF color defintions
func pdfGreen() color.Color {
	return color.Color{
		Red:   minRGB8Bytes,
		Green: maxRGB8Bytes,
		Blue:  minRGB8Bytes,
	}
}

// Maroto PDF color defintions
func pdfBlue() color.Color {
	return color.Color{
		Red:   minRGB8Bytes,
		Green: minRGB8Bytes,
		Blue:  maxRGB8Bytes,
	}
}

// Maroto PDF color defintions
func pdfBlack() color.Color {
	return color.Color{
		Red:   minRGB8Bytes,
		Green: minRGB8Bytes,
		Blue:  minRGB8Bytes,
	}
}

func pdfOrange() color.Color {
	return color.Color{
		Red:   255,
		Green: 165,
		Blue:  0,
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

// New style PDFreportwriter, writes basic metadata
func generate_pdf_report_basic(data Yoda18Metadata, doc pdf.Maroto, fname string) pdf.Maroto {
	var ctime = time.Now().String()
	var colwidth uint = 12
	var rowheight float64 = 4
	var textblock_divider float64 = 20

	pdf_write_header(doc, fmt.Sprintf("\"%s\" metadata", fname), rowheight, colwidth)
	pdf_write_footer(doc, fmt.Sprintf("\"%s\" metadata generated on %s by readYmeta v%s", fname, ctime, _VERSION_), rowheight, colwidth)

	pdf_write_row(doc, "Title", rowheight, colwidth, consts.Bold, pdfBlack())
	pdf_write_row(doc, data.Title, rowheight, colwidth, consts.Normal, pdfBlack())
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_row(doc, "Description", rowheight, colwidth, consts.Bold, pdfBlack())
	if float64(len(data.Description))/textblock_divider > rowheight {
		pdf_write_row(doc, data.Description, float64(len(data.Description))/textblock_divider, colwidth, consts.Normal, pdfBlack())
	} else {
		pdf_write_row(doc, data.Description, rowheight, colwidth, consts.Normal, pdfOrange())
	}
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_row(doc, "Tags", rowheight, colwidth, consts.Bold, pdfBlack())
	pdf_write_list(doc, data.Tag, rowheight, colwidth, consts.Normal, pdfBlack())
	// pdf_write_list_sub1(doc, data.Tag, rowheight, colwidth, consts.Normal, pdfBlack())
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_creators(doc, data, rowheight, colwidth, consts.Normal, pdfBlack())
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_contributors(doc, data, rowheight, colwidth, consts.Normal, pdfBlack())
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_row(doc, "Disciplines", rowheight, colwidth, consts.Bold, pdfBlack())
	pdf_write_list(doc, data.Discipline, rowheight, colwidth, consts.Normal, pdfBlack())
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_row(doc, "Collected", rowheight, colwidth, consts.Bold, pdfBlack())
	pdf_write_row_tuple_indent(doc, "StartDate", data.Collected.StartDate, rowheight, colwidth, consts.Normal, pdfBlack(), 1)
	pdf_write_row_tuple_indent(doc, "EndDate", data.Collected.EndDate, rowheight, colwidth, consts.Normal, pdfBlack(), 1)
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_row(doc, "Covered Period", rowheight, colwidth, consts.Bold, pdfBlack())
	pdf_write_row_tuple_indent(doc, "StartDate", data.CoveredPeriod.StartDate, rowheight, colwidth, consts.Normal, pdfBlack(), 1)
	pdf_write_row_tuple_indent(doc, "EndDate", data.CoveredPeriod.EndDate, rowheight, colwidth, consts.Normal, pdfBlack(), 1)
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_funding(doc, data, rowheight, colwidth, consts.Normal, pdfBlack())
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	pdf_write_related(doc, data, rowheight, colwidth, consts.Normal, pdfBlack())
	pdf_write_row(doc, "", 1, colwidth, consts.Normal, pdfBlack())

	return doc
}

// New style PDFreportwriter header writer
func pdf_write_header(m pdf.Maroto, line string, rowheight float64, colwidth uint) {
	m.RegisterHeader(func() {
		m.Row(rowheight, func() {
			m.Col(colwidth, func() {
				m.Text(line, props.Text{
					Top:         0,
					Size:        12,
					Extrapolate: true,
				})
			})
			m.ColSpace(4)
		})
		m.Line(10)
	})
}

func pdf_write_footer(m pdf.Maroto, line string, rowheight float64, colwidth uint) {
	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				m.Text(line, props.Text{
					Top:   20,
					Style: consts.Italic,
					Size:  6,
					Align: consts.Left,
				})
			})
		})
	})
}

// New style PDFreportwriter row writer
func pdf_write_row(m pdf.Maroto, line string, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color) {
	m.Row(rowheight, func() {
		m.Col(colwidth, func() {
			m.Text(line, props.Text{
				Top:         0,
				Size:        fontsize,
				Extrapolate: false,
				Style:       fontstyle,
				Color:       textcolour,
			})
		})
	})
}

// New style PDFreportwriter row writer
func pdf_write_row_indent(m pdf.Maroto, line string, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color, indent uint) {
	if line == "" || line == " " {
		textcolour = pdfRed()
	}

	m.Row(rowheight, func() {
		m.Col(indent, func() {
			m.Text(indentsymb, props.Text{
				Top:         0,
				Size:        fontsize - 1,
				Extrapolate: false,
				Style:       fontstyle,
				Color:       textcolour,
			})
		})
		m.Col(colwidth-indent, func() {
			m.Text(line, props.Text{
				Top:         0,
				Size:        fontsize - 1,
				Extrapolate: false,
				Style:       fontstyle,
				Color:       textcolour,
			})
		})
	})
}

// New style PDFreportwriter row writer
func pdf_write_row_tuple_indent(m pdf.Maroto, line1 string, line2 string, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color, indent uint) {
	if line2 == "" || line2 == " " {
		textcolour = pdfRed()
	}

	m.Row(rowheight, func() {
		if indent == 0 {
			indent = 1
		} else {
			m.Col(indent, func() {
				m.Text("", props.Text{
					Top:         0,
					Size:        fontsize,
					Extrapolate: false,
					Style:       fontstyle,
					Color:       textcolour,
				})
			})
		}
		m.Col(indent+1, func() {
			m.Text(line1, props.Text{
				Top:         0,
				Size:        fontsize,
				Extrapolate: false,
				Style:       fontstyle,
				Color:       textcolour,
			})
		})
		m.Col(colwidth-indent-1, func() {
			m.Text(line2, props.Text{
				Top:         0,
				Size:        fontsize,
				Extrapolate: false,
				Style:       fontstyle,
				Color:       textcolour,
			})
		})
	})
}

// New style PDFreportwriter list writer
func pdf_write_list(m pdf.Maroto, lines []string, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color) {
	var indent uint = 1

	for line := range lines {
		if lines[line] == "" || lines[line] == " " {
			textcolour = pdfRed()
		} else {
			textcolour = pdfBlack()
		}
		// fmt.Printf("%s\n", lines[line])
		m.Row(rowheight, func() {
			m.Col(indent, func() {
				m.Text(indentsymb, props.Text{
					Top:         0,
					Size:        fontsize - 1,
					Extrapolate: false,
					Style:       fontstyle,
					Color:       textcolour,
				})
			})
			m.Col(colwidth-indent, func() {
				m.Text(lines[line], props.Text{
					Top:         0,
					Size:        fontsize - 1,
					Extrapolate: false,
					Style:       fontstyle,
					Color:       textcolour,
				})
			})
		})
	}
}

// New style PDFreportwriter sublevel 1 list writer
func pdf_write_list_sub1(m pdf.Maroto, lines []string, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color) {
	var indent uint = 2

	for line := range lines {
		if lines[line] == "" || lines[line] == " " {
			textcolour = pdfRed()
		} else {
			textcolour = pdfBlack()
		}
		// fmt.Printf("%s\n", lines[line])
		m.Row(rowheight, func() {
			m.Col(indent, func() {
				m.Text(indentsymb, props.Text{
					Top:         0,
					Size:        fontsize - 1,
					Extrapolate: false,
					Style:       fontstyle,
					Color:       textcolour,
				})
			})
			m.Col(colwidth-indent, func() {
				m.Text(lines[line], props.Text{
					Top:         0,
					Size:        fontsize - 1,
					Extrapolate: false,
					Style:       fontstyle,
					Color:       textcolour,
				})
			})
		})
	}
}

func pdf_write_creators(m pdf.Maroto, data Yoda18Metadata, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color) {
	var ind1 uint = 1
	// var ind2 uint = 2
	pdf_write_row(m, "Creators", rowheight, colwidth, consts.Bold, pdfBlack())
	for i := range data.Creator {
		pdf_write_row(m, fmt.Sprintf("%s %s", data.Creator[i].Name.GivenName, data.Creator[i].Name.FamilyName), rowheight, colwidth, consts.Normal, pdfBlack())
		for j := range data.Creator[i].Affiliation {
			pdf_write_row_indent(m, data.Creator[i].Affiliation[j], rowheight, colwidth, consts.Normal, pdfBlack(), ind1)
		}
		for k := range data.Creator[i].PersonIdentifier {
			pdf_write_row_indent(m, fmt.Sprintf("(%s) %s", data.Creator[i].PersonIdentifier[k].NameIdentifierScheme, data.Creator[i].PersonIdentifier[k].NameIdentifier),
				rowheight, colwidth, consts.Normal, pdfBlack(), ind1)
		}
	}
}

func pdf_write_contributors(m pdf.Maroto, data Yoda18Metadata, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color) {
	var ind1 uint = 1
	// var ind2 uint = 2
	pdf_write_row(m, "Contributors", rowheight, colwidth, consts.Bold, pdfBlack())
	for i := range data.Contributor {
		pdf_write_row(m, fmt.Sprintf("%s %s", data.Contributor[i].Name.GivenName, data.Contributor[i].Name.FamilyName), rowheight, colwidth, consts.Normal, pdfBlack())
		pdf_write_row_indent(m, fmt.Sprintf("%s", data.Contributor[i].ContributorType), rowheight, colwidth, consts.Normal, pdfBlack(), ind1)
		for j := range data.Contributor[i].Affiliation {
			pdf_write_row_indent(m, data.Contributor[i].Affiliation[j], rowheight, colwidth, consts.Normal, pdfBlack(), ind1)
		}
		for k := range data.Contributor[i].PersonIdentifier {
			pdf_write_row_indent(m, fmt.Sprintf("(%s) %s", data.Contributor[i].PersonIdentifier[k].NameIdentifierScheme, data.Contributor[i].PersonIdentifier[k].NameIdentifier),
				rowheight, colwidth, consts.Normal, pdfBlack(), ind1)
		}
	}
}

// new function for writing funders
func pdf_write_funding(m pdf.Maroto, data Yoda18Metadata, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color) {
	pdf_write_row(m, "Funding references", rowheight, colwidth, consts.Bold, pdfBlack())
	for i := range data.FundingReference {
		pdf_write_row_tuple_indent(m, data.FundingReference[i].FunderName, data.FundingReference[i].AwardNumber, rowheight, colwidth, consts.Normal, pdfBlack(), 1)
	}
}

//new functions for writing related data packages
func pdf_write_related(m pdf.Maroto, data Yoda18Metadata, rowheight float64, colwidth uint, fontstyle consts.Style, textcolour color.Color) {
	pdf_write_row(m, "Related datapackages", rowheight, colwidth, consts.Bold, pdfBlack())
	for i := range data.RelatedDatapackage {
		pdf_write_row(m, data.RelatedDatapackage[i].RelationType, rowheight, colwidth, consts.Normal, pdfBlack())
		if data.RelatedDatapackage[i].PersistentIdentifier.IdentifierScheme == "" {
			pdf_write_row_indent(m, "(string)"+data.RelatedDatapackage[i].PersistentIdentifier.Identifier, rowheight, colwidth, consts.Normal, pdfBlack(), 1)
		} else {
			pdf_write_row_indent(m, "("+data.RelatedDatapackage[i].PersistentIdentifier.IdentifierScheme+") "+data.RelatedDatapackage[i].PersistentIdentifier.Identifier, rowheight,
				colwidth, consts.Normal, pdfBlack(), 1)
		}
		pdf_write_row_indent(m, data.RelatedDatapackage[i].Title, rowheight, colwidth, consts.Normal, pdfBlack(), 1)
	}
}

// pdf_write_row(doc, "Tag", rowheight, colwidth, consts.Bold, pdfBlack())
// pdf_write_list(doc, data.Tag, rowheight, colwidth, consts.Normal, pdfBlack())
// pdf_write_list_sub1(doc, data.Tag, rowheight, colwidth, consts.Normal, pdfBlack())

// This is an old function which generates "formatted" strings as output
func pdf_create_and_dump(fname string, sarr []string) {

	// Do things

	var colwidth uint = 12
	var rowheight float64 = 6

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	//m.SetBorder(true)

	m.Row(10, func() {
		m.Col(colwidth, func() {
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
		pdf_write_row(m, fmt.Sprintln(ele), rowheight, colwidth, consts.Normal, pdfBlack())
	}

	run_time := time.Now()

	m.Line(4)
	m.Row(4, func() {
		m.Col(colwidth, func() {
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

// This is an old function which generates "formatted" strings as output
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
