# git-hooks
A GO Lang module to perform actions on git hooks. Say linting or code formating on pre-commit among other possibilities.

## Installation

```
go get github.com/varmakarthik12/git-hooks
```

## Configuration
Root should have a `.git-hooks.yaml` file like below.

``` yaml
# Any supported hooks
pre-commit:
  #optional glob to execute on specific file paths.
  files:
    include:
      - "*.go"
      - "*.md"
    exclude:
      - "*.js"

  # Execute an action / code. 
  run:
    - echo "Single line command "
    - |
      echo "Muti line command"
      gofmt -l -w
```

## Currently support hooks
  * `pre-commit` - Supports files config
