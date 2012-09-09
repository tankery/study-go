/*
 * http://tour.golang.org/#59
 * Exercise: Images
 * 
 * Remember the picture generator you wrote earlier?
 * Let's write another one, but this time it will
 * return an implementation of image.Image instead of
 * a slice of data.
 * 
 * Define your own Image type, implement the necessary
 * methods, and call pic.ShowImage.
 * 
 * Bounds should return a image.Rectangle, like
 * image.Rect(0, 0, w, h).
 * 
 * ColorModel should return color.RGBAModel.
 * 
 * At should return a color; the value v in the last
 * picture generator corresponds to
 * color.RGBA{v, v, 255, 255} in this one.
 */
package main

import (
	"image"
	"image/color"
	"tour/pic"
)

// 定义Image类型
// 类似的定义还可以这样：type MyInt int
// 相当于typedef
type Image struct{
	content [][]uint32	// 二维数组，存储图片内容
				// 包含每个像素点的RGBA值。
	width, height int	// 图片宽度和高度
}

// 自定义的像素点函数，返回给定点的RGBA值
func valueOfPointer(x, y int) uint32 {
	return uint32(0xfffff*x^(0xfffff*y + 0xff))
}

// 自定义的图片生成函数，用于使用给定的像素点函数生成一幅图片
func makePic(w, h int, f func(int,int) uint32) *Image {
	img := new(Image)
	img.width = w
	img.height = h
	
	// 此处先申请一个长度为w，类型为[]uint32的数组
	img.content = make([][]uint32, w)
	for x := 0; x < w; x++ {
		// 再为每个数组的元素申请h长度的uint32型数组
		// 由此而创建出一块 w x h 的二维数组
		img.content[x] = make([]uint32, h)
		
		// 使用f函数为每个像素赋值
		for y := 0; y < h; y++ {
			img.content[x][y] = f(x, y)
		}
	}
	
	return img
}

////////////////////////////////////////////
// 接下来的几个函数是接口image.Image的函数实现
// Go语言中，无需显示的申明实现接口
// 只需要实现接口的所有函数，即实现了接口

// Bounds 函数返回图片的可用区域
// 在 func 和函数名之间加上类型，表示此函数是该类型的成员函数
// 注意，此处的类型不仅限于结构体，比如浮点数、整数都可以。
func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// ColorModel 函数指明图片使用的颜色模式
// 这里，我们选用RGBA模式
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At 函数，返回指定像素点的颜色属性
func (img *Image) At(x, y int) color.Color {
	// 根据练习的说明设置超出范围的点的颜色
	if x >= img.width || y >= img.height {
		return color.RGBA{uint8(x), uint8(y), 0xff, 0xff}
	}
	
	// 根据存储的二维数组，生成RGBA模式的颜色并返回
	var c = img.content[x][y]
	return color.RGBA{
		uint8((c >> 24) & 0xff),
		uint8((c >> 16) & 0xff),
		uint8((c >> 8) & 0xff),
		uint8(c & 0xff) }
}

func main() {
	m := makePic(200, 100, valueOfPointer)
	// 调用pic类的ShowImage来显示生成的图片
	pic.ShowImage(m)
}

