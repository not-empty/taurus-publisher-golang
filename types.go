package bullpublisher

type Options struct {
	Attempts           int    `json:"attempts"`
	Backoff            int    `json:"backoff"`
	Delay              int64  `json:"delay"`
	Lifo               bool   `json:"lifo"`
	PreventParsingData bool   `json:"preventParsingData"`
	Priority           int    `json:"priority"`
	RemoveOnComplete   int    `json:"removeOnComplete"`
	RemoveOnFail       int    `json:"removeOnFail"`
	Timeout            int    `json:"timeout"`
	JobId              string `json:"jobId"`
	Timestamp          int64  `json:"timestamp"`
}
