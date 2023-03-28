package main

import (
	"github.com/ares0516/snake/pkg/component"
	"github.com/ares0516/snake/pkg/define"
	"github.com/hajimehoshi/ebiten/v2"
)
import "image/color"

type MyGame struct {
	screenWidth, screenHeight int
	ball                      *component.Square
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
		g.ball.CollisionDetection2(float64(g.screenWidth), float64(g.screenHeight))
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {

	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {

	}

	return nil
}

// 在屏幕上绘制游戏内容
func (g *MyGame) Draw(screen *ebiten.Image) {
	// 绘制黑色背景
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// 绘制图像
	screen.DrawImage(g.ball.Image, g.ball.Opts)

}

func main() {
	game := NewMyGame()
	//                                颜色       高度   宽度   中心点x    中心点y  步长
	//game.ball = component.NewSquare(define.Red, 240, 320, 157.5, 225, 2)
	game.ball = component.NewSquare(define.Red, 240, 320, 0, 0, 20)

	ebiten.SetWindowTitle("SNAKE")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
