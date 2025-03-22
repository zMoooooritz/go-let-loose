-- Print contents of `tbl`, with indentation.
-- `indent` sets the initial level of indentation.
local function tprint(tbl, indent)
	if not indent then
		indent = 0
	end
	for k, v in pairs(tbl) do
		local formatting = string.rep("  ", indent) .. k .. ": "
		if type(v) == "table" then
			print(formatting)
			tprint(v, indent + 1)
		elseif type(v) == "boolean" then
			print(formatting .. tostring(v))
		else
			print(formatting .. v)
		end
	end
end

local function onKill(event)
	print("Kill: " .. event.Killer.Name .. " -> " .. event.Victim.Name .. " (" .. event.Weapon.Name .. ")")
end

function Init()
	-- Here goes any initialization logic
	print("The available commands are:")
	tprint(listCommands()) -- get a full list of the available commands via help()
	print()
	print("The subscribable events are:")
	tprint(listEvents())
	print()
end

function Run()
	local err, name = getServerName()
	if name then
		print("Connected to the Server: " .. name)
		registerEvent("KILL", onKill)
	else
		print("Error: " .. err)
	end
end

function Stop()
	-- A stop signal will be send to this function from the outside
	-- Do any clean-up steps required
	exit() -- at the end of this function Exit() should always be called
end
