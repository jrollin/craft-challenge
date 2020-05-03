package domain

type StoryScore struct {
	Score   uint8
	StoryID StoryID
}

func NewStoryScore(score uint8, storyID StoryID) *StoryScore {
	return &StoryScore{Score: score, StoryID: storyID}
}
