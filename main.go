package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pw_el116/elements"
	"pw_el116/global"
)

func main() {
	if err := Init(); err != nil {
		fmt.Println(err)
	}

	if err := elements.Init(); err != nil {
		fmt.Println(err)
	}
}

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"ZvxEMecywbttlYx9",
		"172.16.23.153",
		3306,
		"pw",
	)
	if global.MDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("无法连接到MySQL数据库。", err)
		return
	}
	//连接池
	sqlDB, _ := global.MDB.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(5)
	// 自动迁移
	_ = global.MDB.AutoMigrate(
		&elements.EquipmentAddon{},
		&elements.WeaponMajorType{},
		&elements.WeaponSubType{},
		&elements.WeaponEssence{},
		&elements.ArmorMajorType{},
		&elements.ArmorSubType{},
		&elements.ArmorEssence{},
		&elements.DecorationMajorType{},
		&elements.DecorationSubType{},
		&elements.DecorationEssence{},
		&elements.MedicineMajorType{},
		&elements.MedicineSubType{},
		&elements.MedicineEssence{},
		&elements.MaterialMajorType{},
		&elements.MaterialSubType{},
		&elements.MaterialEssence{},
		&elements.DamageruneSubType{},
		&elements.DamageruneEssence{},
		&elements.ArmorruneSubType{},
		&elements.ArmorruneEssence{},
		&elements.SkilltomeSubType{},
		&elements.SkilltomeEssence{},
		&elements.FlyswordEssence{},
		&elements.WingmanwingEssence{},
		&elements.TownscrollEssence{},
		&elements.UnionscrollEssence{},
		&elements.RevivescrollEssence{},
		&elements.ElementEssence{},
		&elements.TaskmatterEssence{},
		&elements.TossmatterEssence{},
		&elements.ProjectileType{},
		&elements.ProjectileEssence{},
		&elements.QuiverSubType{},
		&elements.QuiverEssence{},
		&elements.StoneSubType{},
		&elements.StoneEssence{},
		&elements.MonsterAddon{},
		&elements.MonsterType{},
		&elements.MonsterEssence{},
		&elements.NpcTalkService{},
		&elements.NpcSellService{},
		&elements.NpcBuyService{},
		&elements.NpcRepairService{},
		&elements.NpcInstallService{},
		&elements.NpcUninstallService{},
		&elements.NpcTaskInService{},
		&elements.NpcTaskOutService{},
		&elements.NpcTaskMatterService{},
		&elements.NpcSkillService{},
		&elements.NpcHealService{},
		&elements.NpcTransmitService{},
		&elements.NpcTransportService{},
		&elements.NpcProxyService{},
		&elements.NpcStorageService{},
		&elements.NpcMakeService{},
		&elements.NpcDecomposeService{},
		&elements.NpcType{},
		&elements.NpcEssence{},
		&elements.TalkProc{},
		&elements.FaceTextureEssence{},
		&elements.FaceShapeEssence{},
		&elements.FaceEmotionType{},
		&elements.FaceExpressionEssence{},
		&elements.FaceHairEssence{},
		&elements.FaceMoustacheEssence{},
		&elements.ColorpickerEssence{},
		&elements.CustomizedataEssence{},
		&elements.RecipeMajorType{},
		&elements.RecipeSubType{},
		&elements.RecipeEssence{},
		&elements.EnemyFactionConfig{},
		&elements.CharracterClassConfig{},
		&elements.ParamAdjustConfig{},
		&elements.PlayerActionInfoConfig{},
		&elements.TaskdiceEssence{},
		&elements.TasknormalmatterEssence{},
		&elements.FaceFalingEssence{},
		&elements.PlayerLevelexpConfig{},
		&elements.MineType{},
		&elements.MineEssence{},
		&elements.NpcIdentifyService{},
		&elements.FashionMajorType{},
		&elements.FashionSubType{},
		&elements.FashionEssence{},
		&elements.FaceticketMajorType{},
		&elements.FaceticketSubType{},
		&elements.FaceticketEssence{},
		&elements.FacepillMajorType{},
		&elements.FacepillSubType{},
		&elements.FacepillEssence{},
		&elements.SuiteEssence{},
		&elements.GmGeneratorType{},
		&elements.GmGeneratorEssence{},
		&elements.PetType{},
		&elements.PetEssence{},
		&elements.PetEggEssence{},
		&elements.PetFoodEssence{},
		&elements.PetFaceticketEssence{},
		&elements.FireworksEssence{},
		&elements.WarTankcallinEssence{},
		&elements.NpcWarTowerbuildService{},
		&elements.PlayerSecondlevelConfig{},
		&elements.NpcResetpropService{},
		&elements.NpcPetnameService{},
		&elements.NpcPetlearnskillService{},
		&elements.NpcPetforgetskillService{},
		&elements.SkillmatterEssence{},
		&elements.RefineTicketEssence{},
		&elements.DestroyingEssence{},
		&elements.NpcEquipbindService{},
		&elements.NpcEquipdestroyService{},
		&elements.NpcEquipundestroyService{},
		&elements.BibleEssence{},
		&elements.SpeakerEssence{},
		&elements.AutohpEssence{},
		&elements.AutompEssence{},
		&elements.DoubleExpEssence{},
	)
	return
}
