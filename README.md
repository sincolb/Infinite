无限分类简单GO

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
