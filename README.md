# swagger-filter

`swagger-filter` is a CLI that gets a swagger file as input and filters the endpoints of the API definition depending on the flags passed to the command, then outputs the filtered definition in a new file. Definitions non used in the selected endpoints are also removed from the definition.

## Usage

```sh
swagger-filter \
    --endpoint "/v1/pods" \
    --endpoint-prefix "/v1/" \
    --endpoint-regexp "/v1/.*" \
    input.json output.json
```

## k8s-apis.sh

`k8s-apis.sh` filters the Kubernetes API definition in `swagger.json` and outputs a swagger definition for each API Group in the `output` directory.
