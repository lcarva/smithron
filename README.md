# Smithron

The CI-forging tool.

Smithron forges "step executions" to various CI providers from a common plan. If you have a tool
packaged as a container image and want to provide native integrations with GitHub Actions, GitLab
CI, and Tekton Pipelines, Smithron is for you.

(This project is still experimental. Feedback is highly welcome!)

## Install

Currently, a binary is not yet provided to run Smithron. You must build it from source. The simplest
way to do so is by simply using the `build` (default) target of the [Makefile](./Makefile):

```bash
make
```

## Example Usage

The [examples](./examples/) directory provides a list of plans that can be forged into step
executions.

### GitLab

```bash
./bin/smithron forge --plan examples/simple-plan.yaml --target gitlab
```

```yaml
hello-world:
  image:
    entrypoint:
    - echo
    - Hello, World!
    name: docker.io/busybox:latest
  variables:
    GREETING: Hello
```

### GitHub

```bash
./bin/smithron forge --plan examples/simple-plan.yaml --target github
```

```yaml
name: hello-world
runs:
  using: docker
  image: docker.io/busybox:latest
  entrypoint: echo
  args:
  - Hello, World!
  env:
    GREETING: Hello
```

### Tekton

```bash
./bin/smithron forge --plan examples/simple-plan.yaml --target tekton
```

```yaml
apiVersion: tekton.dev/v1
kind: Task
metadata:
  creationTimestamp: null
  name: hello-world
spec:
  steps:
  - command:
    - echo
    - Hello, World!
    computeResources: {}
    env:
    - name: GREETING
      value: Hello
    image: docker.io/busybox:latest
    name: run
```
