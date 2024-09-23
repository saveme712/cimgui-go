package ebitenbackend

import (
	"fmt"
	"image/color"
	"runtime"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw draws the generated imgui frame to the screen.
// This is usually called inside the game's Draw() function.
func (e *EbitenBackend) Draw(screen *ebiten.Image) {
	// add background color
	if e.bgColor.W == 1 {
		screen.Fill(color.RGBA{
			byte(e.bgColor.X * 255),
			byte(e.bgColor.Y * 255),
			byte(e.bgColor.Z * 255),
			byte(e.bgColor.W * 255),
		})
	}

	if e.debug {
		ebitenutil.DebugPrintAt(
			screen,
			fmt.Sprintf("TPS: %.2f\nFPS: %.2f\n[C]lipMask: %t\nSync [I]nput: %t\nSync C[u]rsor: %t\nControl Cursor [S]hape: %t",
				ebiten.ActualTPS(), ebiten.ActualFPS(), e.ClipMask(), e.syncInputs, e.syncCursor, e.controlCursorShape),
			10, 2,
		)
	}

	imgui.Render()
	if e.clipMask {
		if e.lmask == nil {
			w, h := screen.Size()
			e.lmask = ebiten.NewImage(w, h)
		} else {
			w1, h1 := screen.Size()
			w2, h2 := e.lmask.Size()
			if w1 != w2 || h1 != h2 {
				e.lmask.Dispose()
				e.lmask = ebiten.NewImage(w1, h1)
			}
		}
		RenderMasked(screen, e.lmask, imgui.CurrentDrawData(), e.cache, e.filter)
	} else {
		Render(screen, imgui.CurrentDrawData(), e.cache, e.filter)
	}
}

// Update needs to be called on every frame, before cimgui-go calls.
// This is usually called inside the game's Update() function.
// delta is the time in seconds since the last frame.
func (e *EbitenBackend) Update() error {
	e.BeginFrame()
	e.loop()
	e.EndFrame()

	return nil
}

func (e *EbitenBackend) Layout(outsideWidth, outsideHeight int) (int, int) {
	if e.retina {
		m := ebiten.DeviceScaleFactor()
		e.currentWidth = int(float64(outsideWidth) * m)
		e.currentHeight = int(float64(outsideHeight) * m)
	} else {
		e.currentWidth = outsideWidth
		e.currentHeight = outsideHeight
	}

	return e.currentWidth, e.currentHeight
}

func (e *EbitenBackend) onfinalize() {
	if e.beforeDestroy != nil {
		e.beforeDestroy()
	}

	runtime.SetFinalizer(e, nil)
	e.ctx.Destroy()
}

func (e *EbitenBackend) controlCursorShapeFn() {
	if !e.controlCursorShape {
		return
	}

	switch imgui.CurrentMouseCursor() {
	case imgui.MouseCursorNone:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	case imgui.MouseCursorArrow:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	case imgui.MouseCursorTextInput:
		ebiten.SetCursorShape(ebiten.CursorShapeText)
	case imgui.MouseCursorResizeAll:
		ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	case imgui.MouseCursorResizeEW:
		ebiten.SetCursorShape(ebiten.CursorShapeEWResize)
	case imgui.MouseCursorResizeNS:
		ebiten.SetCursorShape(ebiten.CursorShapeNSResize)
	case imgui.MouseCursorHand:
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
	default:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	}
}
