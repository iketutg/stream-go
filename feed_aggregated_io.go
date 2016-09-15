package getstream

type postAggregatedFeedOutputActivities struct {
	Activities []*Activity `json:"activities"`
}

// GetAggregatedFeedInput is used to Get a list of Activities from a AggregatedFeed
type GetAggregatedFeedInput struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	IDGTE string `json:"id_gte,omitempty"`
	IDGT  string `json:"id_gt,omitempty"`
	IDLTE string `json:"id_lte,omitempty"`
	IDLT  string `json:"id_lt,omitempty"`

	Ranking string `json:"ranking,omitempty"`
}

// GetAggregatedFeedOutput is the response from a AggregatedFeed Activities Get Request
type GetAggregatedFeedOutput struct {
	Duration string
	Next     string
	Results  []*struct {
		Activities    []*Activity
		ActivityCount int
		ActorCount    int
		CreatedAt     string
		Group         string
		ID            string
		UpdatedAt     string
		Verb          string
	}
}

type getAggregatedFeedOutput struct {
	Duration string                           `json:"duration"`
	Next     string                           `json:"next"`
	Results  []*getAggregatedFeedOutputResult `json:"results"`
}

func (a getAggregatedFeedOutput) output() *GetAggregatedFeedOutput {

	output := GetAggregatedFeedOutput{
		Duration: a.Duration,
		Next:     a.Next,
	}

	var results []*struct {
		Activities    []*Activity
		ActivityCount int
		ActorCount    int
		CreatedAt     string
		Group         string
		ID            string
		UpdatedAt     string
		Verb          string
	}

	for _, result := range a.Results {

		outputResult := struct {
			Activities    []*Activity
			ActivityCount int
			ActorCount    int
			CreatedAt     string
			Group         string
			ID            string
			UpdatedAt     string
			Verb          string
		}{
			ActivityCount: result.ActivityCount,
			ActorCount:    result.ActorCount,
			CreatedAt:     result.CreatedAt,
			Group:         result.Group,
			ID:            result.ID,
			UpdatedAt:     result.UpdatedAt,
			Verb:          result.Verb,
		}

		for _, activity := range result.Activities {
			outputResult.Activities = append(outputResult.Activities, activity)
		}

		results = append(results, &outputResult)
	}

	output.Results = results

	return &output
}

type getAggregatedFeedOutputResult struct {
	Activities    []*Activity `json:"activities"`
	ActivityCount int         `json:"activity_count"`
	ActorCount    int         `json:"actor_count"`
	CreatedAt     string      `json:"created_at"`
	Group         string      `json:"group"`
	ID            string      `json:"id"`
	UpdatedAt     string      `json:"updated_at"`
	Verb          string      `json:"verb"`
}

type getAggregatedFeedFollowersInput struct {
	Limit int `json:"limit"`
	Skip  int `json:"offset"`
}

type getAggregatedFeedFollowersOutput struct {
	Duration string                                    `json:"duration"`
	Results  []*getAggregatedFeedFollowersOutputResult `json:"results"`
}

type getAggregatedFeedFollowersOutputResult struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	FeedID    string `json:"feed_id"`
	TargetID  string `json:"target_id"`
}

type postAggregatedFeedFollowingInput struct {
	Target            string `json:"target"`
	ActivityCopyLimit int    `json:"activity_copy_limit"`
}
