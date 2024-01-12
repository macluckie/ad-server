package domain

type IdAd string

type Ad struct {
	Id          string `validate:"required"`
	Title       string `validate:"required"`
	Description string
	URL         string `validate:"required"`
}

type AdServe struct {
	Url             string
	TrackImpression int64
}
