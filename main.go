package main
import(
	"fmt"
	"os"
	"log"
	"strings"
)
import rl "github.com/gen2brain/raylib-go/raylib"
const (
	WIDTH = 800;
	HEIGHT = 600;
)
var (
	mainFile OpenedFile;
)

type OpenedFile struct {
	fullpath string;
	fileName string;
	data []byte;
}

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path);
	if err != nil {
		log.Panic(err);
	}
	return data;
}

func main() {
	if len(os.Args) == 2 {
		mainFile.fullpath = os.Args[1];
		pathSlashtokenized := strings.Split(mainFile.fullpath, "/");
		mainFile.fileName = pathSlashtokenized[len(pathSlashtokenized) - 1];
		mainFile.data = ReadFile(mainFile.fullpath);
	} else if len(os.Args) > 2 {
		fmt.Println("pera: Usage: pera [FILE]")
		os.Exit(2);
	}

	rl.SetTraceLogLevel(rl.LogError);
	icon := rl.LoadImage("resources/icon.png");
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable);
	rl.InitWindow(WIDTH, HEIGHT, "pera");
	rl.SetWindowIcon(*icon);

	if len(os.Args) == 1 {
		for !rl.WindowShouldClose() {
			rl.BeginDrawing();
			warningMessage := "You need to specify input file in console arguments"
			rl.DrawText(warningMessage, int32(rl.GetScreenWidth()/2) - rl.MeasureText("You need to specify input file in console arguments", 30)/2, int32(rl.GetScreenHeight()/2), 30, rl.Red)
			rl.EndDrawing();
		}
	}

	renderLoop();

	rl.UnloadImage(icon);
	rl.CloseWindow();
}
