package app

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

var ErrNoMoreUUIDs = errors.New("you reached the limit of uuids ")

type Uniquer struct {
	source []string
	mutex  sync.Mutex
}

func (u *Uniquer) shuffle() {
	u.source = list
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(u.source), func(i, j int) {
		u.source[i], u.source[j] = u.source[j], u.source[i]
	})
}

func (u *Uniquer) New() (string, error) {
	if u.source == nil {
		u.shuffle()
	}
	u.mutex.Lock()
	defer u.mutex.Unlock()
	last := len(u.source) - 1
	if last == -1 {
		return "", ErrNoMoreUUIDs
	}
	value := u.source[last]
	u.source[last] = ""
	u.source = u.source[:last]
	return value, nil
}

var list = []string{
	"black",
	"navy",
	"darkblue",
	"mediumblue",
	"blue",
	"darkgreen",
	"green",
	"teal",
	"darkcyan",
	"deepskyblue",
	"lime",
	"springgreen",
	"aqua",
	"cyan",
	"dodgerblue",
	"forestgreen",
	"seagreen",
	"limegreen",
	"turquoise",
	"royalblue",
	"steelblue",
	"indigo",
	"cadetblue",
	"dimgray",
	"dimgrey",
	"slateblue",
	"olivedrab",
	"slategray",
	"slategrey",
	"lawngreen",
	"chartreuse",
	"aquamarine",
	"maroon",
	"purple",
	"olive",
	"gray",
	"grey",
	"skyblue",
	"blueviolet",
	"darkred",
	"darkmagenta",
	"saddlebrown",
	"darkseagreen",
	"lightgreen",
	"mediumpurple",
	"darkviolet",
	"palegreen",
	"darkorchid",
	"yellowgreen",
	"sienna",
	"brown",
	"darkgray",
	"darkgrey",
	"lightblue",
	"greenyellow",
	"powderblue",
	"firebrick",
	"darkgoldenrod",
	"mediumorchid",
	"rosybrown",
	"darkkhaki",
	"silver",
	"indianred",
	"peru",
	"chocolate",
	"tan",
	"lightgray",
	"lightgrey",
	"thistle",
	"orchid",
	"goldenrod",
	"crimson",
	"gainsboro",
	"plum",
	"burlywood",
	"lightcyan",
	"lavender",
	"darksalmon",
	"violet",
	"lightcoral",
	"khaki",
	"aliceblue",
	"honeydew",
	"azure",
	"sandybrown",
	"wheat",
	"beige",
	"whitesmoke",
	"mintcream",
	"ghostwhite",
	"salmon",
	"antiquewhite",
	"linen",
	"oldlace",
	"red",
	"fuchsia",
	"magenta",
	"deeppink",
	"orangered",
	"tomato",
	"hotpink",
	"coral",
	"darkorange",
	"lightsalmon",
	"orange",
	"lightpink",
	"pink",
	"gold",
	"peachpuff",
	"navajowhite",
	"moccasin",
	"bisque",
	"mistyrose",
	"blanchedalmond",
	"papayawhip",
	"lavenderblush",
	"seashell",
	"cornsilk",
	"lemonchiffon",
	"floralwhite",
	"snow",
	"yellow",
	"lightyellow",
	"ivory",

	"afghanistan",
	"albania",
	"algeria",
	"america",
	"andorra",
	"angola",
	"antigua",
	"argentina",
	"armenia",
	"australia",
	"austria",
	"azerbaijan",
	"bahamas",
	"bahrain",
	"bangladesh",
	"barbados",
	"belarus",
	"belgium",
	"belize",
	"benin",
	"bhutan",
	"bissau",
	"bolivia",
	"bosnia",
	"botswana",
	"brazil",
	"british",
	"brunei",
	"bulgaria",
	"burkina",
	"burma",
	"burundi",
	"cambodia",
	"cameroon",
	"canada",
	"cape_verde",
	"chad",
	"chile",
	"china",
	"colombia",
	"comoros",
	"congo",
	"costa rica",
	"croatia",
	"cuba",
	"cyprus",
	"czech",
	"denmark",
	"djibouti",
	"dominica",
	"east_timor",
	"ecuador",
	"egypt",
	"el_salvador",
	"emirate",
	"england",
	"eritrea",
	"estonia",
	"ethiopia",
	"fiji",
	"finland",
	"france",
	"gabon",
	"gambia",
	"georgia",
	"germany",
	"ghana",
	"greece",
	"grenada",
	"grenadines",
	"guatemala",
	"guinea",
	"guyana",
	"haiti",
	"herzegovina",
	"holland",
	"honduras",
	"hungary",
	"iceland",
	"india",
	"indonesia",
	"iran",
	"iraq",
	"ireland",
	"israel",
	"italy",
	"ivory coast",
	"jamaica",
	"japan",
	"jordan",
	"kazakhstan",
	"kenya",
	"kiribati",
	"korea",
	"kosovo",
	"kuwait",
	"kyrgyzstan",
	"laos",
	"latvia",
	"lebanon",
	"lesotho",
	"liberia",
	"libya",
	"lithuania",
	"luxembourg",
	"macedonia",
	"madagascar",
	"malawi",
	"malaysia",
	"maldives",
	"mali",
	"malta",
	"marshall",
	"mauritania",
	"mauritius",
	"mexico",
	"micronesia",
	"moldova",
	"monaco",
	"mongolia",
	"montenegro",
	"morocco",
	"mozambique",
	"myanmar",
	"namibia",
	"nauru",
	"nepal",
	"netherlands",
	"new_zealand",
	"nicaragua",
	"niger",
	"nigeria",
	"norway",
	"oman",
	"pakistan",
	"palau",
	"panama",
	"papua",
	"paraguay",
	"peru",
	"philippines",
	"poland",
	"portugal",
	"qatar",
	"romania",
	"russia",
	"rwanda",
	"samoa",
	"san_marino",
	"sao_tome",
	"saudi_arabia",
	"scotland",
	"scottish",
	"senegal",
	"serbia",
	"seychelles",
	"sierra leone",
	"singapore",
	"slovakia",
	"slovenia",
	"solomon",
	"somalia",
	"south_africa",
	"south_sudan",
	"spain",
	"srilanka",
	"kitts",
	"lucia",
	"sudan",
	"suriname",
	"swaziland",
	"sweden",
	"switzerland",
	"syria",
	"taiwan",
	"tajikistan",
	"tanzania",
	"thailand",
	"tobago",
	"togo",
	"tonga",
	"trinidad",
	"tunisia",
	"turkey",
	"turkmenistan",
	"tuvalu",
	"uganda",
	"ukraine",
	"uk",
	"us",
	"uruguay",
	"usa",
	"uzbekistan",
	"vanuatu",
	"vatican",
	"venezuela",
	"vietnam",
	"wales",
	"welsh",
	"yemen",
	"zambia",
	"zimbabwe",
}
var col = []string{
	"black",
	"navy",
	"darkblue",
	"mediumblue",
	"blue",
	"darkgreen",
	"green",
	"teal",
	"darkcyan",
	"deepskyblue",
	"darkturquoise",
	"mediumspringgreen",
	"lime",
	"springgreen",
	"aqua",
	"cyan",
	"midnightblue",
	"dodgerblue",
	"lightseagreen",
	"forestgreen",
	"seagreen",
	"darkslategray",
	"darkslategrey",
	"limegreen",
	"mediumseagreen",
	"turquoise",
	"royalblue",
	"steelblue",
	"darkslateblue",
	"mediumturquoise",
	"indigo",
	"darkolivegreen",
	"cadetblue",
	"cornflowerblue",
	"rebeccapurple",
	"mediumaquamarine",
	"dimgray",
	"dimgrey",
	"slateblue",
	"olivedrab",
	"slategray",
	"slategrey",
	"lightslategray",
	"lightslategrey",
	"mediumslateblue",
	"lawngreen",
	"chartreuse",
	"aquamarine",
	"maroon",
	"purple",
	"olive",
	"gray",
	"grey",
	"skyblue",
	"lightskyblue",
	"blueviolet",
	"darkred",
	"darkmagenta",
	"saddlebrown",
	"darkseagreen",
	"lightgreen",
	"mediumpurple",
	"darkviolet",
	"palegreen",
	"darkorchid",
	"yellowgreen",
	"sienna",
	"brown",
	"darkgray",
	"darkgrey",
	"lightblue",
	"greenyellow",
	"paleturquoise",
	"lightsteelblue",
	"powderblue",
	"firebrick",
	"darkgoldenrod",
	"mediumorchid",
	"rosybrown",
	"darkkhaki",
	"silver",
	"mediumvioletred",
	"indianred",
	"peru",
	"chocolate",
	"tan",
	"lightgray",
	"lightgrey",
	"thistle",
	"orchid",
	"goldenrod",
	"palevioletred",
	"crimson",
	"gainsboro",
	"plum",
	"burlywood",
	"lightcyan",
	"lavender",
	"darksalmon",
	"violet",
	"palegoldenrod",
	"lightcoral",
	"khaki",
	"aliceblue",
	"honeydew",
	"azure",
	"sandybrown",
	"wheat",
	"beige",
	"whitesmoke",
	"mintcream",
	"ghostwhite",
	"salmon",
	"antiquewhite",
	"linen",
	"lightgoldenrodyellow",
	"oldlace",
	"red",
	"fuchsia",
	"magenta",
	"deeppink",
	"orangered",
	"tomato",
	"hotpink",
	"coral",
	"darkorange",
	"lightsalmon",
	"orange",
	"lightpink",
	"pink",
	"gold",
	"peachpuff",
	"navajowhite",
	"moccasin",
	"bisque",
	"mistyrose",
	"blanchedalmond",
	"papayawhip",
	"lavenderblush",
	"seashell",
	"cornsilk",
	"lemonchiffon",
	"floralwhite",
	"snow",
	"yellow",
	"lightyellow",
	"ivory",
}
