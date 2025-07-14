package main

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/Marlliton/temp/pkg/logger"
	"github.com/Marlliton/temp/pkg/stack"
)

func main() {
	logger.Init()
	st := stack.Stack[int]{}
	st.Push(1)
	st.Push(2)
	st.Push(3)
	slog.Info("Mostrando itens", "itens", st.Items)

	e, _ := st.Pop()
	slog.Info("removendo item", "item", e)

	e, _ = st.Peek()
	slog.Info("pegando útilmo item", "item", e)

	size := st.Size()
	slog.Info("Tamando da pilha", "valor", size)

	st.Clear()
	slog.Info("Limpando pilha", "Esta vazio?", st.IsEmpty())
}

func RunStackDecimalToBinay(num int) {
	st := stack.Stack[int]{}
	originalValue := num

	for num > 0 {
		st.Push(num % 2)
		num = num / 2
	}

	var sb strings.Builder
	for !st.IsEmpty() {
		s, _ := st.Pop()
		sb.WriteString(strconv.Itoa(s))
	}

	result := sb.String()

	slog.Info("Resultado da conversão", "valor original", originalValue, "conversão", result)
}
