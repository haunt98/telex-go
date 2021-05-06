package telex

var (
	subRules = map[string]rune{
		"aw": 'ă',
		"aa": 'â',
		"dd": 'đ',
		"ee": 'ê',
		"oo": 'ô',
		"ow": 'ơ',
		"uw": 'ư',
	}

	maskRules = map[rune][]rune{
		'a': {'a', 'à', 'á', 'ả', 'ã', 'ạ'},
		'ă': {'ă', 'ằ', 'ắ', 'ẳ', 'ẵ', 'ặ'},
		'â': {'â', 'ầ', 'ấ', 'ẩ', 'ẫ', 'ậ'},
		'e': {'e', 'è', 'é', 'ẻ', 'ẽ', 'ẹ'},
		'ê': {'ê', 'ề', 'ế', 'ể', 'ễ', 'ệ'},
		'i': {'i', 'ì', 'í', 'ỉ', 'ĩ', 'ị'},
		'o': {'o', 'ò', 'ó', 'ỏ', 'õ', 'ọ'},
		'ô': {'ô', 'ồ', 'ố', 'ổ', 'ỗ', 'ộ'},
		'ơ': {'ơ', 'ờ', 'ớ', 'ở', 'ỡ', 'ợ'},
		'u': {'u', 'ù', 'ú', 'ủ', 'ũ', 'ụ'},
		'ư': {'ư', 'ừ', 'ứ', 'ử', 'ữ', 'ự'},
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
	if c.sub == 0 {
		return c.main
	}

	result, ok := subRules[string(c.main)+string(c.sub)]
	if !ok {
		// main + sub is not valid
		// fallback to main
		return c.main
	}

	if c.mask == 0 {
		return result
	}

	maskRule, ok := maskRules[result]
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
// return true if plus do not cause side effect
// otherwise return false
// side effect: c + s = ắ + s =  ăs
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

	return false
}
