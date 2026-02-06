package main
import rl "github.com/gen2brain/raylib-go/raylib"

var (
	zoom float32 = 15
)

func renderLoop() {
	for !rl.WindowShouldClose() {
		textPosX := int32(rl.GetScreenWidth()/2) - rl.MeasureText(mainFile.fileName, 20)/2

		rl.BeginDrawing();

		rl.ClearBackground(rl.RayWhite);
		rl.DrawText(mainFile.fileName, textPosX, 30, 20, rl.Black);
		rl.DrawFPS(WIDTH-30, 20);

		if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyEqual) {
			zoom++;
		} else if rl.IsKeyDown(rl.KeyLeftControl) && rl.IsKeyPressed(rl.KeyMinus) {
			zoom--;
		}

		font := rl.LoadFont("resources/comic-sans.ttf");
		rl.DrawTextEx(font, string(mainFile.data), rl.Vector2{0,0}, zoom, 1, rl.Black)

		rl.EndDrawing();
	}
}
