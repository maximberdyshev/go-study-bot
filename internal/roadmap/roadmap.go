package roadmap

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Roadmap struct {
	Days      []Day
	TotalDays int
}

type Day struct {
	Number      int    `yaml:"number"`
	Week        int    `yaml:"week"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

func LoadFromFile(path string) (*Roadmap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed read file %s: %w", path, err)
	}

	var roadmap struct {
		Days      []Day `yaml:"days"`
		Totaldays int   `yaml:"total_days"`
	}
	if err := yaml.Unmarshal(data, &roadmap); err != nil {
		return nil, fmt.Errorf("failed parse YAML: %w", err)
	}

	if roadmap.Totaldays != len(roadmap.Days) {
		return nil, fmt.Errorf("mismatch: total_days=%d != has %d days", roadmap.Totaldays, len(roadmap.Days))
	}

	return &Roadmap{
		Days:      roadmap.Days,
		TotalDays: roadmap.Totaldays,
	}, nil
}

func (r *Roadmap) GetByNumber(dayNum int) *Day {
	for _, d := range r.Days {
		if d.Number == dayNum {
			return &d
		}
	}
	return nil
}
