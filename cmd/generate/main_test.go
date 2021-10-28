package main

import (
	"bufio"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_generate(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generate()
		})
	}
}

func Test_fileNamePath(t *testing.T) {
	type args struct {
		filePath    string
		persistPath string
		fileName    string
	}
	tests := []struct {
		name        string
		args        args
		wantFile    string
		wantPersist string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, gotPersist := fileNamePath(tt.args.filePath, tt.args.persistPath, tt.args.fileName)
			if gotFile != tt.wantFile {
				t.Errorf("fileNamePath() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
			if gotPersist != tt.wantPersist {
				t.Errorf("fileNamePath() gotPersist = %v, want %v", gotPersist, tt.wantPersist)
			}
		})
	}
}

func Test_createWriters(t *testing.T) {
	type args struct {
		persistPath string
		filePath    string
		fileName    string
	}
	tests := []struct {
		name string
		args args
		want []*bufio.Writer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createWriters(tt.args.persistPath, tt.args.filePath, tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createWriters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateStruct(t *testing.T) {
	type args struct {
		filePath    string
		persistPath string
		fileName    string
		annotypes   []string
		writers     []*bufio.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateStruct(tt.args.filePath, tt.args.persistPath, tt.args.fileName, tt.args.annotypes, tt.args.writers)
		})
	}
}

func Test_structLine(t *testing.T) {
	type args struct {
		fieldName  string
		jsonEtc    []string
		isSeparate bool
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := structLine(tt.args.fieldName, tt.args.jsonEtc, tt.args.isSeparate)
			if got != tt.want {
				t.Errorf("structLine() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("structLine() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_bestGuessOnType(t *testing.T) {
	type args struct {
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestGuessOnType(tt.args.fieldName); got != tt.want {
				t.Errorf("bestGuessOnType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPackageFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPackageFromPath(tt.args.path); got != tt.want {
				t.Errorf("getPackageFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extract(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name          string
		args          args
		wantFieldType string
		wantFieldName string
		wantOmit      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFieldType, gotFieldName, gotOmit := extract(tt.args.field)
			if gotFieldType != tt.wantFieldType {
				t.Errorf("extract() gotFieldType = %v, want %v", gotFieldType, tt.wantFieldType)
			}
			if gotFieldName != tt.wantFieldName {
				t.Errorf("extract() gotFieldName = %v, want %v", gotFieldName, tt.wantFieldName)
			}
			if gotOmit != tt.wantOmit {
				t.Errorf("extract() gotOmit = %v, want %v", gotOmit, tt.wantOmit)
			}
		})
	}
}
