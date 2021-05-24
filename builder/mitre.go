package builder

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func MitreLookup(mitreID string) []string {
	var empty []string
	mitreID = "Rob"
	// NOTES: column names could be either techID, tactName
	// workflow: check if IDs in input array are subtechnique, if so, query csv for technique (prefix)
	// return tactName
	// IF ID in input array is technique (no dot to split ID on), query csv, return tactName

	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("records is type: %T\ncontents: %v\n", records, records)
	for _, entry := range records {
		for _, value := range entry {
			if value == mitreID {
				return entry
			}
		}
	}
	return empty
}
