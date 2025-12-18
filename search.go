package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func ListIssues(terms []string) {
	result, err := SearchIssues(terms)
	if err != nil {
		fmt.Println("Error searching issues:", err)
		return
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %9.9s %.55s\n",
			item.Number, item.User.Login, item.Age, item.Title)
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	oneYearAgo := time.Now().AddDate(-1, 0, 0)
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	for _, item := range result.Items {
		createdAt, err := time.Parse("2006-01-02T15:04:05Z", item.CreatedAt)
		if err != nil {
			return nil, err
		}
		if createdAt.After(oneMonthAgo) {
			item.Age = LessThanMonth
		} else if createdAt.After(oneYearAgo) {
			item.Age = LessThanYear
		} else {
			item.Age = MoreThanYear
		}
	}

	return &result, nil
}
