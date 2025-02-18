package ai_dto

import "time"

type Provider struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Config           string `json:"config"`
	GetAPIKeyUrl     string `json:"get_apikey_url"`
	DefaultLLM       string `json:"-"`
	DefaultLLMConfig string `json:"-"`
}

type ProviderItem struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	DefaultLLM     string    `json:"default_llm"`
	DefaultLLMLogo string    `json:"default_llm_logo"`
	Logo           string    `json:"logo"`
	Configured     bool      `json:"configured"`
	UpdateTime     time.Time `json:"-"`
}

type SimpleProviderItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	Configured bool   `json:"configured"`
}

type LLMItem struct {
	Id     string   `json:"id"`
	Logo   string   `json:"logo"`
	Config string   `json:"config"`
	Scopes []string `json:"scopes"`
}
