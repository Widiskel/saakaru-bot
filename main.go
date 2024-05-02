package main

import (
	"fmt"
	utils "saakaru/utils"

	//for api call
	api "saakaru/api"
)

func main() {
	// Accessing the accounts array from accounts.go
	for i, token := range utils.Accounts {
		fmt.Printf("Account %d:\n", i+1)
		fmt.Printf("Token: %s\n", token)
		fmt.Println()
		fmt.Println()
		quest, err := api.GetQuest(token)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		user, err := api.GetUser(token)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		userStatistic, err := api.GetUserStatistic(token, quest.ID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		var taskList []string
		for _, task := range quest.Tasks {
			taskList = append(taskList, task.ID)
		}
		var completedTaskList []string
		for _, task := range userStatistic.QuestTaskStatistic {
			completedTaskList = append(completedTaskList, task.QuestTaskID)
		}

		var uncompletedTaskList []string

		for _, taskID := range taskList {
			found := false

			for _, completedTaskID := range completedTaskList {
				if taskID == completedTaskID {
					found = true
					break
				}
			}
			if !found {
				uncompletedTaskList = append(uncompletedTaskList, taskID)
			}
		}

		fmt.Println("===================================================")
		fmt.Println("================ SAAKARU QUEST BOT ================")
		fmt.Println("=================== BY WIDISKEL ===================")
		fmt.Println("===================================================")
		fmt.Println("=> ")
		fmt.Println("=> " + quest.Name)
		fmt.Println("=> Prize Pool :" + quest.TotalTokenPrize + " $SKR")
		fmt.Println("=> Total Quest", len(quest.Tasks))
		fmt.Println("=> ")
		fmt.Println("===================================================")
		fmt.Println("==================== USER INFO ====================")
		fmt.Println("===================================================")
		fmt.Println("=> ")
		fmt.Println("=> " + user.Name)
		fmt.Println("=> " + user.TwitterHandle)
		fmt.Println("=> Uncompleted task : ", len(uncompletedTaskList))
		fmt.Println("=> ")
		fmt.Println("===================================================")

		for _, taskId := range uncompletedTaskList {
			fmt.Println()
			fmt.Println("===================================================")
			fmt.Println("=> Claiming Task " + taskId)
			api.ClaimQuest(token, quest.ID, taskId)
			fmt.Println("===================================================")
			fmt.Println()
		}

	}
}
