/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 24/10/2022
*/

package v1alpha1

import (
    "strings"

    "github.com/w6d-io/x/errorx"
)

//func (in Conditions) Has(triggersID []string) []string {
//	var errorMessages []string
//	for _, cs := range in {
//		for _, c := range cs {
//			if !toolx.InArray(c.Ref, triggersID) {
//				errorMessages = append(errorMessages, c.Ref)
//			}
//		}
//	}
//	return errorMessages
//}

func (in *Pipeline) ConditionValidation() error {

    var errorMessages []string
    var triggersID []string

    for _, t := range in.Triggers {
        triggersID = append(triggersID, t.ID)
    }

    for _, s := range in.Stages {
        var stagesTasksID []string
        stagesTasksID = append(stagesTasksID, s.ID)
        for _, t := range s.Tasks {
            stagesTasksID = append(stagesTasksID, t.ID)
            //errorMessages = append(errorMessages, t.Conditions.Has(triggersID)...)
        }
        triggersID = append(triggersID, stagesTasksID...)
    }
    if len(errorMessages) != 0 {
        return errorx.New(nil, strings.Join(errorMessages, ", "))
    }
    return nil
}
