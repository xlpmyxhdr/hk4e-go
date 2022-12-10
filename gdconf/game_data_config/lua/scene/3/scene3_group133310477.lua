-- 基础信息
local base_info = {
	group_id = 133310477
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 477001, monster_id = 28020108, pos = { x = -2314.108, y = 328.648, z = 4170.665 }, rot = { x = 0.000, y = 95.768, z = 0.000 }, level = 32, drop_tag = "走兽", area_id = 26 },
	{ config_id = 477002, monster_id = 28020108, pos = { x = -2335.673, y = 327.251, z = 4169.644 }, rot = { x = 0.000, y = 258.196, z = 0.000 }, level = 32, drop_tag = "走兽", area_id = 26 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 477003, gadget_id = 70210101, pos = { x = -2312.061, y = 329.524, z = 4169.942 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 26, drop_tag = "搜刮点解谜果蔬须弥", persistent = true, area_id = 26 }
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
		monsters = { 477001, 477002 },
		gadgets = { 477003 },
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