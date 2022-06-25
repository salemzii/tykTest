package files

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Api_Id string `json:"api_id"`
	Hits   uint   `json:"hits"`
}

var (
	datals []Data
)

func Reader() []Data {
	// open file
	file, err := os.Open("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		MakeDataCopy(ParseLine(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return datals
}

func ParseLine(line string) (api_id, hits string) {
	/*
		Given each line of json extracted from data.txt is represented in string,
		and the built-in strconv.Unquote() doesn't cover such use case;

		ParseLine receives each line gotten from data.txt, and formats the said line
		using various delimiters enclosed in each string literal.
		And returns the require API_ID and HITS
	*/

	TrimmedStr := strings.Trim(line, "{}")
	splitstr := strings.Split(TrimmedStr, ":")

	apid := strings.Split(splitstr[1], ",")[0]
	hit := splitstr[2]

	return apid, hit
}

// Recieves individual api_id and hits value and append to a slice of type Data
func MakeDataCopy(api_id, hits string) []Data {
	hit, _ := strconv.Atoi(hits)

	datals = append(datals, Data{Api_Id: api_id, Hits: uint(hit)})
	return datals
}

//https://teivah.medium.com/a-closer-look-at-go-sync-package-9f4e4a28c35a
