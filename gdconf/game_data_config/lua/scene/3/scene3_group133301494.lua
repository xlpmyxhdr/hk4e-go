-- 基础信息
local base_info = {
	group_id = 133301494
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 494002, monster_id = 28060602, pos = { x = -659.171, y = 204.057, z = 3433.825 }, rot = { x = 0.000, y = 21.298, z = 0.000 }, level = 33, drop_tag = "走兽", pose_id = 1, area_id = 22 },
	{ config_id = 494005, monster_id = 28050106, pos = { x = -627.572, y = 222.437, z = 3288.988 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, drop_tag = "魔法生物", area_id = 22 },
	{ config_id = 494006, monster_id = 28010208, pos = { x = -637.675, y = 225.345, z = 3265.052 }, rot = { x = 0.000, y = 32.651, z = 0.000 }, level = 33, drop_tag = "采集动物", area_id = 22 },
	{ config_id = 494007, monster_id = 28020313, pos = { x = -631.418, y = 205.568, z = 3342.431 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, drop_tag = "走兽", area_id = 22 },
	{ config_id = 494008, monster_id = 28020314, pos = { x = -627.106, y = 203.518, z = 3345.456 }, rot = { x = 0.000, y = 329.095, z = 0.000 }, level = 33, drop_tag = "走兽", area_id = 22 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

-- 废弃数据
garbages = {
	monsters = {
		{ config_id = 494001, monster_id = 28060601, pos = { x = -659.003, y = 207.362, z = 3425.579 }, rot = { x = 0.000, y = 246.958, z = 0.000 }, level = 33, drop_tag = "走兽", pose_id = 3, area_id = 22 }
	}
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { 494002, 494005, 494006, 494007, 494008 },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================