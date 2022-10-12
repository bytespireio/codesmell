package golang

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
say we have 2 files
file 1:
-------
#id,name
1,vishnu
2,mitesh
3, varun

file 2:
-------
#name,role
vishnu,senior data engineer
mitesh,principal engineer

we want to a join on name which results in

1,vishnu,senior data engineer
2,mitesh,principal engineer

*/
func Join2files(filePath1, filePath2 string) error {

	// get the num of records
	file, _ := os.Open(filePath1)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	f1Count := lineCount

	// get the num of records
	file, _ = os.Open(filePath2)
	fileScanner = bufio.NewScanner(file)
	lineCount = 0
	for fileScanner.Scan() {
		lineCount++
	}
	f2Count := lineCount

	dataMap := make(map[string]string, 0)
	outputList := make([]string, 0)

	if f1Count <= f2Count {
		f, err := os.Open(filePath1)
		if err != nil {
			//
		}
		defer f.Close()

		fileScanner := bufio.NewScanner(f)
		fileScanner.Split(bufio.ScanLines)

		i := 0
		nameIndex := 0
		otherColIdx := 1
		for fileScanner.Scan() {
			line := fileScanner.Text()
			if i == 0 {
				//first line
				h := strings.TrimPrefix(line, "#")
				cols := strings.Split(h, ",")
				for i, col := range cols {
					if col == "name" {
						nameIndex = i
						if nameIndex == 1 {
							otherColIdx = 0
						}
						break
					}
				}
			} else {
				//other lines
				ss := strings.Split(line, ",")
				dataMap[ss[nameIndex]] = ss[otherColIdx]
			}
			i++
		}
		//read next file

		f, err = os.Open(filePath2)
		if err != nil {
			//
		}
		defer f.Close()
		fileScanner = bufio.NewScanner(f)
		fileScanner.Split(bufio.ScanLines)

		//reset idx
		nameIndex = 0
		otherColIdx = 1
		i = 0
		for fileScanner.Scan() {
			line := fileScanner.Text()
			if i == 0 {
				//first line
				h := strings.TrimPrefix(line, "#")
				cols := strings.Split(h, ",")
				for i, col := range cols {
					if col == "name" {
						nameIndex = i
						if nameIndex == 1 {
							otherColIdx = 0
						}
						break
					}
				}
			} else {
				//other lines
				ss := strings.Split(line, ",")
				name := ss[nameIndex]
				value := ss[otherColIdx]
				if v, ok := dataMap[name]; ok {
					outputList = append(outputList, name+","+v+","+value)
				}
			}
			i++
		}

	} else {
		f, err := os.Open(filePath2)
		if err != nil {
			//
		}
		defer f.Close()

		fileScanner := bufio.NewScanner(f)
		fileScanner.Split(bufio.ScanLines)

		i := 0
		nameIndex := 0
		otherColIdx := 1
		for fileScanner.Scan() {
			line := fileScanner.Text()
			if i == 0 {
				//first line
				h := strings.TrimPrefix(line, "#")
				cols := strings.Split(h, ",")
				for i, col := range cols {
					if col == "name" {
						nameIndex = i
						if nameIndex == 1 {
							otherColIdx = 0
						}
						break
					}
				}
			} else {
				//other lines
				ss := strings.Split(line, ",")
				dataMap[ss[nameIndex]] = ss[otherColIdx]
			}
			i++
		}
		//read next file

		f, err = os.Open(filePath1)
		if err != nil {
			//
		}
		defer f.Close()
		fileScanner = bufio.NewScanner(f)
		fileScanner.Split(bufio.ScanLines)

		//reset idx
		nameIndex = 0
		otherColIdx = 1
		i = 0
		for fileScanner.Scan() {
			line := fileScanner.Text()
			if i == 0 {
				//first line
				h := strings.TrimPrefix(line, "#")
				cols := strings.Split(h, ",")
				for i, col := range cols {
					if col == "name" {
						nameIndex = i
						if nameIndex == 1 {
							otherColIdx = 0
						}
						break
					}
				}
			} else {
				//other lines
				ss := strings.Split(line, ",")
				name := ss[nameIndex]
				value := ss[otherColIdx]
				if v, ok := dataMap[name]; ok {
					outputList = append(outputList, name+","+v+","+value)
				}
			}
			i++
		}

	}

	fmt.Println("----------------")
	for _, val := range outputList {
		fmt.Println(val)
	}
	return nil
}
