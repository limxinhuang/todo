package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/limxinhuang/todo/todo"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("错误：请输入任务名称")
			return
		}
		title := os.Args[2]
		todo.Add(title)
	case "list":
		todo.List()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("错误：请输入任务 ID")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		todo.Completed(id)
	case "del":
		if len(os.Args) < 3 {
			fmt.Println("错误：请输入任务 ID")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		todo.DeleteTodo(id)
	case "edit":
		if len(os.Args) < 4 {
			fmt.Println("错误：请输入任务 ID 和新任务标题")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		title := os.Args[3]
		todo.UpdateTitle(id, title)
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("todo 使用说明:")
	fmt.Println("  add <任务名>   - 添加任务")
	fmt.Println("  list          - 列出所有任务")
	fmt.Println("  done <ID>     - 完成任务")
	fmt.Println("  del <ID>      - 删除任务")
	fmt.Println("  edit <ID> <新标题> - 更新任务标题")
}
