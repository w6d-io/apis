/*
Copyright 2022 WILDCARD.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import "encoding/json"

type RoleItems []*RoleItem

func (x *Trait) ToMap() (map[string]interface{}, error) {
	var m map[string]interface{}

	data, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &m)
	return m, err
}

// Add adds the new role or update the value if the key role already exist
func (in *RoleItems) Add(key int64, value string) {
	m := make(map[int64]string, len(*in)+1)

	for _, v := range *in {
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
	*in = ri
}

// Delete adds the new role or update the value if the key role already exist
func (in *RoleItems) Delete(key int64) {
	var ri []*RoleItem
	for _, v := range in {
		if v.Key != key {
			ri = append(ri, v)
		}
	}
	*in = ri
}
