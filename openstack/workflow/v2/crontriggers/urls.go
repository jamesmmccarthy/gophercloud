package crontriggers

import "github.com/gophercloud/gophercloud"

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("cron_triggers")
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("cron_triggers", id)
}
