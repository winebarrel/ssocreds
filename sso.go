package ssocreds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sso"
)

func getRoleCredentials(cfg aws.Config, account string, permissionSet string, accessToken string, region string) (*sso.GetRoleCredentialsOutput, error) {
	client := sso.NewFromConfig(cfg, func(o *sso.Options) {
		o.Region = region
	})

	input := &sso.GetRoleCredentialsInput{
		AccountId:   aws.String(account),
		RoleName:    aws.String(permissionSet),
		AccessToken: aws.String(accessToken),
	}

	return client.GetRoleCredentials(context.Background(), input)
}

func SsoCredentials(cfg aws.Config, account string, permissionSet string, accessToken string, region string) (string, string, string, error) {
	output, err := getRoleCredentials(cfg, account, permissionSet, accessToken, region)

	if err != nil {
		return "", "", "", err
	}

	return *output.RoleCredentials.AccessKeyId, *output.RoleCredentials.SecretAccessKey, *output.RoleCredentials.SessionToken, nil
}
