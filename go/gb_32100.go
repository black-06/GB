package gb

// GB32100 for "The coding rule of the unified social credit identifier for legal entities and other organizations"
type GB32100 struct {
}

func (p *GB32100) Validate(data string) bool {
	return IsValidUnifiedSocialCreditID(data)
}

func (p *GB32100) Fake() string {
	data := make([]byte, 18)
	sum := 0

	data[0] = RandByte(RegistrationManagementDepartmentCodes)
	sum += UnifiedSocialCreditIDCharMap[data[0]] * UnifiedSocialCreditIDPower[0]

	data[1] = RandByte(OrganizationCategoryCodes)
	sum += UnifiedSocialCreditIDCharMap[data[1]] * UnifiedSocialCreditIDPower[1]

	for i := 2; i < 8; i++ {
		data[i] = RandByte(UnifiedSocialCreditIDChars)
		sum += UnifiedSocialCreditIDCharMap[data[i]] * UnifiedSocialCreditIDPower[i]
	}

	organizationCode := FakeOrganizationCodeInGB32100()
	for i := 8; i < 17; i++ {
		data[i] = organizationCode[i-8]
		sum += UnifiedSocialCreditIDCharMap[data[i]] * UnifiedSocialCreditIDPower[i]
	}

	data[17] = UnifiedSocialCreditIDChars[(31-sum%31)%31]
	return string(data)
}

var (
	RegistrationManagementDepartmentCodes      = []byte("159Y")
	RegistrationManagementDepartmentCheckCodes = []byte("159yY")
	OrganizationCategoryCodes                  = []byte("1239")
	UnifiedSocialCreditIDChars                 = []byte("0123456789ABCDEFGHJKLMNPQRTUWXY")
	UnifiedSocialCreditIDLowerCaseChars        = []byte("0123456789abcdefghjklmnpqrtuwxy")

	RegistrationManagementDepartmentCodeMap = map[byte]struct{}{}
	OrganizationCategoryCodeMap             = map[byte]struct{}{}
	UnifiedSocialCreditIDCharMap            = map[byte]int{}
	UnifiedSocialCreditIDPower              = []int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}
)

func init() {
	empty := struct{}{}

	for _, x := range RegistrationManagementDepartmentCheckCodes {
		RegistrationManagementDepartmentCodeMap[x] = empty
	}

	for _, x := range OrganizationCategoryCodes {
		OrganizationCategoryCodeMap[x] = empty
	}

	for i, x := range UnifiedSocialCreditIDChars {
		UnifiedSocialCreditIDCharMap[x] = i
	}
	for i, x := range UnifiedSocialCreditIDLowerCaseChars {
		UnifiedSocialCreditIDCharMap[x] = i
	}

	RegisterGB("GB32100", &GB32100{})
}

// IsValidUnifiedSocialCreditID Judge the legality of unified social credit id.
// <p>
// number 0:    registration management department code (1/5/9/Y).
// number 1:    organization category code (1/2/3/9).
// number 2-7:  administrative division code [GB/T2260-2007].
// number 8-16: organization code (numbers or letters) [GB 11714-1997].
// number 17:   check code (numbers or letters).
func IsValidUnifiedSocialCreditID(creditID string) bool {
	if len(creditID) != 18 {
		return false
	}
	if _, ok := RegistrationManagementDepartmentCodeMap[creditID[0]]; !ok {
		return false
	}
	if _, ok := OrganizationCategoryCodeMap[creditID[1]]; !ok {
		return false
	}

	if !IsValidOrganizationCode(creditID[8:17]) {
		return false
	}

	sum := 0
	for i := range UnifiedSocialCreditIDPower {
		n, ok := UnifiedSocialCreditIDCharMap[creditID[i]]
		if !ok {
			return false
		}
		sum += n * UnifiedSocialCreditIDPower[i]
	}

	check := UnifiedSocialCreditIDChars[(31-sum%31)%31]
	if check == creditID[17] {
		return true
	}
	// check if lower case
	if creditID[17] < 'a' || creditID[17] > 'z' {
		return false
	}
	// to upper case
	return check == (creditID[17] - 32)
}
