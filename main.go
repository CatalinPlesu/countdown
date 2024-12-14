package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Game struct {
	startTime     time.Time
	duration      time.Duration
	windowWidth   int
	windowHeight  int
	fontFace      font.Face
	windowResized bool // To track if the window size has been set
}

func NewGame(minutes int, fontSize float64) *Game {
	// Load a system font (or embedded font) using opentype
	fontData, err := os.ReadFile("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf") // Example font
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}
	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatalf("Failed to parse font: %v", err)
	}
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("Failed to create font face: %v", err)
	}

	return &Game{
		startTime:    time.Now(),
		duration:     time.Duration(minutes) * time.Minute,
		windowWidth:  800, // Initial window width (can adjust)
		windowHeight: 200, // Initial window height (can adjust)
		fontFace:     face,
		windowResized: false,
	}
}

func (g *Game) Update() error {
	// Exit the game if the countdown is over
	if time.Since(g.startTime) >= g.duration {
		fmt.Println("Countdown complete!")
		return ebiten.Termination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen with transparency (or custom color for debugging)
	screen.Fill(color.RGBA{0, 0, 0, 0})

	// Calculate the remaining time
	remaining := g.duration - time.Since(g.startTime)
	if remaining < 0 {
		remaining = 0
	}
	hours := int(remaining.Hours())
	minutes := int(remaining.Minutes()) % 60
	seconds := int(remaining.Seconds()) % 60

	// Format the countdown timer
	countdownText := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

	// Measure text dimensions to position correctly
	bounds, _ := font.BoundString(g.fontFace, countdownText)
	textWidth := (bounds.Max.X - bounds.Min.X).Ceil()
	textHeight := (bounds.Max.Y - bounds.Min.Y).Ceil()

	// Resize the window only once
	if !g.windowResized {
		// Resize the window to fit the text
		g.windowWidth = textWidth + 40  // Add padding (20 on each side)
		g.windowHeight = textHeight + 40 // Add padding (20 on each side)

		// Set the new window size
		ebiten.SetWindowSize(g.windowWidth, g.windowHeight)

		// Set the window position to the bottom-right corner
		screenWidth, screenHeight := ebiten.Monitor().Size()
		ebiten.SetWindowPosition(screenWidth-g.windowWidth-10, screenHeight-g.windowHeight-10) // 10px margin from bottom-right corner

		// Mark the window as resized
		g.windowResized = true
	}

	// Position for the text (centered horizontally and vertically)
	x := (g.windowWidth - textWidth) / 2
	y := (g.windowHeight - textHeight) / 2 + 20

	// Draw the text on the screen
	text.Draw(screen, countdownText, g.fontFace, x, y, color.RGBA{255, 0, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the fixed window size after resizing
	return g.windowWidth, g.windowHeight
}

func main() {
	// Get the minutes argument from the command line
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <minutes>", os.Args[0])
	}

	// Convert the argument to an integer
	minutes, err := strconv.Atoi(os.Args[1])
	if err != nil || minutes <= 0 {
		log.Fatalf("Invalid minutes: %s", os.Args[1])
	}

	// Set a larger font size for better readability
	fontSize := 30.0

	// Create a new game instance
	game := NewGame(minutes, fontSize)

	// Set up the Ebiten window
	ebiten.SetWindowTitle("Countdown Overlay")
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowDecorated(false) // Removes the title bar and border
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetScreenTransparent(true) // Keep the screen transparent for debugging purposes
	ebiten.SetWindowFloating(true)    // Always on top

	// Start the game loop
	if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
