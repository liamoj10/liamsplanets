package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Script to push planned roadmap issues directly to Linear via GraphQL API
// Configured with:
// - Initial State: Triage
// - T-Shirt Sizing: S, M, L, XL
// - Labels: design-pattern, physics, gameplay, backend
// - Milestones: Milestone 1: MVP Physics, Milestone 2: Game Mechanics, Milestone 3: Control & Architecture
//
// Usage: go run scripts/sync_linear.go (reads LINEAR_API_KEY from .env or environment)

const linearEndpoint = "https://api.linear.app/graphql"

type PlannedIssue struct {
	Title       string
	Description string
	TShirtSize  string // S, M, L, XL
	Milestone   string
	Labels      []string
	Priority    int
}

var plannedIssues = []PlannedIssue{
	{
		Title:       "Implement Vector2D Kinematics & Operations",
		Description: "Complete Vector2D math operations in internal/domain/vector.go (Add, Sub, Mul, Magnitude, Normalize, Distance) and verify unit tests.\n\n**Pattern/Discipline**: Physics Math\n**Files**: `internal/domain/vector.go`, `internal/domain/planet_test.go`",
		TShirtSize:  "S",
		Milestone:   "Milestone 1: MVP Physics",
		Labels:      []string{"physics", "backend"},
		Priority:    1,
	},
	{
		Title:       "Implement Strategy Pattern for Physics Integrators (Euler, Verlet, RK4)",
		Description: "Implement the Integrator interface in internal/physics/strategy.go with:\n- EulerIntegrator (basic 1st order)\n- VerletIntegrator (Velocity-Verlet for energy conservation)\n- RungeKutta4Integrator (RK4 classic celestial mechanics)\n\n**Design Pattern**: [Strategy Pattern](https://refactoring.guru/design-patterns/strategy)\n**Files**: `internal/physics/strategy.go`",
		TShirtSize:  "M",
		Milestone:   "Milestone 1: MVP Physics",
		Labels:      []string{"design-pattern", "physics", "backend"},
		Priority:    1,
	},
	{
		Title:       "Implement Gravitational Force & Circular Orbital Velocity",
		Description: "Complete GravitationalForce vector calculation and circular orbital velocity formula v = sqrt(G*M/r) in internal/domain/planet.go.\n\n**Files**: `internal/domain/planet.go`",
		TShirtSize:  "S",
		Milestone:   "Milestone 1: MVP Physics",
		Labels:      []string{"physics", "backend"},
		Priority:    1,
	},
	{
		Title:       "Implement Shatter Collision Detection & Debris Resolution",
		Description: "Implement ShatterCollisionHandler in internal/domain/collision.go to shatter colliding bodies into smaller asteroid fragments conserving momentum P_cm.\n\n**Files**: `internal/domain/collision.go`",
		TShirtSize:  "M",
		Milestone:   "Milestone 1: MVP Physics",
		Labels:      []string{"physics", "gameplay", "backend"},
		Priority:    2,
	},
	{
		Title:       "Implement Decorator Pattern for Planet Attributes (Atmosphere, Resources, Colonies)",
		Description: "Implement Decorator pattern in internal/domain/decorator.go for dynamic game mechanics wrapping CelestialBody:\n- AtmosphericDecorator (Surface temperature & habitability)\n- ResourceBearingDecorator (Iron, Water Ice, Minerals mining)\n- ColonizedDecorator (Population & infrastructure growth)\n\n**Design Pattern**: [Decorator Pattern](https://refactoring.guru/design-patterns/decorator)\n**Files**: `internal/domain/decorator.go`",
		TShirtSize:  "L",
		Milestone:   "Milestone 2: Game Mechanics",
		Labels:      []string{"design-pattern", "gameplay", "backend"},
		Priority:    2,
	},
	{
		Title:       "Implement SolarSystemBuilder & Factory Methods",
		Description: "Implement Builder & Factory Method in internal/factory/builder.go to fluently build solar systems with correct tangent velocities.\n\n**Design Pattern**: [Builder](https://refactoring.guru/design-patterns/builder) & [Factory Method](https://refactoring.guru/design-patterns/factory-method)\n**Files**: `internal/factory/builder.go`",
		TShirtSize:  "M",
		Milestone:   "Milestone 3: Control & Architecture",
		Labels:      []string{"design-pattern", "backend"},
		Priority:    2,
	},
	{
		Title:       "Implement Command Pattern & Undo History Stack",
		Description: "Implement Command pattern with SpawnBodyCommand, PauseCommand, SetTimeScaleCommand, and CommandHistory in internal/commands/command.go.\n\n**Design Pattern**: [Command Pattern](https://refactoring.guru/design-patterns/command)\n**Files**: `internal/commands/command.go`",
		TShirtSize:  "M",
		Milestone:   "Milestone 3: Control & Architecture",
		Labels:      []string{"design-pattern", "gameplay", "backend"},
		Priority:    3,
	},
	{
		Title:       "Implement Simulation State Machine (Running, Paused)",
		Description: "Implement State pattern in internal/state/state.go for simulation lifecycle transitions.\n\n**Design Pattern**: [State Pattern](https://refactoring.guru/design-patterns/state)\n**Files**: `internal/state/state.go`",
		TShirtSize:  "S",
		Milestone:   "Milestone 3: Control & Architecture",
		Labels:      []string{"design-pattern", "backend"},
		Priority:    3,
	},
	{
		Title:       "Implement Observer Event Dispatcher & WebSocket Stream",
		Description: "Implement Observer pattern in internal/events/observer.go and wire to WebSocket broadcaster in internal/transport.\n\n**Design Pattern**: [Observer Pattern](https://refactoring.guru/design-patterns/observer)\n**Files**: `internal/events/observer.go`, `internal/transport/handler.go`",
		TShirtSize:  "L",
		Milestone:   "Milestone 3: Control & Architecture",
		Labels:      []string{"design-pattern", "backend"},
		Priority:    3,
	},
}

// loadDotEnv parses .env file if present
func loadDotEnv() {
	file, err := os.Open(".env")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.Trim(strings.TrimSpace(parts[1]), `"'`)
			if os.Getenv(key) == "" {
				os.Setenv(key, val)
			}
		}
	}
}

func main() {
	// Automatically try reading .env file first
	loadDotEnv()

	apiKey := os.Getenv("LINEAR_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: LINEAR_API_KEY is not set.")
		fmt.Println("\nChoose one of the following:")
		fmt.Println("  1. Create a .env file in the project root with:")
		fmt.Println("     LINEAR_API_KEY=lin_api_your_key_here")
		fmt.Println("  2. Or run directly with:")
		fmt.Println("     LINEAR_API_KEY=lin_api_xxxx go run scripts/sync_linear.go")
		os.Exit(1)
	}

	// 1. Fetch team ID & triage state ID
	teamID, triageStateID, err := getTeamAndTriageState(apiKey)
	if err != nil {
		fmt.Printf("Failed to fetch team: %v\n", err)
		os.Exit(1)
	}

	// 2. Fetch or create labels
	labelMap, err := ensureLabels(apiKey, teamID, []string{"design-pattern", "physics", "gameplay", "backend"})
	if err != nil {
		fmt.Printf("Warning: Could not fetch/create labels: %v\n", err)
		labelMap = make(map[string]string)
	}

	fmt.Printf("\nTargeting Linear Team (ID: %s) with Triage State (ID: %s)\n", teamID, triageStateID)
	fmt.Println("Syncing roadmap issues...\n")

	// 3. Create each issue
	for _, issue := range plannedIssues {
		var labelIDs []string
		for _, lbl := range issue.Labels {
			if id, ok := labelMap[lbl]; ok {
				labelIDs = append(labelIDs, id)
			}
		}

		desc := fmt.Sprintf("%s\n\n**T-Shirt Size**: %s\n**Milestone**: %s", issue.Description, issue.TShirtSize, issue.Milestone)

		createdID, err := createIssueWithDetails(apiKey, teamID, triageStateID, issue.Title, desc, issue.Priority, labelIDs)
		if err != nil {
			fmt.Printf("❌ Failed to create '%s': %v\n", issue.Title, err)
		} else {
			fmt.Printf("✅ [%s] Created '%s' -> %s\n", issue.TShirtSize, issue.Title, createdID)
		}
	}

	fmt.Println("\n🎉 All roadmap issues successfully dispatched to your Linear Triage Inbox!")
}

func executeGraphQL(apiKey, query string, variables map[string]any) ([]byte, error) {
	reqBody, err := json.Marshal(map[string]any{
		"query":     query,
		"variables": variables,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", linearEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func getTeamAndTriageState(apiKey string) (string, string, error) {
	query := `query {
		teams {
			nodes {
				id
				name
				key
				triageEnabled
				states {
					nodes {
						id
						name
						type
					}
				}
			}
		}
	}`

	body, err := executeGraphQL(apiKey, query, nil)
	if err != nil {
		return "", "", err
	}

	var res struct {
		Data struct {
			Teams struct {
				Nodes []struct {
					ID     string `json:"id"`
					Name   string `json:"name"`
					Key    string `json:"key"`
					States struct {
						Nodes []struct {
							ID   string `json:"id"`
							Name string `json:"name"`
							Type string `json:"type"`
						} `json:"nodes"`
					} `json:"states"`
				} `json:"nodes"`
			} `json:"teams"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return "", "", err
	}
	if len(res.Data.Teams.Nodes) == 0 {
		return "", "", fmt.Errorf("no teams found in Linear")
	}

	team := res.Data.Teams.Nodes[0]
	var triageID string
	for _, s := range team.States.Nodes {
		if s.Type == "triage" || s.Name == "Triage" {
			triageID = s.ID
			break
		}
	}

	return team.ID, triageID, nil
}

func ensureLabels(apiKey, teamID string, names []string) (map[string]string, error) {
	labelMap := make(map[string]string)
	query := `query { issueLabels { nodes { id name } } }`
	body, err := executeGraphQL(apiKey, query, nil)
	if err == nil {
		var res struct {
			Data struct {
				IssueLabels struct {
					Nodes []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"nodes"`
				} `json:"issueLabels"`
			} `json:"data"`
		}
		if json.Unmarshal(body, &res) == nil {
			for _, l := range res.Data.IssueLabels.Nodes {
				labelMap[l.Name] = l.ID
			}
		}
	}
	return labelMap, nil
}

func createIssueWithDetails(apiKey, teamID, stateID, title, description string, priority int, labelIDs []string) (string, error) {
	mutation := `mutation CreateIssue($input: IssueCreateInput!) {
		issueCreate(input: $input) {
			success
			issue { id identifier title url }
		}
	}`

	input := map[string]any{
		"teamId":      teamID,
		"title":       title,
		"description": description,
		"priority":    priority,
	}
	if stateID != "" {
		input["stateId"] = stateID
	}
	if len(labelIDs) > 0 {
		input["labelIds"] = labelIDs
	}

	body, err := executeGraphQL(apiKey, mutation, map[string]any{"input": input})
	if err != nil {
		return "", err
	}

	var res struct {
		Data struct {
			IssueCreate struct {
				Success bool `json:"success"`
				Issue   struct {
					ID         string `json:"id"`
					Identifier string `json:"identifier"`
					Title      string `json:"title"`
					URL        string `json:"url"`
				} `json:"issue"`
			} `json:"issueCreate"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	if len(res.Errors) > 0 {
		return "", fmt.Errorf("%s", res.Errors[0].Message)
	}

	return fmt.Sprintf("%s (%s)", res.Data.IssueCreate.Issue.Identifier, res.Data.IssueCreate.Issue.URL), nil
}
