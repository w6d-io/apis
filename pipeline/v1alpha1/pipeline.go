/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 15/12/2021
*/

package v1alpha1

import "fmt"

type Status int

const (
    Pending Status = iota
    Running
    Cancelled
    Skipped
    Timeout
)

// PipelineStatus string list of status for pipeline
var PipelineStatus = []string{
    "Pending",
    "Running",
    "Cancelled",
    "Skipped",
    "Timeout",
}

func (s Status) ToString() string {
    return PipelineStatus[s]
}

type Data map[string]string

type Trigger struct {
    ID          string `json:"id"          bson:"id"          mapstructure:"id"`
    Ref         string `json:"ref"         bson:"ref"         mapstructure:"ref"`
    Type        string `json:"type"        bson:"type"        mapstructure:"type"`
    ComponentId string `json:"componentId" bson:"componentId" mapstructure:"componentId"`
    Data        Data   `json:"data"        bson:"data"        mapstructure:"data"`
}

type ConditionType string

const (
    TRIGGER ConditionType = "trigger"
    TASK    ConditionType = "task"
    STAGE   ConditionType = "stage"
)

type Condition struct {
    Id   string        `json:"id" bson:"id"`
    Ref  string        `json:"ref" bson:"ref"`
    Type ConditionType `json:"type" bson:"type"`
    When string        `json:"when" bson:"when"`
}

type Conditions [][]Condition

type Action struct {
    ID           string            `json:"id"           bson:"id"           mapstructure:"id"`
    Name         string            `json:"name"         bson:"name"         mapstructure:"name"`
    ComponentID  string            `json:"componentId"  bson:"componentId"  mapstructure:"componentId"`
    Ref          string            `json:"ref"          bson:"ref"          mapstructure:"ref"`
    Data         map[string]string `json:"data"         bson:"data"         mapstructure:"data"`
    Params       map[string]string `json:"params"       bson:"params"       mapstructure:"params"`
    Environments map[string]string `json:"environments" bson:"environments" mapstructure:"environments"`
    Status       string            `json:"status"       bson:"status"       mapstructure:"status"`
    StartTime    int64             `json:"startTime"    bson:"startTime"    mapstructure:"startTime"`
    EndTime      int64             `json:"endTime"      bson:"endTime"      mapstructure:"endTime"`
}

type Task struct {
    ID            string     `json:"id"            bson:"id"            mapstructure:"id"`
    Name          string     `json:"name"          bson:"name"          mapstructure:"name"`
    SkipOnFailure bool       `json:"skipOnFailure" bson:"skipOnFailure" mapstructure:"skipOnFailure"`
    Conditions    Conditions `json:"conditions"    bson:"conditions"    mapstructure:"conditions"`
    Actions       []Action   `json:"actions"       bson:"actions"       mapstructure:"actions"`
    StartTime     int64      `json:"startTime"     bson:"startTime"     mapstructure:"startTime"`
    EndTime       int64      `json:"endTime"       bson:"endTime"       mapstructure:"endTime"`
    Status        string     `json:"status"        bson:"status"        mapstructure:"status"`
}

type Stage struct {
    ID        string `json:"id"         bson:"id"         mapstructure:"id"`
    Name      string `json:"name"       bson:"name"       mapstructure:"name"`
    Tasks     []Task `json:"tasks"      bson:"tasks"      mapstructure:"tasks"`
    Status    string `json:"status"     bson:"status"     mapstructure:"status"`
    EndTime   int64  `json:"endTime"    bson:"endTime"    mapstructure:"endTime"`
    StartTime int64  `json:"startTime"  bson:"startTime"  mapstructure:"startTime"`
}

type ProjectID int64

func (in ProjectID) String() string {
    return fmt.Sprintf("%d", in)
}

type Pipeline struct {
    ID               string    `json:"id"                  bson:"id"                mapstructure:"id"`
    Type             string    `json:"type"                bson:"type"              mapstructure:"type"`
    PipelineIDNumber string    `json:"pipelineIdNumber"    bson:"pipelineIdNumber"  mapstructure:"pipelineIdNumber"`
    ProjectID        ProjectID `json:"projectId"           bson:"projectId"         mapstructure:"projectId"`
    Name             string    `json:"name"                bson:"name"              mapstructure:"name"`
    Triggers         []Trigger `json:"triggers"            bson:"triggers"          mapstructure:"triggers"`
    Stages           []Stage   `json:"stages"              bson:"stages"            mapstructure:"stages"`
    Status           string    `json:"status"              bson:"status"            mapstructure:"status"`
    StartTime        int64     `json:"startTime"           bson:"startTime"         mapstructure:"startTime"`
    EndTime          int64     `json:"endTime"             bson:"endTime"           mapstructure:"endTime"`
    LogUri           string    `json:"logUri"              bson:"logUri"            mapstructure:"logUri"`
    Complete         bool      `json:"complete"            bson:"complete"          mapstructure:"complete"`
    Force            bool      `json:"force"               bson:"-"                 mapstructure:"-"`
    Artifacts        bool      `json:"artifacts"           bson:"artifacts"         mapstructure:"artifacts"`
    TriggerId        string    `json:"triggerId,omitempty" bson:"triggerId"         mapstructure:"triggerId"`
    Commit           Commit    `json:"commit"              bson:"commit"            mapstructure:"commit"`
    EventID          string    `json:"eventId"             bson:"eventId"           mapstructure:"eventId"`
}

type Commit struct {
    ID      string `json:"id"      bson:"id"      mapstructure:"id"`
    Ref     string `json:"ref"     bson:"ref"     mapstructure:"ref"`
    Message string `json:"message" bson:"message" mapstructure:"message"`
}
