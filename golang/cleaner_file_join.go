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

func getSize(filePath string) (int64, error) {
	f, err := os.Stat(filePath)
	if err != nil {
		return -1, err
	}
	return f.Size(), nil
}

func locateNameColumn(line string) int {
	header := strings.TrimPrefix(line, "#")
	cols := strings.Split(header, ",")
	nameIdx := 0
	for i, col := range cols {
		if col == "name" {
			nameIdx = i
			break
		}
	}
	return nameIdx
}

func prepareDataMap(filePath string) (map[string]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dataMap := make(map[string]string, 0)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	nameIndex := 0
	otherColIdx := 1

	lineNum := 1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if lineNum == 1 {
			//first line
			nameIndex = locateNameColumn(line)
			if nameIndex == 1 {
				otherColIdx = 0
			}
		} else {
			//other lines
			splits := strings.Split(line, ",")
			dataMap[splits[nameIndex]] = splits[otherColIdx]
		}
		lineNum++
	}

	return dataMap, nil
}

func join(dataMap map[string]string, otherFilePath string) ([]string, error) {
	f, err := os.Open(otherFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	outputList := make([]string, 0)

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	nameIndex := 0
	otherColIdx := 1
	lineNum := 1

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if lineNum == 1 {
			//first line
			nameIndex = locateNameColumn(line)
			if nameIndex == 1 {
				otherColIdx = 0
			}
		} else {
			//other lines
			splits := strings.Split(line, ",")
			name := splits[nameIndex]
			otherColVal := splits[otherColIdx]
			if v, ok := dataMap[name]; ok {
				outputList = append(outputList, name+","+v+","+otherColVal)
			}
		}
		lineNum++
	}
	return outputList, nil
}

func processOutput(outputList []string) {
	fmt.Println("----------------")
	for _, val := range outputList {
		fmt.Println(val)
	}
}

func Join2filesCleanly(filePath1, filePath2 string) error {

	f1size, err1 := getSize(filePath1)
	if err1 != nil {
		panic(err1.Error())
	}
	f2size, err2 := getSize(filePath2)
	if err2 != nil {
		panic(err2.Error())
	}

	fpath1 := filePath1
	fpath2 := filePath2

	if f1size > f2size {
		fpath1 = filePath2
		fpath2 = filePath1
	}

	dataMap, err := prepareDataMap(fpath1)
	if err != nil {
		panic(err.Error())
	}

	if outputList, err := join(dataMap, fpath2); err != nil {
		panic(err.Error())
	} else {
		processOutput(outputList)
	}

	return nil
}
