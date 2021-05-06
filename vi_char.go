package telex

var (
	subRules = map[string]rune{
		"a":  'a',
		"aw": 'ă',
		"aa": 'â',
		"dd": 'đ',
		"e":  'e',
		"ee": 'ê',
		"i":  'i',
		"o":  'o',
		"oo": 'ô',
		"ow": 'ơ',
		"u":  'u',
		"uw": 'ư',
	}

	maskRules = map[string][]rune{
		"a":  {'a', 'à', 'á', 'ả', 'ã', 'ạ'},
		"aw": {'ă', 'ằ', 'ắ', 'ẳ', 'ẵ', 'ặ'},
		"aa": {'â', 'ầ', 'ấ', 'ẩ', 'ẫ', 'ậ'},
		"e":  {'e', 'è', 'é', 'ẻ', 'ẽ', 'ẹ'},
		"ee": {'ê', 'ề', 'ế', 'ể', 'ễ', 'ệ'},
		"i":  {'i', 'ì', 'í', 'ỉ', 'ĩ', 'ị'},
		"o":  {'o', 'ò', 'ó', 'ỏ', 'õ', 'ọ'},
		"oo": {'ô', 'ồ', 'ố', 'ổ', 'ỗ', 'ộ'},
		"ow": {'ơ', 'ờ', 'ớ', 'ở', 'ỡ', 'ợ'},
		"u":  {'u', 'ù', 'ú', 'ủ', 'ũ', 'ụ'},
		"uw": {'ư', 'ừ', 'ứ', 'ử', 'ữ', 'ự'},
	}

	maskOrders = map[rune]int{
		'z': 0,
		'f': 1,
		's': 2,
		'r': 3,
		'x': 5,
		'j': 6,
	}
)

// main + sub + mask = vietnamese char
// a + w + s = ắ
type viChar struct {
	main rune
	sub  rune
	mask rune
}

func (c *viChar) toRune() rune {
	// Prevent rune 0 convert to string
	mainWithSub := string(c.main)
	if c.sub != 0 {
		mainWithSub += string(c.sub)
	}

	result, ok := subRules[mainWithSub]
	if !ok {
		// main + sub is not valid
		// fallback to main
		return c.main
	}

	if c.mask == 0 {
		return result
	}

	maskRule, ok := maskRules[mainWithSub]
	if !ok {
		// main + sub not in mask rules
		// fallback to result
		return result
	}

	maskOrder, ok := maskOrders[c.mask]
	if !ok {
		// mask is not valid
		// fallback to result
		return result
	}

	return maskRule[maskOrder]
}

// c represent ắ
// c + f = ắ + f = ằ
// return true if r is disappeared
// return false if r is still exist after plus
// example: c + s = ắ + s =  ăs
func (c *viChar) plus(r rune) bool {
	// ắ + w = (a, w, s) + w = (a, 0, s) + w = áw
	if r == c.sub {
		c.sub = 0
		return false
	}

	// ắ + s = (a, w, s) + s = (a, w, 0) + s = ăs
	if r == c.mask {
		c.mask = 0
		return false
	}

	if _, ok := subRules[string(c.main)+string(c.sub)]; ok {
		c.sub = r
		return true
	}

	return false
}
