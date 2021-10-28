package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
)

func main() {
	generate()
}

func generate() {
	fmt.Println("Please enter file path for file generation otherwise output will go to stdout")
	var isFile string
	fmt.Scanln(&isFile)

	fmt.Println("Please enter persistence path for separating persistence layer model")
	var isSeparate string
	fmt.Scanln(&isSeparate)

	fmt.Println("To end generation press ctrl+d on keyboard. To exit program: ctrl+c")
	fmt.Println("Enter your filename (e.g. catalog_item): ")
	var fileName string
	fmt.Scanln(&fileName)

	fmt.Println("Enter annotations desired (e.g. json|jsonapi|dynamodbav): ")
	var annotations string
	fmt.Scanln(&annotations)
	annotypes := strings.Split(annotations, "|")

	fileName = strcase.ToSnake(fileName)
	writers := createWriters(isSeparate, isFile, fileName)
	generateStruct(isFile, isSeparate, fileName, annotypes, writers)
	generate()
}

func fileNamePath(filePath, persistPath, fileName string) (file string, persist string) {
	fileInfo, _ := os.Stat(filePath)
	mode := fileInfo.Mode()
	if mode.IsRegular() {
		file = filePath
	} else {
		file = filePath + "/" + fileName + ".go"
	}
	if persistPath != "" {
		fileInfo, _ := os.Stat(persistPath)
		mode := fileInfo.Mode()
		if mode.IsRegular() {
			persist = persistPath
		} else {
			file = filePath + "/" + fileName + ".go"
		}
	}
	return
}

func createWriters(persistPath, filePath string, fileName string) []*bufio.Writer {
	var writers []*bufio.Writer
	oneOrTwo := 1
	if persistPath != "" {
		oneOrTwo = 2
	}
	for i := 0; i < oneOrTwo; i++ {
		if filePath != "" {
			name, persist := fileNamePath(filePath, persistPath, fileName)
			if i == 1 {
				name = persist
			}
			f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
			if err != nil {
				fmt.Println(err)
				fmt.Println("There was an issue writing to specified directory, will default to stdout")
				writers = append(writers, bufio.NewWriter(os.Stdout))
			} else {
				writers = append(writers, bufio.NewWriter(f))
			}
		} else {
			writers = append(writers, bufio.NewWriter(os.Stdout))
		}
	}
	return writers
}

func generateStruct(filePath, persistPath string, fileName string, annotypes []string, writers []*bufio.Writer) {
	scanner := bufio.NewScanner(os.Stdin)
	structName := strcase.ToCamel(fileName)
	structDeclaration := "\ntype " + structName + " struct { "
	closingBrace := " }\n"
	var fieldName string

	fmt.Println("Enter field name e.g. price_cents, size-SizeType")

	for i, w := range writers {
		var path string
		if i == 1 {
			path = getPackageFromPath(persistPath)
		} else {
			path = getPackageFromPath(filePath)
		}
		if path != "" {
			w.WriteString("package " + path + "\n")
		}
		w.WriteString(structDeclaration + "\n")
	}
	for scanner.Scan() {
		fmt.Println("Enter field name e.g. price_cents")
		fieldName = scanner.Text()
		model, persist := structLine(fieldName, annotypes, persistPath != "")
		for i, w := range writers {
			if i == 1 {
				w.WriteString("\t" + persist + "\n")
			} else {
				w.WriteString("\t" + model + "\n")
			}
		}
	}
	for _, w := range writers {
		w.WriteString(closingBrace)
		w.Flush()
	}

}

// if fieldName has a dash followed by a type or omit we include omitempty and the type in generated code/annotations
// example would be price_cents-int, topic_ID or student-*Student or cost-omit-uint (type should always be last) in which case the
func structLine(fieldName string, jsonEtc []string, isSeparate bool) (string, string) {
	var line strings.Builder
	var dynamoLine strings.Builder
	var annos []string
	sort.Sort(sort.Reverse(sort.StringSlice(jsonEtc)))
	fieldType, name, omit := extract(fieldName)
	strcase.ConfigureAcronym("id", "ID")
	strcase.ConfigureAcronym("pk", "PK")
	strcase.ConfigureAcronym("sk", "SK")

	field := strcase.ToCamel(name)
	name = strings.ToLower(name)

	line.WriteString(field + " ")
	line.WriteString(fieldType)
	line.WriteString(" `")
	for _, an := range jsonEtc {
		var eachAnno strings.Builder

		switch an {
		case "jsonapi":
			eachAnno.WriteString("jsonapi:\"attr," + name)
			if omit {
				eachAnno.WriteString(",omitempty")
			}
			eachAnno.WriteString("\"")
		case "json":
			eachAnno.WriteString("json:\"" + name)
			if omit {
				eachAnno.WriteString(",omitempty")
			}
			eachAnno.WriteString("\"")
		case "dynamodbav":
			if !isSeparate {
				eachAnno.WriteString("dynamodbav:\"" + name)
				if omit {
					eachAnno.WriteString(",-")
				}
				eachAnno.WriteString("\"")
			}
		}
		annos = append(annos, eachAnno.String())
	}
	line.WriteString(strings.Join(annos, " "))
	line.WriteString("`")

	if isSeparate {
		dynamoLine.WriteString(field + " ")
		dynamoLine.WriteString(fieldType)
		dynamoLine.WriteString(" `dynamodbav:\"" + name)
		if omit {
			dynamoLine.WriteString(",-")
		}
		dynamoLine.WriteString("\"`")
	}

	return line.String(), dynamoLine.String()
}

// will return a type based on the field name when a type isn't defined
func bestGuessOnType(fieldName string) string {
	field := strings.ToLower(fieldName)
	if strings.HasSuffix(field, "cents") {
		return "*int"
	}
	if strings.HasPrefix(field, "is_") {
		return "bool"
	}
	if strings.HasSuffix(field, "uuid") {
		return "string"
	}
	if strings.HasSuffix(field, "id") || strings.HasSuffix(field, "count") || strings.HasSuffix(field, "quantity") {
		return "int"
	}
	if field[len(field)-1:] == "s" {
		return "[]" + strcase.ToCamel(field[:len(field)-1])
	}

	return "string"
}

func getPackageFromPath(path string) string {
	lastDir := strings.LastIndex(path, "/")
	pkg := path[lastDir+1:]

	if strings.HasSuffix(pkg, ".go") {
		return ""
	}
	return pkg
}

// will return the fieldname, fieldtype and if to include omit in definition
func extract(field string) (fieldType string, fieldName string, omit bool) {
	if strings.HasSuffix(field, "-omit") {
		omit = true
		fieldName = strings.Replace(field, "-omit", "", 1)
		fieldType = bestGuessOnType(fieldName)
		return
	}

	typePos := strings.LastIndex(field, "-")
	if typePos < 1 {
		fieldName = field
		fieldType = bestGuessOnType(fieldName)
		return
	}

	fieldType = field[typePos+1:]
	fieldWOType := strings.Replace(field, "-"+fieldType, "", 1)
	fieldName = fieldWOType

	if strings.HasSuffix(field, "-omit") {
		omit = true
		fieldType = strings.Replace(fieldType, "-omit", "", 1)
		fieldName = strings.Replace(fieldName, "-omit", "", 1)
		return
	}

	return fieldType, fieldName, omit
}
