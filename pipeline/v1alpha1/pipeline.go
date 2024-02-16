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

// Package v1alpha1 contains API Schema definitions for the pipeline v1alpha1 API group
// +kubebuilder:object:generate=true
package v1alpha1

import (
	"fmt"
)

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
	// ID trigger identifier
	ID string `json:"id"          bson:"id"          mapstructure:"id"`
	// Ref trigger reference
	Ref string `json:"ref"         bson:"ref"         mapstructure:"ref"`
	// Type trigger type
	Type string `json:"type"        bson:"type"        mapstructure:"type"`
	// ComponentId the component id that the trigger is bound with
	ComponentId string `json:"componentId" bson:"componentId" mapstructure:"componentId"`
	// Data the key/value for this trigger
	Data Data `json:"data"        bson:"data"        mapstructure:"data"`
}

type ConditionType string

const (
	TRIGGER ConditionType = "trigger"
	TASK    ConditionType = "task"
	STAGE   ConditionType = "stage"
)

type Condition struct {
	// Id condition identifier
	// +optional
	Id string `json:"id" bson:"id"`
	// Ref condition reference
	// +optional
	Ref string `json:"ref" bson:"ref"`
	// Type condition type
	// +optional
	Type ConditionType `json:"type" bson:"type"`
	// When is the operand when the trigger is true
	// +optional
	When string `json:"when" bson:"when"`
}

type Conditions [][]Condition

type Action struct {
	// ID the identifier of the action
	ID string `json:"id"           bson:"id"           mapstructure:"id"`
	// Name contains the action aname
	Name string `json:"name"         bson:"name"         mapstructure:"name"`
	// ComponentID the component id this action is linked with
	ComponentID string `json:"componentId"  bson:"componentId"  mapstructure:"componentId"`
	// Ref the action reference
	Ref string `json:"ref"          bson:"ref"          mapstructure:"ref"`
	// Data key/value
	// +optional
	Data map[string]string `json:"data"         bson:"data"         mapstructure:"data"`
	// Params key/value
	// +optional
	Params map[string]string `json:"params"       bson:"params"       mapstructure:"params"`
	// Environments key/value
	// +optional
	Environments map[string]string `json:"environments" bson:"environments" mapstructure:"environments"`
	// Status of the action during the execution
	// +optional
	Status string `json:"status"       bson:"status"       mapstructure:"status"`
	// StartTime date when the action has started during the execution
	// +optional
	StartTime int64 `json:"startTime"    bson:"startTime"    mapstructure:"startTime"`
	// EndTime date when the action has ended during the execution
	// +optional
	EndTime int64 `json:"endTime"      bson:"endTime"      mapstructure:"endTime"`
}

type Task struct {
	// ID task identifier
	ID string `json:"id"            bson:"id"            mapstructure:"id"`
	// Name task name
	Name string `json:"name"          bson:"name"          mapstructure:"name"`
	// SkipOnFailure whether to stop the pipeline on raising error
	SkipOnFailure bool `json:"skipOnFailure" bson:"skipOnFailure" mapstructure:"skipOnFailure"`
	// Conditions determines if this task need to be run
	// +optional
	Conditions Conditions `json:"conditions"    bson:"conditions"    mapstructure:"conditions"`
	// Actions list of action of this task
	Actions []Action `json:"actions"       bson:"actions"       mapstructure:"actions"`
	// StartTime timestamp when this task has started during the execution
	// +optional
	StartTime int64 `json:"startTime"     bson:"startTime"     mapstructure:"startTime"`
	// EndTime timestamp when this task has ended during the execution
	// +optional
	EndTime int64 `json:"endTime"       bson:"endTime"       mapstructure:"endTime"`
	// Status of this action during the execution
	// +optional
	Status string `json:"status"        bson:"status"        mapstructure:"status"`
}

type Stage struct {
	// ID the stage identifier
	ID string `json:"id"         bson:"id"         mapstructure:"id"`
	// Name
	Name string `json:"name"       bson:"name"       mapstructure:"name"`
	// Tasks
	Tasks []Task `json:"tasks"      bson:"tasks"      mapstructure:"tasks"`
	// Status of this stage during the execution
	// +optional
	Status string `json:"status"     bson:"status"     mapstructure:"status"`
	// StartTime timestamp when this stage has started during the execution
	// +optional
	StartTime int64 `json:"startTime"  bson:"startTime"  mapstructure:"startTime"`
	// EndTime timestamp when this stage has ended during the execution
	// +optional
	EndTime int64 `json:"endTime"    bson:"endTime"    mapstructure:"endTime"`
}

type ProjectID int64

func (in ProjectID) String() string {
	return fmt.Sprintf("%d", in)
}

// Pipeline defines the desired state of Pipeline
type Pipeline struct {
	// ID is the pipeline identifier
	ID string `json:"id"                             bson:"id"               mapstructure:"id"`
	// Type of the pipeline
	Type string `json:"type"                         bson:"type"             mapstructure:"type"`
	// PipelineIDNumber is the number of this pipeline
	PipelineIDNumber string `json:"pipelineIdNumber" bson:"pipelineIdNumber" mapstructure:"pipelineIdNumber"`
	// ProjectID is the project identifier for this pipeline
	ProjectID ProjectID `json:"projectId"            bson:"projectId"        mapstructure:"projectId"`
	// Name is the pipeline name
	// +optional
	Name string `json:"name"                         bson:"name"             mapstructure:"name"`
	// Triggers is when this pipeline will be triggered
	Triggers []Trigger `json:"triggers"              bson:"triggers"         mapstructure:"triggers"`
	// Stages is the stages within thi pipeline
	Stages []Stage `json:"stages"                    bson:"stages"           mapstructure:"stages"`
	// Status of the pipeline during the execution
	// +optional
	Status string `json:"status"                     bson:"status"           mapstructure:"status"`
	// StartTime timestamp when this pipeline has started during the execution
	// +optional
	StartTime int64 `json:"startTime"                bson:"startTime"        mapstructure:"startTime"`
	// EndTime timestamp when this pipeline has ended during the execution
	// +optional
	EndTime int64 `json:"endTime"                    bson:"endTime"          mapstructure:"endTime"`
	// LogUri url where to get the log of the pipeline execution
	// +optional
	LogUri string `json:"logUri"                     bson:"logUri"           mapstructure:"logUri"`
	// Complete a boolean to know if the pipeline configuration is complete
	// +optional
	Complete bool `json:"complete"                   bson:"complete"         mapstructure:"complete"`
	// Force a boolean to force pipeline recording
	// +optional
	Force bool `json:"force"                         bson:"-"                mapstructure:"-"`
	// FieldError list of error encountered into the pipeline configuration
	// +optional
	FieldError *FieldError `json:"fieldError"        bson:"fieldError"       mapstructure:"-"`
	// Artifacts a boolean set to true if the pipeline has generated artifacts during the execution
	// +optional
	Artifacts bool `json:"artifacts"                 bson:"artifacts"        mapstructure:"artifacts"`
	// TriggerId is the trigger id that trigger this pipeline
	// +optional
	TriggerId string `json:"triggerId,omitempty"     bson:"triggerId"        mapstructure:"triggerId"`
	// Commit is the commit that trigger this pipeline
	// +optional
	Commit Commit `json:"commit"                     bson:"commit"           mapstructure:"commit"`
	// EventID is the event id that trigger this pipeline
	// +optional
	EventID string `json:"eventId"                   bson:"eventId"          mapstructure:"eventId"`
}

type Commit struct {
	// ID the commit identifier
	// +optional
	ID string `json:"id"      bson:"id"      mapstructure:"id"`
	// Ref the commit reference
	// +optional
	Ref string `json:"ref"     bson:"ref"     mapstructure:"ref"`
	// Message the commit message
	// +optional
	Message string `json:"message" bson:"message" mapstructure:"message"`
}

type Field struct {
	// Message holds the main diagnostic message carried by this Field
	// +optional
	Message string `json:"message" bson:"message" mapstructure:"message"`

	// ID of the resource where the error is
	// +optional
	ID string `json:"id" bson:"id" mapstructure:"id"`

	// Path of the resource
	// +optional
	Path string `json:"path" bson:"path" mapstructure:"path"`
}

type FieldError struct {
	// Fields is a list of field where the issue is
	// +optional
	Fields []Field `json:"fields" bson:"fields" mapstructure:"fields"`
}
