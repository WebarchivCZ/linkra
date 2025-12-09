package components

import (
	"encoding/json"
	"linkra/entities"
	"path"
)

func CsvFilename(group *entities.SeedsGroup) string {
	return "export-" + group.ShadowID + ".csv"
}

func ExcelFilename(group *entities.SeedsGroup) string {
	return "export-" + group.ShadowID + ".xlsx"
}

func prettyPrintCaptureState(state entities.CaptureState) string {
	return entities.PrettyPrintCaptureState(state)
}

type seedJsonObject struct {
	Authors      []authorsJsonObject `json:"autoři"`
	URL          string              `json:"url"`
	Webarchive   string              `json:"webarchiv"`
	ArchivalURL  string              `json:"archivní-url"`
	ArchivalDate string              `json:"datum-archivace"`
}

type authorsJsonObject struct {
	Name    string `json:"jméno"`
	Surname string `json:"příjmení"`
}

func groupToJson(group *entities.SeedsGroup) string {
	objects := make([]seedJsonObject, 0, len(group.Seeds))
	for _, seed := range group.Seeds {
		if seed.State != entities.DoneSuccess {
			continue
		}
		seedObject := seedJsonObject{
			Authors:      []authorsJsonObject{{}},
			URL:          seed.URL,
			Webarchive:   "Webarchiv",
			ArchivalURL:  shortWaybackLink(seed),
			ArchivalDate: seed.HarvestedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		objects = append(objects, seedObject)
	}
	data, err := json.Marshal(objects)
	if err != nil {
		return `[{"error":"could not marshal json"}]`
	}
	return string(data)
}

func wrapJsonToScript(data string) string {
	return `<script type="application/json" id="input-data">` + data + `</script>`
}

// NOTE
// Be careful with using url.JoinPath. It does not work unless the base is valid URL with protocol.
// path.Join is better when joining just path segments. Especially when we don't know if they start and end with slashes.

// URL that will redirect to the archived page in wayback when the seed is captured
func shortWaybackLink(seed *entities.Seed) string {
	return Constants().GetFullURL(path.Join(Constants().GetWaybackRedirectPath(), seed.ShadowID))
}

// URL for the seeds detail view
func seedViewLink(seed *entities.Seed) string {
	return Constants().GetFullURL(path.Join(Constants().GetSeedPath(), seed.ShadowID))
}

// URL for the group view
func groupViewLink(group *entities.SeedsGroup) string {
	return Constants().GetFullURL(path.Join(Constants().GetGroupPath(), group.ShadowID))
}

// Get full path to static file "filename"
func fullStaticPath(filename string) string {
	return Constants().GetFullURL(path.Join(Constants().GetStaticPath(), filename))
}

func isCaptureCompleted(group *entities.SeedsGroup) bool {
	for _, seed := range group.Seeds {
		// We could instead test if state is Pending or NotEnqueued which would be more readable
		// but this will be more future proof and bug proof (we rather don't show correct seeds instead of showing incorrect ones)
		if seed.State != entities.DoneSuccess && seed.State != entities.DoneFailure {
			return false
		}
	}
	return true
}
