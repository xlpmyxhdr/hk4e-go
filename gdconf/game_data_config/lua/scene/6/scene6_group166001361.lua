-- 基础信息
local base_info = {
	group_id = 166001361
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 361001, monster_id = 25010201, pos = { x = 957.389, y = 842.599, z = 595.300 }, rot = { x = 0.000, y = 157.687, z = 0.000 }, level = 36, drop_tag = "盗宝团", pose_id = 9004, area_id = 300 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 361002, gadget_id = 70290308, pos = { x = 955.040, y = 842.466, z = 594.654 }, rot = { x = 0.000, y = 322.835, z = 0.000 }, level = 36, area_id = 300 },
	{ config_id = 361003, gadget_id = 70210101, pos = { x = 945.119, y = 843.138, z = 604.289 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜矿石璃月", persistent = true, area_id = 300 }
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
		monsters = { 361001 },
		gadgets = { 361002, 361003 },
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