package model

import "time"

type UserStatisticModel struct {
	TwitterQuestStatistic []struct {
		Tags          string `json:"tags"`
		TotalViews    int    `json:"totalViews"`
		TotalLikes    int    `json:"totalLikes"`
		TotalRetweets int    `json:"totalRetweets"`
		TotalReplies  int    `json:"totalReplies"`
		TotalQuotes   int    `json:"totalQuotes"`
	} `json:"twitterQuestStatistic"`
	QuestTaskStatistic []struct {
		QuestTaskID     string `json:"questTaskId"`
		Category        string `json:"category"`
		Code            string `json:"code"`
		IsFinished      bool   `json:"isFinished"`
		TotalPoints     int    `json:"totalPoints"`
		TotalMultiplier int    `json:"totalMultiplier"`
	} `json:"questTaskStatistic"`
	TotalPoints     int       `json:"totalPoints"`
	TotalMultiplier int       `json:"totalMultiplier"`
	Level           int       `json:"level"`
	LastUpdated     time.Time `json:"lastUpdated"`
	Logs            []struct {
		Added      int    `json:"added"`
		CriteriaID string `json:"criteria_id"`
		ID         string `json:"id"`
		Method     string `json:"method"`
		PolicyCode string `json:"policy_code"`
		Priority   int    `json:"priority"`
		TaskID     string `json:"task_id"`
		Value      int    `json:"value"`
		ValueType  string `json:"value_type"`
	} `json:"logs"`
	OgID int `json:"ogId"`
}
