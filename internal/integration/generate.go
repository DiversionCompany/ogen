package integration

// Sample API matrix:
//
//go:generate go run ../../cmd/ogen -v --debug.ignoreNotImplemented "enum format" --infer-types --clean --generate-tests        --target sample_api      ../../_testdata/positive/sample.json
//go:generate go run ../../cmd/ogen -v --debug.ignoreNotImplemented "enum format" --infer-types --clean --no-server             --target sample_api_ns   ../../_testdata/positive/sample.json
//go:generate go run ../../cmd/ogen -v --debug.ignoreNotImplemented "enum format" --infer-types --clean --no-client             --target sample_api_nc   ../../_testdata/positive/sample.json
//go:generate go run ../../cmd/ogen -v --debug.ignoreNotImplemented "enum format" --infer-types --clean --no-server --no-client --target sample_api_nsnc ../../_testdata/positive/sample.json
//go:generate go run ../../cmd/ogen -v --debug.ignoreNotImplemented "enum format" --infer-types --clean --skip-default-responses --target sample_api_sdr ../../_testdata/positive/sample.json

//go:generate go run ../../cmd/ogen -v --clean --convenient-errors --target sample_err ../../_testdata/positive/convenient_errors/errors.json
//go:generate go run ../../cmd/ogen -v --clean --package techempower --generate-tests --target techempower ../../_testdata/examples/techempower.json

// Tests
//
//go:generate go run ../../cmd/ogen -v --clean --target test_webhooks         ../../_testdata/positive/webhooks.json
//go:generate go run ../../cmd/ogen -v --clean --target test_servers          ../../_testdata/positive/servers.json
//go:generate go run ../../cmd/ogen -v --clean --target test_single_endpoint  ../../_testdata/positive/single_endpoint.json
//go:generate go run ../../cmd/ogen -v --clean --target test_http_responses   ../../_testdata/positive/http_responses.json
//go:generate go run ../../cmd/ogen -v --clean --target test_http_requests    ../../_testdata/positive/http_requests.json
//go:generate go run ../../cmd/ogen -v --clean --target test_form             ../../_testdata/positive/form.json
//go:generate go run ../../cmd/ogen -v --clean --target test_parameters       ../../_testdata/positive/parameters.json
//go:generate go run ../../cmd/ogen -v --clean --target test_security         ../../_testdata/positive/security.json
//
//go:generate go run ./cmd/customformats ../../_testdata/positive/custom_formats.json test_customformats
//
//go:generate go run ../../cmd/ogen -v --clean --target referenced_path_item ../../_testdata/positive/referenced_pathItem.json
//
//go:generate go run ../../cmd/ogen -v --clean --generate-tests --target test_allof ../../_testdata/positive/allOf.yml
//go:generate go run ../../cmd/ogen -v --clean --generate-tests --target test_anyof ../../_testdata/positive/anyOf.json
//
//go:generate go run ../../cmd/ogen -v --clean -target test_enum_naming       ../../_testdata/positive/enum_naming.yml
//go:generate go run ../../cmd/ogen -v --clean -target test_naming_extensions ../../_testdata/positive/naming_extensions.json
