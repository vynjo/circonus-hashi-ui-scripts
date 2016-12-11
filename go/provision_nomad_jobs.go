package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/circonus-labs/circonus-gometrics/api"
)

// Allocation is a struct containing state of a nomad allocation
type Allocation struct {
	// ID of the allocation (UUID)
	ID string
	// Name is a logical name of the allocation.
	Name string
	// NodeID is the node this is being placed on
	NodeID string
	// JobID is the parent job of the task group being allocated.
	// This is copied at allocation time to avoid issues if the job
	// definition is updated.
	JobID string
	// ClientStatus of the allocation on the client
	ClientStatus string
}

type NomadJob struct {
	ID 					string 		`json:"ID"`
	ParentID 			string		`json:"ParentID"`
	Name 				string		`json:"Name"`
	Type 				string		`json:"Type"`
	Status 				string 		`json:"Status"`
	StatusDescription 	string 		`json:"StatusDescription"`
	CreateIndex 		int 		`json:"CreateIndex"`
	ModifyIndex 		int 		`json:"ModifyIndex"`
	JobModifyIndex 		int 		`json:"JobModifyIndex"`
}

type metricSearchResult struct {
	CheckBundleID string   `json:"_check_bundle"`
	Name          string   `json:"_metric_name"`
	Type          string   `json:"_metric_type"`
	Tags          []string `json:"tags"`
	Units         string   `json:"units"`
}

type checkBundleMetric struct {
	Name   string   `json:"name"`
	Status string   `json:"status"`
	Tags   []string `json:"tags"`
	Type   string   `json:"type"`
	Units  string   `json:"units"`
	Result string   `json:"result,omitempty"`
}

type checkBundleMetricList struct {
	Metrics []checkBundleMetric `json:"metrics"`
}

type checkBundleMetricResult struct {
	CID     string              `json:"_cid"`
	Metrics []checkBundleMetric `json:"metrics"`
}

func getAllocationMetrics(id string) ([]metricSearchResult, error) {

	metricSearchURL := fmt.Sprintf("/metric?search=(active:1)*%s*", id)

	metricsJSON, err := circapi.Get(metricSearchURL)
	if err != nil {
		return nil, err
	}

	metrics := []metricSearchResult{}

	err = json.Unmarshal(metricsJSON, &metrics)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

func deactivateMetrics(checkBundleID string, metricList checkBundleMetricList) error {
	reqPath := fmt.Sprintf("/check_bundle_metrics/%s", checkBundleID)

	metricsJSON, err := json.Marshal(metricList)
	if err != nil {
		return err
	}

	response, err := circapi.Put(reqPath, metricsJSON)
	if err != nil {
		return err
	}

	result := checkBundleMetricResult{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}

	for _, metric := range result.Metrics {
		log.Printf("\tmetric: %s, status: %s, result: %s\n", metric.Name, metric.Status, metric.Result)
	}

	log.Println("---")

	return nil
}

func updateMetrics(allocation Allocation) error {

	log.Printf("\tchecking for active metrics\n")

	allocationMetrics, err := getAllocationMetrics(allocation.ID)
	if err != nil {
		return err
	}

	if len(allocationMetrics) == 0 {
		log.Printf("\t0 active metrics, skipping\n")
		return nil
	}

	log.Printf("\tdeactivating %d active metrics\n", len(allocationMetrics))

	checkBundleID := ""
	metrics := []checkBundleMetric{}

	for _, metric := range allocationMetrics {
		if checkBundleID == "" {
			checkBundleID = strings.Replace(metric.CheckBundleID, "/check_bundle/", "", -1)
		}
		metrics = append(metrics, checkBundleMetric{
			Name:   metric.Name,
			Status: "available",
			Tags:   metric.Tags,
			Type:   metric.Type,
			Units:  metric.Units,
		})
	}

	metricList := checkBundleMetricList{metrics}

	err = deactivateMetrics(checkBundleID, metricList)
	if err != nil {
		return err
	}

	return nil
}

var (
	circapi  *api.API
	nomadURL *url.URL
)

func getJobs() ([]NomadJob, error) {
	reqURL := nomadURL.String()
	if !strings.Contains(reqURL, "v1/jobs") {
		reqURL += "v1/jobs"
	}

	res, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	nomadjobs := []NomadJob{}

	err = json.Unmarshal(body, &nomadjobs)
	if err != nil {
		return nil, err
	}

	return nomadjobs, nil
}

func setup() {
	var err error
	var apiKey string
	var apiApp string
	var apiURL string
	var nomadAPIURL string
	var debug bool

	flag.StringVar(&apiKey, "key", "", "Circonus API Token Key [none] (CIRCONUS_API_KEY)")
	flag.StringVar(&apiApp, "app", "", "Circonus API Token App [nomad-metric-reaper] (CIRCONUS_API_APP)")
	flag.StringVar(&apiURL, "apiurl", "", "Base Circonus API URL [https://api.circonus.com/] (CIRCONUS_API_URL)")
	flag.StringVar(&nomadAPIURL, "nomadurl", "", "Base Nomad API URL [http://localhost:4646/] (NOMAD_API_URL)")
	flag.BoolVar(&debug, "debug", false, "Enable Circonus API debugging")

	flag.Parse()

	cfg := &api.Config{}

	if apiKey == "" {
		apiKey = os.Getenv("CIRCONUS_API_KEY")
		if apiKey == "" {
			log.Printf("CIRCONUS_API_KEY is not set, exiting.\n")
			os.Exit(1)
		}
	}
	cfg.TokenKey = apiKey

	if apiApp == "" {
		apiApp := os.Getenv("CIRCONUS_API_APP")
		if apiApp == "" {
			apiApp = "nomad-metrics-reaper"
		}
	}
	cfg.TokenApp = apiApp

	if apiURL == "" {
		apiURL = os.Getenv("CIRCONUS_API_URL")
		if apiURL == "" {
			apiURL = "https://api.circonus.com/"
		}
	}
	cfg.URL = apiURL

	cfg.Debug = debug

	circapi, err = api.NewAPI(cfg)
	if err != nil {
		log.Printf("ERROR: allocating Circonus API %v\n", err)
		os.Exit(1)
	}

	if nomadAPIURL == "" {
		nomadAPIURL = os.Getenv("NOMAD_API_URL")
		if nomadAPIURL == "" {
			nomadAPIURL = "http://localhost:4646/"
		}
	}
	nomadURL, err = url.Parse(nomadAPIURL)
	if err != nil {
		log.Printf("ERROR: parsing Nomad API URL %+v\n", err)
		os.Exit(1)
	}
}

func main() {

	setup()

	log.Println("Retrieving jobs from Nomad")

	nomadjobs, err := getJobs()
	if err != nil {
		log.Printf("ERROR: retrieving jobs %v\n", err)
		os.Exit(1)
	}

	if len(nomadjobs) == 0 {
		log.Println("No jobs found, exiting.")
		os.Exit(0)
	}

	for _, job := range nomadjobs {
		log.Printf("Processing job %s with status -  %s  (%s:%s)\n", job.Name, job.Status, job.Type, job.ID)
// 		err := updateMetrics(allocation)
// 		if err != nil {
// 			log.Printf("ERROR: %+v", err)
// 			os.Exit(1)
// 		}
	}
}

