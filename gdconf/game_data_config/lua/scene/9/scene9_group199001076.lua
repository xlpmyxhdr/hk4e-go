-- 基础信息
local base_info = {
	group_id = 199001076
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
	{ config_id = 76001, gadget_id = 70500036, pos = { x = -389.513, y = 120.254, z = -540.682 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 20, arguments = { 47 }, area_id = 400 },
	{ config_id = 76002, gadget_id = 70710786, pos = { x = -389.125, y = 119.787, z = -540.863 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 20, area_id = 400 },
	{ config_id = 76003, gadget_id = 70710483, pos = { x = -389.554, y = 120.032, z = -540.851 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 20, area_id = 400 },
	{ config_id = 76004, gadget_id = 70710786, pos = { x = -389.513, y = 120.001, z = -540.682 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 20, area_id = 400 },
	{ config_id = 76005, gadget_id = 70710483, pos = { x = -389.147, y = 119.619, z = -540.496 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 20, area_id = 400 },
	{ config_id = 76006, gadget_id = 70710786, pos = { x = -389.531, y = 119.787, z = -540.405 }, rot = { x = 0.000, y = 55.228, z = 0.000 }, level = 20, area_id = 400 }
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
		gadgets = { 76001, 76002, 76003, 76004, 76005, 76006 },
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