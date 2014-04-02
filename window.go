package colorshow

import (
	"bytes"
	"image"
	"image/png"
	"log"

	"github.com/banthar/Go-SDL/sdl"
)

func DisplayImage(src image.Image) {
	// Initialise SDL
	log.Println("Initialising SDL.")
	sdl.Init(sdl.INIT_VIDEO)
	defer sdl.Quit()

	// Transform the image into an SDL surface.
	log.Println("Transforming image into SDL surface.")
	var buf bytes.Buffer
	png.Encode(&buf, src)
	sdlRW := sdl.RWFromReader(&buf)
	sdlSurf := sdl.Load_RW(sdlRW, false)

	// Set the display up.
	log.Println("Setting up screen.")
	sdlScreen := sdl.SetVideoMode(int(sdlSurf.W), int(sdlSurf.H), 32, sdl.SWSURFACE)
	sdlScreen.Blit(nil, sdlSurf, nil)

	// Loop until exit.
	log.Println("Displaying image.")
	loop(sdlScreen)
}

func loop(screen *sdl.Surface) {
	for {
		screen.Flip()
		event := sdl.PollEvent()
		if event != nil {
			if _, ok := event.(*sdl.QuitEvent); ok {
				log.Println("Caught quit event. Quiting SDL.")
				break
			}
		}
	}
}
