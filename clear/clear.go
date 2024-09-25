package clear

import (
	"fmt"
	"log"
	"log-clear/utils"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// DeleteLog 删除过期文件
// /root/apps/node-deploy/.local/bsc/node0
func DeleteLog() error {
	var logDir = "/log"
	files, err := os.ReadDir(logDir)
	if err != nil {
		log.Printf("Error reading logDir: %s\n", err)
		return err
	}

	var fileMap = make(map[int64]os.DirEntry)
	var fileIndexSli utils.Int64Slice

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			log.Printf("Error getting file info: %s\n", err)
			continue
		}
		fileName := fileInfo.Name()
		if !strings.HasPrefix(fileName, "bsc.log.") {
			continue
		}
		fileNameSlic := strings.Split(fileName, ".log.")
		dateIndex := strings.Split(fileNameSlic[1], "_")
		timeStamp, err := parseTime(dateIndex[0])
		if err != nil {
			return err
		}
		i, err := strconv.Atoi(dateIndex[1])
		if err != nil {
			return err
		}
		index := timeStamp + int64(i)

		fileMap[index] = file
		fileIndexSli = append(fileIndexSli, index)
	}

	sort.Sort(fileIndexSli)

	maxNum := fileIndexSli.Len()

	for i, k := range fileIndexSli {
		if i < maxNum-1 {
			err = os.Remove(logDir + "/" + fileMap[k].Name())
			if err != nil {
				log.Printf("delete file error: %s\n", err)
				return err
			}
		}
	}
	return nil
}

func parseTime(date string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, fmt.Sprintf("%s 00:00:00", date))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return t.Unix(), nil
}
