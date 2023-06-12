// Copyright (c) 2023 GPBR Participacoes LTDA.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cloudfront

import (
	"fmt"
	"strings"

	awscloudfront "github.com/aws/aws-sdk-go/service/cloudfront"
)

type OAC struct {
	ID                            string `json:"id"`
	Name                          string `json:"name"`
	OriginName                    string `json:"originName"`
	OriginAccessControlOriginType string `json:"originAccessControlOriginType"`
	SigningBehavior               string `json:"signingBehavior"`
	SigningProtocol               string `json:"signingProtocol"`
}

func NewOAC(distribution, originName string) OAC {
	return OAC{
		Name:                          oacName(distribution, originName),
		OriginName:                    originName,
		OriginAccessControlOriginType: awscloudfront.OriginAccessControlOriginTypesS3,
		SigningBehavior:               awscloudfront.OriginAccessControlSigningBehaviorsAlways,
		SigningProtocol:               awscloudfront.OriginAccessControlSigningProtocolsSigv4,
	}
}

func oacName(distributionName, s3Host string) string {
	s3Name := strings.Split(s3Host, ".")[0]
	return fmt.Sprintf("%s-%s", distributionName, s3Name)
}
