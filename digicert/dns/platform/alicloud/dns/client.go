// Copyright (c) HashiCorp, Inc.

package dns

import (
	"fmt"

	alidns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
)

func NewClient(accessKey, secretKey string) (dns *Alidns, err error) {
	if accessKey == "" {
		return nil, fmt.Errorf("dns.newClient(): missing access_key")
	}
	if secretKey == "" {
		return nil, fmt.Errorf("dns.newClient(): missing secret_key")
	}

	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKey),
		AccessKeySecret: tea.String(secretKey),
	}

	config.Endpoint = tea.String("alidns.cn-hongkong.aliyuncs.com")
	client, err := alidns.NewClient(config)
	return &Alidns{
		Client: client,
	}, err
}
