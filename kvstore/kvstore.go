// Copyright 2016 Cyako Author

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kvstore

import (
	// "fmt"
	cyako "github.com/Cyako/Cyako.go"

	"github.com/Centimitr/namespace"
)

type Interfacce interface {
	Get(string) interface{}
	Set(string, interface{})
	Has(string) bool
	Delete(string)
	Disactive()
	Active()
}

type KVStore struct {
	Interfacce
	Namespace namespace.Namespace
}

func GetScopedKeyString(scope, name string) string {
	return scope + "." + name
}

func (k *KVStore) GetWithScoped(scope, name string) interface{} {
	return k.Get(GetScopedKeyString(scope, name))
}

func (k *KVStore) SetWithScoped(scope, name string, value interface{}) {
	k.Set(GetScopedKeyString(scope, name), value)
}

func (k *KVStore) HasWithScoped(scope, name string) bool {
	return k.Has(GetScopedKeyString(scope, name))
}

func (k *KVStore) DeleteWithScoped(scope, name string) {
	k.Delete(GetScopedKeyString(scope, name))
}

func init() {
	kvstore := &KVStore{
		Namespace: Memory{
			m: make(map[string]interface{}),
		},
	}
	kvstore.Namespace.Init()
	cyako.LoadService(kvstore)
}
