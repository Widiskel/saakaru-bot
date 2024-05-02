package model

import "time"

type QuestModel struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
	TotalTokenPrize string    `json:"totalTokenPrize"`
	Statistic       struct {
		TotalParticipants int   `json:"totalParticipants"`
		TotalPoints       int64 `json:"totalPoints"`
		TotalPosts        int   `json:"totalPosts"`
	} `json:"statistic"`
	Tasks []struct {
		ID          string    `json:"id"`
		Category    string    `json:"category"`
		Code        string    `json:"code"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ExternalURL string    `json:"externalUrl"`
		EndAt       time.Time `json:"endAt"`
		Criteria    []struct {
			Name                   string `json:"name"`
			Field                  string `json:"field"`
			Operator               string `json:"operator"`
			Value                  string `json:"value"`
			RewardType             string `json:"rewardType"`
			RewardQuantity         string `json:"rewardQuantity"`
			CriteriaOverwriteValue string `json:"criteriaOverwriteValue"`
		} `json:"criteria"`
		Metadata struct {
			Tweet struct {
				AuthorID            string    `json:"author_id"`
				CreatedAt           time.Time `json:"created_at"`
				EditHistoryTweetIds []string  `json:"edit_history_tweet_ids"`
				ID                  string    `json:"id"`
				PublicMetrics       struct {
					ImpressionCount int `json:"impression_count"`
					LikeCount       int `json:"like_count"`
					QuoteCount      int `json:"quote_count"`
					ReplyCount      int `json:"reply_count"`
					RetweetCount    int `json:"retweet_count"`
				} `json:"public_metrics"`
				Text string `json:"text"`
			} `json:"tweet"`
			User struct {
				CreatedAt       time.Time `json:"created_at"`
				ID              string    `json:"id"`
				Name            string    `json:"name"`
				ProfileImageURL string    `json:"profile_image_url"`
				PublicMetrics   struct {
					FollowersCount int `json:"followers_count"`
					FollowingCount int `json:"following_count"`
					ListedCount    int `json:"listed_count"`
					TweetCount     int `json:"tweet_count"`
				} `json:"public_metrics"`
				Username string `json:"username"`
				Verified bool   `json:"verified"`
			} `json:"user"`
		} `json:"metadata"`
	} `json:"tasks"`
}
