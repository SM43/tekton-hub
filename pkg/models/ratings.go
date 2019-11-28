package models

import (
	"log"
	"strconv"
)

// Rating represents Rating model in database
type Rating struct {
	ID        int `json:"id"`
	TaskID    int `json:"task_id"`
	OneStar   int `json:"one"`
	TwoStar   int `json:"two"`
	ThreeStar int `json:"three"`
	FourStar  int `json:"four"`
	FiveStar  int `json:"five"`
}

// GetRatingDetialsByTaskID retrieves rating details of a task
func GetRatingDetialsByTaskID(id string) Rating {
	taskID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	sqlStatement := `SELECT * FROM RATING WHERE ID=$1`
	taskRating := Rating{}
	err = DB.QueryRow(sqlStatement, taskID).Scan(&taskRating)
	if err != nil {
		log.Println(err)
	}
	return taskRating
}

func calculateAverageRating(taskID int) float64 {
	rating := Rating{}
	sqlStatement := `SELECT * FROM RATING WHERE TASK_ID=$1`
	err := DB.QueryRow(sqlStatement, taskID).Scan(&rating.ID, &rating.TaskID, &rating.OneStar, &rating.TwoStar, &rating.ThreeStar, &rating.FourStar, &rating.FiveStar)
	if err != nil {
		log.Println(err)
	}
	totalResponses := float64(rating.OneStar + rating.TwoStar + rating.ThreeStar + rating.FourStar + rating.FiveStar)
	averageRating := float64(rating.OneStar+rating.TwoStar*2+rating.ThreeStar*3+rating.FourStar*4+rating.FiveStar*5) / (totalResponses)
	return averageRating
}

func addStars(taskID int, stars int, prevStars int) {
	var sqlStatement string
	switch stars {
	case 1:
		sqlStatement = `INSERT INTO RATING(ONESTAR,TASK_ID) VALUES($1,$2) ON CONFLICT (TASK_ID) DO UPDATE SET ONESTAR=RATING.ONESTAR+1`
	case 2:
		sqlStatement = `INSERT INTO RATING(TWOSTAR,TASK_ID) VALUES($1,$2) ON CONFLICT (TASK_ID) DO UPDATE SET TWOSTAR=RATING.TWOSTAR+1`
	case 3:
		sqlStatement = `INSERT INTO RATING(THREESTAR,TASK_ID) VALUES($1,$2) ON CONFLICT (TASK_ID) DO UPDATE SET THREESTAR=RATING.THREESTAR+1`
	case 4:
		sqlStatement = `INSERT INTO RATING(FOURSTAR,TASK_ID) VALUES($1,$2) ON CONFLICT (TASK_ID) DO UPDATE SET FOURSTAR=RATING.FOURSTAR+1`
	case 5:
		sqlStatement = `INSERT INTO RATING(FIVESTAR,TASK_ID) VALUES($1,$2) ON CONFLICT (TASK_ID) DO UPDATE SET FIVESTAR=RATING.FIVESTAR+1`
	}
	_, err := DB.Exec(sqlStatement, 1, taskID)
	if err != nil {
		log.Println(err)
	}
}

func updateStars(taskID int, stars int, prevStars int) {
	var sqlStatement string
	switch stars {
	case 1:
		sqlStatement = `UPDATE RATING SET ONESTAR=$2 WHERE TASK_ID=$1`
	case 2:
		sqlStatement = `UPDATE RATING SET TWOSTAR=$2 WHERE TASK_ID=$1`
	case 3:
		sqlStatement = `UPDATE RATING SET THREESTAR=$2 WHERE TASK_ID=$1`
	case 4:
		sqlStatement = `UPDATE RATING SET FOURSTAR=$2 WHERE TASK_ID=$1`
	case 5:
		sqlStatement = `UPDATE RATING SET FIVESTAR=$2 WHERE TASK_ID=$1`
	}
	_, err := DB.Exec(sqlStatement, taskID, stars)
	if err != nil {
		log.Println(err)
	}
}

func deleteOldStars(taskID int, prevStars int) {
	var sqlStatement string
	switch prevStars {
	case 1:
		sqlStatement = `UPDATE RATING SET ONESTAR=ONSESTAR-1 WHERE TASK_ID=$1`
	case 2:
		sqlStatement = `UPDATE RATING SET TWOSTAR=TWOSTAR-1 WHERE TASK_ID=$1`
	case 3:
		sqlStatement = `UPDATE RATING SET THREESTAR=THREESTAR-1 WHERE TASK_ID=$1`
	case 4:
		sqlStatement = `UPDATE RATING SET FOURSTAR=FOURSTAR-1 WHERE TASK_ID=$1`
	case 5:
		sqlStatement = `UPDATE RATING SET FIVESTAR=FIVESTAR-1 WHERE TASK_ID=$1`
	}
	_, err := DB.Exec(sqlStatement, taskID, 1)
	if err != nil {
		log.Println(err)
	}
}
