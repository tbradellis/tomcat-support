package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//inclusive lower bound
//each element corresponds to the ordinal for the lower and upper bound versions
var lb = []int{5, 5}

//exclusive upper bound
var ub = []int{9, 0}

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
	if errStr != "" {
		fmt.Println("Error: ", errStr)
	}
	//fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	//outStr output from java -classpath
	testSlice := strings.Split(outStr, "\n")
	var m = make(map[string]string)

	for _, pair := range testSlice {
		z := strings.Split(pair, ":")
		if z[0] != "" {
			m[strings.TrimSpace(z[0])] = strings.TrimSpace(z[1])
		}
	}
	//fmt.Println("Map:", m)
	//fmt.Println("Key: Server number;", "Value:", m["Server number"])

	//Start doing version comparison
	foundVersion := m["Server number"]
	fmt.Println(supported(foundVersion, lb, ub))

}

func filePath(s []string) string {
	current := s[0]
	for i := range s {
		if _, err := os.Stat(s[i]); !os.IsExist(err) {
			//fmt.Println(err)
			if err != nil {
				fmt.Println("Info:", err)
			}
			current = s[i]

		}

	}
	return current

}

//handle the 2nd ordinal case of lower bound if 1st ordinal vf == 1st ordinal lb
//eleminate unsupported versions
//return supported if not eleminated
func supported(vf string, l []int, u []int) bool {

	fv := strings.Split(vf, ".")
	s, err := strconv.Atoi(fv[0])
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if s == l[0] {
		s2, err := strconv.Atoi(fv[1])
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if s2 < l[1] {
			return false
		}

	} else if s < l[0] || s >= u[0] {
		return false
	}
	return true

}
