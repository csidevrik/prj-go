package explorer

import (
	"fmt"
	"time"
)

type TabResult struct {
	TabNumber int
	Path      string
	Status    string        // "success", "failed"
	Error     string        // mensaje de error (vacío si success)
	Duration  time.Duration // tiempo que tomó abrir la pestaña
}

type SessionLog struct {
	TotalTabs   int
	SuccessTabs int
	FailedTabs  int
	Results     []TabResult
}

func NewSessionLog(totalTabs int) *SessionLog {
	return &SessionLog{
		TotalTabs: totalTabs,
		Results:   make([]TabResult, 0, totalTabs),
	}
}

func (sl *SessionLog) AddResult(result TabResult) {
	if result.Status == "success" {
		sl.SuccessTabs++
	} else {
		sl.FailedTabs++
	}
	sl.Results = append(sl.Results, result)
}

func (sl *SessionLog) Print() {
	fmt.Println("\n════════════════════════════════════════════")
	fmt.Printf("📊 Session Summary: %d/%d tabs opened successfully\n",
		sl.SuccessTabs, sl.TotalTabs)
	fmt.Println("════════════════════════════════════════════")

	for _, r := range sl.Results {
		status := "✅"
		if r.Status != "success" {
			status = "❌"
		}

		fmt.Printf("%s [Tab %d] %s\n", status, r.TabNumber, r.Path)
		if r.Error != "" {
			fmt.Printf("   └─ Error: %s\n", r.Error)
		}
		fmt.Printf("   └─ Duration: %dms\n", r.Duration.Milliseconds())
	}

	fmt.Println("════════════════════════════════════════════\n")
}
