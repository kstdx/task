package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/kstdx/task/internal/logger"
	"github.com/kstdx/task/internal/path"
)

func main() {
	flag.Parse()
	args := flag.Args()

	realpath := path.Get("task.json")
	var taskData map[string]string

	if len(args) != 0 {
		taskName := args[0]

		if _, err := os.Stat(realpath); err == nil {
			raw, _ := os.ReadFile(realpath)

			if err := json.Unmarshal([]byte(raw), &taskData); err == nil {
				if _, ok := taskData[taskName]; ok {
					cmd := exec.Command("sh", "-c", taskData[taskName])

					logger.Info("Executing [" + taskData[taskName] + "]")

					stdout, err := cmd.StdoutPipe()
					if err != nil {
						logger.Err("Failed to get stdout pipes")
					}

					if err := cmd.Start(); err != nil {
						logger.Err("Failed to start that command")
					}

					scanner := bufio.NewScanner(stdout)
					for scanner.Scan() {
						fmt.Println(scanner.Text())
					}

					if err := cmd.Wait(); err != nil {
						logger.Err("Failed to execute that command")
					}

					logger.Ok("Succeeded to execute that command")
				} else {
					logger.Err("Requested cmd not found")
				}
			} else {
				logger.Err("Could not understand your json")
			}
		} else {
			logger.Err("Could not found the task.json file")
		}
	} else {
		if _, err := os.Stat(realpath); err == nil {
			raw, _ := os.ReadFile(realpath)
			if err := json.Unmarshal([]byte(raw), &taskData); err == nil {
				for name, cmd := range taskData {
					fmt.Println(name + ": " + cmd)
				}

				logger.Ok("Found " + strconv.Itoa(len(taskData)) + " tasks")
			} else {
				logger.Err("Could not understand your json")
			}
		} else {
			f, _ := os.Create("write.txt")
			f.Write([]byte(`{"hello", "echo Hello, world!"}`))

			logger.Ok("Created task.json!")
		}
	}
}
