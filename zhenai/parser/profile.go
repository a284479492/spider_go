package parser

import (
	"regexp"
	"strconv"

	"crawler/model"
	"crawler/types"
)

// var nameRe = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriedRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var guessFavRe = regexp.MustCompile(`<a class="exp-user-name"[^>]* href="(http://localhost:8080/mock/album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(`http://localhost:8080/mock/album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) types.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	// profile.Name = extractString(contents, nameRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriedRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	} else {
		profile.Age = 0
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	} else {
		profile.Height = 0
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	} else {
		profile.Weight = 0
	}

	result := types.ParseResult{
		Items: []types.Item{{
			Url:     url,
			Type:    "zhenai",
			Id:      extractString([]byte(url), idUrlRe),
			PayLoad: profile,
		}},
	}

	matchs := guessFavRe.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		result.Requests = append(result.Requests, types.Request{
			URL:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

// func profileParser(name string) engine.ParseFunc {
// 	return func(c []byte, url string) engine.ParseResult {
// 		return ParseProfile(c, url, name)
// 	}
// }
