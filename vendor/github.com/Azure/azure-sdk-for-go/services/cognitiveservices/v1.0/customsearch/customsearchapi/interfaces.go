package customsearchapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/customsearch"
)

// CustomInstanceClientAPI contains the set of methods on the CustomInstanceClient type.
type CustomInstanceClientAPI interface {
	Search(ctx context.Context, customConfig string, query string, acceptLanguage string, userAgent string, clientID string, clientIP string, location string, countryCode string, count *int32, market string, offset *int32, safeSearch customsearch.SafeSearch, setLang string, textDecorations *bool, textFormat customsearch.TextFormat) (result customsearch.SearchResponse, err error)
}

var _ CustomInstanceClientAPI = (*customsearch.CustomInstanceClient)(nil)
