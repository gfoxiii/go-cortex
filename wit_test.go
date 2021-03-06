package main

import (
	"bytes"
	"io"
	"testing"
)

type NopCloser struct {
	io.Reader
}

func (NopCloser) Close() error { return nil }

func TestProcessWitResponseGithubMultipleIssues(t *testing.T) {

	withJSON := stringToReadeClosser(githubMultipleIssues)
	numbers := ProcessWitResponse(withJSON).Outcome.Entities.MultipleNumber

	if len(numbers) != 2 {
		t.Errorf("ProcessWitResponse didn't parse the 'numbers' array. We got %+v\n", numbers)
	}
}

func TestProcessWitResponseGithubSingleIssue(t *testing.T) {

	withJSON := stringToReadeClosser(githubSingleIssue)
	number := ProcessWitResponse(withJSON).Outcome.Entities.MultipleNumber

	if len(number) != 1 {
		t.Errorf("ProcessWitResponse didn't parse the 'number' object. We got %+v\n", number)
	}

	if number[0].Value != 45 {
		t.Errorf("Head item's Value is not 45. We got %+v\n", number)
	}
}

func TestProcessWitResponseSingleLight(t *testing.T) {

	withJSON := stringToReadeClosser(lightPayload)
	number := ProcessWitResponse(withJSON).Outcome.Entities.SingleNumber

	if number.Value != 1 {
		t.Errorf("ProcessWitResponse didn't parse the 'number' object. We got %+v\n", number)
	}

}

func stringToReadeClosser(s string) io.ReadCloser {
	return NopCloser{bytes.NewBufferString(s)}
}

func TestSanitizeQuerryStringStringLen(t *testing.T) {
	_, err := sanitizeQuerryString(string300)
	if err == nil {
		t.Error("FetchIntent did not return an error for a string input of 300 chars")
	}
	_, err = sanitizeQuerryString(string254)
	if err != nil {
		t.Errorf("FetchIntent returned an error for a string input of 254 chars %+v", err)
	}
}

const string300 = (`245485328217529591072968367825520430801937353549236235032205454278011159517553408301117871215897624083557692321819508308225339640853054008672033271569751783199322357002915818244872430853340789879400481978383988517251094914866992168126566388692301329752249123938027308855068750472072224632356977779896`)

const string254 = (`24548532821752959107296836782552043080193735354923623503220545427801115951755340830111787121589762408s3557692362408s355769232181950830822533964085305400867203327156975178319932235700291581824487243085334078987940048197838398851725109497222463235697777989`)

const githubMultipleIssues = `{
  "msg_body": "look at #45, #102 and work on those",
  "outcome": {
    "intent": "github",
    "confidence": 0.997,
    "entities": {
      "github_issue": [
        {
          "value": 45,
          "body": "45,",
          "start": 9,
          "end": 11
        },
        {
          "value": 102,
          "body": "102 ",
          "start": 14,
          "end": 17
        }
      ]
    }
  }
}`

const githubSingleIssue = `{
  "msg_body": "look at #45 and work on those",
  "outcome": {
    "intent": "github",
    "confidence": 0.997,
    "entities": {
      "github_issue":
        {
          "value": 45,
          "body": "45,",
          "start": 9,
          "end": 11
        }
    }
  }
}`

const lightPayload = `{
  "msg_body": "turn the light one on please",
  "outcome": {
    "intent": "lights",
    "confidence": 1,
    "entities": {
      "on_off": {
        "value": "on"
      },
      "number": {
        "value": 1,
        "body": "one ",
        "start": 15,
        "end": 18
      }
    }
  }
}`
