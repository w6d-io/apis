/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 14/11/2022
*/

// Package v1alpha1 contains API Schema definitions for notification
// +kubebuilder:object:generate=true
package v1alpha1

type Notification struct {
	Id      string   `json:"id"`
	Type    string   `json:"type"`
	Kind    string   `json:"kind"`
	Scope   []string `json:"scope"`
	Message string   `json:"message"`
	Time    int64    `json:"time"`
}
