package dds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// LogRecords is a nested struct in dds response
type LogRecords struct {
	Category           string `json:"Category" xml:"Category"`
	Id                 int    `json:"Id" xml:"Id"`
	HostAddress        string `json:"HostAddress" xml:"HostAddress"`
	CreateTime         string `json:"CreateTime" xml:"CreateTime"`
	QueryTimes         string `json:"QueryTimes" xml:"QueryTimes"`
	TableName          string `json:"TableName" xml:"TableName"`
	SQLText            string `json:"SQLText" xml:"SQLText"`
	ConnInfo           string `json:"ConnInfo" xml:"ConnInfo"`
	ExecutionStartTime string `json:"ExecutionStartTime" xml:"ExecutionStartTime"`
	ReturnRowCounts    int64  `json:"ReturnRowCounts" xml:"ReturnRowCounts"`
	Content            string `json:"Content" xml:"Content"`
	AccountName        string `json:"AccountName" xml:"AccountName"`
	DocsExamined       int64  `json:"DocsExamined" xml:"DocsExamined"`
	DBName             string `json:"DBName" xml:"DBName"`
	KeysExamined       int64  `json:"KeysExamined" xml:"KeysExamined"`
}
