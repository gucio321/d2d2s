package d2sskills

import (
	"testing"

	"github.com/gucio321/d2d2s/d2senums"
	"github.com/stretchr/testify/assert"
)

type testdata struct {
	class  d2senums.CharacterClass
	skills *Skills
}

func getTestdata() map[[NumSkillBytes]byte]testdata {
	header := []byte(skillsHeaderID)

	return map[[NumSkillBytes]byte]testdata{
		{
			header[0], header[1],
			8, 16, 2, 4, 15, 7, 8, 5, 4, 20, 18, 4, 3, 9, 13, 8, 1, 0,
			19, 8, 2, 0, 4, 3, 6, 1, 16, 15, 14, 7,
		}: {
			d2senums.CharacterClassNecromancer,
			&Skills{
				d2senums.SkillNecromancerAmplifyDamage:    8,
				d2senums.SkillNecromancerTeeth:            16,
				d2senums.SkillNecromancerBoneArmor:        2,
				d2senums.SkillNecromancerSkeletonMastery:  4,
				d2senums.SkillNecromancerRaiseSkeleton:    15,
				d2senums.SkillNecromancerDimVision:        7,
				d2senums.SkillNecromancerWeaken:           8,
				d2senums.SkillNecromancerPoisonDagger:     5,
				d2senums.SkillNecromancerCorpseExplosion:  4,
				d2senums.SkillNecromancerClyGolem:         20,
				d2senums.SkillNecromancerIronMaiden:       18,
				d2senums.SkillNecromancerTerror:           4,
				d2senums.SkillNecromancerBoneWall:         3,
				d2senums.SkillNecromancerGolemMastery:     9,
				d2senums.SkillNecromancerRaiseSkeletalMag: 13,
				d2senums.SkillNecromancerConfuse:          8,
				d2senums.SkillNecromancerLifeTap:          1,
				d2senums.SkillNecromancerPoisonExplosion:  0,
				d2senums.SkillNecromancerBoneSpear:        19,
				d2senums.SkillNecromancerBloodGolem:       8,
				d2senums.SkillNecromancerAttract:          2,
				d2senums.SkillNecromancerDecrepify:        0,
				d2senums.SkillNecromancerBonePrison:       4,
				d2senums.SkillNecromancerSummonResist:     3,
				d2senums.SkillNecromancerIronGolem:        6,
				d2senums.SkillNecromancerLowerResist:      1,
				d2senums.SkillNecromancerPoisonNova:       16,
				d2senums.SkillNecromancerBoneSpirit:       15,
				d2senums.SkillNecromancerFireGolem:        14,
				d2senums.SkillNecromancerRevive:           7,
			},
		},
	}
}

func Test_Load(t *testing.T) {
	td := getTestdata()

	for key, value := range td {
		s := New()

		err := s.Load(key, value.class)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, value.skills, s, "unexpected skills loaded")
	}
}

func Test_Encode(t *testing.T) {
	td := getTestdata()

	for key, value := range td {
		data := value.skills.Encode(value.class)

		assert.Equal(t, key, data, "unexpected skill data encoded")
	}
}