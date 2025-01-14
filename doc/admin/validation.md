# Sourcegraph Instance Validation

>NOTE: **Sourcegraph Instance Validation is currently experimental.** We're exploring this feature set. 
>Let us know what you think! [File an issue](https://github.com/sourcegraph/sourcegraph/issues/new/choose)
>with feedback/problems/questions, or [contact us directly](https://about.sourcegraph.com/contact).

Instance validation provides a quick way to check that a Sourcegraph instance functions properly after a fresh install
 or an update.

The [`src` CLI](https://github.com/sourcegraph/src-cli) has an experimental command `validate` which drives the
 validation from a user-provided configuration file with a validation specification (in JSON or YAML format). if no validation specification file is provided it will execute the following: 
 
* temporarily add an external service
* wait for a repository to be cloned
* perform a search on the cloned repo
* perform a search on a non-indexed branch of the cloned repo
* creates basic code insight
* remove the added external service

### Validation specification
 
Validation specifications can be provided in either a YAML or JSON format. The best way to describe this initial, simple, and experimental validation specification is with the example below:

#### YAML File Specification

```yaml
# creates the first admin user on a fresh install (skips creation if user exists)
firstAdmin:
    email: foo@example.com
    username: foo
    password: "{{ .admin_password }}"

# adds the specified code host
externalService:
  config:
    url: https://github.com
    token: "{{ .github_token }}"
    orgs: []
    repos:
      - sourcegraph-testing/zap
  kind: GITHUB
  displayName: footest
  # set to true if this code host config should be deleted at the end of validation
  deleteWhenDone: true

# checks maxTries if specified repo is cloned and waits sleepBetweenTriesSeconds between checks 
waitRepoCloned:
  repo: github.com/sourcegraph-testing/zap
  maxTries: 5
  sleepBetweenTriesSeconds: 2

# performs the specified search and checks that at least one result is returned
searchQuery: 
  - repo:^github.com/sourcegraph-testing/zap$ test
  - repo:^github.com/sourcegraph-testing/zap$@v1.14.1 test

# checks to see if instance can create code insights
createInsight:
  title: "Javascript to Typescript migration" 
  dataSeries:
    [ {
        "query": "lang:javascript",
        "label": "javascript",
        "repositoryScope": [], # e.g. ["github.com/sourcegraph-testing/zap"]
        "lineColor": "#6495ED",
        "timeScopeUnit": "MONTH",
        "timeScopeValue": 1
      },
     {
        "query": "lang:typescript",
        "label": "typescript",
        "lineColor": "#DE3163",
        "repositoryScope": [
            "github.com/sourcegraph/sourcegraph"
          ],
        "timeScopeUnit": "MONTH",
        "timeScopeValue": 1
      }
    ]
```
#### JSON File Specification

```json
{   
   "firstAdmin": {
    "email": "foo@example.com",
    "username": "foo",
    "password": "{{ .admin_password }}"
    },
    "externalService": {
        "config": {
            "url": "https://github.com",
            "token": "{{ .github_token }}",
            "orgs": [],
            "repos": [
                "sourcegraph-testing/zap"
            ]
        },
        "kind": "GITHUB",
        "displayName": "footest",
        "deleteWhenDone": true
    },
    "waitRepoCloned": {
        "repo": "github.com/sourcegraph-testing/zap",
        "maxTries": 5,
        "sleepBetweenTriesSeconds": 5
    },
    "searchQuery": [
        "repo:^github.com/sourcegraph-testing/zap$ test",
        "repo:^github.com/sourcegraph-testing/zap$@v1.14.1 test"
    ],
   
    "createInsight": {
        "title": "Javascript to Typescript migration", 
        "dataSeries" : [ {
            "query": "lang:javascript",
            "label": "javascript",
            "repositoryScope": [],
            "lineColor": "#6495ED",
            "timeScopeUnit": "MONTH",
            "timeScopeValue": 1
          },
          {
            "query": "lang:typescript",
            "label": "typescript",
            "lineColor": "#DE3163",
            "repositoryScope": [],
            "timeScopeUnit": "MONTH",
            "timeScopeValue": 1
          }
       ]
}
```

With this configuration, the validation command executes the following steps: 

* create the first admin user
* add an external service
* wait for a repository to be cloned
* perform a search
* creates a code insight ( will need to be manually removed )
 
>NOTE: Every step is optional (if the corresponding top-level key is not present then the step is skipped).
> 
### Passing in secrets

It is often the case that the config file with the validation specification needs to declare passwords, tokens, or other
secrets and these secrets should not be exposed or committed to a git repo.

The validation specification can refer to string values that come from a context specified outside the config file
(see the `Usage` section below). References to string values from this outside context are specified like so:
`{{ .some_key }}`. The context will have a string value defined under the key `some_key` and the validation execution will
use that.

### Usage

Use the [`src` CLI](https://github.com/sourcegraph/src-cli) to validate with a validation specification file:
```shell script
src validate -context github_token=$GITHUB_TOKEN validate.yaml
```
```shell script
src validate -context github_token=$GITHUB_TOKEN validate.json
```
To execute default validation checks:

```shell script
src validate -context github_token=$GITHUB_TOKEN
```

The `src` binary finds the Sourcegraph instance to validate from the environment variables 
[`SRC_ENDPOINT` and `SRC_ACCESS_TOKEN`](https://github.com/sourcegraph/src-cli#setup-with-your-sourcegraph-instance). 

> Note: The `SRC_ACCESS_TOKEN` is not needed when a first admin user is declared in the validation specification.


