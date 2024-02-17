package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func RetrieveParameterValues(inputArg string) ([]string, error) {
	if fileInfo, err := os.Stat(inputArg); err == nil && !fileInfo.IsDir() {
		file, err := os.Open(inputArg)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines, scanner.Err()
	}

	return strings.Split(inputArg, ","), nil
}

func ExecuteHPPTests(targetURL, parameterKey string, parameterValues []string) {
	var waitGroup sync.WaitGroup
	paramCombinations := [][]string{
		{parameterValues[0], parameterValues[1]},
		{parameterValues[1], parameterValues[0]},
		{parameterValues[0]},
		{parameterValues[1]},
	}

	for _, combo := range paramCombinations {
		waitGroup.Add(1)
		go func(combo []string) {
			defer waitGroup.Done()

			queryParams := url.Values{}
			for _, val := range combo {
				queryParams.Add(parameterKey, val)
			}

			fullRequestURL := fmt.Sprintf("%s?%s", targetURL, queryParams.Encode())
			response, err := http.Get(fullRequestURL)
			if err != nil {
				color.Red(fmt.Sprintf("Error ❌: Request to %s failed with error: %v", fullRequestURL, err))
				return
			}
			defer response.Body.Close()

			if response.StatusCode == 200 {
				color.Green(fmt.Sprintf("Success ✅: URL %s responded with status code %d", fullRequestURL, response.StatusCode))
			} else {
				color.Yellow(fmt.Sprintf("Warning ⚠️: URL %s responded with status code %d", fullRequestURL, response.StatusCode))
			}
		}(combo)
	}

	waitGroup.Wait()
}

func displayBanner(c *cli.Context) error {
	banner := `
	██╗  ██╗██████╗ ██████╗ ███████╗██╗   ██╗███████╗███████╗██████╗ ██╗   ██╗███████╗████████╗██████╗ ██████╗ 
	██║  ██║██╔══██╗██╔══██╗██╔════╝██║   ██║╚══███╔╝╚══███╔╝██╔══██╗██║   ██║██╔════╝╚══██╔══╝╚════██╗██╔══██╗
	███████║██████╔╝██████╔╝█████╗  ██║   ██║  ███╔╝   ███╔╝ ██████╔╝██║   ██║███████╗   ██║    █████╔╝██████╔╝
	██╔══██║██╔═══╝ ██╔═══╝ ██╔══╝  ██║   ██║ ███╔╝   ███╔╝  ██╔══██╗██║   ██║╚════██║   ██║    ╚═══██╗██╔══██╗
	██║  ██║██║     ██║     ██║     ╚██████╔╝███████╗███████╗██████╔╝╚██████╔╝███████║   ██║   ██████╔╝██║  ██║
	╚═╝  ╚═╝╚═╝     ╚═╝     ╚═╝      ╚═════╝ ╚══════╝╚══════╝╚═════╝  ╚═════╝ ╚══════╝   ╚═╝   ╚═════╝ ╚═╝  ╚═╝
	Fuzzing tool for HTTP Parameter Pollution vulnerabilities. Written By 1337-SIGMA
`
	fmt.Println(banner)
	return nil
}

func main() {
	appConfig := &cli.App{
		Name:   "HPPFuzZBu5t3R",
		Usage:  "Detects HTTP Parameter Pollution vulnerabilities.",
		Before: displayBanner,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "target",
				Aliases:  []string{"t"},
				Usage:    "Target URL for HPP testing.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "param",
				Aliases:  []string{"p"},
				Usage:    "Query parameter to test for pollution.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "data",
				Aliases:  []string{"d"},
				Usage:    "Values for testing, either comma-separated or file path.",
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {
			targetURL := context.String("target")
			parameterKey := context.String("param")
			inputArg := context.String("data")

			parameterValues, err := RetrieveParameterValues(inputArg)
			if err != nil {
				return err
			}

			ExecuteHPPTests(targetURL, parameterKey, parameterValues)
			return nil
		},
	}

	if err := appConfig.Run(os.Args); err != nil {
		color.Red("Execution Error: %v", err)
	}
}
