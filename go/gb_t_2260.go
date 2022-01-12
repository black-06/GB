package gb

func init() {
	RegisterGB("GBT2260", NewGBT2260FromMCA202012())
}

// GBT2260 for "Codes for the administrative divisions of the People's Republic of China"
type GBT2260 struct {
	DivisionMap map[string]Division
}

func NewGBT2260FromMCA202012() *GBT2260 {
	return &GBT2260{DivisionMap: MCA202012}
}

func (p *GBT2260) Validate(data string) bool {
	_, ok := p.DivisionMap[data]
	return ok
}

func (p *GBT2260) Fake() string {
	var data string
	for data = range p.DivisionMap {
		break
	}
	return data
}

// Division is administrative division

type Division struct {
	Province string
	City     string
	County   string
}

var (
	MCA202012 = map[string]Division{
		"110000": {
			Province: "北京市",
			City:     "北京市",
			County:   "北京市",
		},
		"110101": {
			Province: "北京市",
			City:     "北京市",
			County:   "东城区",
		},
		"110102": {
			Province: "北京市",
			City:     "北京市",
			County:   "西城区",
		},
		"110105": {
			Province: "北京市",
			City:     "北京市",
			County:   "朝阳区",
		},
		"110106": {
			Province: "北京市",
			City:     "北京市",
			County:   "丰台区",
		},
		"110107": {
			Province: "北京市",
			City:     "北京市",
			County:   "石景山区",
		},
		"110108": {
			Province: "北京市",
			City:     "北京市",
			County:   "海淀区",
		},
		"110109": {
			Province: "北京市",
			City:     "北京市",
			County:   "门头沟区",
		},
		"110111": {
			Province: "北京市",
			City:     "北京市",
			County:   "房山区",
		},
		"110112": {
			Province: "北京市",
			City:     "北京市",
			County:   "通州区",
		},
		"110113": {
			Province: "北京市",
			City:     "北京市",
			County:   "顺义区",
		},
		"110114": {
			Province: "北京市",
			City:     "北京市",
			County:   "昌平区",
		},
		"110115": {
			Province: "北京市",
			City:     "北京市",
			County:   "大兴区",
		},
		"110116": {
			Province: "北京市",
			City:     "北京市",
			County:   "怀柔区",
		},
		"110117": {
			Province: "北京市",
			City:     "北京市",
			County:   "平谷区",
		},
		"110118": {
			Province: "北京市",
			City:     "北京市",
			County:   "密云区",
		},
		"110119": {
			Province: "北京市",
			City:     "北京市",
			County:   "延庆区",
		},
		"120000": {
			Province: "天津市",
			City:     "天津市",
			County:   "天津市",
		},
		"120101": {
			Province: "天津市",
			City:     "天津市",
			County:   "和平区",
		},
		"120102": {
			Province: "天津市",
			City:     "天津市",
			County:   "河东区",
		},
		"120103": {
			Province: "天津市",
			City:     "天津市",
			County:   "河西区",
		},
		"120104": {
			Province: "天津市",
			City:     "天津市",
			County:   "南开区",
		},
		"120105": {
			Province: "天津市",
			City:     "天津市",
			County:   "河北区",
		},
		"120106": {
			Province: "天津市",
			City:     "天津市",
			County:   "红桥区",
		},
		"120110": {
			Province: "天津市",
			City:     "天津市",
			County:   "东丽区",
		},
		"120111": {
			Province: "天津市",
			City:     "天津市",
			County:   "西青区",
		},
		"120112": {
			Province: "天津市",
			City:     "天津市",
			County:   "津南区",
		},
		"120113": {
			Province: "天津市",
			City:     "天津市",
			County:   "北辰区",
		},
		"120114": {
			Province: "天津市",
			City:     "天津市",
			County:   "武清区",
		},
		"120115": {
			Province: "天津市",
			City:     "天津市",
			County:   "宝坻区",
		},
		"120116": {
			Province: "天津市",
			City:     "天津市",
			County:   "滨海新区",
		},
		"120117": {
			Province: "天津市",
			City:     "天津市",
			County:   "宁河区",
		},
		"120118": {
			Province: "天津市",
			City:     "天津市",
			County:   "静海区",
		},
		"120119": {
			Province: "天津市",
			City:     "天津市",
			County:   "蓟州区",
		},
		"130000": {
			Province: "河北省",
			City:     "河北省",
			County:   "河北省",
		},
		"130100": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "石家庄市",
		},
		"130102": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "长安区",
		},
		"130104": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "桥西区",
		},
		"130105": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "新华区",
		},
		"130107": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "井陉矿区",
		},
		"130108": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "裕华区",
		},
		"130109": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "藁城区",
		},
		"130110": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "鹿泉区",
		},
		"130111": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "栾城区",
		},
		"130121": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "井陉县",
		},
		"130123": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "正定县",
		},
		"130125": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "行唐县",
		},
		"130126": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "灵寿县",
		},
		"130127": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "高邑县",
		},
		"130128": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "深泽县",
		},
		"130129": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "赞皇县",
		},
		"130130": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "无极县",
		},
		"130131": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "平山县",
		},
		"130132": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "元氏县",
		},
		"130133": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "赵县",
		},
		"130181": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "辛集市",
		},
		"130183": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "晋州市",
		},
		"130184": {
			Province: "河北省",
			City:     "石家庄市",
			County:   "新乐市",
		},
		"130200": {
			Province: "河北省",
			City:     "唐山市",
			County:   "唐山市",
		},
		"130202": {
			Province: "河北省",
			City:     "唐山市",
			County:   "路南区",
		},
		"130203": {
			Province: "河北省",
			City:     "唐山市",
			County:   "路北区",
		},
		"130204": {
			Province: "河北省",
			City:     "唐山市",
			County:   "古冶区",
		},
		"130205": {
			Province: "河北省",
			City:     "唐山市",
			County:   "开平区",
		},
		"130207": {
			Province: "河北省",
			City:     "唐山市",
			County:   "丰南区",
		},
		"130208": {
			Province: "河北省",
			City:     "唐山市",
			County:   "丰润区",
		},
		"130209": {
			Province: "河北省",
			City:     "唐山市",
			County:   "曹妃甸区",
		},
		"130224": {
			Province: "河北省",
			City:     "唐山市",
			County:   "滦南县",
		},
		"130225": {
			Province: "河北省",
			City:     "唐山市",
			County:   "乐亭县",
		},
		"130227": {
			Province: "河北省",
			City:     "唐山市",
			County:   "迁西县",
		},
		"130229": {
			Province: "河北省",
			City:     "唐山市",
			County:   "玉田县",
		},
		"130281": {
			Province: "河北省",
			City:     "唐山市",
			County:   "遵化市",
		},
		"130283": {
			Province: "河北省",
			City:     "唐山市",
			County:   "迁安市",
		},
		"130284": {
			Province: "河北省",
			City:     "唐山市",
			County:   "滦州市",
		},
		"130300": {
			Province: "河北省",
			City:     "秦皇岛市",
			County:   "秦皇岛市",
		},
		"130302": {
			Province: "河北省",
			City:     "秦皇岛市",
			County:   "海港区",
		},
		"130303": {
			Province: "河北省",
			City:     "秦皇岛市",
			County:   "山海关区",
		},
		"130304": {
			Province: "河北省",
			City:     "秦皇岛市",
			County:   "北戴河区",
		},
		"130306": {
			Province: "河北省",
			City:     "秦皇岛市",
			County:   "抚宁区",
		},
		"130321": {
			Province: "河北省",
			City:     "秦皇岛市",
			County:   "青龙满族自治县",
		},
		"130322": {
			Province: "河北省",
			City:     "秦皇岛市",
			County:   "昌黎县",
		},
		"130324": {
			Province: "河北省",
			City:     "秦皇岛市",
			County:   "卢龙县",
		},
		"130400": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "邯郸市",
		},
		"130402": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "邯山区",
		},
		"130403": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "丛台区",
		},
		"130404": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "复兴区",
		},
		"130406": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "峰峰矿区",
		},
		"130407": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "肥乡区",
		},
		"130408": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "永年区",
		},
		"130423": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "临漳县",
		},
		"130424": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "成安县",
		},
		"130425": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "大名县",
		},
		"130426": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "涉县",
		},
		"130427": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "磁县",
		},
		"130430": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "邱县",
		},
		"130431": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "鸡泽县",
		},
		"130432": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "广平县",
		},
		"130433": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "馆陶县",
		},
		"130434": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "魏县",
		},
		"130435": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "曲周县",
		},
		"130481": {
			Province: "河北省",
			City:     "邯郸市",
			County:   "武安市",
		},
		"130500": {
			Province: "河北省",
			City:     "邢台市",
			County:   "邢台市",
		},
		"130502": {
			Province: "河北省",
			City:     "邢台市",
			County:   "襄都区",
		},
		"130503": {
			Province: "河北省",
			City:     "邢台市",
			County:   "信都区",
		},
		"130505": {
			Province: "河北省",
			City:     "邢台市",
			County:   "任泽区",
		},
		"130506": {
			Province: "河北省",
			City:     "邢台市",
			County:   "南和区",
		},
		"130522": {
			Province: "河北省",
			City:     "邢台市",
			County:   "临城县",
		},
		"130523": {
			Province: "河北省",
			City:     "邢台市",
			County:   "内丘县",
		},
		"130524": {
			Province: "河北省",
			City:     "邢台市",
			County:   "柏乡县",
		},
		"130525": {
			Province: "河北省",
			City:     "邢台市",
			County:   "隆尧县",
		},
		"130528": {
			Province: "河北省",
			City:     "邢台市",
			County:   "宁晋县",
		},
		"130529": {
			Province: "河北省",
			City:     "邢台市",
			County:   "巨鹿县",
		},
		"130530": {
			Province: "河北省",
			City:     "邢台市",
			County:   "新河县",
		},
		"130531": {
			Province: "河北省",
			City:     "邢台市",
			County:   "广宗县",
		},
		"130532": {
			Province: "河北省",
			City:     "邢台市",
			County:   "平乡县",
		},
		"130533": {
			Province: "河北省",
			City:     "邢台市",
			County:   "威县",
		},
		"130534": {
			Province: "河北省",
			City:     "邢台市",
			County:   "清河县",
		},
		"130535": {
			Province: "河北省",
			City:     "邢台市",
			County:   "临西县",
		},
		"130581": {
			Province: "河北省",
			City:     "邢台市",
			County:   "南宫市",
		},
		"130582": {
			Province: "河北省",
			City:     "邢台市",
			County:   "沙河市",
		},
		"130600": {
			Province: "河北省",
			City:     "保定市",
			County:   "保定市",
		},
		"130602": {
			Province: "河北省",
			City:     "保定市",
			County:   "竞秀区",
		},
		"130606": {
			Province: "河北省",
			City:     "保定市",
			County:   "莲池区",
		},
		"130607": {
			Province: "河北省",
			City:     "保定市",
			County:   "满城区",
		},
		"130608": {
			Province: "河北省",
			City:     "保定市",
			County:   "清苑区",
		},
		"130609": {
			Province: "河北省",
			City:     "保定市",
			County:   "徐水区",
		},
		"130623": {
			Province: "河北省",
			City:     "保定市",
			County:   "涞水县",
		},
		"130624": {
			Province: "河北省",
			City:     "保定市",
			County:   "阜平县",
		},
		"130626": {
			Province: "河北省",
			City:     "保定市",
			County:   "定兴县",
		},
		"130627": {
			Province: "河北省",
			City:     "保定市",
			County:   "唐县",
		},
		"130628": {
			Province: "河北省",
			City:     "保定市",
			County:   "高阳县",
		},
		"130629": {
			Province: "河北省",
			City:     "保定市",
			County:   "容城县",
		},
		"130630": {
			Province: "河北省",
			City:     "保定市",
			County:   "涞源县",
		},
		"130631": {
			Province: "河北省",
			City:     "保定市",
			County:   "望都县",
		},
		"130632": {
			Province: "河北省",
			City:     "保定市",
			County:   "安新县",
		},
		"130633": {
			Province: "河北省",
			City:     "保定市",
			County:   "易县",
		},
		"130634": {
			Province: "河北省",
			City:     "保定市",
			County:   "曲阳县",
		},
		"130635": {
			Province: "河北省",
			City:     "保定市",
			County:   "蠡县",
		},
		"130636": {
			Province: "河北省",
			City:     "保定市",
			County:   "顺平县",
		},
		"130637": {
			Province: "河北省",
			City:     "保定市",
			County:   "博野县",
		},
		"130638": {
			Province: "河北省",
			City:     "保定市",
			County:   "雄县",
		},
		"130681": {
			Province: "河北省",
			City:     "保定市",
			County:   "涿州市",
		},
		"130682": {
			Province: "河北省",
			City:     "保定市",
			County:   "定州市",
		},
		"130683": {
			Province: "河北省",
			City:     "保定市",
			County:   "安国市",
		},
		"130684": {
			Province: "河北省",
			City:     "保定市",
			County:   "高碑店市",
		},
		"130700": {
			Province: "河北省",
			City:     "张家口市",
			County:   "张家口市",
		},
		"130702": {
			Province: "河北省",
			City:     "张家口市",
			County:   "桥东区",
		},
		"130703": {
			Province: "河北省",
			City:     "张家口市",
			County:   "桥西区",
		},
		"130705": {
			Province: "河北省",
			City:     "张家口市",
			County:   "宣化区",
		},
		"130706": {
			Province: "河北省",
			City:     "张家口市",
			County:   "下花园区",
		},
		"130708": {
			Province: "河北省",
			City:     "张家口市",
			County:   "万全区",
		},
		"130709": {
			Province: "河北省",
			City:     "张家口市",
			County:   "崇礼区",
		},
		"130722": {
			Province: "河北省",
			City:     "张家口市",
			County:   "张北县",
		},
		"130723": {
			Province: "河北省",
			City:     "张家口市",
			County:   "康保县",
		},
		"130724": {
			Province: "河北省",
			City:     "张家口市",
			County:   "沽源县",
		},
		"130725": {
			Province: "河北省",
			City:     "张家口市",
			County:   "尚义县",
		},
		"130726": {
			Province: "河北省",
			City:     "张家口市",
			County:   "蔚县",
		},
		"130727": {
			Province: "河北省",
			City:     "张家口市",
			County:   "阳原县",
		},
		"130728": {
			Province: "河北省",
			City:     "张家口市",
			County:   "怀安县",
		},
		"130730": {
			Province: "河北省",
			City:     "张家口市",
			County:   "怀来县",
		},
		"130731": {
			Province: "河北省",
			City:     "张家口市",
			County:   "涿鹿县",
		},
		"130732": {
			Province: "河北省",
			City:     "张家口市",
			County:   "赤城县",
		},
		"130800": {
			Province: "河北省",
			City:     "承德市",
			County:   "承德市",
		},
		"130802": {
			Province: "河北省",
			City:     "承德市",
			County:   "双桥区",
		},
		"130803": {
			Province: "河北省",
			City:     "承德市",
			County:   "双滦区",
		},
		"130804": {
			Province: "河北省",
			City:     "承德市",
			County:   "鹰手营子矿区",
		},
		"130821": {
			Province: "河北省",
			City:     "承德市",
			County:   "承德县",
		},
		"130822": {
			Province: "河北省",
			City:     "承德市",
			County:   "兴隆县",
		},
		"130824": {
			Province: "河北省",
			City:     "承德市",
			County:   "滦平县",
		},
		"130825": {
			Province: "河北省",
			City:     "承德市",
			County:   "隆化县",
		},
		"130826": {
			Province: "河北省",
			City:     "承德市",
			County:   "丰宁满族自治县",
		},
		"130827": {
			Province: "河北省",
			City:     "承德市",
			County:   "宽城满族自治县",
		},
		"130828": {
			Province: "河北省",
			City:     "承德市",
			County:   "围场满族蒙古族自治县",
		},
		"130881": {
			Province: "河北省",
			City:     "承德市",
			County:   "平泉市",
		},
		"130900": {
			Province: "河北省",
			City:     "沧州市",
			County:   "沧州市",
		},
		"130902": {
			Province: "河北省",
			City:     "沧州市",
			County:   "新华区",
		},
		"130903": {
			Province: "河北省",
			City:     "沧州市",
			County:   "运河区",
		},
		"130921": {
			Province: "河北省",
			City:     "沧州市",
			County:   "沧县",
		},
		"130922": {
			Province: "河北省",
			City:     "沧州市",
			County:   "青县",
		},
		"130923": {
			Province: "河北省",
			City:     "沧州市",
			County:   "东光县",
		},
		"130924": {
			Province: "河北省",
			City:     "沧州市",
			County:   "海兴县",
		},
		"130925": {
			Province: "河北省",
			City:     "沧州市",
			County:   "盐山县",
		},
		"130926": {
			Province: "河北省",
			City:     "沧州市",
			County:   "肃宁县",
		},
		"130927": {
			Province: "河北省",
			City:     "沧州市",
			County:   "南皮县",
		},
		"130928": {
			Province: "河北省",
			City:     "沧州市",
			County:   "吴桥县",
		},
		"130929": {
			Province: "河北省",
			City:     "沧州市",
			County:   "献县",
		},
		"130930": {
			Province: "河北省",
			City:     "沧州市",
			County:   "孟村回族自治县",
		},
		"130981": {
			Province: "河北省",
			City:     "沧州市",
			County:   "泊头市",
		},
		"130982": {
			Province: "河北省",
			City:     "沧州市",
			County:   "任丘市",
		},
		"130983": {
			Province: "河北省",
			City:     "沧州市",
			County:   "黄骅市",
		},
		"130984": {
			Province: "河北省",
			City:     "沧州市",
			County:   "河间市",
		},
		"131000": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "廊坊市",
		},
		"131002": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "安次区",
		},
		"131003": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "广阳区",
		},
		"131022": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "固安县",
		},
		"131023": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "永清县",
		},
		"131024": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "香河县",
		},
		"131025": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "大城县",
		},
		"131026": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "文安县",
		},
		"131028": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "大厂回族自治县",
		},
		"131081": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "霸州市",
		},
		"131082": {
			Province: "河北省",
			City:     "廊坊市",
			County:   "三河市",
		},
		"131100": {
			Province: "河北省",
			City:     "衡水市",
			County:   "衡水市",
		},
		"131102": {
			Province: "河北省",
			City:     "衡水市",
			County:   "桃城区",
		},
		"131103": {
			Province: "河北省",
			City:     "衡水市",
			County:   "冀州区",
		},
		"131121": {
			Province: "河北省",
			City:     "衡水市",
			County:   "枣强县",
		},
		"131122": {
			Province: "河北省",
			City:     "衡水市",
			County:   "武邑县",
		},
		"131123": {
			Province: "河北省",
			City:     "衡水市",
			County:   "武强县",
		},
		"131124": {
			Province: "河北省",
			City:     "衡水市",
			County:   "饶阳县",
		},
		"131125": {
			Province: "河北省",
			City:     "衡水市",
			County:   "安平县",
		},
		"131126": {
			Province: "河北省",
			City:     "衡水市",
			County:   "故城县",
		},
		"131127": {
			Province: "河北省",
			City:     "衡水市",
			County:   "景县",
		},
		"131128": {
			Province: "河北省",
			City:     "衡水市",
			County:   "阜城县",
		},
		"131182": {
			Province: "河北省",
			City:     "衡水市",
			County:   "深州市",
		},
		"140000": {
			Province: "山西省",
			City:     "山西省",
			County:   "山西省",
		},
		"140100": {
			Province: "山西省",
			City:     "太原市",
			County:   "太原市",
		},
		"140105": {
			Province: "山西省",
			City:     "太原市",
			County:   "小店区",
		},
		"140106": {
			Province: "山西省",
			City:     "太原市",
			County:   "迎泽区",
		},
		"140107": {
			Province: "山西省",
			City:     "太原市",
			County:   "杏花岭区",
		},
		"140108": {
			Province: "山西省",
			City:     "太原市",
			County:   "尖草坪区",
		},
		"140109": {
			Province: "山西省",
			City:     "太原市",
			County:   "万柏林区",
		},
		"140110": {
			Province: "山西省",
			City:     "太原市",
			County:   "晋源区",
		},
		"140121": {
			Province: "山西省",
			City:     "太原市",
			County:   "清徐县",
		},
		"140122": {
			Province: "山西省",
			City:     "太原市",
			County:   "阳曲县",
		},
		"140123": {
			Province: "山西省",
			City:     "太原市",
			County:   "娄烦县",
		},
		"140181": {
			Province: "山西省",
			City:     "太原市",
			County:   "古交市",
		},
		"140200": {
			Province: "山西省",
			City:     "大同市",
			County:   "大同市",
		},
		"140212": {
			Province: "山西省",
			City:     "大同市",
			County:   "新荣区",
		},
		"140213": {
			Province: "山西省",
			City:     "大同市",
			County:   "平城区",
		},
		"140214": {
			Province: "山西省",
			City:     "大同市",
			County:   "云冈区",
		},
		"140215": {
			Province: "山西省",
			City:     "大同市",
			County:   "云州区",
		},
		"140221": {
			Province: "山西省",
			City:     "大同市",
			County:   "阳高县",
		},
		"140222": {
			Province: "山西省",
			City:     "大同市",
			County:   "天镇县",
		},
		"140223": {
			Province: "山西省",
			City:     "大同市",
			County:   "广灵县",
		},
		"140224": {
			Province: "山西省",
			City:     "大同市",
			County:   "灵丘县",
		},
		"140225": {
			Province: "山西省",
			City:     "大同市",
			County:   "浑源县",
		},
		"140226": {
			Province: "山西省",
			City:     "大同市",
			County:   "左云县",
		},
		"140300": {
			Province: "山西省",
			City:     "阳泉市",
			County:   "阳泉市",
		},
		"140302": {
			Province: "山西省",
			City:     "阳泉市",
			County:   "城区",
		},
		"140303": {
			Province: "山西省",
			City:     "阳泉市",
			County:   "矿区",
		},
		"140311": {
			Province: "山西省",
			City:     "阳泉市",
			County:   "郊区",
		},
		"140321": {
			Province: "山西省",
			City:     "阳泉市",
			County:   "平定县",
		},
		"140322": {
			Province: "山西省",
			City:     "阳泉市",
			County:   "盂县",
		},
		"140400": {
			Province: "山西省",
			City:     "长治市",
			County:   "长治市",
		},
		"140403": {
			Province: "山西省",
			City:     "长治市",
			County:   "潞州区",
		},
		"140404": {
			Province: "山西省",
			City:     "长治市",
			County:   "上党区",
		},
		"140405": {
			Province: "山西省",
			City:     "长治市",
			County:   "屯留区",
		},
		"140406": {
			Province: "山西省",
			City:     "长治市",
			County:   "潞城区",
		},
		"140423": {
			Province: "山西省",
			City:     "长治市",
			County:   "襄垣县",
		},
		"140425": {
			Province: "山西省",
			City:     "长治市",
			County:   "平顺县",
		},
		"140426": {
			Province: "山西省",
			City:     "长治市",
			County:   "黎城县",
		},
		"140427": {
			Province: "山西省",
			City:     "长治市",
			County:   "壶关县",
		},
		"140428": {
			Province: "山西省",
			City:     "长治市",
			County:   "长子县",
		},
		"140429": {
			Province: "山西省",
			City:     "长治市",
			County:   "武乡县",
		},
		"140430": {
			Province: "山西省",
			City:     "长治市",
			County:   "沁县",
		},
		"140431": {
			Province: "山西省",
			City:     "长治市",
			County:   "沁源县",
		},
		"140500": {
			Province: "山西省",
			City:     "晋城市",
			County:   "晋城市",
		},
		"140502": {
			Province: "山西省",
			City:     "晋城市",
			County:   "城区",
		},
		"140521": {
			Province: "山西省",
			City:     "晋城市",
			County:   "沁水县",
		},
		"140522": {
			Province: "山西省",
			City:     "晋城市",
			County:   "阳城县",
		},
		"140524": {
			Province: "山西省",
			City:     "晋城市",
			County:   "陵川县",
		},
		"140525": {
			Province: "山西省",
			City:     "晋城市",
			County:   "泽州县",
		},
		"140581": {
			Province: "山西省",
			City:     "晋城市",
			County:   "高平市",
		},
		"140600": {
			Province: "山西省",
			City:     "朔州市",
			County:   "朔州市",
		},
		"140602": {
			Province: "山西省",
			City:     "朔州市",
			County:   "朔城区",
		},
		"140603": {
			Province: "山西省",
			City:     "朔州市",
			County:   "平鲁区",
		},
		"140621": {
			Province: "山西省",
			City:     "朔州市",
			County:   "山阴县",
		},
		"140622": {
			Province: "山西省",
			City:     "朔州市",
			County:   "应县",
		},
		"140623": {
			Province: "山西省",
			City:     "朔州市",
			County:   "右玉县",
		},
		"140681": {
			Province: "山西省",
			City:     "朔州市",
			County:   "怀仁市",
		},
		"140700": {
			Province: "山西省",
			City:     "晋中市",
			County:   "晋中市",
		},
		"140702": {
			Province: "山西省",
			City:     "晋中市",
			County:   "榆次区",
		},
		"140703": {
			Province: "山西省",
			City:     "晋中市",
			County:   "太谷区",
		},
		"140721": {
			Province: "山西省",
			City:     "晋中市",
			County:   "榆社县",
		},
		"140722": {
			Province: "山西省",
			City:     "晋中市",
			County:   "左权县",
		},
		"140723": {
			Province: "山西省",
			City:     "晋中市",
			County:   "和顺县",
		},
		"140724": {
			Province: "山西省",
			City:     "晋中市",
			County:   "昔阳县",
		},
		"140725": {
			Province: "山西省",
			City:     "晋中市",
			County:   "寿阳县",
		},
		"140727": {
			Province: "山西省",
			City:     "晋中市",
			County:   "祁县",
		},
		"140728": {
			Province: "山西省",
			City:     "晋中市",
			County:   "平遥县",
		},
		"140729": {
			Province: "山西省",
			City:     "晋中市",
			County:   "灵石县",
		},
		"140781": {
			Province: "山西省",
			City:     "晋中市",
			County:   "介休市",
		},
		"140800": {
			Province: "山西省",
			City:     "运城市",
			County:   "运城市",
		},
		"140802": {
			Province: "山西省",
			City:     "运城市",
			County:   "盐湖区",
		},
		"140821": {
			Province: "山西省",
			City:     "运城市",
			County:   "临猗县",
		},
		"140822": {
			Province: "山西省",
			City:     "运城市",
			County:   "万荣县",
		},
		"140823": {
			Province: "山西省",
			City:     "运城市",
			County:   "闻喜县",
		},
		"140824": {
			Province: "山西省",
			City:     "运城市",
			County:   "稷山县",
		},
		"140825": {
			Province: "山西省",
			City:     "运城市",
			County:   "新绛县",
		},
		"140826": {
			Province: "山西省",
			City:     "运城市",
			County:   "绛县",
		},
		"140827": {
			Province: "山西省",
			City:     "运城市",
			County:   "垣曲县",
		},
		"140828": {
			Province: "山西省",
			City:     "运城市",
			County:   "夏县",
		},
		"140829": {
			Province: "山西省",
			City:     "运城市",
			County:   "平陆县",
		},
		"140830": {
			Province: "山西省",
			City:     "运城市",
			County:   "芮城县",
		},
		"140881": {
			Province: "山西省",
			City:     "运城市",
			County:   "永济市",
		},
		"140882": {
			Province: "山西省",
			City:     "运城市",
			County:   "河津市",
		},
		"140900": {
			Province: "山西省",
			City:     "忻州市",
			County:   "忻州市",
		},
		"140902": {
			Province: "山西省",
			City:     "忻州市",
			County:   "忻府区",
		},
		"140921": {
			Province: "山西省",
			City:     "忻州市",
			County:   "定襄县",
		},
		"140922": {
			Province: "山西省",
			City:     "忻州市",
			County:   "五台县",
		},
		"140923": {
			Province: "山西省",
			City:     "忻州市",
			County:   "代县",
		},
		"140924": {
			Province: "山西省",
			City:     "忻州市",
			County:   "繁峙县",
		},
		"140925": {
			Province: "山西省",
			City:     "忻州市",
			County:   "宁武县",
		},
		"140926": {
			Province: "山西省",
			City:     "忻州市",
			County:   "静乐县",
		},
		"140927": {
			Province: "山西省",
			City:     "忻州市",
			County:   "神池县",
		},
		"140928": {
			Province: "山西省",
			City:     "忻州市",
			County:   "五寨县",
		},
		"140929": {
			Province: "山西省",
			City:     "忻州市",
			County:   "岢岚县",
		},
		"140930": {
			Province: "山西省",
			City:     "忻州市",
			County:   "河曲县",
		},
		"140931": {
			Province: "山西省",
			City:     "忻州市",
			County:   "保德县",
		},
		"140932": {
			Province: "山西省",
			City:     "忻州市",
			County:   "偏关县",
		},
		"140981": {
			Province: "山西省",
			City:     "忻州市",
			County:   "原平市",
		},
		"141000": {
			Province: "山西省",
			City:     "临汾市",
			County:   "临汾市",
		},
		"141002": {
			Province: "山西省",
			City:     "临汾市",
			County:   "尧都区",
		},
		"141021": {
			Province: "山西省",
			City:     "临汾市",
			County:   "曲沃县",
		},
		"141022": {
			Province: "山西省",
			City:     "临汾市",
			County:   "翼城县",
		},
		"141023": {
			Province: "山西省",
			City:     "临汾市",
			County:   "襄汾县",
		},
		"141024": {
			Province: "山西省",
			City:     "临汾市",
			County:   "洪洞县",
		},
		"141025": {
			Province: "山西省",
			City:     "临汾市",
			County:   "古县",
		},
		"141026": {
			Province: "山西省",
			City:     "临汾市",
			County:   "安泽县",
		},
		"141027": {
			Province: "山西省",
			City:     "临汾市",
			County:   "浮山县",
		},
		"141028": {
			Province: "山西省",
			City:     "临汾市",
			County:   "吉县",
		},
		"141029": {
			Province: "山西省",
			City:     "临汾市",
			County:   "乡宁县",
		},
		"141030": {
			Province: "山西省",
			City:     "临汾市",
			County:   "大宁县",
		},
		"141031": {
			Province: "山西省",
			City:     "临汾市",
			County:   "隰县",
		},
		"141032": {
			Province: "山西省",
			City:     "临汾市",
			County:   "永和县",
		},
		"141033": {
			Province: "山西省",
			City:     "临汾市",
			County:   "蒲县",
		},
		"141034": {
			Province: "山西省",
			City:     "临汾市",
			County:   "汾西县",
		},
		"141081": {
			Province: "山西省",
			City:     "临汾市",
			County:   "侯马市",
		},
		"141082": {
			Province: "山西省",
			City:     "临汾市",
			County:   "霍州市",
		},
		"141100": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "吕梁市",
		},
		"141102": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "离石区",
		},
		"141121": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "文水县",
		},
		"141122": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "交城县",
		},
		"141123": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "兴县",
		},
		"141124": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "临县",
		},
		"141125": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "柳林县",
		},
		"141126": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "石楼县",
		},
		"141127": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "岚县",
		},
		"141128": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "方山县",
		},
		"141129": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "中阳县",
		},
		"141130": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "交口县",
		},
		"141181": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "孝义市",
		},
		"141182": {
			Province: "山西省",
			City:     "吕梁市",
			County:   "汾阳市",
		},
		"150000": {
			Province: "内蒙古自治区",
			City:     "内蒙古自治区",
			County:   "内蒙古自治区",
		},
		"150100": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "呼和浩特市",
		},
		"150102": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "新城区",
		},
		"150103": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "回民区",
		},
		"150104": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "玉泉区",
		},
		"150105": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "赛罕区",
		},
		"150121": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "土默特左旗",
		},
		"150122": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "托克托县",
		},
		"150123": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "和林格尔县",
		},
		"150124": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "清水河县",
		},
		"150125": {
			Province: "内蒙古自治区",
			City:     "呼和浩特市",
			County:   "武川县",
		},
		"150200": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "包头市",
		},
		"150202": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "东河区",
		},
		"150203": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "昆都仑区",
		},
		"150204": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "青山区",
		},
		"150205": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "石拐区",
		},
		"150206": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "白云鄂博矿区",
		},
		"150207": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "九原区",
		},
		"150221": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "土默特右旗",
		},
		"150222": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "固阳县",
		},
		"150223": {
			Province: "内蒙古自治区",
			City:     "包头市",
			County:   "达尔罕茂明安联合旗",
		},
		"150300": {
			Province: "内蒙古自治区",
			City:     "乌海市",
			County:   "乌海市",
		},
		"150302": {
			Province: "内蒙古自治区",
			City:     "乌海市",
			County:   "海勃湾区",
		},
		"150303": {
			Province: "内蒙古自治区",
			City:     "乌海市",
			County:   "海南区",
		},
		"150304": {
			Province: "内蒙古自治区",
			City:     "乌海市",
			County:   "乌达区",
		},
		"150400": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "赤峰市",
		},
		"150402": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "红山区",
		},
		"150403": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "元宝山区",
		},
		"150404": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "松山区",
		},
		"150421": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "阿鲁科尔沁旗",
		},
		"150422": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "巴林左旗",
		},
		"150423": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "巴林右旗",
		},
		"150424": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "林西县",
		},
		"150425": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "克什克腾旗",
		},
		"150426": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "翁牛特旗",
		},
		"150428": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "喀喇沁旗",
		},
		"150429": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "宁城县",
		},
		"150430": {
			Province: "内蒙古自治区",
			City:     "赤峰市",
			County:   "敖汉旗",
		},
		"150500": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "通辽市",
		},
		"150502": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "科尔沁区",
		},
		"150521": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "科尔沁左翼中旗",
		},
		"150522": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "科尔沁左翼后旗",
		},
		"150523": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "开鲁县",
		},
		"150524": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "库伦旗",
		},
		"150525": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "奈曼旗",
		},
		"150526": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "扎鲁特旗",
		},
		"150581": {
			Province: "内蒙古自治区",
			City:     "通辽市",
			County:   "霍林郭勒市",
		},
		"150600": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "鄂尔多斯市",
		},
		"150602": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "东胜区",
		},
		"150603": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "康巴什区",
		},
		"150621": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "达拉特旗",
		},
		"150622": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "准格尔旗",
		},
		"150623": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "鄂托克前旗",
		},
		"150624": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "鄂托克旗",
		},
		"150625": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "杭锦旗",
		},
		"150626": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "乌审旗",
		},
		"150627": {
			Province: "内蒙古自治区",
			City:     "鄂尔多斯市",
			County:   "伊金霍洛旗",
		},
		"150700": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "呼伦贝尔市",
		},
		"150702": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "海拉尔区",
		},
		"150703": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "扎赉诺尔区",
		},
		"150721": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "阿荣旗",
		},
		"150722": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "莫力达瓦达斡尔族自治旗",
		},
		"150723": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "鄂伦春自治旗",
		},
		"150724": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "鄂温克族自治旗",
		},
		"150725": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "陈巴尔虎旗",
		},
		"150726": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "新巴尔虎左旗",
		},
		"150727": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "新巴尔虎右旗",
		},
		"150781": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "满洲里市",
		},
		"150782": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "牙克石市",
		},
		"150783": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "扎兰屯市",
		},
		"150784": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "额尔古纳市",
		},
		"150785": {
			Province: "内蒙古自治区",
			City:     "呼伦贝尔市",
			County:   "根河市",
		},
		"150800": {
			Province: "内蒙古自治区",
			City:     "巴彦淖尔市",
			County:   "巴彦淖尔市",
		},
		"150802": {
			Province: "内蒙古自治区",
			City:     "巴彦淖尔市",
			County:   "临河区",
		},
		"150821": {
			Province: "内蒙古自治区",
			City:     "巴彦淖尔市",
			County:   "五原县",
		},
		"150822": {
			Province: "内蒙古自治区",
			City:     "巴彦淖尔市",
			County:   "磴口县",
		},
		"150823": {
			Province: "内蒙古自治区",
			City:     "巴彦淖尔市",
			County:   "乌拉特前旗",
		},
		"150824": {
			Province: "内蒙古自治区",
			City:     "巴彦淖尔市",
			County:   "乌拉特中旗",
		},
		"150825": {
			Province: "内蒙古自治区",
			City:     "巴彦淖尔市",
			County:   "乌拉特后旗",
		},
		"150826": {
			Province: "内蒙古自治区",
			City:     "巴彦淖尔市",
			County:   "杭锦后旗",
		},
		"150900": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "乌兰察布市",
		},
		"150902": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "集宁区",
		},
		"150921": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "卓资县",
		},
		"150922": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "化德县",
		},
		"150923": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "商都县",
		},
		"150924": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "兴和县",
		},
		"150925": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "凉城县",
		},
		"150926": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "察哈尔右翼前旗",
		},
		"150927": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "察哈尔右翼中旗",
		},
		"150928": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "察哈尔右翼后旗",
		},
		"150929": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "四子王旗",
		},
		"150981": {
			Province: "内蒙古自治区",
			City:     "乌兰察布市",
			County:   "丰镇市",
		},
		"152200": {
			Province: "内蒙古自治区",
			City:     "兴安盟",
			County:   "兴安盟",
		},
		"152201": {
			Province: "内蒙古自治区",
			City:     "兴安盟",
			County:   "乌兰浩特市",
		},
		"152202": {
			Province: "内蒙古自治区",
			City:     "兴安盟",
			County:   "阿尔山市",
		},
		"152221": {
			Province: "内蒙古自治区",
			City:     "兴安盟",
			County:   "科尔沁右翼前旗",
		},
		"152222": {
			Province: "内蒙古自治区",
			City:     "兴安盟",
			County:   "科尔沁右翼中旗",
		},
		"152223": {
			Province: "内蒙古自治区",
			City:     "兴安盟",
			County:   "扎赉特旗",
		},
		"152224": {
			Province: "内蒙古自治区",
			City:     "兴安盟",
			County:   "突泉县",
		},
		"152500": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "锡林郭勒盟",
		},
		"152501": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "二连浩特市",
		},
		"152502": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "锡林浩特市",
		},
		"152522": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "阿巴嘎旗",
		},
		"152523": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "苏尼特左旗",
		},
		"152524": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "苏尼特右旗",
		},
		"152525": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "东乌珠穆沁旗",
		},
		"152526": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "西乌珠穆沁旗",
		},
		"152527": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "太仆寺旗",
		},
		"152528": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "镶黄旗",
		},
		"152529": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "正镶白旗",
		},
		"152530": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "正蓝旗",
		},
		"152531": {
			Province: "内蒙古自治区",
			City:     "锡林郭勒盟",
			County:   "多伦县",
		},
		"152900": {
			Province: "内蒙古自治区",
			City:     "阿拉善盟",
			County:   "阿拉善盟",
		},
		"152921": {
			Province: "内蒙古自治区",
			City:     "阿拉善盟",
			County:   "阿拉善左旗",
		},
		"152922": {
			Province: "内蒙古自治区",
			City:     "阿拉善盟",
			County:   "阿拉善右旗",
		},
		"152923": {
			Province: "内蒙古自治区",
			City:     "阿拉善盟",
			County:   "额济纳旗",
		},
		"210000": {
			Province: "辽宁省",
			City:     "辽宁省",
			County:   "辽宁省",
		},
		"210100": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "沈阳市",
		},
		"210102": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "和平区",
		},
		"210103": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "沈河区",
		},
		"210104": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "大东区",
		},
		"210105": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "皇姑区",
		},
		"210106": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "铁西区",
		},
		"210111": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "苏家屯区",
		},
		"210112": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "浑南区",
		},
		"210113": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "沈北新区",
		},
		"210114": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "于洪区",
		},
		"210115": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "辽中区",
		},
		"210123": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "康平县",
		},
		"210124": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "法库县",
		},
		"210181": {
			Province: "辽宁省",
			City:     "沈阳市",
			County:   "新民市",
		},
		"210200": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "大连市",
		},
		"210202": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "中山区",
		},
		"210203": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "西岗区",
		},
		"210204": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "沙河口区",
		},
		"210211": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "甘井子区",
		},
		"210212": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "旅顺口区",
		},
		"210213": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "金州区",
		},
		"210214": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "普兰店区",
		},
		"210224": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "长海县",
		},
		"210281": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "瓦房店市",
		},
		"210283": {
			Province: "辽宁省",
			City:     "大连市",
			County:   "庄河市",
		},
		"210300": {
			Province: "辽宁省",
			City:     "鞍山市",
			County:   "鞍山市",
		},
		"210302": {
			Province: "辽宁省",
			City:     "鞍山市",
			County:   "铁东区",
		},
		"210303": {
			Province: "辽宁省",
			City:     "鞍山市",
			County:   "铁西区",
		},
		"210304": {
			Province: "辽宁省",
			City:     "鞍山市",
			County:   "立山区",
		},
		"210311": {
			Province: "辽宁省",
			City:     "鞍山市",
			County:   "千山区",
		},
		"210321": {
			Province: "辽宁省",
			City:     "鞍山市",
			County:   "台安县",
		},
		"210323": {
			Province: "辽宁省",
			City:     "鞍山市",
			County:   "岫岩满族自治县",
		},
		"210381": {
			Province: "辽宁省",
			City:     "鞍山市",
			County:   "海城市",
		},
		"210400": {
			Province: "辽宁省",
			City:     "抚顺市",
			County:   "抚顺市",
		},
		"210402": {
			Province: "辽宁省",
			City:     "抚顺市",
			County:   "新抚区",
		},
		"210403": {
			Province: "辽宁省",
			City:     "抚顺市",
			County:   "东洲区",
		},
		"210404": {
			Province: "辽宁省",
			City:     "抚顺市",
			County:   "望花区",
		},
		"210411": {
			Province: "辽宁省",
			City:     "抚顺市",
			County:   "顺城区",
		},
		"210421": {
			Province: "辽宁省",
			City:     "抚顺市",
			County:   "抚顺县",
		},
		"210422": {
			Province: "辽宁省",
			City:     "抚顺市",
			County:   "新宾满族自治县",
		},
		"210423": {
			Province: "辽宁省",
			City:     "抚顺市",
			County:   "清原满族自治县",
		},
		"210500": {
			Province: "辽宁省",
			City:     "本溪市",
			County:   "本溪市",
		},
		"210502": {
			Province: "辽宁省",
			City:     "本溪市",
			County:   "平山区",
		},
		"210503": {
			Province: "辽宁省",
			City:     "本溪市",
			County:   "溪湖区",
		},
		"210504": {
			Province: "辽宁省",
			City:     "本溪市",
			County:   "明山区",
		},
		"210505": {
			Province: "辽宁省",
			City:     "本溪市",
			County:   "南芬区",
		},
		"210521": {
			Province: "辽宁省",
			City:     "本溪市",
			County:   "本溪满族自治县",
		},
		"210522": {
			Province: "辽宁省",
			City:     "本溪市",
			County:   "桓仁满族自治县",
		},
		"210600": {
			Province: "辽宁省",
			City:     "丹东市",
			County:   "丹东市",
		},
		"210602": {
			Province: "辽宁省",
			City:     "丹东市",
			County:   "元宝区",
		},
		"210603": {
			Province: "辽宁省",
			City:     "丹东市",
			County:   "振兴区",
		},
		"210604": {
			Province: "辽宁省",
			City:     "丹东市",
			County:   "振安区",
		},
		"210624": {
			Province: "辽宁省",
			City:     "丹东市",
			County:   "宽甸满族自治县",
		},
		"210681": {
			Province: "辽宁省",
			City:     "丹东市",
			County:   "东港市",
		},
		"210682": {
			Province: "辽宁省",
			City:     "丹东市",
			County:   "凤城市",
		},
		"210700": {
			Province: "辽宁省",
			City:     "锦州市",
			County:   "锦州市",
		},
		"210702": {
			Province: "辽宁省",
			City:     "锦州市",
			County:   "古塔区",
		},
		"210703": {
			Province: "辽宁省",
			City:     "锦州市",
			County:   "凌河区",
		},
		"210711": {
			Province: "辽宁省",
			City:     "锦州市",
			County:   "太和区",
		},
		"210726": {
			Province: "辽宁省",
			City:     "锦州市",
			County:   "黑山县",
		},
		"210727": {
			Province: "辽宁省",
			City:     "锦州市",
			County:   "义县",
		},
		"210781": {
			Province: "辽宁省",
			City:     "锦州市",
			County:   "凌海市",
		},
		"210782": {
			Province: "辽宁省",
			City:     "锦州市",
			County:   "北镇市",
		},
		"210800": {
			Province: "辽宁省",
			City:     "营口市",
			County:   "营口市",
		},
		"210802": {
			Province: "辽宁省",
			City:     "营口市",
			County:   "站前区",
		},
		"210803": {
			Province: "辽宁省",
			City:     "营口市",
			County:   "西市区",
		},
		"210804": {
			Province: "辽宁省",
			City:     "营口市",
			County:   "鲅鱼圈区",
		},
		"210811": {
			Province: "辽宁省",
			City:     "营口市",
			County:   "老边区",
		},
		"210881": {
			Province: "辽宁省",
			City:     "营口市",
			County:   "盖州市",
		},
		"210882": {
			Province: "辽宁省",
			City:     "营口市",
			County:   "大石桥市",
		},
		"210900": {
			Province: "辽宁省",
			City:     "阜新市",
			County:   "阜新市",
		},
		"210902": {
			Province: "辽宁省",
			City:     "阜新市",
			County:   "海州区",
		},
		"210903": {
			Province: "辽宁省",
			City:     "阜新市",
			County:   "新邱区",
		},
		"210904": {
			Province: "辽宁省",
			City:     "阜新市",
			County:   "太平区",
		},
		"210905": {
			Province: "辽宁省",
			City:     "阜新市",
			County:   "清河门区",
		},
		"210911": {
			Province: "辽宁省",
			City:     "阜新市",
			County:   "细河区",
		},
		"210921": {
			Province: "辽宁省",
			City:     "阜新市",
			County:   "阜新蒙古族自治县",
		},
		"210922": {
			Province: "辽宁省",
			City:     "阜新市",
			County:   "彰武县",
		},
		"211000": {
			Province: "辽宁省",
			City:     "辽阳市",
			County:   "辽阳市",
		},
		"211002": {
			Province: "辽宁省",
			City:     "辽阳市",
			County:   "白塔区",
		},
		"211003": {
			Province: "辽宁省",
			City:     "辽阳市",
			County:   "文圣区",
		},
		"211004": {
			Province: "辽宁省",
			City:     "辽阳市",
			County:   "宏伟区",
		},
		"211005": {
			Province: "辽宁省",
			City:     "辽阳市",
			County:   "弓长岭区",
		},
		"211011": {
			Province: "辽宁省",
			City:     "辽阳市",
			County:   "太子河区",
		},
		"211021": {
			Province: "辽宁省",
			City:     "辽阳市",
			County:   "辽阳县",
		},
		"211081": {
			Province: "辽宁省",
			City:     "辽阳市",
			County:   "灯塔市",
		},
		"211100": {
			Province: "辽宁省",
			City:     "盘锦市",
			County:   "盘锦市",
		},
		"211102": {
			Province: "辽宁省",
			City:     "盘锦市",
			County:   "双台子区",
		},
		"211103": {
			Province: "辽宁省",
			City:     "盘锦市",
			County:   "兴隆台区",
		},
		"211104": {
			Province: "辽宁省",
			City:     "盘锦市",
			County:   "大洼区",
		},
		"211122": {
			Province: "辽宁省",
			City:     "盘锦市",
			County:   "盘山县",
		},
		"211200": {
			Province: "辽宁省",
			City:     "铁岭市",
			County:   "铁岭市",
		},
		"211202": {
			Province: "辽宁省",
			City:     "铁岭市",
			County:   "银州区",
		},
		"211204": {
			Province: "辽宁省",
			City:     "铁岭市",
			County:   "清河区",
		},
		"211221": {
			Province: "辽宁省",
			City:     "铁岭市",
			County:   "铁岭县",
		},
		"211223": {
			Province: "辽宁省",
			City:     "铁岭市",
			County:   "西丰县",
		},
		"211224": {
			Province: "辽宁省",
			City:     "铁岭市",
			County:   "昌图县",
		},
		"211281": {
			Province: "辽宁省",
			City:     "铁岭市",
			County:   "调兵山市",
		},
		"211282": {
			Province: "辽宁省",
			City:     "铁岭市",
			County:   "开原市",
		},
		"211300": {
			Province: "辽宁省",
			City:     "朝阳市",
			County:   "朝阳市",
		},
		"211302": {
			Province: "辽宁省",
			City:     "朝阳市",
			County:   "双塔区",
		},
		"211303": {
			Province: "辽宁省",
			City:     "朝阳市",
			County:   "龙城区",
		},
		"211321": {
			Province: "辽宁省",
			City:     "朝阳市",
			County:   "朝阳县",
		},
		"211322": {
			Province: "辽宁省",
			City:     "朝阳市",
			County:   "建平县",
		},
		"211324": {
			Province: "辽宁省",
			City:     "朝阳市",
			County:   "喀喇沁左翼蒙古族自治县",
		},
		"211381": {
			Province: "辽宁省",
			City:     "朝阳市",
			County:   "北票市",
		},
		"211382": {
			Province: "辽宁省",
			City:     "朝阳市",
			County:   "凌源市",
		},
		"211400": {
			Province: "辽宁省",
			City:     "葫芦岛市",
			County:   "葫芦岛市",
		},
		"211402": {
			Province: "辽宁省",
			City:     "葫芦岛市",
			County:   "连山区",
		},
		"211403": {
			Province: "辽宁省",
			City:     "葫芦岛市",
			County:   "龙港区",
		},
		"211404": {
			Province: "辽宁省",
			City:     "葫芦岛市",
			County:   "南票区",
		},
		"211421": {
			Province: "辽宁省",
			City:     "葫芦岛市",
			County:   "绥中县",
		},
		"211422": {
			Province: "辽宁省",
			City:     "葫芦岛市",
			County:   "建昌县",
		},
		"211481": {
			Province: "辽宁省",
			City:     "葫芦岛市",
			County:   "兴城市",
		},
		"220000": {
			Province: "吉林省",
			City:     "吉林省",
			County:   "吉林省",
		},
		"220100": {
			Province: "吉林省",
			City:     "长春市",
			County:   "长春市",
		},
		"220102": {
			Province: "吉林省",
			City:     "长春市",
			County:   "南关区",
		},
		"220103": {
			Province: "吉林省",
			City:     "长春市",
			County:   "宽城区",
		},
		"220104": {
			Province: "吉林省",
			City:     "长春市",
			County:   "朝阳区",
		},
		"220105": {
			Province: "吉林省",
			City:     "长春市",
			County:   "二道区",
		},
		"220106": {
			Province: "吉林省",
			City:     "长春市",
			County:   "绿园区",
		},
		"220112": {
			Province: "吉林省",
			City:     "长春市",
			County:   "双阳区",
		},
		"220113": {
			Province: "吉林省",
			City:     "长春市",
			County:   "九台区",
		},
		"220122": {
			Province: "吉林省",
			City:     "长春市",
			County:   "农安县",
		},
		"220182": {
			Province: "吉林省",
			City:     "长春市",
			County:   "榆树市",
		},
		"220183": {
			Province: "吉林省",
			City:     "长春市",
			County:   "德惠市",
		},
		"220184": {
			Province: "吉林省",
			City:     "长春市",
			County:   "公主岭市",
		},
		"220200": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "吉林市",
		},
		"220202": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "昌邑区",
		},
		"220203": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "龙潭区",
		},
		"220204": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "船营区",
		},
		"220211": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "丰满区",
		},
		"220221": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "永吉县",
		},
		"220281": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "蛟河市",
		},
		"220282": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "桦甸市",
		},
		"220283": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "舒兰市",
		},
		"220284": {
			Province: "吉林省",
			City:     "吉林市",
			County:   "磐石市",
		},
		"220300": {
			Province: "吉林省",
			City:     "四平市",
			County:   "四平市",
		},
		"220302": {
			Province: "吉林省",
			City:     "四平市",
			County:   "铁西区",
		},
		"220303": {
			Province: "吉林省",
			City:     "四平市",
			County:   "铁东区",
		},
		"220322": {
			Province: "吉林省",
			City:     "四平市",
			County:   "梨树县",
		},
		"220323": {
			Province: "吉林省",
			City:     "四平市",
			County:   "伊通满族自治县",
		},
		"220382": {
			Province: "吉林省",
			City:     "四平市",
			County:   "双辽市",
		},
		"220400": {
			Province: "吉林省",
			City:     "辽源市",
			County:   "辽源市",
		},
		"220402": {
			Province: "吉林省",
			City:     "辽源市",
			County:   "龙山区",
		},
		"220403": {
			Province: "吉林省",
			City:     "辽源市",
			County:   "西安区",
		},
		"220421": {
			Province: "吉林省",
			City:     "辽源市",
			County:   "东丰县",
		},
		"220422": {
			Province: "吉林省",
			City:     "辽源市",
			County:   "东辽县",
		},
		"220500": {
			Province: "吉林省",
			City:     "通化市",
			County:   "通化市",
		},
		"220502": {
			Province: "吉林省",
			City:     "通化市",
			County:   "东昌区",
		},
		"220503": {
			Province: "吉林省",
			City:     "通化市",
			County:   "二道江区",
		},
		"220521": {
			Province: "吉林省",
			City:     "通化市",
			County:   "通化县",
		},
		"220523": {
			Province: "吉林省",
			City:     "通化市",
			County:   "辉南县",
		},
		"220524": {
			Province: "吉林省",
			City:     "通化市",
			County:   "柳河县",
		},
		"220581": {
			Province: "吉林省",
			City:     "通化市",
			County:   "梅河口市",
		},
		"220582": {
			Province: "吉林省",
			City:     "通化市",
			County:   "集安市",
		},
		"220600": {
			Province: "吉林省",
			City:     "白山市",
			County:   "白山市",
		},
		"220602": {
			Province: "吉林省",
			City:     "白山市",
			County:   "浑江区",
		},
		"220605": {
			Province: "吉林省",
			City:     "白山市",
			County:   "江源区",
		},
		"220621": {
			Province: "吉林省",
			City:     "白山市",
			County:   "抚松县",
		},
		"220622": {
			Province: "吉林省",
			City:     "白山市",
			County:   "靖宇县",
		},
		"220623": {
			Province: "吉林省",
			City:     "白山市",
			County:   "长白朝鲜族自治县",
		},
		"220681": {
			Province: "吉林省",
			City:     "白山市",
			County:   "临江市",
		},
		"220700": {
			Province: "吉林省",
			City:     "松原市",
			County:   "松原市",
		},
		"220702": {
			Province: "吉林省",
			City:     "松原市",
			County:   "宁江区",
		},
		"220721": {
			Province: "吉林省",
			City:     "松原市",
			County:   "前郭尔罗斯蒙古族自治县",
		},
		"220722": {
			Province: "吉林省",
			City:     "松原市",
			County:   "长岭县",
		},
		"220723": {
			Province: "吉林省",
			City:     "松原市",
			County:   "乾安县",
		},
		"220781": {
			Province: "吉林省",
			City:     "松原市",
			County:   "扶余市",
		},
		"220800": {
			Province: "吉林省",
			City:     "白城市",
			County:   "白城市",
		},
		"220802": {
			Province: "吉林省",
			City:     "白城市",
			County:   "洮北区",
		},
		"220821": {
			Province: "吉林省",
			City:     "白城市",
			County:   "镇赉县",
		},
		"220822": {
			Province: "吉林省",
			City:     "白城市",
			County:   "通榆县",
		},
		"220881": {
			Province: "吉林省",
			City:     "白城市",
			County:   "洮南市",
		},
		"220882": {
			Province: "吉林省",
			City:     "白城市",
			County:   "大安市",
		},
		"222400": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "延边朝鲜族自治州",
		},
		"222401": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "延吉市",
		},
		"222402": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "图们市",
		},
		"222403": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "敦化市",
		},
		"222404": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "珲春市",
		},
		"222405": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "龙井市",
		},
		"222406": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "和龙市",
		},
		"222424": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "汪清县",
		},
		"222426": {
			Province: "吉林省",
			City:     "延边朝鲜族自治州",
			County:   "安图县",
		},
		"230000": {
			Province: "黑龙江省",
			City:     "黑龙江省",
			County:   "黑龙江省",
		},
		"230100": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "哈尔滨市",
		},
		"230102": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "道里区",
		},
		"230103": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "南岗区",
		},
		"230104": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "道外区",
		},
		"230108": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "平房区",
		},
		"230109": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "松北区",
		},
		"230110": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "香坊区",
		},
		"230111": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "呼兰区",
		},
		"230112": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "阿城区",
		},
		"230113": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "双城区",
		},
		"230123": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "依兰县",
		},
		"230124": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "方正县",
		},
		"230125": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "宾县",
		},
		"230126": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "巴彦县",
		},
		"230127": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "木兰县",
		},
		"230128": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "通河县",
		},
		"230129": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "延寿县",
		},
		"230183": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "尚志市",
		},
		"230184": {
			Province: "黑龙江省",
			City:     "哈尔滨市",
			County:   "五常市",
		},
		"230200": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "齐齐哈尔市",
		},
		"230202": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "龙沙区",
		},
		"230203": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "建华区",
		},
		"230204": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "铁锋区",
		},
		"230205": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "昂昂溪区",
		},
		"230206": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "富拉尔基区",
		},
		"230207": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "碾子山区",
		},
		"230208": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "梅里斯达斡尔族区",
		},
		"230221": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "龙江县",
		},
		"230223": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "依安县",
		},
		"230224": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "泰来县",
		},
		"230225": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "甘南县",
		},
		"230227": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "富裕县",
		},
		"230229": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "克山县",
		},
		"230230": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "克东县",
		},
		"230231": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "拜泉县",
		},
		"230281": {
			Province: "黑龙江省",
			City:     "齐齐哈尔市",
			County:   "讷河市",
		},
		"230300": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "鸡西市",
		},
		"230302": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "鸡冠区",
		},
		"230303": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "恒山区",
		},
		"230304": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "滴道区",
		},
		"230305": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "梨树区",
		},
		"230306": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "城子河区",
		},
		"230307": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "麻山区",
		},
		"230321": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "鸡东县",
		},
		"230381": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "虎林市",
		},
		"230382": {
			Province: "黑龙江省",
			City:     "鸡西市",
			County:   "密山市",
		},
		"230400": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "鹤岗市",
		},
		"230402": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "向阳区",
		},
		"230403": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "工农区",
		},
		"230404": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "南山区",
		},
		"230405": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "兴安区",
		},
		"230406": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "东山区",
		},
		"230407": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "兴山区",
		},
		"230421": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "萝北县",
		},
		"230422": {
			Province: "黑龙江省",
			City:     "鹤岗市",
			County:   "绥滨县",
		},
		"230500": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "双鸭山市",
		},
		"230502": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "尖山区",
		},
		"230503": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "岭东区",
		},
		"230505": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "四方台区",
		},
		"230506": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "宝山区",
		},
		"230521": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "集贤县",
		},
		"230522": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "友谊县",
		},
		"230523": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "宝清县",
		},
		"230524": {
			Province: "黑龙江省",
			City:     "双鸭山市",
			County:   "饶河县",
		},
		"230600": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "大庆市",
		},
		"230602": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "萨尔图区",
		},
		"230603": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "龙凤区",
		},
		"230604": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "让胡路区",
		},
		"230605": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "红岗区",
		},
		"230606": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "大同区",
		},
		"230621": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "肇州县",
		},
		"230622": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "肇源县",
		},
		"230623": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "林甸县",
		},
		"230624": {
			Province: "黑龙江省",
			City:     "大庆市",
			County:   "杜尔伯特蒙古族自治县",
		},
		"230700": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "伊春市",
		},
		"230717": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "伊美区",
		},
		"230718": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "乌翠区",
		},
		"230719": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "友好区",
		},
		"230722": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "嘉荫县",
		},
		"230723": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "汤旺县",
		},
		"230724": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "丰林县",
		},
		"230725": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "大箐山县",
		},
		"230726": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "南岔县",
		},
		"230751": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "金林区",
		},
		"230781": {
			Province: "黑龙江省",
			City:     "伊春市",
			County:   "铁力市",
		},
		"230800": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "佳木斯市",
		},
		"230803": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "向阳区",
		},
		"230804": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "前进区",
		},
		"230805": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "东风区",
		},
		"230811": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "郊区",
		},
		"230822": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "桦南县",
		},
		"230826": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "桦川县",
		},
		"230828": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "汤原县",
		},
		"230881": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "同江市",
		},
		"230882": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "富锦市",
		},
		"230883": {
			Province: "黑龙江省",
			City:     "佳木斯市",
			County:   "抚远市",
		},
		"230900": {
			Province: "黑龙江省",
			City:     "七台河市",
			County:   "七台河市",
		},
		"230902": {
			Province: "黑龙江省",
			City:     "七台河市",
			County:   "新兴区",
		},
		"230903": {
			Province: "黑龙江省",
			City:     "七台河市",
			County:   "桃山区",
		},
		"230904": {
			Province: "黑龙江省",
			City:     "七台河市",
			County:   "茄子河区",
		},
		"230921": {
			Province: "黑龙江省",
			City:     "七台河市",
			County:   "勃利县",
		},
		"231000": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "牡丹江市",
		},
		"231002": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "东安区",
		},
		"231003": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "阳明区",
		},
		"231004": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "爱民区",
		},
		"231005": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "西安区",
		},
		"231025": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "林口县",
		},
		"231081": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "绥芬河市",
		},
		"231083": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "海林市",
		},
		"231084": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "宁安市",
		},
		"231085": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "穆棱市",
		},
		"231086": {
			Province: "黑龙江省",
			City:     "牡丹江市",
			County:   "东宁市",
		},
		"231100": {
			Province: "黑龙江省",
			City:     "黑河市",
			County:   "黑河市",
		},
		"231102": {
			Province: "黑龙江省",
			City:     "黑河市",
			County:   "爱辉区",
		},
		"231123": {
			Province: "黑龙江省",
			City:     "黑河市",
			County:   "逊克县",
		},
		"231124": {
			Province: "黑龙江省",
			City:     "黑河市",
			County:   "孙吴县",
		},
		"231181": {
			Province: "黑龙江省",
			City:     "黑河市",
			County:   "北安市",
		},
		"231182": {
			Province: "黑龙江省",
			City:     "黑河市",
			County:   "五大连池市",
		},
		"231183": {
			Province: "黑龙江省",
			City:     "黑河市",
			County:   "嫩江市",
		},
		"231200": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "绥化市",
		},
		"231202": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "北林区",
		},
		"231221": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "望奎县",
		},
		"231222": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "兰西县",
		},
		"231223": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "青冈县",
		},
		"231224": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "庆安县",
		},
		"231225": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "明水县",
		},
		"231226": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "绥棱县",
		},
		"231281": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "安达市",
		},
		"231282": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "肇东市",
		},
		"231283": {
			Province: "黑龙江省",
			City:     "绥化市",
			County:   "海伦市",
		},
		"232700": {
			Province: "黑龙江省",
			City:     "大兴安岭地区",
			County:   "大兴安岭地区",
		},
		"232701": {
			Province: "黑龙江省",
			City:     "大兴安岭地区",
			County:   "漠河市",
		},
		"232721": {
			Province: "黑龙江省",
			City:     "大兴安岭地区",
			County:   "呼玛县",
		},
		"232722": {
			Province: "黑龙江省",
			City:     "大兴安岭地区",
			County:   "塔河县",
		},
		"310000": {
			Province: "上海市",
			City:     "上海市",
			County:   "上海市",
		},
		"310101": {
			Province: "上海市",
			City:     "上海市",
			County:   "黄浦区",
		},
		"310104": {
			Province: "上海市",
			City:     "上海市",
			County:   "徐汇区",
		},
		"310105": {
			Province: "上海市",
			City:     "上海市",
			County:   "长宁区",
		},
		"310106": {
			Province: "上海市",
			City:     "上海市",
			County:   "静安区",
		},
		"310107": {
			Province: "上海市",
			City:     "上海市",
			County:   "普陀区",
		},
		"310109": {
			Province: "上海市",
			City:     "上海市",
			County:   "虹口区",
		},
		"310110": {
			Province: "上海市",
			City:     "上海市",
			County:   "杨浦区",
		},
		"310112": {
			Province: "上海市",
			City:     "上海市",
			County:   "闵行区",
		},
		"310113": {
			Province: "上海市",
			City:     "上海市",
			County:   "宝山区",
		},
		"310114": {
			Province: "上海市",
			City:     "上海市",
			County:   "嘉定区",
		},
		"310115": {
			Province: "上海市",
			City:     "上海市",
			County:   "浦东新区",
		},
		"310116": {
			Province: "上海市",
			City:     "上海市",
			County:   "金山区",
		},
		"310117": {
			Province: "上海市",
			City:     "上海市",
			County:   "松江区",
		},
		"310118": {
			Province: "上海市",
			City:     "上海市",
			County:   "青浦区",
		},
		"310120": {
			Province: "上海市",
			City:     "上海市",
			County:   "奉贤区",
		},
		"310151": {
			Province: "上海市",
			City:     "上海市",
			County:   "崇明区",
		},
		"320000": {
			Province: "江苏省",
			City:     "江苏省",
			County:   "江苏省",
		},
		"320100": {
			Province: "江苏省",
			City:     "南京市",
			County:   "南京市",
		},
		"320102": {
			Province: "江苏省",
			City:     "南京市",
			County:   "玄武区",
		},
		"320104": {
			Province: "江苏省",
			City:     "南京市",
			County:   "秦淮区",
		},
		"320105": {
			Province: "江苏省",
			City:     "南京市",
			County:   "建邺区",
		},
		"320106": {
			Province: "江苏省",
			City:     "南京市",
			County:   "鼓楼区",
		},
		"320111": {
			Province: "江苏省",
			City:     "南京市",
			County:   "浦口区",
		},
		"320113": {
			Province: "江苏省",
			City:     "南京市",
			County:   "栖霞区",
		},
		"320114": {
			Province: "江苏省",
			City:     "南京市",
			County:   "雨花台区",
		},
		"320115": {
			Province: "江苏省",
			City:     "南京市",
			County:   "江宁区",
		},
		"320116": {
			Province: "江苏省",
			City:     "南京市",
			County:   "六合区",
		},
		"320117": {
			Province: "江苏省",
			City:     "南京市",
			County:   "溧水区",
		},
		"320118": {
			Province: "江苏省",
			City:     "南京市",
			County:   "高淳区",
		},
		"320200": {
			Province: "江苏省",
			City:     "无锡市",
			County:   "无锡市",
		},
		"320205": {
			Province: "江苏省",
			City:     "无锡市",
			County:   "锡山区",
		},
		"320206": {
			Province: "江苏省",
			City:     "无锡市",
			County:   "惠山区",
		},
		"320211": {
			Province: "江苏省",
			City:     "无锡市",
			County:   "滨湖区",
		},
		"320213": {
			Province: "江苏省",
			City:     "无锡市",
			County:   "梁溪区",
		},
		"320214": {
			Province: "江苏省",
			City:     "无锡市",
			County:   "新吴区",
		},
		"320281": {
			Province: "江苏省",
			City:     "无锡市",
			County:   "江阴市",
		},
		"320282": {
			Province: "江苏省",
			City:     "无锡市",
			County:   "宜兴市",
		},
		"320300": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "徐州市",
		},
		"320302": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "鼓楼区",
		},
		"320303": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "云龙区",
		},
		"320305": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "贾汪区",
		},
		"320311": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "泉山区",
		},
		"320312": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "铜山区",
		},
		"320321": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "丰县",
		},
		"320322": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "沛县",
		},
		"320324": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "睢宁县",
		},
		"320381": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "新沂市",
		},
		"320382": {
			Province: "江苏省",
			City:     "徐州市",
			County:   "邳州市",
		},
		"320400": {
			Province: "江苏省",
			City:     "常州市",
			County:   "常州市",
		},
		"320402": {
			Province: "江苏省",
			City:     "常州市",
			County:   "天宁区",
		},
		"320404": {
			Province: "江苏省",
			City:     "常州市",
			County:   "钟楼区",
		},
		"320411": {
			Province: "江苏省",
			City:     "常州市",
			County:   "新北区",
		},
		"320412": {
			Province: "江苏省",
			City:     "常州市",
			County:   "武进区",
		},
		"320413": {
			Province: "江苏省",
			City:     "常州市",
			County:   "金坛区",
		},
		"320481": {
			Province: "江苏省",
			City:     "常州市",
			County:   "溧阳市",
		},
		"320500": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "苏州市",
		},
		"320505": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "虎丘区",
		},
		"320506": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "吴中区",
		},
		"320507": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "相城区",
		},
		"320508": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "姑苏区",
		},
		"320509": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "吴江区",
		},
		"320581": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "常熟市",
		},
		"320582": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "张家港市",
		},
		"320583": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "昆山市",
		},
		"320585": {
			Province: "江苏省",
			City:     "苏州市",
			County:   "太仓市",
		},
		"320600": {
			Province: "江苏省",
			City:     "南通市",
			County:   "南通市",
		},
		"320612": {
			Province: "江苏省",
			City:     "南通市",
			County:   "通州区",
		},
		"320613": {
			Province: "江苏省",
			City:     "南通市",
			County:   "崇川区",
		},
		"320614": {
			Province: "江苏省",
			City:     "南通市",
			County:   "海门区",
		},
		"320623": {
			Province: "江苏省",
			City:     "南通市",
			County:   "如东县",
		},
		"320681": {
			Province: "江苏省",
			City:     "南通市",
			County:   "启东市",
		},
		"320682": {
			Province: "江苏省",
			City:     "南通市",
			County:   "如皋市",
		},
		"320685": {
			Province: "江苏省",
			City:     "南通市",
			County:   "海安市",
		},
		"320700": {
			Province: "江苏省",
			City:     "连云港市",
			County:   "连云港市",
		},
		"320703": {
			Province: "江苏省",
			City:     "连云港市",
			County:   "连云区",
		},
		"320706": {
			Province: "江苏省",
			City:     "连云港市",
			County:   "海州区",
		},
		"320707": {
			Province: "江苏省",
			City:     "连云港市",
			County:   "赣榆区",
		},
		"320722": {
			Province: "江苏省",
			City:     "连云港市",
			County:   "东海县",
		},
		"320723": {
			Province: "江苏省",
			City:     "连云港市",
			County:   "灌云县",
		},
		"320724": {
			Province: "江苏省",
			City:     "连云港市",
			County:   "灌南县",
		},
		"320800": {
			Province: "江苏省",
			City:     "淮安市",
			County:   "淮安市",
		},
		"320803": {
			Province: "江苏省",
			City:     "淮安市",
			County:   "淮安区",
		},
		"320804": {
			Province: "江苏省",
			City:     "淮安市",
			County:   "淮阴区",
		},
		"320812": {
			Province: "江苏省",
			City:     "淮安市",
			County:   "清江浦区",
		},
		"320813": {
			Province: "江苏省",
			City:     "淮安市",
			County:   "洪泽区",
		},
		"320826": {
			Province: "江苏省",
			City:     "淮安市",
			County:   "涟水县",
		},
		"320830": {
			Province: "江苏省",
			City:     "淮安市",
			County:   "盱眙县",
		},
		"320831": {
			Province: "江苏省",
			City:     "淮安市",
			County:   "金湖县",
		},
		"320900": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "盐城市",
		},
		"320902": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "亭湖区",
		},
		"320903": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "盐都区",
		},
		"320904": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "大丰区",
		},
		"320921": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "响水县",
		},
		"320922": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "滨海县",
		},
		"320923": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "阜宁县",
		},
		"320924": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "射阳县",
		},
		"320925": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "建湖县",
		},
		"320981": {
			Province: "江苏省",
			City:     "盐城市",
			County:   "东台市",
		},
		"321000": {
			Province: "江苏省",
			City:     "扬州市",
			County:   "扬州市",
		},
		"321002": {
			Province: "江苏省",
			City:     "扬州市",
			County:   "广陵区",
		},
		"321003": {
			Province: "江苏省",
			City:     "扬州市",
			County:   "邗江区",
		},
		"321012": {
			Province: "江苏省",
			City:     "扬州市",
			County:   "江都区",
		},
		"321023": {
			Province: "江苏省",
			City:     "扬州市",
			County:   "宝应县",
		},
		"321081": {
			Province: "江苏省",
			City:     "扬州市",
			County:   "仪征市",
		},
		"321084": {
			Province: "江苏省",
			City:     "扬州市",
			County:   "高邮市",
		},
		"321100": {
			Province: "江苏省",
			City:     "镇江市",
			County:   "镇江市",
		},
		"321102": {
			Province: "江苏省",
			City:     "镇江市",
			County:   "京口区",
		},
		"321111": {
			Province: "江苏省",
			City:     "镇江市",
			County:   "润州区",
		},
		"321112": {
			Province: "江苏省",
			City:     "镇江市",
			County:   "丹徒区",
		},
		"321181": {
			Province: "江苏省",
			City:     "镇江市",
			County:   "丹阳市",
		},
		"321182": {
			Province: "江苏省",
			City:     "镇江市",
			County:   "扬中市",
		},
		"321183": {
			Province: "江苏省",
			City:     "镇江市",
			County:   "句容市",
		},
		"321200": {
			Province: "江苏省",
			City:     "泰州市",
			County:   "泰州市",
		},
		"321202": {
			Province: "江苏省",
			City:     "泰州市",
			County:   "海陵区",
		},
		"321203": {
			Province: "江苏省",
			City:     "泰州市",
			County:   "高港区",
		},
		"321204": {
			Province: "江苏省",
			City:     "泰州市",
			County:   "姜堰区",
		},
		"321281": {
			Province: "江苏省",
			City:     "泰州市",
			County:   "兴化市",
		},
		"321282": {
			Province: "江苏省",
			City:     "泰州市",
			County:   "靖江市",
		},
		"321283": {
			Province: "江苏省",
			City:     "泰州市",
			County:   "泰兴市",
		},
		"321300": {
			Province: "江苏省",
			City:     "宿迁市",
			County:   "宿迁市",
		},
		"321302": {
			Province: "江苏省",
			City:     "宿迁市",
			County:   "宿城区",
		},
		"321311": {
			Province: "江苏省",
			City:     "宿迁市",
			County:   "宿豫区",
		},
		"321322": {
			Province: "江苏省",
			City:     "宿迁市",
			County:   "沭阳县",
		},
		"321323": {
			Province: "江苏省",
			City:     "宿迁市",
			County:   "泗阳县",
		},
		"321324": {
			Province: "江苏省",
			City:     "宿迁市",
			County:   "泗洪县",
		},
		"330000": {
			Province: "浙江省",
			City:     "浙江省",
			County:   "浙江省",
		},
		"330100": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "杭州市",
		},
		"330102": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "上城区",
		},
		"330103": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "下城区",
		},
		"330104": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "江干区",
		},
		"330105": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "拱墅区",
		},
		"330106": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "西湖区",
		},
		"330108": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "滨江区",
		},
		"330109": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "萧山区",
		},
		"330110": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "余杭区",
		},
		"330111": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "富阳区",
		},
		"330112": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "临安区",
		},
		"330122": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "桐庐县",
		},
		"330127": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "淳安县",
		},
		"330182": {
			Province: "浙江省",
			City:     "杭州市",
			County:   "建德市",
		},
		"330200": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "宁波市",
		},
		"330203": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "海曙区",
		},
		"330205": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "江北区",
		},
		"330206": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "北仑区",
		},
		"330211": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "镇海区",
		},
		"330212": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "鄞州区",
		},
		"330213": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "奉化区",
		},
		"330225": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "象山县",
		},
		"330226": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "宁海县",
		},
		"330281": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "余姚市",
		},
		"330282": {
			Province: "浙江省",
			City:     "宁波市",
			County:   "慈溪市",
		},
		"330300": {
			Province: "浙江省",
			City:     "温州市",
			County:   "温州市",
		},
		"330302": {
			Province: "浙江省",
			City:     "温州市",
			County:   "鹿城区",
		},
		"330303": {
			Province: "浙江省",
			City:     "温州市",
			County:   "龙湾区",
		},
		"330304": {
			Province: "浙江省",
			City:     "温州市",
			County:   "瓯海区",
		},
		"330305": {
			Province: "浙江省",
			City:     "温州市",
			County:   "洞头区",
		},
		"330324": {
			Province: "浙江省",
			City:     "温州市",
			County:   "永嘉县",
		},
		"330326": {
			Province: "浙江省",
			City:     "温州市",
			County:   "平阳县",
		},
		"330327": {
			Province: "浙江省",
			City:     "温州市",
			County:   "苍南县",
		},
		"330328": {
			Province: "浙江省",
			City:     "温州市",
			County:   "文成县",
		},
		"330329": {
			Province: "浙江省",
			City:     "温州市",
			County:   "泰顺县",
		},
		"330381": {
			Province: "浙江省",
			City:     "温州市",
			County:   "瑞安市",
		},
		"330382": {
			Province: "浙江省",
			City:     "温州市",
			County:   "乐清市",
		},
		"330383": {
			Province: "浙江省",
			City:     "温州市",
			County:   "龙港市",
		},
		"330400": {
			Province: "浙江省",
			City:     "嘉兴市",
			County:   "嘉兴市",
		},
		"330402": {
			Province: "浙江省",
			City:     "嘉兴市",
			County:   "南湖区",
		},
		"330411": {
			Province: "浙江省",
			City:     "嘉兴市",
			County:   "秀洲区",
		},
		"330421": {
			Province: "浙江省",
			City:     "嘉兴市",
			County:   "嘉善县",
		},
		"330424": {
			Province: "浙江省",
			City:     "嘉兴市",
			County:   "海盐县",
		},
		"330481": {
			Province: "浙江省",
			City:     "嘉兴市",
			County:   "海宁市",
		},
		"330482": {
			Province: "浙江省",
			City:     "嘉兴市",
			County:   "平湖市",
		},
		"330483": {
			Province: "浙江省",
			City:     "嘉兴市",
			County:   "桐乡市",
		},
		"330500": {
			Province: "浙江省",
			City:     "湖州市",
			County:   "湖州市",
		},
		"330502": {
			Province: "浙江省",
			City:     "湖州市",
			County:   "吴兴区",
		},
		"330503": {
			Province: "浙江省",
			City:     "湖州市",
			County:   "南浔区",
		},
		"330521": {
			Province: "浙江省",
			City:     "湖州市",
			County:   "德清县",
		},
		"330522": {
			Province: "浙江省",
			City:     "湖州市",
			County:   "长兴县",
		},
		"330523": {
			Province: "浙江省",
			City:     "湖州市",
			County:   "安吉县",
		},
		"330600": {
			Province: "浙江省",
			City:     "绍兴市",
			County:   "绍兴市",
		},
		"330602": {
			Province: "浙江省",
			City:     "绍兴市",
			County:   "越城区",
		},
		"330603": {
			Province: "浙江省",
			City:     "绍兴市",
			County:   "柯桥区",
		},
		"330604": {
			Province: "浙江省",
			City:     "绍兴市",
			County:   "上虞区",
		},
		"330624": {
			Province: "浙江省",
			City:     "绍兴市",
			County:   "新昌县",
		},
		"330681": {
			Province: "浙江省",
			City:     "绍兴市",
			County:   "诸暨市",
		},
		"330683": {
			Province: "浙江省",
			City:     "绍兴市",
			County:   "嵊州市",
		},
		"330700": {
			Province: "浙江省",
			City:     "金华市",
			County:   "金华市",
		},
		"330702": {
			Province: "浙江省",
			City:     "金华市",
			County:   "婺城区",
		},
		"330703": {
			Province: "浙江省",
			City:     "金华市",
			County:   "金东区",
		},
		"330723": {
			Province: "浙江省",
			City:     "金华市",
			County:   "武义县",
		},
		"330726": {
			Province: "浙江省",
			City:     "金华市",
			County:   "浦江县",
		},
		"330727": {
			Province: "浙江省",
			City:     "金华市",
			County:   "磐安县",
		},
		"330781": {
			Province: "浙江省",
			City:     "金华市",
			County:   "兰溪市",
		},
		"330782": {
			Province: "浙江省",
			City:     "金华市",
			County:   "义乌市",
		},
		"330783": {
			Province: "浙江省",
			City:     "金华市",
			County:   "东阳市",
		},
		"330784": {
			Province: "浙江省",
			City:     "金华市",
			County:   "永康市",
		},
		"330800": {
			Province: "浙江省",
			City:     "衢州市",
			County:   "衢州市",
		},
		"330802": {
			Province: "浙江省",
			City:     "衢州市",
			County:   "柯城区",
		},
		"330803": {
			Province: "浙江省",
			City:     "衢州市",
			County:   "衢江区",
		},
		"330822": {
			Province: "浙江省",
			City:     "衢州市",
			County:   "常山县",
		},
		"330824": {
			Province: "浙江省",
			City:     "衢州市",
			County:   "开化县",
		},
		"330825": {
			Province: "浙江省",
			City:     "衢州市",
			County:   "龙游县",
		},
		"330881": {
			Province: "浙江省",
			City:     "衢州市",
			County:   "江山市",
		},
		"330900": {
			Province: "浙江省",
			City:     "舟山市",
			County:   "舟山市",
		},
		"330902": {
			Province: "浙江省",
			City:     "舟山市",
			County:   "定海区",
		},
		"330903": {
			Province: "浙江省",
			City:     "舟山市",
			County:   "普陀区",
		},
		"330921": {
			Province: "浙江省",
			City:     "舟山市",
			County:   "岱山县",
		},
		"330922": {
			Province: "浙江省",
			City:     "舟山市",
			County:   "嵊泗县",
		},
		"331000": {
			Province: "浙江省",
			City:     "台州市",
			County:   "台州市",
		},
		"331002": {
			Province: "浙江省",
			City:     "台州市",
			County:   "椒江区",
		},
		"331003": {
			Province: "浙江省",
			City:     "台州市",
			County:   "黄岩区",
		},
		"331004": {
			Province: "浙江省",
			City:     "台州市",
			County:   "路桥区",
		},
		"331022": {
			Province: "浙江省",
			City:     "台州市",
			County:   "三门县",
		},
		"331023": {
			Province: "浙江省",
			City:     "台州市",
			County:   "天台县",
		},
		"331024": {
			Province: "浙江省",
			City:     "台州市",
			County:   "仙居县",
		},
		"331081": {
			Province: "浙江省",
			City:     "台州市",
			County:   "温岭市",
		},
		"331082": {
			Province: "浙江省",
			City:     "台州市",
			County:   "临海市",
		},
		"331083": {
			Province: "浙江省",
			City:     "台州市",
			County:   "玉环市",
		},
		"331100": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "丽水市",
		},
		"331102": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "莲都区",
		},
		"331121": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "青田县",
		},
		"331122": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "缙云县",
		},
		"331123": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "遂昌县",
		},
		"331124": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "松阳县",
		},
		"331125": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "云和县",
		},
		"331126": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "庆元县",
		},
		"331127": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "景宁畲族自治县",
		},
		"331181": {
			Province: "浙江省",
			City:     "丽水市",
			County:   "龙泉市",
		},
		"340000": {
			Province: "安徽省",
			City:     "安徽省",
			County:   "安徽省",
		},
		"340100": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "合肥市",
		},
		"340102": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "瑶海区",
		},
		"340103": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "庐阳区",
		},
		"340104": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "蜀山区",
		},
		"340111": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "包河区",
		},
		"340121": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "长丰县",
		},
		"340122": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "肥东县",
		},
		"340123": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "肥西县",
		},
		"340124": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "庐江县",
		},
		"340181": {
			Province: "安徽省",
			City:     "合肥市",
			County:   "巢湖市",
		},
		"340200": {
			Province: "安徽省",
			City:     "芜湖市",
			County:   "芜湖市",
		},
		"340202": {
			Province: "安徽省",
			City:     "芜湖市",
			County:   "镜湖区",
		},
		"340207": {
			Province: "安徽省",
			City:     "芜湖市",
			County:   "鸠江区",
		},
		"340209": {
			Province: "安徽省",
			City:     "芜湖市",
			County:   "弋江区",
		},
		"340210": {
			Province: "安徽省",
			City:     "芜湖市",
			County:   "湾沚区",
		},
		"340212": {
			Province: "安徽省",
			City:     "芜湖市",
			County:   "繁昌区",
		},
		"340223": {
			Province: "安徽省",
			City:     "芜湖市",
			County:   "南陵县",
		},
		"340281": {
			Province: "安徽省",
			City:     "芜湖市",
			County:   "无为市",
		},
		"340300": {
			Province: "安徽省",
			City:     "蚌埠市",
			County:   "蚌埠市",
		},
		"340302": {
			Province: "安徽省",
			City:     "蚌埠市",
			County:   "龙子湖区",
		},
		"340303": {
			Province: "安徽省",
			City:     "蚌埠市",
			County:   "蚌山区",
		},
		"340304": {
			Province: "安徽省",
			City:     "蚌埠市",
			County:   "禹会区",
		},
		"340311": {
			Province: "安徽省",
			City:     "蚌埠市",
			County:   "淮上区",
		},
		"340321": {
			Province: "安徽省",
			City:     "蚌埠市",
			County:   "怀远县",
		},
		"340322": {
			Province: "安徽省",
			City:     "蚌埠市",
			County:   "五河县",
		},
		"340323": {
			Province: "安徽省",
			City:     "蚌埠市",
			County:   "固镇县",
		},
		"340400": {
			Province: "安徽省",
			City:     "淮南市",
			County:   "淮南市",
		},
		"340402": {
			Province: "安徽省",
			City:     "淮南市",
			County:   "大通区",
		},
		"340403": {
			Province: "安徽省",
			City:     "淮南市",
			County:   "田家庵区",
		},
		"340404": {
			Province: "安徽省",
			City:     "淮南市",
			County:   "谢家集区",
		},
		"340405": {
			Province: "安徽省",
			City:     "淮南市",
			County:   "八公山区",
		},
		"340406": {
			Province: "安徽省",
			City:     "淮南市",
			County:   "潘集区",
		},
		"340421": {
			Province: "安徽省",
			City:     "淮南市",
			County:   "凤台县",
		},
		"340422": {
			Province: "安徽省",
			City:     "淮南市",
			County:   "寿县",
		},
		"340500": {
			Province: "安徽省",
			City:     "马鞍山市",
			County:   "马鞍山市",
		},
		"340503": {
			Province: "安徽省",
			City:     "马鞍山市",
			County:   "花山区",
		},
		"340504": {
			Province: "安徽省",
			City:     "马鞍山市",
			County:   "雨山区",
		},
		"340506": {
			Province: "安徽省",
			City:     "马鞍山市",
			County:   "博望区",
		},
		"340521": {
			Province: "安徽省",
			City:     "马鞍山市",
			County:   "当涂县",
		},
		"340522": {
			Province: "安徽省",
			City:     "马鞍山市",
			County:   "含山县",
		},
		"340523": {
			Province: "安徽省",
			City:     "马鞍山市",
			County:   "和县",
		},
		"340600": {
			Province: "安徽省",
			City:     "淮北市",
			County:   "淮北市",
		},
		"340602": {
			Province: "安徽省",
			City:     "淮北市",
			County:   "杜集区",
		},
		"340603": {
			Province: "安徽省",
			City:     "淮北市",
			County:   "相山区",
		},
		"340604": {
			Province: "安徽省",
			City:     "淮北市",
			County:   "烈山区",
		},
		"340621": {
			Province: "安徽省",
			City:     "淮北市",
			County:   "濉溪县",
		},
		"340700": {
			Province: "安徽省",
			City:     "铜陵市",
			County:   "铜陵市",
		},
		"340705": {
			Province: "安徽省",
			City:     "铜陵市",
			County:   "铜官区",
		},
		"340706": {
			Province: "安徽省",
			City:     "铜陵市",
			County:   "义安区",
		},
		"340711": {
			Province: "安徽省",
			City:     "铜陵市",
			County:   "郊区",
		},
		"340722": {
			Province: "安徽省",
			City:     "铜陵市",
			County:   "枞阳县",
		},
		"340800": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "安庆市",
		},
		"340802": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "迎江区",
		},
		"340803": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "大观区",
		},
		"340811": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "宜秀区",
		},
		"340822": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "怀宁县",
		},
		"340825": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "太湖县",
		},
		"340826": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "宿松县",
		},
		"340827": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "望江县",
		},
		"340828": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "岳西县",
		},
		"340881": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "桐城市",
		},
		"340882": {
			Province: "安徽省",
			City:     "安庆市",
			County:   "潜山市",
		},
		"341000": {
			Province: "安徽省",
			City:     "黄山市",
			County:   "黄山市",
		},
		"341002": {
			Province: "安徽省",
			City:     "黄山市",
			County:   "屯溪区",
		},
		"341003": {
			Province: "安徽省",
			City:     "黄山市",
			County:   "黄山区",
		},
		"341004": {
			Province: "安徽省",
			City:     "黄山市",
			County:   "徽州区",
		},
		"341021": {
			Province: "安徽省",
			City:     "黄山市",
			County:   "歙县",
		},
		"341022": {
			Province: "安徽省",
			City:     "黄山市",
			County:   "休宁县",
		},
		"341023": {
			Province: "安徽省",
			City:     "黄山市",
			County:   "黟县",
		},
		"341024": {
			Province: "安徽省",
			City:     "黄山市",
			County:   "祁门县",
		},
		"341100": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "滁州市",
		},
		"341102": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "琅琊区",
		},
		"341103": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "南谯区",
		},
		"341122": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "来安县",
		},
		"341124": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "全椒县",
		},
		"341125": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "定远县",
		},
		"341126": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "凤阳县",
		},
		"341181": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "天长市",
		},
		"341182": {
			Province: "安徽省",
			City:     "滁州市",
			County:   "明光市",
		},
		"341200": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "阜阳市",
		},
		"341202": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "颍州区",
		},
		"341203": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "颍东区",
		},
		"341204": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "颍泉区",
		},
		"341221": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "临泉县",
		},
		"341222": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "太和县",
		},
		"341225": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "阜南县",
		},
		"341226": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "颍上县",
		},
		"341282": {
			Province: "安徽省",
			City:     "阜阳市",
			County:   "界首市",
		},
		"341300": {
			Province: "安徽省",
			City:     "宿州市",
			County:   "宿州市",
		},
		"341302": {
			Province: "安徽省",
			City:     "宿州市",
			County:   "埇桥区",
		},
		"341321": {
			Province: "安徽省",
			City:     "宿州市",
			County:   "砀山县",
		},
		"341322": {
			Province: "安徽省",
			City:     "宿州市",
			County:   "萧县",
		},
		"341323": {
			Province: "安徽省",
			City:     "宿州市",
			County:   "灵璧县",
		},
		"341324": {
			Province: "安徽省",
			City:     "宿州市",
			County:   "泗县",
		},
		"341500": {
			Province: "安徽省",
			City:     "六安市",
			County:   "六安市",
		},
		"341502": {
			Province: "安徽省",
			City:     "六安市",
			County:   "金安区",
		},
		"341503": {
			Province: "安徽省",
			City:     "六安市",
			County:   "裕安区",
		},
		"341504": {
			Province: "安徽省",
			City:     "六安市",
			County:   "叶集区",
		},
		"341522": {
			Province: "安徽省",
			City:     "六安市",
			County:   "霍邱县",
		},
		"341523": {
			Province: "安徽省",
			City:     "六安市",
			County:   "舒城县",
		},
		"341524": {
			Province: "安徽省",
			City:     "六安市",
			County:   "金寨县",
		},
		"341525": {
			Province: "安徽省",
			City:     "六安市",
			County:   "霍山县",
		},
		"341600": {
			Province: "安徽省",
			City:     "亳州市",
			County:   "亳州市",
		},
		"341602": {
			Province: "安徽省",
			City:     "亳州市",
			County:   "谯城区",
		},
		"341621": {
			Province: "安徽省",
			City:     "亳州市",
			County:   "涡阳县",
		},
		"341622": {
			Province: "安徽省",
			City:     "亳州市",
			County:   "蒙城县",
		},
		"341623": {
			Province: "安徽省",
			City:     "亳州市",
			County:   "利辛县",
		},
		"341700": {
			Province: "安徽省",
			City:     "池州市",
			County:   "池州市",
		},
		"341702": {
			Province: "安徽省",
			City:     "池州市",
			County:   "贵池区",
		},
		"341721": {
			Province: "安徽省",
			City:     "池州市",
			County:   "东至县",
		},
		"341722": {
			Province: "安徽省",
			City:     "池州市",
			County:   "石台县",
		},
		"341723": {
			Province: "安徽省",
			City:     "池州市",
			County:   "青阳县",
		},
		"341800": {
			Province: "安徽省",
			City:     "宣城市",
			County:   "宣城市",
		},
		"341802": {
			Province: "安徽省",
			City:     "宣城市",
			County:   "宣州区",
		},
		"341821": {
			Province: "安徽省",
			City:     "宣城市",
			County:   "郎溪县",
		},
		"341823": {
			Province: "安徽省",
			City:     "宣城市",
			County:   "泾县",
		},
		"341824": {
			Province: "安徽省",
			City:     "宣城市",
			County:   "绩溪县",
		},
		"341825": {
			Province: "安徽省",
			City:     "宣城市",
			County:   "旌德县",
		},
		"341881": {
			Province: "安徽省",
			City:     "宣城市",
			County:   "宁国市",
		},
		"341882": {
			Province: "安徽省",
			City:     "宣城市",
			County:   "广德市",
		},
		"350000": {
			Province: "福建省",
			City:     "福建省",
			County:   "福建省",
		},
		"350100": {
			Province: "福建省",
			City:     "福州市",
			County:   "福州市",
		},
		"350102": {
			Province: "福建省",
			City:     "福州市",
			County:   "鼓楼区",
		},
		"350103": {
			Province: "福建省",
			City:     "福州市",
			County:   "台江区",
		},
		"350104": {
			Province: "福建省",
			City:     "福州市",
			County:   "仓山区",
		},
		"350105": {
			Province: "福建省",
			City:     "福州市",
			County:   "马尾区",
		},
		"350111": {
			Province: "福建省",
			City:     "福州市",
			County:   "晋安区",
		},
		"350112": {
			Province: "福建省",
			City:     "福州市",
			County:   "长乐区",
		},
		"350121": {
			Province: "福建省",
			City:     "福州市",
			County:   "闽侯县",
		},
		"350122": {
			Province: "福建省",
			City:     "福州市",
			County:   "连江县",
		},
		"350123": {
			Province: "福建省",
			City:     "福州市",
			County:   "罗源县",
		},
		"350124": {
			Province: "福建省",
			City:     "福州市",
			County:   "闽清县",
		},
		"350125": {
			Province: "福建省",
			City:     "福州市",
			County:   "永泰县",
		},
		"350128": {
			Province: "福建省",
			City:     "福州市",
			County:   "平潭县",
		},
		"350181": {
			Province: "福建省",
			City:     "福州市",
			County:   "福清市",
		},
		"350200": {
			Province: "福建省",
			City:     "厦门市",
			County:   "厦门市",
		},
		"350203": {
			Province: "福建省",
			City:     "厦门市",
			County:   "思明区",
		},
		"350205": {
			Province: "福建省",
			City:     "厦门市",
			County:   "海沧区",
		},
		"350206": {
			Province: "福建省",
			City:     "厦门市",
			County:   "湖里区",
		},
		"350211": {
			Province: "福建省",
			City:     "厦门市",
			County:   "集美区",
		},
		"350212": {
			Province: "福建省",
			City:     "厦门市",
			County:   "同安区",
		},
		"350213": {
			Province: "福建省",
			City:     "厦门市",
			County:   "翔安区",
		},
		"350300": {
			Province: "福建省",
			City:     "莆田市",
			County:   "莆田市",
		},
		"350302": {
			Province: "福建省",
			City:     "莆田市",
			County:   "城厢区",
		},
		"350303": {
			Province: "福建省",
			City:     "莆田市",
			County:   "涵江区",
		},
		"350304": {
			Province: "福建省",
			City:     "莆田市",
			County:   "荔城区",
		},
		"350305": {
			Province: "福建省",
			City:     "莆田市",
			County:   "秀屿区",
		},
		"350322": {
			Province: "福建省",
			City:     "莆田市",
			County:   "仙游县",
		},
		"350400": {
			Province: "福建省",
			City:     "三明市",
			County:   "三明市",
		},
		"350402": {
			Province: "福建省",
			City:     "三明市",
			County:   "梅列区",
		},
		"350403": {
			Province: "福建省",
			City:     "三明市",
			County:   "三元区",
		},
		"350421": {
			Province: "福建省",
			City:     "三明市",
			County:   "明溪县",
		},
		"350423": {
			Province: "福建省",
			City:     "三明市",
			County:   "清流县",
		},
		"350424": {
			Province: "福建省",
			City:     "三明市",
			County:   "宁化县",
		},
		"350425": {
			Province: "福建省",
			City:     "三明市",
			County:   "大田县",
		},
		"350426": {
			Province: "福建省",
			City:     "三明市",
			County:   "尤溪县",
		},
		"350427": {
			Province: "福建省",
			City:     "三明市",
			County:   "沙县",
		},
		"350428": {
			Province: "福建省",
			City:     "三明市",
			County:   "将乐县",
		},
		"350429": {
			Province: "福建省",
			City:     "三明市",
			County:   "泰宁县",
		},
		"350430": {
			Province: "福建省",
			City:     "三明市",
			County:   "建宁县",
		},
		"350481": {
			Province: "福建省",
			City:     "三明市",
			County:   "永安市",
		},
		"350500": {
			Province: "福建省",
			City:     "泉州市",
			County:   "泉州市",
		},
		"350502": {
			Province: "福建省",
			City:     "泉州市",
			County:   "鲤城区",
		},
		"350503": {
			Province: "福建省",
			City:     "泉州市",
			County:   "丰泽区",
		},
		"350504": {
			Province: "福建省",
			City:     "泉州市",
			County:   "洛江区",
		},
		"350505": {
			Province: "福建省",
			City:     "泉州市",
			County:   "泉港区",
		},
		"350521": {
			Province: "福建省",
			City:     "泉州市",
			County:   "惠安县",
		},
		"350524": {
			Province: "福建省",
			City:     "泉州市",
			County:   "安溪县",
		},
		"350525": {
			Province: "福建省",
			City:     "泉州市",
			County:   "永春县",
		},
		"350526": {
			Province: "福建省",
			City:     "泉州市",
			County:   "德化县",
		},
		"350527": {
			Province: "福建省",
			City:     "泉州市",
			County:   "金门县",
		},
		"350581": {
			Province: "福建省",
			City:     "泉州市",
			County:   "石狮市",
		},
		"350582": {
			Province: "福建省",
			City:     "泉州市",
			County:   "晋江市",
		},
		"350583": {
			Province: "福建省",
			City:     "泉州市",
			County:   "南安市",
		},
		"350600": {
			Province: "福建省",
			City:     "漳州市",
			County:   "漳州市",
		},
		"350602": {
			Province: "福建省",
			City:     "漳州市",
			County:   "芗城区",
		},
		"350603": {
			Province: "福建省",
			City:     "漳州市",
			County:   "龙文区",
		},
		"350622": {
			Province: "福建省",
			City:     "漳州市",
			County:   "云霄县",
		},
		"350623": {
			Province: "福建省",
			City:     "漳州市",
			County:   "漳浦县",
		},
		"350624": {
			Province: "福建省",
			City:     "漳州市",
			County:   "诏安县",
		},
		"350625": {
			Province: "福建省",
			City:     "漳州市",
			County:   "长泰县",
		},
		"350626": {
			Province: "福建省",
			City:     "漳州市",
			County:   "东山县",
		},
		"350627": {
			Province: "福建省",
			City:     "漳州市",
			County:   "南靖县",
		},
		"350628": {
			Province: "福建省",
			City:     "漳州市",
			County:   "平和县",
		},
		"350629": {
			Province: "福建省",
			City:     "漳州市",
			County:   "华安县",
		},
		"350681": {
			Province: "福建省",
			City:     "漳州市",
			County:   "龙海市",
		},
		"350700": {
			Province: "福建省",
			City:     "南平市",
			County:   "南平市",
		},
		"350702": {
			Province: "福建省",
			City:     "南平市",
			County:   "延平区",
		},
		"350703": {
			Province: "福建省",
			City:     "南平市",
			County:   "建阳区",
		},
		"350721": {
			Province: "福建省",
			City:     "南平市",
			County:   "顺昌县",
		},
		"350722": {
			Province: "福建省",
			City:     "南平市",
			County:   "浦城县",
		},
		"350723": {
			Province: "福建省",
			City:     "南平市",
			County:   "光泽县",
		},
		"350724": {
			Province: "福建省",
			City:     "南平市",
			County:   "松溪县",
		},
		"350725": {
			Province: "福建省",
			City:     "南平市",
			County:   "政和县",
		},
		"350781": {
			Province: "福建省",
			City:     "南平市",
			County:   "邵武市",
		},
		"350782": {
			Province: "福建省",
			City:     "南平市",
			County:   "武夷山市",
		},
		"350783": {
			Province: "福建省",
			City:     "南平市",
			County:   "建瓯市",
		},
		"350800": {
			Province: "福建省",
			City:     "龙岩市",
			County:   "龙岩市",
		},
		"350802": {
			Province: "福建省",
			City:     "龙岩市",
			County:   "新罗区",
		},
		"350803": {
			Province: "福建省",
			City:     "龙岩市",
			County:   "永定区",
		},
		"350821": {
			Province: "福建省",
			City:     "龙岩市",
			County:   "长汀县",
		},
		"350823": {
			Province: "福建省",
			City:     "龙岩市",
			County:   "上杭县",
		},
		"350824": {
			Province: "福建省",
			City:     "龙岩市",
			County:   "武平县",
		},
		"350825": {
			Province: "福建省",
			City:     "龙岩市",
			County:   "连城县",
		},
		"350881": {
			Province: "福建省",
			City:     "龙岩市",
			County:   "漳平市",
		},
		"350900": {
			Province: "福建省",
			City:     "宁德市",
			County:   "宁德市",
		},
		"350902": {
			Province: "福建省",
			City:     "宁德市",
			County:   "蕉城区",
		},
		"350921": {
			Province: "福建省",
			City:     "宁德市",
			County:   "霞浦县",
		},
		"350922": {
			Province: "福建省",
			City:     "宁德市",
			County:   "古田县",
		},
		"350923": {
			Province: "福建省",
			City:     "宁德市",
			County:   "屏南县",
		},
		"350924": {
			Province: "福建省",
			City:     "宁德市",
			County:   "寿宁县",
		},
		"350925": {
			Province: "福建省",
			City:     "宁德市",
			County:   "周宁县",
		},
		"350926": {
			Province: "福建省",
			City:     "宁德市",
			County:   "柘荣县",
		},
		"350981": {
			Province: "福建省",
			City:     "宁德市",
			County:   "福安市",
		},
		"350982": {
			Province: "福建省",
			City:     "宁德市",
			County:   "福鼎市",
		},
		"360000": {
			Province: "江西省",
			City:     "江西省",
			County:   "江西省",
		},
		"360100": {
			Province: "江西省",
			City:     "南昌市",
			County:   "南昌市",
		},
		"360102": {
			Province: "江西省",
			City:     "南昌市",
			County:   "东湖区",
		},
		"360103": {
			Province: "江西省",
			City:     "南昌市",
			County:   "西湖区",
		},
		"360104": {
			Province: "江西省",
			City:     "南昌市",
			County:   "青云谱区",
		},
		"360111": {
			Province: "江西省",
			City:     "南昌市",
			County:   "青山湖区",
		},
		"360112": {
			Province: "江西省",
			City:     "南昌市",
			County:   "新建区",
		},
		"360113": {
			Province: "江西省",
			City:     "南昌市",
			County:   "红谷滩区",
		},
		"360121": {
			Province: "江西省",
			City:     "南昌市",
			County:   "南昌县",
		},
		"360123": {
			Province: "江西省",
			City:     "南昌市",
			County:   "安义县",
		},
		"360124": {
			Province: "江西省",
			City:     "南昌市",
			County:   "进贤县",
		},
		"360200": {
			Province: "江西省",
			City:     "景德镇市",
			County:   "景德镇市",
		},
		"360202": {
			Province: "江西省",
			City:     "景德镇市",
			County:   "昌江区",
		},
		"360203": {
			Province: "江西省",
			City:     "景德镇市",
			County:   "珠山区",
		},
		"360222": {
			Province: "江西省",
			City:     "景德镇市",
			County:   "浮梁县",
		},
		"360281": {
			Province: "江西省",
			City:     "景德镇市",
			County:   "乐平市",
		},
		"360300": {
			Province: "江西省",
			City:     "萍乡市",
			County:   "萍乡市",
		},
		"360302": {
			Province: "江西省",
			City:     "萍乡市",
			County:   "安源区",
		},
		"360313": {
			Province: "江西省",
			City:     "萍乡市",
			County:   "湘东区",
		},
		"360321": {
			Province: "江西省",
			City:     "萍乡市",
			County:   "莲花县",
		},
		"360322": {
			Province: "江西省",
			City:     "萍乡市",
			County:   "上栗县",
		},
		"360323": {
			Province: "江西省",
			City:     "萍乡市",
			County:   "芦溪县",
		},
		"360400": {
			Province: "江西省",
			City:     "九江市",
			County:   "九江市",
		},
		"360402": {
			Province: "江西省",
			City:     "九江市",
			County:   "濂溪区",
		},
		"360403": {
			Province: "江西省",
			City:     "九江市",
			County:   "浔阳区",
		},
		"360404": {
			Province: "江西省",
			City:     "九江市",
			County:   "柴桑区",
		},
		"360423": {
			Province: "江西省",
			City:     "九江市",
			County:   "武宁县",
		},
		"360424": {
			Province: "江西省",
			City:     "九江市",
			County:   "修水县",
		},
		"360425": {
			Province: "江西省",
			City:     "九江市",
			County:   "永修县",
		},
		"360426": {
			Province: "江西省",
			City:     "九江市",
			County:   "德安县",
		},
		"360428": {
			Province: "江西省",
			City:     "九江市",
			County:   "都昌县",
		},
		"360429": {
			Province: "江西省",
			City:     "九江市",
			County:   "湖口县",
		},
		"360430": {
			Province: "江西省",
			City:     "九江市",
			County:   "彭泽县",
		},
		"360481": {
			Province: "江西省",
			City:     "九江市",
			County:   "瑞昌市",
		},
		"360482": {
			Province: "江西省",
			City:     "九江市",
			County:   "共青城市",
		},
		"360483": {
			Province: "江西省",
			City:     "九江市",
			County:   "庐山市",
		},
		"360500": {
			Province: "江西省",
			City:     "新余市",
			County:   "新余市",
		},
		"360502": {
			Province: "江西省",
			City:     "新余市",
			County:   "渝水区",
		},
		"360521": {
			Province: "江西省",
			City:     "新余市",
			County:   "分宜县",
		},
		"360600": {
			Province: "江西省",
			City:     "鹰潭市",
			County:   "鹰潭市",
		},
		"360602": {
			Province: "江西省",
			City:     "鹰潭市",
			County:   "月湖区",
		},
		"360603": {
			Province: "江西省",
			City:     "鹰潭市",
			County:   "余江区",
		},
		"360681": {
			Province: "江西省",
			City:     "鹰潭市",
			County:   "贵溪市",
		},
		"360700": {
			Province: "江西省",
			City:     "赣州市",
			County:   "赣州市",
		},
		"360702": {
			Province: "江西省",
			City:     "赣州市",
			County:   "章贡区",
		},
		"360703": {
			Province: "江西省",
			City:     "赣州市",
			County:   "南康区",
		},
		"360704": {
			Province: "江西省",
			City:     "赣州市",
			County:   "赣县区",
		},
		"360722": {
			Province: "江西省",
			City:     "赣州市",
			County:   "信丰县",
		},
		"360723": {
			Province: "江西省",
			City:     "赣州市",
			County:   "大余县",
		},
		"360724": {
			Province: "江西省",
			City:     "赣州市",
			County:   "上犹县",
		},
		"360725": {
			Province: "江西省",
			City:     "赣州市",
			County:   "崇义县",
		},
		"360726": {
			Province: "江西省",
			City:     "赣州市",
			County:   "安远县",
		},
		"360728": {
			Province: "江西省",
			City:     "赣州市",
			County:   "定南县",
		},
		"360729": {
			Province: "江西省",
			City:     "赣州市",
			County:   "全南县",
		},
		"360730": {
			Province: "江西省",
			City:     "赣州市",
			County:   "宁都县",
		},
		"360731": {
			Province: "江西省",
			City:     "赣州市",
			County:   "于都县",
		},
		"360732": {
			Province: "江西省",
			City:     "赣州市",
			County:   "兴国县",
		},
		"360733": {
			Province: "江西省",
			City:     "赣州市",
			County:   "会昌县",
		},
		"360734": {
			Province: "江西省",
			City:     "赣州市",
			County:   "寻乌县",
		},
		"360735": {
			Province: "江西省",
			City:     "赣州市",
			County:   "石城县",
		},
		"360781": {
			Province: "江西省",
			City:     "赣州市",
			County:   "瑞金市",
		},
		"360783": {
			Province: "江西省",
			City:     "赣州市",
			County:   "龙南市",
		},
		"360800": {
			Province: "江西省",
			City:     "吉安市",
			County:   "吉安市",
		},
		"360802": {
			Province: "江西省",
			City:     "吉安市",
			County:   "吉州区",
		},
		"360803": {
			Province: "江西省",
			City:     "吉安市",
			County:   "青原区",
		},
		"360821": {
			Province: "江西省",
			City:     "吉安市",
			County:   "吉安县",
		},
		"360822": {
			Province: "江西省",
			City:     "吉安市",
			County:   "吉水县",
		},
		"360823": {
			Province: "江西省",
			City:     "吉安市",
			County:   "峡江县",
		},
		"360824": {
			Province: "江西省",
			City:     "吉安市",
			County:   "新干县",
		},
		"360825": {
			Province: "江西省",
			City:     "吉安市",
			County:   "永丰县",
		},
		"360826": {
			Province: "江西省",
			City:     "吉安市",
			County:   "泰和县",
		},
		"360827": {
			Province: "江西省",
			City:     "吉安市",
			County:   "遂川县",
		},
		"360828": {
			Province: "江西省",
			City:     "吉安市",
			County:   "万安县",
		},
		"360829": {
			Province: "江西省",
			City:     "吉安市",
			County:   "安福县",
		},
		"360830": {
			Province: "江西省",
			City:     "吉安市",
			County:   "永新县",
		},
		"360881": {
			Province: "江西省",
			City:     "吉安市",
			County:   "井冈山市",
		},
		"360900": {
			Province: "江西省",
			City:     "宜春市",
			County:   "宜春市",
		},
		"360902": {
			Province: "江西省",
			City:     "宜春市",
			County:   "袁州区",
		},
		"360921": {
			Province: "江西省",
			City:     "宜春市",
			County:   "奉新县",
		},
		"360922": {
			Province: "江西省",
			City:     "宜春市",
			County:   "万载县",
		},
		"360923": {
			Province: "江西省",
			City:     "宜春市",
			County:   "上高县",
		},
		"360924": {
			Province: "江西省",
			City:     "宜春市",
			County:   "宜丰县",
		},
		"360925": {
			Province: "江西省",
			City:     "宜春市",
			County:   "靖安县",
		},
		"360926": {
			Province: "江西省",
			City:     "宜春市",
			County:   "铜鼓县",
		},
		"360981": {
			Province: "江西省",
			City:     "宜春市",
			County:   "丰城市",
		},
		"360982": {
			Province: "江西省",
			City:     "宜春市",
			County:   "樟树市",
		},
		"360983": {
			Province: "江西省",
			City:     "宜春市",
			County:   "高安市",
		},
		"361000": {
			Province: "江西省",
			City:     "抚州市",
			County:   "抚州市",
		},
		"361002": {
			Province: "江西省",
			City:     "抚州市",
			County:   "临川区",
		},
		"361003": {
			Province: "江西省",
			City:     "抚州市",
			County:   "东乡区",
		},
		"361021": {
			Province: "江西省",
			City:     "抚州市",
			County:   "南城县",
		},
		"361022": {
			Province: "江西省",
			City:     "抚州市",
			County:   "黎川县",
		},
		"361023": {
			Province: "江西省",
			City:     "抚州市",
			County:   "南丰县",
		},
		"361024": {
			Province: "江西省",
			City:     "抚州市",
			County:   "崇仁县",
		},
		"361025": {
			Province: "江西省",
			City:     "抚州市",
			County:   "乐安县",
		},
		"361026": {
			Province: "江西省",
			City:     "抚州市",
			County:   "宜黄县",
		},
		"361027": {
			Province: "江西省",
			City:     "抚州市",
			County:   "金溪县",
		},
		"361028": {
			Province: "江西省",
			City:     "抚州市",
			County:   "资溪县",
		},
		"361030": {
			Province: "江西省",
			City:     "抚州市",
			County:   "广昌县",
		},
		"361100": {
			Province: "江西省",
			City:     "上饶市",
			County:   "上饶市",
		},
		"361102": {
			Province: "江西省",
			City:     "上饶市",
			County:   "信州区",
		},
		"361103": {
			Province: "江西省",
			City:     "上饶市",
			County:   "广丰区",
		},
		"361104": {
			Province: "江西省",
			City:     "上饶市",
			County:   "广信区",
		},
		"361123": {
			Province: "江西省",
			City:     "上饶市",
			County:   "玉山县",
		},
		"361124": {
			Province: "江西省",
			City:     "上饶市",
			County:   "铅山县",
		},
		"361125": {
			Province: "江西省",
			City:     "上饶市",
			County:   "横峰县",
		},
		"361126": {
			Province: "江西省",
			City:     "上饶市",
			County:   "弋阳县",
		},
		"361127": {
			Province: "江西省",
			City:     "上饶市",
			County:   "余干县",
		},
		"361128": {
			Province: "江西省",
			City:     "上饶市",
			County:   "鄱阳县",
		},
		"361129": {
			Province: "江西省",
			City:     "上饶市",
			County:   "万年县",
		},
		"361130": {
			Province: "江西省",
			City:     "上饶市",
			County:   "婺源县",
		},
		"361181": {
			Province: "江西省",
			City:     "上饶市",
			County:   "德兴市",
		},
		"370000": {
			Province: "山东省",
			City:     "山东省",
			County:   "山东省",
		},
		"370100": {
			Province: "山东省",
			City:     "济南市",
			County:   "济南市",
		},
		"370102": {
			Province: "山东省",
			City:     "济南市",
			County:   "历下区",
		},
		"370103": {
			Province: "山东省",
			City:     "济南市",
			County:   "市中区",
		},
		"370104": {
			Province: "山东省",
			City:     "济南市",
			County:   "槐荫区",
		},
		"370105": {
			Province: "山东省",
			City:     "济南市",
			County:   "天桥区",
		},
		"370112": {
			Province: "山东省",
			City:     "济南市",
			County:   "历城区",
		},
		"370113": {
			Province: "山东省",
			City:     "济南市",
			County:   "长清区",
		},
		"370114": {
			Province: "山东省",
			City:     "济南市",
			County:   "章丘区",
		},
		"370115": {
			Province: "山东省",
			City:     "济南市",
			County:   "济阳区",
		},
		"370116": {
			Province: "山东省",
			City:     "济南市",
			County:   "莱芜区",
		},
		"370117": {
			Province: "山东省",
			City:     "济南市",
			County:   "钢城区",
		},
		"370124": {
			Province: "山东省",
			City:     "济南市",
			County:   "平阴县",
		},
		"370126": {
			Province: "山东省",
			City:     "济南市",
			County:   "商河县",
		},
		"370200": {
			Province: "山东省",
			City:     "青岛市",
			County:   "青岛市",
		},
		"370202": {
			Province: "山东省",
			City:     "青岛市",
			County:   "市南区",
		},
		"370203": {
			Province: "山东省",
			City:     "青岛市",
			County:   "市北区",
		},
		"370211": {
			Province: "山东省",
			City:     "青岛市",
			County:   "黄岛区",
		},
		"370212": {
			Province: "山东省",
			City:     "青岛市",
			County:   "崂山区",
		},
		"370213": {
			Province: "山东省",
			City:     "青岛市",
			County:   "李沧区",
		},
		"370214": {
			Province: "山东省",
			City:     "青岛市",
			County:   "城阳区",
		},
		"370215": {
			Province: "山东省",
			City:     "青岛市",
			County:   "即墨区",
		},
		"370281": {
			Province: "山东省",
			City:     "青岛市",
			County:   "胶州市",
		},
		"370283": {
			Province: "山东省",
			City:     "青岛市",
			County:   "平度市",
		},
		"370285": {
			Province: "山东省",
			City:     "青岛市",
			County:   "莱西市",
		},
		"370300": {
			Province: "山东省",
			City:     "淄博市",
			County:   "淄博市",
		},
		"370302": {
			Province: "山东省",
			City:     "淄博市",
			County:   "淄川区",
		},
		"370303": {
			Province: "山东省",
			City:     "淄博市",
			County:   "张店区",
		},
		"370304": {
			Province: "山东省",
			City:     "淄博市",
			County:   "博山区",
		},
		"370305": {
			Province: "山东省",
			City:     "淄博市",
			County:   "临淄区",
		},
		"370306": {
			Province: "山东省",
			City:     "淄博市",
			County:   "周村区",
		},
		"370321": {
			Province: "山东省",
			City:     "淄博市",
			County:   "桓台县",
		},
		"370322": {
			Province: "山东省",
			City:     "淄博市",
			County:   "高青县",
		},
		"370323": {
			Province: "山东省",
			City:     "淄博市",
			County:   "沂源县",
		},
		"370400": {
			Province: "山东省",
			City:     "枣庄市",
			County:   "枣庄市",
		},
		"370402": {
			Province: "山东省",
			City:     "枣庄市",
			County:   "市中区",
		},
		"370403": {
			Province: "山东省",
			City:     "枣庄市",
			County:   "薛城区",
		},
		"370404": {
			Province: "山东省",
			City:     "枣庄市",
			County:   "峄城区",
		},
		"370405": {
			Province: "山东省",
			City:     "枣庄市",
			County:   "台儿庄区",
		},
		"370406": {
			Province: "山东省",
			City:     "枣庄市",
			County:   "山亭区",
		},
		"370481": {
			Province: "山东省",
			City:     "枣庄市",
			County:   "滕州市",
		},
		"370500": {
			Province: "山东省",
			City:     "东营市",
			County:   "东营市",
		},
		"370502": {
			Province: "山东省",
			City:     "东营市",
			County:   "东营区",
		},
		"370503": {
			Province: "山东省",
			City:     "东营市",
			County:   "河口区",
		},
		"370505": {
			Province: "山东省",
			City:     "东营市",
			County:   "垦利区",
		},
		"370522": {
			Province: "山东省",
			City:     "东营市",
			County:   "利津县",
		},
		"370523": {
			Province: "山东省",
			City:     "东营市",
			County:   "广饶县",
		},
		"370600": {
			Province: "山东省",
			City:     "烟台市",
			County:   "烟台市",
		},
		"370602": {
			Province: "山东省",
			City:     "烟台市",
			County:   "芝罘区",
		},
		"370611": {
			Province: "山东省",
			City:     "烟台市",
			County:   "福山区",
		},
		"370612": {
			Province: "山东省",
			City:     "烟台市",
			County:   "牟平区",
		},
		"370613": {
			Province: "山东省",
			City:     "烟台市",
			County:   "莱山区",
		},
		"370614": {
			Province: "山东省",
			City:     "烟台市",
			County:   "蓬莱区",
		},
		"370681": {
			Province: "山东省",
			City:     "烟台市",
			County:   "龙口市",
		},
		"370682": {
			Province: "山东省",
			City:     "烟台市",
			County:   "莱阳市",
		},
		"370683": {
			Province: "山东省",
			City:     "烟台市",
			County:   "莱州市",
		},
		"370685": {
			Province: "山东省",
			City:     "烟台市",
			County:   "招远市",
		},
		"370686": {
			Province: "山东省",
			City:     "烟台市",
			County:   "栖霞市",
		},
		"370687": {
			Province: "山东省",
			City:     "烟台市",
			County:   "海阳市",
		},
		"370700": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "潍坊市",
		},
		"370702": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "潍城区",
		},
		"370703": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "寒亭区",
		},
		"370704": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "坊子区",
		},
		"370705": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "奎文区",
		},
		"370724": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "临朐县",
		},
		"370725": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "昌乐县",
		},
		"370781": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "青州市",
		},
		"370782": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "诸城市",
		},
		"370783": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "寿光市",
		},
		"370784": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "安丘市",
		},
		"370785": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "高密市",
		},
		"370786": {
			Province: "山东省",
			City:     "潍坊市",
			County:   "昌邑市",
		},
		"370800": {
			Province: "山东省",
			City:     "济宁市",
			County:   "济宁市",
		},
		"370811": {
			Province: "山东省",
			City:     "济宁市",
			County:   "任城区",
		},
		"370812": {
			Province: "山东省",
			City:     "济宁市",
			County:   "兖州区",
		},
		"370826": {
			Province: "山东省",
			City:     "济宁市",
			County:   "微山县",
		},
		"370827": {
			Province: "山东省",
			City:     "济宁市",
			County:   "鱼台县",
		},
		"370828": {
			Province: "山东省",
			City:     "济宁市",
			County:   "金乡县",
		},
		"370829": {
			Province: "山东省",
			City:     "济宁市",
			County:   "嘉祥县",
		},
		"370830": {
			Province: "山东省",
			City:     "济宁市",
			County:   "汶上县",
		},
		"370831": {
			Province: "山东省",
			City:     "济宁市",
			County:   "泗水县",
		},
		"370832": {
			Province: "山东省",
			City:     "济宁市",
			County:   "梁山县",
		},
		"370881": {
			Province: "山东省",
			City:     "济宁市",
			County:   "曲阜市",
		},
		"370883": {
			Province: "山东省",
			City:     "济宁市",
			County:   "邹城市",
		},
		"370900": {
			Province: "山东省",
			City:     "泰安市",
			County:   "泰安市",
		},
		"370902": {
			Province: "山东省",
			City:     "泰安市",
			County:   "泰山区",
		},
		"370911": {
			Province: "山东省",
			City:     "泰安市",
			County:   "岱岳区",
		},
		"370921": {
			Province: "山东省",
			City:     "泰安市",
			County:   "宁阳县",
		},
		"370923": {
			Province: "山东省",
			City:     "泰安市",
			County:   "东平县",
		},
		"370982": {
			Province: "山东省",
			City:     "泰安市",
			County:   "新泰市",
		},
		"370983": {
			Province: "山东省",
			City:     "泰安市",
			County:   "肥城市",
		},
		"371000": {
			Province: "山东省",
			City:     "威海市",
			County:   "威海市",
		},
		"371002": {
			Province: "山东省",
			City:     "威海市",
			County:   "环翠区",
		},
		"371003": {
			Province: "山东省",
			City:     "威海市",
			County:   "文登区",
		},
		"371082": {
			Province: "山东省",
			City:     "威海市",
			County:   "荣成市",
		},
		"371083": {
			Province: "山东省",
			City:     "威海市",
			County:   "乳山市",
		},
		"371100": {
			Province: "山东省",
			City:     "日照市",
			County:   "日照市",
		},
		"371102": {
			Province: "山东省",
			City:     "日照市",
			County:   "东港区",
		},
		"371103": {
			Province: "山东省",
			City:     "日照市",
			County:   "岚山区",
		},
		"371121": {
			Province: "山东省",
			City:     "日照市",
			County:   "五莲县",
		},
		"371122": {
			Province: "山东省",
			City:     "日照市",
			County:   "莒县",
		},
		"371300": {
			Province: "山东省",
			City:     "临沂市",
			County:   "临沂市",
		},
		"371302": {
			Province: "山东省",
			City:     "临沂市",
			County:   "兰山区",
		},
		"371311": {
			Province: "山东省",
			City:     "临沂市",
			County:   "罗庄区",
		},
		"371312": {
			Province: "山东省",
			City:     "临沂市",
			County:   "河东区",
		},
		"371321": {
			Province: "山东省",
			City:     "临沂市",
			County:   "沂南县",
		},
		"371322": {
			Province: "山东省",
			City:     "临沂市",
			County:   "郯城县",
		},
		"371323": {
			Province: "山东省",
			City:     "临沂市",
			County:   "沂水县",
		},
		"371324": {
			Province: "山东省",
			City:     "临沂市",
			County:   "兰陵县",
		},
		"371325": {
			Province: "山东省",
			City:     "临沂市",
			County:   "费县",
		},
		"371326": {
			Province: "山东省",
			City:     "临沂市",
			County:   "平邑县",
		},
		"371327": {
			Province: "山东省",
			City:     "临沂市",
			County:   "莒南县",
		},
		"371328": {
			Province: "山东省",
			City:     "临沂市",
			County:   "蒙阴县",
		},
		"371329": {
			Province: "山东省",
			City:     "临沂市",
			County:   "临沭县",
		},
		"371400": {
			Province: "山东省",
			City:     "德州市",
			County:   "德州市",
		},
		"371402": {
			Province: "山东省",
			City:     "德州市",
			County:   "德城区",
		},
		"371403": {
			Province: "山东省",
			City:     "德州市",
			County:   "陵城区",
		},
		"371422": {
			Province: "山东省",
			City:     "德州市",
			County:   "宁津县",
		},
		"371423": {
			Province: "山东省",
			City:     "德州市",
			County:   "庆云县",
		},
		"371424": {
			Province: "山东省",
			City:     "德州市",
			County:   "临邑县",
		},
		"371425": {
			Province: "山东省",
			City:     "德州市",
			County:   "齐河县",
		},
		"371426": {
			Province: "山东省",
			City:     "德州市",
			County:   "平原县",
		},
		"371427": {
			Province: "山东省",
			City:     "德州市",
			County:   "夏津县",
		},
		"371428": {
			Province: "山东省",
			City:     "德州市",
			County:   "武城县",
		},
		"371481": {
			Province: "山东省",
			City:     "德州市",
			County:   "乐陵市",
		},
		"371482": {
			Province: "山东省",
			City:     "德州市",
			County:   "禹城市",
		},
		"371500": {
			Province: "山东省",
			City:     "聊城市",
			County:   "聊城市",
		},
		"371502": {
			Province: "山东省",
			City:     "聊城市",
			County:   "东昌府区",
		},
		"371503": {
			Province: "山东省",
			City:     "聊城市",
			County:   "茌平区",
		},
		"371521": {
			Province: "山东省",
			City:     "聊城市",
			County:   "阳谷县",
		},
		"371522": {
			Province: "山东省",
			City:     "聊城市",
			County:   "莘县",
		},
		"371524": {
			Province: "山东省",
			City:     "聊城市",
			County:   "东阿县",
		},
		"371525": {
			Province: "山东省",
			City:     "聊城市",
			County:   "冠县",
		},
		"371526": {
			Province: "山东省",
			City:     "聊城市",
			County:   "高唐县",
		},
		"371581": {
			Province: "山东省",
			City:     "聊城市",
			County:   "临清市",
		},
		"371600": {
			Province: "山东省",
			City:     "滨州市",
			County:   "滨州市",
		},
		"371602": {
			Province: "山东省",
			City:     "滨州市",
			County:   "滨城区",
		},
		"371603": {
			Province: "山东省",
			City:     "滨州市",
			County:   "沾化区",
		},
		"371621": {
			Province: "山东省",
			City:     "滨州市",
			County:   "惠民县",
		},
		"371622": {
			Province: "山东省",
			City:     "滨州市",
			County:   "阳信县",
		},
		"371623": {
			Province: "山东省",
			City:     "滨州市",
			County:   "无棣县",
		},
		"371625": {
			Province: "山东省",
			City:     "滨州市",
			County:   "博兴县",
		},
		"371681": {
			Province: "山东省",
			City:     "滨州市",
			County:   "邹平市",
		},
		"371700": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "菏泽市",
		},
		"371702": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "牡丹区",
		},
		"371703": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "定陶区",
		},
		"371721": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "曹县",
		},
		"371722": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "单县",
		},
		"371723": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "成武县",
		},
		"371724": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "巨野县",
		},
		"371725": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "郓城县",
		},
		"371726": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "鄄城县",
		},
		"371728": {
			Province: "山东省",
			City:     "菏泽市",
			County:   "东明县",
		},
		"410000": {
			Province: "河南省",
			City:     "河南省",
			County:   "河南省",
		},
		"410100": {
			Province: "河南省",
			City:     "郑州市",
			County:   "郑州市",
		},
		"410102": {
			Province: "河南省",
			City:     "郑州市",
			County:   "中原区",
		},
		"410103": {
			Province: "河南省",
			City:     "郑州市",
			County:   "二七区",
		},
		"410104": {
			Province: "河南省",
			City:     "郑州市",
			County:   "管城回族区",
		},
		"410105": {
			Province: "河南省",
			City:     "郑州市",
			County:   "金水区",
		},
		"410106": {
			Province: "河南省",
			City:     "郑州市",
			County:   "上街区",
		},
		"410108": {
			Province: "河南省",
			City:     "郑州市",
			County:   "惠济区",
		},
		"410122": {
			Province: "河南省",
			City:     "郑州市",
			County:   "中牟县",
		},
		"410181": {
			Province: "河南省",
			City:     "郑州市",
			County:   "巩义市",
		},
		"410182": {
			Province: "河南省",
			City:     "郑州市",
			County:   "荥阳市",
		},
		"410183": {
			Province: "河南省",
			City:     "郑州市",
			County:   "新密市",
		},
		"410184": {
			Province: "河南省",
			City:     "郑州市",
			County:   "新郑市",
		},
		"410185": {
			Province: "河南省",
			City:     "郑州市",
			County:   "登封市",
		},
		"410200": {
			Province: "河南省",
			City:     "开封市",
			County:   "开封市",
		},
		"410202": {
			Province: "河南省",
			City:     "开封市",
			County:   "龙亭区",
		},
		"410203": {
			Province: "河南省",
			City:     "开封市",
			County:   "顺河回族区",
		},
		"410204": {
			Province: "河南省",
			City:     "开封市",
			County:   "鼓楼区",
		},
		"410205": {
			Province: "河南省",
			City:     "开封市",
			County:   "禹王台区",
		},
		"410212": {
			Province: "河南省",
			City:     "开封市",
			County:   "祥符区",
		},
		"410221": {
			Province: "河南省",
			City:     "开封市",
			County:   "杞县",
		},
		"410222": {
			Province: "河南省",
			City:     "开封市",
			County:   "通许县",
		},
		"410223": {
			Province: "河南省",
			City:     "开封市",
			County:   "尉氏县",
		},
		"410225": {
			Province: "河南省",
			City:     "开封市",
			County:   "兰考县",
		},
		"410300": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "洛阳市",
		},
		"410302": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "老城区",
		},
		"410303": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "西工区",
		},
		"410304": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "瀍河回族区",
		},
		"410305": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "涧西区",
		},
		"410306": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "吉利区",
		},
		"410311": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "洛龙区",
		},
		"410322": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "孟津县",
		},
		"410323": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "新安县",
		},
		"410324": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "栾川县",
		},
		"410325": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "嵩县",
		},
		"410326": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "汝阳县",
		},
		"410327": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "宜阳县",
		},
		"410328": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "洛宁县",
		},
		"410329": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "伊川县",
		},
		"410381": {
			Province: "河南省",
			City:     "洛阳市",
			County:   "偃师市",
		},
		"410400": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "平顶山市",
		},
		"410402": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "新华区",
		},
		"410403": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "卫东区",
		},
		"410404": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "石龙区",
		},
		"410411": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "湛河区",
		},
		"410421": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "宝丰县",
		},
		"410422": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "叶县",
		},
		"410423": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "鲁山县",
		},
		"410425": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "郏县",
		},
		"410481": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "舞钢市",
		},
		"410482": {
			Province: "河南省",
			City:     "平顶山市",
			County:   "汝州市",
		},
		"410500": {
			Province: "河南省",
			City:     "安阳市",
			County:   "安阳市",
		},
		"410502": {
			Province: "河南省",
			City:     "安阳市",
			County:   "文峰区",
		},
		"410503": {
			Province: "河南省",
			City:     "安阳市",
			County:   "北关区",
		},
		"410505": {
			Province: "河南省",
			City:     "安阳市",
			County:   "殷都区",
		},
		"410506": {
			Province: "河南省",
			City:     "安阳市",
			County:   "龙安区",
		},
		"410522": {
			Province: "河南省",
			City:     "安阳市",
			County:   "安阳县",
		},
		"410523": {
			Province: "河南省",
			City:     "安阳市",
			County:   "汤阴县",
		},
		"410526": {
			Province: "河南省",
			City:     "安阳市",
			County:   "滑县",
		},
		"410527": {
			Province: "河南省",
			City:     "安阳市",
			County:   "内黄县",
		},
		"410581": {
			Province: "河南省",
			City:     "安阳市",
			County:   "林州市",
		},
		"410600": {
			Province: "河南省",
			City:     "鹤壁市",
			County:   "鹤壁市",
		},
		"410602": {
			Province: "河南省",
			City:     "鹤壁市",
			County:   "鹤山区",
		},
		"410603": {
			Province: "河南省",
			City:     "鹤壁市",
			County:   "山城区",
		},
		"410611": {
			Province: "河南省",
			City:     "鹤壁市",
			County:   "淇滨区",
		},
		"410621": {
			Province: "河南省",
			City:     "鹤壁市",
			County:   "浚县",
		},
		"410622": {
			Province: "河南省",
			City:     "鹤壁市",
			County:   "淇县",
		},
		"410700": {
			Province: "河南省",
			City:     "新乡市",
			County:   "新乡市",
		},
		"410702": {
			Province: "河南省",
			City:     "新乡市",
			County:   "红旗区",
		},
		"410703": {
			Province: "河南省",
			City:     "新乡市",
			County:   "卫滨区",
		},
		"410704": {
			Province: "河南省",
			City:     "新乡市",
			County:   "凤泉区",
		},
		"410711": {
			Province: "河南省",
			City:     "新乡市",
			County:   "牧野区",
		},
		"410721": {
			Province: "河南省",
			City:     "新乡市",
			County:   "新乡县",
		},
		"410724": {
			Province: "河南省",
			City:     "新乡市",
			County:   "获嘉县",
		},
		"410725": {
			Province: "河南省",
			City:     "新乡市",
			County:   "原阳县",
		},
		"410726": {
			Province: "河南省",
			City:     "新乡市",
			County:   "延津县",
		},
		"410727": {
			Province: "河南省",
			City:     "新乡市",
			County:   "封丘县",
		},
		"410781": {
			Province: "河南省",
			City:     "新乡市",
			County:   "卫辉市",
		},
		"410782": {
			Province: "河南省",
			City:     "新乡市",
			County:   "辉县市",
		},
		"410783": {
			Province: "河南省",
			City:     "新乡市",
			County:   "长垣市",
		},
		"410800": {
			Province: "河南省",
			City:     "焦作市",
			County:   "焦作市",
		},
		"410802": {
			Province: "河南省",
			City:     "焦作市",
			County:   "解放区",
		},
		"410803": {
			Province: "河南省",
			City:     "焦作市",
			County:   "中站区",
		},
		"410804": {
			Province: "河南省",
			City:     "焦作市",
			County:   "马村区",
		},
		"410811": {
			Province: "河南省",
			City:     "焦作市",
			County:   "山阳区",
		},
		"410821": {
			Province: "河南省",
			City:     "焦作市",
			County:   "修武县",
		},
		"410822": {
			Province: "河南省",
			City:     "焦作市",
			County:   "博爱县",
		},
		"410823": {
			Province: "河南省",
			City:     "焦作市",
			County:   "武陟县",
		},
		"410825": {
			Province: "河南省",
			City:     "焦作市",
			County:   "温县",
		},
		"410882": {
			Province: "河南省",
			City:     "焦作市",
			County:   "沁阳市",
		},
		"410883": {
			Province: "河南省",
			City:     "焦作市",
			County:   "孟州市",
		},
		"410900": {
			Province: "河南省",
			City:     "濮阳市",
			County:   "濮阳市",
		},
		"410902": {
			Province: "河南省",
			City:     "濮阳市",
			County:   "华龙区",
		},
		"410922": {
			Province: "河南省",
			City:     "濮阳市",
			County:   "清丰县",
		},
		"410923": {
			Province: "河南省",
			City:     "濮阳市",
			County:   "南乐县",
		},
		"410926": {
			Province: "河南省",
			City:     "濮阳市",
			County:   "范县",
		},
		"410927": {
			Province: "河南省",
			City:     "濮阳市",
			County:   "台前县",
		},
		"410928": {
			Province: "河南省",
			City:     "濮阳市",
			County:   "濮阳县",
		},
		"411000": {
			Province: "河南省",
			City:     "许昌市",
			County:   "许昌市",
		},
		"411002": {
			Province: "河南省",
			City:     "许昌市",
			County:   "魏都区",
		},
		"411003": {
			Province: "河南省",
			City:     "许昌市",
			County:   "建安区",
		},
		"411024": {
			Province: "河南省",
			City:     "许昌市",
			County:   "鄢陵县",
		},
		"411025": {
			Province: "河南省",
			City:     "许昌市",
			County:   "襄城县",
		},
		"411081": {
			Province: "河南省",
			City:     "许昌市",
			County:   "禹州市",
		},
		"411082": {
			Province: "河南省",
			City:     "许昌市",
			County:   "长葛市",
		},
		"411100": {
			Province: "河南省",
			City:     "漯河市",
			County:   "漯河市",
		},
		"411102": {
			Province: "河南省",
			City:     "漯河市",
			County:   "源汇区",
		},
		"411103": {
			Province: "河南省",
			City:     "漯河市",
			County:   "郾城区",
		},
		"411104": {
			Province: "河南省",
			City:     "漯河市",
			County:   "召陵区",
		},
		"411121": {
			Province: "河南省",
			City:     "漯河市",
			County:   "舞阳县",
		},
		"411122": {
			Province: "河南省",
			City:     "漯河市",
			County:   "临颍县",
		},
		"411200": {
			Province: "河南省",
			City:     "三门峡市",
			County:   "三门峡市",
		},
		"411202": {
			Province: "河南省",
			City:     "三门峡市",
			County:   "湖滨区",
		},
		"411203": {
			Province: "河南省",
			City:     "三门峡市",
			County:   "陕州区",
		},
		"411221": {
			Province: "河南省",
			City:     "三门峡市",
			County:   "渑池县",
		},
		"411224": {
			Province: "河南省",
			City:     "三门峡市",
			County:   "卢氏县",
		},
		"411281": {
			Province: "河南省",
			City:     "三门峡市",
			County:   "义马市",
		},
		"411282": {
			Province: "河南省",
			City:     "三门峡市",
			County:   "灵宝市",
		},
		"411300": {
			Province: "河南省",
			City:     "南阳市",
			County:   "南阳市",
		},
		"411302": {
			Province: "河南省",
			City:     "南阳市",
			County:   "宛城区",
		},
		"411303": {
			Province: "河南省",
			City:     "南阳市",
			County:   "卧龙区",
		},
		"411321": {
			Province: "河南省",
			City:     "南阳市",
			County:   "南召县",
		},
		"411322": {
			Province: "河南省",
			City:     "南阳市",
			County:   "方城县",
		},
		"411323": {
			Province: "河南省",
			City:     "南阳市",
			County:   "西峡县",
		},
		"411324": {
			Province: "河南省",
			City:     "南阳市",
			County:   "镇平县",
		},
		"411325": {
			Province: "河南省",
			City:     "南阳市",
			County:   "内乡县",
		},
		"411326": {
			Province: "河南省",
			City:     "南阳市",
			County:   "淅川县",
		},
		"411327": {
			Province: "河南省",
			City:     "南阳市",
			County:   "社旗县",
		},
		"411328": {
			Province: "河南省",
			City:     "南阳市",
			County:   "唐河县",
		},
		"411329": {
			Province: "河南省",
			City:     "南阳市",
			County:   "新野县",
		},
		"411330": {
			Province: "河南省",
			City:     "南阳市",
			County:   "桐柏县",
		},
		"411381": {
			Province: "河南省",
			City:     "南阳市",
			County:   "邓州市",
		},
		"411400": {
			Province: "河南省",
			City:     "商丘市",
			County:   "商丘市",
		},
		"411402": {
			Province: "河南省",
			City:     "商丘市",
			County:   "梁园区",
		},
		"411403": {
			Province: "河南省",
			City:     "商丘市",
			County:   "睢阳区",
		},
		"411421": {
			Province: "河南省",
			City:     "商丘市",
			County:   "民权县",
		},
		"411422": {
			Province: "河南省",
			City:     "商丘市",
			County:   "睢县",
		},
		"411423": {
			Province: "河南省",
			City:     "商丘市",
			County:   "宁陵县",
		},
		"411424": {
			Province: "河南省",
			City:     "商丘市",
			County:   "柘城县",
		},
		"411425": {
			Province: "河南省",
			City:     "商丘市",
			County:   "虞城县",
		},
		"411426": {
			Province: "河南省",
			City:     "商丘市",
			County:   "夏邑县",
		},
		"411481": {
			Province: "河南省",
			City:     "商丘市",
			County:   "永城市",
		},
		"411500": {
			Province: "河南省",
			City:     "信阳市",
			County:   "信阳市",
		},
		"411502": {
			Province: "河南省",
			City:     "信阳市",
			County:   "浉河区",
		},
		"411503": {
			Province: "河南省",
			City:     "信阳市",
			County:   "平桥区",
		},
		"411521": {
			Province: "河南省",
			City:     "信阳市",
			County:   "罗山县",
		},
		"411522": {
			Province: "河南省",
			City:     "信阳市",
			County:   "光山县",
		},
		"411523": {
			Province: "河南省",
			City:     "信阳市",
			County:   "新县",
		},
		"411524": {
			Province: "河南省",
			City:     "信阳市",
			County:   "商城县",
		},
		"411525": {
			Province: "河南省",
			City:     "信阳市",
			County:   "固始县",
		},
		"411526": {
			Province: "河南省",
			City:     "信阳市",
			County:   "潢川县",
		},
		"411527": {
			Province: "河南省",
			City:     "信阳市",
			County:   "淮滨县",
		},
		"411528": {
			Province: "河南省",
			City:     "信阳市",
			County:   "息县",
		},
		"411600": {
			Province: "河南省",
			City:     "周口市",
			County:   "周口市",
		},
		"411602": {
			Province: "河南省",
			City:     "周口市",
			County:   "川汇区",
		},
		"411603": {
			Province: "河南省",
			City:     "周口市",
			County:   "淮阳区",
		},
		"411621": {
			Province: "河南省",
			City:     "周口市",
			County:   "扶沟县",
		},
		"411622": {
			Province: "河南省",
			City:     "周口市",
			County:   "西华县",
		},
		"411623": {
			Province: "河南省",
			City:     "周口市",
			County:   "商水县",
		},
		"411624": {
			Province: "河南省",
			City:     "周口市",
			County:   "沈丘县",
		},
		"411625": {
			Province: "河南省",
			City:     "周口市",
			County:   "郸城县",
		},
		"411627": {
			Province: "河南省",
			City:     "周口市",
			County:   "太康县",
		},
		"411628": {
			Province: "河南省",
			City:     "周口市",
			County:   "鹿邑县",
		},
		"411681": {
			Province: "河南省",
			City:     "周口市",
			County:   "项城市",
		},
		"411700": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "驻马店市",
		},
		"411702": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "驿城区",
		},
		"411721": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "西平县",
		},
		"411722": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "上蔡县",
		},
		"411723": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "平舆县",
		},
		"411724": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "正阳县",
		},
		"411725": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "确山县",
		},
		"411726": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "泌阳县",
		},
		"411727": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "汝南县",
		},
		"411728": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "遂平县",
		},
		"411729": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "新蔡县",
		},
		"419001": {
			Province: "河南省",
			City:     "驻马店市",
			County:   "济源市",
		},
		"420000": {
			Province: "湖北省",
			City:     "湖北省",
			County:   "湖北省",
		},
		"420100": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "武汉市",
		},
		"420102": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "江岸区",
		},
		"420103": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "江汉区",
		},
		"420104": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "硚口区",
		},
		"420105": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "汉阳区",
		},
		"420106": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "武昌区",
		},
		"420107": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "青山区",
		},
		"420111": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "洪山区",
		},
		"420112": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "东西湖区",
		},
		"420113": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "汉南区",
		},
		"420114": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "蔡甸区",
		},
		"420115": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "江夏区",
		},
		"420116": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "黄陂区",
		},
		"420117": {
			Province: "湖北省",
			City:     "武汉市",
			County:   "新洲区",
		},
		"420200": {
			Province: "湖北省",
			City:     "黄石市",
			County:   "黄石市",
		},
		"420202": {
			Province: "湖北省",
			City:     "黄石市",
			County:   "黄石港区",
		},
		"420203": {
			Province: "湖北省",
			City:     "黄石市",
			County:   "西塞山区",
		},
		"420204": {
			Province: "湖北省",
			City:     "黄石市",
			County:   "下陆区",
		},
		"420205": {
			Province: "湖北省",
			City:     "黄石市",
			County:   "铁山区",
		},
		"420222": {
			Province: "湖北省",
			City:     "黄石市",
			County:   "阳新县",
		},
		"420281": {
			Province: "湖北省",
			City:     "黄石市",
			County:   "大冶市",
		},
		"420300": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "十堰市",
		},
		"420302": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "茅箭区",
		},
		"420303": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "张湾区",
		},
		"420304": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "郧阳区",
		},
		"420322": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "郧西县",
		},
		"420323": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "竹山县",
		},
		"420324": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "竹溪县",
		},
		"420325": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "房县",
		},
		"420381": {
			Province: "湖北省",
			City:     "十堰市",
			County:   "丹江口市",
		},
		"420500": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "宜昌市",
		},
		"420502": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "西陵区",
		},
		"420503": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "伍家岗区",
		},
		"420504": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "点军区",
		},
		"420505": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "猇亭区",
		},
		"420506": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "夷陵区",
		},
		"420525": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "远安县",
		},
		"420526": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "兴山县",
		},
		"420527": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "秭归县",
		},
		"420528": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "长阳土家族自治县",
		},
		"420529": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "五峰土家族自治县",
		},
		"420581": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "宜都市",
		},
		"420582": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "当阳市",
		},
		"420583": {
			Province: "湖北省",
			City:     "宜昌市",
			County:   "枝江市",
		},
		"420600": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "襄阳市",
		},
		"420602": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "襄城区",
		},
		"420606": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "樊城区",
		},
		"420607": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "襄州区",
		},
		"420624": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "南漳县",
		},
		"420625": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "谷城县",
		},
		"420626": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "保康县",
		},
		"420682": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "老河口市",
		},
		"420683": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "枣阳市",
		},
		"420684": {
			Province: "湖北省",
			City:     "襄阳市",
			County:   "宜城市",
		},
		"420700": {
			Province: "湖北省",
			City:     "鄂州市",
			County:   "鄂州市",
		},
		"420702": {
			Province: "湖北省",
			City:     "鄂州市",
			County:   "梁子湖区",
		},
		"420703": {
			Province: "湖北省",
			City:     "鄂州市",
			County:   "华容区",
		},
		"420704": {
			Province: "湖北省",
			City:     "鄂州市",
			County:   "鄂城区",
		},
		"420800": {
			Province: "湖北省",
			City:     "荆门市",
			County:   "荆门市",
		},
		"420802": {
			Province: "湖北省",
			City:     "荆门市",
			County:   "东宝区",
		},
		"420804": {
			Province: "湖北省",
			City:     "荆门市",
			County:   "掇刀区",
		},
		"420822": {
			Province: "湖北省",
			City:     "荆门市",
			County:   "沙洋县",
		},
		"420881": {
			Province: "湖北省",
			City:     "荆门市",
			County:   "钟祥市",
		},
		"420882": {
			Province: "湖北省",
			City:     "荆门市",
			County:   "京山市",
		},
		"420900": {
			Province: "湖北省",
			City:     "孝感市",
			County:   "孝感市",
		},
		"420902": {
			Province: "湖北省",
			City:     "孝感市",
			County:   "孝南区",
		},
		"420921": {
			Province: "湖北省",
			City:     "孝感市",
			County:   "孝昌县",
		},
		"420922": {
			Province: "湖北省",
			City:     "孝感市",
			County:   "大悟县",
		},
		"420923": {
			Province: "湖北省",
			City:     "孝感市",
			County:   "云梦县",
		},
		"420981": {
			Province: "湖北省",
			City:     "孝感市",
			County:   "应城市",
		},
		"420982": {
			Province: "湖北省",
			City:     "孝感市",
			County:   "安陆市",
		},
		"420984": {
			Province: "湖北省",
			City:     "孝感市",
			County:   "汉川市",
		},
		"421000": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "荆州市",
		},
		"421002": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "沙市区",
		},
		"421003": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "荆州区",
		},
		"421022": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "公安县",
		},
		"421024": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "江陵县",
		},
		"421081": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "石首市",
		},
		"421083": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "洪湖市",
		},
		"421087": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "松滋市",
		},
		"421088": {
			Province: "湖北省",
			City:     "荆州市",
			County:   "监利市",
		},
		"421100": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "黄冈市",
		},
		"421102": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "黄州区",
		},
		"421121": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "团风县",
		},
		"421122": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "红安县",
		},
		"421123": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "罗田县",
		},
		"421124": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "英山县",
		},
		"421125": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "浠水县",
		},
		"421126": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "蕲春县",
		},
		"421127": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "黄梅县",
		},
		"421181": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "麻城市",
		},
		"421182": {
			Province: "湖北省",
			City:     "黄冈市",
			County:   "武穴市",
		},
		"421200": {
			Province: "湖北省",
			City:     "咸宁市",
			County:   "咸宁市",
		},
		"421202": {
			Province: "湖北省",
			City:     "咸宁市",
			County:   "咸安区",
		},
		"421221": {
			Province: "湖北省",
			City:     "咸宁市",
			County:   "嘉鱼县",
		},
		"421222": {
			Province: "湖北省",
			City:     "咸宁市",
			County:   "通城县",
		},
		"421223": {
			Province: "湖北省",
			City:     "咸宁市",
			County:   "崇阳县",
		},
		"421224": {
			Province: "湖北省",
			City:     "咸宁市",
			County:   "通山县",
		},
		"421281": {
			Province: "湖北省",
			City:     "咸宁市",
			County:   "赤壁市",
		},
		"421300": {
			Province: "湖北省",
			City:     "随州市",
			County:   "随州市",
		},
		"421303": {
			Province: "湖北省",
			City:     "随州市",
			County:   "曾都区",
		},
		"421321": {
			Province: "湖北省",
			City:     "随州市",
			County:   "随县",
		},
		"421381": {
			Province: "湖北省",
			City:     "随州市",
			County:   "广水市",
		},
		"422800": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "恩施土家族苗族自治州",
		},
		"422801": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "恩施市",
		},
		"422802": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "利川市",
		},
		"422822": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "建始县",
		},
		"422823": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "巴东县",
		},
		"422825": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "宣恩县",
		},
		"422826": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "咸丰县",
		},
		"422827": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "来凤县",
		},
		"422828": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "鹤峰县",
		},
		"429004": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "仙桃市",
		},
		"429005": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "潜江市",
		},
		"429006": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "天门市",
		},
		"429021": {
			Province: "湖北省",
			City:     "恩施土家族苗族自治州",
			County:   "神农架林区",
		},
		"430000": {
			Province: "湖南省",
			City:     "湖南省",
			County:   "湖南省",
		},
		"430100": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "长沙市",
		},
		"430102": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "芙蓉区",
		},
		"430103": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "天心区",
		},
		"430104": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "岳麓区",
		},
		"430105": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "开福区",
		},
		"430111": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "雨花区",
		},
		"430112": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "望城区",
		},
		"430121": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "长沙县",
		},
		"430181": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "浏阳市",
		},
		"430182": {
			Province: "湖南省",
			City:     "长沙市",
			County:   "宁乡市",
		},
		"430200": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "株洲市",
		},
		"430202": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "荷塘区",
		},
		"430203": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "芦淞区",
		},
		"430204": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "石峰区",
		},
		"430211": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "天元区",
		},
		"430212": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "渌口区",
		},
		"430223": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "攸县",
		},
		"430224": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "茶陵县",
		},
		"430225": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "炎陵县",
		},
		"430281": {
			Province: "湖南省",
			City:     "株洲市",
			County:   "醴陵市",
		},
		"430300": {
			Province: "湖南省",
			City:     "湘潭市",
			County:   "湘潭市",
		},
		"430302": {
			Province: "湖南省",
			City:     "湘潭市",
			County:   "雨湖区",
		},
		"430304": {
			Province: "湖南省",
			City:     "湘潭市",
			County:   "岳塘区",
		},
		"430321": {
			Province: "湖南省",
			City:     "湘潭市",
			County:   "湘潭县",
		},
		"430381": {
			Province: "湖南省",
			City:     "湘潭市",
			County:   "湘乡市",
		},
		"430382": {
			Province: "湖南省",
			City:     "湘潭市",
			County:   "韶山市",
		},
		"430400": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "衡阳市",
		},
		"430405": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "珠晖区",
		},
		"430406": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "雁峰区",
		},
		"430407": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "石鼓区",
		},
		"430408": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "蒸湘区",
		},
		"430412": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "南岳区",
		},
		"430421": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "衡阳县",
		},
		"430422": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "衡南县",
		},
		"430423": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "衡山县",
		},
		"430424": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "衡东县",
		},
		"430426": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "祁东县",
		},
		"430481": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "耒阳市",
		},
		"430482": {
			Province: "湖南省",
			City:     "衡阳市",
			County:   "常宁市",
		},
		"430500": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "邵阳市",
		},
		"430502": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "双清区",
		},
		"430503": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "大祥区",
		},
		"430511": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "北塔区",
		},
		"430522": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "新邵县",
		},
		"430523": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "邵阳县",
		},
		"430524": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "隆回县",
		},
		"430525": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "洞口县",
		},
		"430527": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "绥宁县",
		},
		"430528": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "新宁县",
		},
		"430529": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "城步苗族自治县",
		},
		"430581": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "武冈市",
		},
		"430582": {
			Province: "湖南省",
			City:     "邵阳市",
			County:   "邵东市",
		},
		"430600": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "岳阳市",
		},
		"430602": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "岳阳楼区",
		},
		"430603": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "云溪区",
		},
		"430611": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "君山区",
		},
		"430621": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "岳阳县",
		},
		"430623": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "华容县",
		},
		"430624": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "湘阴县",
		},
		"430626": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "平江县",
		},
		"430681": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "汨罗市",
		},
		"430682": {
			Province: "湖南省",
			City:     "岳阳市",
			County:   "临湘市",
		},
		"430700": {
			Province: "湖南省",
			City:     "常德市",
			County:   "常德市",
		},
		"430702": {
			Province: "湖南省",
			City:     "常德市",
			County:   "武陵区",
		},
		"430703": {
			Province: "湖南省",
			City:     "常德市",
			County:   "鼎城区",
		},
		"430721": {
			Province: "湖南省",
			City:     "常德市",
			County:   "安乡县",
		},
		"430722": {
			Province: "湖南省",
			City:     "常德市",
			County:   "汉寿县",
		},
		"430723": {
			Province: "湖南省",
			City:     "常德市",
			County:   "澧县",
		},
		"430724": {
			Province: "湖南省",
			City:     "常德市",
			County:   "临澧县",
		},
		"430725": {
			Province: "湖南省",
			City:     "常德市",
			County:   "桃源县",
		},
		"430726": {
			Province: "湖南省",
			City:     "常德市",
			County:   "石门县",
		},
		"430781": {
			Province: "湖南省",
			City:     "常德市",
			County:   "津市市",
		},
		"430800": {
			Province: "湖南省",
			City:     "张家界市",
			County:   "张家界市",
		},
		"430802": {
			Province: "湖南省",
			City:     "张家界市",
			County:   "永定区",
		},
		"430811": {
			Province: "湖南省",
			City:     "张家界市",
			County:   "武陵源区",
		},
		"430821": {
			Province: "湖南省",
			City:     "张家界市",
			County:   "慈利县",
		},
		"430822": {
			Province: "湖南省",
			City:     "张家界市",
			County:   "桑植县",
		},
		"430900": {
			Province: "湖南省",
			City:     "益阳市",
			County:   "益阳市",
		},
		"430902": {
			Province: "湖南省",
			City:     "益阳市",
			County:   "资阳区",
		},
		"430903": {
			Province: "湖南省",
			City:     "益阳市",
			County:   "赫山区",
		},
		"430921": {
			Province: "湖南省",
			City:     "益阳市",
			County:   "南县",
		},
		"430922": {
			Province: "湖南省",
			City:     "益阳市",
			County:   "桃江县",
		},
		"430923": {
			Province: "湖南省",
			City:     "益阳市",
			County:   "安化县",
		},
		"430981": {
			Province: "湖南省",
			City:     "益阳市",
			County:   "沅江市",
		},
		"431000": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "郴州市",
		},
		"431002": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "北湖区",
		},
		"431003": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "苏仙区",
		},
		"431021": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "桂阳县",
		},
		"431022": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "宜章县",
		},
		"431023": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "永兴县",
		},
		"431024": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "嘉禾县",
		},
		"431025": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "临武县",
		},
		"431026": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "汝城县",
		},
		"431027": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "桂东县",
		},
		"431028": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "安仁县",
		},
		"431081": {
			Province: "湖南省",
			City:     "郴州市",
			County:   "资兴市",
		},
		"431100": {
			Province: "湖南省",
			City:     "永州市",
			County:   "永州市",
		},
		"431102": {
			Province: "湖南省",
			City:     "永州市",
			County:   "零陵区",
		},
		"431103": {
			Province: "湖南省",
			City:     "永州市",
			County:   "冷水滩区",
		},
		"431121": {
			Province: "湖南省",
			City:     "永州市",
			County:   "祁阳县",
		},
		"431122": {
			Province: "湖南省",
			City:     "永州市",
			County:   "东安县",
		},
		"431123": {
			Province: "湖南省",
			City:     "永州市",
			County:   "双牌县",
		},
		"431124": {
			Province: "湖南省",
			City:     "永州市",
			County:   "道县",
		},
		"431125": {
			Province: "湖南省",
			City:     "永州市",
			County:   "江永县",
		},
		"431126": {
			Province: "湖南省",
			City:     "永州市",
			County:   "宁远县",
		},
		"431127": {
			Province: "湖南省",
			City:     "永州市",
			County:   "蓝山县",
		},
		"431128": {
			Province: "湖南省",
			City:     "永州市",
			County:   "新田县",
		},
		"431129": {
			Province: "湖南省",
			City:     "永州市",
			County:   "江华瑶族自治县",
		},
		"431200": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "怀化市",
		},
		"431202": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "鹤城区",
		},
		"431221": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "中方县",
		},
		"431222": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "沅陵县",
		},
		"431223": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "辰溪县",
		},
		"431224": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "溆浦县",
		},
		"431225": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "会同县",
		},
		"431226": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "麻阳苗族自治县",
		},
		"431227": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "新晃侗族自治县",
		},
		"431228": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "芷江侗族自治县",
		},
		"431229": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "靖州苗族侗族自治县",
		},
		"431230": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "通道侗族自治县",
		},
		"431281": {
			Province: "湖南省",
			City:     "怀化市",
			County:   "洪江市",
		},
		"431300": {
			Province: "湖南省",
			City:     "娄底市",
			County:   "娄底市",
		},
		"431302": {
			Province: "湖南省",
			City:     "娄底市",
			County:   "娄星区",
		},
		"431321": {
			Province: "湖南省",
			City:     "娄底市",
			County:   "双峰县",
		},
		"431322": {
			Province: "湖南省",
			City:     "娄底市",
			County:   "新化县",
		},
		"431381": {
			Province: "湖南省",
			City:     "娄底市",
			County:   "冷水江市",
		},
		"431382": {
			Province: "湖南省",
			City:     "娄底市",
			County:   "涟源市",
		},
		"433100": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "湘西土家族苗族自治州",
		},
		"433101": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "吉首市",
		},
		"433122": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "泸溪县",
		},
		"433123": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "凤凰县",
		},
		"433124": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "花垣县",
		},
		"433125": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "保靖县",
		},
		"433126": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "古丈县",
		},
		"433127": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "永顺县",
		},
		"433130": {
			Province: "湖南省",
			City:     "湘西土家族苗族自治州",
			County:   "龙山县",
		},
		"440000": {
			Province: "广东省",
			City:     "广东省",
			County:   "广东省",
		},
		"440100": {
			Province: "广东省",
			City:     "广州市",
			County:   "广州市",
		},
		"440103": {
			Province: "广东省",
			City:     "广州市",
			County:   "荔湾区",
		},
		"440104": {
			Province: "广东省",
			City:     "广州市",
			County:   "越秀区",
		},
		"440105": {
			Province: "广东省",
			City:     "广州市",
			County:   "海珠区",
		},
		"440106": {
			Province: "广东省",
			City:     "广州市",
			County:   "天河区",
		},
		"440111": {
			Province: "广东省",
			City:     "广州市",
			County:   "白云区",
		},
		"440112": {
			Province: "广东省",
			City:     "广州市",
			County:   "黄埔区",
		},
		"440113": {
			Province: "广东省",
			City:     "广州市",
			County:   "番禺区",
		},
		"440114": {
			Province: "广东省",
			City:     "广州市",
			County:   "花都区",
		},
		"440115": {
			Province: "广东省",
			City:     "广州市",
			County:   "南沙区",
		},
		"440117": {
			Province: "广东省",
			City:     "广州市",
			County:   "从化区",
		},
		"440118": {
			Province: "广东省",
			City:     "广州市",
			County:   "增城区",
		},
		"440200": {
			Province: "广东省",
			City:     "韶关市",
			County:   "韶关市",
		},
		"440203": {
			Province: "广东省",
			City:     "韶关市",
			County:   "武江区",
		},
		"440204": {
			Province: "广东省",
			City:     "韶关市",
			County:   "浈江区",
		},
		"440205": {
			Province: "广东省",
			City:     "韶关市",
			County:   "曲江区",
		},
		"440222": {
			Province: "广东省",
			City:     "韶关市",
			County:   "始兴县",
		},
		"440224": {
			Province: "广东省",
			City:     "韶关市",
			County:   "仁化县",
		},
		"440229": {
			Province: "广东省",
			City:     "韶关市",
			County:   "翁源县",
		},
		"440232": {
			Province: "广东省",
			City:     "韶关市",
			County:   "乳源瑶族自治县",
		},
		"440233": {
			Province: "广东省",
			City:     "韶关市",
			County:   "新丰县",
		},
		"440281": {
			Province: "广东省",
			City:     "韶关市",
			County:   "乐昌市",
		},
		"440282": {
			Province: "广东省",
			City:     "韶关市",
			County:   "南雄市",
		},
		"440300": {
			Province: "广东省",
			City:     "深圳市",
			County:   "深圳市",
		},
		"440303": {
			Province: "广东省",
			City:     "深圳市",
			County:   "罗湖区",
		},
		"440304": {
			Province: "广东省",
			City:     "深圳市",
			County:   "福田区",
		},
		"440305": {
			Province: "广东省",
			City:     "深圳市",
			County:   "南山区",
		},
		"440306": {
			Province: "广东省",
			City:     "深圳市",
			County:   "宝安区",
		},
		"440307": {
			Province: "广东省",
			City:     "深圳市",
			County:   "龙岗区",
		},
		"440308": {
			Province: "广东省",
			City:     "深圳市",
			County:   "盐田区",
		},
		"440309": {
			Province: "广东省",
			City:     "深圳市",
			County:   "龙华区",
		},
		"440310": {
			Province: "广东省",
			City:     "深圳市",
			County:   "坪山区",
		},
		"440311": {
			Province: "广东省",
			City:     "深圳市",
			County:   "光明区",
		},
		"440400": {
			Province: "广东省",
			City:     "珠海市",
			County:   "珠海市",
		},
		"440402": {
			Province: "广东省",
			City:     "珠海市",
			County:   "香洲区",
		},
		"440403": {
			Province: "广东省",
			City:     "珠海市",
			County:   "斗门区",
		},
		"440404": {
			Province: "广东省",
			City:     "珠海市",
			County:   "金湾区",
		},
		"440500": {
			Province: "广东省",
			City:     "汕头市",
			County:   "汕头市",
		},
		"440507": {
			Province: "广东省",
			City:     "汕头市",
			County:   "龙湖区",
		},
		"440511": {
			Province: "广东省",
			City:     "汕头市",
			County:   "金平区",
		},
		"440512": {
			Province: "广东省",
			City:     "汕头市",
			County:   "濠江区",
		},
		"440513": {
			Province: "广东省",
			City:     "汕头市",
			County:   "潮阳区",
		},
		"440514": {
			Province: "广东省",
			City:     "汕头市",
			County:   "潮南区",
		},
		"440515": {
			Province: "广东省",
			City:     "汕头市",
			County:   "澄海区",
		},
		"440523": {
			Province: "广东省",
			City:     "汕头市",
			County:   "南澳县",
		},
		"440600": {
			Province: "广东省",
			City:     "佛山市",
			County:   "佛山市",
		},
		"440604": {
			Province: "广东省",
			City:     "佛山市",
			County:   "禅城区",
		},
		"440605": {
			Province: "广东省",
			City:     "佛山市",
			County:   "南海区",
		},
		"440606": {
			Province: "广东省",
			City:     "佛山市",
			County:   "顺德区",
		},
		"440607": {
			Province: "广东省",
			City:     "佛山市",
			County:   "三水区",
		},
		"440608": {
			Province: "广东省",
			City:     "佛山市",
			County:   "高明区",
		},
		"440700": {
			Province: "广东省",
			City:     "江门市",
			County:   "江门市",
		},
		"440703": {
			Province: "广东省",
			City:     "江门市",
			County:   "蓬江区",
		},
		"440704": {
			Province: "广东省",
			City:     "江门市",
			County:   "江海区",
		},
		"440705": {
			Province: "广东省",
			City:     "江门市",
			County:   "新会区",
		},
		"440781": {
			Province: "广东省",
			City:     "江门市",
			County:   "台山市",
		},
		"440783": {
			Province: "广东省",
			City:     "江门市",
			County:   "开平市",
		},
		"440784": {
			Province: "广东省",
			City:     "江门市",
			County:   "鹤山市",
		},
		"440785": {
			Province: "广东省",
			City:     "江门市",
			County:   "恩平市",
		},
		"440800": {
			Province: "广东省",
			City:     "湛江市",
			County:   "湛江市",
		},
		"440802": {
			Province: "广东省",
			City:     "湛江市",
			County:   "赤坎区",
		},
		"440803": {
			Province: "广东省",
			City:     "湛江市",
			County:   "霞山区",
		},
		"440804": {
			Province: "广东省",
			City:     "湛江市",
			County:   "坡头区",
		},
		"440811": {
			Province: "广东省",
			City:     "湛江市",
			County:   "麻章区",
		},
		"440823": {
			Province: "广东省",
			City:     "湛江市",
			County:   "遂溪县",
		},
		"440825": {
			Province: "广东省",
			City:     "湛江市",
			County:   "徐闻县",
		},
		"440881": {
			Province: "广东省",
			City:     "湛江市",
			County:   "廉江市",
		},
		"440882": {
			Province: "广东省",
			City:     "湛江市",
			County:   "雷州市",
		},
		"440883": {
			Province: "广东省",
			City:     "湛江市",
			County:   "吴川市",
		},
		"440900": {
			Province: "广东省",
			City:     "茂名市",
			County:   "茂名市",
		},
		"440902": {
			Province: "广东省",
			City:     "茂名市",
			County:   "茂南区",
		},
		"440904": {
			Province: "广东省",
			City:     "茂名市",
			County:   "电白区",
		},
		"440981": {
			Province: "广东省",
			City:     "茂名市",
			County:   "高州市",
		},
		"440982": {
			Province: "广东省",
			City:     "茂名市",
			County:   "化州市",
		},
		"440983": {
			Province: "广东省",
			City:     "茂名市",
			County:   "信宜市",
		},
		"441200": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "肇庆市",
		},
		"441202": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "端州区",
		},
		"441203": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "鼎湖区",
		},
		"441204": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "高要区",
		},
		"441223": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "广宁县",
		},
		"441224": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "怀集县",
		},
		"441225": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "封开县",
		},
		"441226": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "德庆县",
		},
		"441284": {
			Province: "广东省",
			City:     "肇庆市",
			County:   "四会市",
		},
		"441300": {
			Province: "广东省",
			City:     "惠州市",
			County:   "惠州市",
		},
		"441302": {
			Province: "广东省",
			City:     "惠州市",
			County:   "惠城区",
		},
		"441303": {
			Province: "广东省",
			City:     "惠州市",
			County:   "惠阳区",
		},
		"441322": {
			Province: "广东省",
			City:     "惠州市",
			County:   "博罗县",
		},
		"441323": {
			Province: "广东省",
			City:     "惠州市",
			County:   "惠东县",
		},
		"441324": {
			Province: "广东省",
			City:     "惠州市",
			County:   "龙门县",
		},
		"441400": {
			Province: "广东省",
			City:     "梅州市",
			County:   "梅州市",
		},
		"441402": {
			Province: "广东省",
			City:     "梅州市",
			County:   "梅江区",
		},
		"441403": {
			Province: "广东省",
			City:     "梅州市",
			County:   "梅县区",
		},
		"441422": {
			Province: "广东省",
			City:     "梅州市",
			County:   "大埔县",
		},
		"441423": {
			Province: "广东省",
			City:     "梅州市",
			County:   "丰顺县",
		},
		"441424": {
			Province: "广东省",
			City:     "梅州市",
			County:   "五华县",
		},
		"441426": {
			Province: "广东省",
			City:     "梅州市",
			County:   "平远县",
		},
		"441427": {
			Province: "广东省",
			City:     "梅州市",
			County:   "蕉岭县",
		},
		"441481": {
			Province: "广东省",
			City:     "梅州市",
			County:   "兴宁市",
		},
		"441500": {
			Province: "广东省",
			City:     "汕尾市",
			County:   "汕尾市",
		},
		"441502": {
			Province: "广东省",
			City:     "汕尾市",
			County:   "城区",
		},
		"441521": {
			Province: "广东省",
			City:     "汕尾市",
			County:   "海丰县",
		},
		"441523": {
			Province: "广东省",
			City:     "汕尾市",
			County:   "陆河县",
		},
		"441581": {
			Province: "广东省",
			City:     "汕尾市",
			County:   "陆丰市",
		},
		"441600": {
			Province: "广东省",
			City:     "河源市",
			County:   "河源市",
		},
		"441602": {
			Province: "广东省",
			City:     "河源市",
			County:   "源城区",
		},
		"441621": {
			Province: "广东省",
			City:     "河源市",
			County:   "紫金县",
		},
		"441622": {
			Province: "广东省",
			City:     "河源市",
			County:   "龙川县",
		},
		"441623": {
			Province: "广东省",
			City:     "河源市",
			County:   "连平县",
		},
		"441624": {
			Province: "广东省",
			City:     "河源市",
			County:   "和平县",
		},
		"441625": {
			Province: "广东省",
			City:     "河源市",
			County:   "东源县",
		},
		"441700": {
			Province: "广东省",
			City:     "阳江市",
			County:   "阳江市",
		},
		"441702": {
			Province: "广东省",
			City:     "阳江市",
			County:   "江城区",
		},
		"441704": {
			Province: "广东省",
			City:     "阳江市",
			County:   "阳东区",
		},
		"441721": {
			Province: "广东省",
			City:     "阳江市",
			County:   "阳西县",
		},
		"441781": {
			Province: "广东省",
			City:     "阳江市",
			County:   "阳春市",
		},
		"441800": {
			Province: "广东省",
			City:     "清远市",
			County:   "清远市",
		},
		"441802": {
			Province: "广东省",
			City:     "清远市",
			County:   "清城区",
		},
		"441803": {
			Province: "广东省",
			City:     "清远市",
			County:   "清新区",
		},
		"441821": {
			Province: "广东省",
			City:     "清远市",
			County:   "佛冈县",
		},
		"441823": {
			Province: "广东省",
			City:     "清远市",
			County:   "阳山县",
		},
		"441825": {
			Province: "广东省",
			City:     "清远市",
			County:   "连山壮族瑶族自治县",
		},
		"441826": {
			Province: "广东省",
			City:     "清远市",
			County:   "连南瑶族自治县",
		},
		"441881": {
			Province: "广东省",
			City:     "清远市",
			County:   "英德市",
		},
		"441882": {
			Province: "广东省",
			City:     "清远市",
			County:   "连州市",
		},
		"441900": {
			Province: "广东省",
			City:     "东莞市",
			County:   "东莞市",
		},
		"442000": {
			Province: "广东省",
			City:     "中山市",
			County:   "中山市",
		},
		"445100": {
			Province: "广东省",
			City:     "潮州市",
			County:   "潮州市",
		},
		"445102": {
			Province: "广东省",
			City:     "潮州市",
			County:   "湘桥区",
		},
		"445103": {
			Province: "广东省",
			City:     "潮州市",
			County:   "潮安区",
		},
		"445122": {
			Province: "广东省",
			City:     "潮州市",
			County:   "饶平县",
		},
		"445200": {
			Province: "广东省",
			City:     "揭阳市",
			County:   "揭阳市",
		},
		"445202": {
			Province: "广东省",
			City:     "揭阳市",
			County:   "榕城区",
		},
		"445203": {
			Province: "广东省",
			City:     "揭阳市",
			County:   "揭东区",
		},
		"445222": {
			Province: "广东省",
			City:     "揭阳市",
			County:   "揭西县",
		},
		"445224": {
			Province: "广东省",
			City:     "揭阳市",
			County:   "惠来县",
		},
		"445281": {
			Province: "广东省",
			City:     "揭阳市",
			County:   "普宁市",
		},
		"445300": {
			Province: "广东省",
			City:     "云浮市",
			County:   "云浮市",
		},
		"445302": {
			Province: "广东省",
			City:     "云浮市",
			County:   "云城区",
		},
		"445303": {
			Province: "广东省",
			City:     "云浮市",
			County:   "云安区",
		},
		"445321": {
			Province: "广东省",
			City:     "云浮市",
			County:   "新兴县",
		},
		"445322": {
			Province: "广东省",
			City:     "云浮市",
			County:   "郁南县",
		},
		"445381": {
			Province: "广东省",
			City:     "云浮市",
			County:   "罗定市",
		},
		"450000": {
			Province: "广西壮族自治区",
			City:     "广西壮族自治区",
			County:   "广西壮族自治区",
		},
		"450100": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "南宁市",
		},
		"450102": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "兴宁区",
		},
		"450103": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "青秀区",
		},
		"450105": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "江南区",
		},
		"450107": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "西乡塘区",
		},
		"450108": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "良庆区",
		},
		"450109": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "邕宁区",
		},
		"450110": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "武鸣区",
		},
		"450123": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "隆安县",
		},
		"450124": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "马山县",
		},
		"450125": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "上林县",
		},
		"450126": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "宾阳县",
		},
		"450127": {
			Province: "广西壮族自治区",
			City:     "南宁市",
			County:   "横县",
		},
		"450200": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "柳州市",
		},
		"450202": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "城中区",
		},
		"450203": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "鱼峰区",
		},
		"450204": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "柳南区",
		},
		"450205": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "柳北区",
		},
		"450206": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "柳江区",
		},
		"450222": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "柳城县",
		},
		"450223": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "鹿寨县",
		},
		"450224": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "融安县",
		},
		"450225": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "融水苗族自治县",
		},
		"450226": {
			Province: "广西壮族自治区",
			City:     "柳州市",
			County:   "三江侗族自治县",
		},
		"450300": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "桂林市",
		},
		"450302": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "秀峰区",
		},
		"450303": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "叠彩区",
		},
		"450304": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "象山区",
		},
		"450305": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "七星区",
		},
		"450311": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "雁山区",
		},
		"450312": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "临桂区",
		},
		"450321": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "阳朔县",
		},
		"450323": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "灵川县",
		},
		"450324": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "全州县",
		},
		"450325": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "兴安县",
		},
		"450326": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "永福县",
		},
		"450327": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "灌阳县",
		},
		"450328": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "龙胜各族自治县",
		},
		"450329": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "资源县",
		},
		"450330": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "平乐县",
		},
		"450332": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "恭城瑶族自治县",
		},
		"450381": {
			Province: "广西壮族自治区",
			City:     "桂林市",
			County:   "荔浦市",
		},
		"450400": {
			Province: "广西壮族自治区",
			City:     "梧州市",
			County:   "梧州市",
		},
		"450403": {
			Province: "广西壮族自治区",
			City:     "梧州市",
			County:   "万秀区",
		},
		"450405": {
			Province: "广西壮族自治区",
			City:     "梧州市",
			County:   "长洲区",
		},
		"450406": {
			Province: "广西壮族自治区",
			City:     "梧州市",
			County:   "龙圩区",
		},
		"450421": {
			Province: "广西壮族自治区",
			City:     "梧州市",
			County:   "苍梧县",
		},
		"450422": {
			Province: "广西壮族自治区",
			City:     "梧州市",
			County:   "藤县",
		},
		"450423": {
			Province: "广西壮族自治区",
			City:     "梧州市",
			County:   "蒙山县",
		},
		"450481": {
			Province: "广西壮族自治区",
			City:     "梧州市",
			County:   "岑溪市",
		},
		"450500": {
			Province: "广西壮族自治区",
			City:     "北海市",
			County:   "北海市",
		},
		"450502": {
			Province: "广西壮族自治区",
			City:     "北海市",
			County:   "海城区",
		},
		"450503": {
			Province: "广西壮族自治区",
			City:     "北海市",
			County:   "银海区",
		},
		"450512": {
			Province: "广西壮族自治区",
			City:     "北海市",
			County:   "铁山港区",
		},
		"450521": {
			Province: "广西壮族自治区",
			City:     "北海市",
			County:   "合浦县",
		},
		"450600": {
			Province: "广西壮族自治区",
			City:     "防城港市",
			County:   "防城港市",
		},
		"450602": {
			Province: "广西壮族自治区",
			City:     "防城港市",
			County:   "港口区",
		},
		"450603": {
			Province: "广西壮族自治区",
			City:     "防城港市",
			County:   "防城区",
		},
		"450621": {
			Province: "广西壮族自治区",
			City:     "防城港市",
			County:   "上思县",
		},
		"450681": {
			Province: "广西壮族自治区",
			City:     "防城港市",
			County:   "东兴市",
		},
		"450700": {
			Province: "广西壮族自治区",
			City:     "钦州市",
			County:   "钦州市",
		},
		"450702": {
			Province: "广西壮族自治区",
			City:     "钦州市",
			County:   "钦南区",
		},
		"450703": {
			Province: "广西壮族自治区",
			City:     "钦州市",
			County:   "钦北区",
		},
		"450721": {
			Province: "广西壮族自治区",
			City:     "钦州市",
			County:   "灵山县",
		},
		"450722": {
			Province: "广西壮族自治区",
			City:     "钦州市",
			County:   "浦北县",
		},
		"450800": {
			Province: "广西壮族自治区",
			City:     "贵港市",
			County:   "贵港市",
		},
		"450802": {
			Province: "广西壮族自治区",
			City:     "贵港市",
			County:   "港北区",
		},
		"450803": {
			Province: "广西壮族自治区",
			City:     "贵港市",
			County:   "港南区",
		},
		"450804": {
			Province: "广西壮族自治区",
			City:     "贵港市",
			County:   "覃塘区",
		},
		"450821": {
			Province: "广西壮族自治区",
			City:     "贵港市",
			County:   "平南县",
		},
		"450881": {
			Province: "广西壮族自治区",
			City:     "贵港市",
			County:   "桂平市",
		},
		"450900": {
			Province: "广西壮族自治区",
			City:     "玉林市",
			County:   "玉林市",
		},
		"450902": {
			Province: "广西壮族自治区",
			City:     "玉林市",
			County:   "玉州区",
		},
		"450903": {
			Province: "广西壮族自治区",
			City:     "玉林市",
			County:   "福绵区",
		},
		"450921": {
			Province: "广西壮族自治区",
			City:     "玉林市",
			County:   "容县",
		},
		"450922": {
			Province: "广西壮族自治区",
			City:     "玉林市",
			County:   "陆川县",
		},
		"450923": {
			Province: "广西壮族自治区",
			City:     "玉林市",
			County:   "博白县",
		},
		"450924": {
			Province: "广西壮族自治区",
			City:     "玉林市",
			County:   "兴业县",
		},
		"450981": {
			Province: "广西壮族自治区",
			City:     "玉林市",
			County:   "北流市",
		},
		"451000": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "百色市",
		},
		"451002": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "右江区",
		},
		"451003": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "田阳区",
		},
		"451022": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "田东县",
		},
		"451024": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "德保县",
		},
		"451026": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "那坡县",
		},
		"451027": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "凌云县",
		},
		"451028": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "乐业县",
		},
		"451029": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "田林县",
		},
		"451030": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "西林县",
		},
		"451031": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "隆林各族自治县",
		},
		"451081": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "靖西市",
		},
		"451082": {
			Province: "广西壮族自治区",
			City:     "百色市",
			County:   "平果市",
		},
		"451100": {
			Province: "广西壮族自治区",
			City:     "贺州市",
			County:   "贺州市",
		},
		"451102": {
			Province: "广西壮族自治区",
			City:     "贺州市",
			County:   "八步区",
		},
		"451103": {
			Province: "广西壮族自治区",
			City:     "贺州市",
			County:   "平桂区",
		},
		"451121": {
			Province: "广西壮族自治区",
			City:     "贺州市",
			County:   "昭平县",
		},
		"451122": {
			Province: "广西壮族自治区",
			City:     "贺州市",
			County:   "钟山县",
		},
		"451123": {
			Province: "广西壮族自治区",
			City:     "贺州市",
			County:   "富川瑶族自治县",
		},
		"451200": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "河池市",
		},
		"451202": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "金城江区",
		},
		"451203": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "宜州区",
		},
		"451221": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "南丹县",
		},
		"451222": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "天峨县",
		},
		"451223": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "凤山县",
		},
		"451224": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "东兰县",
		},
		"451225": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "罗城仫佬族自治县",
		},
		"451226": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "环江毛南族自治县",
		},
		"451227": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "巴马瑶族自治县",
		},
		"451228": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "都安瑶族自治县",
		},
		"451229": {
			Province: "广西壮族自治区",
			City:     "河池市",
			County:   "大化瑶族自治县",
		},
		"451300": {
			Province: "广西壮族自治区",
			City:     "来宾市",
			County:   "来宾市",
		},
		"451302": {
			Province: "广西壮族自治区",
			City:     "来宾市",
			County:   "兴宾区",
		},
		"451321": {
			Province: "广西壮族自治区",
			City:     "来宾市",
			County:   "忻城县",
		},
		"451322": {
			Province: "广西壮族自治区",
			City:     "来宾市",
			County:   "象州县",
		},
		"451323": {
			Province: "广西壮族自治区",
			City:     "来宾市",
			County:   "武宣县",
		},
		"451324": {
			Province: "广西壮族自治区",
			City:     "来宾市",
			County:   "金秀瑶族自治县",
		},
		"451381": {
			Province: "广西壮族自治区",
			City:     "来宾市",
			County:   "合山市",
		},
		"451400": {
			Province: "广西壮族自治区",
			City:     "崇左市",
			County:   "崇左市",
		},
		"451402": {
			Province: "广西壮族自治区",
			City:     "崇左市",
			County:   "江州区",
		},
		"451421": {
			Province: "广西壮族自治区",
			City:     "崇左市",
			County:   "扶绥县",
		},
		"451422": {
			Province: "广西壮族自治区",
			City:     "崇左市",
			County:   "宁明县",
		},
		"451423": {
			Province: "广西壮族自治区",
			City:     "崇左市",
			County:   "龙州县",
		},
		"451424": {
			Province: "广西壮族自治区",
			City:     "崇左市",
			County:   "大新县",
		},
		"451425": {
			Province: "广西壮族自治区",
			City:     "崇左市",
			County:   "天等县",
		},
		"451481": {
			Province: "广西壮族自治区",
			City:     "崇左市",
			County:   "凭祥市",
		},
		"460000": {
			Province: "海南省",
			City:     "海南省",
			County:   "海南省",
		},
		"460100": {
			Province: "海南省",
			City:     "海口市",
			County:   "海口市",
		},
		"460105": {
			Province: "海南省",
			City:     "海口市",
			County:   "秀英区",
		},
		"460106": {
			Province: "海南省",
			City:     "海口市",
			County:   "龙华区",
		},
		"460107": {
			Province: "海南省",
			City:     "海口市",
			County:   "琼山区",
		},
		"460108": {
			Province: "海南省",
			City:     "海口市",
			County:   "美兰区",
		},
		"460200": {
			Province: "海南省",
			City:     "三亚市",
			County:   "三亚市",
		},
		"460202": {
			Province: "海南省",
			City:     "三亚市",
			County:   "海棠区",
		},
		"460203": {
			Province: "海南省",
			City:     "三亚市",
			County:   "吉阳区",
		},
		"460204": {
			Province: "海南省",
			City:     "三亚市",
			County:   "天涯区",
		},
		"460205": {
			Province: "海南省",
			City:     "三亚市",
			County:   "崖州区",
		},
		"460300": {
			Province: "海南省",
			City:     "三沙市",
			County:   "三沙市",
		},
		"460400": {
			Province: "海南省",
			City:     "儋州市",
			County:   "儋州市",
		},
		"469001": {
			Province: "海南省",
			City:     "儋州市",
			County:   "五指山市",
		},
		"469002": {
			Province: "海南省",
			City:     "儋州市",
			County:   "琼海市",
		},
		"469005": {
			Province: "海南省",
			City:     "儋州市",
			County:   "文昌市",
		},
		"469006": {
			Province: "海南省",
			City:     "儋州市",
			County:   "万宁市",
		},
		"469007": {
			Province: "海南省",
			City:     "儋州市",
			County:   "东方市",
		},
		"469021": {
			Province: "海南省",
			City:     "儋州市",
			County:   "定安县",
		},
		"469022": {
			Province: "海南省",
			City:     "儋州市",
			County:   "屯昌县",
		},
		"469023": {
			Province: "海南省",
			City:     "儋州市",
			County:   "澄迈县",
		},
		"469024": {
			Province: "海南省",
			City:     "儋州市",
			County:   "临高县",
		},
		"469025": {
			Province: "海南省",
			City:     "儋州市",
			County:   "白沙黎族自治县",
		},
		"469026": {
			Province: "海南省",
			City:     "儋州市",
			County:   "昌江黎族自治县",
		},
		"469027": {
			Province: "海南省",
			City:     "儋州市",
			County:   "乐东黎族自治县",
		},
		"469028": {
			Province: "海南省",
			City:     "儋州市",
			County:   "陵水黎族自治县",
		},
		"469029": {
			Province: "海南省",
			City:     "儋州市",
			County:   "保亭黎族苗族自治县",
		},
		"469030": {
			Province: "海南省",
			City:     "儋州市",
			County:   "琼中黎族苗族自治县",
		},
		"500000": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "重庆市",
		},
		"500101": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "万州区",
		},
		"500102": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "涪陵区",
		},
		"500103": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "渝中区",
		},
		"500104": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "大渡口区",
		},
		"500105": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "江北区",
		},
		"500106": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "沙坪坝区",
		},
		"500107": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "九龙坡区",
		},
		"500108": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "南岸区",
		},
		"500109": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "北碚区",
		},
		"500110": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "綦江区",
		},
		"500111": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "大足区",
		},
		"500112": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "渝北区",
		},
		"500113": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "巴南区",
		},
		"500114": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "黔江区",
		},
		"500115": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "长寿区",
		},
		"500116": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "江津区",
		},
		"500117": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "合川区",
		},
		"500118": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "永川区",
		},
		"500119": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "南川区",
		},
		"500120": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "璧山区",
		},
		"500151": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "铜梁区",
		},
		"500152": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "潼南区",
		},
		"500153": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "荣昌区",
		},
		"500154": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "开州区",
		},
		"500155": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "梁平区",
		},
		"500156": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "武隆区",
		},
		"500229": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "城口县",
		},
		"500230": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "丰都县",
		},
		"500231": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "垫江县",
		},
		"500233": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "忠县",
		},
		"500235": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "云阳县",
		},
		"500236": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "奉节县",
		},
		"500237": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "巫山县",
		},
		"500238": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "巫溪县",
		},
		"500240": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "石柱土家族自治县",
		},
		"500241": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "秀山土家族苗族自治县",
		},
		"500242": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "酉阳土家族苗族自治县",
		},
		"500243": {
			Province: "重庆市",
			City:     "重庆市",
			County:   "彭水苗族土家族自治县",
		},
		"510000": {
			Province: "四川省",
			City:     "四川省",
			County:   "四川省",
		},
		"510100": {
			Province: "四川省",
			City:     "成都市",
			County:   "成都市",
		},
		"510104": {
			Province: "四川省",
			City:     "成都市",
			County:   "锦江区",
		},
		"510105": {
			Province: "四川省",
			City:     "成都市",
			County:   "青羊区",
		},
		"510106": {
			Province: "四川省",
			City:     "成都市",
			County:   "金牛区",
		},
		"510107": {
			Province: "四川省",
			City:     "成都市",
			County:   "武侯区",
		},
		"510108": {
			Province: "四川省",
			City:     "成都市",
			County:   "成华区",
		},
		"510112": {
			Province: "四川省",
			City:     "成都市",
			County:   "龙泉驿区",
		},
		"510113": {
			Province: "四川省",
			City:     "成都市",
			County:   "青白江区",
		},
		"510114": {
			Province: "四川省",
			City:     "成都市",
			County:   "新都区",
		},
		"510115": {
			Province: "四川省",
			City:     "成都市",
			County:   "温江区",
		},
		"510116": {
			Province: "四川省",
			City:     "成都市",
			County:   "双流区",
		},
		"510117": {
			Province: "四川省",
			City:     "成都市",
			County:   "郫都区",
		},
		"510118": {
			Province: "四川省",
			City:     "成都市",
			County:   "新津区",
		},
		"510121": {
			Province: "四川省",
			City:     "成都市",
			County:   "金堂县",
		},
		"510129": {
			Province: "四川省",
			City:     "成都市",
			County:   "大邑县",
		},
		"510131": {
			Province: "四川省",
			City:     "成都市",
			County:   "蒲江县",
		},
		"510181": {
			Province: "四川省",
			City:     "成都市",
			County:   "都江堰市",
		},
		"510182": {
			Province: "四川省",
			City:     "成都市",
			County:   "彭州市",
		},
		"510183": {
			Province: "四川省",
			City:     "成都市",
			County:   "邛崃市",
		},
		"510184": {
			Province: "四川省",
			City:     "成都市",
			County:   "崇州市",
		},
		"510185": {
			Province: "四川省",
			City:     "成都市",
			County:   "简阳市",
		},
		"510300": {
			Province: "四川省",
			City:     "自贡市",
			County:   "自贡市",
		},
		"510302": {
			Province: "四川省",
			City:     "自贡市",
			County:   "自流井区",
		},
		"510303": {
			Province: "四川省",
			City:     "自贡市",
			County:   "贡井区",
		},
		"510304": {
			Province: "四川省",
			City:     "自贡市",
			County:   "大安区",
		},
		"510311": {
			Province: "四川省",
			City:     "自贡市",
			County:   "沿滩区",
		},
		"510321": {
			Province: "四川省",
			City:     "自贡市",
			County:   "荣县",
		},
		"510322": {
			Province: "四川省",
			City:     "自贡市",
			County:   "富顺县",
		},
		"510400": {
			Province: "四川省",
			City:     "攀枝花市",
			County:   "攀枝花市",
		},
		"510402": {
			Province: "四川省",
			City:     "攀枝花市",
			County:   "东区",
		},
		"510403": {
			Province: "四川省",
			City:     "攀枝花市",
			County:   "西区",
		},
		"510411": {
			Province: "四川省",
			City:     "攀枝花市",
			County:   "仁和区",
		},
		"510421": {
			Province: "四川省",
			City:     "攀枝花市",
			County:   "米易县",
		},
		"510422": {
			Province: "四川省",
			City:     "攀枝花市",
			County:   "盐边县",
		},
		"510500": {
			Province: "四川省",
			City:     "泸州市",
			County:   "泸州市",
		},
		"510502": {
			Province: "四川省",
			City:     "泸州市",
			County:   "江阳区",
		},
		"510503": {
			Province: "四川省",
			City:     "泸州市",
			County:   "纳溪区",
		},
		"510504": {
			Province: "四川省",
			City:     "泸州市",
			County:   "龙马潭区",
		},
		"510521": {
			Province: "四川省",
			City:     "泸州市",
			County:   "泸县",
		},
		"510522": {
			Province: "四川省",
			City:     "泸州市",
			County:   "合江县",
		},
		"510524": {
			Province: "四川省",
			City:     "泸州市",
			County:   "叙永县",
		},
		"510525": {
			Province: "四川省",
			City:     "泸州市",
			County:   "古蔺县",
		},
		"510600": {
			Province: "四川省",
			City:     "德阳市",
			County:   "德阳市",
		},
		"510603": {
			Province: "四川省",
			City:     "德阳市",
			County:   "旌阳区",
		},
		"510604": {
			Province: "四川省",
			City:     "德阳市",
			County:   "罗江区",
		},
		"510623": {
			Province: "四川省",
			City:     "德阳市",
			County:   "中江县",
		},
		"510681": {
			Province: "四川省",
			City:     "德阳市",
			County:   "广汉市",
		},
		"510682": {
			Province: "四川省",
			City:     "德阳市",
			County:   "什邡市",
		},
		"510683": {
			Province: "四川省",
			City:     "德阳市",
			County:   "绵竹市",
		},
		"510700": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "绵阳市",
		},
		"510703": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "涪城区",
		},
		"510704": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "游仙区",
		},
		"510705": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "安州区",
		},
		"510722": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "三台县",
		},
		"510723": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "盐亭县",
		},
		"510725": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "梓潼县",
		},
		"510726": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "北川羌族自治县",
		},
		"510727": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "平武县",
		},
		"510781": {
			Province: "四川省",
			City:     "绵阳市",
			County:   "江油市",
		},
		"510800": {
			Province: "四川省",
			City:     "广元市",
			County:   "广元市",
		},
		"510802": {
			Province: "四川省",
			City:     "广元市",
			County:   "利州区",
		},
		"510811": {
			Province: "四川省",
			City:     "广元市",
			County:   "昭化区",
		},
		"510812": {
			Province: "四川省",
			City:     "广元市",
			County:   "朝天区",
		},
		"510821": {
			Province: "四川省",
			City:     "广元市",
			County:   "旺苍县",
		},
		"510822": {
			Province: "四川省",
			City:     "广元市",
			County:   "青川县",
		},
		"510823": {
			Province: "四川省",
			City:     "广元市",
			County:   "剑阁县",
		},
		"510824": {
			Province: "四川省",
			City:     "广元市",
			County:   "苍溪县",
		},
		"510900": {
			Province: "四川省",
			City:     "遂宁市",
			County:   "遂宁市",
		},
		"510903": {
			Province: "四川省",
			City:     "遂宁市",
			County:   "船山区",
		},
		"510904": {
			Province: "四川省",
			City:     "遂宁市",
			County:   "安居区",
		},
		"510921": {
			Province: "四川省",
			City:     "遂宁市",
			County:   "蓬溪县",
		},
		"510923": {
			Province: "四川省",
			City:     "遂宁市",
			County:   "大英县",
		},
		"510981": {
			Province: "四川省",
			City:     "遂宁市",
			County:   "射洪市",
		},
		"511000": {
			Province: "四川省",
			City:     "内江市",
			County:   "内江市",
		},
		"511002": {
			Province: "四川省",
			City:     "内江市",
			County:   "市中区",
		},
		"511011": {
			Province: "四川省",
			City:     "内江市",
			County:   "东兴区",
		},
		"511024": {
			Province: "四川省",
			City:     "内江市",
			County:   "威远县",
		},
		"511025": {
			Province: "四川省",
			City:     "内江市",
			County:   "资中县",
		},
		"511083": {
			Province: "四川省",
			City:     "内江市",
			County:   "隆昌市",
		},
		"511100": {
			Province: "四川省",
			City:     "乐山市",
			County:   "乐山市",
		},
		"511102": {
			Province: "四川省",
			City:     "乐山市",
			County:   "市中区",
		},
		"511111": {
			Province: "四川省",
			City:     "乐山市",
			County:   "沙湾区",
		},
		"511112": {
			Province: "四川省",
			City:     "乐山市",
			County:   "五通桥区",
		},
		"511113": {
			Province: "四川省",
			City:     "乐山市",
			County:   "金口河区",
		},
		"511123": {
			Province: "四川省",
			City:     "乐山市",
			County:   "犍为县",
		},
		"511124": {
			Province: "四川省",
			City:     "乐山市",
			County:   "井研县",
		},
		"511126": {
			Province: "四川省",
			City:     "乐山市",
			County:   "夹江县",
		},
		"511129": {
			Province: "四川省",
			City:     "乐山市",
			County:   "沐川县",
		},
		"511132": {
			Province: "四川省",
			City:     "乐山市",
			County:   "峨边彝族自治县",
		},
		"511133": {
			Province: "四川省",
			City:     "乐山市",
			County:   "马边彝族自治县",
		},
		"511181": {
			Province: "四川省",
			City:     "乐山市",
			County:   "峨眉山市",
		},
		"511300": {
			Province: "四川省",
			City:     "南充市",
			County:   "南充市",
		},
		"511302": {
			Province: "四川省",
			City:     "南充市",
			County:   "顺庆区",
		},
		"511303": {
			Province: "四川省",
			City:     "南充市",
			County:   "高坪区",
		},
		"511304": {
			Province: "四川省",
			City:     "南充市",
			County:   "嘉陵区",
		},
		"511321": {
			Province: "四川省",
			City:     "南充市",
			County:   "南部县",
		},
		"511322": {
			Province: "四川省",
			City:     "南充市",
			County:   "营山县",
		},
		"511323": {
			Province: "四川省",
			City:     "南充市",
			County:   "蓬安县",
		},
		"511324": {
			Province: "四川省",
			City:     "南充市",
			County:   "仪陇县",
		},
		"511325": {
			Province: "四川省",
			City:     "南充市",
			County:   "西充县",
		},
		"511381": {
			Province: "四川省",
			City:     "南充市",
			County:   "阆中市",
		},
		"511400": {
			Province: "四川省",
			City:     "眉山市",
			County:   "眉山市",
		},
		"511402": {
			Province: "四川省",
			City:     "眉山市",
			County:   "东坡区",
		},
		"511403": {
			Province: "四川省",
			City:     "眉山市",
			County:   "彭山区",
		},
		"511421": {
			Province: "四川省",
			City:     "眉山市",
			County:   "仁寿县",
		},
		"511423": {
			Province: "四川省",
			City:     "眉山市",
			County:   "洪雅县",
		},
		"511424": {
			Province: "四川省",
			City:     "眉山市",
			County:   "丹棱县",
		},
		"511425": {
			Province: "四川省",
			City:     "眉山市",
			County:   "青神县",
		},
		"511500": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "宜宾市",
		},
		"511502": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "翠屏区",
		},
		"511503": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "南溪区",
		},
		"511504": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "叙州区",
		},
		"511523": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "江安县",
		},
		"511524": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "长宁县",
		},
		"511525": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "高县",
		},
		"511526": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "珙县",
		},
		"511527": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "筠连县",
		},
		"511528": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "兴文县",
		},
		"511529": {
			Province: "四川省",
			City:     "宜宾市",
			County:   "屏山县",
		},
		"511600": {
			Province: "四川省",
			City:     "广安市",
			County:   "广安市",
		},
		"511602": {
			Province: "四川省",
			City:     "广安市",
			County:   "广安区",
		},
		"511603": {
			Province: "四川省",
			City:     "广安市",
			County:   "前锋区",
		},
		"511621": {
			Province: "四川省",
			City:     "广安市",
			County:   "岳池县",
		},
		"511622": {
			Province: "四川省",
			City:     "广安市",
			County:   "武胜县",
		},
		"511623": {
			Province: "四川省",
			City:     "广安市",
			County:   "邻水县",
		},
		"511681": {
			Province: "四川省",
			City:     "广安市",
			County:   "华蓥市",
		},
		"511700": {
			Province: "四川省",
			City:     "达州市",
			County:   "达州市",
		},
		"511702": {
			Province: "四川省",
			City:     "达州市",
			County:   "通川区",
		},
		"511703": {
			Province: "四川省",
			City:     "达州市",
			County:   "达川区",
		},
		"511722": {
			Province: "四川省",
			City:     "达州市",
			County:   "宣汉县",
		},
		"511723": {
			Province: "四川省",
			City:     "达州市",
			County:   "开江县",
		},
		"511724": {
			Province: "四川省",
			City:     "达州市",
			County:   "大竹县",
		},
		"511725": {
			Province: "四川省",
			City:     "达州市",
			County:   "渠县",
		},
		"511781": {
			Province: "四川省",
			City:     "达州市",
			County:   "万源市",
		},
		"511800": {
			Province: "四川省",
			City:     "雅安市",
			County:   "雅安市",
		},
		"511802": {
			Province: "四川省",
			City:     "雅安市",
			County:   "雨城区",
		},
		"511803": {
			Province: "四川省",
			City:     "雅安市",
			County:   "名山区",
		},
		"511822": {
			Province: "四川省",
			City:     "雅安市",
			County:   "荥经县",
		},
		"511823": {
			Province: "四川省",
			City:     "雅安市",
			County:   "汉源县",
		},
		"511824": {
			Province: "四川省",
			City:     "雅安市",
			County:   "石棉县",
		},
		"511825": {
			Province: "四川省",
			City:     "雅安市",
			County:   "天全县",
		},
		"511826": {
			Province: "四川省",
			City:     "雅安市",
			County:   "芦山县",
		},
		"511827": {
			Province: "四川省",
			City:     "雅安市",
			County:   "宝兴县",
		},
		"511900": {
			Province: "四川省",
			City:     "巴中市",
			County:   "巴中市",
		},
		"511902": {
			Province: "四川省",
			City:     "巴中市",
			County:   "巴州区",
		},
		"511903": {
			Province: "四川省",
			City:     "巴中市",
			County:   "恩阳区",
		},
		"511921": {
			Province: "四川省",
			City:     "巴中市",
			County:   "通江县",
		},
		"511922": {
			Province: "四川省",
			City:     "巴中市",
			County:   "南江县",
		},
		"511923": {
			Province: "四川省",
			City:     "巴中市",
			County:   "平昌县",
		},
		"512000": {
			Province: "四川省",
			City:     "资阳市",
			County:   "资阳市",
		},
		"512002": {
			Province: "四川省",
			City:     "资阳市",
			County:   "雁江区",
		},
		"512021": {
			Province: "四川省",
			City:     "资阳市",
			County:   "安岳县",
		},
		"512022": {
			Province: "四川省",
			City:     "资阳市",
			County:   "乐至县",
		},
		"513200": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "阿坝藏族羌族自治州",
		},
		"513201": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "马尔康市",
		},
		"513221": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "汶川县",
		},
		"513222": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "理县",
		},
		"513223": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "茂县",
		},
		"513224": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "松潘县",
		},
		"513225": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "九寨沟县",
		},
		"513226": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "金川县",
		},
		"513227": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "小金县",
		},
		"513228": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "黑水县",
		},
		"513230": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "壤塘县",
		},
		"513231": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "阿坝县",
		},
		"513232": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "若尔盖县",
		},
		"513233": {
			Province: "四川省",
			City:     "阿坝藏族羌族自治州",
			County:   "红原县",
		},
		"513300": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "甘孜藏族自治州",
		},
		"513301": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "康定市",
		},
		"513322": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "泸定县",
		},
		"513323": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "丹巴县",
		},
		"513324": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "九龙县",
		},
		"513325": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "雅江县",
		},
		"513326": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "道孚县",
		},
		"513327": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "炉霍县",
		},
		"513328": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "甘孜县",
		},
		"513329": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "新龙县",
		},
		"513330": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "德格县",
		},
		"513331": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "白玉县",
		},
		"513332": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "石渠县",
		},
		"513333": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "色达县",
		},
		"513334": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "理塘县",
		},
		"513335": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "巴塘县",
		},
		"513336": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "乡城县",
		},
		"513337": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "稻城县",
		},
		"513338": {
			Province: "四川省",
			City:     "甘孜藏族自治州",
			County:   "得荣县",
		},
		"513400": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "凉山彝族自治州",
		},
		"513401": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "西昌市",
		},
		"513422": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "木里藏族自治县",
		},
		"513423": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "盐源县",
		},
		"513424": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "德昌县",
		},
		"513425": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "会理县",
		},
		"513426": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "会东县",
		},
		"513427": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "宁南县",
		},
		"513428": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "普格县",
		},
		"513429": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "布拖县",
		},
		"513430": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "金阳县",
		},
		"513431": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "昭觉县",
		},
		"513432": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "喜德县",
		},
		"513433": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "冕宁县",
		},
		"513434": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "越西县",
		},
		"513435": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "甘洛县",
		},
		"513436": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "美姑县",
		},
		"513437": {
			Province: "四川省",
			City:     "凉山彝族自治州",
			County:   "雷波县",
		},
		"520000": {
			Province: "贵州省",
			City:     "贵州省",
			County:   "贵州省",
		},
		"520100": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "贵阳市",
		},
		"520102": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "南明区",
		},
		"520103": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "云岩区",
		},
		"520111": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "花溪区",
		},
		"520112": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "乌当区",
		},
		"520113": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "白云区",
		},
		"520115": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "观山湖区",
		},
		"520121": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "开阳县",
		},
		"520122": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "息烽县",
		},
		"520123": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "修文县",
		},
		"520181": {
			Province: "贵州省",
			City:     "贵阳市",
			County:   "清镇市",
		},
		"520200": {
			Province: "贵州省",
			City:     "六盘水市",
			County:   "六盘水市",
		},
		"520201": {
			Province: "贵州省",
			City:     "六盘水市",
			County:   "钟山区",
		},
		"520203": {
			Province: "贵州省",
			City:     "六盘水市",
			County:   "六枝特区",
		},
		"520204": {
			Province: "贵州省",
			City:     "六盘水市",
			County:   "水城区",
		},
		"520281": {
			Province: "贵州省",
			City:     "六盘水市",
			County:   "盘州市",
		},
		"520300": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "遵义市",
		},
		"520302": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "红花岗区",
		},
		"520303": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "汇川区",
		},
		"520304": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "播州区",
		},
		"520322": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "桐梓县",
		},
		"520323": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "绥阳县",
		},
		"520324": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "正安县",
		},
		"520325": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "道真仡佬族苗族自治县",
		},
		"520326": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "务川仡佬族苗族自治县",
		},
		"520327": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "凤冈县",
		},
		"520328": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "湄潭县",
		},
		"520329": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "余庆县",
		},
		"520330": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "习水县",
		},
		"520381": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "赤水市",
		},
		"520382": {
			Province: "贵州省",
			City:     "遵义市",
			County:   "仁怀市",
		},
		"520400": {
			Province: "贵州省",
			City:     "安顺市",
			County:   "安顺市",
		},
		"520402": {
			Province: "贵州省",
			City:     "安顺市",
			County:   "西秀区",
		},
		"520403": {
			Province: "贵州省",
			City:     "安顺市",
			County:   "平坝区",
		},
		"520422": {
			Province: "贵州省",
			City:     "安顺市",
			County:   "普定县",
		},
		"520423": {
			Province: "贵州省",
			City:     "安顺市",
			County:   "镇宁布依族苗族自治县",
		},
		"520424": {
			Province: "贵州省",
			City:     "安顺市",
			County:   "关岭布依族苗族自治县",
		},
		"520425": {
			Province: "贵州省",
			City:     "安顺市",
			County:   "紫云苗族布依族自治县",
		},
		"520500": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "毕节市",
		},
		"520502": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "七星关区",
		},
		"520521": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "大方县",
		},
		"520522": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "黔西县",
		},
		"520523": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "金沙县",
		},
		"520524": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "织金县",
		},
		"520525": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "纳雍县",
		},
		"520526": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "威宁彝族回族苗族自治县",
		},
		"520527": {
			Province: "贵州省",
			City:     "毕节市",
			County:   "赫章县",
		},
		"520600": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "铜仁市",
		},
		"520602": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "碧江区",
		},
		"520603": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "万山区",
		},
		"520621": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "江口县",
		},
		"520622": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "玉屏侗族自治县",
		},
		"520623": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "石阡县",
		},
		"520624": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "思南县",
		},
		"520625": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "印江土家族苗族自治县",
		},
		"520626": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "德江县",
		},
		"520627": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "沿河土家族自治县",
		},
		"520628": {
			Province: "贵州省",
			City:     "铜仁市",
			County:   "松桃苗族自治县",
		},
		"522300": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "黔西南布依族苗族自治州",
		},
		"522301": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "兴义市",
		},
		"522302": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "兴仁市",
		},
		"522323": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "普安县",
		},
		"522324": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "晴隆县",
		},
		"522325": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "贞丰县",
		},
		"522326": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "望谟县",
		},
		"522327": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "册亨县",
		},
		"522328": {
			Province: "贵州省",
			City:     "黔西南布依族苗族自治州",
			County:   "安龙县",
		},
		"522600": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "黔东南苗族侗族自治州",
		},
		"522601": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "凯里市",
		},
		"522622": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "黄平县",
		},
		"522623": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "施秉县",
		},
		"522624": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "三穗县",
		},
		"522625": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "镇远县",
		},
		"522626": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "岑巩县",
		},
		"522627": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "天柱县",
		},
		"522628": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "锦屏县",
		},
		"522629": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "剑河县",
		},
		"522630": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "台江县",
		},
		"522631": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "黎平县",
		},
		"522632": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "榕江县",
		},
		"522633": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "从江县",
		},
		"522634": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "雷山县",
		},
		"522635": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "麻江县",
		},
		"522636": {
			Province: "贵州省",
			City:     "黔东南苗族侗族自治州",
			County:   "丹寨县",
		},
		"522700": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "黔南布依族苗族自治州",
		},
		"522701": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "都匀市",
		},
		"522702": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "福泉市",
		},
		"522722": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "荔波县",
		},
		"522723": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "贵定县",
		},
		"522725": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "瓮安县",
		},
		"522726": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "独山县",
		},
		"522727": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "平塘县",
		},
		"522728": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "罗甸县",
		},
		"522729": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "长顺县",
		},
		"522730": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "龙里县",
		},
		"522731": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "惠水县",
		},
		"522732": {
			Province: "贵州省",
			City:     "黔南布依族苗族自治州",
			County:   "三都水族自治县",
		},
		"530000": {
			Province: "云南省",
			City:     "云南省",
			County:   "云南省",
		},
		"530100": {
			Province: "云南省",
			City:     "昆明市",
			County:   "昆明市",
		},
		"530102": {
			Province: "云南省",
			City:     "昆明市",
			County:   "五华区",
		},
		"530103": {
			Province: "云南省",
			City:     "昆明市",
			County:   "盘龙区",
		},
		"530111": {
			Province: "云南省",
			City:     "昆明市",
			County:   "官渡区",
		},
		"530112": {
			Province: "云南省",
			City:     "昆明市",
			County:   "西山区",
		},
		"530113": {
			Province: "云南省",
			City:     "昆明市",
			County:   "东川区",
		},
		"530114": {
			Province: "云南省",
			City:     "昆明市",
			County:   "呈贡区",
		},
		"530115": {
			Province: "云南省",
			City:     "昆明市",
			County:   "晋宁区",
		},
		"530124": {
			Province: "云南省",
			City:     "昆明市",
			County:   "富民县",
		},
		"530125": {
			Province: "云南省",
			City:     "昆明市",
			County:   "宜良县",
		},
		"530126": {
			Province: "云南省",
			City:     "昆明市",
			County:   "石林彝族自治县",
		},
		"530127": {
			Province: "云南省",
			City:     "昆明市",
			County:   "嵩明县",
		},
		"530128": {
			Province: "云南省",
			City:     "昆明市",
			County:   "禄劝彝族苗族自治县",
		},
		"530129": {
			Province: "云南省",
			City:     "昆明市",
			County:   "寻甸回族彝族自治县",
		},
		"530181": {
			Province: "云南省",
			City:     "昆明市",
			County:   "安宁市",
		},
		"530300": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "曲靖市",
		},
		"530302": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "麒麟区",
		},
		"530303": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "沾益区",
		},
		"530304": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "马龙区",
		},
		"530322": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "陆良县",
		},
		"530323": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "师宗县",
		},
		"530324": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "罗平县",
		},
		"530325": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "富源县",
		},
		"530326": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "会泽县",
		},
		"530381": {
			Province: "云南省",
			City:     "曲靖市",
			County:   "宣威市",
		},
		"530400": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "玉溪市",
		},
		"530402": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "红塔区",
		},
		"530403": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "江川区",
		},
		"530423": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "通海县",
		},
		"530424": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "华宁县",
		},
		"530425": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "易门县",
		},
		"530426": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "峨山彝族自治县",
		},
		"530427": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "新平彝族傣族自治县",
		},
		"530428": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "元江哈尼族彝族傣族自治县",
		},
		"530481": {
			Province: "云南省",
			City:     "玉溪市",
			County:   "澄江市",
		},
		"530500": {
			Province: "云南省",
			City:     "保山市",
			County:   "保山市",
		},
		"530502": {
			Province: "云南省",
			City:     "保山市",
			County:   "隆阳区",
		},
		"530521": {
			Province: "云南省",
			City:     "保山市",
			County:   "施甸县",
		},
		"530523": {
			Province: "云南省",
			City:     "保山市",
			County:   "龙陵县",
		},
		"530524": {
			Province: "云南省",
			City:     "保山市",
			County:   "昌宁县",
		},
		"530581": {
			Province: "云南省",
			City:     "保山市",
			County:   "腾冲市",
		},
		"530600": {
			Province: "云南省",
			City:     "昭通市",
			County:   "昭通市",
		},
		"530602": {
			Province: "云南省",
			City:     "昭通市",
			County:   "昭阳区",
		},
		"530621": {
			Province: "云南省",
			City:     "昭通市",
			County:   "鲁甸县",
		},
		"530622": {
			Province: "云南省",
			City:     "昭通市",
			County:   "巧家县",
		},
		"530623": {
			Province: "云南省",
			City:     "昭通市",
			County:   "盐津县",
		},
		"530624": {
			Province: "云南省",
			City:     "昭通市",
			County:   "大关县",
		},
		"530625": {
			Province: "云南省",
			City:     "昭通市",
			County:   "永善县",
		},
		"530626": {
			Province: "云南省",
			City:     "昭通市",
			County:   "绥江县",
		},
		"530627": {
			Province: "云南省",
			City:     "昭通市",
			County:   "镇雄县",
		},
		"530628": {
			Province: "云南省",
			City:     "昭通市",
			County:   "彝良县",
		},
		"530629": {
			Province: "云南省",
			City:     "昭通市",
			County:   "威信县",
		},
		"530681": {
			Province: "云南省",
			City:     "昭通市",
			County:   "水富市",
		},
		"530700": {
			Province: "云南省",
			City:     "丽江市",
			County:   "丽江市",
		},
		"530702": {
			Province: "云南省",
			City:     "丽江市",
			County:   "古城区",
		},
		"530721": {
			Province: "云南省",
			City:     "丽江市",
			County:   "玉龙纳西族自治县",
		},
		"530722": {
			Province: "云南省",
			City:     "丽江市",
			County:   "永胜县",
		},
		"530723": {
			Province: "云南省",
			City:     "丽江市",
			County:   "华坪县",
		},
		"530724": {
			Province: "云南省",
			City:     "丽江市",
			County:   "宁蒗彝族自治县",
		},
		"530800": {
			Province: "云南省",
			City:     "普洱市",
			County:   "普洱市",
		},
		"530802": {
			Province: "云南省",
			City:     "普洱市",
			County:   "思茅区",
		},
		"530821": {
			Province: "云南省",
			City:     "普洱市",
			County:   "宁洱哈尼族彝族自治县",
		},
		"530822": {
			Province: "云南省",
			City:     "普洱市",
			County:   "墨江哈尼族自治县",
		},
		"530823": {
			Province: "云南省",
			City:     "普洱市",
			County:   "景东彝族自治县",
		},
		"530824": {
			Province: "云南省",
			City:     "普洱市",
			County:   "景谷傣族彝族自治县",
		},
		"530825": {
			Province: "云南省",
			City:     "普洱市",
			County:   "镇沅彝族哈尼族拉祜族自治县",
		},
		"530826": {
			Province: "云南省",
			City:     "普洱市",
			County:   "江城哈尼族彝族自治县",
		},
		"530827": {
			Province: "云南省",
			City:     "普洱市",
			County:   "孟连傣族拉祜族佤族自治县",
		},
		"530828": {
			Province: "云南省",
			City:     "普洱市",
			County:   "澜沧拉祜族自治县",
		},
		"530829": {
			Province: "云南省",
			City:     "普洱市",
			County:   "西盟佤族自治县",
		},
		"530900": {
			Province: "云南省",
			City:     "临沧市",
			County:   "临沧市",
		},
		"530902": {
			Province: "云南省",
			City:     "临沧市",
			County:   "临翔区",
		},
		"530921": {
			Province: "云南省",
			City:     "临沧市",
			County:   "凤庆县",
		},
		"530922": {
			Province: "云南省",
			City:     "临沧市",
			County:   "云县",
		},
		"530923": {
			Province: "云南省",
			City:     "临沧市",
			County:   "永德县",
		},
		"530924": {
			Province: "云南省",
			City:     "临沧市",
			County:   "镇康县",
		},
		"530925": {
			Province: "云南省",
			City:     "临沧市",
			County:   "双江拉祜族佤族布朗族傣族自治县",
		},
		"530926": {
			Province: "云南省",
			City:     "临沧市",
			County:   "耿马傣族佤族自治县",
		},
		"530927": {
			Province: "云南省",
			City:     "临沧市",
			County:   "沧源佤族自治县",
		},
		"532300": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "楚雄彝族自治州",
		},
		"532301": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "楚雄市",
		},
		"532322": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "双柏县",
		},
		"532323": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "牟定县",
		},
		"532324": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "南华县",
		},
		"532325": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "姚安县",
		},
		"532326": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "大姚县",
		},
		"532327": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "永仁县",
		},
		"532328": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "元谋县",
		},
		"532329": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "武定县",
		},
		"532331": {
			Province: "云南省",
			City:     "楚雄彝族自治州",
			County:   "禄丰县",
		},
		"532500": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "红河哈尼族彝族自治州",
		},
		"532501": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "个旧市",
		},
		"532502": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "开远市",
		},
		"532503": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "蒙自市",
		},
		"532504": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "弥勒市",
		},
		"532523": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "屏边苗族自治县",
		},
		"532524": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "建水县",
		},
		"532525": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "石屏县",
		},
		"532527": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "泸西县",
		},
		"532528": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "元阳县",
		},
		"532529": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "红河县",
		},
		"532530": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "金平苗族瑶族傣族自治县",
		},
		"532531": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "绿春县",
		},
		"532532": {
			Province: "云南省",
			City:     "红河哈尼族彝族自治州",
			County:   "河口瑶族自治县",
		},
		"532600": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "文山壮族苗族自治州",
		},
		"532601": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "文山市",
		},
		"532622": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "砚山县",
		},
		"532623": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "西畴县",
		},
		"532624": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "麻栗坡县",
		},
		"532625": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "马关县",
		},
		"532626": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "丘北县",
		},
		"532627": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "广南县",
		},
		"532628": {
			Province: "云南省",
			City:     "文山壮族苗族自治州",
			County:   "富宁县",
		},
		"532800": {
			Province: "云南省",
			City:     "西双版纳傣族自治州",
			County:   "西双版纳傣族自治州",
		},
		"532801": {
			Province: "云南省",
			City:     "西双版纳傣族自治州",
			County:   "景洪市",
		},
		"532822": {
			Province: "云南省",
			City:     "西双版纳傣族自治州",
			County:   "勐海县",
		},
		"532823": {
			Province: "云南省",
			City:     "西双版纳傣族自治州",
			County:   "勐腊县",
		},
		"532900": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "大理白族自治州",
		},
		"532901": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "大理市",
		},
		"532922": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "漾濞彝族自治县",
		},
		"532923": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "祥云县",
		},
		"532924": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "宾川县",
		},
		"532925": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "弥渡县",
		},
		"532926": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "南涧彝族自治县",
		},
		"532927": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "巍山彝族回族自治县",
		},
		"532928": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "永平县",
		},
		"532929": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "云龙县",
		},
		"532930": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "洱源县",
		},
		"532931": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "剑川县",
		},
		"532932": {
			Province: "云南省",
			City:     "大理白族自治州",
			County:   "鹤庆县",
		},
		"533100": {
			Province: "云南省",
			City:     "德宏傣族景颇族自治州",
			County:   "德宏傣族景颇族自治州",
		},
		"533102": {
			Province: "云南省",
			City:     "德宏傣族景颇族自治州",
			County:   "瑞丽市",
		},
		"533103": {
			Province: "云南省",
			City:     "德宏傣族景颇族自治州",
			County:   "芒市",
		},
		"533122": {
			Province: "云南省",
			City:     "德宏傣族景颇族自治州",
			County:   "梁河县",
		},
		"533123": {
			Province: "云南省",
			City:     "德宏傣族景颇族自治州",
			County:   "盈江县",
		},
		"533124": {
			Province: "云南省",
			City:     "德宏傣族景颇族自治州",
			County:   "陇川县",
		},
		"533300": {
			Province: "云南省",
			City:     "怒江傈僳族自治州",
			County:   "怒江傈僳族自治州",
		},
		"533301": {
			Province: "云南省",
			City:     "怒江傈僳族自治州",
			County:   "泸水市",
		},
		"533323": {
			Province: "云南省",
			City:     "怒江傈僳族自治州",
			County:   "福贡县",
		},
		"533324": {
			Province: "云南省",
			City:     "怒江傈僳族自治州",
			County:   "贡山独龙族怒族自治县",
		},
		"533325": {
			Province: "云南省",
			City:     "怒江傈僳族自治州",
			County:   "兰坪白族普米族自治县",
		},
		"533400": {
			Province: "云南省",
			City:     "迪庆藏族自治州",
			County:   "迪庆藏族自治州",
		},
		"533401": {
			Province: "云南省",
			City:     "迪庆藏族自治州",
			County:   "香格里拉市",
		},
		"533422": {
			Province: "云南省",
			City:     "迪庆藏族自治州",
			County:   "德钦县",
		},
		"533423": {
			Province: "云南省",
			City:     "迪庆藏族自治州",
			County:   "维西傈僳族自治县",
		},
		"540000": {
			Province: "西藏自治区",
			City:     "西藏自治区",
			County:   "西藏自治区",
		},
		"540100": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "拉萨市",
		},
		"540102": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "城关区",
		},
		"540103": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "堆龙德庆区",
		},
		"540104": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "达孜区",
		},
		"540121": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "林周县",
		},
		"540122": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "当雄县",
		},
		"540123": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "尼木县",
		},
		"540124": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "曲水县",
		},
		"540127": {
			Province: "西藏自治区",
			City:     "拉萨市",
			County:   "墨竹工卡县",
		},
		"540200": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "日喀则市",
		},
		"540202": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "桑珠孜区",
		},
		"540221": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "南木林县",
		},
		"540222": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "江孜县",
		},
		"540223": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "定日县",
		},
		"540224": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "萨迦县",
		},
		"540225": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "拉孜县",
		},
		"540226": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "昂仁县",
		},
		"540227": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "谢通门县",
		},
		"540228": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "白朗县",
		},
		"540229": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "仁布县",
		},
		"540230": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "康马县",
		},
		"540231": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "定结县",
		},
		"540232": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "仲巴县",
		},
		"540233": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "亚东县",
		},
		"540234": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "吉隆县",
		},
		"540235": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "聂拉木县",
		},
		"540236": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "萨嘎县",
		},
		"540237": {
			Province: "西藏自治区",
			City:     "日喀则市",
			County:   "岗巴县",
		},
		"540300": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "昌都市",
		},
		"540302": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "卡若区",
		},
		"540321": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "江达县",
		},
		"540322": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "贡觉县",
		},
		"540323": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "类乌齐县",
		},
		"540324": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "丁青县",
		},
		"540325": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "察雅县",
		},
		"540326": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "八宿县",
		},
		"540327": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "左贡县",
		},
		"540328": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "芒康县",
		},
		"540329": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "洛隆县",
		},
		"540330": {
			Province: "西藏自治区",
			City:     "昌都市",
			County:   "边坝县",
		},
		"540400": {
			Province: "西藏自治区",
			City:     "林芝市",
			County:   "林芝市",
		},
		"540402": {
			Province: "西藏自治区",
			City:     "林芝市",
			County:   "巴宜区",
		},
		"540421": {
			Province: "西藏自治区",
			City:     "林芝市",
			County:   "工布江达县",
		},
		"540422": {
			Province: "西藏自治区",
			City:     "林芝市",
			County:   "米林县",
		},
		"540423": {
			Province: "西藏自治区",
			City:     "林芝市",
			County:   "墨脱县",
		},
		"540424": {
			Province: "西藏自治区",
			City:     "林芝市",
			County:   "波密县",
		},
		"540425": {
			Province: "西藏自治区",
			City:     "林芝市",
			County:   "察隅县",
		},
		"540426": {
			Province: "西藏自治区",
			City:     "林芝市",
			County:   "朗县",
		},
		"540500": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "山南市",
		},
		"540502": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "乃东区",
		},
		"540521": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "扎囊县",
		},
		"540522": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "贡嘎县",
		},
		"540523": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "桑日县",
		},
		"540524": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "琼结县",
		},
		"540525": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "曲松县",
		},
		"540526": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "措美县",
		},
		"540527": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "洛扎县",
		},
		"540528": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "加查县",
		},
		"540529": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "隆子县",
		},
		"540530": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "错那县",
		},
		"540531": {
			Province: "西藏自治区",
			City:     "山南市",
			County:   "浪卡子县",
		},
		"540600": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "那曲市",
		},
		"540602": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "色尼区",
		},
		"540621": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "嘉黎县",
		},
		"540622": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "比如县",
		},
		"540623": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "聂荣县",
		},
		"540624": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "安多县",
		},
		"540625": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "申扎县",
		},
		"540626": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "索县",
		},
		"540627": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "班戈县",
		},
		"540628": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "巴青县",
		},
		"540629": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "尼玛县",
		},
		"540630": {
			Province: "西藏自治区",
			City:     "那曲市",
			County:   "双湖县",
		},
		"542500": {
			Province: "西藏自治区",
			City:     "阿里地区",
			County:   "阿里地区",
		},
		"542521": {
			Province: "西藏自治区",
			City:     "阿里地区",
			County:   "普兰县",
		},
		"542522": {
			Province: "西藏自治区",
			City:     "阿里地区",
			County:   "札达县",
		},
		"542523": {
			Province: "西藏自治区",
			City:     "阿里地区",
			County:   "噶尔县",
		},
		"542524": {
			Province: "西藏自治区",
			City:     "阿里地区",
			County:   "日土县",
		},
		"542525": {
			Province: "西藏自治区",
			City:     "阿里地区",
			County:   "革吉县",
		},
		"542526": {
			Province: "西藏自治区",
			City:     "阿里地区",
			County:   "改则县",
		},
		"542527": {
			Province: "西藏自治区",
			City:     "阿里地区",
			County:   "措勤县",
		},
		"610000": {
			Province: "陕西省",
			City:     "陕西省",
			County:   "陕西省",
		},
		"610100": {
			Province: "陕西省",
			City:     "西安市",
			County:   "西安市",
		},
		"610102": {
			Province: "陕西省",
			City:     "西安市",
			County:   "新城区",
		},
		"610103": {
			Province: "陕西省",
			City:     "西安市",
			County:   "碑林区",
		},
		"610104": {
			Province: "陕西省",
			City:     "西安市",
			County:   "莲湖区",
		},
		"610111": {
			Province: "陕西省",
			City:     "西安市",
			County:   "灞桥区",
		},
		"610112": {
			Province: "陕西省",
			City:     "西安市",
			County:   "未央区",
		},
		"610113": {
			Province: "陕西省",
			City:     "西安市",
			County:   "雁塔区",
		},
		"610114": {
			Province: "陕西省",
			City:     "西安市",
			County:   "阎良区",
		},
		"610115": {
			Province: "陕西省",
			City:     "西安市",
			County:   "临潼区",
		},
		"610116": {
			Province: "陕西省",
			City:     "西安市",
			County:   "长安区",
		},
		"610117": {
			Province: "陕西省",
			City:     "西安市",
			County:   "高陵区",
		},
		"610118": {
			Province: "陕西省",
			City:     "西安市",
			County:   "鄠邑区",
		},
		"610122": {
			Province: "陕西省",
			City:     "西安市",
			County:   "蓝田县",
		},
		"610124": {
			Province: "陕西省",
			City:     "西安市",
			County:   "周至县",
		},
		"610200": {
			Province: "陕西省",
			City:     "铜川市",
			County:   "铜川市",
		},
		"610202": {
			Province: "陕西省",
			City:     "铜川市",
			County:   "王益区",
		},
		"610203": {
			Province: "陕西省",
			City:     "铜川市",
			County:   "印台区",
		},
		"610204": {
			Province: "陕西省",
			City:     "铜川市",
			County:   "耀州区",
		},
		"610222": {
			Province: "陕西省",
			City:     "铜川市",
			County:   "宜君县",
		},
		"610300": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "宝鸡市",
		},
		"610302": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "渭滨区",
		},
		"610303": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "金台区",
		},
		"610304": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "陈仓区",
		},
		"610322": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "凤翔县",
		},
		"610323": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "岐山县",
		},
		"610324": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "扶风县",
		},
		"610326": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "眉县",
		},
		"610327": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "陇县",
		},
		"610328": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "千阳县",
		},
		"610329": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "麟游县",
		},
		"610330": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "凤县",
		},
		"610331": {
			Province: "陕西省",
			City:     "宝鸡市",
			County:   "太白县",
		},
		"610400": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "咸阳市",
		},
		"610402": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "秦都区",
		},
		"610403": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "杨陵区",
		},
		"610404": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "渭城区",
		},
		"610422": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "三原县",
		},
		"610423": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "泾阳县",
		},
		"610424": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "乾县",
		},
		"610425": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "礼泉县",
		},
		"610426": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "永寿县",
		},
		"610428": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "长武县",
		},
		"610429": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "旬邑县",
		},
		"610430": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "淳化县",
		},
		"610431": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "武功县",
		},
		"610481": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "兴平市",
		},
		"610482": {
			Province: "陕西省",
			City:     "咸阳市",
			County:   "彬州市",
		},
		"610500": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "渭南市",
		},
		"610502": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "临渭区",
		},
		"610503": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "华州区",
		},
		"610522": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "潼关县",
		},
		"610523": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "大荔县",
		},
		"610524": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "合阳县",
		},
		"610525": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "澄城县",
		},
		"610526": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "蒲城县",
		},
		"610527": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "白水县",
		},
		"610528": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "富平县",
		},
		"610581": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "韩城市",
		},
		"610582": {
			Province: "陕西省",
			City:     "渭南市",
			County:   "华阴市",
		},
		"610600": {
			Province: "陕西省",
			City:     "延安市",
			County:   "延安市",
		},
		"610602": {
			Province: "陕西省",
			City:     "延安市",
			County:   "宝塔区",
		},
		"610603": {
			Province: "陕西省",
			City:     "延安市",
			County:   "安塞区",
		},
		"610621": {
			Province: "陕西省",
			City:     "延安市",
			County:   "延长县",
		},
		"610622": {
			Province: "陕西省",
			City:     "延安市",
			County:   "延川县",
		},
		"610625": {
			Province: "陕西省",
			City:     "延安市",
			County:   "志丹县",
		},
		"610626": {
			Province: "陕西省",
			City:     "延安市",
			County:   "吴起县",
		},
		"610627": {
			Province: "陕西省",
			City:     "延安市",
			County:   "甘泉县",
		},
		"610628": {
			Province: "陕西省",
			City:     "延安市",
			County:   "富县",
		},
		"610629": {
			Province: "陕西省",
			City:     "延安市",
			County:   "洛川县",
		},
		"610630": {
			Province: "陕西省",
			City:     "延安市",
			County:   "宜川县",
		},
		"610631": {
			Province: "陕西省",
			City:     "延安市",
			County:   "黄龙县",
		},
		"610632": {
			Province: "陕西省",
			City:     "延安市",
			County:   "黄陵县",
		},
		"610681": {
			Province: "陕西省",
			City:     "延安市",
			County:   "子长市",
		},
		"610700": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "汉中市",
		},
		"610702": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "汉台区",
		},
		"610703": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "南郑区",
		},
		"610722": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "城固县",
		},
		"610723": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "洋县",
		},
		"610724": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "西乡县",
		},
		"610725": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "勉县",
		},
		"610726": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "宁强县",
		},
		"610727": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "略阳县",
		},
		"610728": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "镇巴县",
		},
		"610729": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "留坝县",
		},
		"610730": {
			Province: "陕西省",
			City:     "汉中市",
			County:   "佛坪县",
		},
		"610800": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "榆林市",
		},
		"610802": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "榆阳区",
		},
		"610803": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "横山区",
		},
		"610822": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "府谷县",
		},
		"610824": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "靖边县",
		},
		"610825": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "定边县",
		},
		"610826": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "绥德县",
		},
		"610827": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "米脂县",
		},
		"610828": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "佳县",
		},
		"610829": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "吴堡县",
		},
		"610830": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "清涧县",
		},
		"610831": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "子洲县",
		},
		"610881": {
			Province: "陕西省",
			City:     "榆林市",
			County:   "神木市",
		},
		"610900": {
			Province: "陕西省",
			City:     "安康市",
			County:   "安康市",
		},
		"610902": {
			Province: "陕西省",
			City:     "安康市",
			County:   "汉滨区",
		},
		"610921": {
			Province: "陕西省",
			City:     "安康市",
			County:   "汉阴县",
		},
		"610922": {
			Province: "陕西省",
			City:     "安康市",
			County:   "石泉县",
		},
		"610923": {
			Province: "陕西省",
			City:     "安康市",
			County:   "宁陕县",
		},
		"610924": {
			Province: "陕西省",
			City:     "安康市",
			County:   "紫阳县",
		},
		"610925": {
			Province: "陕西省",
			City:     "安康市",
			County:   "岚皋县",
		},
		"610926": {
			Province: "陕西省",
			City:     "安康市",
			County:   "平利县",
		},
		"610927": {
			Province: "陕西省",
			City:     "安康市",
			County:   "镇坪县",
		},
		"610928": {
			Province: "陕西省",
			City:     "安康市",
			County:   "旬阳县",
		},
		"610929": {
			Province: "陕西省",
			City:     "安康市",
			County:   "白河县",
		},
		"611000": {
			Province: "陕西省",
			City:     "商洛市",
			County:   "商洛市",
		},
		"611002": {
			Province: "陕西省",
			City:     "商洛市",
			County:   "商州区",
		},
		"611021": {
			Province: "陕西省",
			City:     "商洛市",
			County:   "洛南县",
		},
		"611022": {
			Province: "陕西省",
			City:     "商洛市",
			County:   "丹凤县",
		},
		"611023": {
			Province: "陕西省",
			City:     "商洛市",
			County:   "商南县",
		},
		"611024": {
			Province: "陕西省",
			City:     "商洛市",
			County:   "山阳县",
		},
		"611025": {
			Province: "陕西省",
			City:     "商洛市",
			County:   "镇安县",
		},
		"611026": {
			Province: "陕西省",
			City:     "商洛市",
			County:   "柞水县",
		},
		"620000": {
			Province: "甘肃省",
			City:     "甘肃省",
			County:   "甘肃省",
		},
		"620100": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "兰州市",
		},
		"620102": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "城关区",
		},
		"620103": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "七里河区",
		},
		"620104": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "西固区",
		},
		"620105": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "安宁区",
		},
		"620111": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "红古区",
		},
		"620121": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "永登县",
		},
		"620122": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "皋兰县",
		},
		"620123": {
			Province: "甘肃省",
			City:     "兰州市",
			County:   "榆中县",
		},
		"620200": {
			Province: "甘肃省",
			City:     "嘉峪关市",
			County:   "嘉峪关市",
		},
		"620300": {
			Province: "甘肃省",
			City:     "金昌市",
			County:   "金昌市",
		},
		"620302": {
			Province: "甘肃省",
			City:     "金昌市",
			County:   "金川区",
		},
		"620321": {
			Province: "甘肃省",
			City:     "金昌市",
			County:   "永昌县",
		},
		"620400": {
			Province: "甘肃省",
			City:     "白银市",
			County:   "白银市",
		},
		"620402": {
			Province: "甘肃省",
			City:     "白银市",
			County:   "白银区",
		},
		"620403": {
			Province: "甘肃省",
			City:     "白银市",
			County:   "平川区",
		},
		"620421": {
			Province: "甘肃省",
			City:     "白银市",
			County:   "靖远县",
		},
		"620422": {
			Province: "甘肃省",
			City:     "白银市",
			County:   "会宁县",
		},
		"620423": {
			Province: "甘肃省",
			City:     "白银市",
			County:   "景泰县",
		},
		"620500": {
			Province: "甘肃省",
			City:     "天水市",
			County:   "天水市",
		},
		"620502": {
			Province: "甘肃省",
			City:     "天水市",
			County:   "秦州区",
		},
		"620503": {
			Province: "甘肃省",
			City:     "天水市",
			County:   "麦积区",
		},
		"620521": {
			Province: "甘肃省",
			City:     "天水市",
			County:   "清水县",
		},
		"620522": {
			Province: "甘肃省",
			City:     "天水市",
			County:   "秦安县",
		},
		"620523": {
			Province: "甘肃省",
			City:     "天水市",
			County:   "甘谷县",
		},
		"620524": {
			Province: "甘肃省",
			City:     "天水市",
			County:   "武山县",
		},
		"620525": {
			Province: "甘肃省",
			City:     "天水市",
			County:   "张家川回族自治县",
		},
		"620600": {
			Province: "甘肃省",
			City:     "武威市",
			County:   "武威市",
		},
		"620602": {
			Province: "甘肃省",
			City:     "武威市",
			County:   "凉州区",
		},
		"620621": {
			Province: "甘肃省",
			City:     "武威市",
			County:   "民勤县",
		},
		"620622": {
			Province: "甘肃省",
			City:     "武威市",
			County:   "古浪县",
		},
		"620623": {
			Province: "甘肃省",
			City:     "武威市",
			County:   "天祝藏族自治县",
		},
		"620700": {
			Province: "甘肃省",
			City:     "张掖市",
			County:   "张掖市",
		},
		"620702": {
			Province: "甘肃省",
			City:     "张掖市",
			County:   "甘州区",
		},
		"620721": {
			Province: "甘肃省",
			City:     "张掖市",
			County:   "肃南裕固族自治县",
		},
		"620722": {
			Province: "甘肃省",
			City:     "张掖市",
			County:   "民乐县",
		},
		"620723": {
			Province: "甘肃省",
			City:     "张掖市",
			County:   "临泽县",
		},
		"620724": {
			Province: "甘肃省",
			City:     "张掖市",
			County:   "高台县",
		},
		"620725": {
			Province: "甘肃省",
			City:     "张掖市",
			County:   "山丹县",
		},
		"620800": {
			Province: "甘肃省",
			City:     "平凉市",
			County:   "平凉市",
		},
		"620802": {
			Province: "甘肃省",
			City:     "平凉市",
			County:   "崆峒区",
		},
		"620821": {
			Province: "甘肃省",
			City:     "平凉市",
			County:   "泾川县",
		},
		"620822": {
			Province: "甘肃省",
			City:     "平凉市",
			County:   "灵台县",
		},
		"620823": {
			Province: "甘肃省",
			City:     "平凉市",
			County:   "崇信县",
		},
		"620825": {
			Province: "甘肃省",
			City:     "平凉市",
			County:   "庄浪县",
		},
		"620826": {
			Province: "甘肃省",
			City:     "平凉市",
			County:   "静宁县",
		},
		"620881": {
			Province: "甘肃省",
			City:     "平凉市",
			County:   "华亭市",
		},
		"620900": {
			Province: "甘肃省",
			City:     "酒泉市",
			County:   "酒泉市",
		},
		"620902": {
			Province: "甘肃省",
			City:     "酒泉市",
			County:   "肃州区",
		},
		"620921": {
			Province: "甘肃省",
			City:     "酒泉市",
			County:   "金塔县",
		},
		"620922": {
			Province: "甘肃省",
			City:     "酒泉市",
			County:   "瓜州县",
		},
		"620923": {
			Province: "甘肃省",
			City:     "酒泉市",
			County:   "肃北蒙古族自治县",
		},
		"620924": {
			Province: "甘肃省",
			City:     "酒泉市",
			County:   "阿克塞哈萨克族自治县",
		},
		"620981": {
			Province: "甘肃省",
			City:     "酒泉市",
			County:   "玉门市",
		},
		"620982": {
			Province: "甘肃省",
			City:     "酒泉市",
			County:   "敦煌市",
		},
		"621000": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "庆阳市",
		},
		"621002": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "西峰区",
		},
		"621021": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "庆城县",
		},
		"621022": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "环县",
		},
		"621023": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "华池县",
		},
		"621024": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "合水县",
		},
		"621025": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "正宁县",
		},
		"621026": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "宁县",
		},
		"621027": {
			Province: "甘肃省",
			City:     "庆阳市",
			County:   "镇原县",
		},
		"621100": {
			Province: "甘肃省",
			City:     "定西市",
			County:   "定西市",
		},
		"621102": {
			Province: "甘肃省",
			City:     "定西市",
			County:   "安定区",
		},
		"621121": {
			Province: "甘肃省",
			City:     "定西市",
			County:   "通渭县",
		},
		"621122": {
			Province: "甘肃省",
			City:     "定西市",
			County:   "陇西县",
		},
		"621123": {
			Province: "甘肃省",
			City:     "定西市",
			County:   "渭源县",
		},
		"621124": {
			Province: "甘肃省",
			City:     "定西市",
			County:   "临洮县",
		},
		"621125": {
			Province: "甘肃省",
			City:     "定西市",
			County:   "漳县",
		},
		"621126": {
			Province: "甘肃省",
			City:     "定西市",
			County:   "岷县",
		},
		"621200": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "陇南市",
		},
		"621202": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "武都区",
		},
		"621221": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "成县",
		},
		"621222": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "文县",
		},
		"621223": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "宕昌县",
		},
		"621224": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "康县",
		},
		"621225": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "西和县",
		},
		"621226": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "礼县",
		},
		"621227": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "徽县",
		},
		"621228": {
			Province: "甘肃省",
			City:     "陇南市",
			County:   "两当县",
		},
		"622900": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "临夏回族自治州",
		},
		"622901": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "临夏市",
		},
		"622921": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "临夏县",
		},
		"622922": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "康乐县",
		},
		"622923": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "永靖县",
		},
		"622924": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "广河县",
		},
		"622925": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "和政县",
		},
		"622926": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "东乡族自治县",
		},
		"622927": {
			Province: "甘肃省",
			City:     "临夏回族自治州",
			County:   "积石山保安族东乡族撒拉族自治县",
		},
		"623000": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "甘南藏族自治州",
		},
		"623001": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "合作市",
		},
		"623021": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "临潭县",
		},
		"623022": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "卓尼县",
		},
		"623023": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "舟曲县",
		},
		"623024": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "迭部县",
		},
		"623025": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "玛曲县",
		},
		"623026": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "碌曲县",
		},
		"623027": {
			Province: "甘肃省",
			City:     "甘南藏族自治州",
			County:   "夏河县",
		},
		"630000": {
			Province: "青海省",
			City:     "青海省",
			County:   "青海省",
		},
		"630100": {
			Province: "青海省",
			City:     "西宁市",
			County:   "西宁市",
		},
		"630102": {
			Province: "青海省",
			City:     "西宁市",
			County:   "城东区",
		},
		"630103": {
			Province: "青海省",
			City:     "西宁市",
			County:   "城中区",
		},
		"630104": {
			Province: "青海省",
			City:     "西宁市",
			County:   "城西区",
		},
		"630105": {
			Province: "青海省",
			City:     "西宁市",
			County:   "城北区",
		},
		"630106": {
			Province: "青海省",
			City:     "西宁市",
			County:   "湟中区",
		},
		"630121": {
			Province: "青海省",
			City:     "西宁市",
			County:   "大通回族土族自治县",
		},
		"630123": {
			Province: "青海省",
			City:     "西宁市",
			County:   "湟源县",
		},
		"630200": {
			Province: "青海省",
			City:     "海东市",
			County:   "海东市",
		},
		"630202": {
			Province: "青海省",
			City:     "海东市",
			County:   "乐都区",
		},
		"630203": {
			Province: "青海省",
			City:     "海东市",
			County:   "平安区",
		},
		"630222": {
			Province: "青海省",
			City:     "海东市",
			County:   "民和回族土族自治县",
		},
		"630223": {
			Province: "青海省",
			City:     "海东市",
			County:   "互助土族自治县",
		},
		"630224": {
			Province: "青海省",
			City:     "海东市",
			County:   "化隆回族自治县",
		},
		"630225": {
			Province: "青海省",
			City:     "海东市",
			County:   "循化撒拉族自治县",
		},
		"632200": {
			Province: "青海省",
			City:     "海北藏族自治州",
			County:   "海北藏族自治州",
		},
		"632221": {
			Province: "青海省",
			City:     "海北藏族自治州",
			County:   "门源回族自治县",
		},
		"632222": {
			Province: "青海省",
			City:     "海北藏族自治州",
			County:   "祁连县",
		},
		"632223": {
			Province: "青海省",
			City:     "海北藏族自治州",
			County:   "海晏县",
		},
		"632224": {
			Province: "青海省",
			City:     "海北藏族自治州",
			County:   "刚察县",
		},
		"632300": {
			Province: "青海省",
			City:     "黄南藏族自治州",
			County:   "黄南藏族自治州",
		},
		"632301": {
			Province: "青海省",
			City:     "黄南藏族自治州",
			County:   "同仁市",
		},
		"632322": {
			Province: "青海省",
			City:     "黄南藏族自治州",
			County:   "尖扎县",
		},
		"632323": {
			Province: "青海省",
			City:     "黄南藏族自治州",
			County:   "泽库县",
		},
		"632324": {
			Province: "青海省",
			City:     "黄南藏族自治州",
			County:   "河南蒙古族自治县",
		},
		"632500": {
			Province: "青海省",
			City:     "海南藏族自治州",
			County:   "海南藏族自治州",
		},
		"632521": {
			Province: "青海省",
			City:     "海南藏族自治州",
			County:   "共和县",
		},
		"632522": {
			Province: "青海省",
			City:     "海南藏族自治州",
			County:   "同德县",
		},
		"632523": {
			Province: "青海省",
			City:     "海南藏族自治州",
			County:   "贵德县",
		},
		"632524": {
			Province: "青海省",
			City:     "海南藏族自治州",
			County:   "兴海县",
		},
		"632525": {
			Province: "青海省",
			City:     "海南藏族自治州",
			County:   "贵南县",
		},
		"632600": {
			Province: "青海省",
			City:     "果洛藏族自治州",
			County:   "果洛藏族自治州",
		},
		"632621": {
			Province: "青海省",
			City:     "果洛藏族自治州",
			County:   "玛沁县",
		},
		"632622": {
			Province: "青海省",
			City:     "果洛藏族自治州",
			County:   "班玛县",
		},
		"632623": {
			Province: "青海省",
			City:     "果洛藏族自治州",
			County:   "甘德县",
		},
		"632624": {
			Province: "青海省",
			City:     "果洛藏族自治州",
			County:   "达日县",
		},
		"632625": {
			Province: "青海省",
			City:     "果洛藏族自治州",
			County:   "久治县",
		},
		"632626": {
			Province: "青海省",
			City:     "果洛藏族自治州",
			County:   "玛多县",
		},
		"632700": {
			Province: "青海省",
			City:     "玉树藏族自治州",
			County:   "玉树藏族自治州",
		},
		"632701": {
			Province: "青海省",
			City:     "玉树藏族自治州",
			County:   "玉树市",
		},
		"632722": {
			Province: "青海省",
			City:     "玉树藏族自治州",
			County:   "杂多县",
		},
		"632723": {
			Province: "青海省",
			City:     "玉树藏族自治州",
			County:   "称多县",
		},
		"632724": {
			Province: "青海省",
			City:     "玉树藏族自治州",
			County:   "治多县",
		},
		"632725": {
			Province: "青海省",
			City:     "玉树藏族自治州",
			County:   "囊谦县",
		},
		"632726": {
			Province: "青海省",
			City:     "玉树藏族自治州",
			County:   "曲麻莱县",
		},
		"632800": {
			Province: "青海省",
			City:     "海西蒙古族藏族自治州",
			County:   "海西蒙古族藏族自治州",
		},
		"632801": {
			Province: "青海省",
			City:     "海西蒙古族藏族自治州",
			County:   "格尔木市",
		},
		"632802": {
			Province: "青海省",
			City:     "海西蒙古族藏族自治州",
			County:   "德令哈市",
		},
		"632803": {
			Province: "青海省",
			City:     "海西蒙古族藏族自治州",
			County:   "茫崖市",
		},
		"632821": {
			Province: "青海省",
			City:     "海西蒙古族藏族自治州",
			County:   "乌兰县",
		},
		"632822": {
			Province: "青海省",
			City:     "海西蒙古族藏族自治州",
			County:   "都兰县",
		},
		"632823": {
			Province: "青海省",
			City:     "海西蒙古族藏族自治州",
			County:   "天峻县",
		},
		"640000": {
			Province: "宁夏回族自治区",
			City:     "宁夏回族自治区",
			County:   "宁夏回族自治区",
		},
		"640100": {
			Province: "宁夏回族自治区",
			City:     "银川市",
			County:   "银川市",
		},
		"640104": {
			Province: "宁夏回族自治区",
			City:     "银川市",
			County:   "兴庆区",
		},
		"640105": {
			Province: "宁夏回族自治区",
			City:     "银川市",
			County:   "西夏区",
		},
		"640106": {
			Province: "宁夏回族自治区",
			City:     "银川市",
			County:   "金凤区",
		},
		"640121": {
			Province: "宁夏回族自治区",
			City:     "银川市",
			County:   "永宁县",
		},
		"640122": {
			Province: "宁夏回族自治区",
			City:     "银川市",
			County:   "贺兰县",
		},
		"640181": {
			Province: "宁夏回族自治区",
			City:     "银川市",
			County:   "灵武市",
		},
		"640200": {
			Province: "宁夏回族自治区",
			City:     "石嘴山市",
			County:   "石嘴山市",
		},
		"640202": {
			Province: "宁夏回族自治区",
			City:     "石嘴山市",
			County:   "大武口区",
		},
		"640205": {
			Province: "宁夏回族自治区",
			City:     "石嘴山市",
			County:   "惠农区",
		},
		"640221": {
			Province: "宁夏回族自治区",
			City:     "石嘴山市",
			County:   "平罗县",
		},
		"640300": {
			Province: "宁夏回族自治区",
			City:     "吴忠市",
			County:   "吴忠市",
		},
		"640302": {
			Province: "宁夏回族自治区",
			City:     "吴忠市",
			County:   "利通区",
		},
		"640303": {
			Province: "宁夏回族自治区",
			City:     "吴忠市",
			County:   "红寺堡区",
		},
		"640323": {
			Province: "宁夏回族自治区",
			City:     "吴忠市",
			County:   "盐池县",
		},
		"640324": {
			Province: "宁夏回族自治区",
			City:     "吴忠市",
			County:   "同心县",
		},
		"640381": {
			Province: "宁夏回族自治区",
			City:     "吴忠市",
			County:   "青铜峡市",
		},
		"640400": {
			Province: "宁夏回族自治区",
			City:     "固原市",
			County:   "固原市",
		},
		"640402": {
			Province: "宁夏回族自治区",
			City:     "固原市",
			County:   "原州区",
		},
		"640422": {
			Province: "宁夏回族自治区",
			City:     "固原市",
			County:   "西吉县",
		},
		"640423": {
			Province: "宁夏回族自治区",
			City:     "固原市",
			County:   "隆德县",
		},
		"640424": {
			Province: "宁夏回族自治区",
			City:     "固原市",
			County:   "泾源县",
		},
		"640425": {
			Province: "宁夏回族自治区",
			City:     "固原市",
			County:   "彭阳县",
		},
		"640500": {
			Province: "宁夏回族自治区",
			City:     "中卫市",
			County:   "中卫市",
		},
		"640502": {
			Province: "宁夏回族自治区",
			City:     "中卫市",
			County:   "沙坡头区",
		},
		"640521": {
			Province: "宁夏回族自治区",
			City:     "中卫市",
			County:   "中宁县",
		},
		"640522": {
			Province: "宁夏回族自治区",
			City:     "中卫市",
			County:   "海原县",
		},
		"650000": {
			Province: "新疆维吾尔自治区",
			City:     "新疆维吾尔自治区",
			County:   "新疆维吾尔自治区",
		},
		"650100": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "乌鲁木齐市",
		},
		"650102": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "天山区",
		},
		"650103": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "沙依巴克区",
		},
		"650104": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "新市区",
		},
		"650105": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "水磨沟区",
		},
		"650106": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "头屯河区",
		},
		"650107": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "达坂城区",
		},
		"650109": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "米东区",
		},
		"650121": {
			Province: "新疆维吾尔自治区",
			City:     "乌鲁木齐市",
			County:   "乌鲁木齐县",
		},
		"650200": {
			Province: "新疆维吾尔自治区",
			City:     "克拉玛依市",
			County:   "克拉玛依市",
		},
		"650202": {
			Province: "新疆维吾尔自治区",
			City:     "克拉玛依市",
			County:   "独山子区",
		},
		"650203": {
			Province: "新疆维吾尔自治区",
			City:     "克拉玛依市",
			County:   "克拉玛依区",
		},
		"650204": {
			Province: "新疆维吾尔自治区",
			City:     "克拉玛依市",
			County:   "白碱滩区",
		},
		"650205": {
			Province: "新疆维吾尔自治区",
			City:     "克拉玛依市",
			County:   "乌尔禾区",
		},
		"650400": {
			Province: "新疆维吾尔自治区",
			City:     "吐鲁番市",
			County:   "吐鲁番市",
		},
		"650402": {
			Province: "新疆维吾尔自治区",
			City:     "吐鲁番市",
			County:   "高昌区",
		},
		"650421": {
			Province: "新疆维吾尔自治区",
			City:     "吐鲁番市",
			County:   "鄯善县",
		},
		"650422": {
			Province: "新疆维吾尔自治区",
			City:     "吐鲁番市",
			County:   "托克逊县",
		},
		"650500": {
			Province: "新疆维吾尔自治区",
			City:     "哈密市",
			County:   "哈密市",
		},
		"650502": {
			Province: "新疆维吾尔自治区",
			City:     "哈密市",
			County:   "伊州区",
		},
		"650521": {
			Province: "新疆维吾尔自治区",
			City:     "哈密市",
			County:   "巴里坤哈萨克自治县",
		},
		"650522": {
			Province: "新疆维吾尔自治区",
			City:     "哈密市",
			County:   "伊吾县",
		},
		"652300": {
			Province: "新疆维吾尔自治区",
			City:     "昌吉回族自治州",
			County:   "昌吉回族自治州",
		},
		"652301": {
			Province: "新疆维吾尔自治区",
			City:     "昌吉回族自治州",
			County:   "昌吉市",
		},
		"652302": {
			Province: "新疆维吾尔自治区",
			City:     "昌吉回族自治州",
			County:   "阜康市",
		},
		"652323": {
			Province: "新疆维吾尔自治区",
			City:     "昌吉回族自治州",
			County:   "呼图壁县",
		},
		"652324": {
			Province: "新疆维吾尔自治区",
			City:     "昌吉回族自治州",
			County:   "玛纳斯县",
		},
		"652325": {
			Province: "新疆维吾尔自治区",
			City:     "昌吉回族自治州",
			County:   "奇台县",
		},
		"652327": {
			Province: "新疆维吾尔自治区",
			City:     "昌吉回族自治州",
			County:   "吉木萨尔县",
		},
		"652328": {
			Province: "新疆维吾尔自治区",
			City:     "昌吉回族自治州",
			County:   "木垒哈萨克自治县",
		},
		"652700": {
			Province: "新疆维吾尔自治区",
			City:     "博尔塔拉蒙古自治州",
			County:   "博尔塔拉蒙古自治州",
		},
		"652701": {
			Province: "新疆维吾尔自治区",
			City:     "博尔塔拉蒙古自治州",
			County:   "博乐市",
		},
		"652702": {
			Province: "新疆维吾尔自治区",
			City:     "博尔塔拉蒙古自治州",
			County:   "阿拉山口市",
		},
		"652722": {
			Province: "新疆维吾尔自治区",
			City:     "博尔塔拉蒙古自治州",
			County:   "精河县",
		},
		"652723": {
			Province: "新疆维吾尔自治区",
			City:     "博尔塔拉蒙古自治州",
			County:   "温泉县",
		},
		"652800": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "巴音郭楞蒙古自治州",
		},
		"652801": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "库尔勒市",
		},
		"652822": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "轮台县",
		},
		"652823": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "尉犁县",
		},
		"652824": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "若羌县",
		},
		"652825": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "且末县",
		},
		"652826": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "焉耆回族自治县",
		},
		"652827": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "和静县",
		},
		"652828": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "和硕县",
		},
		"652829": {
			Province: "新疆维吾尔自治区",
			City:     "巴音郭楞蒙古自治州",
			County:   "博湖县",
		},
		"652900": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "阿克苏地区",
		},
		"652901": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "阿克苏市",
		},
		"652902": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "库车市",
		},
		"652922": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "温宿县",
		},
		"652924": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "沙雅县",
		},
		"652925": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "新和县",
		},
		"652926": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "拜城县",
		},
		"652927": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "乌什县",
		},
		"652928": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "阿瓦提县",
		},
		"652929": {
			Province: "新疆维吾尔自治区",
			City:     "阿克苏地区",
			County:   "柯坪县",
		},
		"653000": {
			Province: "新疆维吾尔自治区",
			City:     "克孜勒苏柯尔克孜自治州",
			County:   "克孜勒苏柯尔克孜自治州",
		},
		"653001": {
			Province: "新疆维吾尔自治区",
			City:     "克孜勒苏柯尔克孜自治州",
			County:   "阿图什市",
		},
		"653022": {
			Province: "新疆维吾尔自治区",
			City:     "克孜勒苏柯尔克孜自治州",
			County:   "阿克陶县",
		},
		"653023": {
			Province: "新疆维吾尔自治区",
			City:     "克孜勒苏柯尔克孜自治州",
			County:   "阿合奇县",
		},
		"653024": {
			Province: "新疆维吾尔自治区",
			City:     "克孜勒苏柯尔克孜自治州",
			County:   "乌恰县",
		},
		"653100": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "喀什地区",
		},
		"653101": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "喀什市",
		},
		"653121": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "疏附县",
		},
		"653122": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "疏勒县",
		},
		"653123": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "英吉沙县",
		},
		"653124": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "泽普县",
		},
		"653125": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "莎车县",
		},
		"653126": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "叶城县",
		},
		"653127": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "麦盖提县",
		},
		"653128": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "岳普湖县",
		},
		"653129": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "伽师县",
		},
		"653130": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "巴楚县",
		},
		"653131": {
			Province: "新疆维吾尔自治区",
			City:     "喀什地区",
			County:   "塔什库尔干塔吉克自治县",
		},
		"653200": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "和田地区",
		},
		"653201": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "和田市",
		},
		"653221": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "和田县",
		},
		"653222": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "墨玉县",
		},
		"653223": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "皮山县",
		},
		"653224": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "洛浦县",
		},
		"653225": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "策勒县",
		},
		"653226": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "于田县",
		},
		"653227": {
			Province: "新疆维吾尔自治区",
			City:     "和田地区",
			County:   "民丰县",
		},
		"654000": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "伊犁哈萨克自治州",
		},
		"654002": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "伊宁市",
		},
		"654003": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "奎屯市",
		},
		"654004": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "霍尔果斯市",
		},
		"654021": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "伊宁县",
		},
		"654022": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "察布查尔锡伯自治县",
		},
		"654023": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "霍城县",
		},
		"654024": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "巩留县",
		},
		"654025": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "新源县",
		},
		"654026": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "昭苏县",
		},
		"654027": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "特克斯县",
		},
		"654028": {
			Province: "新疆维吾尔自治区",
			City:     "伊犁哈萨克自治州",
			County:   "尼勒克县",
		},
		"654200": {
			Province: "新疆维吾尔自治区",
			City:     "塔城地区",
			County:   "塔城地区",
		},
		"654201": {
			Province: "新疆维吾尔自治区",
			City:     "塔城地区",
			County:   "塔城市",
		},
		"654202": {
			Province: "新疆维吾尔自治区",
			City:     "塔城地区",
			County:   "乌苏市",
		},
		"654221": {
			Province: "新疆维吾尔自治区",
			City:     "塔城地区",
			County:   "额敏县",
		},
		"654223": {
			Province: "新疆维吾尔自治区",
			City:     "塔城地区",
			County:   "沙湾县",
		},
		"654224": {
			Province: "新疆维吾尔自治区",
			City:     "塔城地区",
			County:   "托里县",
		},
		"654225": {
			Province: "新疆维吾尔自治区",
			City:     "塔城地区",
			County:   "裕民县",
		},
		"654226": {
			Province: "新疆维吾尔自治区",
			City:     "塔城地区",
			County:   "和布克赛尔蒙古自治县",
		},
		"654300": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "阿勒泰地区",
		},
		"654301": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "阿勒泰市",
		},
		"654321": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "布尔津县",
		},
		"654322": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "富蕴县",
		},
		"654323": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "福海县",
		},
		"654324": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "哈巴河县",
		},
		"654325": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "青河县",
		},
		"654326": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "吉木乃县",
		},
		"659001": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "石河子市",
		},
		"659002": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "阿拉尔市",
		},
		"659003": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "图木舒克市",
		},
		"659004": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "五家渠市",
		},
		"659005": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "北屯市",
		},
		"659006": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "铁门关市",
		},
		"659007": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "双河市",
		},
		"659008": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "可克达拉市",
		},
		"659009": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "昆玉市",
		},
		"659010": {
			Province: "新疆维吾尔自治区",
			City:     "阿勒泰地区",
			County:   "胡杨河市",
		},
		"710000": {
			Province: "台湾省",
			City:     "台湾省",
			County:   "台湾省",
		},
		"810000": {
			Province: "香港特别行政区",
			City:     "香港特别行政区",
			County:   "香港特别行政区",
		},
		"820000": {
			Province: "澳门特别行政区",
			City:     "澳门特别行政区",
			County:   "澳门特别行政区",
		},
	}
)
