package main

import (
	"github.com/ares0516/snake/pkg/component"
	"github.com/ares0516/snake/pkg/define"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand"
	"time"
)
import "image/color"

type MyGame struct {
	screenWidth  int
	screenHeight int
	ball         *component.Square
	board        *component.Square
	awards       []*component.Square
	cls          chan int
}

func NewMyGame() *MyGame {
	return &MyGame{
		screenWidth:  640,
		screenHeight: 480,
	}
}

func (g *MyGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *MyGame) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.ball.IsRun = true
	}

	if g.ball.IsRun {
		g.ball.CollisionDetection3(float64(g.screenWidth), float64(g.screenHeight), g.board)
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.board.Move(g.screenWidth, -5, g.ball)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.board.Move(g.screenWidth, 5, g.ball)
	}

	return nil
}

// Draw 在屏幕上绘制游戏内容
func (g *MyGame) Draw(screen *ebiten.Image) {
	// 绘制黑色背景
	screen.Fill(color.RGBA{0, 0, 0, 255})
	// 绘制图像
	screen.DrawImage(g.ball.Image, g.ball.Opts)
	screen.DrawImage(g.board.Image, g.board.Opts)
	// 绘制奖励
	for _, i := range g.awards {
		screen.DrawImage(i.Image, i.Opts)
	}
	if !g.ball.IsAlive() {
		ebitenutil.DebugPrint(screen, "Game Over")
	} else {
		ebitenutil.DebugPrint(screen, "Score: "+g.ball.GetScore())
	}

}

func (g *MyGame) AwardGenerator() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if !g.ball.IsAlive() {
				return
			}
			if len(g.awards) == 5 {
				g.awards = g.awards[1:5]
			}
			g.awards = append(g.awards, component.NewSquare(define.Yellow, 5, 5, float64(rand.Intn(300)+10), float64(rand.Intn(200)+10), 0))
		case <-g.cls:
			return
		}
	}
}

func main() {
	game := NewMyGame()
	//                                颜色       高度   宽度   中心点x    中心点y  步长
	//game.ball = component.NewSquare(define.Red, 240, 320, 157.5, 225, 2)
	game.ball = component.NewSquare(define.Red, 5, 5, 300, 395, 3)
	game.board = component.NewSquare(define.White, 5, 40, 285, 400, 0)
	game.cls = make(chan int, 0)
	game.awards = make([]*component.Square, 0)

	go game.AwardGenerator()

	ebiten.SetWindowTitle("SNAKE")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
