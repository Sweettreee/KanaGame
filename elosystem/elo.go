package elosystem

import (
	"math"
)

const maxK = 40     // 최대 K값, 클수록 초반에 레이팅이 크게 변화
const minK = 32     // 최소 K값, 작을수록 후반에 레이팅이 적게 변화
const newRound = 20 // 신규유저로 인식할 라운드수
const scail = 400.0 // 작을수록 레이팅 차이에 대해 더 민감하게 반응

func GetKValue(gameCount int) int {
	if gameCount == 0 {
		return maxK
	}
	if gameCount >= newRound {
		return minK
	}
	return int(math.Round(maxK - float64(gameCount)*(maxK-minK)/newRound))
}

func ExpectedScore(ratingA, ratingB int) float64 {
	return 1.0 / (1.0 + math.Pow(10, float64(ratingB-ratingA)/scail))
}

func UpdateElo(ratingA, ratingB int, scoreA float64, gameCountA, gameCountB int) (int, int) {
	K_A := float64(GetKValue(gameCountA))
	K_B := float64(GetKValue(gameCountB))

	Ea := ExpectedScore(ratingA, ratingB)

	newRatingA := int(math.Round(float64(ratingA) + K_A*(scoreA-Ea)))
	newRatingB := int(math.Round(float64(ratingB) + K_B*((1-scoreA)-(1-Ea))))

	return newRatingA, newRatingB
}
