package steamid

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	steamid64Identifier uint64 = 76561197960265728
)

func IDToID3(steamId string) string {
	steamid_split := strings.Split(steamId, ":")
	id3 := []string{"[U:1:"}

	y, _ := strconv.Atoi(steamid_split[1])
	z, _ := strconv.Atoi(steamid_split[2])

	accountId := z*2 + y

	id3 = append(id3, fmt.Sprintf("%d]", accountId))

	return strings.Join(id3, "")
}

func IDToID64(steamId string) string {
	steamid_split := strings.Split(steamId, ":")
	id64, _ := strconv.Atoi(steamid_split[2])
	id64 = id64 * 2

	if steamid_split[1] == "1" {
		id64 += 1
	}

	id64 += int(steamid64Identifier)
	return fmt.Sprintf("%d", id64)
}

func ID3ToID(steamId3 string) string {
	steamId3 = strings.Replace(steamId3, "[", "", 1)
	steamId3 = strings.Replace(steamId3, "]", "", 1)

	steamId3_split := strings.Split(steamId3, ":")
	id := []string{"STEAM_0:"}

	z, _ := strconv.Atoi(steamId3_split[2])
	if z%2 == 0 {
		id = append(id, "0:")
	} else {
		id = append(id, "1:")
	}

	accountId := z / 2
	id = append(id, fmt.Sprintf("%d", accountId))

	return strings.Join(id, "")
}

func ID3ToID64(steamId3 string) string {
	steamId3 = strings.Replace(steamId3, "[", "", 1)
	steamId3 = strings.Replace(steamId3, "]", "", 1)

	steamId3_split := strings.Split(steamId3, ":")
	id, _ := strconv.Atoi(steamId3_split[2])
	id64 := id + int(steamid64Identifier)
	return fmt.Sprintf("%d", id64)
}

func ID64ToID(steamId64 string) string {
	id := []string{"STEAM_0:"}

	id64_int, _ := strconv.Atoi(steamId64)
	accountId := id64_int - int(steamid64Identifier)
	if accountId%2 == 0 {
		id = append(id, "0:")
	} else {
		id = append(id, "1:")
	}

	id = append(id, fmt.Sprintf("%d", accountId/2))
	return strings.Join(id, "")
}

func ID64ToID3(steamId64 string) string {
	id3 := []string{"[U:1:"}
	id, _ := strconv.Atoi(steamId64)
	accountId := id - int(steamid64Identifier)

	id3 = append(id3, fmt.Sprintf("%d]", accountId))
	return strings.Join(id3, "")
}
