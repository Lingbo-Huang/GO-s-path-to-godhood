package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math"
)

func main() {
	// 创建一个绘图
	p := plot.New()

	// 设置图像的标题和标签
	p.Title.Text = "Sin Function"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// 创建一个正弦函数的绘图数据
	points := make(plotter.XYs, 100)
	for i := range points {
		x := float64(i) * 0.1 // X 轴范围
		y := math.Sin(x)      // 正弦函数
		points[i].X = x
		points[i].Y = y
	}

	// 创建线条并添加到绘图
	line, points1, err := plotter.NewLinePoints(points)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Add(line, points1)

	// 保存图像到文件
	if err := p.Save(4*vg.Inch, 4*vg.Inch, ".\\test\\test3\\sin_function.png"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Plot saved to sin_function.png")
}
