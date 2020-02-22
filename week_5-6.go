package week5_6

import "strconv"

//FizzWord รับพารามิเตอร์ 2 ตัว เมื่อนำตัวแรกมาหารด้วยตัวที่สองลงตัวจะคืนค่าเป็น "Fizz" หากไม่ลงตัวจะคืนค่าเป็นพารามิตเอร์ตัวแรกแต่มีไทป์เป็น string
func FizzWord(number int, mod int) string {
	return strconv.Itoa(number)
}

//MultiplicationTable จะรับพารามิเตอร์ 1 ตัวแล้วคืนค่าเป็นจำนวนสูตรคูณตั้งแต่ 1-12 กลับไป เช่น รับ n มาเป็น 2 จะคืนค่า []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24}
func MultiplicationTable(n int) []int {
	return []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24}
}

//StringMultiplicationTable เหมือนกับ MultiplicationTable ในรูปแบบ string เช่น 2 x 1 = 2 ...
func StringMultiplicationTable(n int) []string {
	return []string{"2 x 1 = 2", "2 x 2 = 4", "2 x 3 = 6", "2 x 4 = 8",
		"2 x 5 = 10", "2 x 6 = 12", "2 x 7 = 14", "2 x 8 = 16", "2 x 9 = 18",
		"2 x 10 = 20", "2 x 11 = 22", "2 x 12 = 24"}
}

//DeleteIntItem รับพารามิเตอร์ 2 ตัว ตัวแรกเป็น array ตัวที่ 2 เป็นตัวเลข หากในอาเรย์มีตัวเลขเหมือนพารามิเตอร์ตัวที่ 2 ให้ลบทิ้งและคืนค่าที่ได้กลับไป
func DeleteIntItem([]int, int) []int {
	return []int{}
}

//Grade รับค่าเป็นคะแนน แล้วคืนค่าเป็นเกรด A,B,C,D,F โดยใช้เกณฑ์มาตฐาน A มากกว่าหรือเท่ากับ 80
func Grade(point float32) string {
	return "F"
}
