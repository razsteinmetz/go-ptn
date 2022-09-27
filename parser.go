package ptn

import (
	"encoding/json"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strconv"
	"strings"
)

// TorrentInfo is the resulting structure returned by Parse
// important, season/episode are 0,0 for movies (so you can't have S00E00 file!)
type TorrentInfo struct {
	Title      string `json:"title,omitempty"`
	Season     int    `json:"season,omitempty"`
	Episode    int    `json:"episode,omitempty"`
	Year       int    `json:"year,omitempty"`
	Resolution string `json:"resolution,omitempty"` //1080p etc
	Quality    string `json:"quality,omitempty"`
	Codec      string `json:"codec,omitempty"`
	Audio      string `json:"audio,omitempty"`
	Service    string `json:"service,omitempty"` // NF etc
	Group      string `json:"group,omitempty"`
	Region     string `json:"region,omitempty"`
	Extended   bool   `json:"extended,omitempty"`
	Hardcoded  bool   `json:"hardcoded,omitempty"`
	Proper     bool   `json:"proper,omitempty"`
	Repack     bool   `json:"repack,omitempty"`
	Container  string `json:"container,omitempty"`
	Widescreen bool   `json:"widescreen,omitempty"`
	Website    string `json:"website,omitempty"`
	Language   string `json:"language,omitempty"`
	Sbs        string `json:"sbs,omitempty"`
	Unrated    bool   `json:"unrated,omitempty"`
	Size       string `json:"size,omitempty"`
	Threed     bool   `json:"3d,omitempty"`
	Country    string `json:"country,omitempty"`
	IsMovie    bool   `json:"ismovie"` // true if this is a movie, false if tv show
}

func (t TorrentInfo) Tojson() (string, error) {
	s, e := json.Marshal(t)
	if e != nil {
		return "", e
	}
	return string(s), nil
}

func Efunc(t *TorrentInfo) (string, error) {
	s, e := json.Marshal(t)
	if e != nil {
		return "", e
	}
	return string(s), nil
}

func setField(tor *TorrentInfo, field, raw, val string) {
	// set the Field by reflecting its info
	ttor := reflect.TypeOf(tor)
	torV := reflect.ValueOf(tor)
	//field = strings.Title(field)
	// Title was deprecated, so need to use cases.
	caser := cases.Title(language.English)
	field = caser.String(field)
	v, _ := ttor.Elem().FieldByName(field)
	//fmt.Printf("    field=%v, type=%+v, value=%v\n", field, v.Type, val)
	switch v.Type.Kind() {
	case reflect.Bool:
		torV.Elem().FieldByName(field).SetBool(true)
	case reflect.Int:
		clean, _ := strconv.ParseInt(val, 10, 64)
		torV.Elem().FieldByName(field).SetInt(clean)
	case reflect.Uint:
		clean, _ := strconv.ParseUint(val, 10, 64)
		torV.Elem().FieldByName(field).SetUint(clean)
	case reflect.String:
		torV.Elem().FieldByName(field).SetString(val)
	}
}

// Parse breaks up the given filename in TorrentInfo
// algo - remove the file extention if its one of known, then parse the rest
// the title is the last part.
func Parse(filename string) (*TorrentInfo, error) {
	tor := &TorrentInfo{}
	//fmt.Printf("filename %q\n", filename)
	var startIndex, endIndex = 0, len(filename)
	// remove any underline and replace with Spaces
	cleanName := strings.Replace(filename, "_", " ", -1)
	if matches := container.FindAllStringSubmatch(cleanName, -1); len(matches) != 0 {
		tor.Container = matches[0][1]
		cleanName = cleanName[0 : len(cleanName)-4]
	} else if matches := otherExtensions.FindAllStringSubmatch(cleanName, -1); len(matches) != 0 {
		cleanName = cleanName[0 : len(cleanName)-4] // remove the . and the extension from the checked strings.
	}
	// go over all patterns
	for _, pattern := range patterns {
		matches := pattern.re.FindAllStringSubmatch(cleanName, -1)
		if len(matches) == 0 {
			continue
		}
		matchIdx := 0
		if pattern.last {
			// Take last occurrence of element.
			matchIdx = len(matches) - 1
		}
		//fmt.Printf("  %s: pattern:%q match:%#v\n", pattern.name, pattern.re, matches[matchIdx])

		index := strings.Index(cleanName, matches[matchIdx][1])
		if index == 0 {
			startIndex = len(matches[matchIdx][1])
			//fmt.Printf("    startIndex moved to %d [%q]\n", startIndex, filename[startIndex:endIndex])
		} else if index < endIndex {
			endIndex = index
			//fmt.Printf("    endIndex moved to %d [%q]\n", endIndex, filename[startIndex:endIndex])
		}
		setField(tor, pattern.name, matches[matchIdx][1], matches[matchIdx][2])
	}

	// Start process for title and remove all dots/underscore from it
	//fmt.Println("  title: <internal>")
	raw := strings.Split(filename[startIndex:endIndex], "(")[0]
	cleanName = raw
	if strings.HasPrefix(cleanName, "- ") {
		cleanName = raw[2:]
	}
	// clean out the title remove any starting chars
	cleanName = strings.Trim(cleanName, " -_.^/\\(){}[]")
	// only remove the dots if there are no spaces for some titles have dots and spaces
	if strings.ContainsRune(cleanName, '.') && !strings.ContainsRune(cleanName, ' ') {
		cleanName = strings.Replace(cleanName, ".", " ", -1)
	}
	//cleanName = strings.ReplaceAll(cleanName, ".", " ")
	cleanName = strings.ReplaceAll(cleanName, "_", " ")
	cleanName = strings.ReplaceAll(cleanName, "  ", " ")
	cleanName = strings.Trim(cleanName, " -_.^/\\(){}[]")
	//if strings.ContainsRune(cleanName, '.') && !strings.ContainsRune(cleanName, ' ') {
	//	cleanName = strings.Replace(cleanName, ".", " ", -1)
	//}
	//cleanName = strings.Replace(cleanName, "_", " ", -1)
	//cleanName = re.sub('([\[\(_]|- )$', '', cleanName).strip()
	if matches := countryre.FindAllStringSubmatch(cleanName, -1); len(matches) != 0 {
		tor.Country = matches[0][1]
		cleanName = cleanName[0 : len(cleanName)-3] // remove the coutnry from th name
	}
	setField(tor, "title", raw, cleanName)
	tor.IsMovie = tor.Episode == 0 && tor.Season == 0
	return tor, nil
}
