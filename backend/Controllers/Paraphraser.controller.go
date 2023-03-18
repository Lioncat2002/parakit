package controllers

import (
	"context"
	"fmt"
	models "main/Models"
	"net/http"

	gogpt "github.com/sashabaranov/go-openai"

	"github.com/gin-gonic/gin"
)

func GetSummary(c *gin.Context, client *gogpt.Client) {
	ctx := context.Background()

	var paraphraser models.Paraphrase
	c.BindJSON(&paraphraser)
	fmt.Println(paraphraser.Text)
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   100,
		Prompt:      "Paraphrase the following in less than 100 words: \"" + paraphraser.Text + "\"",
		Temperature: 1,
	}
	resp, err := client.CreateCompletion(ctx, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get fact",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"paraphrased": resp.Choices[0].Text,
	})

}
