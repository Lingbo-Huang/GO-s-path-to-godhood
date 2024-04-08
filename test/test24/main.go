package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func main() {
	// 创建一个新的动画
	anim := gif.GIF{}

	// 定义一些颜色
	palette := []color.Color{
		color.White,
		color.RGBA{0, 255, 0, 255}, // 绿色
		color.RGBA{255, 0, 0, 255}, // 红色
		color.RGBA{0, 0, 255, 255}, // 蓝色
	}

	// 添加第一帧
	addFrame(&anim, palette, 0)

	// 添加第二帧
	addFrame(&anim, palette, 1)

	// 保存 GIF 文件
	file, err := os.Create("./toys/test24/animation.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	gif.EncodeAll(file, &anim)
}

func addFrame(anim *gif.GIF, palette []color.Color, colorIndex int) {
	const (
		width  = 100
		height = 100
	)

	// 创建一个新的图像
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	// 在图像上绘制一些东西（这里简单地画一个矩形）
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			img.Set(i, j, palette[colorIndex])
		}
	}

	// 添加图像到动画中
	anim.Delay = append(anim.Delay, 100) // 每一帧的延迟（以100毫秒为单位）
	anim.Image = append(anim.Image, img)
}
