package main

// Type for Gfycat API responses

type GfycatReaction struct {
	Tags               []string `json:"tags"`
	LanguageCategories []string `json:"languageCategories"`
	GifUrl             string   `json:"gifUrl"`
	Views              int64    `json:"views"`
	Likes              int64    `json:"likes"`
	Dislikes           int64    `json:"dislikes"`
	GfyId              string   `json:"gfyId"`
	GfyName            string   `json:"gfyName"`
	AvgColor           string   `json:"avgColor"`
	Width              int      `json:"width"`
	Height             int      `json:"height"`
	Framerate          int      `json:"frameRate"`
	NumFrames          int64    `json:"numFrames"`
	GifSize            int64    `json:"gifSize"`
}

type GfycatReactions struct {
	Cursor  string           `json:"string"`
	Gfycats []GfycatReaction `json:"gfycats"`
}
