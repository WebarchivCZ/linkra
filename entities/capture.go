package entities

import "golang.org/x/text/language"

// Information necessary for capturing/crawling a page.
type CaptureRequest struct {
	// The URL address we want to capture.
	SeedURL string `json:"seedURL"`
	// The ShadowID of Seed we want to capture. This will be used in CaptureResult.
	SeedShadowID string `json:"seedShadowID"`
	// The status of the request.
	State CaptureState `json:"state"`
}

func NewRequestFromSeed(seed *Seed) *CaptureRequest {
	return &CaptureRequest{
		SeedURL:      seed.URL,
		SeedShadowID: seed.ShadowID,
		State:        NotEnqueued,
	}
}

type CaptureState string

const (
	// The request is newly created and is not enqueued.
	NotEnqueued CaptureState = "NotEnqueued"
	// The request was enqueued for processing.
	Pending CaptureState = "Pending"
	// The capture was successful.
	DoneSuccess CaptureState = "DoneSuccess"
	// The capture failed.
	DoneFailure CaptureState = "DoneFailure"
)

func (state CaptureState) IsCaptureState() bool {
	return state == NotEnqueued ||
		state == Pending ||
		state == DoneSuccess ||
		state == DoneFailure
}

type CaptureResult struct {
	// ShadowID of the seed to which the result belongs to.
	SeedShadowID string `json:"seedShadowID"`
	// Was the capture completed
	Done bool `json:"done"`
	// Received errors
	ErrorMessages []string `json:"errorMessages"`

	CaptureMetadata *CaptureMetadata `json:"captureMetadata"`
}

type CaptureMetadata struct {
	// CDXJ timestamp of the instant the capture was taken as recorded in index/WARC https://specs.webrecorder.net/cdxj/0.1.0/#timestamp
	Timestamp string `json:"timestamp"`
	// URL from CDXJ JSON block.
	CapturedUrl string `json:"capturedUrl"`
}

func PrettyPrintCaptureState(state CaptureState, lang language.Tag) string {
	czechStates := map[CaptureState]string{
		NotEnqueued: "Nezařazeno",
		Pending:     "Čeká na archivaci",
		DoneSuccess: "Úspěšně archivováno",
		DoneFailure: "Chyba při archivaci",
	}
	englishStates := map[CaptureState]string{
		NotEnqueued: "Not enqueued",
		Pending:     "Waiting for capture",
		DoneSuccess: "Success",
		DoneFailure: "Failure",
	}

	var prettyMap map[CaptureState]string
	var unknown string
	if lang == language.Czech {
		prettyMap = czechStates
		unknown = "Neznámý stav"
	} else {
		prettyMap = englishStates
		unknown = "Unknown state"
	}

	if pretty, ok := prettyMap[state]; ok {
		return pretty
	}
	return unknown
}
