# Go Action

A GitHub Action written in Go.

## Example

Below is an example workflow that uses this action:

```yaml
---
name: Go Action Example

on:
  workflow_dispatch:
    inputs:
      name:
        description: 'Name to greet'
        required: true
        default: 'World'
        type: string

jobs:
  greet-me:
    runs-on: ubuntu-latest
    steps:
      - name: Run Go Action
        uses: dustindortch/go-action@main
        with:
          name: 'World'
      - name: Print Greeting
        shell: bash
        run: |
          echo "${{ steps.greet-me.outputs.greeting }}"
...
```

## Inputs

### Required

`name` - The name to greet.

## Outputs

`greeting` - The greeting message.
