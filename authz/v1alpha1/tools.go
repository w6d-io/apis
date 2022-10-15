/*
Copyright 2020 WILDCARD SA.

Licensed under the WILDCARD SA License, Version 1.0 (the "License");
WILDCARD SA is register in french corporation.
You may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.w6d.io/licenses/LICENSE-1.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is prohibited.
Created on 25/07/2022
*/

package v1alpha1

import "encoding/json"

func (x *Trait) ToMap() (map[string]interface{}, error) {
    var m map[string]interface{}

    data, err := json.Marshal(x)
    if err != nil {
        return nil, err
    }
    err = json.Unmarshal(data, &m)
    return m, err
}

// SetRole adds the new role or update the value if the key role already exist
func SetRole(roles []*RoleItem, key int64, value string) []*RoleItem {
    m := make(map[int64]string, len(roles)+1)

    for _, v := range roles {
        m[v.Key] = v.Value
    }
    m[key] = value

    var ri []*RoleItem
    for k, v := range m {
        ri = append(ri, &RoleItem{
            Key:   k,
            Value: v,
        })
    }
    return ri
}

// DelRole adds the new role or update the value if the key role already exist
func DelRole(roles []*RoleItem, key int64) []*RoleItem {
    var ri []*RoleItem
    for _, v := range roles {
        if v.Key != key {
            ri = append(ri, v)
        }
    }
    return ri
}
