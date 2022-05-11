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

// RouteMap is a nested struct in cbn response
type RouteMap struct {
	RouteMapId                         string                        `json:"RouteMapId" xml:"RouteMapId"`
	Status                             string                        `json:"Status" xml:"Status"`
	TransmitDirection                  string                        `json:"TransmitDirection" xml:"TransmitDirection"`
	SourceInstanceIdsReverseMatch      bool                          `json:"SourceInstanceIdsReverseMatch" xml:"SourceInstanceIdsReverseMatch"`
	CenRegionId                        string                        `json:"CenRegionId" xml:"CenRegionId"`
	CenId                              string                        `json:"CenId" xml:"CenId"`
	Priority                           int                           `json:"Priority" xml:"Priority"`
	TransitRouterRouteTableId          string                        `json:"TransitRouterRouteTableId" xml:"TransitRouterRouteTableId"`
	CommunityOperateMode               string                        `json:"CommunityOperateMode" xml:"CommunityOperateMode"`
	MapResult                          string                        `json:"MapResult" xml:"MapResult"`
	CommunityMatchMode                 string                        `json:"CommunityMatchMode" xml:"CommunityMatchMode"`
	Description                        string                        `json:"Description" xml:"Description"`
	AsPathMatchMode                    string                        `json:"AsPathMatchMode" xml:"AsPathMatchMode"`
	Preference                         int                           `json:"Preference" xml:"Preference"`
	DestinationInstanceIdsReverseMatch bool                          `json:"DestinationInstanceIdsReverseMatch" xml:"DestinationInstanceIdsReverseMatch"`
	CidrMatchMode                      string                        `json:"CidrMatchMode" xml:"CidrMatchMode"`
	NextPriority                       int                           `json:"NextPriority" xml:"NextPriority"`
	MatchAddressType                   string                        `json:"MatchAddressType" xml:"MatchAddressType"`
	GatewayZoneId                      string                        `json:"GatewayZoneId" xml:"GatewayZoneId"`
	SourceRegionIds                    SourceRegionIds               `json:"SourceRegionIds" xml:"SourceRegionIds"`
	SourceChildInstanceTypes           SourceChildInstanceTypes      `json:"SourceChildInstanceTypes" xml:"SourceChildInstanceTypes"`
	DestinationRouteTableIds           DestinationRouteTableIds      `json:"DestinationRouteTableIds" xml:"DestinationRouteTableIds"`
	SourceInstanceIds                  SourceInstanceIds             `json:"SourceInstanceIds" xml:"SourceInstanceIds"`
	DestinationCidrBlocks              DestinationCidrBlocks         `json:"DestinationCidrBlocks" xml:"DestinationCidrBlocks"`
	DestinationRegionIds               DestinationRegionIds          `json:"DestinationRegionIds" xml:"DestinationRegionIds"`
	SourceRouteTableIds                SourceRouteTableIds           `json:"SourceRouteTableIds" xml:"SourceRouteTableIds"`
	MatchCommunitySet                  MatchCommunitySet             `json:"MatchCommunitySet" xml:"MatchCommunitySet"`
	PrependAsPath                      PrependAsPath                 `json:"PrependAsPath" xml:"PrependAsPath"`
	RouteTypes                         RouteTypes                    `json:"RouteTypes" xml:"RouteTypes"`
	OriginalRouteTableIds              OriginalRouteTableIds         `json:"OriginalRouteTableIds" xml:"OriginalRouteTableIds"`
	DestinationChildInstanceTypes      DestinationChildInstanceTypes `json:"DestinationChildInstanceTypes" xml:"DestinationChildInstanceTypes"`
	DestinationInstanceIds             DestinationInstanceIds        `json:"DestinationInstanceIds" xml:"DestinationInstanceIds"`
	MatchAsns                          MatchAsns                     `json:"MatchAsns" xml:"MatchAsns"`
	OperateCommunitySet                OperateCommunitySet           `json:"OperateCommunitySet" xml:"OperateCommunitySet"`
	SrcZoneIds                         SrcZoneIds                    `json:"SrcZoneIds" xml:"SrcZoneIds"`
}
