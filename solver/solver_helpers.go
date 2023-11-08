package solver


var sizeToString = map[int]string{
	2: "pair",
	3: "triple",
	4: "quadruple",
	5: "quintuple",
}

func isInArray(n int, arr []int) bool {
	for _, v := range arr {
		if n == v {
			return true
		}
	}
	return false
}

var Easy []string = []string{
	"32168....64721.58.....3.1.65...7....49...176....9...5..74..8.311..34..7883.19....",
}
var Medium []string = []string{
	"1..4..5....8.62...764.389....56...4..83...7.1...913..5....4....3......582....1.7.",
}
var Hard []string = []string{
	"..2.945.....1..82.3....7.4.1..7....2....51..46...4..85..7..2.9............4.....3",
	".1....7..7...84..6......83..49...2..8......1.53.6.29.49...1..78.8....4..3...2...1",
	".935.87.4.6...251......7....127......4.....3.35.2....7.8...3........19....592.1..",
}
var Expert []string = []string{
	"......5...7.9....2.4.68.1......32.....6.9...45.3..4........1..67.......1..9....7.",
	"76.1..5.85...74....3..........658..............3...2.4.7.2........5....6.....3.91",
	".8.....9..1..863.2...31......4..............5...261..4...54...63.9...8..2........",
}

var NakedPairTest = []string{
	"84.2...91...1694.81.94...57718324.6...2..6..............169..8..9.8..1..58.741...",
	".9.1..2......9..6.5........2...67.9.1....3.....3.1.5.8.54.2.8....2..89.7.7....6..",
}

var HiddenSetTest = []string{
	".49132....81479...327685914.96.518...75.28....38.46..5853267...712894563964513...",
	"5286...4913649..257942.563....1..2....78263....25.9.6.24.3..9768.97.2413.7.9.4582",
}