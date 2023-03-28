package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"math"
	"math/rand"
)

type Square struct {
	bgc   color.RGBA
	h     float64
	w     float64
	x     float64
	y     float64
	step  float64
	angle float64
	stepX float64
	stepY float64
	alive bool
	IsRun bool
	Image *ebiten.Image
	Opts  *ebiten.DrawImageOptions
}

func NewSquare(bgc color.RGBA, h, w int, x, y, step float64) *Square {
	image := ebiten.NewImage(w, h)
	image.Fill(bgc)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	angle := float64(rand.Intn(120) + 30)
	//angle := float64(0)
	return &Square{
		bgc:   bgc,
		h:     float64(h),
		w:     float64(w),
		x:     x,
		y:     y,
		step:  step,
		angle: angle,
		stepX: step * math.Cos(angle),
		stepY: step * math.Sin(angle),
		alive: true,
		Image: image,
		Opts:  opts,
	}
}

func (s *Square) CollisionDetection2(w, h float64) {
	x, y := s.x+s.stepX, s.y+s.stepY // 小球下一步位置坐标
	log.Printf("x[%f]y[%f]", x, y)
	if x <= 0 {
		s.stepX *= -1
	} else if x+s.w >= w {
		s.stepX *= -1
	} else if y <= 0 {
		s.stepY *= -1
	} else if y+s.h >= h {
		s.stepY *= -1
	}
	tx, ty := s.stepX, s.stepY
	s.x += tx
	s.y += ty
	s.Opts.GeoM.Translate(tx, ty)
}

func (s *Square) CollisionDetection(w, h float64) {
	x, y := s.x+s.stepX, s.y+s.stepY // 小球下一步位置坐标
	tx, ty := s.stepX, s.stepY
	if x <= 0 {
		tx = -s.x
		ty = tx * math.Tan(s.angle) // y轴镜像反射
		s.stepX *= -1               // x轴转向
	} else if x >= w {
		//tx = -(w - s.x)
		tx = -s.x
		ty = tx * math.Tan(s.angle)
		s.stepX *= -1
	} else if y <= 0 {
		ty = -s.y
		tx = ty / math.Tan(s.angle)
		s.stepY *= -1
	} else if y >= h {
		//ty = -(h - s.y)
		ty = -s.y
		tx = ty / math.Tan(s.angle)
		s.stepY *= -1
	}
	log.Printf("x : %f , tx %f, y: %f, ty: %f", x, tx, y, ty)
	s.x += tx
	s.y += ty
	s.Opts.GeoM.Translate(tx, ty)
}
