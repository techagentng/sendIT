package handlers

import (
	"net/http"

	"mailer/internal/email"
	"mailer/internal/postmark"

	"github.com/gin-gonic/gin"
)

type EmailHandler struct {
	service *email.Service
}

type SendEmailRequest struct {
	To       string `json:"to" binding:"required,email"`
	Subject  string `json:"subject" binding:"required"`
	HTMLBody string `json:"html_body" binding:"required"`
	TextBody string `json:"text_body"`
	Tag      string `json:"tag"`
}

type SendWelcomeRequest struct {
	To   string `json:"to" binding:"required,email"`
	Name string `json:"name" binding:"required"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewEmailHandler(service *email.Service) *EmailHandler {
	return &EmailHandler{service: service}
}

func (h *EmailHandler) SendEmail(c *gin.Context) {
	var req SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	emailMsg := postmark.Email{
		To:         req.To,
		Subject:    req.Subject,
		HTMLBody:   req.HTMLBody,
		TextBody:   req.TextBody,
		Tag:        req.Tag,
		TrackOpens: true,
	}

	if err := h.service.SendEmail(c.Request.Context(), emailMsg); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Email sent successfully",
	})
}

func (h *EmailHandler) SendWelcome(c *gin.Context) {
	var req SendWelcomeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := h.service.SendWelcomeEmail(c.Request.Context(), req.To, req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Welcome email sent successfully",
	})
}

func (h *EmailHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Email service is healthy",
	})
}
