package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"time"

	"math"
	"image"
	"image/png"
	"draw2d"
)

const (
	w, h = 256, 256
)

var (
	lastTime int64
	folder   = "../../test_results/"
)

func initGc(w, h int) (image.Image, *draw2d.GraphicContext) {
	i := image.NewRGBA(w, h)
	gc := draw2d.NewGraphicContext(i)
	lastTime = time.Nanoseconds()

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	// fill the background 
	//gc.Clear()

	return i, gc
}

func saveToPngFile(TestName string, m image.Image) {
	dt := time.Nanoseconds() - lastTime
	fmt.Printf("%s during: %f ms\n", TestName, float(dt)*10e-6)
	filePath := folder + TestName + ".png"
	f, err := os.Open(filePath, os.O_CREAT|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}

/*
  <img src="../test_results/TestPath.png"/>
*/
func TestPath() {
	i, gc := initGc(w, h)
	gc.MoveTo(10.0, 10.0)
	gc.LineTo(100.0, 10.0)
	gc.LineTo(100.0, 100.0)
	gc.LineTo(10.0, 100.0)
	gc.LineTo(10.0, 10.0)
	gc.FillStroke()
	saveToPngFile("TestPath", i)
}


func cos(f float) float {
	return float(math.Cos(float64(f)))
}
func sin(f float) float {
	return float(math.Sin(float64(f)))
}
/*
  <img src="../test_results/TestDrawArc.png"/>
*/
func TestDrawArc() {
	i, gc := initGc(w, h)
	// draw an arc
	xc, yc := 128.0, 128.0
	radiusX, radiusY := 100.0, 100.0
	startAngle := 45.0 * (math.Pi / 180.0) /* angles are specified */
	angle := 135 * (math.Pi / 180.0)       /* in radians           */
	gc.SetLineWidth(10)
	gc.SetLineCap(draw2d.ButtCap)
	gc.SetStrokeColor(image.Black)
	gc.ArcTo(xc, yc, radiusX, radiusY, startAngle, angle)
	gc.Stroke()
	// fill a circle
	gc.SetStrokeColor(image.RGBAColor{255, 0x33, 0x33, 0x80})
	gc.SetFillColor(image.RGBAColor{255, 0x33, 0x33, 0x80})
	gc.SetLineWidth(6)

	gc.MoveTo(xc, yc)
	gc.LineTo(xc+cos(startAngle)*radiusX, yc+sin(startAngle)*radiusY)
	gc.MoveTo(xc, yc)
	gc.LineTo(xc-radiusX, yc)
	gc.Stroke()

	gc.ArcTo(xc, yc, 10.0, 10.0, 0, 2*math.Pi)
	gc.Fill()
	saveToPngFile("TestDrawArc", i)
}
/*
  <img src="../test_results/TestDrawArc.png"/>
*/
func TestDrawArcNegative() {
	i, gc := initGc(w, h)
	// draw an arc
	xc, yc := 128.0, 128.0
	radiusX, radiusY := 100.0, 100.0
	startAngle := 45.0 * (math.Pi / 180.0) /* angles are specified */
	angle := -225 * (math.Pi / 180.0)      /* in radians           */
	gc.SetLineWidth(10)
	gc.SetLineCap(draw2d.ButtCap)
	gc.SetStrokeColor(image.Black)
	gc.ArcTo(xc, yc, radiusX, radiusY, startAngle, angle)
	gc.Stroke()
	// fill a circle
	gc.SetStrokeColor(image.RGBAColor{255, 0x33, 0x33, 0x80})
	gc.SetFillColor(image.RGBAColor{255, 0x33, 0x33, 0x80})
	gc.SetLineWidth(6)

	gc.MoveTo(xc, yc)
	gc.LineTo(xc+cos(startAngle)*radiusX, yc+sin(startAngle)*radiusY)
	gc.MoveTo(xc, yc)
	gc.LineTo(xc-radiusX, yc)
	gc.Stroke()

	gc.ArcTo(xc, yc, 10.0, 10.0, 0, 2*math.Pi)
	gc.Fill()
	saveToPngFile("TestDrawArcNegative", i)
}

func TestCurveRectangle() {
	i, gc := initGc(w, h)

	/* a custom shape that could be wrapped in a function */
	x0, y0 := 25.6, 25.6 /* parameters like cairo_rectangle */
	rect_width, rect_height := 204.8, 204.8
	radius := 102.4 /* and an approximate curvature radius */

	x1 := x0 + rect_width
	y1 := y0 + rect_height
	if rect_width/2 < radius {
		if rect_height/2 < radius {
			gc.MoveTo(x0, (y0+y1)/2)
			gc.CubicCurveTo(x0, y0, x0, y0, (x0+x1)/2, y0)
			gc.CubicCurveTo(x1, y0, x1, y0, x1, (y0+y1)/2)
			gc.CubicCurveTo(x1, y1, x1, y1, (x1+x0)/2, y1)
			gc.CubicCurveTo(x0, y1, x0, y1, x0, (y0+y1)/2)
		} else {
			gc.MoveTo(x0, y0+radius)
			gc.CubicCurveTo(x0, y0, x0, y0, (x0+x1)/2, y0)
			gc.CubicCurveTo(x1, y0, x1, y0, x1, y0+radius)
			gc.LineTo(x1, y1-radius)
			gc.CubicCurveTo(x1, y1, x1, y1, (x1+x0)/2, y1)
			gc.CubicCurveTo(x0, y1, x0, y1, x0, y1-radius)
		}
	} else {
		if rect_height/2 < radius {
			gc.MoveTo(x0, (y0+y1)/2)
			gc.CubicCurveTo(x0, y0, x0, y0, x0+radius, y0)
			gc.LineTo(x1-radius, y0)
			gc.CubicCurveTo(x1, y0, x1, y0, x1, (y0+y1)/2)
			gc.CubicCurveTo(x1, y1, x1, y1, x1-radius, y1)
			gc.LineTo(x0+radius, y1)
			gc.CubicCurveTo(x0, y1, x0, y1, x0, (y0+y1)/2)
		} else {
			gc.MoveTo(x0, y0+radius)
			gc.CubicCurveTo(x0, y0, x0, y0, x0+radius, y0)
			gc.LineTo(x1-radius, y0)
			gc.CubicCurveTo(x1, y0, x1, y0, x1, y0+radius)
			gc.LineTo(x1, y1-radius)
			gc.CubicCurveTo(x1, y1, x1, y1, x1-radius, y1)
			gc.LineTo(x0+radius, y1)
			gc.CubicCurveTo(x0, y1, x0, y1, x0, y1-radius)
		}
	}
	gc.Close()

	gc.SetFillColor(image.RGBAColor{0x80, 0x80, 0xFF, 0xFF})
	gc.SetStrokeColor(image.RGBAColor{0x80, 0, 0, 0x80})
	gc.SetLineWidth(10.0)
	gc.FillStroke()

	saveToPngFile("TestCurveRectangle", i)
}
/*
  <img src="../test_results/TestDrawCubicCurve.png"/>
*/
func TestDrawCubicCurve() {
	i, gc := initGc(w, h)
	// draw a cubic curve
	x, y := 25.6, 128.0
	x1, y1 := 102.4, 230.4
	x2, y2 := 153.6, 25.6
	x3, y3 := 230.4, 128.0

	gc.SetFillColor(image.RGBAColor{0xAA, 0xAA, 0xAA, 0xFF})
	gc.SetLineWidth(10)
	gc.MoveTo(x, y)
	gc.CubicCurveTo(x1, y1, x2, y2, x3, y3)
	gc.Stroke()

	gc.SetStrokeColor(image.RGBAColor{0xFF, 0x33, 0x33, 0x88})

	gc.SetLineWidth(6)
	// draw segment of curve
	gc.MoveTo(x, y)
	gc.LineTo(x1, y1)
	gc.LineTo(x2, y2)
	gc.LineTo(x3, y3)
	gc.Stroke()
	saveToPngFile("TestDrawCubicCurve", i)
}

/*
  <img src="../test_results/TestDash.png"/>
*/
func TestDash() {
	i, gc := initGc(w, h)
	gc.SetLineDash([]float{50, 10, 10, 10}, -50.0)
	gc.SetLineCap(draw2d.ButtCap)
	gc.SetLineJoin(draw2d.BevelJoin)
	gc.SetLineWidth(10)

	gc.MoveTo(128.0, 25.6)
	gc.LineTo(128.0, 25.6)
	gc.LineTo(230.4, 230.4)
	gc.RLineTo(-102.4, 0.0)
	gc.CubicCurveTo(51.2, 230.4, 51.2, 128.0, 128.0, 128.0)
	gc.Stroke()
	gc.SetLineDash(nil, 0.0)
	saveToPngFile("TestDash", i)
}


/*
  <img src="../test_results/TestFillStroke.png"/>
*/
func TestFillStroke() {
	i, gc := initGc(w, h)
	gc.MoveTo(128.0, 25.6)
	gc.LineTo(230.4, 230.4)
	gc.RLineTo(-102.4, 0.0)
	gc.CubicCurveTo(51.2, 230.4, 51.2, 128.0, 128.0, 128.0)
	gc.Close()

	gc.MoveTo(64.0, 25.6)
	gc.RLineTo(51.2, 51.2)
	gc.RLineTo(-51.2, 51.2)
	gc.RLineTo(-51.2, -51.2)
	gc.Close()

	gc.SetLineWidth(10.0)
	gc.SetFillColor(image.RGBAColor{0, 0, 0xFF, 0xFF})
	gc.SetStrokeColor(image.Black)
	gc.FillStroke()
	saveToPngFile("TestFillStroke", i)
}

/*
  <img src="../test_results/TestFillStyle.png"/>
*/
func TestFillStyle() {
	i, gc := initGc(w, h)
	gc.SetLineWidth(6)

	gc.Rect(12, 12, 244, 70)

	wheel1 := new(draw2d.Path)
	wheel1.ArcTo(64, 64, 40, 40, 0, 2*math.Pi)
	wheel2 := new(draw2d.Path)
	wheel2.ArcTo(192, 64, 40, 40, 0, 2*math.Pi)

	gc.SetFillRule(draw2d.FillRuleEvenOdd)
	gc.SetFillColor(image.RGBAColor{0, 0xB2, 0, 0xFF})

	gc.SetStrokeColor(image.Black)
	gc.FillStroke(wheel1, wheel2)

	gc.Rect(12, 140, 244, 198)
	wheel1 = new(draw2d.Path)
	wheel1.ArcTo(64, 192, 40, 40, 0, 2*math.Pi)
	wheel2 = new(draw2d.Path)
	wheel2.ArcTo(192, 192, 40, 40, 0, -2*math.Pi)

	gc.SetFillRule(draw2d.FillRuleWinding)
	gc.SetFillColor(image.RGBAColor{0, 0, 0xE5, 0xFF})
	gc.FillStroke(wheel1, wheel2)
	saveToPngFile("TestFillStyle", i)
}

func TestMultiSegmentCaps() {
	i, gc := initGc(w, h)
	gc.MoveTo(50.0, 75.0)
	gc.LineTo(200.0, 75.0)

	gc.MoveTo(50.0, 125.0)
	gc.LineTo(200.0, 125.0)

	gc.MoveTo(50.0, 175.0)
	gc.LineTo(200.0, 175.0)

	gc.SetLineWidth(30.0)
	gc.SetLineCap(draw2d.RoundCap)
	gc.Stroke()
	saveToPngFile("TestMultiSegmentCaps", i)
}


func TestRoundRectangle() {
	i, gc := initGc(w, h)
	/* a custom shape that could be wrapped in a function */
	x, y := 25.6, 25.6
	width, height := 204.8, 204.8
	aspect := 1.0                  /* aspect ratio */
	corner_radius := height / 10.0 /* and corner curvature radius */

	radius := corner_radius / aspect
	degrees := math.Pi / 180.0

	gc.ArcTo(x+width-radius, y+radius, radius, radius, -90*degrees, 90*degrees)
	gc.ArcTo(x+width-radius, y+height-radius, radius, radius, 0*degrees, 90*degrees)
	gc.ArcTo(x+radius, y+height-radius, radius, radius, 90*degrees, 90*degrees)
	gc.ArcTo(x+radius, y+radius, radius, radius, 180*degrees, 90*degrees)
	gc.Close()

	gc.SetFillColor(image.RGBAColor{0x80, 0x80, 0xFF, 0xFF})
	gc.SetStrokeColor(image.RGBAColor{0x80, 0, 0, 0x80})
	gc.SetLineWidth(10.0)
	gc.FillStroke()

	saveToPngFile("TestRoundRectangle", i)
}

func TestLineCap() {
	i, gc := initGc(w, h)
	gc.SetLineWidth(30.0)
	gc.SetLineCap(draw2d.ButtCap)
	gc.MoveTo(64.0, 50.0)
	gc.LineTo(64.0, 200.0)
	gc.Stroke()
	gc.SetLineCap(draw2d.RoundCap)
	gc.MoveTo(128.0, 50.0)
	gc.LineTo(128.0, 200.0)
	gc.Stroke()
	gc.SetLineCap(draw2d.SquareCap)
	gc.MoveTo(192.0, 50.0)
	gc.LineTo(192.0, 200.0)
	gc.Stroke()

	/* draw helping lines */
	gc.SetStrokeColor(image.RGBAColor{0xFF, 0x33, 0x33, 0xFF})
	gc.SetLineWidth(2.56)
	gc.MoveTo(64.0, 50.0)
	gc.LineTo(64.0, 200.0)
	gc.MoveTo(128.0, 50.0)
	gc.LineTo(128.0, 200.0)
	gc.MoveTo(192.0, 50.0)
	gc.LineTo(192.0, 200.0)
	gc.Stroke()
	saveToPngFile("TestLineCap", i)
}
func TestLineJoin() {
	i, gc := initGc(w, h)
	gc.SetLineWidth(40.96)
	gc.MoveTo(76.8, 84.48)
	gc.RLineTo(51.2, -51.2)
	gc.RLineTo(51.2, 51.2)
	gc.SetLineJoin(draw2d.MiterJoin) /* default */
	gc.Stroke()

	gc.MoveTo(76.8, 161.28)
	gc.RLineTo(51.2, -51.2)
	gc.RLineTo(51.2, 51.2)
	gc.SetLineJoin(draw2d.BevelJoin)
	gc.Stroke()

	gc.MoveTo(76.8, 238.08)
	gc.RLineTo(51.2, -51.2)
	gc.RLineTo(51.2, 51.2)
	gc.SetLineJoin(draw2d.RoundJoin)
	gc.Stroke()
	saveToPngFile("TestLineJoin", i)
}


func main() {
	TestPath()
	TestDrawArc()
	TestDrawArcNegative()
	TestDrawCubicCurve()
	TestCurveRectangle()
	TestDash()
	TestFillStroke()
	TestFillStyle()
	TestMultiSegmentCaps()
	TestRoundRectangle()
	TestLineCap()
	TestLineJoin()
}