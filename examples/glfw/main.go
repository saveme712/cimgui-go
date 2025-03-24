package main

import (
	"fmt"
	"runtime"

	"github.com/saveme712/cimgui-go/backend"
	"github.com/saveme712/cimgui-go/backend/glfwbackend"
	"github.com/saveme712/cimgui-go/examples/common"
	"github.com/saveme712/cimgui-go/imgui"
	_ "github.com/saveme712/cimgui-go/immarkdown"
	_ "github.com/saveme712/cimgui-go/imnodes"
)

var currentBackend backend.Backend[glfwbackend.GLFWWindowFlags]

func init() {
	runtime.LockOSThread()
}

func main() {
	common.Initialize()
	currentBackend, _ = backend.CreateBackend(glfwbackend.NewGLFWBackend())
	currentBackend.SetAfterCreateContextHook(common.AfterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(common.BeforeDestroyContext)

	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	currentBackend.CreateWindow("Hello from cimgui-go", 1200, 900)

	currentBackend.SetDropCallback(func(p []string) {
		fmt.Printf("drop triggered: %v", p)
	})

	currentBackend.SetCloseCallback(func() {
		fmt.Println("window is closing")
	})

	currentBackend.SetIcons(common.Image())

	currentBackend.Run(common.Loop)
}
