-- 基础信息
local base_info = {
	group_id = 133301349
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 349003, monster_id = 28050106, pos = { x = -524.512, y = 161.750, z = 3559.623 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, drop_tag = "魔法生物", area_id = 22 },
	{ config_id = 349004, monster_id = 28050106, pos = { x = -523.423, y = 158.866, z = 3566.904 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, drop_tag = "魔法生物", area_id = 22 },
	{ config_id = 349005, monster_id = 28030101, pos = { x = -498.499, y = 156.955, z = 3585.731 }, rot = { x = 0.000, y = 330.200, z = 0.000 }, level = 33, drop_tag = "鸟类", disableWander = true, area_id = 22 },
	{ config_id = 349006, monster_id = 28030101, pos = { x = -476.819, y = 158.089, z = 3634.914 }, rot = { x = 0.000, y = 196.435, z = 0.000 }, level = 33, drop_tag = "鸟类", disableWander = true, area_id = 22 },
	{ config_id = 349007, monster_id = 28030101, pos = { x = -453.973, y = 156.980, z = 3603.280 }, rot = { x = 0.000, y = 205.041, z = 0.000 }, level = 33, drop_tag = "鸟类", disableWander = true, area_id = 22 },
	{ config_id = 349008, monster_id = 28030101, pos = { x = -486.034, y = 156.472, z = 3589.754 }, rot = { x = 0.000, y = 0.573, z = 0.000 }, level = 33, drop_tag = "鸟类", disableWander = true, area_id = 22 }
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
		{ config_id = 349001, monster_id = 28050106, pos = { x = -508.780, y = 160.863, z = 3563.923 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, drop_tag = "魔法生物", area_id = 22 },
		{ config_id = 349002, monster_id = 28050106, pos = { x = -516.988, y = 159.193, z = 3564.854 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 33, drop_tag = "魔法生物", area_id = 22 }
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
		monsters = { 349003, 349004, 349005, 349006, 349007, 349008 },
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