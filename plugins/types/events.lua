---@meta

---Base event structure
---@class BaseEvent
---@field EventType string The type of event
---@field EventTime string ISO timestamp when the event occurred
local BaseEvent = {}

---Kill event - fired when a player kills another player
---@class KillEvent : BaseEvent
---@field Killer PlayerInfo The player who made the kill
---@field Victim PlayerInfo The player who was killed
---@field Weapon Weapon The weapon used for the kill
local KillEvent = {}

---Death event - fired when a player dies
---@class DeathEvent : BaseEvent
---@field Victim PlayerInfo The player who died
---@field Killer PlayerInfo The player who caused the death
---@field Weapon Weapon The weapon that caused the death
local DeathEvent = {}

---Team kill event - fired when a player team kills
---@class TeamKillEvent : BaseEvent
---@field Killer PlayerInfo The player who made the team kill
---@field Victim PlayerInfo The teammate who was killed
---@field Weapon Weapon The weapon used for the team kill
local TeamKillEvent = {}

---Team death event - fired when a player dies to a teammate
---@class TeamDeathEvent : BaseEvent
---@field Victim PlayerInfo The player who died to a teammate
---@field Killer PlayerInfo The teammate who caused the death
---@field Weapon Weapon The weapon that caused the team death
local TeamDeathEvent = {}

---Chat event - fired when a player sends a chat message
---@class ChatEvent : BaseEvent
---@field Player PlayerInfo The player who sent the message
---@field Team string The team scope ("Allies"/"Axis")
---@field Scope string Chat scope ("Team"/"Unit"/"None")
---@field Message string The chat message content
local ChatEvent = {}

---Connection event - fired when a player connects
---@class ConnectEvent : BaseEvent
---@field Player PlayerInfo The player who connected
local ConnectEvent = {}

---Disconnection event - fired when a player disconnects
---@class DisconnectEvent : BaseEvent
---@field Player PlayerInfo The player who disconnected
local DisconnectEvent = {}

---Ban event - fired when a player is banned
---@class BanEvent : BaseEvent
---@field Player PlayerInfo The player who was banned
---@field Reason string The reason for the ban
local BanEvent = {}

---Kick event - fired when a player is kicked
---@class KickEvent : BaseEvent
---@field Player PlayerInfo The player who was kicked
---@field Reason string The reason for the kick
local KickEvent = {}

---Message event - fired for admin messages
---@class MessageEvent : BaseEvent
---@field Player PlayerInfo The player associated with the message
---@field Message string The message content
local MessageEvent = {}

---Match start event - fired when a match begins
---@class MatchStartEvent : BaseEvent
---@field Map GameMap Information about the map
local MatchStartEvent = {}

---Match end event - fired when a match ends
---@class MatchEndEvent : BaseEvent
---@field Map GameMap Information about the map
---@field Score TeamData The final scores for each team
local MatchEndEvent = {}

---Admin camera entered events - fired when entering admin camera
---@class AdminCamEnteredEvent : BaseEvent
---@field Player PlayerInfo The admin who entered camera mode
local AdminCamEnteredEvent = {}

---Admin camera left events - fired when leaving admin camera
---@class AdminCamLeftEvent : BaseEvent
---@field Player PlayerInfo The admin who left camera mode
local AdminCamLeftEvent = {}

---Vote started event - fired when a vote kick is initiated
---@class VoteStartedEvent : BaseEvent
---@field ID integer The unique vote ID
---@field Initiator PlayerInfo The player who initiated the vote
---@field Target PlayerInfo The player targeted by the vote
local VoteStartedEvent = {}

---Vote submitted event - fired when a player submits a vote
---@class VoteSubmittedEvent : BaseEvent
---@field Submitter PlayerInfo The player who submitted a vote
---@field ID integer The unique vote ID
---@field Vote string The vote choice
local VoteSubmittedEvent = {}

---Vote completed event - fired when a vote kick concludes
---@class VoteCompletedEvent : BaseEvent
---@field Reason string The reason for the vote
---@field Result string The result of the vote
---@field ID integer The unique vote ID
---@field Initiator PlayerInfo The player who initiated the vote
---@field Target PlayerInfo The player who was the target of the vote
local VoteCompletedEvent = {}

---Team switch event - fired when a player switches teams
---@class TeamSwitchEvent : BaseEvent
---@field Player PlayerInfo The player who switched teams
---@field From string The team they left
---@field To string The team they joined
local TeamSwitchEvent = {}

---Squad switch event - fired when a player switches squads
---@class SquadSwitchEvent : BaseEvent
---@field Player PlayerInfo The player who switched squads
---@field OldSquad Unit The squad they left
---@field NewSquad Unit The squad they joined
local SquadSwitchEvent = {}

---Score update event - fired when player scores are updated
---@class ScoreUpdateEvent : BaseEvent
---@field Player PlayerInfo The player whose score was updated
---@field OldScore Score The previous score
---@field NewScore Score The new score
local ScoreUpdateEvent = {}

---Role change event - fired when a player changes roles
---@class RoleChangeEvent : BaseEvent
---@field Player PlayerInfo The player who changed roles
---@field OldRole string The previous role
---@field NewRole string The new role
local RoleChangeEvent = {}

---Loadout change event - fired when a player changes loadout
---@class LoadoutChangeEvent : BaseEvent
---@field Player PlayerInfo The player who changed loadout
---@field OldLoadout string The previous loadout
---@field NewLoadout string The new loadout
local LoadoutChangeEvent = {}

---Objective captured event - fired when an objective is captured
---@class ObjectiveEvent : BaseEvent
---@field Player PlayerInfo The player who captured the objective
local ObjectiveEvent = {}

---Position change event - fired when a player moves significantly
---@class PositionEvent : BaseEvent
---@field Player PlayerInfo The player who moved
---@field OldPos Position The previous position
---@field NewPos Position The new position
local PositionEvent = {}

---Clan tag change event - fired when a player changes clan tag
---@class ClanTagChangeEvent : BaseEvent
---@field Player PlayerInfo The player who changed their clan tag
---@field OldClanTag string The previous clan tag
---@field NewClanTag string The new clan tag
local ClanTagChangeEvent = {}

---Register a kill event handler
---@param callback fun(event: KillEvent): nil
function onKill(callback) end

---Register a death event handler  
---@param callback fun(event: DeathEvent): nil
function onDeath(callback) end

---Register a team kill event handler
---@param callback fun(event: TeamKillEvent): nil
function onTeamKill(callback) end

---Register a team death event handler
---@param callback fun(event: TeamDeathEvent): nil
function onTeamDeath(callback) end

---Register a player connect event handler
---@param callback fun(event: ConnectEvent): nil
function onConnected(callback) end

---Register a player disconnect event handler
---@param callback fun(event: DisconnectEvent): nil
function onDisconnected(callback) end

---Register a chat event handler
---@param callback fun(event: ChatEvent): nil
function onChat(callback) end

---Register a ban event handler
---@param callback fun(event: BanEvent): nil
function onBan(callback) end

---Register a kick event handler
---@param callback fun(event: KickEvent): nil
function onKick(callback) end

---Register a message event handler
---@param callback fun(event: MessageEvent): nil
function onMessage(callback) end

---Register a match start event handler
---@param callback fun(event: MatchStartEvent): nil
function onMatchStart(callback) end

---Register a match end event handler
---@param callback fun(event: MatchEndEvent): nil
function onMatchEnd(callback) end

---Register an admin camera enter event handler
---@param callback fun(event: AdminCamEnteredEvent): nil
function onEnterAdminCam(callback) end

---Register an admin camera leave event handler
---@param callback fun(event: AdminCamLeftEvent): nil
function onLeaveAdminCam(callback) end

---Register a vote kick started event handler
---@param callback fun(event: VoteStartedEvent): nil
function onVoteKickStarted(callback) end

---Register a vote submitted event handler
---@param callback fun(event: VoteSubmittedEvent): nil
function onVoteSubmitted(callback) end

---Register a vote kick completed event handler
---@param callback fun(event: VoteCompletedEvent): nil
function onVoteKickCompleted(callback) end

---Register a team switch event handler
---@param callback fun(event: TeamSwitchEvent): nil
function onTeamSwitched(callback) end

---Register a squad switch event handler
---@param callback fun(event: SquadSwitchEvent): nil
function onSquadSwitched(callback) end

---Register a score update event handler
---@param callback fun(event: ScoreUpdateEvent): nil
function onScoreUpdate(callback) end

---Register a role change event handler
---@param callback fun(event: RoleChangeEvent): nil
function onRoleChanged(callback) end

---Register a loadout change event handler
---@param callback fun(event: LoadoutChangeEvent): nil
function onLoadoutChanged(callback) end

---Register an objective captured event handler
---@param callback fun(event: ObjectiveEvent): nil
function onObjectiveCapped(callback) end

---Register a position change event handler
---@param callback fun(event: PositionEvent): nil
function onPositionChanged(callback) end

---Register a clan tag change event handler
---@param callback fun(event: ClanTagChangeEvent): nil
function onClanTagChanged(callback) end