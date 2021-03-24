import(
	"fmt"
	// "bufio"
	"os"
	"path/filepath"
)

func files(){
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("What is your name? ")
	// // var text string

	// scanner.Scan()
	// text := scanner.Text()
	// fmt.Print(text)

	// scanner.Scan()
	// text = scanner.Text()
	// fmt.Print(text)


	// // var first string;
	// // first = "Kabir"
	// // fmt.Print("What is your name? ")
	// // // fmt.Scan
	// // fmt.Print("Your first name is ", first)
	var files []string
	fileMap := make(map[string]os.FileInfo)

	root := "C:/Users/theaw/Desktop/PractingFiles"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			fileMap[path] = info;
			return nil;
	})
	if err != nil {
			panic(err);
	}
	for _, file := range files {
		fmt.Println(file);
		fmt.Println(fileMap[file].ModTime())
	}
}