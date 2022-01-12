package gb

// GB11714 for "Rules of coding for the representation of organization"
type GB11714 struct {
}

func (p *GB11714) Validate(data string) bool {
	return IsValidOrganizationCode(data)
}

func (p *GB11714) Fake() string {
	data := make([]byte, 9)
	sum := 0
	for i := 0; i < 8; i++ {
		data[i] = RandByte(OrganizationCodeChars)
		sum += OrganizationCodeCharMap[data[i]] * OrganizationCodePower[i]
	}
	data[8] = OrganizationCodeCheckCodes[sum%11]
	return string(data)
}

var (
	OrganizationCodeCheckCodes = []byte("0X987654321")
	OrganizationCodeChars      = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	OrganizationCodeCharMap    = map[byte]int{}
	OrganizationCodePower      = []int{3, 7, 9, 10, 5, 8, 4, 2}
)

func init() {
	for i, x := range OrganizationCodeChars {
		OrganizationCodeCharMap[x] = i
		if i > 9 {
			OrganizationCodeCharMap[x+32] = i
		}
	}

	RegisterGB("GB11714", &GB11714{})
}

// IsValidOrganizationCode Judge the legality of organization code.
// <p>
// number 8: check code, this can be the number 0-9 or the letter x
func IsValidOrganizationCode(organizationCode string) bool {
	if len(organizationCode) != 9 {
		return false
	}
	sum := 0
	for i := range OrganizationCodePower {
		n, ok := OrganizationCodeCharMap[organizationCode[i]]
		if !ok {
			return false
		}
		sum += n * OrganizationCodePower[i]
	}
	check := OrganizationCodeCheckCodes[sum%11]
	return check == organizationCode[8] || (check == 'X' && organizationCode[8] == 'x')
}

func FakeOrganizationCodeInGB32100() string {
	data := make([]byte, 9)
	sum := 0
	for i := 0; i < 8; i++ {
		for true {
			data[i] = RandByte(OrganizationCodeChars)
			if _, ok := UnifiedSocialCreditIDCharMap[data[i]]; ok {
				break
			}
		}
		sum += OrganizationCodeCharMap[data[i]] * OrganizationCodePower[i]
	}
	data[8] = OrganizationCodeCheckCodes[sum%11]
	return string(data)
}
