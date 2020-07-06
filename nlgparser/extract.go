package nlgparser

import (
	"bufio"
	"github.com/golang-collections/collections/stack"
	"os"
	"strings"
)

const NewLineCharacter = '\n'
const SearchTerm = "macrodef"

func Extract(base, file, out string) {

	baseFile, err := os.Open(base)
	check(err)
	outFile, err := os.Create(out)
	check(err)
	nlgFile, err := os.Open(file)
	check(err)
	r := bufio.NewReader(baseFile)
	r2 := bufio.NewReader(nlgFile)
	w := bufio.NewWriter(outFile)
	m := fillMacrosAvailable(r2)
	defer nlgFile.Close()

	for {
		line, e := r.ReadString(NewLineCharacter)
		if e != nil {
			break
		}
		if strings.Contains(line, SearchTerm) {
			macroName := getMacroName(line)
			_, ok := m[macroName]
			if !ok {
				//fmt.Println("macro not present: ", macroName)
				chunk := line
				stk := stack.New()
				if strings.Contains(line, "{") {
					stk.Push("{")
				}
				for {
					b, e := r.ReadByte()
					if e != nil {
						break
					}
					s := string(b)
					chunk += s
					//fmt.Println(s)
					if s == "{" {
						stk.Push("{")
					} else if s == "}" {
						stk.Pop()
					}
					if stk.Len() == 0 {
						break
					}
				}
				chunk += string(NewLineCharacter)
				//fmt.Println(chunk)
				_, err := w.WriteString(chunk)
				check(err)
				w.Flush()
			}
		}
	}
	defer baseFile.Close()
	defer outFile.Close()
}

func getMacroName(line string) string {
	fields := strings.Fields(line)
	macroName := strings.ReplaceAll(fields[1], "{", "")
	return macroName
}

func fillMacrosAvailable(r *bufio.Reader) map[string]string {
	m := make(map[string]string)
	for {
		line, e := r.ReadString(NewLineCharacter)
		if e != nil {
			break
		}
		if strings.Contains(line, SearchTerm) {
			macroName := getMacroName(line)
			m[macroName] = macroName
		}
	}
	return m
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
