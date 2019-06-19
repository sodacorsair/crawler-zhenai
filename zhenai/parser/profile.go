package parser

import (
	"github.com/Eru/crawler/engine"
	"github.com/Eru/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>([0-9]+)岁</div>`)
var marriageRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>(未婚|离异|丧偶)</div>`)
var nameRe = regexp.MustCompile(
	`<meta name="keywords" content="(.+)照片, .+资料, .+征婚交友">`)
var genderRe = regexp.MustCompile(
	`<div class="m-title" data-v-bff6f798>(.)的动态</div>`)
var heightRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>([0-9]+)cm</div>`)
var weightRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>([0-9]+)kg</div>`)
var incomeRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)
var educationRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>(博士|硕士|大学本科|大专|中专|高中及以下)</div>`)
var occupationRe = regexp.MustCompile(
	`月收入:[^<]+</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)
var hukouRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-bff6f798>籍贯:([^<]+)</div>`)
var constellationRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>([^>]+座)\([\d]{2}\.[\d]{2}-[\d]{2}\.[\d]{2}\)</div>`)
var carRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-bff6f798>(已买车|未买车)</div>`)
var houseRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-bff6f798>(已购房|租房)</div>`)

func ParseProfile(
	contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

    age, err := strconv.Atoi(
    	extractString(contents, ageRe))
    if err == nil {
    	profile.Age = age
	}
    profile.Marriage = extractString(
    	contents, marriageRe)
    profile.Name = extractString(
    	contents, nameRe)
    if (extractString(contents, genderRe) == "她") {
    	profile.Gender = "女"
	} else {
		profile.Gender = "男"
	}
	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Income = extractString(
		contents, incomeRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Hukou = extractString(
		contents, hukouRe)
	profile.Constellation = extractString(
		contents, constellationRe)
	profile.Car = extractString(
		contents, carRe)
	profile.House = extractString(
		contents, houseRe)

	result := engine.ParseResult{
		Items: []interface{} {profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}