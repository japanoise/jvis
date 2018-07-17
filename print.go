package main

import (
	"fmt"
	"sort"

	termutil "github.com/japanoise/termbox-util"
)

type kvList struct {
	items    []kv
	array    bool
	keyWidth int
}

type kv struct {
	key   string
	value string
	child interface{}
}

func printJSONNode(node interface{}) kvList {
	var items []kv
	keyWidth := 0
	var ret kvList
	switch v := node.(type) {
	case []interface{}:
		ret.array = true
		for i, child := range v {
			items = append(items, kv{key: fmt.Sprint(i), value: printJSONNodeShort(child), child: child})
			w := termutil.RunewidthStr(items[i].key)
			if w > keyWidth {
				keyWidth = w
			}
		}
	case map[string]interface{}:
		for key, child := range v {
			items = append(items, kv{key: key, value: printJSONNodeShort(child), child: child})
			w := termutil.RunewidthStr(key)
			if w > keyWidth {
				keyWidth = w
			}
		}
		sort.Slice(items, func(i, j int) bool {
			return items[i].key < items[j].key
		})
	default:
		items = []kv{kv{key: "", value: printJSONNodeShort(node), child: v}}
	}
	ret.items = items
	ret.keyWidth = keyWidth
	return ret
}

func printJSONNodeShort(node interface{}) string {
	switch v := node.(type) {
	case []interface{}:
		return fmt.Sprintf("[ (%d) ]", len(v))
	case map[string]interface{}:
		return "{ ... }"
	case string:
		return fmt.Sprintf("\"%s\"", v)
	default:
		return fmt.Sprint(v)
	}
}
