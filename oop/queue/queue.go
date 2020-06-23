package queue

import "fmt"

//interface 代表任何类型
//需求修改为只可以push 任何类型
//pop int
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	for i, v := range *q {
		if value, ok := v.(int); ok {
			*q = append((*q)[0:i], (*q)[i+1:]...)
			return value
		} else {
			fmt.Println("ERROR")
		}
	}
	panic("has no int value")
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

//type Queue []int
//
//func (q *Queue) Push(v int)  {
//	*q = append(*q, v)
//}
//
//func (q *Queue) Pop() int {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head
//}
//
//func (q Queue) IsEmpty() bool {
//	return len(q) == 0
//}
