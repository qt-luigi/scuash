// スカッシュゲーム（壁打ちテニス）
package main

// ライブラリーのインポート
import (
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	clr "github.com/qt-luigi/ebiten-samples/color"
	"github.com/qt-luigi/ebiten-samples/draw"
)

func init() {
	// ウィンドウの作成
	ebiten.SetWindowSize(640, 480)
}

type Game struct {
	is_gameover   bool
	ball_ichi_x   int
	ball_ichi_y   int
	ball_idou_x   int
	ball_idou_y   int
	ball_size     int
	racket_ichi_x int
	racket_size   int
	point         int
	speed         int

	racket_idou int
	baseTime    time.Time
}

// ゲームの初期化
func (g *Game) init() {
	g.is_gameover = false
	g.ball_ichi_x = 0
	g.ball_ichi_y = 250
	g.ball_idou_x = 15
	g.ball_idou_y = -15
	g.ball_size = 10
	g.racket_ichi_x = 0
	g.racket_size = 100
	g.point = 0
	g.speed = 50
	ebiten.SetWindowTitle("スカッシュゲーム：スタート！")

	g.racket_idou = 15
	g.baseTime = time.Now()

	rand.Seed(time.Now().UnixNano())
}

func (g *Game) after() bool {
	now := time.Now()
	if now.Sub(g.baseTime) > (time.Duration(g.speed) * time.Millisecond) {
		g.baseTime = now
		return false
	}
	return true
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func (g *Game) Update() error {
	if g.is_gameover {
		g.click()
		return nil
	}
	if !g.after() {
		g.motion()
		g.moveBall()
	}
	return nil
}

// 画面の描画
func (g *Game) drawScreen(screen *ebiten.Image) {
	// 画面クリア＆キャンバス（画面）の作成
	screen.Fill(color.White)
}

func (g *Game) drawBall(screen *ebiten.Image) {
	// ボールを描く
	draw.ArcFill(screen, draw.SubImage, g.ball_ichi_x, g.ball_ichi_y, g.ball_size, clr.Red)
}

func (g *Game) drawRacket(screen *ebiten.Image) {
	// ラケットを描く
	draw.SquareFill(screen, g.racket_ichi_x, 470, g.racket_ichi_x+g.racket_size, 480, clr.Yellow)
}

// ボールの移動
func (g *Game) moveBall() {
	// （※ ゲームオーバー判定はUpdate()に切り出し）

	// 左右の壁に当たったかの判定（※ ball_sizeを追加）
	if g.ball_ichi_x+g.ball_idou_x-g.ball_size < 0 || g.ball_ichi_x+g.ball_idou_x+g.ball_size > 640 {
		g.ball_idou_x *= -1
		// winsound.Beep(1320, 50)
	}
	// 天井に当たったかの判定（※ ball_sizeを追加）
	if g.ball_ichi_y+g.ball_idou_y-g.ball_size < 0 {
		g.ball_idou_y *= -1
		// winsound.Beep(1320, 50)
	}
	// ラケットに当たったかの判定（※ ball_sizeを追加）
	if g.ball_ichi_y+g.ball_idou_y+g.ball_size > 470 &&
		(g.racket_ichi_x <= (g.ball_ichi_x+g.ball_idou_x) &&
			(g.ball_ichi_x+g.ball_idou_x) <= (g.racket_ichi_x+g.racket_size)) {
		g.ball_idou_y *= -1
		if rand.Intn(2) == 0 {
			g.ball_idou_x *= -1
		}
		// winsound.Beep(2000, 50)
		var message string
		mes := rand.Intn(5)
		if mes == 0 {
			message = "うまい！"
		}
		if mes == 1 {
			message = "グッド！"
		}
		if mes == 2 {
			message = "ナイス！"
		}
		if mes == 3 {
			message = "よしッ！"
		}
		if mes == 4 {
			message = "すてき！"
		}
		g.point += 10
		ebiten.SetWindowTitle(message + "　得点＝" + strconv.Itoa(g.point))
	}
	// ミスしたときの判定
	if g.ball_ichi_y+g.ball_idou_y >= 480 {
		var message string
		mes := rand.Intn(3)
		if mes == 0 {
			message = "ヘタくそ！"
		}
		if mes == 1 {
			message = "ミスしたね！"
		}
		if mes == 2 {
			message = "あーあ、見てられないね！"
		}
		ebiten.SetWindowTitle(message + "　得点＝" + strconv.Itoa(g.point))
		// winsound.Beep(200, 800)
		g.is_gameover = true
		return // （※ 追加）
	}
	if 0 <= g.ball_ichi_x+g.ball_idou_x && g.ball_ichi_x+g.ball_idou_x <= 640 {
		g.ball_ichi_x = g.ball_ichi_x + g.ball_idou_x
	}
	if 0 <= g.ball_ichi_y+g.ball_idou_y && g.ball_ichi_y+g.ball_idou_y <= 480 {
		g.ball_ichi_y = g.ball_ichi_y + g.ball_idou_y
	}
}

// マウスの動きの処理
// マウスポインタの位置確認
func (g *Game) motion() {
	x, _ := ebiten.CursorPosition()
	g.racket_ichi_x = x
}

// クリックで再スタート
func (g *Game) click() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.init()
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ゲームの繰り返し処理の命令
	g.drawScreen(screen)
	g.drawBall(screen)
	g.drawRacket(screen)
}

func main() {
	// ゲームのメイン処理
	game := &Game{}
	game.init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
