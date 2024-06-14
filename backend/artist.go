package backend

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func InfoAboutArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists/" {
		err := "404 Page not found"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		err := "405 method not allowed"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusMethodNotAllowed)
		return
	}
	artists, err := GetAllArtists()
	if err != nil {
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if id <= 0 || id > len(artists) || err != nil {
		err := "404 Page not found"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusNotFound)
		return
	}
	infoAboutOne, err := OneArtist(id)
	if err != nil {
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
	rel, err := Relations(id)
	if err != nil {
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
	artist := ArtistInfo{infoAboutOne, rel}
	err = tmp.ExecuteTemplate(w, "artist.html", artist)
	if err != nil {
		err := "500 Internal Server Error"
		fmt.Println(err)
		ErrorPage(w, err, http.StatusInternalServerError)
		return
	}
}


func CreationDate(data Everything, range1 string, range2 string) (Everything, error) {
	var output Everything
	year1, err := strconv.Atoi(range1)
	if err != nil {
		fmt.Println("error in years")
		return output, err
	}
	year2, err := strconv.Atoi(range2)
	if err != nil {
		fmt.Println("error in years")
		return output, err
	}
	var artists []Artist
	for _, result := range data.Everyone {
		if year1 <= result.CreationDate && year2 >= result.CreationDate {
			artists = append(artists, result)
		}
	}

	// fmt.Println(artists)
	output.Everyone = artists
	return output, nil
	//}
}

func FirstAlbumDates(data Everything, range1 string, range2 string) (Everything, error) {
	var output Everything
	year1, err := strconv.Atoi(range1)
	if err != nil {
		fmt.Println("error in years")
		return output, err
	}
	year2, err := strconv.Atoi(range2)
	if err != nil {
		fmt.Println("error in years")
		return output, err
	}
	var artists []Artist
	for _, result := range data.Everyone {
		data := strings.Split(result.FirstAlbum, "-")
		dataNumber, err := strconv.Atoi(data[len(data)-1])
		if err != nil {
			fmt.Println("error in years")
			return output, err
		}
		if year1 <= dataNumber && year2 >= dataNumber {
			artists = append(artists, result)
		}
	}

	// fmt.Println(artists)
	output.Everyone = artists
	return output, nil
	//}
}

func MembersNumber(data Everything, arr []string) (Everything, error) {
	var output Everything
	var memberNum []int
	for _, val := range arr {
		n, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("error in years")
			return output, err
		}
		memberNum = append(memberNum, n)
	}
	var artists []Artist
	for _, result := range data.Everyone {
		for _, val := range memberNum {
			if len(result.Members) == val {
				artists = append(artists, result)
			}
		}
	}

	// fmt.Println(artists)
	output.Everyone = artists
	return output, nil
	//}
}

func LocationSearch(data Everything, search string) (Everything, error) {
	var output Everything
	var artists []Artist
	ids := make(map[int]int)
	for _, result := range data.Location.Ind {
		for _, location := range result.Locations {
			if strings.Contains(strings.ToLower(location), strings.ToLower(search)) {
				if _, ok := ids[result.ID]; ok {
					continue
				} else {
					art, err := OneArtist(result.ID)
					if err != nil {
						fmt.Println("error in getting artist")
						return output, err
					}
					artists = append(artists, art)
					ids[result.ID] += 1
				}
			}
		}
	}
	output.Everyone = artists
	return output, nil
	//}
}