package gb

import (
	mapset "github.com/deckarep/golang-set"
)

func init() {
	RegisterGB("GBT4762", &GBT4762{})
}

type GBT4762 struct {
}

func (p *GBT4762) Validate(data string) bool {
	return PoliticalAffiliationSet.Contains(data)
}

func (p *GBT4762) Fake() string {
	return PoliticalAffiliations[random.Intn(len(PoliticalAffiliations))].(string)
}

var (
	PoliticalAffiliations = []interface{}{
		"中国共产党党员", "中共党员",
		"中国共产党预备党员", "中共预备党员",
		"中国共产主义青年团团员", "共青团员", "中国新民主主义青年团团员", "青年团员",
		"中国国民党革命委员会会员", "民革会员",
		"中国民主同盟盟员", "民盟盟员",
		"中国民主建国会会员", "民建会员",
		"中国民主促进会会员", "民进会员",
		"中国农工民主党单元", "农工党党员",
		"中国致公党党员", "致公党党员",
		"九三学社社员",
		"台湾民主自治同盟盟员", "台盟盟员",
		"无党派民主人士",
		"群众",
	}
	PoliticalAffiliationSet = mapset.NewThreadUnsafeSetFromSlice(PoliticalAffiliations)
)
