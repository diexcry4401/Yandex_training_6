package main

import "fmt"

func min(X, Y int) (R int) {
	if X != Y {
		if X < Y {
			return X
		} else {
			return Y
		}
	} else {
		return -1
	}
}

func max(X, Y int) (R int) {
	if X != Y {
		if X < Y {
			return Y
		} else {
			return X
		}
	} else {
		return -1
	}
}

func solution(A, B, C, D int) (M, N int) {
	if A == B {
		if A+B > C+D {
			M = 1
			N = max(C, D) + 1
		} else {
			M = A + 1
			N = 1
		}
	} else {
		if C == D {
			if C+D > A+B {
				M = max(A, B) + 1
				N = 1
			} else {
				M = 1
				N = C + 1
			}
		}
		if A == 0 || B == 0 {
			M = 1
			if C == 0 || D == 0 {
				N = 1
			} else {
				if A == 0 {
					N = C + 1
				} else {
					N = D + 1
				}
			}
		} else {
			if C == 0 || D == 0 {
				N = 1
				if C == 0 {
					M = A + 1
				} else {
					M = B + 1
				}
			} else {
				if (min(A, B) == A && min(C, D) == C) || (min(A, B) == B && min(C, D) == D) {
					var temp_M, temp_N int
					M = min(A, B) + 1
					N = min(C, D) + 1
					if max(A, B) > max(C, D) {
						temp_M = 1
						temp_N = max(C, D) + 1
					} else {
						if max(A, B) < max(C, D) {
							temp_M = max(A, B) + 1
							temp_N = 1
						} else {
							temp_M = max(A, B) + 1
							temp_N = 1
						}
					}
					if M+N > temp_M+temp_N {
						M = temp_M
						N = temp_N
					}
				} else {
					if (min(A, B) == B && min(C, D) == C) || (min(A, B) == A && min(C, D) == D) {
						if max(A, B) < max(C, D) {
							M = max(A, B) + 1
							N = 1
						} else {
							M = 1
							N = max(C, D) + 1
						}
					}
				}
			}
		}
	}

	return M, N
}

func main() {

	var A, B, C, D int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&D)

	fmt.Print(solution(A, B, C, D))

	// // 1 тест
	// A := 6
	// B := 2
	// C := 7
	// D := 3
	// fmt.Print(solution(A, B, C, D))
	// fmt.Printf("  \nans - 3 4 for values %d %d %d %d\n", A, B, C, D)

	// // 2 тест
	// A = 1
	// B = 1
	// C = 1
	// D = 1
	// fmt.Print(solution(A, B, C, D))
	// fmt.Printf(" \nans - 2 1 for values %d %d %d %d\n", A, B, C, D)

	// // 6 тест
	// A = 10
	// B = 7
	// C = 0
	// D = 4
	// fmt.Print(solution(A, B, C, D))
	// fmt.Printf(" \nans - 11 1 for values %d %d %d %d\n", A, B, C, D)

	// // 16 тест
	// A = 3
	// B = 3
	// C = 4
	// D = 7
	// fmt.Print(solution(A, B, C, D))
	// fmt.Printf(" \nans - 4 1 for values %d %d %d %d\n", A, B, C, D)

	// // 19 тест
	// A = 8
	// B = 9
	// C = 5
	// D = 9
	// fmt.Print(solution(A, B, C, D))
	// fmt.Printf("  \nans - 10 1 for values %d %d %d %d\n", A, B, C, D)

	// // 19 тест
	// A = 458
	// B = 865
	// C = 226
	// D = 791
	// fmt.Print(solution(A, B, C, D))
	// fmt.Printf("  \nans - 10 1 for values %d %d %d %d\n", A, B, C, D)

	// // 44 тест
	// A = 443908681
	// B = 899837908
	// C = 446114765
	// D = 894217789
	// fmt.Print(solution(A, B, C, D))
	// fmt.Printf("  \nans - 10 1 for values %d %d %d %d\n", A, B, C, D)

	// // Свой тест тест
	// A = 6
	// B = 2
	// C = 3
	// D = 7
	// fmt.Print(solution(A, B, C, D))
	// fmt.Printf("  \nans - 7 1 for values %d %d %d %d\n", A, B, C, D)
}
