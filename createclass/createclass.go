package createclass

import (
	"fmt"
	"strings"

	"github.com/kagankuscu/CSharpClass/types"
)

func CreateCSharpClassWithType(data *string, className *string, properties *[]types.Class) string {
	var b strings.Builder

	splitedLine := strings.Split(*data, "\n")
	for _, line := range splitedLine {
		values := strings.Split(line, ",")

		if len(values) != 1 {
			newClass := fmt.Sprintf("new %v { ", *className)
			b.WriteString(newClass)

			properties := *properties

			for j, value := range values {
				s := fmt.Sprintf(checkType(properties[j]), properties[j].PropertyName, value)
				b.WriteString(s)
			}
			b.WriteString("} \n")
		}
	}

	return b.String()
}

func checkType(data types.Class) string {
	switch strings.ToLower(data.PropertyType) {
	case "int":
		return "%v = %v, "
	case "string":
		return "%v = \"%v\", "
	default:
		fmt.Println("Supported Type is int, string")
		return "Supported Type is int, string"
	}
}

func CreateDataCsv(data *[][]string) (*[]types.Class, string) {
	var arr []types.Class
	f := *data
	var sb strings.Builder

	for _, d := range f[1:] {
		for i, v := range f[0] {
			s := strings.Split(d[i], ":")
			arr = append(arr, types.Class{PropertyName: v, PropertyType: s[1]})
			if len(f[0])-1 == i {
				sb.WriteString(fmt.Sprintf("%v", s[0]))
			} else {
				sb.WriteString(fmt.Sprintf("%v,", s[0]))
			}
		}
		sb.WriteString("\n")
	}
	return &arr, sb.String()
}
