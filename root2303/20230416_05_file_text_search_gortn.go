package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	fmt.Println("---------------------------------")

	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 인수 필요")
		return
	}
	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}
	// fmt.Println(findInfos)
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("---------------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("---------------------------------------")
		fmt.Println()
	}
}

func PrintFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을수 없다.", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word string, path string) []FindInfo {
	findInfos := []FindInfo{}
	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다.err:", err, ", path:", path)
		return findInfos
	}
	// ----------------------------------
	// 개선전
	// for _, filename := range filelist {
	// 	findInfos = append(findInfos, FindWordInFile(word, filename))
	// }
	// ----------------------------------
	// 개선후
	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0
	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch) // 고루틴 실행
	}

	for FindInfo := range ch {
		findInfos = append(findInfos, FindInfo) // 결과수직
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	// ----------------------------------
	return findInfos
}

// 전
// func FindWordInFile(word string, filename string) FindInfo {
// 후
func FindWordInFile(word string, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을수 없다.", filename)
		ch <- findInfo // *
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}

// D:\200.dev\__git_repo\go\root2303>go run "d:\200.dev\__git_repo\go\root2303\20230416_04_file_text_search.go" love ../_txt/hamlet.txt
// ---------------------------------
// ../_txt/hamlet.txt
// ---------------------------------------
//          253         As needful in our loves, fitting our duty?
//          379         And with no less nobility of love
//          472       Ham. For God's love let me hear!
//          538         I will requite your loves. So, fare you well.
//          542       Ham. Your loves, as mine to you. Farewell.
//          574         Grows wide withal. Perhaps he loves you now,
//          584         Whereof he is the head. Then if he says he loves you,
//          675       Oph. My lord, he hath importun'd me with love
//          855         If thou didst ever thy dear father love-
//          862         As meditation or the thoughts of love,
//          882         From me, whose love was of that dignity
//          1038        With all my love I do commend me to you;
//          1040        May do t' express his love and friending to you,
//          1152      Pol. Mad for thy love?
//          1171        This is the very ecstasy of love,
//          1189        More grief to hide than hate to utter love.
//          1338                But never doubt I love.
//          1340        reckon my groans; but that I love thee best, O most best, believe
//          1350        Receiv'd his love?
//          1354        When I had seen this hot love on the wing
//          1360        Or look'd upon this love with idle sight?
//          1391        Mark the encounter. If he love her not,
//          1421        for love- very near this. I'll speak to him again.- What do you
//          1511        obligation of our ever-preserved love, and by what more dear a
//          1515      Ham. [aside] Nay then, I have an eye of you.- If you love me, hold
//          1540        target; the lover shall not sigh gratis; the humorous man shall
//          1622               The which he loved passing well.'
//          1627        love passing well.
//          1884        If't be th' affliction of his love, or no,
//          1928        The pangs of despis'd love, the law's delay,
//          1971        but now the time gives it proof. I did love you once.
//          1974        inoculate our old stock but we shall relish of it. I loved you
//          2035        Sprung from neglected love.- How now, Ophelia?
//          2208        his love.
//          2230      Ham. As woman's love.
//          2238          Since love our hearts, and Hymen did our hands,
//          2241          Make us again count o'er ere love be done!
//          2246          For women's fear and love holds quantity,
//          2248          Now what my love is, proof hath made you know;
//          2249          And as my love is siz'd, my fear is so.
//          2250          Where love is great, the littlest doubts are fear;
//          2251          Where little fears grow great, great love grows there.
//          2252        King. Faith, I must leave thee, love, and shortly too;
//          2258          Such love must needs be treason in my breast.
//          2265          Are base respects of thrift, but none of love.
//          2283          That even our loves should with our fortunes change;
//          2285          Whether love lead fortune, or else fortune love.
//          2288          And hitherto doth love on fortune tend,
//          2335      Ham. I could interpret between you and your love, if I could see
//          2354        shall see anon how the murtherer gets the love of Gonzago's wife.
//          2421      Ros. My lord, you once did love me.
//          2437      Guil. O my lord, if my duty be too bold, my love is too unmannerly.
//          2665        From the fair forehead of an innocent love,
//          2691        You cannot call it love; for at your age
//          2718        Stew'd in corruption, honeying and making love
//          2780        Would gambol from. Mother, for love of grace,
//          2888        This mad young man. But so much was our love
//          3038        And, England, if my love thou hold'st at aught,-
//          3170             How should I your true-love know
//          3195               With true-love showers.
//          3347        Nature is fine in love, and where 'tis fine,
//          3349        After the thing it loves.
//          3363      Oph. There's rosemary, that's for remembrance. Pray you, love,
//          3485        Is the great love the general gender bear him,
//          3501        I lov'd your father, and we love ourself,
//          3592      King. Not that I think you did not love your father;
//          3593        But that I know love is begun by time,
//          3596        There lives within the very flame of love
//          3757           In youth when I did love, did love,
//          3976        Could not (with all their quantity of love)
//          3979      Queen. For love of God, forbear him!
//          4062        As love between them like the palm might flourish,
//          4080      Ham. Why, man, they did make love to this employment!
//          4266        I do receive your offer'd love like love,
// ---------------------------------------
