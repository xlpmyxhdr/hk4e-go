-- 基础信息
local base_info = {
	group_id = 133221073
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
	-- 百人一揆二期活动NPC
	{ config_id = 73001, npc_id = 30065, pos = { x = -3378.854, y = 200.492, z = -4276.824 }, rot = { x = 0.000, y = 156.772, z = 0.000 }, area_id = 11 }
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
		gadgets = { },
		regions = { },
		triggers = { },
		npcs = { 73001 },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================