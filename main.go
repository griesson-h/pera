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
	} else if len(os.Args) > 2 {
		fmt.Println("pera: Usage: pera [FILE]")
		os.Exit(2);
	}
	mainFile.data = ReadFile(mainFile.fullpath);

	rl.SetTraceLogLevel(rl.LogError);
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable);
	rl.InitWindow(WIDTH, HEIGHT, "title");

	renderLoop();
}
