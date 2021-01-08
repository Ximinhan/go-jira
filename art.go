package main

import (
	"fmt"
	"context"
	"io/ioutil"
	"github.com/andygrunwald/go-jira"
)

func main() {
	base := "https://issues.redhat.com/"
	tp := jira.BasicAuthTransport{
		Username: "ximhan",
		Password: "",
	}

	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}


	issueID := "ART-2146"
	remotelink := jira.RemoteLink{
		Object: &jira.RemoteLinkObject{
			URL: "https://github.com/openshift/art-ci-toolkit/pull/13",
			Title: "update image",
			Summary: "reduce image size",
			Icon: &jira.RemoteLinkIcon{
				Url16x16: "https://saml.buildvm.openshift.eng.bos.redhat.com:8888/favicon.ico",
				Title: "Jenkins Link",
			},
			Status: &jira.RemoteLinkStatus{
				Resolved: false,
				Icon: &jira.RemoteLinkIcon{
					Url16x16: "https://saml.buildvm.openshift.eng.bos.redhat.com:8888/favicon.ico",
					Title: "Jenkins Link",
				},
			},
		},
	}
	responseRemotelink, resp, err := jiraClient.Issue.AddRemoteLinkWithContext(context.Background(), issueID, &remotelink)
	if err != nil {
		panic(err)
	}
  body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s: %+v\n", responseRemotelink.Self, body)
}
