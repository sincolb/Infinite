package main

import (
	"encoding/json"
	"fmt"
)

type Menu struct {
	Id    int
	Pid   int
	Name  string
	Child []*Menu
}

func loop(list []*Menu) []*Menu {
	tree := make(map[int]*Menu, len(list))
	for _, ele := range list {
		tree[ele.Id] = ele
	}
	var Menus []*Menu
	for _, item := range tree {
		if temp, ok := tree[item.Pid]; ok {
			child := temp.Child
			child = append(child, tree[item.Id])
			temp.Child = child
			tree[item.Pid] = temp
		} else {
			Menus = append(Menus, item)
		}
	}
	return Menus
}

func recurrence(list []*Menu) []*Menu {
	var data map[int]map[int]*Menu = make(map[int]map[int]*Menu)
	for _, v := range list {
		if _, ok := data[v.Pid]; !ok {
			data[v.Pid] = make(map[int]*Menu)
		}
		data[v.Pid][v.Id] = v
	}
	result := buildTree(0, data)
	return result
}

func buildTree(index int, data map[int]map[int]*Menu) []*Menu {
	tmp := make([]*Menu, 0, len(data[index]))
	for id, item := range data[index] {
		if data[id] != nil {
			item.Child = buildTree(id, data)
		}
		tmp = append(tmp, item)
	}
	return tmp
}

var menus = []byte(`
[
    {
        "id": 5,
        "name": "test5",
        "pid": 0
    },
    {
        "id": 6,
        "name": "test6",
        "pid": 3
    },
    {
        "id": 7,
        "name": "test7",
        "pid": 4
    },
    {
        "id": 1,
        "name": "test0",
        "pid": 0
    },
    {
        "id": 2,
        "name": "test1",
        "pid": 1
    },
    {
        "id": 3,
        "name": "test2",
        "pid": 2
    },
    {
        "id": 4,
        "name": "test3",
        "pid": 1
    }
]`)

func main() {
	var data []*Menu
	if err := json.Unmarshal(menus, &data); err != nil {
		panic(err)
	}

	tree1, err := json.MarshalIndent(loop(data), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tree1))

	tree2, err := json.MarshalIndent(recurrence(data), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(tree2))
}
