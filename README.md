# ssocreds

A tool to get temporary credentials from AWS SSO.

cf. https://aws.amazon.com/premiumsupport/knowledge-center/sso-temporary-credentials/

[![build](https://github.com/winebarrel/ssocreds/actions/workflows/build.yml/badge.svg)](https://github.com/winebarrel/ssocreds/actions/workflows/build.yml)

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
```
