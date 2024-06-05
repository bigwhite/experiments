package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func drawDots(input [][]float64, featureIndex int, xlabel, ylabel string) (err error) {
	pts := plotter.XYs{}
	for _, sample := range input {
		pts = append(pts, plotter.XY{X: sample[featureIndex], Y: sample[len(sample)-1]})
	}

	// 创建一个新的plot
	p := plot.New()
	if err != nil {
		return
	}

	// 设置plot的标题和轴标签
	p.Title.Text = "Housing Dataset"
	p.X.Label.Text = xlabel
	p.Y.Label.Text = ylabel

	// 创建一个新的散点图,并将数据点添加到散点图中
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		return
	}
	scatter.GlyphStyle.Color = plotutil.Color(1)
	scatter.GlyphStyle.Radius = vg.Points(3)

	// 将散点图添加到plot中
	p.Add(scatter)

	// 保存图像到文件
	if err = p.Save(4*vg.Inch, 4*vg.Inch, "housing_"+xlabel+"_"+ylabel+".png"); err != nil {
		return
	}
	return
}
