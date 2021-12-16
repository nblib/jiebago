package jiebago_test

import (
	"bufio"
	"fmt"
	"github.com/wangbin/jiebago"
	"io"
	"os"
	"testing"
	"time"
)

func Example() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	print := func(ch []string) {
		for _, word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}

	fmt.Print("【全模式】：")
	print(seg.CutAll("我来到北京清华大学"))

	fmt.Print("【精确模式】：")
	print(seg.CutSync("我来到北京清华大学", false))

	fmt.Print("【新词识别】：")
	words := seg.CutSync("他来到了网易杭研大厦", true)
	for _, word := range words {
		fmt.Printf(" %s /", word)
	}
	fmt.Println()

	fmt.Print("【搜索引擎模式】：")
	print(seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true))
	// Output:
	// 【全模式】： 我 / 来到 / 北京 / 清华 / 清华大学 / 华大 / 大学 /
	// 【精确模式】： 我 / 来到 / 北京 / 清华大学 /
	// 【新词识别】： 他 / 来到 / 了 / 网易 / 杭研 / 大厦 /
	// 【搜索引擎模式】： 小明 / 硕士 / 毕业 / 于 / 中国 / 科学 / 学院 / 科学院 / 中国科学院 / 计算 / 计算所 / ， / 后 / 在 / 日本 / 京都 / 大学 / 日本京都大学 / 深造 /
}
func Example_weicheng() {
	data := "第一章红海早过了，船在印度洋面上开驶着，但是太阳依然不饶人地迟落早起，侵占去大部分的夜。夜仿佛纸浸了油变成半透明体；它给太阳拥抱住了，分不出身来，也许是给太阳陶醉了，所以夕照晚霞褪后的夜色也带着酡红。到红消醉醒，船舱里的睡人也一身腻汗地醒来，洗了澡赶到甲板上吹海风，又是一天开始。这是七月下旬，合中国旧历的三伏，一年最热的时候。在中国热得更比常年利害，事后大家都说是兵戈之象，因为这就是民国二十六年【一九三七年】。"
	var seg jiebago.Segmenter
	seg.LoadDictionary("./dict/jieba.dict.utf8")

	print := func(ch []string) {
		for _, word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}
	print(seg.CutSync(data, true))

	//Output:
	//
}
func Example_suggestFrequency() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	print := func(ch []string) {
		for _, word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}
	sentence := "超敏C反应蛋白是什么？"
	fmt.Print("Before:")
	print(seg.CutSync(sentence, false))
	word := "超敏C反应蛋白"
	oldFrequency, _ := seg.Frequency(word)
	frequency := seg.SuggestFrequency(word)
	fmt.Printf("%s current frequency: %f, suggest: %f.\n", word, oldFrequency, frequency)
	seg.AddWord(word, frequency)
	fmt.Print("After:")
	print(seg.CutSync(sentence, false))

	sentence = "如果放到post中将出错"
	fmt.Print("Before:")
	print(seg.CutSync(sentence, false))
	word = "中将"
	oldFrequency, _ = seg.Frequency(word)
	frequency = seg.SuggestFrequency("中", "将")
	fmt.Printf("%s current frequency: %f, suggest: %f.\n", word, oldFrequency, frequency)
	seg.AddWord(word, frequency)
	fmt.Print("After:")
	print(seg.CutSync(sentence, false))

	sentence = "今天天气不错"
	fmt.Print("Before:")
	print(seg.CutSync(sentence, false))
	word = "今天天气"
	oldFrequency, _ = seg.Frequency(word)
	frequency = seg.SuggestFrequency("今天", "天气")
	fmt.Printf("%s current frequency: %f, suggest: %f.\n", word, oldFrequency, frequency)
	seg.AddWord(word, frequency)
	fmt.Print("After:")
	print(seg.CutSync(sentence, false))
	// Output:
	// Before: 超敏 / C / 反应 / 蛋白 / 是 / 什么 / ？ /
	// 超敏C反应蛋白 current frequency: 0.000000, suggest: 1.000000.
	// After: 超敏C反应蛋白 / 是 / 什么 / ？ /
	// Before: 如果 / 放到 / post / 中将 / 出错 /
	// 中将 current frequency: 763.000000, suggest: 494.000000.
	// After: 如果 / 放到 / post / 中 / 将 / 出错 /
	// Before: 今天天气 / 不错 /
	// 今天天气 current frequency: 3.000000, suggest: 0.000000.
	// After: 今天 / 天气 / 不错 /
}

func Example_loadUserDictionary() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	print := func(ch []string) {
		for _, word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}
	sentence := "李小福是创新办主任也是云计算方面的专家"
	fmt.Print("Before:")
	print(seg.CutSync(sentence, true))

	seg.LoadUserDictionary("userdict.txt")

	fmt.Print("After:")
	print(seg.CutSync(sentence, true))
	// Output:
	// Before: 李小福 / 是 / 创新 / 办 / 主任 / 也 / 是 / 云 / 计算 / 方面 / 的 / 专家 /
	// After: 李小福 / 是 / 创新办 / 主任 / 也 / 是 / 云计算 / 方面 / 的 / 专家 /
}

var t string
var ta []string

func print(ch []string) {
	for _, word := range ch {
		//fmt.Printf(" %s /", word)
		t = word
	}
	//fmt.Println()
}
func TestBench(t *testing.T) {
	var seg jiebago.Segmenter
	//cpuFile, err := os.Create("./cpu.prof")
	//if err != nil {
	//	panic(err)
	//}
	//defer cpuFile.Close()

	seg.LoadDictionary("./dict/jieba.dict.utf8")

	lines := []string{}
	f, err := os.Open("./weicheng.utf8")
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			lines = append(lines, line)
		}
	}
	fmt.Println(len(lines))
	//err = pprof.StartCPUProfile(cpuFile)
	//if err != nil {
	//	panic(err)
	//}
	//minTime := math.MaxFloat64
	//for v := 0; v < 5; v++ {
	//	startTime := time.Now()
	//	for i := 0; i < 10; i++ {
	//		for j := 0; j < len(lines); j++ {
	//			seg.CutSync(lines[j], true)
	//		}
	//	}
	//	dura := time.Now().Sub(startTime).Seconds()
	//	if dura < minTime {
	//		minTime = dura
	//	}
	//}
	startTime := time.Now()
	for i := 0; i < 50; i++ {
		for j := 0; j < len(lines); j++ {
			seg.CutSync(lines[j], true)
		}
	}
	dura := time.Now().Sub(startTime).Seconds()
	//pprof.StopCPUProfile()

	fmt.Println("最佳耗时: ", dura)

}
