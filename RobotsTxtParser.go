package RobotsTxtParser

import (
	"fmt"
	"regexp"
	"strings"
	// "golang.org/x/exp/slices"
)

type KV struct {
    Key   string
    Value string
}

type UserAgentRule struct {
	UserAgents []string
	Allow []string
	Disallow []string
}

type RobotsTxt struct {
	UserAgentRules []UserAgentRule
	Sitemaps []string
}

func ParseTxt(txt string) (RobotsTxt, error) {

	kv_of_kvs := [][]KV{}
	pairs := splitByPairs(txt)

	for _, current_pair := range pairs {
		kvs := getLinesWithKeys(current_pair)

		if kvs != nil {
			kv_of_kvs = append(kv_of_kvs, kvs)
		}
	}

	rt := transformToRobotsTxt(kv_of_kvs)
	
	return rt, nil
}

func transformToRobotsTxt(kv_of_kvs [][]KV) RobotsTxt {
	var userAgentRules []UserAgentRule
	var sitemaps []string

	for _, kv_group := range kv_of_kvs {
		sitemapKvs := filterForKey(kv_group, "sitemap")

		if sitemapKvs != nil {
			sitemaps = append(sitemaps, sitemapKvs...)
		}

		if kv_group != nil && kvGroupContainsKeys(kv_group, []string{"user-agent", "allow", "disallow"}){
			ua_rule := UserAgentRule{
				UserAgents: filterForKey(kv_group, "user-agent"),
				Allow: filterForKey(kv_group, "allow"),
				Disallow: filterForKey(kv_group, "disallow"),
			}
			userAgentRules = append(userAgentRules, ua_rule)
		}
	}


	rt := RobotsTxt{
		UserAgentRules: userAgentRules,
		Sitemaps: sitemaps,
	}
	return rt
}

func splitByPairs(text string) []string {
	pattern := regexp.MustCompile(`(?m)\n(\n|#\n)`)
	pairs := pattern.Split(text, -1)

	var unescaped_pairs []string
	for _, item := range pairs {
		escaped_item := strings.Replace(item, `\n`, "\n", -1)
		unescaped_pairs = append(unescaped_pairs, escaped_item)
	}
	return unescaped_pairs
}

func getLinesWithKeys(text string) []KV {
	const KEY_POSITION = 2
	const VALUE_POSITION = 4

	pattern := regexp.MustCompile(`(?m)^([^#].*)?^\s*([\w|-]*):(\t|.|)(.*|)$`)
	key_lines := pattern.FindAllStringSubmatch(text, -1)

	var key_value []KV
	for _, element := range key_lines {
		if !strings.Contains(element[KEY_POSITION], "#") {
			key_value = append(key_value, KV{
				Key: strings.ToLower(element[KEY_POSITION]), 
				Value: element[VALUE_POSITION],
			})
		}
	}
	return key_value
}

func filterForKey(kvs []KV, key string) []string {
    var out []string
    for _, kv := range kvs {
        if kv.Key == key {
            out = append(out, kv.Value)
        }
    }
    return out
}

func kvGroupContainsKeys(kvs []KV, keys []string) bool {
	containsKeys := false
	for _, key := range keys {
		for _, kv := range kvs {
			if kv.Key == key {
				containsKeys = true
			}
		}
	}
	return containsKeys
}