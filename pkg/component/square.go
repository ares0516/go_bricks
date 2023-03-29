package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
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
	score int
	Image *ebiten.Image
	Opts  *ebiten.DrawImageOptions
}

func NewSquare(bgc color.RGBA, h, w int, x, y, step float64) *Square {
	image := ebiten.NewImage(w, h)
	image.Fill(bgc)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	angle := float64(240+rand.Intn(60)) * math.Pi / 180
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

// game.ball = component.NewSquare(define.Red, 5, 5, 300, 390, 3)
// game.board = component.NewSquare(define.White, 5, 40, 285, 400, 0)
func (s *Square) CollisionDetection3(w, h float64, board *Square) {
	x, y := s.x+s.stepX, s.y+s.stepY // 小球下一步位置坐标
	log.Printf("current squre x[%f]y[%f]", s.x, s.y)
	log.Printf("suqre stepX[%f]stepY[%f]", s.stepX, s.stepY)
	log.Printf("squre next x[%f]y[%f]", x, y)
	log.Printf("w[%f]h[%f]", w, h)
	if x <= 0 {
		s.stepX *= -1
		log.Printf("window left")
	} else if x+s.w >= w {
		s.stepX *= -1
		log.Printf("window right")
	} else if y <= 0 {
		s.stepY *= -1
		log.Printf("window top")
	} else if (x+s.w >= board.x) && (x <= board.x+board.w) && y+s.h >= board.y { // 小球与板子上表面发生碰撞
		// 小球右侧坐标大于等于板子左侧坐标，小球左侧坐标小于等于板子右侧坐标，小球下侧坐标小于等于板子上侧坐标
		s.score++
		s.stepY *= -1
		log.Printf("board top")
	} else if y+s.h >= h { // 小球落地
		s.alive = false
		log.Printf("window bottom")
		return
	}
	tx, ty := s.stepX, s.stepY
	s.x += tx
	s.y += ty
	s.Opts.GeoM.Translate(tx, ty)
}

func (s *Square) CollisionDetection2(w, h float64, board *Square) {
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

// Move board 的移动逻辑，在未开球的时候，小球也跟着board移动
func (s *Square) Move(w int, step float64, ball *Square) {
	W := float64(w)
	if s.x+step < 0 {
		s.Opts.GeoM.Translate(0-s.x, 0)
		if !ball.IsRun {
			ball.Opts.GeoM.Translate(0-s.x, 0)
			ball.x += -s.x
		}
		s.x = 0
		return
	}
	if s.x+step+s.w > W {
		s.Opts.GeoM.Translate(W-s.x-s.w, 0)
		if !ball.IsRun {
			ball.Opts.GeoM.Translate(W-s.x-s.w, 0)
			ball.x += W - s.x - s.w
		}
		s.x = W - s.w
		return
	}
	s.x += step
	s.Opts.GeoM.Translate(step, 0)
	if !ball.IsRun {
		ball.Opts.GeoM.Translate(step, 0)
		ball.x += step
	}
}

// HitDetection 碰撞检测，是否撞击到奖励。撞击到奖励后，奖励消失，分数增加
func (s *Square) HitDetection(awards *[]*Square) {
	for i := 0; i < len(*awards); i++ {
		award := (*awards)[i]
		if math.Abs(s.x-award.x) <= s.w && math.Abs(s.y-award.y) <= s.h {
			*awards = append((*awards)[0:i], (*awards)[i+1:]...)
			i--
			s.score += 10
		}
	}
}

func (s *Square) IsAlive() bool {
	return s.alive
}

func (s *Square) GetScore() string {
	return strconv.Itoa(s.score)
}
