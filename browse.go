package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	termutil "github.com/japanoise/termbox-util"
	termbox "github.com/nsf/termbox-go"
)

const topOffset int = 2

func browse(data []byte) {
	var rootNode interface{}
	json.Unmarshal(data, &rootNode)
	browseNode(rootNode, "/")
}

func drawNodeBrowser(sx, sy, scroll, sel int, breadcrumb string, list *kvList) {
	termbox.Clear(0, 0)
	termutil.Printstring(breadcrumb, 0, 0)
	for i, item := range list.items[scroll:] {
		y := topOffset + i
		if y > sy {
			break
		}
		if scroll+i == sel {
			for x := 0; x < sx; x++ {
				termutil.PrintRune(x, y, ' ', termbox.AttrReverse)
			}
			termutil.PrintstringColored(termbox.AttrReverse, item.key, 0, y)
			termutil.PrintstringColored(termbox.AttrReverse, item.value, list.keyWidth+2, y)
		} else {
			termutil.Printstring(item.key, 0, y)
			termutil.Printstring(item.value, list.keyWidth+2, y)
		}
	}
	termbox.Flush()
}

func browseNode(node interface{}, breadcrumb string) {
	list := printJSONNode(node)
	running := true
	sel := 0
	scroll := 0
	listlen := len(list.items)
	search := ""

	for running {
		if scroll < 0 {
			scroll = 0
		}
		sx, sy := termbox.Size()
		drawNodeBrowser(sx, sy, scroll, sel, breadcrumb, &list)

		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			pev := termutil.ParseTermboxEvent(ev)
			switch pev {
			case "q", "C-c":
				termbox.Close()
				os.Exit(0)
			case "C-l", "C-g", "C-b", "h", "LEFT":
				return
			case "C-n", "j", "DOWN":
				if sel < listlen-1 {
					sel++
					if (sel+topOffset)-scroll >= sy {
						scroll++
					}
				}
			case "C-p", "k", "UP":
				if sel > 0 {
					sel--
					if sel < scroll {
						scroll--
					}
				}
			case "M-v", "prior":
				sel -= sy
				if sel < 0 {
					sel = 0
				}
				if sel < scroll {
					scroll = sel
				}
			case "C-v", "next":
				sel += sy
				if sel >= listlen {
					sel = listlen - 1
				}
				if (sel+topOffset)-scroll >= sy {
					scroll = sel
				}
			case "M-<", "g":
				sel = 0
				scroll = 0
			case "M->", "G":
				sel = listlen - 1
				scroll = listlen - (sy - topOffset)
			case "C-s", "/", "n":
				for search == "" {
					search = termutil.Prompt("search term", func(ssx, ssy int) {
						termutil.ClearLine(ssx, ssy-1)
						drawNodeBrowser(ssx, ssy, scroll, sel, breadcrumb, &list)
					})
					termbox.HideCursor()
				}
				if sel < listlen-1 {
					for i, item := range list.items[sel+1:] {
						if strings.Contains(item.key, search) {
							sel += i + 1
							scroll = sel
							break
						}
					}
				}
			case "C-r", "?", "p":
				for search == "" {
					search = termutil.Prompt("reverse search term", func(ssx, ssy int) {
						termutil.ClearLine(ssx, ssy-1)
						drawNodeBrowser(ssx, ssy, scroll, sel, breadcrumb, &list)
					})
					termbox.HideCursor()
				}
				if sel > 0 {
					for i, item := range list.items[:sel] {
						if strings.Contains(item.key, search) {
							sel = i
							scroll = sel
						}
					}
				}
			case "RET", "l", "C-f", "RIGHT":
				// useless switch, but it's the only way to get this reflection to compile without whinging
				switch list.items[sel].child.(type) {
				case map[string]interface{}, []interface{}:
					browseNode(list.items[sel].child, breadcrumb+list.items[sel].key+"/")
				}
			case "C-x", "x":
				filename := termutil.Prompt("Export to", func(ssx, ssy int) {
					termutil.ClearLine(ssx, ssy-1)
					drawNodeBrowser(ssx, ssy, scroll, sel, breadcrumb, &list)
				})
				termbox.HideCursor()
				if filename == "" {
					break
				}

				file, err := os.Create(filename)
				if err != nil {
					break
				}

				for _, item := range list.items {
					fmt.Fprintf(file, "%s\t%s\n", item.key, item.value)
				}

				file.Close()
			}
			if search != "" && pev != "C-s" && pev != "/" && pev != "n" && pev != "C-r" && pev != "?" && pev != "p" {
				search = ""
			}
		}
	}
}
