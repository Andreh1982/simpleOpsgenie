package models

import "time"

type PayloadListMirror struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		Actions         []interface{} `json:"actions"`
		CreatedAt       time.Time     `json:"createdAt"`
		Description     string        `json:"description"`
		ExtraProperties struct {
		} `json:"extraProperties"`
		ID               string        `json:"id"`
		ImpactStartDate  time.Time     `json:"impactStartDate"`
		ImpactedServices []interface{} `json:"impactedServices"`
		Links            struct {
			API string `json:"api"`
			Web string `json:"web"`
		} `json:"links"`
		Message    string `json:"message"`
		OwnerTeam  string `json:"ownerTeam"`
		Priority   string `json:"priority"`
		Responders []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"responders"`
		Status    string        `json:"status"`
		Tags      []interface{} `json:"tags"`
		TinyID    string        `json:"tinyId"`
		UpdatedAt time.Time     `json:"updatedAt"`
	} `json:"data"`
	Paging struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"paging"`
	Took      float64 `json:"took"`
	RequestID string  `json:"requestId"`
}

type PayloadUnitMirror struct {
	Data struct {
		ID               string    `json:"id"`
		Description      string    `json:"description"`
		ImpactedServices []string  `json:"impactedServices"`
		TinyID           string    `json:"tinyId"`
		Message          string    `json:"message"`
		Status           string    `json:"status"`
		Tags             []string  `json:"tags"`
		CreatedAt        time.Time `json:"createdAt"`
		UpdatedAt        time.Time `json:"updatedAt"`
		Priority         string    `json:"priority"`
		OwnerTeam        string    `json:"ownerTeam"`
		Responders       []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"responders"`
		ExtraProperties struct {
		} `json:"extraProperties"`
		Links struct {
			Web string `json:"web"`
			API string `json:"api"`
		} `json:"links"`
		ImpactStartDate time.Time     `json:"impactStartDate"`
		Actions         []interface{} `json:"actions"`
	} `json:"data"`
	Took      float64 `json:"took"`
	RequestID string  `json:"requestId"`
}

type CreateIncident struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Responders  []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"responders"`
	Tags    []string `json:"tags"`
	Details struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
	} `json:"details"`
	Priority         string   `json:"priority"`
	ImpactedServices []string `json:"impactedServices"`
	StatusPageEntry  struct {
		Title  string `json:"title"`
		Detail string `json:"detail"`
	} `json:"statusPageEntry"`
	NotifyStakeholders bool `json:"notifyStakeholders"`
}
