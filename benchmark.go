package main

import (
	"time"
  "os"
  "github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)


func benchMark(upto int,increaseAmount int){
  var start time.Time

  threadedRes := make([]opts.BarData,0)
  noThreadedRes := make([]opts.BarData,0)
  originalArr := make([]int,0)
  xAxis := make([]int,0)
  

  current := increaseAmount
    
   for current < upto {
      originalArr = GenerateRandomSlice(current)

      xAxis = append(xAxis,current)

      start = time.Now()
      MergeSortThreaded(originalArr)
      threadedRes = append(threadedRes,opts.BarData{Value:time.Since(start).Milliseconds()})

      start = time.Now()
      MergeSort(originalArr)
      noThreadedRes = append(noThreadedRes,opts.BarData{Value:time.Since(start).Milliseconds()})

      current += increaseAmount
   }

  bar := charts.NewBar()
  bar.SetXAxis(xAxis).AddSeries("Multi-Threaded",threadedRes).AddSeries("Single-Threaded", noThreadedRes)
  f, _ := os.Create("bar.html")
  bar.Render(f)

}

