package main

import (
	"github.com/Devopsengineer75/copyfastgo/internal/fssync"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "copyfast",
		Short: "Sync folder to target",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			origin := args[0]
			target := args[1]
			fssync.CopyRecursive(origin, target)
		},
	}
	rootCmd.Execute()
}

/*import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func scan_recursive(dir_path string, ignore []string) ([]string, []string) {

	folders := []string{}
	files := []string{}

	// Scan
	filepath.Walk(dir_path, func(path string, f os.FileInfo, err error) error {

		_continue := false

		// Loop : Ignore Files & Folders
		for _, i := range ignore {

			// If ignored path
			if strings.Index(path, i) != -1 {

				// Continue
				_continue = true
			}
		}

		if _continue == false {

			f, err = os.Stat(path)

			// If no error
			if err != nil {
				log.Fatal(err)
			}

			// File & Folder Mode
			f_mode := f.Mode()

			// Is folder
			if f_mode.IsDir() {

				// Append to Folders Array
				folders = append(folders, path)

				// Is file
			} else if f_mode.IsRegular() {

				// Append to Files Array
				files = append(files, path)
			}
		}

		return nil
	})

	return folders, files
}

func main() {

	folders, files := scan_recursive("/Users/aur95/copyfast", []string{".git", "/.git", "/.git/", ".gitignore", ".DS_Store", ".idea", "/.idea/", "/.idea"})

	// Files
	for _, i := range files {
		fmt.Println(i)
	}

	// Folders
	for _, i := range folders {
		fmt.Println(i)
	}
}*/

/*func main() {
	a := "a" // 1Mo
	Hello(&a)
	// a 1.5Mo
	fmt.Println(a)
}

//using only pointer interaction on variables and get a memory gain
func Hello(s *string) *string { // 1Mo
	*s = "b"
	return s // 1Ko
}*/

/*package main

import (
	"fmt"
	"io/ioutil"
)

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

func main() {
	/ / Traverse all file names
	var s []string
	s, _ = GetAllFile("/root/go/src/test", s)

	fmt.Printf("slice: %v", s)
*/

// 1.5 Mo + 1Ko => 1.5Mo

/*import (
	"fmt"

	"github.com/Devopsengineer75/copyfastgo/internal/lib"
	"github.com/Devopsengineer75/copyfastgo/internal/zoo"
)

func main() {

	//var name string //nil
	//var name2 string = "Florent"
	//var name3 = "Florent"
	name4 := "Florent"
	age := 30

	fmt.Println("Hello " + name4 + " j'ai " + lib.ConvertIntToString(age) + " ans")

	fmt.Println(lib.Division(3, 0))
	fmt.Println(lib.Division(3, 2))

	chien := zoo.MakeChien("Olaf", "Woaf !")
	chat := zoo.MakeChien("Gros Minet", "Miaou !")
	oiseau := zoo.MakeChien("titi", "Piou !")
	//fmt.Println(chien, chien.Pattes)
	//chien.Say()

	adopte(chien)
	adopte(chat)
	adopte(oiseau)
}

func adopte(animal zoo.IsAnimal) {
	switch a := animal.(type) {
	case *zoo.Chat:
		fmt.Println("Super je peux adopter", a.Name)
	case *zoo.Chien:
		fmt.Println("Super je peux adopter", a.Name)
	case *zoo.Oiseau:
		fmt.Println("Super je peux adopter", a.Name)
	}

}*/
