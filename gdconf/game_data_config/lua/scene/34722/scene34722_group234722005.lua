-- 基础信息
local base_info = {
	group_id = 234722005
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 5001, monster_id = 23050101, pos = { x = 0.023, y = -0.102, z = -10.008 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, disableWander = true },
	{ config_id = 5008, monster_id = 23040101, pos = { x = 0.026, y = -0.102, z = 10.001 }, rot = { x = 0.000, y = 180.000, z = 0.000 }, level = 1, disableWander = true }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 5011, gadget_id = 70900205, pos = { x = 6.204, y = -1.386, z = 3.136 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
	{ config_id = 1005005, name = "ANY_MONSTER_LIVE_5005", event = EventType.EVENT_ANY_MONSTER_LIVE, source = "", condition = "condition_EVENT_ANY_MONSTER_LIVE_5005", action = "action_EVENT_ANY_MONSTER_LIVE_5005" },
	{ config_id = 1005015, name = "CHALLENGE_SUCCESS_5015", event = EventType.EVENT_CHALLENGE_SUCCESS, source = "2", condition = "", action = "action_EVENT_CHALLENGE_SUCCESS_5015" },
	{ config_id = 1005016, name = "CHALLENGE_FAIL_5016", event = EventType.EVENT_CHALLENGE_FAIL, source = "2", condition = "", action = "action_EVENT_CHALLENGE_FAIL_5016" }
}

-- 变量
variables = {
	{ config_id = 1, name = "monster_wave", value = 0, no_refresh = false }
}

-- 废弃数据
garbages = {
	triggers = {
		{ config_id = 1005002, name = "TIME_AXIS_PASS_5002", event = EventType.EVENT_TIME_AXIS_PASS, source = "", condition = "condition_EVENT_TIME_AXIS_PASS_5002", action = "action_EVENT_TIME_AXIS_PASS_5002", trigger_count = 0 },
		{ config_id = 1005003, name = "TIME_AXIS_PASS_5003", event = EventType.EVENT_TIME_AXIS_PASS, source = "", condition = "condition_EVENT_TIME_AXIS_PASS_5003", action = "action_EVENT_TIME_AXIS_PASS_5003", trigger_count = 0 },
		{ config_id = 1005004, name = "TIME_AXIS_PASS_5004", event = EventType.EVENT_TIME_AXIS_PASS, source = "", condition = "condition_EVENT_TIME_AXIS_PASS_5004", action = "action_EVENT_TIME_AXIS_PASS_5004", trigger_count = 0 },
		{ config_id = 1005006, name = "TIME_AXIS_PASS_5006", event = EventType.EVENT_TIME_AXIS_PASS, source = "", condition = "condition_EVENT_TIME_AXIS_PASS_5006", action = "action_EVENT_TIME_AXIS_PASS_5006", trigger_count = 0 },
		{ config_id = 1005007, name = "TIME_AXIS_PASS_5007", event = EventType.EVENT_TIME_AXIS_PASS, source = "", condition = "condition_EVENT_TIME_AXIS_PASS_5007", action = "action_EVENT_TIME_AXIS_PASS_5007", trigger_count = 0 }
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
		monsters = { },
		gadgets = { 5011 },
		regions = { },
		triggers = { },
		rand_weight = 100
	},
	{
		-- suite_id = 2,
		-- description = suite_2,
		monsters = { 5001, 5008 },
		gadgets = { 5011 },
		regions = { },
		triggers = { "ANY_MONSTER_LIVE_5005", "CHALLENGE_SUCCESS_5015", "CHALLENGE_FAIL_5016" },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================

-- 触发条件
function condition_EVENT_ANY_MONSTER_LIVE_5005(context, evt)
	if 5001 ~= evt.param1 then
		return false
	end
	
	return true
end

-- 触发操作
function action_EVENT_ANY_MONSTER_LIVE_5005(context, evt)
	-- 创建编号为2（该挑战的识别id),挑战内容为198的区域挑战，param1必须为时间
	-- 从234722002的变量TPL_TIME中取出对应值并开启挑战
	  local tpl_time = ScriptLib.GetGroupVariableValueByGroup(context, "TPL_TIME", 234722002)
	  if tpl_time == nil or tpl_time < 0 then
	    ScriptLib.PrintContextLog(context, "@@ LUA_WARNING : active_challenge_by_remainTime")
	    return -1
	  elseif tpl_time < 1 then
	    tpl_time = 0
	  end
	  if 0 ~= ScriptLib.ActiveChallenge(context, 2, 198, tpl_time, 234722005, 2, 0) then
	    ScriptLib.PrintContextLog(context, "@@ LUA_WARNING : active_challenge_by_remainTime")
	    return -1
	  end
	
	-- 针对当前group内变量名为 "monster_wave" 的变量，进行修改，变化值为 1
	if 0 ~= ScriptLib.ChangeGroupVariableValue(context, "monster_wave", 1) then
	  ScriptLib.PrintContextLog(context, "@@ LUA_WARNING : change_GroupVariable")
	  return -1
	end
	
	return 0
end

-- 触发操作
function action_EVENT_CHALLENGE_SUCCESS_5015(context, evt)
	-- 将本组内变量名为 "stage" 的变量设置为 2
	if 0 ~= ScriptLib.SetGroupVariableValueByGroup(context, "stage", 2, 234722002) then
	  ScriptLib.PrintContextLog(context, "@@ LUA_WARNING : set_groupVariable_by_group")
	  return -1
	end
	
	return 0
end

-- 触发操作
function action_EVENT_CHALLENGE_FAIL_5016(context, evt)
		-- 重新生成指定group，指定suite
		if 0 ~= ScriptLib.RefreshGroup(context, { group_id = 234722005, suite = 1 }) then
	    ScriptLib.PrintContextLog(context, "@@ LUA_WARNING : refresh_group_to_suite")
			return -1
		end
	
	-- 地城失败结算
	if 0 ~= ScriptLib.CauseDungeonFail(context) then
		ScriptLib.PrintContextLog(context, "@@ LUA_WARNING : cause_dungeonfail")
		return -1
	end
	
	return 0
end