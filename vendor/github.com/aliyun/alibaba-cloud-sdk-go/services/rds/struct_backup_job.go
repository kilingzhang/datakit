package rds

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

// BackupJob is a nested struct in rds response
type BackupJob struct {
	BackupProgressStatus string `json:"BackupProgressStatus" xml:"BackupProgressStatus"`
	BackupStatus         string `json:"BackupStatus" xml:"BackupStatus"`
	JobMode              string `json:"JobMode" xml:"JobMode"`
	Process              string `json:"Process" xml:"Process"`
	TaskAction           string `json:"TaskAction" xml:"TaskAction"`
	BackupJobId          string `json:"BackupJobId" xml:"BackupJobId"`
	BackupId             string `json:"BackupId" xml:"BackupId"`
}
