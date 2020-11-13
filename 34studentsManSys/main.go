package main

import (
	"fmt"
	"os"
)

var (
	allStudent map[int64]*student // 全局变量声明
)

type student struct {
	id   int64
	name string
}

func newStu(name string, id int64) *student {
	return &student{
		name: name,
		id: id,
	}
}

func all() {
	for k,v := range allStudent {
		fmt.Printf("学号：%d, 姓名：%s\n", k, v.name)
	}
}

func add() {
	var (
		name string
		id int64
	)

	print("输入名字：")
	fmt.Scanln(&name)
	print("输入学号：")
	fmt.Scanln(&id)
	stu := newStu(name, id)
	allStudent[id] = stu
}

func del() {
	var deleteID int64
	print("输入删除id：")
	fmt.Scanln(&deleteID)
	delete(allStudent,deleteID)
}

func main() {
	// 查看编辑新增删除学生

	// 全局变量实例化，开辟内存
	allStudent = make(map[int64]*student, 48) 
	for {
		fmt.Println(`
			1. 查看所有
			2. 新增
			3. 删除
			4. 退出
		`)

		choice := 1
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			all()
		case 2:
			add()
		case 3:
			del()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("input error")
		}
	}
}
