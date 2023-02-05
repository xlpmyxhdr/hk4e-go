package game

import (
	"strconv"

	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/object"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

func (g *GameManager) GetAllAvatarDataConfig() map[int32]*gdconf.AvatarData {
	allAvatarDataConfig := make(map[int32]*gdconf.AvatarData)
	for avatarId, avatarData := range gdconf.CONF.AvatarDataMap {
		if avatarId < 10000002 || avatarId >= 11000000 {
			// 跳过无效角色
			continue
		}
		if avatarId == 10000005 || avatarId == 10000007 {
			// 跳过主角
			continue
		}
		if avatarId >= 10000079 {
			// 跳过后续版本的角色
			continue
		}
		allAvatarDataConfig[avatarId] = avatarData
	}
	return allAvatarDataConfig
}

func (g *GameManager) AddUserAvatar(userId uint32, avatarId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	// 判断玩家是否已有该角色
	_, ok := player.AvatarMap[avatarId]
	if ok {
		// TODO 如果已有转换命座材料
		return
	}
	player.AddAvatar(avatarId)

	// 添加初始武器
	avatarDataConfig, ok := gdconf.CONF.AvatarDataMap[int32(avatarId)]
	if !ok {
		logger.Error("config is nil, itemId: %v", avatarId)
		return
	}
	weaponId := g.AddUserWeapon(player.PlayerID, uint32(avatarDataConfig.InitialWeapon))

	// 角色装上初始武器
	g.WearUserAvatarEquip(player.PlayerID, avatarId, weaponId)

	g.UpdateUserAvatarFightProp(player.PlayerID, avatarId)

	avatarAddNotify := &proto.AvatarAddNotify{
		Avatar:   g.PacketAvatarInfo(player.AvatarMap[avatarId]),
		IsInTeam: false,
	}
	g.SendMsg(cmd.AvatarAddNotify, userId, player.ClientSeq, avatarAddNotify)
}

// AvatarPromoteGetRewardReq 角色突破获取奖励请求
func (g *GameManager) AvatarPromoteGetRewardReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user promote get reward, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.AvatarPromoteGetRewardReq)
	// 是否拥有角色
	avatar, ok := player.AvatarMap[player.GetAvatarIdByGuid(req.AvatarGuid)]
	if !ok {
		logger.Error("avatar error, avatarGuid: %v", req.AvatarGuid)
		g.CommonRetError(cmd.AvatarPromoteGetRewardRsp, player, &proto.AvatarPromoteGetRewardRsp{}, proto.Retcode_RET_CAN_NOT_FIND_AVATAR)
		return
	}
	// 获取角色配置表
	avatarDataConfig, ok := gdconf.CONF.AvatarDataMap[int32(avatar.AvatarId)]
	if !ok {
		logger.Error("avatar config error, avatarId: %v", avatar.AvatarId)
		g.CommonRetError(cmd.AvatarPromoteGetRewardRsp, player, &proto.AvatarPromoteGetRewardRsp{})
		return
	}
	// 角色是否获取过该突破等级的奖励
	if avatar.PromoteRewardMap[req.PromoteLevel] {
		logger.Error("avatar config error, avatarId: %v", avatar.AvatarId)
		g.CommonRetError(cmd.AvatarPromoteGetRewardRsp, player, &proto.AvatarPromoteGetRewardRsp{}, proto.Retcode_RET_REWARD_HAS_TAKEN)
		return
	}
	// 获取奖励配置表
	rewardConfig, ok := gdconf.CONF.RewardDataMap[int32(avatarDataConfig.PromoteRewardMap[req.PromoteLevel])]
	if !ok {
		logger.Error("reward config error, rewardId: %v", avatarDataConfig.PromoteRewardMap[req.PromoteLevel])
		g.CommonRetError(cmd.AvatarPromoteGetRewardRsp, player, &proto.AvatarPromoteGetRewardRsp{})
		return
	}
	// 设置该奖励为已被获取状态
	avatar.PromoteRewardMap[req.PromoteLevel] = true
	// 给予突破奖励
	rewardItemList := make([]*UserItem, 0, len(rewardConfig.RewardItemMap))
	for itemId, count := range rewardConfig.RewardItemMap {
		rewardItemList = append(rewardItemList, &UserItem{
			ItemId:      itemId,
			ChangeCount: count,
		})
	}
	g.AddUserItem(player.PlayerID, rewardItemList, false, 0)

	avatarPromoteGetRewardRsp := &proto.AvatarPromoteGetRewardRsp{
		RewardId:     uint32(rewardConfig.RewardID),
		AvatarGuid:   req.AvatarGuid,
		PromoteLevel: req.PromoteLevel,
	}
	g.SendMsg(cmd.AvatarPromoteGetRewardRsp, player.PlayerID, player.ClientSeq, avatarPromoteGetRewardRsp)
}

// AvatarPromoteReq 角色突破请求
func (g *GameManager) AvatarPromoteReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user promote, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.AvatarPromoteReq)
	// 是否拥有角色
	avatar, ok := player.AvatarMap[player.GetAvatarIdByGuid(req.Guid)]
	if !ok {
		logger.Error("avatar error, avatarGuid: %v", req.Guid)
		g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{}, proto.Retcode_RET_CAN_NOT_FIND_AVATAR)
		return
	}
	// 获取角色配置表
	avatarDataConfig, ok := gdconf.CONF.AvatarDataMap[int32(avatar.AvatarId)]
	if !ok {
		logger.Error("avatar config error, avatarId: %v", avatar.AvatarId)
		g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{})
		return
	}
	// 获取角色突破配置表
	avatarPromoteDataMap, ok := gdconf.CONF.AvatarPromoteDataMap[avatarDataConfig.PromoteId]
	if !ok {
		logger.Error("avatar promote config error, promoteId: %v", avatarDataConfig.PromoteId)
		g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{})
		return
	}
	// 获取角色突破等级的配置表
	avatarPromoteConfig, ok := avatarPromoteDataMap[int32(avatar.Promote)]
	if !ok {
		logger.Error("avatar promote config error, promoteLevel: %v", avatar.Promote)
		g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{})
		return
	}
	// 角色等级是否达到限制
	if avatar.Level < uint8(avatarPromoteConfig.LevelLimit) {
		logger.Error("avatar level le level limit, level: %v", avatar.Level)
		g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{}, proto.Retcode_RET_AVATAR_LEVEL_LESS_THAN)
		return
	}
	// 获取角色突破下一级的配置表
	avatarPromoteConfig, ok = avatarPromoteDataMap[int32(avatar.Promote+1)]
	if !ok {
		logger.Error("avatar promote config error, next promoteLevel: %v", avatar.Promote+1)
		g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{}, proto.Retcode_RET_AVATAR_ON_MAX_BREAK_LEVEL)
		return
	}
	// 将被消耗的物品列表
	costItemList := make([]*UserItem, 0, len(avatarPromoteConfig.CostItemMap)+1)
	// 突破材料是否足够并添加到消耗物品列表
	for itemId, count := range avatarPromoteConfig.CostItemMap {
		costItemList = append(costItemList, &UserItem{
			ItemId:      itemId,
			ChangeCount: count,
		})
	}
	// 消耗列表添加摩拉的消耗
	costItemList = append(costItemList, &UserItem{
		ItemId:      constant.ItemConstantConst.SCOIN,
		ChangeCount: uint32(avatarPromoteConfig.CostCoin),
	})
	// 突破材料以及摩拉是否足够
	for _, item := range costItemList {
		if player.GetItemCount(item.ItemId) < item.ChangeCount {
			logger.Error("item count not enough, itemId: %v", item.ItemId)
			// 摩拉的错误提示与材料不同
			if item.ItemId == constant.ItemConstantConst.SCOIN {
				g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{}, proto.Retcode_RET_SCOIN_NOT_ENOUGH)
			}
			g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{}, proto.Retcode_RET_ITEM_COUNT_NOT_ENOUGH)
			return
		}
	}
	// 冒险等级是否符合要求
	if player.PropertiesMap[constant.PlayerPropertyConst.PROP_PLAYER_LEVEL] < uint32(avatarPromoteConfig.MinPlayerLevel) {
		logger.Error("player level not enough, level: %v", player.PropertiesMap[constant.PlayerPropertyConst.PROP_PLAYER_LEVEL])
		g.CommonRetError(cmd.AvatarPromoteRsp, player, &proto.AvatarPromoteRsp{}, proto.Retcode_RET_PLAYER_LEVEL_LESS_THAN)
		return
	}
	// 消耗突破材料和摩拉
	g.CostUserItem(player.PlayerID, costItemList)

	// 角色突破等级+1
	avatar.Promote++
	// 角色更新面板
	g.UpdateUserAvatarFightProp(player.PlayerID, avatar.AvatarId)
	// 角色属性表更新通知
	g.SendMsg(cmd.AvatarPropNotify, player.PlayerID, player.ClientSeq, g.PacketAvatarPropNotify(avatar))

	avatarPromoteRsp := &proto.AvatarPromoteRsp{
		Guid: req.Guid,
	}
	g.SendMsg(cmd.AvatarPromoteRsp, player.PlayerID, player.ClientSeq, avatarPromoteRsp)
}

// AvatarUpgradeReq 角色升级请求
func (g *GameManager) AvatarUpgradeReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user upgrade, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.AvatarUpgradeReq)
	// 是否拥有角色
	avatar, ok := player.AvatarMap[player.GetAvatarIdByGuid(req.AvatarGuid)]
	if !ok {
		logger.Error("avatar error, avatarGuid: %v", req.AvatarGuid)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{}, proto.Retcode_RET_CAN_NOT_FIND_AVATAR)
		return
	}
	// 经验书数量是否足够
	if player.GetItemCount(req.ItemId) < req.Count {
		logger.Error("item count not enough, itemId: %v", req.ItemId)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{}, proto.Retcode_RET_ITEM_COUNT_NOT_ENOUGH)
		return
	}
	// 获取经验书物品配置表
	itemDataConfig, ok := gdconf.CONF.ItemDataMap[int32(req.ItemId)]
	if !ok {
		logger.Error("item data config error, itemId: %v", constant.ItemConstantConst.SCOIN)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{}, proto.Retcode_RET_ITEM_NOT_EXIST)
		return
	}
	// 经验书将给予的经验数
	itemParam, err := strconv.Atoi(itemDataConfig.Use1Param1)
	if err != nil {
		logger.Error("parse item param error: %v", err)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{})
		return
	}
	// 角色获得的经验
	expCount := uint32(itemParam) * req.Count
	// 摩拉数量是否足够
	if player.GetItemCount(constant.ItemConstantConst.SCOIN) < expCount/5 {
		logger.Error("item count not enough, itemId: %v", constant.ItemConstantConst.SCOIN)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{}, proto.Retcode_RET_SCOIN_NOT_ENOUGH)
		return
	}
	// 获取角色配置表
	avatarDataConfig, ok := gdconf.CONF.AvatarDataMap[int32(avatar.AvatarId)]
	if !ok {
		logger.Error("avatar config error, avatarId: %v", avatar.AvatarId)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{})
		return
	}
	// 获取角色突破配置表
	avatarPromoteDataMap, ok := gdconf.CONF.AvatarPromoteDataMap[avatarDataConfig.PromoteId]
	if !ok {
		logger.Error("avatar promote config error, promoteId: %v", avatarDataConfig.PromoteId)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{})
		return
	}
	// 获取角色突破等级对应的配置表
	avatarPromoteConfig, ok := avatarPromoteDataMap[int32(avatar.Promote)]
	if !ok {
		logger.Error("avatar promote config error, promoteLevel: %v", avatar.Promote)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{})
		return
	}
	// 角色等级是否达到限制
	if avatar.Level >= uint8(avatarPromoteConfig.LevelLimit) {
		logger.Error("avatar level ge level limit, level: %v", avatar.Level)
		g.CommonRetError(cmd.AvatarUpgradeRsp, player, &proto.AvatarUpgradeRsp{}, proto.Retcode_RET_AVATAR_BREAK_LEVEL_LESS_THAN)
		return
	}
	// 消耗升级材料以及摩拉
	g.CostUserItem(player.PlayerID, []*UserItem{
		{
			ItemId:      req.ItemId,
			ChangeCount: req.Count,
		},
		{
			ItemId:      constant.ItemConstantConst.SCOIN,
			ChangeCount: expCount / 5,
		},
	})
	// 角色升级前的信息
	oldLevel := avatar.Level
	oldFightPropMap := make(map[uint32]float32, len(avatar.FightPropMap))
	for propType, propValue := range avatar.FightPropMap {
		oldFightPropMap[propType] = propValue
	}

	// 角色添加经验
	g.UpgradePlayerAvatar(player, avatar, expCount)

	avatarUpgradeRsp := &proto.AvatarUpgradeRsp{
		CurLevel:        uint32(avatar.Level),
		OldLevel:        uint32(oldLevel),
		OldFightPropMap: oldFightPropMap,
		CurFightPropMap: avatar.FightPropMap,
		AvatarGuid:      req.AvatarGuid,
	}
	g.SendMsg(cmd.AvatarUpgradeRsp, player.PlayerID, player.ClientSeq, avatarUpgradeRsp)
}

// UpgradePlayerAvatar 玩家角色升级
func (g *GameManager) UpgradePlayerAvatar(player *model.Player, avatar *model.Avatar, expCount uint32) {
	// 获取角色配置表
	avatarDataConfig, ok := gdconf.CONF.AvatarDataMap[int32(avatar.AvatarId)]
	if !ok {
		logger.Error("avatar config error, avatarId: %v", avatar.AvatarId)
		return
	}
	// 获取角色突破配置表
	avatarPromoteDataMap, ok := gdconf.CONF.AvatarPromoteDataMap[avatarDataConfig.PromoteId]
	if !ok {
		logger.Error("avatar promote config error, promoteId: %v", avatarDataConfig.PromoteId)
		return
	}
	// 获取角色突破等级对应的配置表
	avatarPromoteConfig, ok := avatarPromoteDataMap[int32(avatar.Promote)]
	if !ok {
		logger.Error("avatar promote config error, promoteLevel: %v", avatar.Promote)
		return
	}
	// 角色增加经验
	avatar.Exp += expCount
	// 角色升级
	for {
		// 获取角色等级配置表
		avatarLevelConfig, ok := gdconf.CONF.AvatarLevelDataMap[int32(avatar.Level)]
		if !ok {
			// 获取不到代表已经到达最大等级
			break
		}
		// 角色当前等级未突破则跳出循环
		if avatar.Level >= uint8(avatarPromoteConfig.LevelLimit) {
			// 角色未突破溢出的经验处理
			avatar.Exp = 0
			break
		}
		// 角色经验小于升级所需的经验则跳出循环
		if avatar.Exp < uint32(avatarLevelConfig.Exp) {
			break
		}
		// 角色等级提升
		avatar.Exp -= uint32(avatarLevelConfig.Exp)
		avatar.Level++
	}
	// 角色更新面板
	g.UpdateUserAvatarFightProp(player.PlayerID, avatar.AvatarId)
	// 角色属性表更新通知
	g.SendMsg(cmd.AvatarPropNotify, player.PlayerID, player.ClientSeq, g.PacketAvatarPropNotify(avatar))
}

// PacketAvatarPropNotify 角色属性表更新通知
func (g *GameManager) PacketAvatarPropNotify(avatar *model.Avatar) *proto.AvatarPropNotify {
	avatarPropNotify := &proto.AvatarPropNotify{
		PropMap:    make(map[uint32]int64, 5),
		AvatarGuid: avatar.Guid,
	}
	// 角色等级
	avatarPropNotify.PropMap[uint32(constant.PlayerPropertyConst.PROP_LEVEL)] = int64(avatar.Level)
	// 角色经验
	avatarPropNotify.PropMap[uint32(constant.PlayerPropertyConst.PROP_EXP)] = int64(avatar.Exp)
	// 角色突破等级
	avatarPropNotify.PropMap[uint32(constant.PlayerPropertyConst.PROP_BREAK_LEVEL)] = int64(avatar.Promote)
	// 角色饱食度
	avatarPropNotify.PropMap[uint32(constant.PlayerPropertyConst.PROP_SATIATION_VAL)] = int64(avatar.Satiation)
	// 角色饱食度溢出
	avatarPropNotify.PropMap[uint32(constant.PlayerPropertyConst.PROP_SATIATION_PENALTY_TIME)] = int64(avatar.SatiationPenalty)

	return avatarPropNotify
}

func (g *GameManager) WearEquipReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user wear equip, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.WearEquipReq)
	avatarGuid := req.AvatarGuid
	equipGuid := req.EquipGuid
	avatar, ok := player.GameObjectGuidMap[avatarGuid].(*model.Avatar)
	if !ok {
		logger.Error("avatar error, avatarGuid: %v", avatarGuid)
		g.CommonRetError(cmd.WearEquipRsp, player, &proto.WearEquipRsp{}, proto.Retcode_RET_CAN_NOT_FIND_AVATAR)
		return
	}
	weapon, ok := player.GameObjectGuidMap[equipGuid].(*model.Weapon)
	if !ok {
		logger.Error("equip error, equipGuid: %v", equipGuid)
		g.CommonRetError(cmd.WearEquipRsp, player, &proto.WearEquipRsp{})
		return
	}
	g.WearUserAvatarEquip(player.PlayerID, avatar.AvatarId, weapon.WeaponId)

	wearEquipRsp := &proto.WearEquipRsp{
		AvatarGuid: avatarGuid,
		EquipGuid:  equipGuid,
	}
	g.SendMsg(cmd.WearEquipRsp, player.PlayerID, player.ClientSeq, wearEquipRsp)
}

func (g *GameManager) WearUserAvatarEquip(userId uint32, avatarId uint32, weaponId uint64) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	avatar := player.AvatarMap[avatarId]
	weapon := player.WeaponMap[weaponId]

	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)

	if weapon.AvatarId != 0 {
		// 武器在别的角色身上
		weakAvatarId := weapon.AvatarId
		weakWeaponId := weaponId
		strongAvatarId := avatarId
		strongWeaponId := avatar.EquipWeapon.WeaponId
		player.TakeOffWeapon(weakAvatarId, weakWeaponId)
		player.TakeOffWeapon(strongAvatarId, strongWeaponId)
		player.WearWeapon(weakAvatarId, strongWeaponId)
		player.WearWeapon(strongAvatarId, weakWeaponId)

		weakAvatar := player.AvatarMap[weakAvatarId]
		weakWeapon := player.WeaponMap[weakAvatar.EquipWeapon.WeaponId]

		weakWorldAvatar := world.GetPlayerWorldAvatar(player, weakAvatarId)
		if weakWorldAvatar != nil {
			weakWorldAvatar.weaponEntityId = scene.CreateEntityWeapon()
			avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotify(weakAvatar, weakWeapon, weakWorldAvatar.weaponEntityId)
			g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
		} else {
			avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotify(weakAvatar, weakWeapon, 0)
			g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
		}
	} else if avatar.EquipWeapon != nil {
		// 角色当前有武器
		player.TakeOffWeapon(avatarId, avatar.EquipWeapon.WeaponId)
		player.WearWeapon(avatarId, weaponId)
	} else {
		// 是新角色还没有武器
		player.WearWeapon(avatarId, weaponId)
	}

	worldAvatar := world.GetPlayerWorldAvatar(player, avatarId)
	if worldAvatar != nil {
		worldAvatar.weaponEntityId = scene.CreateEntityWeapon()
		avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotify(avatar, weapon, worldAvatar.weaponEntityId)
		g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
	} else {
		avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotify(avatar, weapon, 0)
		g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
	}
}

func (g *GameManager) AvatarChangeCostumeReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user change avatar costume, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.AvatarChangeCostumeReq)
	avatarGuid := req.AvatarGuid
	costumeId := req.CostumeId

	exist := false
	for _, v := range player.CostumeList {
		if v == costumeId {
			exist = true
		}
	}
	if costumeId == 0 {
		exist = true
	}
	if !exist {
		return
	}

	avatar := player.GameObjectGuidMap[avatarGuid].(*model.Avatar)
	avatar.Costume = req.CostumeId

	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)

	avatarChangeCostumeNotify := new(proto.AvatarChangeCostumeNotify)
	avatarChangeCostumeNotify.EntityInfo = g.PacketSceneEntityInfoAvatar(scene, player, avatar.AvatarId)
	for _, scenePlayer := range scene.playerMap {
		g.SendMsg(cmd.AvatarChangeCostumeNotify, scenePlayer.PlayerID, scenePlayer.ClientSeq, avatarChangeCostumeNotify)
	}

	avatarChangeCostumeRsp := &proto.AvatarChangeCostumeRsp{
		AvatarGuid: req.AvatarGuid,
		CostumeId:  req.CostumeId,
	}
	g.SendMsg(cmd.AvatarChangeCostumeRsp, player.PlayerID, player.ClientSeq, avatarChangeCostumeRsp)
}

func (g *GameManager) AvatarWearFlycloakReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user change avatar fly cloak, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.AvatarWearFlycloakReq)
	avatarGuid := req.AvatarGuid
	flycloakId := req.FlycloakId

	exist := false
	for _, v := range player.FlyCloakList {
		if v == flycloakId {
			exist = true
		}
	}
	if !exist {
		return
	}

	avatar := player.GameObjectGuidMap[avatarGuid].(*model.Avatar)
	avatar.FlyCloak = req.FlycloakId

	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)

	avatarFlycloakChangeNotify := &proto.AvatarFlycloakChangeNotify{
		AvatarGuid: avatarGuid,
		FlycloakId: flycloakId,
	}
	for _, scenePlayer := range scene.playerMap {
		g.SendMsg(cmd.AvatarFlycloakChangeNotify, scenePlayer.PlayerID, scenePlayer.ClientSeq, avatarFlycloakChangeNotify)
	}

	avatarWearFlycloakRsp := &proto.AvatarWearFlycloakRsp{
		AvatarGuid: req.AvatarGuid,
		FlycloakId: req.FlycloakId,
	}
	g.SendMsg(cmd.AvatarWearFlycloakRsp, player.PlayerID, player.ClientSeq, avatarWearFlycloakRsp)
}

func (g *GameManager) PacketAvatarEquipChangeNotify(avatar *model.Avatar, weapon *model.Weapon, entityId uint32) *proto.AvatarEquipChangeNotify {
	itemDataConfig, ok := gdconf.CONF.ItemDataMap[int32(weapon.ItemId)]
	if !ok {
		logger.Error("item data config error, itemId: %v", weapon.ItemId)
		return new(proto.AvatarEquipChangeNotify)
	}
	avatarEquipChangeNotify := &proto.AvatarEquipChangeNotify{
		AvatarGuid: avatar.Guid,
		ItemId:     weapon.ItemId,
		EquipGuid:  weapon.Guid,
	}
	switch itemDataConfig.Type {
	case int32(constant.ItemTypeConst.ITEM_WEAPON):
		avatarEquipChangeNotify.EquipType = uint32(constant.EquipTypeConst.EQUIP_WEAPON)
	case int32(constant.ItemTypeConst.ITEM_RELIQUARY):
		avatarEquipChangeNotify.EquipType = uint32(itemDataConfig.ReliquaryType)
	}
	avatarEquipChangeNotify.Weapon = &proto.SceneWeaponInfo{
		EntityId:    entityId,
		GadgetId:    uint32(itemDataConfig.GadgetId),
		ItemId:      weapon.ItemId,
		Guid:        weapon.Guid,
		Level:       uint32(weapon.Level),
		AbilityInfo: new(proto.AbilitySyncStateInfo),
	}
	return avatarEquipChangeNotify
}

func (g *GameManager) PacketAvatarEquipTakeOffNotify(avatar *model.Avatar, weapon *model.Weapon) *proto.AvatarEquipChangeNotify {
	avatarEquipChangeNotify := &proto.AvatarEquipChangeNotify{
		AvatarGuid: avatar.Guid,
	}
	itemDataConfig, exist := gdconf.CONF.ItemDataMap[int32(weapon.ItemId)]
	if exist {
		avatarEquipChangeNotify.EquipType = uint32(itemDataConfig.Type)
	}
	return avatarEquipChangeNotify
}

func (g *GameManager) UpdateUserAvatarFightProp(userId uint32, avatarId uint32) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	avatar, ok := player.AvatarMap[avatarId]
	if !ok {
		logger.Error("avatar is nil, avatarId: %v", avatar)
		return
	}
	// 角色初始化面板
	player.InitAvatarFightProp(avatar)

	avatarFightPropNotify := &proto.AvatarFightPropNotify{
		AvatarGuid:   avatar.Guid,
		FightPropMap: avatar.FightPropMap,
	}
	g.SendMsg(cmd.AvatarFightPropNotify, userId, player.ClientSeq, avatarFightPropNotify)
}

func (g *GameManager) PacketAvatarInfo(avatar *model.Avatar) *proto.AvatarInfo {
	isFocus := false
	if avatar.AvatarId == 10000005 || avatar.AvatarId == 10000007 {
		isFocus = true
	}
	pbAvatar := &proto.AvatarInfo{
		IsFocus:  isFocus,
		AvatarId: avatar.AvatarId,
		Guid:     avatar.Guid,
		PropMap: map[uint32]*proto.PropValue{
			uint32(constant.PlayerPropertyConst.PROP_LEVEL): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_LEVEL),
				Val:   int64(avatar.Level),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Level)},
			},
			uint32(constant.PlayerPropertyConst.PROP_EXP): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_EXP),
				Val:   int64(avatar.Exp),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Exp)},
			},
			uint32(constant.PlayerPropertyConst.PROP_BREAK_LEVEL): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_BREAK_LEVEL),
				Val:   int64(avatar.Promote),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Promote)},
			},
			uint32(constant.PlayerPropertyConst.PROP_SATIATION_VAL): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_SATIATION_VAL),
				Val:   int64(avatar.Satiation),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.Satiation)},
			},
			uint32(constant.PlayerPropertyConst.PROP_SATIATION_PENALTY_TIME): {
				Type:  uint32(constant.PlayerPropertyConst.PROP_SATIATION_PENALTY_TIME),
				Val:   int64(avatar.SatiationPenalty),
				Value: &proto.PropValue_Ival{Ival: int64(avatar.SatiationPenalty)},
			},
		},
		LifeState:     uint32(avatar.LifeState),
		EquipGuidList: object.ConvMapToList(avatar.EquipGuidMap),
		FightPropMap:  avatar.FightPropMap,
		SkillDepotId:  avatar.SkillDepotId,
		FetterInfo: &proto.AvatarFetterInfo{
			ExpLevel:                uint32(avatar.FetterLevel),
			ExpNumber:               avatar.FetterExp,
			FetterList:              nil,
			RewardedFetterLevelList: []uint32{10},
		},
		SkillLevelMap:            avatar.SkillLevelMap,
		AvatarType:               1,
		WearingFlycloakId:        avatar.FlyCloak,
		CostumeId:                avatar.Costume,
		BornTime:                 uint32(avatar.BornTime),
		PendingPromoteRewardList: make([]uint32, 0, len(avatar.PromoteRewardMap)),
	}
	for _, v := range avatar.FetterList {
		pbAvatar.FetterInfo.FetterList = append(pbAvatar.FetterInfo.FetterList, &proto.FetterData{
			FetterId:    v,
			FetterState: uint32(constant.FetterStateConst.FINISH),
		})
	}
	// 解锁全部资料
	for _, v := range gdconf.CONF.FetterDataAvatarIdMap[int32(avatar.AvatarId)] {
		pbAvatar.FetterInfo.FetterList = append(pbAvatar.FetterInfo.FetterList, &proto.FetterData{
			FetterId:    uint32(v),
			FetterState: uint32(constant.FetterStateConst.FINISH),
		})
	}
	// 突破等级奖励
	for promoteLevel, isTaken := range avatar.PromoteRewardMap {
		if !isTaken {
			pbAvatar.PendingPromoteRewardList = append(pbAvatar.PendingPromoteRewardList, promoteLevel)
		}
	}
	return pbAvatar
}
