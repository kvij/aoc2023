package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kvij/aoc/internal/app"
)

type number struct {
	i    int
	used bool
}

type coordinate struct {
	x, y int
}

const input = `
.......................153..988....502..842.........588.....441.468......481..........314...715.57............................163..992..512.
............805............*......#.............%...............*........=......%......................#......*.............-....#....*.....
........914.........#...617..201.........271.....671......52..898................847..........*230..215......393..%751....537...............
..........#......361..........*...........*............-4.............165..609........922..133...........706..................*....552*127..
490*..........................350...*...664........806................../..*...514.31...=........../25....%.................83..............
....245...............805...........467.......449...+..............313*....115....*.......611.343............$...237..229...................
........................*.....150.............*..............8.........511...................*........#.....837.*......*....*...............
435.................688.8..............$......330.................474.......9.......736...........*...787........336...245...446.....916....
...*.*....920........%................666..90.....786$..221.......*.....173.%.......*........@..25.98.....186#.......................*......
.718.120..@.............$.931................*...........*....15...10....*.........649....&.542.....................-.......@.......106.....
.............-.......794...................486.........209....*..2........252..........517...................250.519......274...............
...636.....500........................878.......106.........487..*....819........353.........678.........../..-...............539...........
.............................43......*...........$.......44.....881....*...662*...%................407#...612.......452@.286................
..................................195..543.../.......75..............887.......43...396..................................=.............813..
........$.735......590.998.............=...862............228......................*................+...%...........979.....500=......*.....
......597..........=...+.........61..................878../....596......*........542../815....567.834..672.=378.....*...&.........#..981....
...................................*618..998.........*...............716..........................................495.682..62..945..........
.......915..............156.................$.......298........479........170............985............+90..../.............@.....&7.......
...800*.............246*.........370..............................-.........*....639....*.....362.368.......818................687..........
.........$.173............383=......*........389......................273....785.....514....=...*....*...........@..............*....446*833
.433...227...*...-.-...............770.988.................738.523...*....................397....599..499...&..385.........*199.35..........
...%........40.437..300../....740.........*...................*...........................................192......&....272............*....
........#.................305.......318.608..609....*..............448......@...959*604....16*...../..............131............=..325.215.
.980...85....*48......883.............@..............663.................128..................944.420...203.....................381.........
..........887............*....434........725................958......367......178..........................@.779.........................986
.......................591...*......964.*....713.140..........*..698....*......................772.......................914........223.....
...........960.....*....................80.....#..%....456..824....*...138...939.........362*......488...572........663..&..........*.......
.......+.....*...645.............%12.....................*.........692...............922.....902..@.................*........27..542........
...792..242.702.........256..668..............613*479...329...................134#....*...................*545...375........................
.......................&......*.............&..........................298.........696.................844.....=....................184.....
......886=..................129....187.......850..*.....&................*...............358...130............500................@.....*....
.............*54..112.............*..............791..215...787.........437..889..............*........584........../226......366...483.....
.......47..........*....826.500..48......665..................*.................*.377....#...49.......$.....................................
.196...*..........815......*..............*.....*...........571........297*...153.*....416...............+.....73...126.............111.....
.......164............547#................69.503.743............108...............22......../70....36.....470................516....*.......
....................................634................&.357......*...................275.........+................592......%....52.204.....
....#943....843............544.......-.....96*139...640........694................@.........558......................*...........*.......104
..............*....731.......*...382...................................../...%557..857......*....-................716..638./165...338.......
......926....342../..........861............/......................#...768...................569..192...................*...................
..272....................500.......*.........784......558..38.....786..........238.932%................/...827....841............76.........
.....*25....=..780.74...........457....%..%.......806........*........&........*...........+..........53..*..........%.............*........
..99.....697...$....+................601.757.600...........751........801......795..998....362...........431......&............797.860......
....*...................50.579...............&...-.878.....................586.....*............587..216.......321...783....................
...859.................@.....+.317/............34../...499.....904.991.......$....938......337@...*..#.................*.............130*312
............707.266.....................156..9............*.......*..............................880.......787.........102..................
.........*.....*................938*743...*.*..............195.........425..........869*.....448..........*.....73.689.....48..469..........
......568.164....938........................964.......899......641.....*........930.............*...891....104........*...&.......*719......
....................*..................#...........+..-...381......36.749...744..*......659...935......+..............694.....370...........
........174........460............+..435........893........%................+.....500......$.............#......397.............*...549.677.
....981.*.....&355.....960.......963..................543#.........................................971...483....%............428.....=......
354..*................*.....752..................69........676..54.......531*......78....877..-950....%............346...........699........
......265.....#......959.................846.......*...93............263.....974....=.......*............&........................*.........
............857...........527............../.......29....@.846.......*.....................768..134&...571..........302...........512.......
....$254.........397......*............./....................%....128.........*984.....693..................991...+...*.@....@...........228
............335...*.....751......626.699............$118..............991..............$....................*...327.93..508.301......348....
....387........@.266..............*.......................-...........*.......%...................260....756................................
....*.....817..................726...391.....112..........273..997...197....667......./......932*....$........290...............711......560
...488.....&....961.....170.........*..........*................*.........*........883..825......71..............&...215...........*.766....
...............+.........#....$.....600.&.......674.42.................730.998.........*..................139.........*.........461....=....
...#97...........817........86..........155..............759.....200.............827...163..........................499.....................
.....................#.........................793*722...........................$..............&....725&......................442..........
.............812....613.....892..........................464......430..%.....462.....%.......599.............555..204....219......*637......
...544........./............&...............................*.......*.840.......*...994.797........52.........*....*..........739...........
....=..................................713.304...161........186.................824.......-....=....*.......118.844..113......*.....872.....
......................75.963......244.....*.....*.......466..........590....................208....363...............*.....872..............
.............391.....*...............*.......843......#..........184..-.................................441......448.597....................
979......416......114................306.%.........719.....792......$...........................#..........*..............603.989.914*......
.........-..................975..670.....890...621...........*..........145..659...369=.@978....935........938...............*........107...
................259..%632....=...................*......173...871........*...................4..........*..........&108.........*281........
...............*................326.....732...528.../....#.............698.............316..*........188.888.................918............
..512........507...............*...........*........440.........................938......./.254..313...............................422......
....*............300.........430..354$.....884.701.......%....368............../....@480...........*..142.............305...145.61......106.
.363..939.639..................................$...../....704.....835@...-776....................655.........645......*......*....*892......
.........*..........*338................../........741........................113..895...............%..403.....*..40..992..2...............
.................986...........927..%449.17............................182....*.....=......870......899........665.#.................949.189
...841.........................*.......................................*...849...........=....*295.........420..............441.882.........
......$..+.............124...993..........*...........*................365..........79.364................*.............414*.....-.....360..
........159...........*...........457$....199......960.967...&...................................+......181.....+.......................*...
.....................859...................................962.............912....#......330.....168...........904.....958*....213*583..49..
........%.........*..........421..340*196..892.........409.................*....488.645...*..........#......................................
.....802........116...........*...........................*...227........791.........$....203.....963..418................599+..........744.
..............#.....-966.23.94...58..283...........476....722.......730...........................................@..700.......%788.........
240%..........844..................*.-.........486............307............#49...288........&.329.............61..*.....................95
.........#............148..446...........402/..........256.......................&..........573..*...................319.......763..........
......936..............*...%.................................-....709...371...955..352............553.629....445..................*.........
.121......445.347...455.......=....#....92..................367..*......*............@........881........*..*....267/.....-89......843.*782.
......519.=.....*.........401.996.452...-.......-387.886.........982...149..910%..........81..-.........653..407......819.......@...........
.....*.......950...448#....*...............&691......$......................................*....538..............681*.........560..........
.....25....................60............*......................304*863...../.............103......*.........811..........923.......426.....
.800..................../................974..............255..............958....738..............465.........*......181...+........./.....
....*.....-......824....782........90...........+505.......*....903...................544.712.826...........800..........*..................
...408..123...-.........................................557.......*....................@.....*....561.743............310..642...97*475......
...............900.#494.....915.624............../..............508.....431.*500.........986...../......=....187........*...................
...........................*......*....#741.@957.979.*................/..%................%.............................218.805......505*...
...........................59.....595.................948...+........885.....*808...554.........70......827=.............................690
...95........+......455.....................................776...........795......=..............&...$........349...............621*80.....
...*....../.793..$...#.....................842..........898.........729...................698.......273...704...*.......451..%..............
.......340......911......865.....876..........-...........*...%......*....676*498..424...-.......=.........*........241..#...967.......323..
........................-...........*271..578...........734..8..692..435....................698.548.................@...................*...
.....=......*..................131...........=..=..............*.............523.659...221...+........................468..263..395.408..757
.....236....661...565*...........#...............228............287...448.......*......&......................=.........#.+........*........
......................982.................$..........403*317.............*103.................454...#...311....588...................105....
...72...............................246..991...../...........12................%.........271...=....668.................................*...
.......348...259......$..../384.221...............873.......*.....560....&...686............*..................241............763........884
........*...+..........848.................=........................*.....76......*....889..417.404...............&...538.663..=......61....
.......446....948..805.....615.....303+.372..........893.......945...849.........861....*........*...187.980...&....../...........838.+.....
...453...............*....*.....................357...../.........+....................506....129..@.......*..633.822............*..........
......*188........495....732............635*.....$............219........65....475................811...231..........*..332%...212..........
.................................*..........575.....-....591...*...420.........*...............................897.317.............165...762
.............764..410.............438....*..........588....*.245..%......*.....468..........32..-714...456.........................*....*...
..............*.......................516.552...366......416........166...552.................*.............../..126.........396.386....921.
.......&.............67&........468................*.............................403.221.......433.....411..299..........457*...............
.....539.......*287.............*............657/.457..............614........34*......*.....$......12*.........693.........................
............908......@602..176..358..778...............72...-..............$..........161.579..................*................571*602.....
...........................*............*799...........*.....118........788....689................$....../702..493.......117............%...
.................337..496..504.......................519............#..........=....744........782...................625.*.......499.....576
.......459..%..............................................28.....398...............*..............*19....#935.........$..845......#........
.............707.....................992........289...481+...*..................3..234....822$.......................................#......
..513.914*.........161......664.....*..........@..............922......$.........*......................330..............@...........798....
..........378.................*..258...............219+.313.........836.....=226.566.....748....&...147*......872.......786..538.446........
.....*827...........907....844.............684..........@..................................*..263............../....8..............=...*....
835........$...88......*.........308.........*.....161.....621*327..-184....333.........@.473.....205..................................872..
...*....340.............744........-.........961....../........................*449...740........*................867.427..251..............
..988........@...................................................765..........................743......228...............*........942.......
.............125..#...........406.........%.....837.................*.....609*.........%..........644...=..465*643....880..509...*.......453
....544..........480.............=.........274.....*..148..583...821..........880.....492....535...@.......................*............*...
.....$.......306................................287....#..*...........647.....................$........-......166.714....942..675.....657...
................*....................673...780...........620..........*.........&..........=......997.465..........$.............*..........
.456........920..855.913..............*.../.....325*316.......6.......208....486........715.....=...%.............................67........
......*178..+.............686.59....871.................$......*.............................149..............238........634................
....56.........385...648*...............*968..&.......732.....921.467........*.........479.........695*284.../..........*...................
............56.....$.....675..746#.............384.................*......451.709................................%..939..265.........290*891
................501.......................620.............900.......825.................258...387......21*40...273...*......................
.......532..................................*....*....956....=....................711.........&.....................415.307....411......9...
..............................314../......692.214.718.............762*461.....844*.....&.............973.675...80.................*143......
...........365.......460.329.@......476....................................*..........926.....799...*.............-.........................
..............*.297...*...*.........................753........700......473.765..............*....967........111.629.932....125.............
.......+...261...&.......786.283....695.....486.....*......565.+................536...../....380....../124..%........*.........%............
......339..............*..............*....*.......996.......-.....+..............*......752......................@...............-.........
.......................716...........551.631........................279..555.....373...................691......114.215............515......
`

const gridSize = 140

func main() {
	// Parse lines into coordinate number lookup matrix and symbol coordinate list.
	// The matrix should hold pointers to number struct
	//
	// Loop over symbol coordinates and find adjacent number(s) that are not used yet
	matrix := [gridSize][gridSize]*number{}
	symbolCoordinates := make([]coordinate, 0, gridSize/4)
	asterixCoordinates := make([]coordinate, 0, gridSize/6)

	var lineNo int
	err := app.ProcessLines(input, func(line string) error {
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '.':
				continue
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				l := numLen(line[i:])
				n, err := strconv.Atoi(line[i : i+l])
				if err != nil {
					return err
				}
				num := &number{
					i: n,
				}
				for it := 0; it < l; it++ {
					matrix[lineNo][i+it] = num
				}
				i += l - 1
			case '*': // For part two
				asterixCoordinates = append(asterixCoordinates, coordinate{
					x: i,
					y: lineNo,
				})
				fallthrough
			default:
				symbolCoordinates = append(symbolCoordinates, coordinate{
					x: i,
					y: lineNo,
				})
			}
		}

		lineNo++
		return nil
	})

	if err != nil {
		log.Fatalln("Failed to parse lines", err)
	}

	partOne(matrix, symbolCoordinates)
	partTwo(matrix, asterixCoordinates)
}

func partOne(matrix [gridSize][gridSize]*number, symbolCoordinates []coordinate) {
	var answer int
	find := func(x, y int) int {
		// Out of bounds
		if x < 0 || x > gridSize || y < 0 || y > gridSize {
			return 0
		}

		num := matrix[y][x]
		if num == nil || num.used {
			return 0
		}

		num.used = true
		return num.i
	}

	for _, sc := range symbolCoordinates {
		surounding := []coordinate{
			// same row left, right
			{sc.x - 1, sc.y},
			{sc.x + 1, sc.y},
			// row above, left, above, right
			{sc.x - 1, sc.y - 1},
			{sc.x, sc.y - 1},
			{sc.x + 1, sc.y - 1},
			// row below, left, below, right
			{sc.x - 1, sc.y + 1},
			{sc.x, sc.y + 1},
			{sc.x + 1, sc.y + 1},
		}

		for _, c := range surounding {
			answer += find(c.x, c.y)
		}
	}

	fmt.Println("The answer is:", answer)
}

func partTwo(matrix [gridSize][gridSize]*number, asterixCoordinates []coordinate) {
	find := func(x, y int) *number {
		// Out of bounds
		if x < 0 || x > gridSize || y < 0 || y > gridSize {
			return nil
		}

		return matrix[y][x]
	}
	var answer int
	for _, sc := range asterixCoordinates {
		found := make(map[*number]struct{}, 4)

		surounding := []coordinate{
			// same row left, right
			{sc.x - 1, sc.y},
			{sc.x + 1, sc.y},
			// row above, left, above, right
			{sc.x - 1, sc.y - 1},
			{sc.x, sc.y - 1},
			{sc.x + 1, sc.y - 1},
			// row below, left, below, right
			{sc.x - 1, sc.y + 1},
			{sc.x, sc.y + 1},
			{sc.x + 1, sc.y + 1},
		}

		for _, c := range surounding {
			num := find(c.x, c.y)
			if num != nil {
				found[num] = struct{}{}
			}
		}

		if len(found) == 2 {
			set := Keys(found)
			answer += set[0].i * set[1].i
		}
	}

	fmt.Println("The answer is:", answer)
}

func numLen(s string) (length int) {
	for i, v := range s {
		if v < '0' || v > '9' {
			return length
		}
		length = i + 1
	}

	return length
}

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}