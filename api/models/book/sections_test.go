package book

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestSection(t *testing.T) {
	// addDaqing(t)
	// addEryue(t)
	var bid int32 = NewBook().GetBookByName("大秦帝国").ID
	// var bid int32 = NewBook().GetBookByName("康雍乾帝王").ID
	s := NewSection()

	Convey("测试小节操作", t, func() {
		ss1 := s.GetSectionsByBid(bid)
		So(len(ss1), ShouldBeGreaterThan, 0)
		maxIdx := s.GetMaxIdxByBid(bid)
		So(maxIdx, ShouldBeGreaterThan, 0)
		sContent1 := s.GetById(ss1[0].ID, bid)
		So(sContent1, ShouldNotBeNil)
	})
}

func addDaqing(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	file := path.Dir(filename) + "/../../conf/01novel.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	type sec struct {
		chapter string
		list    []struct {
			title   string
			content string
		}
	}
	var scc []sec

	arr := strings.Split(string(content), "\n")
	var chapter string
	var secTemp sec
	var sectionPrefix string
	for _, v := range arr {
		if matched, err := regexp.MatchString("^ *第[一|二|三|四|五|六|七|八|九|十|百|千|零]{1,}部 .*?", v); err != nil {
			panic(err)
		} else if matched {
			if chapter != v {
				if chapter != "" {
					var se sec
					se.chapter = secTemp.chapter
					se.list = make([]struct {
						title   string
						content string
					}, len(secTemp.list))
					copy(se.list, secTemp.list)
					scc = append(scc, se)
				}

				chapter = v
				secTemp.chapter = v
				secTemp.list = nil
				sectionPrefix = ""
			}
			continue
		}
		if secTemp.chapter == "" {
			continue
		}

		if matched, err := regexp.MatchString("^ *第[一|二|三|四|五|六|七|八|九|十|百|千|零]{1,}章 .*?", v); err != nil {
			panic(err)
		} else if matched {
			sectionPrefix = v
			continue
		}

		if matched, err := regexp.MatchString("^ *[一|二|三|四|五|六|七|八|九|十|百|千|零]{1,}、.*|^楔子", v); err != nil {
			panic(err)
		} else if matched {
			// t.Log(secTemp)
			secTemp.list = append(secTemp.list, struct {
				title   string
				content string
			}{title: sectionPrefix + " " + v, content: ""})
			// t.Log(chapter + " " + sectionPrefix + " " + v)
			continue
		}

		if len(secTemp.list) > 0 {
			secTemp.list[len(secTemp.list)-1].content += strings.Trim(v, "\r") + "\r\n"
		}
	}
	scc = append(scc, secTemp)
	// t.Log(scc)

	var bid int32 = NewBook().GetBookByName("大秦帝国").ID
	for _, v := range scc {
		cid, err := NewChapter().Add(bid, v.chapter, "")
		if err != nil {
			panic(err)
		}
		for _, v1 := range v.list {

			_, err := NewSection().Add(bid, cid, v1.title, v1.content, 1, 0)
			if err != nil {
				t.Log(utf8.RuneCountInString(v1.content), v1.title, v.chapter)
			}

		}
	}
}

func addEryue(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	file := path.Dir(filename) + "/../../conf/02novel.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	type sec struct {
		chapter string
		list    []struct {
			title   string
			content string
		}
	}
	var scc []sec

	arr := strings.Split(string(content), "\n")
	var chapter string
	var secTemp sec
	for _, v := range arr {
		// t.Log(v)
		if matched, err := regexp.MatchString("^卷[一|二|三|四|五|六|七|八|九|十|百|千|零]{1,} .*?", v); err != nil {
			panic(err)
		} else if matched {
			if chapter != v {
				if chapter != "" {
					var se sec
					se.chapter = secTemp.chapter
					se.list = make([]struct {
						title   string
						content string
					}, len(secTemp.list))
					copy(se.list, secTemp.list)
					scc = append(scc, se)
				}
				chapter = v
				secTemp.chapter = v
				secTemp.list = nil

			}
			continue
		}
		if secTemp.chapter == "" {
			continue
		}

		if matched, err := regexp.MatchString("^ *第[一|二|三|四|五|六|七|八|九|十|百|千|零]{1,}回.*?", v); err != nil {
			panic(err)
		} else if matched {
			v = strings.Replace(v, "回", "回 ", 1)
			ttt := []rune(v)
			tttt := ttt[: len(ttt)-8 : len(ttt)-8]
			tttt = append(tttt, ' ')
			tttt = append(tttt, ttt[len(ttt)-8:]...)
			v = string(tttt)
			// t.Log(secTemp)
			secTemp.list = append(secTemp.list, struct {
				title   string
				content string
			}{title: v, content: ""})
			// t.Log(chapter+" "+v)
			continue
		}

		if len(secTemp.list) > 0 {
			secTemp.list[len(secTemp.list)-1].content += strings.Trim(v, "\r") + "\r\n"
		}
	}
	scc = append(scc, secTemp)
	// t.Log(scc)

	// return

	var bid int32 = NewBook().GetBookByName("康雍乾帝王").ID
	for _, v := range scc {
		cid, err := NewChapter().Add(bid, v.chapter, "")
		if err != nil {
			panic(err)
		}
		for _, v1 := range v.list {
			_, err := NewSection().Add(bid, cid, v1.title, v1.content, 1, 0)
			if err != nil {
				t.Log(utf8.RuneCountInString(v1.content), v1.title, v.chapter)
			}

		}
	}
}
