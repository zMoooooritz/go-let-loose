---@meta

-- Admin Management API

---Add an admin user
---@param id string Player ID to make admin
---@param comment string Comment describing the admin
---@param role string Admin role ("owner", "senior", "junior", "spectator")
---@return string|nil error Error message if any
function addAdmin(id, comment, role) end

---Remove an admin user
---@param id string Player ID to remove admin status from
---@return string|nil error Error message if any
function removeAdmin(id) end

---Get list of admin roles
---@return string|nil error Error message if any
---@return string[]|nil roles Array of admin role names if successful
function getAdminRoles() end

---Get list of admins
---@return string|nil error Error message if any
---@return Admin[]|nil admins Array of admin information if successful
function getAdmins() end

-- VIP Management API

---Add a VIP user
---@param id string Player ID to make VIP
---@param comment string Comment describing the VIP
---@return string|nil error Error message if any
function addVip(id, comment) end

---Remove a VIP user
---@param id string Player ID to remove VIP status from
---@return string|nil error Error message if any
function removeVip(id) end

---Get list of VIP players
---@return string|nil error Error message if any
---@return PlayerInfo[]|nil vips Array of VIP player information if successful
function getVIPs() end

-- Map Management API

---Add map to rotation
---@param layer Layer Layer to add to rotation
---@param index number Position in rotation to insert at
---@return string|nil error Error message if any
function addMapToRotation(layer, index) end

---Add map to sequence
---@param layer Layer Layer to add to sequence
---@param index number Position in sequence to insert at
---@return string|nil error Error message if any
function addMapToSequence(layer, index) end

---Remove map from rotation
---@param index number Index of map to remove from rotation
---@return string|nil error Error message if any
function removeMapFromRotation(index) end

---Remove map from sequence
---@param index number Index of map to remove from sequence
---@return string|nil error Error message if any
function removeMapToSequence(index) end

---Move map position in sequence
---@param from number Current position
---@param to number New position
---@return string|nil error Error message if any
function moveMapInSequence(from, to) end

---Set current map
---@param layer Layer Layer to switch to
---@return string|nil error Error message if any
function setCurrentMap(layer) end

---Shuffle map sequence
---@param enabled boolean Whether to enable sequence shuffling
---@return string|nil error Error message if any
function shuffleMapSequence(enabled) end

---Get all available maps
---@return string|nil error Error message if any
---@return Layer[]|nil maps Array of available layers if successful
function getAllMaps() end

---Get current layer
---@return string|nil error Error message if any
---@return Layer|nil layer Current layer if successful
function getCurrentLayer() end

---Get current map
---@return string|nil error Error message if any
---@return GameMap|nil map Current map if successful
function getCurrentMap() end

---Get current map rotation
---@return string|nil error Error message if any
---@return Layer[]|nil rotation Array of layers in rotation if successful
function getCurrentMapRotation() end

---Get current map sequence
---@return string|nil error Error message if any
---@return Layer[]|nil sequence Array of layers in sequence if successful
function getCurrentMapSequence() end

---Get current map objectives
---@return string|nil error Error message if any
---@return string[][]|nil objectives 2D array of objective information if successful
function getCurrentMapObjectives() end

---Get current game mode
---@return string|nil error Error message if any
---@return string|nil gameMode Current game mode if successful
function getGameMode() end

-- Player Information API

---Get basic player list
---@return string|nil error Error message if any
---@return PlayerInfo[]|nil players Array of player information if successful
function getPlayers() end

---Get detailed player information
---@return string|nil error Error message if any
---@return DetailedPlayerInfo[]|nil players Array of detailed player information if successful
function getPlayersInfo() end

---Get player names only
---@return string|nil error Error message if any
---@return string[]|nil names Array of player names if successful
function getPlayerNames() end

---Get player IDs only
---@return string|nil error Error message if any
---@return string[]|nil ids Array of player IDs if successful
function getPlayerIDs() end

---Get information for specific player
---@param playerID string Player ID to get info for
---@return string|nil error Error message if any
---@return DetailedPlayerInfo|nil player Player information if successful
function getPlayerInfo(playerID) end

---Get player counts per team
---@return string|nil error Error message if any
---@return TeamData|nil counts Player counts if successful
function getPlayerCounts() end

---Get server view with team/squad organization
---@return string|nil error Error message if any
---@return ServerView|nil view Server view if successful
function getServerView() end

-- Player Moderation API

---Kick a player
---@param player string Player name or ID
---@param reason string Reason for kick
---@return string|nil error Error message if any
function kickPlayer(player, reason) end

---Temporarily ban a player
---@param player string Player name or ID
---@param duration number Ban duration in minutes
---@param reason string Reason for ban
---@param admin string Admin name issuing the ban
---@return string|nil error Error message if any
function tempBanPlayer(player, duration, reason, admin) end

---Permanently ban a player
---@param player string Player name or ID
---@param reason string Reason for ban
---@param admin string Admin name issuing the ban
---@return string|nil error Error message if any
function permaBanPlayer(player, reason, admin) end

---Pardon a temporary ban
---@param ban ServerBan Ban to pardon
---@return string|nil error Error message if any
function pardonTempBanPlayer(ban) end

---Pardon a permanent ban
---@param ban ServerBan Ban to pardon
---@return string|nil error Error message if any
function pardonPermaBanPlayer(ban) end

---Get temporary bans list
---@return string|nil error Error message if any
---@return ServerBan[]|nil bans Array of temporary bans if successful
function getTempBans() end

---Get permanent bans list
---@return string|nil error Error message if any
---@return ServerBan[]|nil bans Array of permanent bans if successful
function getPermaBans() end

---Send message to specific player
---@param playerID string Player ID to message
---@param message string Message to send
---@return string|nil error Error message if any
function messagePlayer(playerID, message) end

---Punish a player
---@param player string Player name or ID
---@param reason string Reason for punishment
---@return string|nil error Error message if any
function punishPlayer(player, reason) end

---Remove player from platoon
---@param player string Player name or ID
---@param reason string Reason for removal
---@return string|nil error Error message if any
function removePlayerFromPlatoon(player, reason) end

---Disband a platoon
---@param team string Team name
---@param unit Unit Unit to disband
---@param reason string Reason for disbanding
---@return string|nil error Error message if any
function disbandPlatoon(team, unit, reason) end

---Switch player team immediately
---@param player string Player name or ID
---@return string|nil error Error message if any
function switchPlayerNow(player) end

---Switch player team on death
---@param player string Player name or ID
---@return string|nil error Error message if any
function switchPlayerOnDeath(player) end

-- Server Information API

---Get server name
---@return string|nil error Error message if any
---@return string|nil name Server name if successful
function getServerName() end

---Get server configuration
---@return string|nil error Error message if any
---@return ServerConfig|nil config Server configuration if successful
function getServerConfig() end

---Get session information
---@return string|nil error Error message if any
---@return SessionInfo|nil info Session information if successful
function getSessionInfo() end

---Get current game state
---@return string|nil error Error message if any
---@return GameState|nil state Game state if successful
function getGameState() end

---Get player slots (current, max)
---@return string|nil error Error message if any
---@return number|nil current Current player count if successful
---@return number|nil max Maximum player count if successful
function getSlots() end

---Get team scores
---@return string|nil error Error message if any
---@return TeamData|nil scores Team scores if successful
function getScore() end

---Get server changelist/version
---@return string|nil error Error message if any
---@return string|nil changelist Server changelist if successful
function getServerChangelist() end

-- Server Configuration API

---Set broadcast message
---@param message string Message to broadcast
---@return string|nil error Error message if any
function setBroadcastMessage(message) end

---Clear broadcast message
---@param message string Message to clear
---@return string|nil error Error message if any
function clearBroadcastMessage(message) end

---Set welcome message
---@param message string Welcome message to set
---@return string|nil error Error message if any
function setWelcomeMessage(message) end

---Enable/disable auto-balance
---@param enabled boolean Whether to enable auto-balance
---@return string|nil error Error message if any
function setAutoBalanceEnabled(enabled) end

---Check if auto-balance is enabled
---@return string|nil error Error message if any
---@return boolean|nil enabled Whether auto-balance is enabled if successful
function isAutoBalanceEnabled() end

---Set auto-balance threshold
---@param threshold number Player difference threshold for auto-balance
---@return string|nil error Error message if any
function setAutoBalanceThreshold(threshold) end

---Get auto-balance threshold
---@return string|nil error Error message if any
---@return number|nil threshold Auto-balance threshold if successful
function getAutoBalanceThreshold() end

---Enable/disable vote kick
---@param enabled boolean Whether to enable vote kick
---@return string|nil error Error message if any
function setVoteKickEnabled(enabled) end

---Check if vote kick is enabled
---@return string|nil error Error message if any
---@return boolean|nil enabled Whether vote kick is enabled if successful
function isVoteKickEnabled() end

---Set vote kick threshold
---@param thresholdPairs string Threshold pairs configuration
---@return string|nil error Error message if any
function setVoteKickThreshold(thresholdPairs) end

---Get vote kick threshold
---@return string|nil error Error message if any
---@return number|nil threshold Vote kick threshold if successful
function getVoteKickThreshold() end

---Reset vote kick threshold to default
---@return string|nil error Error message if any
function resetVoteKickThreshold() end

---Set team switch cooldown
---@param cooldown number Cooldown in seconds
---@return string|nil error Error message if any
function setTeamSwitchCooldown(cooldown) end

---Get team switch cooldown
---@return string|nil error Error message if any
---@return number|nil cooldown Team switch cooldown if successful
function getTeamSwitchCooldown() end

---Set high ping threshold
---@param threshold number Ping threshold in milliseconds
---@return string|nil error Error message if any
function setHighPing(threshold) end

---Get high ping threshold
---@return string|nil error Error message if any
---@return number|nil threshold High ping threshold if successful
function getHighPing() end

---Set idle kick time
---@param threshold number Idle time in seconds before kick
---@return string|nil error Error message if any
function setKickIdleTime(threshold) end

---Get idle kick time
---@return string|nil error Error message if any
---@return number|nil time Idle kick time if successful
function getIdleTime() end

---Set maximum queued players
---@param size number Maximum queue size
---@return string|nil error Error message if any
function setMaxQueuedPlayers(size) end

---Get maximum queued players
---@return string|nil error Error message if any
---@return number|nil size Maximum queue size if successful
function getMaxQueuedPlayers() end

---Get number of queued players
---@return string|nil error Error message if any
---@return number|nil count Number of queued players if successful
function getQueuedPlayers() end

---Set number of VIP slots
---@param amount number Number of VIP slots
---@return string|nil error Error message if any
function setNumVipSlots(amount) end

---Get number of VIP slots
---@return string|nil error Error message if any
---@return number|nil slots Number of VIP slots if successful
function getNumVipSlots() end

---Get number of queued VIPs
---@return string|nil error Error message if any
---@return number|nil count Number of queued VIPs if successful
function getQueuedVips() end

-- Game Management API

---Set game layout with objective names
---@param objs string[] Array of objective names
---@return string|nil error Error message if any
function setGameLayout(objs) end

---Set game layout with objective indices
---@param objs number[] Array of objective indices
---@return string|nil error Error message if any
function setGameLayoutIndexed(objs) end

---Set match timer for game mode
---@param gameMode string Game mode
---@param duration number Timer duration in seconds
---@return string|nil error Error message if any
function setMatchTimer(gameMode, duration) end

---Remove match timer for game mode
---@param gameMode string Game mode
---@return string|nil error Error message if any
function removeMatchTimer(gameMode) end

---Set warmup timer for game mode
---@param gameMode string Game mode
---@param duration number Timer duration in seconds
---@return string|nil error Error message if any
function setWarmupTimer(gameMode, duration) end

---Remove warmup timer for game mode
---@param gameMode string Game mode
---@return string|nil error Error message if any
function removeWarmupTimer(gameMode) end

---Toggle dynamic weather for layer
---@param layer Layer Layer to toggle weather for
---@param enabled boolean Whether to enable dynamic weather
---@return string|nil error Error message if any
function setDynamicWeatherToggle(layer, enabled) end

-- Profanity Management API

---Ban profanities (add to banned words list)
---@param profanities string[] Array of words to ban
---@return string|nil error Error message if any
function banProfanities(profanities) end

---Unban profanities (remove from banned words list)
---@param profanities string[] Array of words to unban
---@return string|nil error Error message if any
function unbanProfanities(profanities) end

---Get list of banned profanities
---@return string|nil error Error message if any
---@return string[]|nil words Array of banned words if successful
function getProfanities() end

-- Logging API

---Get log entries
---@param seconds number Number of seconds back to retrieve logs
---@param filters string Log filters to apply
---@return string|nil error Error message if any
---@return LogEntry[]|nil entries Array of log entries if successful
function getLogEntries(seconds, filters) end

---Get raw logs
---@param spanMins number Number of minutes back to retrieve logs
---@return string|nil error Error message if any
---@return string[]|nil logs Array of raw log lines if successful
function getLogs(spanMins) end

-- Command Information API

---Get available commands
---@return string|nil error Error message if any
---@return Command[]|nil commands Array of available commands if successful
function getCommands() end

---Get command details
---@param commandID string Command ID to get details for
---@return string|nil error Error message if any
---@return CommandDetails|nil details Command details if successful
function getCommandDetails(commandID) end

-- Utility Functions

---Get list of available commands (returns function signatures)
---@return table commands Array of command signatures
function listCommands() end

---Exit the plugin (must be called in Stop function)
function exit() end