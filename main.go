package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//	*/lib
	//	*/server/lib/
	//check to see if the file exists in the path
	//eventually use regex up until the */lib or */server/lib directory (using my full path for testing)
	paths := []string{"/usr/local/Cellar/tomcat/8.5.13/libexec/server/lib/catalina.jar", "/usr/local/Cellar/tomcat/8.5.13/libexec/lib/catalina.jar"}
	si := "org.apache.catalina.util.ServerInfo"
	p := filePath(paths)
	cmd := exec.Command("java", "-classpath", p, si)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	fmt.Println(
		"**************",
	)

	testSlice := strings.Split(outStr, "\n")
	var m = make(map[string]string)

	for _, pair := range testSlice {
		z := strings.Split(pair, ":")
		if z[0] != "" {
			m[strings.TrimSpace(z[0])] = strings.TrimSpace(z[1])
		}
	}
	fmt.Println("Map:", m)
	fmt.Println("Key: Server number;", "Value:", m["Server number"])

	//Start doing version comparison
	foundVersion := m["Server number"]
	//We shouldn't need to check dot releases as our documentation
	supportedVersions := []string{"5.5", "6", "7", "9"}
	fmt.Println(supportedVersions[0])
	fv := strings.Split(foundVersion, ".")
	fmt.Println(fv[0])
	fmt.Println(fv[0] < supportedVersions[0])
	//for i, val := range supportedVersions {

	//}

}

func filePath(s []string) string {
	current := s[0]
	for i := range s {
		if _, err := os.Stat(s[i]); !os.IsExist(err) {
			fmt.Println(err)
			current = s[i]

		}

	}
	return current

}

//Now returning based on an address that matches

//func inRange(t string) bool {
//	pf := strings.Split(t, ".")
// 	piecesSupport := strings.Split(supportedVersions, ".")

// 	for i, val := range piecesSupport {
// 		if val == 'x' {
// 			return true
// 		}
// 		if val != piecesFound[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
