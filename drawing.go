package main
import rl "github.com/gen2brain/raylib-go/raylib"
//import "fmt"

var (
	zoom float32 = 15
	scroll = rl.Vector2{0, 0};
)

func renderLoop() {
	for !rl.WindowShouldClose() {
		textPosX := int32(rl.GetScreenWidth()/2) - rl.MeasureText(mainFile.fileName, 20)/2
		font := rl.LoadFont("resources/comic-sans.ttf");

		rl.BeginDrawing();

		rl.ClearBackground(rl.RayWhite);
		rl.DrawText(mainFile.fileName, textPosX, 30, 20, rl.Black);
		// rl.DrawFPS(WIDTH-30, 20);


		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyEqual) {
			zoom++;
		} else if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyMinus) {
			zoom--;
		} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyPageDown) {
			if !(scroll.Y - 10 < -(rl.MeasureTextEx(font, string(mainFile.data), zoom, 1).Y - float32(rl.GetScreenHeight())/2)) {
				scroll.Y -= 10;
			}
		} else if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyPageUp) {
			if !(scroll.Y + 10 > 0) {
				scroll.Y += 10;
			}
		} else if rl.GetMouseWheelMove() != 0 {
			if !(scroll.Y + rl.GetMouseWheelMove() * 50 > 0 && scroll.Y + rl.GetMouseWheelMove() * 50 < -(rl.MeasureTextEx(font, string(mainFile.data), zoom, 1).Y - float32(rl.GetScreenHeight())/2)) {
				scroll.Y += rl.GetMouseWheelMove() * 50;
			}
		}

		rl.DrawTextEx(font, string(mainFile.data), scroll, zoom, 1, rl.Black)

		rl.EndDrawing();
		rl.UnloadFont(font);
	}
}
