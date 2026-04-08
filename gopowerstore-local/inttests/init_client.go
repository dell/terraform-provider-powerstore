/*
 *
 * Copyright © 2020-2024 Dell Inc. or its subsidiaries. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package inttests

import (
	"log"

	"github.com/dell/gopowerstore"
	"github.com/joho/godotenv"
)

const envVarsFile = "GOPOWERSTORE_TEST.env"

// C is global powerstore Client instance for testing
var C gopowerstore.Client

func initClient() {
	err := godotenv.Load(envVarsFile)
	if err != nil {
		log.Printf("%s file not found.", envVarsFile)
	}
	C, err = gopowerstore.NewClient()
	if err != nil {
		panic(err)
	}
}

func init() {
	initClient()
}

func GetNewClient() (client gopowerstore.Client) {
	err := godotenv.Load(envVarsFile)
	if err != nil {
		log.Printf("%s file not found.", envVarsFile)
	}
	client, err = gopowerstore.NewClient()
	if err != nil {
		panic(err)
	}

	return client
}
