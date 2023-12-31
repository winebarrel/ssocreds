# ssocreds

A tool to get temporary credentials from AWS SSO.

cf. https://aws.amazon.com/premiumsupport/knowledge-center/sso-temporary-credentials/

[![CI](https://github.com/winebarrel/ssocreds/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/ssocreds/actions/workflows/ci.yml)

## Installation

```
brew install winebarrel/ssocreds/ssocreds
```

## Usage

```
$ export AWS_PROFILE=my-profile
$ aws sso login

$ ssocreds
export AWS_ACCESS_KEY_ID='ASIA*************'
export AWS_SECRET_ACCESS_KEY='**********************************'
export AWS_SESSION_TOKEN='****************************************'

$ ssocreds json
{
  "accessKeyId": "ASIA*************",
  "secretAccessKey": "**********************************',
  "sessionToken": "****************************************"
}

$ eval $(ssocreds)
$ ... # Run any command that do not support AWS SSO
```
