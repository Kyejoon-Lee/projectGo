package service

import (
	"github.com/markcheno/go-quote"
	"time"
)

func FindDate(q quote.Quote) (time.Time,time.Time, float64){
	leftLow := 0
	leftHigh := 0
	left := 0
	diff := 0.00
	answer1 := q.Date[0]
	answer2 := q.Date[0]
	for left < len(q.High)-1{

		if(q.Low[left] < q.Low[leftLow]){
			leftLow = left
			leftHigh = left
			//fmt.Println(left)
		}
		if(q.High[left] >= q.High[leftHigh]){
			//fmt.Println("check", left)
			if(diff < q.High[left] - q.Low[leftLow]){
				diff = q.High[left] - q.Low[leftLow]
				answer1 = q.Date[leftLow]
				answer2 = q.Date[left]
				//fmt.Println("point",left, q.High[left])
			}
			leftHigh = left
		}
		left++
	}
	return answer1,answer2, diff
}

func FindDateClose(q quote.Quote) (time.Time,time.Time, float64) {
	left := 0
	diff := 0.00
	answer1 := q.Date[0]
	answer2 := q.Date[0]
	for i := 1;i < len(q.High);i++{
		if(q.Close[left]> q.Close[i]){
			left = i
		}else if(q.Close[i] - q.Close[left]> diff){
			diff = q.Close[i] - q.Close[left]
			answer1 = q.Date[left]
			answer2 = q.Date[i]
		}
	}
	return answer1,answer2,diff
}

