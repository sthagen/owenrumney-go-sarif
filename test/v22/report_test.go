package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_new_simple_report_with_single_run(t *testing.T) {
	given, _, then := newReportTest(t)

	given.a_new_report().
		with_a_run_added("tfsec", "https://tfsec.dev")
	then.report_text_is(`{"$schema":"https://raw.githubusercontent.com/oasis-tcs/sarif-spec/refs/heads/main/sarif-2.2/schema/sarif-2-2.schema.json","version":"2.2","runs":[{"tool":{"driver":{"notifications":[],"rules":[],"taxa":[],"locations":[],"informationURI":"https://tfsec.dev","name":"tfsec","contents":["localizedData","nonLocalizedData"],"supportedTaxonomies":[]},"extensions":[]},"results":[],"newlineSequences":["\r\n","\n"],"versionControlProvenance":[],"artifacts":[],"graphs":[],"threadFlowLocations":[],"translations":[],"policies":[],"invocations":[],"runAggregates":[],"taxonomies":[],"addresses":[],"webResponses":[],"logicalLocations":[],"redactionTokens":[],"webRequests":[]}],"inlineExternalProperties":[],"guid":"516d8cc4-18b1-463e-9e0e-417473149927","properties":{"tags":[]}}`)
}

func Test_new_report_with_empty_run(t *testing.T) {
	given, _, then := newReportTest(t)

	given.a_new_report().
		with_a_run_with_empty_result_added("tfsec", "https://tfsec.dev")
	then.report_text_is(`{"$schema":"https://raw.githubusercontent.com/oasis-tcs/sarif-spec/refs/heads/main/sarif-2.2/schema/sarif-2-2.schema.json","version":"2.2","runs":[{"tool":{"driver":{"notifications":[],"rules":[],"taxa":[],"locations":[],"informationURI":"https://tfsec.dev","name":"tfsec","contents":["localizedData","nonLocalizedData"],"supportedTaxonomies":[]},"extensions":[]},"results":[],"newlineSequences":["\r\n","\n"],"versionControlProvenance":[],"artifacts":[],"graphs":[],"threadFlowLocations":[],"translations":[],"policies":[],"invocations":[],"runAggregates":[],"taxonomies":[],"addresses":[],"webResponses":[],"logicalLocations":[],"redactionTokens":[],"webRequests":[]}],"inlineExternalProperties":[],"guid":"516d8cc4-18b1-463e-9e0e-417473149927","properties":{"tags":[]}}`)
}

func Test_new_simple_report_with_artifact(t *testing.T) {
	given, when, then := newReportTest(t)

	run := given.a_new_report().
		with_a_run_added("tfsec", "https://tfsec.dev")
	when.an_artifact_is_added_to_the_run(run, "file://broken.go")
	then.report_text_is(`{"$schema":"https://raw.githubusercontent.com/oasis-tcs/sarif-spec/refs/heads/main/sarif-2.2/schema/sarif-2-2.schema.json","version":"2.2","runs":[{"tool":{"driver":{"notifications":[],"rules":[],"taxa":[],"locations":[],"informationURI":"https://tfsec.dev","name":"tfsec","contents":["localizedData","nonLocalizedData"],"supportedTaxonomies":[]},"extensions":[]},"results":[],"newlineSequences":["\r\n","\n"],"versionControlProvenance":[],"artifacts":[{"location":{"uri":"file://broken.go"},"roles":[]}],"graphs":[],"threadFlowLocations":[],"translations":[],"policies":[],"invocations":[],"runAggregates":[],"taxonomies":[],"addresses":[],"webResponses":[],"logicalLocations":[],"redactionTokens":[],"webRequests":[]}],"inlineExternalProperties":[],"guid":"516d8cc4-18b1-463e-9e0e-417473149927","properties":{"tags":[]}}`)
}

func Test_new_simple_report_with_propertybag(t *testing.T) {
	given, when, then := newReportTest(t)

	run := given.a_new_report().
		with_a_run_added("tfsec", "https://tfsec.dev")
	when.some_properties_are_added_to_the_run(run)
	then.report_text_is(`{"$schema":"https://raw.githubusercontent.com/oasis-tcs/sarif-spec/refs/heads/main/sarif-2.2/schema/sarif-2-2.schema.json","version":"2.2","runs":[{"tool":{"driver":{"notifications":[],"rules":[],"taxa":[],"locations":[],"informationURI":"https://tfsec.dev","name":"tfsec","contents":["localizedData","nonLocalizedData"],"supportedTaxonomies":[]},"extensions":[]},"results":[],"newlineSequences":["\r\n","\n"],"versionControlProvenance":[],"artifacts":[],"graphs":[],"threadFlowLocations":[],"translations":[],"policies":[],"invocations":[],"runAggregates":[],"taxonomies":[],"addresses":[],"webResponses":[],"logicalLocations":[],"redactionTokens":[],"webRequests":[],"properties":{"properties":{"integer_property":10,"string_property":"this is a string"},"tags":[]}}],"inlineExternalProperties":[],"guid":"516d8cc4-18b1-463e-9e0e-417473149927","properties":{"tags":[]}}`)
}

func Test_new_simple_report_with_duplicate_artifact(t *testing.T) {
	given, when, then := newReportTest(t)

	run := given.a_new_report().
		with_a_run_added("tfsec", "https://tfsec.dev")
	when.an_artifact_is_added_to_the_run(run, "file://broken.go").
		and().
		an_artifact_is_added_to_the_run(run, "file://broken.go")
	then.report_text_is(`{"$schema":"https://raw.githubusercontent.com/oasis-tcs/sarif-spec/refs/heads/main/sarif-2.2/schema/sarif-2-2.schema.json","version":"2.2","runs":[{"tool":{"driver":{"notifications":[],"rules":[],"taxa":[],"locations":[],"informationURI":"https://tfsec.dev","name":"tfsec","contents":["localizedData","nonLocalizedData"],"supportedTaxonomies":[]},"extensions":[]},"results":[],"newlineSequences":["\r\n","\n"],"versionControlProvenance":[],"artifacts":[{"location":{"uri":"file://broken.go"},"roles":[]},{"location":{"uri":"file://broken.go"},"roles":[]}],"graphs":[],"threadFlowLocations":[],"translations":[],"policies":[],"invocations":[],"runAggregates":[],"taxonomies":[],"addresses":[],"webResponses":[],"logicalLocations":[],"redactionTokens":[],"webRequests":[]}],"inlineExternalProperties":[],"guid":"516d8cc4-18b1-463e-9e0e-417473149927","properties":{"tags":[]}}`)
}

func Test_load_sarif_from_string(t *testing.T) {
	given, _, then := newReportTest(t)

	content := `{
  "version": "2.1.0",
  "$schema": "https://docs.oasis-open.org/sarif/sarif/v2.1.0/errata01/os/schemas/sarif-schema-2.1.0.json",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "ESLint",
          "informationURI": "https://eslint.org"
        }
      }
    }
  ]
}`

	given.a_report_is_loaded_from_a_string(content)
	then.the_report_has_expected_driver_name_and_information_uri("ESLint", "https://eslint.org")
}

func Test_load_sarif_report_from_file(t *testing.T) {
	given, _, then := newReportTest(t)

	content := `{
  "version": "2.1.0",
  "$schema": "https://docs.oasis-open.org/sarif/sarif/v2.1.0/errata01/os/schemas/sarif-schema-2.1.0.json",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "ESLint",
          "informationURI": "https://eslint.org"
        }
      }
    }
  ]
}`

	file, err := os.CreateTemp(t.TempDir(), "sarifReport")
	assert.NoError(t, err)
	defer file.Close()

	os.WriteFile(file.Name(), []byte(content), 755)

	given.a_report_is_loaded_from_a_file(file.Name())
	then.the_report_has_expected_driver_name_and_information_uri("ESLint", "https://eslint.org")
}

func Test_err_on_load_sarif_report_from_file_when_not_exists(t *testing.T) {
	given, _, _ := newReportTest(t)
	defer func() {
		if err := recover().(error); err != nil {
			assert.Equal(t, "the provided file path doesn't have a file", err.Error())
		}
	}()
	given.a_report_is_loaded_from_a_file("")
}

func Test_err_on_load_sarif_report_from_file_when_file_not_legit(t *testing.T) {
	given, _, _ := newReportTest(t)
	defer func() {
		if err := recover().(error); err != nil {
			assert.Equal(t, "the provided filepath could not be opened. read /tmp: is a directory", err.Error())
		}
	}()
	given.a_report_is_loaded_from_a_file("/tmp")
}
