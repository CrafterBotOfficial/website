package services

import (
	"fmt"
	"log"
	"slices"
	"time"
)

type VideoData struct {
	Id			int
	Title 		string
	DownloadUrl string
	Summary 	string
	Date		time.Time
}

var cachedVideoRows []VideoData

func GetChunkedVideos(count int, chunk int) ([][]VideoData, error) {
	allVideos, err := GetVideos()
	if err != nil {
		return nil, err
	}

	chunks := slices.Collect(slices.Chunk(allVideos, count))
	if len(chunks) < chunk {
		return nil, &IndexOutOfRange { message: "Index out of range" }
	}

	return chunks, nil
}

func GetVideos() ([]VideoData, error) {
	if cachedVideoRows == nil {
		db := GetDatabase()

		rows, err := db.Query("SELECT * FROM trailcam")
		if err != nil {
			log.Println(err)
			return nil, err
		}

		defer rows.Close()

		var r []VideoData
		for rows.Next() {
			var b VideoData 
			rows.Scan(&b.Id, &b.Title, &b.DownloadUrl, &b.Summary, &b.Date)
			b.DownloadUrl = fmt.Sprintf("/trailcam/%s", b.DownloadUrl)
			r = append(r, b)
		}
		cachedVideoRows = r
		return r, nil
	}
	return cachedVideoRows, nil
}

type IndexOutOfRange struct {
	arg 	int
	message string
}

func (e *IndexOutOfRange) Error() string {
	return "Index out of range"
}
