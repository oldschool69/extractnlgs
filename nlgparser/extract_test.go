package nlgparser

import "testing"

func BenchmarkExtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Extract("../nlg1.txt", "../nlg2.txt", "../out.txt")
	}
}

func BenchmarkExtract2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Extract("../nlg3.txt", "../nlg4.txt", "../out2.txt")
	}
}
