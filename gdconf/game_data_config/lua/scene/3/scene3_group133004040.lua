-- 基础信息
local base_info = {
	group_id = 133004040
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 40001, gadget_id = 70500000, pos = { x = 2780.590, y = 211.468, z = -771.630 }, rot = { x = 0.000, y = 106.302, z = 0.000 }, level = 20, point_type = 2002, area_id = 4 },
	{ config_id = 40002, gadget_id = 70500000, pos = { x = 2785.104, y = 210.664, z = -773.682 }, rot = { x = 0.000, y = 262.614, z = 0.000 }, level = 20, point_type = 2001, area_id = 4 }
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
		monsters = { },
		gadgets = { 40001, 40002 },
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