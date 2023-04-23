package main

import (
	"log"
	"os"

	"github.com/hashicorp/logutils"
	"github.com/mothbe/httpclient/lib"
	"github.com/urfave/cli"
)

func main() {
	// os.Args = []string{"httpclient", "--url", "https://google.com/"}

	app := cli.NewApp()
	app.Name = "httpclient"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "url", Value: "https://google.com", Usage: "URL", Required: false},
		cli.StringFlag{Name: "loglevel", Value: "WARN", Usage: "Mininum loglevel: DEBUG, WARN, ERROR, INFO", Required: false},
		cli.StringFlag{Name: "dir", Usage: "Path to file or directory", Required: true},
	}
	app.Action = func(c *cli.Context) error {
		filter := &logutils.LevelFilter{
			Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
			MinLevel: logutils.LogLevel(c.String("loglevel")),
			Writer:   os.Stderr,
		}
		log.SetOutput(filter)

		dir := c.String("dir")
		txt_files := lib.Get_text_files(dir)
		log.Printf("[INFO] Text files: %s", txt_files)
		// fmt.Println(filepath.Base(dir))
		// fmt.Println(filepath.Dir(dir))

		url := c.String("url")
		log.Printf("[INFO] Checking: %v\n", url)

		// a := read_line_by_line("google.com.txt")
		for _, f := range txt_files {
			a := lib.Read_line_by_line(f)
			for url, status_code := range a {
				log.Printf("[DEBUG] url: %s, status Code: %d", url, status_code)
				if lib.GetStatusCode(url) == status_code {
					log.Printf("[INFO] Status code math for %s - %d", url, status_code)
				} else {
					log.Printf("[ERROR] No match for %s - %d", url, status_code)
				}
			}
		}

		return nil
	}
	app.UsageText = "httpclient --url https://www.google.com/"
	app.Author = "mars3n"
	app.Email = "mars3n@protonmail.com"
	_ = app.Run(os.Args)
}
