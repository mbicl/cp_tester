package server

import (
	"errors"
	"fmt"
	"strings"
)

func ExtractPlatformAndProblemName(url string) (string, string, string, error) {
	url = plainUrl(strings.TrimSpace(url))
	parts := strings.Split(url, "/")
	if parser, ok := platformParsers[parts[0]]; ok {
		return parser(parts)
	}
	return "", "", "", errors.New("unsupported platform")
}

func plainUrl(url string) string {
	if len(url) == 0 {
		return ""
	}

	if url[0] == '/' {
		return plainUrl(url[1:])
	}

	if url[len(url)-1] == '/' {
		return plainUrl(url[:len(url)-1])
	}

	if url[:8] == "https://" {
		return plainUrl(url[8:])
	}

	if url[:7] == "http://" {
		return plainUrl(url[7:])
	}

	if url[:4] == "www." {
		return plainUrl(url[4:])
	}

	return url
}

var platformParsers = map[string]func(parts []string) (platform string, contestId string, problemId string, err error){
	"codeforces.com":    codeforcesParser,
	"atcoder.jp":        atcoderParser,
	"acmp.ru":           acmpParser,
	"robocontest.uz":    robocontestParser,
	"kep.uz":            kepParser,
	"cses.fi":           csesParser,
	"codechef.com":      codechefParser,
	"coderun.yandex.ru": coderunParser,
	"contest.yandex.ru": yandexParser,
	"acm.timus.ru":      timusParser,
	"dmoj.ca":           dmojParser,
	"eolymp.com":        eolympParser,
	"hackerrank.com":    hackerrankParser,
	"kattis.com":        kattisParser,
	"luogu.com.cn":      luoguParser,
	"sortme.org":        sortmeParser,
	"spoj.com":          spojParser,
	"usaco.org":         usacoParser,
	"onlinejudge.org":   onlinejudgeParser,
}

func codeforcesParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "codeforces"
	// codeforces.com/problemset/problem/2203/A - problemset
	// codeforces.com/contest/2218/problem/A - contest
	// codeforces.com/gym/106403/problem/A - gym
	// codeforces.com/edu/course/2/lesson/2/1/practice/contest/269100/problem/A - edu
	// codeforces.com/group/t2cQ2lPSyS/contest/604276/problem/C - group
	if len(parts) < 5 {
		return "", "", "", errors.New("invalid url format for codeforces")
	}
	switch parts[1] {
	case "problemset":
		return platform, "", parts[3] + parts[4], nil
	case "contest":
		return platform, parts[2], parts[4], nil
	case "gym":
		return platform, "", "gym" + parts[2] + parts[4], nil
	case "edu":
		if len(parts) < 12 {
			return "", "", "", errors.New("invalid url format for codeforces edu")
		}
		return platform, "", fmt.Sprintf("edu_%s_%s_%s%s", parts[3], parts[5], parts[6], parts[11]), nil
	case "group":
		if len(parts) < 7 {
			return "", "", "", errors.New("invalid url format for codeforces group")
		}
		return platform, "", "group" + parts[4] + parts[6], nil
	}
	return "", "", "", errors.New("unsupported codeforces url format")
}

func atcoderParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "atcoder"
	// atcoder.jp/contests/abc452/tasks/abc452_a - contest
	if len(parts) < 5 {
		return "", "", "", errors.New("invalid url format for atcoder")
	}
	return platform, parts[2], parts[4], nil
}

func acmpParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "acmp"
	// acmp.ru/asp/do/index.asp?main=task&id_course=1&id_section=1&id_topic=26&id_problem=142 - from course
	// acmp.ru/index.asp?main=task&id_task=1 - from archive
	if len(parts) < 2 {
		return "", "", "", errors.New("invalid url format for acmp")
	}
	if parts[1] == "asp" {
		if len(parts) < 4 {
			return "", "", "", errors.New("invalid url format for acmp course")
		}
		queries := strings.Split(parts[3], "&")
		courseId := queries[0][len("id_course="):]
		problemId := queries[3][len("id_problem="):]
		return platform, "", "course" + courseId + "_" + problemId, nil
	} else {
		task_id := strings.TrimPrefix(parts[1], "index.asp?main=task&id_task=")
		if task_id == parts[1] {
			return "", "", "", errors.New("invalid url format for acmp archive")
		}
		return platform, "", task_id, nil
	}
}

func robocontestParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "robocontest"
	// robocontest.uz/tasks/A0001 - problemset
	// robocontest.uz/olympiads/3209/tasks/B - contest

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for robocontest")
	}

	switch parts[1] {
	case "tasks":
		return platform, "", parts[2], nil
	case "olympiads":
		if len(parts) < 4 {
			return "", "", "", errors.New("invalid url format for robocontest olympiad")
		}
		return platform, parts[2], parts[4], nil
	}

	return "", "", "", errors.New("unsupported robocontest url format")
}

func kepParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "kep"
	// kep.uz/problems/1 - problemset
	// kep.uz/contests/488/problem/A - contest

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for kep")
	}

	switch parts[1] {
	case "problems":
		return platform, "", parts[2], nil
	case "contests":
		if len(parts) < 5 {
			return "", "", "", errors.New("invalid url format for kep contest")
		}
		return platform, parts[2], parts[4], nil
	}

	return "", "", "", errors.New("unsupported kep url format")
}

func csesParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "cses"
	// cses.fi/641/task/A - contest
	// cses.fi/problemset/task/1068 - problemset

	if len(parts) < 4 {
		return "", "", "", errors.New("invalid url format for cses")
	}

	if parts[1] == "problemset" {
		return platform, "", parts[3], nil
	} else {
		return platform, parts[1], parts[3], nil
	}
}

func codechefParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "codechef"
	// codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/CWC23QUALIF - practice
	// codechef.com/problems/MXSZ - problemset/contest

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for codechef")
	}

	return platform, "", parts[len(parts)-1], nil
}

func coderunParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "coderun"
	// coderun.yandex.ru/selections/quickstart/problems/gcd-and-lcm - selections/selection_id/problems/problem_id
	// coderun.yandex.ru/problem/histogram-and-rectangle - problem/problem_id

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for coderun")
	}

	switch parts[1] {
	case "selections":
		if len(parts) < 5 {
			return "", "", "", errors.New("invalid url format for coderun selection")
		}
		return platform, parts[2], parts[4], nil
	case "problem":
		return platform, "", parts[2], nil
	}

	return "", "", "", errors.New("unsupported coderun url format")
}

func yandexParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "yandex"
	// contest.yandex.ru/contest/79803/problems/A/ - contest

	if len(parts) < 5 {
		return "", "", "", errors.New("invalid url format for yandex contest")
	}

	return platform, parts[2], parts[4], nil
}

func timusParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "timus"
	// acm.timus.ru/problem.aspx?space=1&num=1005 - problemset, space problemset id, num problem id

	if len(parts) < 2 {
		return "", "", "", errors.New("invalid url format for timus")
	}

	queries := strings.Split(parts[1], "?")
	if len(queries) < 2 {
		return "", "", "", errors.New("invalid url format for timus")
	}

	params := strings.Split(queries[1], "&")
	if len(params) < 2 {
		return "", "", "", errors.New("invalid url format for timus")
	}

	space := ""
	num := ""

	for _, param := range params {
		if strings.HasPrefix(param, "space=") {
			space = param[len("space="):]
		} else if strings.HasPrefix(param, "num=") {
			num = param[len("num="):]
		}
	}

	if space == "" || num == "" {
		return "", "", "", errors.New("invalid url format for timus")
	}

	return platform, space, num, nil
}

func dmojParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "dmoj"
	// dmoj.ca/problem/a4b1 - problemset

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for dmoj")
	}

	return platform, "", parts[2], nil
}

func eolympParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "eolymp"
	// eolymp.com/en/problems/10 - problemset
	// eolymp.com/en/compete/4pe1cne4r571j089budf65juq4/problem/3 - contest
	// eolymp.com/en/courses/5k3e59fd7t4il2r0tprott0npc/modules/p9ecb/items/l3bph - course

	if len(parts) < 4 {
		return "", "", "", errors.New("invalid url format for eolymp")
	}

	switch parts[1] {
	case "problems":
		return platform, "", parts[3], nil
	case "compete":
		if len(parts) < 6 {
			return "", "", "", errors.New("invalid url format for eolymp compete")
		}
		return platform, parts[3], parts[5], nil
	case "courses":
		if len(parts) < 8 {
			return "", "", "", errors.New("invalid url format for eolymp course")
		}
		return platform, parts[2], parts[7], nil
	}

	return "", "", "", errors.New("unsupported eolymp url format")
}

func hackerrankParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "hackerrank"
	// hackerrank.com/challenges/birthday-cake-candles/problem?isFullScreen=true - problemset

	if len(parts) < 4 {
		return "", "", "", errors.New("invalid url format for hackerrank")
	}

	return platform, "", parts[2], nil
}

func kattisParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "kattis"
	// open.kattis.com/problems/sequences - problemset
	// open.kattis.com/contests/xcr52d/problems/nodup - contest
	// open.kattis.com/contests/uzcy9d/problems/squaredeal - contest

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for kattis")
	}

	switch parts[1] {
	case "problems":
		return platform, "", parts[2], nil
	case "contests":
		if len(parts) < 5 {
			return "", "", "", errors.New("invalid url format for kattis contest")
		}
		return platform, parts[2], parts[4], nil
	}

	return "", "", "", errors.New("unsupported kattis url format")
}

func luoguParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "luogu"
	// luogu.com.cn/problem/P1001 - problemset
	// luogu.com.cn/problem/P16213?contestId=319886 - contest

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for luogu")
	}

	if strings.Contains(parts[2], "contestId") {
		subparts := strings.Split(parts[2], "?")
		if len(subparts) < 2 {
			return "", "", "", errors.New("invalid url format for luogu contest")
		}
		problemId := subparts[0]
		contestId := strings.TrimPrefix(subparts[1], "contestId=")
		return platform, contestId, problemId, nil
	} else {
		return platform, "", parts[2], nil
	}
}

func sortmeParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "sortme"
	// sort-me.org/tasks/1122?archive=6&hidesolved=1&category=0 - contest
	// sort-me.org/tasks/22 - problemset

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for sortme")
	}

	if strings.Contains(parts[2], "?") {
		subparts := strings.Split(parts[2], "?")
		return platform, "", subparts[0], nil
	}

	return platform, "", parts[2], nil
}

func spojParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "spoj"
	// spoj.com/problems/HOTLINE/ - problemset
	// spoj.com/ALGO24S1/problems/ALGOUPT08/ - contest

	if len(parts) < 3 {
		return "", "", "", errors.New("invalid url format for spoj")
	}

	if parts[1] == "problems" {
		return platform, "", parts[2], nil
	} else {
		if len(parts) < 4 {
			return "", "", "", errors.New("invalid url format for spoj contest")
		}
		return platform, parts[1], parts[3], nil
	}
}

func usacoParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "usaco"
	// usaco.org/index.php?page=viewproblem2&cpid=1542 - problemset

	if len(parts) < 2 {
		return "", "", "", errors.New("invalid url format for usaco")
	}

	queries := strings.Split(parts[1], "?")
	if len(queries) < 2 {
		return "", "", "", errors.New("invalid url format for usaco")
	}

	params := strings.SplitSeq(queries[1], "&")
	for param := range params {
		if strings.HasPrefix(param, "cpid=") {
			problemId := param[len("cpid="):]
			return platform, "", problemId, nil
		}
	}

	return "", "", "", errors.New("invalid url format for usaco")
}

func onlinejudgeParser(parts []string) (platform string, contestId string, problemId string, err error) {
	platform = "onlinejudge"

	// onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=3&page=show_problem&problem=123 - problemset

	if len(parts) < 2 {
		return "", "", "", errors.New("invalid url format for onlinejudge")
	}

	queries := strings.Split(parts[1], "?")
	if len(queries) < 2 {
		return "", "", "", errors.New("invalid url format for onlinejudge")
	}

	params := strings.SplitSeq(queries[1], "&")
	for param := range params {
		if strings.HasPrefix(param, "problem=") {
			problemId := param[len("problem="):]
			return platform, "", problemId, nil
		}
	}

	return "", "", "", errors.New("invalid url format for onlinejudge")
}
