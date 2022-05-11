package cbn

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

// TrafficMatchRule is a nested struct in cbn response
type TrafficMatchRule struct {
	MatchDscp                   int    `json:"MatchDscp" xml:"MatchDscp"`
	DstCidr                     string `json:"DstCidr" xml:"DstCidr"`
	TrafficMatchRuleDescription string `json:"TrafficMatchRuleDescription" xml:"TrafficMatchRuleDescription"`
	Protocol                    string `json:"Protocol" xml:"Protocol"`
	TrafficMatchRuleId          string `json:"TrafficMatchRuleId" xml:"TrafficMatchRuleId"`
	SrcCidr                     string `json:"SrcCidr" xml:"SrcCidr"`
	TrafficMatchRuleName        string `json:"TrafficMatchRuleName" xml:"TrafficMatchRuleName"`
	TrafficMatchRuleStatus      string `json:"TrafficMatchRuleStatus" xml:"TrafficMatchRuleStatus"`
	DstPortRange                []int  `json:"DstPortRange" xml:"DstPortRange"`
	SrcPortRange                []int  `json:"SrcPortRange" xml:"SrcPortRange"`
}
