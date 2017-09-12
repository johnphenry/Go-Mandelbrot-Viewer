package main
import(
"github.com/veandco/go-sdl2/sdl"
"./mandel"
"os"
)

var(
offset complex128 = 0
scale complex128 = 1
)

const(
MOVE_SPEED = 4.0 / 60.0
SCALE_SPEED = 0.1
)
func main(){
	sdl.Init(sdl.INIT_VIDEO)
	defer sdl.Quit()
	const(
	W=640
	H=480
	POS = sdl.WINDOWPOS_UNDEFINED)
	win, _ := sdl.CreateWindow("Mandelbrot Set", POS, POS, W, H, sdl.WINDOW_SHOWN)
	defer win.Destroy()
	ren, _:=sdl.CreateRenderer(win,-1,sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC )
	defer ren.Destroy()
	Frame:= func(){
		ren.Present()
		var event sdl.Event
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e:= event.(type) {
				case *sdl.QuitEvent:
				os.Exit(0)
				case *sdl.KeyDownEvent:
					switch( e.Keysym.Sym){
						case sdl.K_LEFT:
							offset-= MOVE_SPEED*scale
						case sdl.K_RIGHT:
							offset+= MOVE_SPEED*scale
						case sdl.K_UP:
							offset-= MOVE_SPEED*1i*scale
						case sdl.K_DOWN:
							offset+= MOVE_SPEED*1i*scale
						case sdl.K_w:
							scale*=1.0-SCALE_SPEED
						case sdl.K_s:
							scale*=1.0+SCALE_SPEED
					}
			}
		}
	}
	for{
		for y :=0;y<H;y++{
			for x:=0;x<W;x++{
				r,g,b:=mandel.Calc_pixel(x,y,W,H, offset, scale)
				ren.SetDrawColor(r,g,b,255)
				ren.DrawPoint(x,y);
			}
		}
		Frame()
	}
}
