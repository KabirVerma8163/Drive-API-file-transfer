package main

import(
	"fmt"
	"os"
	"path/filepath"
	"net/http"
	// "github.com/h2non/filetype"
	// "io/ioutil"
)


func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DetectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func filing(){
	// var files [][]string;
	// var file string
	// var filedata string;
	files := make(map[string]os.FileInfo)

	root := "/home/toastedwaffle8163/Desktop/PracticeFiles";
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			files[path] = info
			// fmt.Println(path)
			return nil
	})
	if err != nil {
			panic(err)
	}

	for path, _ := range files{ // the  _ is supposed to be 'info'
		fmt.Println(path)
		// info, _:= os.Stat(path)
		// if info.IsDir(){
		// 	fmt.Println("This is a directory. ")
		// } else {
		// 	buf, _ := ioutil.ReadFile(path)

		// 	kind, _ := filetype.Match(buf)
			
		// 	if kind == filetype.Unknown {
		// 		fmt.Println("Unknown file type")
		// 	} else{
		// 		fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
		// 	}
		// }

		// f, err := os.Open(path)
		// if err != nil {
		// 	panic(err)
		// }
		// defer f.Close()

		// contentType, err := GetFileContentType(f)
		// if err != nil {
		// 	panic(err)
		// }
	
		// fmt.Println("Content Type: " + contentType)
		// fmt.Println(info.ModTime())
		// info, _:= os.Stat(path)
		// if info.IsDir(){
		// 	fmt.Println("This is a directory. ")
		// } else{
			f, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			contentType, err := GetFileContentType(f)
			if err != nil {
				panic(err)
			}
		
			fmt.Println("Content Type: " + contentType)
		}
	}

	// fmt.Print("This thing works")
	// var firstName, secondName string

	// fmt.Print("What is  your first name? ")
	// fmt.Scan(&firstName)
	// fmt.Println("What is your second name? ")
	// fmt.Scan(&secondName)
	// fmt.Printf("Your full name is %s %s.", firstName, secondName)