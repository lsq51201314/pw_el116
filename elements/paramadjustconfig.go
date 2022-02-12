package elements

import (
	"pw_el116/utils/structEx"
)

type ParamAdjustConfig struct {
	ID                             int     `elements:"name:id;type:int;size:4;text:未知;" gorm:"column:id;type:integer;autoIncrement:false;uniqueIndex;not null;" json:"id"`
	Name                           string  `elements:"name:name;type:unicode;size:64;text:未知;" gorm:"column:name;type:varchar(64);not null;" json:"name"`
	LevelDiffAdjust1LevelDiff      int     `elements:"name:level_diff_adjust_1_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_1_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_1_level_diff"`
	LevelDiffAdjust1AdjustExp      float32 `elements:"name:level_diff_adjust_1_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_1_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_1_adjust_exp"`
	LevelDiffAdjust1AdjustSp       float32 `elements:"name:level_diff_adjust_1_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_1_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_1_adjust_sp"`
	LevelDiffAdjust1AdjustMoney    float32 `elements:"name:level_diff_adjust_1_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_1_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_1_adjust_money"`
	LevelDiffAdjust1AdjustMatter   float32 `elements:"name:level_diff_adjust_1_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_1_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_1_adjust_matter"`
	LevelDiffAdjust1AdjustAttack   float32 `elements:"name:level_diff_adjust_1_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_1_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_1_adjust_attack"`
	LevelDiffAdjust2LevelDiff      int     `elements:"name:level_diff_adjust_2_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_2_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_2_level_diff"`
	LevelDiffAdjust2AdjustExp      float32 `elements:"name:level_diff_adjust_2_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_2_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_2_adjust_exp"`
	LevelDiffAdjust2AdjustSp       float32 `elements:"name:level_diff_adjust_2_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_2_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_2_adjust_sp"`
	LevelDiffAdjust2AdjustMoney    float32 `elements:"name:level_diff_adjust_2_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_2_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_2_adjust_money"`
	LevelDiffAdjust2AdjustMatter   float32 `elements:"name:level_diff_adjust_2_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_2_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_2_adjust_matter"`
	LevelDiffAdjust2AdjustAttack   float32 `elements:"name:level_diff_adjust_2_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_2_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_2_adjust_attack"`
	LevelDiffAdjust3LevelDiff      int     `elements:"name:level_diff_adjust_3_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_3_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_3_level_diff"`
	LevelDiffAdjust3AdjustExp      float32 `elements:"name:level_diff_adjust_3_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_3_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_3_adjust_exp"`
	LevelDiffAdjust3AdjustSp       float32 `elements:"name:level_diff_adjust_3_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_3_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_3_adjust_sp"`
	LevelDiffAdjust3AdjustMoney    float32 `elements:"name:level_diff_adjust_3_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_3_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_3_adjust_money"`
	LevelDiffAdjust3AdjustMatter   float32 `elements:"name:level_diff_adjust_3_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_3_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_3_adjust_matter"`
	LevelDiffAdjust3AdjustAttack   float32 `elements:"name:level_diff_adjust_3_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_3_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_3_adjust_attack"`
	LevelDiffAdjust4LevelDiff      int     `elements:"name:level_diff_adjust_4_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_4_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_4_level_diff"`
	LevelDiffAdjust4AdjustExp      float32 `elements:"name:level_diff_adjust_4_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_4_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_4_adjust_exp"`
	LevelDiffAdjust4AdjustSp       float32 `elements:"name:level_diff_adjust_4_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_4_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_4_adjust_sp"`
	LevelDiffAdjust4AdjustMoney    float32 `elements:"name:level_diff_adjust_4_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_4_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_4_adjust_money"`
	LevelDiffAdjust4AdjustMatter   float32 `elements:"name:level_diff_adjust_4_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_4_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_4_adjust_matter"`
	LevelDiffAdjust4AdjustAttack   float32 `elements:"name:level_diff_adjust_4_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_4_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_4_adjust_attack"`
	LevelDiffAdjust5LevelDiff      int     `elements:"name:level_diff_adjust_5_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_5_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_5_level_diff"`
	LevelDiffAdjust5AdjustExp      float32 `elements:"name:level_diff_adjust_5_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_5_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_5_adjust_exp"`
	LevelDiffAdjust5AdjustSp       float32 `elements:"name:level_diff_adjust_5_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_5_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_5_adjust_sp"`
	LevelDiffAdjust5AdjustMoney    float32 `elements:"name:level_diff_adjust_5_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_5_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_5_adjust_money"`
	LevelDiffAdjust5AdjustMatter   float32 `elements:"name:level_diff_adjust_5_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_5_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_5_adjust_matter"`
	LevelDiffAdjust5AdjustAttack   float32 `elements:"name:level_diff_adjust_5_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_5_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_5_adjust_attack"`
	LevelDiffAdjust6LevelDiff      int     `elements:"name:level_diff_adjust_6_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_6_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_6_level_diff"`
	LevelDiffAdjust6AdjustExp      float32 `elements:"name:level_diff_adjust_6_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_6_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_6_adjust_exp"`
	LevelDiffAdjust6AdjustSp       float32 `elements:"name:level_diff_adjust_6_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_6_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_6_adjust_sp"`
	LevelDiffAdjust6AdjustMoney    float32 `elements:"name:level_diff_adjust_6_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_6_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_6_adjust_money"`
	LevelDiffAdjust6AdjustMatter   float32 `elements:"name:level_diff_adjust_6_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_6_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_6_adjust_matter"`
	LevelDiffAdjust6AdjustAttack   float32 `elements:"name:level_diff_adjust_6_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_6_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_6_adjust_attack"`
	LevelDiffAdjust7LevelDiff      int     `elements:"name:level_diff_adjust_7_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_7_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_7_level_diff"`
	LevelDiffAdjust7AdjustExp      float32 `elements:"name:level_diff_adjust_7_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_7_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_7_adjust_exp"`
	LevelDiffAdjust7AdjustSp       float32 `elements:"name:level_diff_adjust_7_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_7_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_7_adjust_sp"`
	LevelDiffAdjust7AdjustMoney    float32 `elements:"name:level_diff_adjust_7_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_7_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_7_adjust_money"`
	LevelDiffAdjust7AdjustMatter   float32 `elements:"name:level_diff_adjust_7_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_7_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_7_adjust_matter"`
	LevelDiffAdjust7AdjustAttack   float32 `elements:"name:level_diff_adjust_7_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_7_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_7_adjust_attack"`
	LevelDiffAdjust8LevelDiff      int     `elements:"name:level_diff_adjust_8_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_8_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_8_level_diff"`
	LevelDiffAdjust8AdjustExp      float32 `elements:"name:level_diff_adjust_8_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_8_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_8_adjust_exp"`
	LevelDiffAdjust8AdjustSp       float32 `elements:"name:level_diff_adjust_8_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_8_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_8_adjust_sp"`
	LevelDiffAdjust8AdjustMoney    float32 `elements:"name:level_diff_adjust_8_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_8_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_8_adjust_money"`
	LevelDiffAdjust8AdjustMatter   float32 `elements:"name:level_diff_adjust_8_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_8_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_8_adjust_matter"`
	LevelDiffAdjust8AdjustAttack   float32 `elements:"name:level_diff_adjust_8_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_8_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_8_adjust_attack"`
	LevelDiffAdjust9LevelDiff      int     `elements:"name:level_diff_adjust_9_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_9_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_9_level_diff"`
	LevelDiffAdjust9AdjustExp      float32 `elements:"name:level_diff_adjust_9_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_9_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_9_adjust_exp"`
	LevelDiffAdjust9AdjustSp       float32 `elements:"name:level_diff_adjust_9_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_9_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_9_adjust_sp"`
	LevelDiffAdjust9AdjustMoney    float32 `elements:"name:level_diff_adjust_9_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_9_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_9_adjust_money"`
	LevelDiffAdjust9AdjustMatter   float32 `elements:"name:level_diff_adjust_9_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_9_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_9_adjust_matter"`
	LevelDiffAdjust9AdjustAttack   float32 `elements:"name:level_diff_adjust_9_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_9_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_9_adjust_attack"`
	LevelDiffAdjust10LevelDiff     int     `elements:"name:level_diff_adjust_10_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_10_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_10_level_diff"`
	LevelDiffAdjust10AdjustExp     float32 `elements:"name:level_diff_adjust_10_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_10_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_10_adjust_exp"`
	LevelDiffAdjust10AdjustSp      float32 `elements:"name:level_diff_adjust_10_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_10_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_10_adjust_sp"`
	LevelDiffAdjust10AdjustMoney   float32 `elements:"name:level_diff_adjust_10_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_10_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_10_adjust_money"`
	LevelDiffAdjust10AdjustMatter  float32 `elements:"name:level_diff_adjust_10_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_10_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_10_adjust_matter"`
	LevelDiffAdjust10AdjustAttack  float32 `elements:"name:level_diff_adjust_10_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_10_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_10_adjust_attack"`
	LevelDiffAdjust11LevelDiff     int     `elements:"name:level_diff_adjust_11_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_11_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_11_level_diff"`
	LevelDiffAdjust11AdjustExp     float32 `elements:"name:level_diff_adjust_11_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_11_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_11_adjust_exp"`
	LevelDiffAdjust11AdjustSp      float32 `elements:"name:level_diff_adjust_11_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_11_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_11_adjust_sp"`
	LevelDiffAdjust11AdjustMoney   float32 `elements:"name:level_diff_adjust_11_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_11_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_11_adjust_money"`
	LevelDiffAdjust11AdjustMatter  float32 `elements:"name:level_diff_adjust_11_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_11_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_11_adjust_matter"`
	LevelDiffAdjust11AdjustAttack  float32 `elements:"name:level_diff_adjust_11_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_11_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_11_adjust_attack"`
	LevelDiffAdjust12LevelDiff     int     `elements:"name:level_diff_adjust_12_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_12_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_12_level_diff"`
	LevelDiffAdjust12AdjustExp     float32 `elements:"name:level_diff_adjust_12_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_12_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_12_adjust_exp"`
	LevelDiffAdjust12AdjustSp      float32 `elements:"name:level_diff_adjust_12_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_12_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_12_adjust_sp"`
	LevelDiffAdjust12AdjustMoney   float32 `elements:"name:level_diff_adjust_12_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_12_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_12_adjust_money"`
	LevelDiffAdjust12AdjustMatter  float32 `elements:"name:level_diff_adjust_12_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_12_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_12_adjust_matter"`
	LevelDiffAdjust12AdjustAttack  float32 `elements:"name:level_diff_adjust_12_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_12_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_12_adjust_attack"`
	LevelDiffAdjust13LevelDiff     int     `elements:"name:level_diff_adjust_13_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_13_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_13_level_diff"`
	LevelDiffAdjust13AdjustExp     float32 `elements:"name:level_diff_adjust_13_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_13_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_13_adjust_exp"`
	LevelDiffAdjust13AdjustSp      float32 `elements:"name:level_diff_adjust_13_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_13_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_13_adjust_sp"`
	LevelDiffAdjust13AdjustMoney   float32 `elements:"name:level_diff_adjust_13_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_13_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_13_adjust_money"`
	LevelDiffAdjust13AdjustMatter  float32 `elements:"name:level_diff_adjust_13_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_13_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_13_adjust_matter"`
	LevelDiffAdjust13AdjustAttack  float32 `elements:"name:level_diff_adjust_13_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_13_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_13_adjust_attack"`
	LevelDiffAdjust14LevelDiff     int     `elements:"name:level_diff_adjust_14_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_14_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_14_level_diff"`
	LevelDiffAdjust14AdjustExp     float32 `elements:"name:level_diff_adjust_14_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_14_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_14_adjust_exp"`
	LevelDiffAdjust14AdjustSp      float32 `elements:"name:level_diff_adjust_14_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_14_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_14_adjust_sp"`
	LevelDiffAdjust14AdjustMoney   float32 `elements:"name:level_diff_adjust_14_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_14_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_14_adjust_money"`
	LevelDiffAdjust14AdjustMatter  float32 `elements:"name:level_diff_adjust_14_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_14_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_14_adjust_matter"`
	LevelDiffAdjust14AdjustAttack  float32 `elements:"name:level_diff_adjust_14_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_14_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_14_adjust_attack"`
	LevelDiffAdjust15LevelDiff     int     `elements:"name:level_diff_adjust_15_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_15_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_15_level_diff"`
	LevelDiffAdjust15AdjustExp     float32 `elements:"name:level_diff_adjust_15_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_15_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_15_adjust_exp"`
	LevelDiffAdjust15AdjustSp      float32 `elements:"name:level_diff_adjust_15_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_15_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_15_adjust_sp"`
	LevelDiffAdjust15AdjustMoney   float32 `elements:"name:level_diff_adjust_15_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_15_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_15_adjust_money"`
	LevelDiffAdjust15AdjustMatter  float32 `elements:"name:level_diff_adjust_15_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_15_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_15_adjust_matter"`
	LevelDiffAdjust15AdjustAttack  float32 `elements:"name:level_diff_adjust_15_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_15_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_15_adjust_attack"`
	LevelDiffAdjust16LevelDiff     int     `elements:"name:level_diff_adjust_16_level_diff;type:int;size:4;text:未知;" gorm:"column:level_diff_adjust_16_level_diff;type:integer;default:0;not null;" json:"level_diff_adjust_16_level_diff"`
	LevelDiffAdjust16AdjustExp     float32 `elements:"name:level_diff_adjust_16_adjust_exp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_16_adjust_exp;type:float;default:0;not null;" json:"level_diff_adjust_16_adjust_exp"`
	LevelDiffAdjust16AdjustSp      float32 `elements:"name:level_diff_adjust_16_adjust_sp;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_16_adjust_sp;type:float;default:0;not null;" json:"level_diff_adjust_16_adjust_sp"`
	LevelDiffAdjust16AdjustMoney   float32 `elements:"name:level_diff_adjust_16_adjust_money;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_16_adjust_money;type:float;default:0;not null;" json:"level_diff_adjust_16_adjust_money"`
	LevelDiffAdjust16AdjustMatter  float32 `elements:"name:level_diff_adjust_16_adjust_matter;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_16_adjust_matter;type:float;default:0;not null;" json:"level_diff_adjust_16_adjust_matter"`
	LevelDiffAdjust16AdjustAttack  float32 `elements:"name:level_diff_adjust_16_adjust_attack;type:float;size:4;text:未知;" gorm:"column:level_diff_adjust_16_adjust_attack;type:float;default:0;not null;" json:"level_diff_adjust_16_adjust_attack"`
	TeamAdjust1AdjustExp           float32 `elements:"name:team_adjust_1_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_1_adjust_exp;type:float;default:0;not null;" json:"team_adjust_1_adjust_exp"`
	TeamAdjust1AdjustSp            float32 `elements:"name:team_adjust_1_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_1_adjust_sp;type:float;default:0;not null;" json:"team_adjust_1_adjust_sp"`
	TeamAdjust2AdjustExp           float32 `elements:"name:team_adjust_2_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_2_adjust_exp;type:float;default:0;not null;" json:"team_adjust_2_adjust_exp"`
	TeamAdjust2AdjustSp            float32 `elements:"name:team_adjust_2_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_2_adjust_sp;type:float;default:0;not null;" json:"team_adjust_2_adjust_sp"`
	TeamAdjust3AdjustExp           float32 `elements:"name:team_adjust_3_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_3_adjust_exp;type:float;default:0;not null;" json:"team_adjust_3_adjust_exp"`
	TeamAdjust3AdjustSp            float32 `elements:"name:team_adjust_3_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_3_adjust_sp;type:float;default:0;not null;" json:"team_adjust_3_adjust_sp"`
	TeamAdjust4AdjustExp           float32 `elements:"name:team_adjust_4_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_4_adjust_exp;type:float;default:0;not null;" json:"team_adjust_4_adjust_exp"`
	TeamAdjust4AdjustSp            float32 `elements:"name:team_adjust_4_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_4_adjust_sp;type:float;default:0;not null;" json:"team_adjust_4_adjust_sp"`
	TeamAdjust5AdjustExp           float32 `elements:"name:team_adjust_5_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_5_adjust_exp;type:float;default:0;not null;" json:"team_adjust_5_adjust_exp"`
	TeamAdjust5AdjustSp            float32 `elements:"name:team_adjust_5_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_5_adjust_sp;type:float;default:0;not null;" json:"team_adjust_5_adjust_sp"`
	TeamAdjust6AdjustExp           float32 `elements:"name:team_adjust_6_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_6_adjust_exp;type:float;default:0;not null;" json:"team_adjust_6_adjust_exp"`
	TeamAdjust6AdjustSp            float32 `elements:"name:team_adjust_6_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_6_adjust_sp;type:float;default:0;not null;" json:"team_adjust_6_adjust_sp"`
	TeamAdjust7AdjustExp           float32 `elements:"name:team_adjust_7_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_7_adjust_exp;type:float;default:0;not null;" json:"team_adjust_7_adjust_exp"`
	TeamAdjust7AdjustSp            float32 `elements:"name:team_adjust_7_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_7_adjust_sp;type:float;default:0;not null;" json:"team_adjust_7_adjust_sp"`
	TeamAdjust8AdjustExp           float32 `elements:"name:team_adjust_8_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_8_adjust_exp;type:float;default:0;not null;" json:"team_adjust_8_adjust_exp"`
	TeamAdjust8AdjustSp            float32 `elements:"name:team_adjust_8_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_8_adjust_sp;type:float;default:0;not null;" json:"team_adjust_8_adjust_sp"`
	TeamAdjust9AdjustExp           float32 `elements:"name:team_adjust_9_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_9_adjust_exp;type:float;default:0;not null;" json:"team_adjust_9_adjust_exp"`
	TeamAdjust9AdjustSp            float32 `elements:"name:team_adjust_9_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_9_adjust_sp;type:float;default:0;not null;" json:"team_adjust_9_adjust_sp"`
	TeamAdjust10AdjustExp          float32 `elements:"name:team_adjust_10_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_10_adjust_exp;type:float;default:0;not null;" json:"team_adjust_10_adjust_exp"`
	TeamAdjust10AdjustSp           float32 `elements:"name:team_adjust_10_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_10_adjust_sp;type:float;default:0;not null;" json:"team_adjust_10_adjust_sp"`
	TeamAdjust11AdjustExp          float32 `elements:"name:team_adjust_11_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_adjust_11_adjust_exp;type:float;default:0;not null;" json:"team_adjust_11_adjust_exp"`
	TeamAdjust11AdjustSp           float32 `elements:"name:team_adjust_11_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_adjust_11_adjust_sp;type:float;default:0;not null;" json:"team_adjust_11_adjust_sp"`
	TeamProfessionAdjust1AdjustExp float32 `elements:"name:team_profession_adjust_1_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_1_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_1_adjust_exp"`
	TeamProfessionAdjust1AdjustSp  float32 `elements:"name:team_profession_adjust_1_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_1_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_1_adjust_sp"`
	TeamProfessionAdjust2AdjustExp float32 `elements:"name:team_profession_adjust_2_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_2_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_2_adjust_exp"`
	TeamProfessionAdjust2AdjustSp  float32 `elements:"name:team_profession_adjust_2_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_2_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_2_adjust_sp"`
	TeamProfessionAdjust3AdjustExp float32 `elements:"name:team_profession_adjust_3_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_3_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_3_adjust_exp"`
	TeamProfessionAdjust3AdjustSp  float32 `elements:"name:team_profession_adjust_3_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_3_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_3_adjust_sp"`
	TeamProfessionAdjust4AdjustExp float32 `elements:"name:team_profession_adjust_4_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_4_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_4_adjust_exp"`
	TeamProfessionAdjust4AdjustSp  float32 `elements:"name:team_profession_adjust_4_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_4_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_4_adjust_sp"`
	TeamProfessionAdjust5AdjustExp float32 `elements:"name:team_profession_adjust_5_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_5_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_5_adjust_exp"`
	TeamProfessionAdjust5AdjustSp  float32 `elements:"name:team_profession_adjust_5_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_5_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_5_adjust_sp"`
	TeamProfessionAdjust6AdjustExp float32 `elements:"name:team_profession_adjust_6_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_6_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_6_adjust_exp"`
	TeamProfessionAdjust6AdjustSp  float32 `elements:"name:team_profession_adjust_6_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_6_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_6_adjust_sp"`
	TeamProfessionAdjust7AdjustExp float32 `elements:"name:team_profession_adjust_7_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_7_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_7_adjust_exp"`
	TeamProfessionAdjust7AdjustSp  float32 `elements:"name:team_profession_adjust_7_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_7_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_7_adjust_sp"`
	TeamProfessionAdjust8AdjustExp float32 `elements:"name:team_profession_adjust_8_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_8_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_8_adjust_exp"`
	TeamProfessionAdjust8AdjustSp  float32 `elements:"name:team_profession_adjust_8_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_8_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_8_adjust_sp"`
	TeamProfessionAdjust9AdjustExp float32 `elements:"name:team_profession_adjust_9_adjust_exp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_9_adjust_exp;type:float;default:0;not null;" json:"team_profession_adjust_9_adjust_exp"`
	TeamProfessionAdjust9AdjustSp  float32 `elements:"name:team_profession_adjust_9_adjust_sp;type:float;size:4;text:未知;" gorm:"column:team_profession_adjust_9_adjust_sp;type:float;default:0;not null;" json:"team_profession_adjust_9_adjust_sp"`
}

func (ParamAdjustConfig) TableName() string {
	return "el_paramadjustconfig"
}

func (e *ParamAdjustConfig) GetSize() (size int, err error) {
	return structEx.GetSize(e)
}

func (e *ParamAdjustConfig) SetData(data []byte) (err error) {
	return structEx.SetData(e, data)
}
