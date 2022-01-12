package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/opesun/goquery"
)

func UpdateSource(id string) error {
	gbMap, err := ReadSource()
	if err != nil {
		return err
	}
	rawMsg, ok := gbMap[id]
	if !ok {
		return fmt.Errorf("gb %v not found in source.json", id)
	}

	gb := Gb{}
	err = json.Unmarshal(rawMsg, &gb)
	if err != nil {
		return err
	}
	if gb.Source == "" {
		return fmt.Errorf("gb %v source not found in source.json", id)
	}

	err = GoQuery(&gb)
	if err != nil {
		return err
	}

	rawMsg, err = json.Marshal(gb)
	if err != nil {
		return err
	}
	gbMap[id] = rawMsg
	source, err := json.Marshal(gbMap)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("source.json", source, 0666)
}

type Gb struct {
	ID                 string `json:"id"`
	Source             string `json:"source"`
	NameCN             string `json:"name_cn"`
	NameUS             string `json:"name_us"`
	Status             string `json:"status"`
	CCS                string `json:"ccs"`
	ICS                string `json:"ics"`
	PublishingDate     string `json:"publishing_date"`
	ImplementationDate string `json:"implementation_date"`
	ManagementUnit     string `json:"management_unit"`
	ResponsibleUnit    string `json:"responsible_unit"`
	PublishingUnit     string `json:"publishing_unit"`
	Remarks            string `json:"remarks"`
}

func (p *Gb) String() string {
	builder := strings.Builder{}
	builder.WriteString("### " + p.ID + "  \n\n")
	builder.WriteString("来源/source : " + p.Source + "  \n\n")
	builder.WriteString("| | |\n|:---|:---|\n")
	builder.WriteString("| 中文标准名称 | " + p.NameCN + " |\n")
	builder.WriteString("| 英文标准名称 | " + p.NameUS + " |\n")
	builder.WriteString("| 标准状态 | " + p.Status + " |\n")
	builder.WriteString("| 中国标准分类号 | " + p.CCS + " |\n")
	builder.WriteString("| 国际标准分类号 | " + p.ICS + " |\n")
	builder.WriteString("| 发布日期 | " + p.PublishingDate + " |\n")
	builder.WriteString("| 实施日期 | " + p.ImplementationDate + " |\n")
	builder.WriteString("| 主管部门 | " + p.ManagementUnit + " |\n")
	builder.WriteString("| 归口单位 | " + p.ResponsibleUnit + " |\n")
	builder.WriteString("| 发布单位 | " + p.PublishingUnit + " |\n")
	builder.WriteString("| 备注 | " + p.Remarks + " |\n")
	return builder.String()
}

func ReadSource() (map[string]json.RawMessage, error) {
	file, err := os.Open("source.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	gbMap := make(map[string]json.RawMessage)
	err = json.Unmarshal(data, &gbMap)
	if err != nil {
		return nil, err
	}
	return gbMap, nil
}

func GoQuery(gb *Gb) error {
	doc, err := goquery.ParseUrl(gb.Source)
	if err != nil {
		return err
	}
	nodes := doc.Find(".bor2")
	table := nodes.Find(".tdlist")
	tds := table.Find("td")

	gb.NameCN = strings.Split(tds.Eq(0).Text(), "：")[1]
	gb.NameUS = strings.Split(tds.Eq(2).Text(), "：")[1]
	gb.Status = strings.Split(tds.Eq(3).Text(), "：")[1]

	rows := nodes.Find(".col-xs-12")
	gb.CCS = SplitTable(rows.Eq(1).Text())
	gb.ICS = SplitTable(rows.Eq(3).Text())
	gb.PublishingDate = SplitTable(rows.Eq(5).Text())
	gb.ImplementationDate = SplitTable(rows.Eq(7).Text())
	gb.ManagementUnit = SplitTable(rows.Eq(7).Text())
	gb.ResponsibleUnit = SplitTable(rows.Eq(9).Text())
	gb.PublishingUnit = SplitTable(rows.Eq(11).Text())
	gb.Remarks = SplitTable(rows.Eq(13).Text())
	return nil
}

func SplitTable(text string) string {
	row := strings.ReplaceAll(text, "\t", "")
	row = strings.ReplaceAll(row, "\n", "")
	return row
}
