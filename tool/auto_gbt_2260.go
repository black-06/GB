package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

func AutoGBT2260() error {
	file, err := os.Create("../go/gb_t_2260.go")
	if err != nil {
		return err
	}
	mca, err := os.Open("../data/GB_T_2260_2007/mca_202012.csv")
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("package gb\n\n" +
		"// GBT2260 for \"Codes for the administrative divisions of the People's Republic of China\"\n" +
		"type GBT2260 struct {\n\tDivisionMap map[string]Division\n}\n\n" +
		"func NewGBT2260FromMCA202012() *GBT2260 {\n\treturn &GBT2260{DivisionMap: MCA202012}\n}\n\n" +
		"func (p *GBT2260) Validate(data string) bool {\n\t_, ok := p.DivisionMap[data]\n\treturn ok\n}\n\n" +
		"func (p *GBT2260) Fake() string {\n\tvar data string\n\tfor data = range p.DivisionMap {\n\t\tbreak\n\t}\n\treturn data\n}\n\n" +
		"// Division is administrative division\n\n" +
		"type Division struct {\n\tProvince string\n\tCity     string\n\tCounty   string\n}\n\n" +
		"var (\n\tMCA202012 = map[string]Division{")
	if err != nil {
		return err
	}

	var province, city string
	reader := csv.NewReader(mca)
	// skip title
	_, _ = reader.Read()
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		code, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			return err
		}
		if code%10000 == 0 {
			province = row[1]
			city = row[1]
		} else if code%100 == 0 {
			city = row[1]
		}
		_, err = writer.WriteString("\n\t\t\"" + row[0] + "\": {" +
			"\n\t\t\tProvince: \"" + province + "\"," +
			"\n\t\t\tCity:     \"" + city + "\"," +
			"\n\t\t\tCounty:   \"" + row[1] + "\"," +
			"\n\t\t},")
	}

	_, err = writer.WriteString("\n\t}\n)\n")
	if err != nil {
		return err
	}
	return writer.Flush()
}
