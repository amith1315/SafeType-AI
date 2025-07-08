package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/genai"
)

var routes = map[string]string{
	"getshoes": "GET",
}

func getRoutes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, routes)
}

func sendDataToGemini(textData string) (interface{}, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(textData+" the username field is from the json data. decode meaningful words, names of people, all content as important_information, sensitive_information, nsfw_content and mental state(ranging from 1 to 5, 1 is good and 5 is worst) and current data and time in ist format. only return me the sql query to insert it into a database table content_analysis with fields (username, names, important_information sensitive_information, content, nsfw_content, mental_state, cur_data_time) here the username field refers to the username inside the input json data and nothing else just give the sql insert data query if null the make it empty string"),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	var data interface{}
	jsonBytes, err := result.MarshalJSON()
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}
	return data, err
	// return "", nil
}

func extractSQLQuery(response interface{}) (string, error) {
	responseMap, ok := response.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("response is not a map")
	}

	candidates, ok := responseMap["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		return "", fmt.Errorf("candidates field is missing or empty")
	}

	firstCandidate, ok := candidates[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("first candidate is not a map")
	}

	content, ok := firstCandidate["content"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("content field is not a map")
	}

	parts, ok := content["parts"].([]interface{})
	if !ok || len(parts) == 0 {
		return "", fmt.Errorf("parts field is missing or empty")
	}

	firstPart, ok := parts[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("first part is not a map")
	}

	sqlQuery, ok := firstPart["text"].(string)
	if !ok {
		return "", fmt.Errorf("text field is not a string")
	}

	return sqlQuery, nil
}

var DB *sql.DB

func logData(c *gin.Context) {
	var data interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
	}
	str := fmt.Sprintf("%v", data)

	response, err := sendDataToGemini(str)
	if err != nil {
		fmt.Println(response)
	}

	sqlQuery, err := extractSQLQuery(response)
	if err != nil {
		log.Println("Error extracting SQL query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract SQL query"})
		return
	}

	q := sqlQuery[6:]
	query := strings.Split(q, "`")[0]
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	_, err = DB.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()
	c.JSON(http.StatusOK, gin.H{})
}

type ContentAnalysis struct {
	ID                   int    `json:"id"`
	Username             string `json:"username"`
	Names                string `json:"names"`
	ImportantInformation string `json:"important_information"`
	SensitiveInformation string `json:"sensitive_information"`
	Content              string `json:"content"`
	NSFWContent          string `json:"nsfw_content"`
	MentalState          *int   `json:"mental_state"`
	DataTime             string `json:"cur_data_time"`
}

func getUserData(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	username := queryParams.Get("name")

	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println("Error opening database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT id, username, names, important_information, sensitive_information, content, nsfw_content, mental_state, cur_data_time FROM content_analysis where username =" + `"` + username + `"` + ";")
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer rows.Close()

	var results []ContentAnalysis

	for rows.Next() {
		var record ContentAnalysis
		err := rows.Scan(
			&record.ID,
			&record.Username,
			&record.Names,
			&record.ImportantInformation,
			&record.SensitiveInformation,
			&record.Content,
			&record.NSFWContent,
			&record.MentalState,
			&record.DataTime,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse data"})
			return
		}
		results = append(results, record)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process data"})
		return
	}
	fmt.Println(results)
	c.JSON(http.StatusOK, results)
}

func login(c *gin.Context) {
	type Cred struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var cred Cred
	if err := c.ShouldBindJSON(&cred); err != nil {
		log.Println(err)
	}
	query := "SELECT * FROM users  WHERE " + `"` + cred.Username + `", "` + cred.Password + `"` + ";"
	fmt.Println(query)
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	_, err = DB.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	defer DB.Close()
	c.JSON(http.StatusOK, gin.H{"status": true})
}

var URL = "http://127.0.0.1:8000/"

func main() {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	router.GET("/routes", getRoutes)
	router.POST("/logdata", logData)
	router.GET("/getdata", getUserData)
	router.POST("/login", login)

	router.Run(":8000")
}
