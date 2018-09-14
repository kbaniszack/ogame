package ogame

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFusionReactorCapacity(t *testing.T) {
	fr := newFusionReactor()
	assert.Equal(t, 38*time.Second, fr.ConstructionTime(2, 7, Facilities{RoboticsFactory: 3}))
}

func TestFusionReactor_IsAvailable(t *testing.T) {
	fr := newFusionReactor()
	assert.False(t, fr.IsAvailable(ResourcesBuildings{}, Facilities{}, Researches{}, 0))
	assert.False(t, fr.IsAvailable(ResourcesBuildings{DeuteriumSynthesizer: 4}, Facilities{}, Researches{EnergyTechnology: 3}, 0))
	assert.False(t, fr.IsAvailable(ResourcesBuildings{DeuteriumSynthesizer: 5}, Facilities{}, Researches{EnergyTechnology: 2}, 0))
	assert.True(t, fr.IsAvailable(ResourcesBuildings{DeuteriumSynthesizer: 5}, Facilities{}, Researches{EnergyTechnology: 3}, 0))
	assert.True(t, fr.IsAvailable(ResourcesBuildings{DeuteriumSynthesizer: 6}, Facilities{}, Researches{EnergyTechnology: 4}, 0))
}
